package types

import "context"

type Payment struct {
	SessionID string `bson:"_id"`
	UserEmail string `bson:"user_email"`
	// "can be 'paid' or 'unpaid'"
	Status string `bson:"status"`
	Item   string `bson:"item"`
	// Price is measured in cents
	Price    uint `bson:"price"`
	Quantity uint `bson:"quantity"`
}

type Database interface {
	GetPayment(SessionID string) (Payment, error)
	CreatePayment(ctx context.Context, payment Payment) error
	MarkStatusAsPaid(SessionID string) error
}

type RedisDB interface {
	LockSessionID(ctx context.Context, sessionID string) (bool, error)
	DelSessionID(ctx context.Context, sessionID string) error
}
