package storage

import (
	"context"

	"github.com/EduartePaiva/payment-gateways/types"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (d *db) GetPayment(ctx context.Context, SessionID string) (types.Payment, error) {
	var payment types.Payment
	if err := d.
		client.
		Database(mongoDbName).
		Collection(mongoColPayments).
		FindOne(ctx, bson.D{{Key: "_id", Value: SessionID}}).
		Decode(&payment); err != nil {
		return types.Payment{}, err
	}
	return payment, nil
}

func (d *db) CreatePayment(ctx context.Context, payment types.Payment) error {
	// ensures that the status is unpaid when creating a payment
	payment.Status = "unpaid"
	coll := d.client.Database(mongoDbName).Collection(mongoColPayments)
	_, err := coll.InsertOne(ctx, payment)
	return err
}

func (d *db) MarkStatusAsPaid(ctx context.Context, SessionID string) error {
	coll := d.client.Database(mongoDbName).Collection(mongoColPayments)
	update := bson.D{
		{Key: "$set",
			Value: bson.D{{Key: "status", Value: "paid"}},
		},
	}
	_, err := coll.UpdateByID(ctx, SessionID, update)
	return err
}
