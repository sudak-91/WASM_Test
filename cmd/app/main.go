package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	wSocket "github.com/gorilla/websocket"
)

var Upgrader = wSocket.Upgrader{}

type SpaHanler struct {
	StaticPath string
	IndexPath  string
}

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("../../template")))
	router.Handle("/post/", http.StripPrefix("/post/", http.FileServer(http.Dir("../../template"))))
	router.HandleFunc("/ws", getWebSocket)
	//spaHandler := SpaHanler{StaticPath: "../../template", IndexPath: "index.html"}
	//router.Handle("/", http.FileServer(http.Dir("../../template")))
	//router.PathPrefix("/").Handler(spaHandler)
	go func() {
		if err := http.ListenAndServe(":8000", router); err != nil {
			panic(err.Error())
		}
	}()

	b := make(chan bool)
	<-b

}

func (h SpaHanler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("inter to Serv")
	log.Println(r.URL.Host)
	//w.Header().Set("Content-Type", "application/wasm")
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.StaticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		log.Println("serve index.html")
		// file does not exist, serve index.html
		w.Header().Add("Content-Type", "application/wasm")
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("here")
	w.Header().Add("Content-Type", "application/wasm")
	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

func getWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("run")
	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader", err.Error())
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}
		log.Println(string(message[:]))
		err = c.WriteMessage(mt, []byte("right right right"))
		if err != nil {
			log.Println(err.Error())
			break
		}
	}

}
