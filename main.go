// Save as quote-service/main.go
package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket" // You'll need to install this: go get github.com/gorilla/websocket
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for now (CORS for websockets)
		return true 
	},
}

// Mock price data
var mockPrices = map[string]float64{
	"AAPL": 170.00,
	"MSFT": 410.00,
	"TSLA": 200.00,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Start a goroutine to continuously send mock stock updates
	for {
		// Simulate a price change
		mockPrices["AAPL"] += 0.05
		mockPrices["MSFT"] -= 0.02
		
		message := map[string]interface{}{
			"event": "price_update",
			"data":  mockPrices,
		}
		
		err = ws.WriteJSON(message)
		if err != nil {
			log.Printf("error: %v", err)
			break // Close connection if writing fails
		}
		
		time.Sleep(1 * time.Second) // Send update every second
	}
}

func main() {
	log.Println("Quote Service started on :8080")
	// WebSocket endpoint
	http.HandleFunc("/ws/quotes", handleConnections) 

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
