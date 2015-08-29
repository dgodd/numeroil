package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/dgodd/numeroil/numeroil"
)

func main() {
	top := " <html> <head> <link href='https://fonts.googleapis.com/css?family=Poiret+One' rel='stylesheet' type='text/css'> <style>body{font-family: 'Poiret One', cursive;font-size:30px;background:linen;text-align:center;margin:100px 0;}</style> </head><body> <form action='/word'> <label for='q'>Word</label> <input name='q' autofocus> <input type='submit'> </form><br><br>"
	bottom := "</body></html>"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, top+bottom)
	})

	http.HandleFunc("/word", func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("q")
		big := numeroil.AddLetters(word)
		small := numeroil.Reduce(big)

		fmt.Fprintf(w, top)
		fmt.Fprintf(w, "Word: %q <br> Big: %d <br> small: %d", html.EscapeString(word), big, small)
		fmt.Fprintf(w, bottom)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listen on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
