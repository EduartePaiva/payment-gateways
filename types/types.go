package types

type Payment struct {
	UserEmail string
	Status    string
	SessionID string
}

type Database interface {
	GetPayment(SessionID string) (Payment, error)
	CreatePayment(SessionID string, UserEmail string) error
	MarkStatusAsPaid(SessionID string) error
}
