package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	wSocket "github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/sudak-91/wasm-test/internal/pkg/handler"
	mongorepository "github.com/sudak-91/wasm-test/internal/pkg/mongo_repository"
	"github.com/sudak-91/wasm-test/internal/pkg/updater"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	spa := spaHandler{staticPath: "../../template", indexPath: "index.html"}
	client := connectToMongoDb()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	DB := client.Database("blog")
	MongoRepository := mongorepository.NewMongoRepository(DB)
	Updater := updater.NewUpdater(MongoRepository.Users)
	wshandler := handler.NewWSHandler(Upgrader, Updater)
	router.Handle("/ws", wshandler)
	router.PathPrefix("/").Handler(spa)

	go func() {

		if err := http.ListenAndServe("0.0.0.0:8000", router); err != nil {
			panic(err.Error())
		}

	}()
	b := make(chan (bool))
	<-b

}

func connectToMongoDb() *mongo.Client {
	login := os.Getenv("MONGO_DB_USER")
	pass := os.Getenv("MONGO_DB_PASS")
	host := os.Getenv("MONGO_DB_HOST")
	ConnectionString := fmt.Sprintf("mongodb://%s:%s@%s/?maxPoolSize=20&w=majority", login, pass, host)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ConnectionString))
	if err != nil {
		log.Fatal("MongoDB connection error")
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Ping MongoDB error")
	}
	log.Println("Successfully connected and pinged")
	return client
}
