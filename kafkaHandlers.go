package handlers

import (
	"net/http"
	"fmt"
	"github.com/jonsteck/kafka"
)

func AdHandlerKafka(w http.ResponseWriter, r *http.Request, s *kafka.Server) {
	topic := "ad"
	value := r.URL.Path[4:]
	s.KafkaSend(w, topic, value)
	//fmt.Fprintf(w, "Ad: %s", r.URL.Path[4:])
}

func ClickHandlerKafka(w http.ResponseWriter, r *http.Request, s *kafka.Server) {
	topic := "click" 
	value := r.URL.Path[7:]
	s.KafkaSend(w, topic, value)
	//fmt.Fprintf(w, "click: %s", r.URL.Path[7:])
}

func WinHandlerKafka(w http.ResponseWriter, r *http.Request, s *kafka.Server) {
        topic := "win"
	value := r.URL.Path[5:]
	s.KafkaSend(w, topic, value)
	//fmt.Fprintf(w, "win: %s", r.URL.Path[5:])
}

func OtherHandlerKafka(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Failed to store your data:\n")
	http.NotFound(w, r)
}
