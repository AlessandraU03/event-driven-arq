package adapters

import (
	"encoding/json"
	"eventdriven/src/internal/domain/entities"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("amqp://ale:ale05@54.156.170.232/")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error al abrir el canal: %w", err)
	}

	return &RabbitMQAdapter{conn: conn, channel: channel}, nil
}

func (r *RabbitMQAdapter) NotifyPedidoCreation(pedido *entities.Pedido) error {
	// Declarar la cola
	queue, err := r.channel.QueueDeclare(
		"pedidos", 
		true,     
		false,     
		false,     
		false,     
		nil,       
	)
	if err != nil {
		return fmt.Errorf("error al declarar la cola: %w", err)
	}

	pedidoJSON, err := json.Marshal(pedido)
	if err != nil {
		return fmt.Errorf("error al convertir el pedido a JSON: %w", err)
	}

	// Publicar el mensaje en RabbitMQ
	err = r.channel.Publish(
		"",        
		queue.Name, 
		false,     
		false,    
		amqp.Publishing{
			ContentType: "application/json",
			Body:        pedidoJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
	}

	log.Printf("📤 Pedido enviado a RabbitMQ: %s", pedidoJSON)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
