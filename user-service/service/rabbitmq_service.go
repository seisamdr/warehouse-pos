package service

import (
	"context"
	"encoding/json"
	"fmt"
	"micro-warehouse/user-service/config"

	"github.com/gofiber/fiber/v2/log"
	"github.com/streadway/amqp"
)

type EmailPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
}

type RabbitMQServiceInterface interface {
	PublishEmail(ctx context.Context, payload EmailPayload) error
	Close() error
}

type rabbitMQService struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	config config.Config
}

// Close implements [RabbitMQServiceInterface].
func (r *rabbitMQService) Close() error {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	return nil
}

// PublishEmail implements [RabbitMQServiceInterface].
func (r *rabbitMQService) PublishEmail(ctx context.Context, payload EmailPayload) error {
	// Convert body to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("failed to marshal email payload: %v", err)
		return err
	}

	// Declare queue if not exist
	queue, err := r.ch.QueueDeclare(
		"email_queue", // Queue name
		true,          // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare email queue: %v", err)
	}

	// Publish ke email queue langsung (tanpa exchange)
	err = r.ch.Publish(
		"",           // Exchange (empty for default)
		queue.Name,   // Routing key (queue name)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish email message: %v", err)
	}

	return nil
}

func NewRabbitMQService(config config.Config) (RabbitMQServiceInterface, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port))
	if err != nil {
		log.Errorf("[RabbitMQ Service] NewRabbitMQService - 1: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[RabbitMQ Service] NewRabbitMQService - 2: %v", err)
		return nil, err
	}

	return &rabbitMQService{
		conn:   conn,
		ch:     ch,
		config: config,
	}, nil
}
