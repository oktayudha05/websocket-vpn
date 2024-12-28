package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Izinkan semua koneksi
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade koneksi HTTP menjadi WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected!")

	// Loop untuk menerima pesan dari client
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		// Kirim pesan kembali ke client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/wss", handleWebSocket)

	port := ":9000"
	fmt.Printf("WebSocket server started at ws://localhost%s/wss\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
