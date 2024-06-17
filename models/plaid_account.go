package models

import "gorm.io/gorm"

type PlaidAccount struct {
    gorm.Model
    Email      string `gorm:"uniqueIndex:idx_item_account_email;type:varchar(255)" json:"email"`
    AccessToken       string `gorm:"type:varchar(255)" json:"access_token"`
    ItemID            string `gorm:"uniqueIndex:idx_item_account_email;type:varchar(255)" json:"item_id"`
    AccountID         string `gorm:"uniqueIndex:idx_item_account_email;type:varchar(255)" json:"account_id"`
    BalanceAvailable  string `gorm:"type:varchar(255)" json:"balance_available"`
    BalanceCurrent    string `gorm:"type:varchar(255)" json:"balance_current"`
    BalanceISOCurrencyCode    string `gorm:"type:varchar(3)" json:"balance_iso_currency_code"`
    BalanceLimit      string `gorm:"type:varchar(255)" json:"balance_limit"`
    BalanceUnofficialCurrencyCode string `gorm:"type:varchar(3)" json:"balance_unofficial_currency_code"`
    Mask              string `gorm:"type:varchar(255)" json:"mask"`
    Name              string `gorm:"type:varchar(255)" json:"name"`
    OfficialName      string `gorm:"type:varchar(255)" json:"official_name"`
    Subtype           string `gorm:"type:varchar(255)" json:"subtype"`
    Type              string `gorm:"type:varchar(255)" json:"type"`
    InstitutionID     string `gorm:"type:varchar(255)" json:"institution_id"`
}


type Account struct {
    AccountID    string   `json:"account_id"`
    Balances     Balances `json:"balances"`
    Mask         string   `json:"mask"`
    Name         string   `json:"name"`
    OfficialName string   `json:"official_name"`
    Subtype      string   `json:"subtype"`
    Type         string   `json:"type"`
}

type Balances struct {
    Available              float64 `json:"available"`
    Current                float64 `json:"current"`
    IsoCurrencyCode        string  `json:"iso_currency_code"`
    Limit                  *float64 `json:"limit,omitempty"`
    UnofficialCurrencyCode *string `json:"unofficial_currency_code,omitempty"`
}

type Added struct {
    AccountID                string           `json:"account_id"`
    AccountOwner             *string          `json:"account_owner,omitempty"`
    Amount                   float64          `json:"amount"`
    IsoCurrencyCode          string           `json:"iso_currency_code"`
    UnofficialCurrencyCode   *string          `json:"unofficial_currency_code,omitempty"`
    Category                 []string         `json:"category"`
    CategoryID               string           `json:"category_id"`
    CheckNumber              *string          `json:"check_number,omitempty"`
    Counterparties          []Counterparty   `json:"counterparties"`
    Date                     string           `json:"date"`
    Datetime                 string           `json:"datetime"`
    AuthorizedDate           string           `json:"authorized_date"`
    AuthorizedDatetime       string           `json:"authorized_datetime"`
    Location                 Location         `json:"location"`
    Name                     string           `json:"name"`
    MerchantName             string           `json:"merchant_name"`
    MerchantEntityID         string           `json:"merchant_entity_id"`
    LogoURL                  string           `json:"logo_url"`
    Website                  string           `json:"website"`
    PaymentMeta              PaymentMeta      `json:"payment_meta"`
    PaymentChannel           string           `json:"payment_channel"`
    Pending                  bool             `json:"pending"`
    PendingTransactionID     *string          `json:"pending_transaction_id,omitempty"`
    PersonalFinanceCategory  FinanceCategory  `json:"personal_finance_category"`
    PersonalFinanceCategoryIconURL string     `json:"personal_finance_category_icon_url"`
    TransactionID            string           `json:"transaction_id"`
    TransactionCode          *string          `json:"transaction_code,omitempty"`
    TransactionType          string           `json:"transaction_type"`
}

type Counterparty struct {
    Name            string  `json:"name"`
    Type            string  `json:"type"`
    LogoURL         string  `json:"logo_url"`
    Website         string  `json:"website"`
    EntityID        string  `json:"entity_id"`
    ConfidenceLevel string  `json:"confidence_level"`
}

type Location struct {
    Address      *string  `json:"address,omitempty"`
    City         *string  `json:"city,omitempty"`
    Region       *string  `json:"region,omitempty"`
    PostalCode   *string  `json:"postal_code,omitempty"`
    Country      *string  `json:"country,omitempty"`
    Lat          *float64 `json:"lat,omitempty"`
    Lon          *float64 `json:"lon,omitempty"`
    StoreNumber  *string  `json:"store_number,omitempty"`
}

type PaymentMeta struct {
    ByOrderOf         *string `json:"by_order_of,omitempty"`
    Payee             *string `json:"payee,omitempty"`
    Payer             *string `json:"payer,omitempty"`
    PaymentMethod     *string `json:"payment_method,omitempty"`
    PaymentProcessor  *string `json:"payment_processor,omitempty"`
    PpdID             *string `json:"ppd_id,omitempty"`
    Reason            *string `json:"reason,omitempty"`
    ReferenceNumber   *string `json:"reference_number,omitempty"`
}

type FinanceCategory struct {
    Primary          string `json:"primary"`
    Detailed         string `json:"detailed"`
    ConfidenceLevel  string `json:"confidence_level"`
}

type Removed struct {
    AccountID     string `json:"account_id"`
    TransactionID string `json:"transaction_id"`
}

type ApiResponse struct {
    Accounts               []Account      `json:"accounts"`
    Added                  []Added        `json:"added"`
    Modified               []Added        `json:"modified"`
    Removed                []Removed      `json:"removed"`
    NextCursor             string         `json:"next_cursor"`
    HasMore                bool           `json:"has_more"`
    RequestID              string         `json:"request_id"`
    TransactionsUpdateStatus string       `json:"transactions_update_status"`
    Item                   Item           `json:"item"`
}