package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dgodd/numeroil/numeroil"
)

type Data struct {
	Word  string
	Big   int
	Small int
}

func printData(w io.Writer, data []Data) {
	fmt.Fprintf(w, "<center><table>")
	for i := len(data) - 1; i >= 0; i-- {
		d := data[i]
		fmt.Fprintf(w, "<tr><td align=center>%s</td><td align=right>%d</td><td align=right>%d</td></tr>", d.Word, d.Big, d.Small)
	}
	fmt.Fprintf(w, "</table></center>")
}

func main() {
	top := " <html> <head> <link href='https://fonts.googleapis.com/css?family=Poiret+One' rel='stylesheet' type='text/css'> <style>body{font-family: 'Poiret One', cursive;font-size:30px;background:linen;text-align:center;margin:100px 0;}</style> </head><body> <a href='/clear' style='float:right;'>Clear</a> <form action='/word'> <label for='q'>Word</label> <input name='q' autofocus> <input type='submit'> </form><br><br>"
	bottom := "</body></html>"

	var data []Data

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, top)
		printData(w, data)
		fmt.Fprintf(w, bottom)
	})

	http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		data = data[:0]
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.HandleFunc("/word", func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("q")
		big := numeroil.AddLetters(word)
		small := numeroil.Reduce(big)
		data = append(data, Data{Word: word, Big: big, Small: small})

		fmt.Fprintf(w, top)
		fmt.Fprintf(w, "%q <br> Big: %d <br> small: %d", html.EscapeString(word), big, small)
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
