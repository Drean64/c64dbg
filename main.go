package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/Drean64/c64"
)

type Res = http.ResponseWriter
type Req = http.Request

func main() {

	com64 := c64.Make(c64.NTSC)
	// commands := make(chan interface{})
	// play := make(chan bool)

	// go com64.Run(commands, play)

	http.HandleFunc("/", func(res Res, req *Req) {
		fmt.Printf("[%s]\n", req.URL.Path)
		res.Header().Add("Cache-Control", "No-Store, Max-Age=0")

		if req.URL.Path == "/" {
			http.ServeFile(res, req, "./web/index.html")
		} else if req.URL.Path == "/index.js" {
			http.ServeFile(res, req, "./web/index.js")
		} else {
			res.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/guistart", func(w Res, r *Req) {
		//
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
	fileHeader := func(name string) textproto.MIMEHeader {
		return textproto.MIMEHeader{
			"Content-Type":        []string{"application/octet-stream"},
			"Content-Disposition": []string{fmt.Sprintf(`form-data; name="%s"; filename="%s"`, name, name)}}
	}

	multi := multipart.NewWriter(res)
	defer multi.Close()
	res.Header().Set("Content-type", multi.FormDataContentType())
	ram, _ := multi.CreatePart(fileHeader("RAM"))
	ram.Write(c64.RAM[:])
	io, _ := multi.CreatePart(fileHeader("IO"))
	io.Write(c64.IO[:])
}

//Content-Disposition: form-data; name="fieldName"
