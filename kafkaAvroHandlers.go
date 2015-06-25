package handlers

import (
	"net/http"
	"github.com/jonsteck/myAvro"
	"github.com/jonsteck/kafka"
)

//wrap handler functions to make required handler format: func(http.ResponseWriter, *http.Request)
func MakeHandler(fn func(http.ResponseWriter, *http.Request, *kafka.Server), s *kafka.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s)
	}
}

//ad handler - calls avro to encode value
func AdHandlerKafkaAvro(w http.ResponseWriter, r *http.Request, s *kafka.Server) {
	topic := "ad"
	value := myAvro.MyEncode(r.URL.Path[4:])
	s.KafkaSendAvro(w, topic, value)
}

//click handler - calls avro to encode value
func ClickHandlerKafkaAvro(w http.ResponseWriter, r *http.Request, s *kafka.Server) {
	topic := "click" 
	value := myAvro.MyEncode(r.URL.Path[7:])
	s.KafkaSendAvro(w, topic, value)
}

//win handler - calls avro to encode value
func WinHandlerKafkaAvro(w http.ResponseWriter, r *http.Request, s *kafka.Server) {
	topic := "win"
	value := myAvro.MyEncode(r.URL.Path[5:])
	s.KafkaSendAvro(w, topic, value)
}
