package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Drean64/c64"
)

func main() {

	com64 := c64.Make(c64.NTSC)
	events := make(chan interface{})
	commands := make(chan interface{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/index.html")
	})

	go com64.Run(events, commands)
	go func() {
		time.Sleep(1 * time.Second)
		close(commands)
	}()

	port := ":6464"
	fmt.Printf("go64 debugger listening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
