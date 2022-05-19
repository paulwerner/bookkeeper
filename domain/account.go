package domain

type AccountID string

type AccountType string

const (
	CHECKING    AccountType = "CHECKING"
	SAVINGS     AccountType = "SAVINGS"
	CREDIT_CARD AccountType = "CREDIT_CARD"
)

type Account struct {
	ID              AccountID
	User            User
	Name            string
	Description     *string
	Type            AccountType
	BalanceValue    int64
	BalanceCurrency string
}

func NewAccount(
	id AccountID,
	u User,
	name string,
	description *string,
	accountType AccountType,
	balanceValue int64,
	balanceCurrency string,
) *Account {
	return &Account{
		ID:              id,
		User:            u,
		Name:            name,
		Description:     description,
		Type:            accountType,
		BalanceValue:    balanceValue,
		BalanceCurrency: balanceCurrency,
	}
}
