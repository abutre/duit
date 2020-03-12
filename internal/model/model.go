package model

import (
	"encoding/json"

	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v3"
)

// Config is content of configuration file
type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

// User is container for user's data
type User struct {
	ID         int64           `db:"id"          json:"id"`
	Username   string          `db:"username"    json:"username"`
	Name       string          `db:"name"        json:"name"`
	Password   string          `db:"password"    json:"password,omitempty"`
	Permission json.RawMessage `db:"permission"  json:"permission"`
}

// Account is container for financial account
type Account struct {
	ID            int64           `db:"id"             json:"id"`
	Name          string          `db:"name"           json:"name"`
	InitialAmount decimal.Decimal `db:"initial_amount" json:"initialAmount"`

	// Additional info fields
	Total decimal.Decimal `db:"total" json:"total"`
}

// Entry is container for book entries
type Entry struct {
	ID                int64           `db:"id"                  json:"id"`
	AccountID         int64           `db:"account_id"          json:"accountId"`
	AffectedAccountID null.Int        `db:"affected_account_id" json:"affectedAccountId"`
	EntryType         int             `db:"entry_type"          json:"entryType"`
	Description       null.String     `db:"description"         json:"description"`
	Amount            decimal.Decimal `db:"amount"              json:"amount"`
	EntryDate         string          `db:"entry_date"          json:"entryDate"`
}