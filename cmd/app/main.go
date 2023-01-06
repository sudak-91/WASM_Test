package main

import (
	"fmt"
	"log"
	"net/http"

	wSocket "github.com/gorilla/websocket"
)

var Upgrader = wSocket.Upgrader{}

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("../../template")))
	//router.Handle("/post/", http.StripPrefix("/post/", http.FileServer(http.Dir("../../template"))))
	//router.HandleFunc("/ws", getWebSocket)
	go func() {
		if err := http.ListenAndServe(":8000", router); err != nil {
			panic(err.Error())
		}
	}()

	b := make(chan bool)
	<-b

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
