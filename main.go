package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	rabbitmq "Robo/pkg/rabittMq"
	"Robo/service"
	"Robo/skimas"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"

)

func Publich(ch *amqp.Channel, whoisData skimas.WhoisData) error {
	body, err := json.Marshal(whoisData) 
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	Publich(ch, skimas.WhoisData{
		Domain:       "teste.com",
		Name:         "teste",
		Email:        "mail@gmail.com",
		Phone:        "123456789",
		Country:      "BR",
		Organization: "teste",
		CNPJ:         "123456789",
	})

	r := mux.NewRouter()

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Referrer-Policy", "no-referrer")
			next.ServeHTTP(w, r)
		})
	})

	ch, err = rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consumer(ch, out, "concurrentInformation")

	for i := 1; i <= 10; i++ {
		go func() {
			for {
				msg := <-out
				
				var body skimas.WhoisData
				err := json.Unmarshal(msg.Body, &body) // Converte um JSON
				if err != nil {
					panic(err)
				}
				
				service.SendEmail(body)
		
				msg.Ack(false)
			}
		}()
	}

	log.Fatal(http.ListenAndServe(":8084", corsHandler(r)))
}
