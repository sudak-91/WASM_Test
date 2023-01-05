package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("../../template")))
	router.Handle("/post/", http.StripPrefix("/post/", http.FileServer(http.Dir("../../template"))))
	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err.Error())
	}
	b := make(chan bool)
	<-b

}

func getIndex(w http.ResponseWriter, r *http.Request) {
}
