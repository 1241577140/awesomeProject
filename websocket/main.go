package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	if _, err := os.Stat("/sts"); !os.IsExist(err) {
		os.Mkdir("/sts", os.ModePerm)
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			sh := "echo %s >> /sts/sts.text"
			if string(msg) != "" {
				output, err := exec.Command("/bin/bash", "-c", fmt.Sprintf(sh, string(msg))).Output()
				if err != nil {
					fmt.Println("保存文件", string(output), "err:", err)
				}
				bytes, err := ioutil.ReadFile("/sts/sts.text")
				if err != nil {
					fmt.Println("读取文件err:", err)
				}
				// Write message back to browser
				if err = conn.WriteMessage(msgType, []byte("back:["+string(bytes)+"]")); err != nil {
					return
				}
			}
		}
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	})
	http.ListenAndServe(":80", nil)
}
