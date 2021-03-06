package main

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/Drean64/c64"
)

type Res = http.ResponseWriter
type Req = http.Request
type Headers = textproto.MIMEHeader

func main() {

	com64 := c64.Make(c64.NTSC)
	// commands := make(chan interface{})
	// play := make(chan bool)

	// go com64.Run(commands, play)

	http.HandleFunc("/", func(res Res, req *Req) {
		res.Header().Add("Cache-Control", "No-Store, Max-Age=0")

		if req.URL.Path == "/" {
			http.ServeFile(res, req, "../web/index.html")
		} else if req.URL.Path == "/index.js" {
			http.ServeFile(res, req, "../web/index.js")
		} else {
			res.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/state", func(res Res, _ *Req) {
		res.Header().Add("Cache-Control", "No-Store, Max-Age=0")
		sendState(com64, res)
	})

	port := ":6464"
	fmt.Printf("go64 debugger listening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func sendState(c64 *c64.C64, res Res) {
	multi := multipart.NewWriter(res)
	defer multi.Close()

	res.Header().Set("Content-type", multi.FormDataContentType())
	cpu, _ := multi.CreatePart(Headers{"Content-Type": []string{"application/json"}, "Content-Disposition": []string{`form-data; name="CPU"`}})
	cpuJson, err := json.Marshal(c64.CPU)
	if err != nil {
		res.Header().Set("Content-type", "text/plain")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	cpu.Write(cpuJson)
	ram, _ := multi.CreatePart(Headers{"Content-Type": []string{"application/octet-stream"}, "Content-Disposition": []string{`form-data; name="RAM"; filename="RAM"`}})
	ram.Write(c64.RAM[:])
	io, _ := multi.CreatePart(Headers{"Content-Type": []string{"application/octet-stream"}, "Content-Disposition": []string{`form-data; name="IO"; filename="IO"`}})
	io.Write(c64.IO[:])
}
