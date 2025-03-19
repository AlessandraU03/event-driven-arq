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
	queue   amqp.Queue
}

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	// Conectar a RabbitMQ
	conn, err := amqp.Dial("amqp://ale:ale05@54.156.170.232/")
	if err != nil {
		return nil, fmt.Errorf("❌ Error al conectar con RabbitMQ: %w", err)
	}

	// Crear un canal
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("❌ Error al abrir el canal: %w", err)
	}

	// Declarar la cola (se declara solo una vez aquí, no en cada envío)
	queue, err := channel.QueueDeclare(
		"pedidos", 
		true,     
		false,    
		false,    
		false,    
		nil,      
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("❌ Error al declarar la cola: %w", err)
	}

	log.Println("✅ Conectado a RabbitMQ y cola 'pedidos' declarada.")

	return &RabbitMQAdapter{conn: conn, channel: channel, queue: queue}, nil
}

// Implementa la interfaz NotificationPort
func (r *RabbitMQAdapter) NotifyPedidoCreation(pedido *entities.Pedido) error {
	if r.channel == nil {
		return fmt.Errorf("❌ Canal RabbitMQ no inicializado")
	}

	pedidoJSON, err := json.Marshal(pedido)
	if err != nil {
		return fmt.Errorf("❌ Error al convertir el pedido a JSON: %w", err)
	}

	// Publicar el mensaje en la cola ya declarada
	err = r.channel.Publish(
		"",         
		r.queue.Name, 
		false,      
		false,      
		amqp.Publishing{
			ContentType: "application/json",
			Body:        pedidoJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("❌ Error al publicar el mensaje en RabbitMQ: %w", err)
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
	log.Println("🔌 Conexión a RabbitMQ cerrada.")
}
