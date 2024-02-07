package bridge

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"setu-engine/models"
)

func StartSocket(bridge models.Bridge, exit chan bool) {
	var firstChannel = make(chan models.Message, 100)
	var secondChannel = make(chan models.Message, 100)
	var firstExit = make(chan bool)
	var secondExit = make(chan bool)
	go ListenAndForward(bridge.FirstURL, firstChannel, bridge.FromChatID, secondChannel, bridge.SecondChatID, firstExit)
	go ListenAndForward(bridge.SecondURL, secondChannel, bridge.SecondChatID, firstChannel, bridge.FromChatID, secondExit)
	_ = <-exit
	firstExit <- true
	secondExit <- true
}

func ListenAndForward(host string, selfChannel chan models.Message, sourceChatId string, targetChannel chan models.Message, targetChatID string, exit chan bool) {
	URL := url.URL{Scheme: "ws", Host: host, Path: "connect",
		RawQuery: fmt.Sprintf("sub=%s&API_KEY=%s", sourceChatId, "API_KEY")}
	log.Println("ATTEMPTING CONNECTION AT:", URL.String())
	header := http.Header{}
	header.Add("API_KEY", "API_KEY")
	conn, _, err := websocket.DefaultDialer.Dial(URL.String(), header)
	if err != nil {
		log.Println("ERROR CONNECTING TO SERVER:", err)
		return
	}
	log.Println("CONNECTION OPENED AT:", URL.String())
	var counter = 0
	go func() {
		for msg := range selfChannel {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("ERROR WRITING MESSAGE TO CHANNEL", sourceChatId, err)
				if counter > 3 {
					break
				}
				conn, _, err = websocket.DefaultDialer.Dial(URL.String(), header)
				if err != nil {
					exit <- true
					break
				}
				counter++
				continue
			}
		}
	}()
	go func() {
		for {
			var msg models.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("ERROR READING MESSAGE FROM CHANNEL", sourceChatId, err)
				if counter > 3 {
					break
				}
				conn, _, err = websocket.DefaultDialer.Dial(URL.String(), header)
				if err != nil {
					exit <- true
					break
				}
				counter++
				continue
			}
			msg.ChatId = targetChatID
			targetChannel <- msg
		}
	}()
	_ = <-exit
	log.Println("CONNECTION CLOSED AT:", URL.String())
}