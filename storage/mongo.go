package storage

import (
	"github.com/EduartePaiva/payment-gateways/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type db struct {
	client *mongo.Client
}

func NewDatabase(client *mongo.Client) *db {
	return &db{client: client}
}

func (d *db) GetPayment(SessionID string) (types.Payment, error) {
	return types.Payment{}, nil
}
func (d *db) CreatePayment(SessionID string, UserEmail string) error {
	return nil
}
func (d *db) MarkStatusAsPaid(SessionID string) error {
	return nil
}
