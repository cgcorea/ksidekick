package server

import (
	"encoding/json"
	"log"
	"net/http"

	kannel2 "github.com/cgcorea/ksidekick/kannel"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Router *chi.Mux
	client *kannel2.Client
}

type Message struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Text     string `json:"text"`
	Priority int    `json:"priority"`
}

func NewServer() *Server {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	c := kannel2.NewClient("localhost", 4103, "sender", "sender")

	return &Server{Router: r, client: c}
}

func (s *Server) sendMessage(w http.ResponseWriter, r *http.Request) {
	m := &Message{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println("Error decoding message: ", err)
		return
	}

	log.Printf("sendMessage: %#v", m)

	req, err := s.client.NewRequest(m.From, m.To, m.Text)
	if err != nil {
		log.Println("Error creating Kannel request:", err)
	}
	response, err := s.client.Send(req)
	log.Printf("response: %#v", response)

	if err != nil {
		log.Println("Error sending message:", err)
	} else {
		renderJSON(w, response)
	}
}

func renderJSON(w http.ResponseWriter, response interface{}) {
	data, _ := json.Marshal(response)
	w.Header().Set("content-type", "aplication/json")
	_, err := w.Write(data)
	if err != nil {
		log.Println("Response write error:", err)
	}
}

func (s *Server) SetRoutes() {
	s.Router.Post("/message", s.sendMessage)
}
