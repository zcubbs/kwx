package awxclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	httpClient *resty.Client
}

func NewClient(baseURL, token string) *Client {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetHeader("Authorization", "Bearer "+token)
	client.SetHeader("Content-Type", "application/json")

	return &Client{
		httpClient: client,
	}
}

func (c *Client) ListenForEvents() {
	url := fmt.Sprintf("ws://%s/websocket/", c.httpClient.BaseURL)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		if messageType == websocket.TextMessage {
			// Handle the incoming message (parse it, act upon it, etc.)
			log.Println("Received:", string(p))
		}
	}
}
