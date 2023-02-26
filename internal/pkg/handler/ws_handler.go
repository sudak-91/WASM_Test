package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sudak-91/wasm-test/internal/pkg/updater"
	pubupdater "github.com/sudak-91/wasm-test/pkg/updater"
)

type WSHandler struct {
	Upgrader websocket.Upgrader
	Updater  updater.Updater
}

func NewWSHandler(upgareder websocket.Upgrader, updater updater.Updater) *WSHandler {
	return &WSHandler{
		Upgrader: upgareder,
		Updater:  updater,
	}
}

func (ws *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("run")
	ws.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader", err.Error())
		return
	}
	log.Println("UpgrADE")
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}
		log.Println(string(message[:]))
		var update pubupdater.Update
		err = json.Unmarshal(message, &update)
		if err != nil {
			log.Println("Unmwrshal error")
			continue
		}
		err = ws.Updater.Controler(update)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		err = c.WriteMessage(mt, []byte("right right right"))
		if err != nil {
			log.Println(err.Error())
		}
	}
}
