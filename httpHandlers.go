package handlers

import (
	"net/http"
	"fmt"
)

func AdHandler(w http.ResponseWriter, r *http.Request) {
	//topic = "ad"
	//value = r.URL.Path[4:]
	//return
	fmt.Fprintf(w, "Ad: %s", r.URL.Path[4:])
}

func ClickHandler(w http.ResponseWriter, r *http.Request) {
	//topic = "click" 
	//value = r.URL.Path[7:]
	//return
	fmt.Fprintf(w, "click: %s", r.URL.Path[7:])
}

func WinHandler(w http.ResponseWriter, r *http.Request) {
        //topic = "win"
	//value = r.URL.Path[5:]
	//return
	fmt.Fprintf(w, "win: %s", r.URL.Path[5:])
}

func OtherHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
