package domain

type AppConfig struct {
	SupportedAccountTypes []AccountType
}

func Config() *AppConfig {
	return &AppConfig{
		SupportedAccountTypes: supportedAccountTypes(),
	}
}

func supportedAccountTypes() []AccountType {
	return []AccountType{CHECKING, SAVINGS, CREDIT_CARD}
}
