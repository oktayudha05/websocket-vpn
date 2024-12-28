package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Izinkan semua koneksi (untuk testing, gunakan dengan hati-hati di produksi)
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade koneksi HTTP menjadi WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	log.Println("Client connected!")

	// Loop untuk menerima pesan dari client
	for {
		// Membaca pesan dari client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s\n", message)

		// Kirim pesan kembali ke client (echo)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	// Menangani route /wss untuk WebSocket
	http.HandleFunc("/wss", handleWebSocket)

	// Menjalankan server HTTP pada port 9000
	port := ":9000"
	fmt.Printf("WebSocket server started at ws://localhost%s/wss\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Server error:", err)
	}
}
