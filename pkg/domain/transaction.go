package domain

type TransactionID string

type Transaction struct {
	ID          TransactionID
	Account     Account
	Description *string
	Amount      int64
	Currency    string
}

func NewTransaction(
	id TransactionID,
	a Account,
	description *string,
	amount int64,
	currency string,
) *Transaction {
	return &Transaction{
		ID:          id,
		Account:     a,
		Description: description,
		Amount:      amount,
		Currency:    currency,
	}
}

func (tx *Transaction) UpdateAccountBalance() {
	tx.Account.BalanceValue += tx.Amount
}
