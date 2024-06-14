package models

type PlaidLinkToken struct {
	Expiration string
	LinkToken  string
	RequestID  uint64
}
