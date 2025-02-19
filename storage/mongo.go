package storage

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	mongoDbName      = "db"
	mongoColPayments = "payments"
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

func (d *db) CreatePayment(ctx context.Context, payment types.Payment) error {
	coll := d.client.Database(mongoDbName).Collection(mongoColPayments)
	_, err := coll.InsertOne(ctx, payment)
	return err
}

func (d *db) MarkStatusAsPaid(SessionID string) error {
	return nil
}
