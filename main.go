package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

var Conns = make(map[string][]*websocket.Conn)

type WSMsg struct {
    ID string `json:"id"`
    Content string `json:"content"`
    WhitesTurn bool `json:"whitesTurn"`
    Checkmate bool `json:"checkmate"`
}


func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func game(w http.ResponseWriter, r *http.Request) {
    unfilteredPathParts := strings.Split(r.URL.Path, "/")
    var pathParts []string
    for _, str := range unfilteredPathParts {
        if str != "" {
            pathParts = append(pathParts, str)
        }
    }
    if len(pathParts) == 2 {
        http.ServeFile(w, r, "game.html")
    } else if len(pathParts) == 3 {
        id := pathParts[1]
        board := NewBoard()
        board.ID = id
        Boards[id] = board
        http.ServeFile(w, r, "game.html")
    } else {
        panic("game: invalid URL")
    }
    
}

func handleWS(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        panic(err)
    }
    defer ws.Close()

    unfilteredPathParts := strings.Split(r.URL.Path, "/")
    var pathParts []string
    for _, str := range unfilteredPathParts {
        if str != "" {
            pathParts = append(pathParts, str)
        }
    }
    if len(pathParts) != 2 {
        panic("handleWS: invalid URL")
    }
    id := pathParts[1]

    Conns[id] = append(Conns[id], ws)

    broadcast(id)
    readLoop(id, ws)
}

func readLoop(id string, ws *websocket.Conn) {
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
            if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
                fmt.Println("read error: websocket closed")
                break
            }
			fmt.Println("read error:", err)
			continue
		}
        fmt.Println("message: ", msg)
        i1 := int(msg[0] - '0')
        j1 := int(msg[1] - '0')
        i2 := int(msg[2] - '0')
        j2 := int(msg[3] - '0')
		Move(i1, j1, i2, j2, Boards[id])
		broadcast(id)
	}
}

func broadcast(id string) {
    board := Boards[id]
	for _, ws := range Conns[id] {
        go func(ws *websocket.Conn) {
            if ws != nil {    
                customMsg := WSMsg {
                    ID: id,
                    Content: board.Content,
                    WhitesTurn: board.WhitesTurn,
                    Checkmate: board.Checkmate,
                }
                msgBytes, err := json.Marshal(customMsg)
                if err != nil {
                    fmt.Println("JSON marshaling error:", err)
                    return
                }
                if err := ws.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
                    fmt.Println("write error:", err)
                }
            }
        }(ws)
	}
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/game/", game)
    http.HandleFunc("/ws/", handleWS)

    Boards["1234"] = NewBoard()

    static := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", static))
    fmt.Println("Server at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}