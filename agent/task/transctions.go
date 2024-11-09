package task

import "time"

type Transactions struct {
	ReferenceID     string
	TransactionType string
	Amount          string
	TransactionDate time.Time
}

func (t Transactions) ToProto() string {
	return ""
}
