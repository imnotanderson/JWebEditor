package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var htmBytes []byte

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
}

func startSvr() {
	http.HandleFunc("/js", Handle)
	http.ListenAndServe(":4000", nil)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	info := r.FormValue("a")
	fmt.Println(info)
	w.Write([]byte(fmt.Sprintf(js, string(htmBytes))))
}

var js string = " var div = document.createElement('div');div.innerHTML = \"%v\";document.body.appendChild(div);"
