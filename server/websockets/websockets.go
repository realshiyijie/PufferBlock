//Package websockets ...
package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

//Request ...
type Request struct {
	Mode     string `json:"requestmode"`
	Name     string `json:"username"`
	Function string `json:"command"`
	OpName   string `json:"operatname"`
	OpAmount int    `json:"operatamount"`
}

//Websockets returns ...
func Websockets() {

	http.Handle("/", websocket.Handler(echo))
	if err := http.ListenAndServe(":6001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func echo(ws *websocket.Conn) {
	for {
		requestAsBytes := []byte{}
		if err := websocket.JSON.Receive(ws, &requestAsBytes); err != nil {
			fmt.Println("Can't receive")
			break
		}

		request := &Request{}
		json.Unmarshal(requestAsBytes, request)
		fmt.Println("Received back from client: " + string(requestAsBytes))

		response, err := request.doSelect()
		if err != nil {
			fmt.Println("Can't understand requst")
			break
		}
		responseAsBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Can't understand response")
			break
		}
		fmt.Println("Sending to client: " + string(responseAsBytes))
		if err := websocket.JSON.Send(ws, responseAsBytes); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
