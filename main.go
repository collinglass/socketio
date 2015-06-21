package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	yo := server.Of("/profile")

	log.Println(server.Name())
	log.Println(yo.Name())

	// namespaced
	yo.On("connection", func(so socketio.Socket) {
		log.Println("on connection /profile")
		so.On("authorization", func(msg string) {
			log.Println("authorized /profile " + msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect /profile")
		})
	})

	// namespaced
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.On("authorization", func(msg string) {
			log.Println("authorized " + msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
