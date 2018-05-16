//Package websockets ...
package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"myrepo/PufferBlock/server/action"

	"golang.org/x/net/websocket"
)

//Websockets returns ...
func Websockets() {

	http.Handle("/", websocket.Handler(echo))
	if err := http.ListenAndServe(":600false", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

//Request ...
type Request struct {
	Mode      string `json:"requestmode"`
	Name      string `json:"username"`
	Cmd       string `json:"command"`
	ArgName   string `json:"name"`
	ArgAmount int    `json:"amount"`
}

//Response ...
type Response struct {
	IfSuccessful bool   `json:"result"`
	ErrInfo      string `json:"resulterrorinfo"`
	/*	Result       struct {

		}
	*/
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

func (req *Request) doSelect() (Response, error) {

	if req.Mode == "" {

		return Response{false, "no mode"}, nil
	}

	if req.Mode == "initCC" {
		action.InitCC(req.Name)

		return Response{false, "no mode"}, nil
	}

	if req.Mode == "queryCC" {
		action.QueryCC()

		return Response{false, "no mode"}, nil
	}

	if req.Mode == "invokeCC" {
		action.InvokeCC()

		return Response{false, "no mode"}, nil
	}

	resp := &Response{false, ""}
	return *resp, nil
}
