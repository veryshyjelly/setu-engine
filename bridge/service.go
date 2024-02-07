package bridge

import (
	"log"
	"setu-engine/models"
)

type Service interface {
	Subscribe() chan models.Bridge
	Unsubscribe() chan models.Bridge
	Run()
}

type service struct {
	Subscribers map[models.Bridge]chan bool
	subscribe   chan models.Bridge
	unsubscribe chan models.Bridge
}

func NewService() Service {
	return &service{
		Subscribers: make(map[models.Bridge]chan bool),
		subscribe:   make(chan models.Bridge, 100),
		unsubscribe: make(chan models.Bridge, 100),
	}
}

func (s *service) Subscribe() chan models.Bridge {
	return s.subscribe
}

func (s *service) Unsubscribe() chan models.Bridge {
	return s.unsubscribe
}

func (s *service) Run() {
	go s.subscriber()
	go s.unSubscriber()
	log.Println("STARTED BRIDGE SERVICE")
}

func (s *service) subscriber() {
	for c := range s.subscribe {
		s.Subscribers[c] = make(chan bool)
		log.Println("NEW BRIDGE SUBSCRIPTION:", c)
		StartSocket(c, s.Subscribers[c])
	}
}

func (s *service) unSubscriber() {
	for c := range s.unsubscribe {
		s.Subscribers[c] <- true
		delete(s.Subscribers, c)
	}
}