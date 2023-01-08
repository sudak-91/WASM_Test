package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	wSocket "github.com/gorilla/websocket"
)

var Upgrader = wSocket.Upgrader{}

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)

	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(path)
	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)
	fmt.Println(path)
	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	router := mux.NewRouter()
	//router.Handle("/", http.FileServer(http.Dir("../../template")))
	//router.Handle("/post/", http.StripPrefix("/post/", http.FileServer(http.Dir("../../template"))))
	spa := spaHandler{staticPath: "../../template", indexPath: "index.html"}
	router.HandleFunc("/ws", getWebSocket)
	router.PathPrefix("/").Handler(spa)
	go func() {
		if err := http.ListenAndServe("0.0.0.0:8000", router); err != nil {
			panic(err.Error())
		}
	}()

	b := make(chan bool)
	<-b

}

func getWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("run")
	Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
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
