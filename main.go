package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/dgodd/numeroil/numeroil"
)

type Data struct {
	Word  string
	Big   int
	Small int
}
type Datum []Data

func (d Datum) Len() int { return len(d) }
func (d Datum) Less(i, j int) bool {
	return d[i].Small < d[j].Small || (d[i].Small == d[j].Small && d[i].Big < d[j].Big)
}
func (d Datum) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func printData(w io.Writer, data Datum) {
	fmt.Fprintf(w, "<center><table>")
	for _, d := range data {
		fmt.Fprintf(w, "<tr><td align=center>%s</td><td align=right>%d</td><td align=right>%d</td><td><a href='/delete?q=%s' onclick='return confirm(\"Delete? %s\");'>x</a></tr>", d.Word, d.Big, d.Small, url.QueryEscape(d.Word), url.QueryEscape(d.Word))
	}
	fmt.Fprintf(w, "</table></center>")
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readData(db *bolt.DB) Datum {
	var data Datum
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Words"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			a := strings.Split(string(v), ",")
			big, e1 := strconv.Atoi(a[0])
			small, e2 := strconv.Atoi(a[1])
			if e1 == nil && e2 == nil {
				data = append(data, Data{Word: string(k), Big: big, Small: small})
			}
		}

		return nil
	})
	sort.Stable(data)
	return data
}

func add(db *bolt.DB, d Data) Datum {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Words"))
		err := b.Put([]byte(d.Word), []byte(fmt.Sprintf("%d,%d", d.Big, d.Small)))
		return err
	})

	return readData(db)
}

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("Words"))
		tx.CreateBucket([]byte("ClearedWords"))
		return nil
	})
	defer db.Close()

	top := " <html> <head> <link href='https://fonts.googleapis.com/css?family=Poiret+One' rel='stylesheet' type='text/css'> <style>body{font-family: 'Poiret One', cursive;font-size:30px;background:linen;text-align:center;margin:100px 0;}</style> </head><body> <form action='/word'> <label for='q'>Word</label> <input name='q' autofocus> <input type='submit'> </form><br><br>"
	bottom := "</body></html>"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, top)
		data := readData(db)
		printData(w, data)
		fmt.Fprintf(w, bottom)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		word := []byte(r.FormValue("q"))
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Words"))
			dead := tx.Bucket([]byte("ClearedWords"))
			data := b.Get(word)
			b.Delete(word)
			err := dead.Put(word, data)
			return err
		})
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.HandleFunc("/word", func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("q")
		big := numeroil.AddLetters(word)
		small := numeroil.Reduce(big)
		data := add(db, Data{Word: word, Big: big, Small: small})

		fmt.Fprintf(w, top)
		fmt.Fprintf(w, "%q <br> Big: %d <br> small: %d <br><br>", html.EscapeString(word), big, small)
		printData(w, data)
		fmt.Fprintf(w, bottom)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listen on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
