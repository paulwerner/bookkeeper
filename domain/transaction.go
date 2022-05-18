package domain

type TransactionID string

func (self TransactionID) String() string {
	return string(self)
}

type Transaction struct {
	ID          TransactionID
	Amount      int64
	Currency    string
	Description *string
}

func NewTransaction(
	id TransactionID,
	amount int64,
	currency string,
	description *string,
) *Transaction {
	return &Transaction{
		ID:          id,
		Amount:      amount,
		Currency:    currency,
		Description: description,
	}
}
