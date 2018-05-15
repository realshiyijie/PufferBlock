//Package websockets ...
package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

//Websockets returns ...
func Websockets() {
	//	action.Shell()
	http.Handle("/", websocket.Handler(echo))

	if err := http.ListenAndServe(":6001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

//Response ...
type Response struct {
}

func echo(ws *websocket.Conn) {
	for {
		requstAsBytes := []byte{}
		if err := websocket.JSON.Receive(ws, &requstAsBytes); err != nil {
			fmt.Println("Can't receive")
			break
		}

		requst := &Response{}
		json.Unmarshal(requstAsBytes, requst)
		fmt.Println("Received back from client: " + string(requstAsBytes))

		response, err := doSelect(requst)
		if err != nil {
			fmt.Println("Can't understand requst")
			break
		}
		responseAsBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Can't understand response")
		}
		fmt.Println("Sending to client: " + string(responseAsBytes))
		if err := websocket.JSON.Send(ws, responseAsBytes); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func doSelect(req *Response) (Response, error) {

	return *req, nil
}
