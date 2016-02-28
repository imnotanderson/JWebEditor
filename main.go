package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//	"strings"
)

var htmBytes []byte
var jsBytes []byte

func main() {
	initHtm()
	startSvr()
}

func initHtm() {
	fileName := "editor.htm"
	file1, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(file1)
	if err != nil {
		return
	}
	htmBytes = b

	file1.Close()
	file1, _ = os.OpenFile("drag.js", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
	jsBytes, _ = ioutil.ReadAll(file1)
}

func startSvr() {
	http.HandleFunc("/dragjs", dragHandle)
	http.HandleFunc("/js", Handle)
	http.ListenAndServe(":4000", nil)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	info := r.FormValue("a")
	fmt.Println(info)

	w.Write([]byte(addJS + addDiv))
}

func dragHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	info := r.FormValue("a")
	fmt.Println(info)

	w.Write(jsBytes)
}

var addDiv string = "var div = document.createElement('div');" +
	"div.id = 'dlgTest';" +
	"div.className = 'dialog';" +
	"var div2 = document.createElement('div');" +
	"div.style.position='absolute';" +
	"div2.style.backgroundColor='#ccc';" +
	"div2.className = 'dialog-title';div2.innerHTML='Editor';var text = document.createElement('textarea');text.style = 'width:690px; height:600px;';div.appendChild(div2);div.appendChild(text);document.body.appendChild(div);"

var addJS string = "var s = document.createElement('script'); s.type = 'text/javascript'; s.src = 'http://127.0.0.1:4000/dragjs'; document.body.insertBefore(s, document.body.firstChild);"
