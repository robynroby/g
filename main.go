package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {
	welcome := Welcome{"sale begins now", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/template.html"))

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sale := r.FormValue("sale"); sale != "" {
			welcome.Sale = sale

		}
		if err := template.Execute(w, welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}
