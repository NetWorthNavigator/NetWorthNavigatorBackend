package models


type Institution struct {
    InstitutionID string `json:"institution_id"`
    Name          string `json:"name"`
    Logo          string `json:"logo"`
    PrimaryColor  string `json:"primary_color"`
    URL           string `json:"url"`
}

type InstitutionDetail struct {
    CountryCodes   []string            `json:"country_codes"`
    InstitutionID  string              `json:"institution_id"`
    Name           string              `json:"name"`
    Products       []string            `json:"products"`
    RoutingNumbers []string            `json:"routing_numbers"`
    DTCNumbers     []string            `json:"dtc_numbers"`
    OAuth          bool                `json:"oauth"`
    Status         InstitutionStatuses `json:"status"`
}

type InstitutionStatuses struct {
    ItemLogins          StatusDetail `json:"item_logins"`
    TransactionsUpdates StatusDetail `json:"transactions_updates"`
    Auth                StatusDetail `json:"auth"`
    Identity            StatusDetail `json:"identity"`
    Investments         StatusDetail `json:"investments"`
    InvestmentsUpdates  StatusDetail `json:"investments_updates"`
    LiabilitiesUpdates  StatusDetail `json:"liabilities_updates"`
}

type StatusDetail struct {
    Status            string            `json:"status"`
    LastStatusChange  string            `json:"last_status_change"`
    Breakdown         StatusBreakdown   `json:"breakdown"`
}

type StatusBreakdown struct {
    Success          float64 `json:"success"`
    ErrorPlaid       float64 `json:"error_plaid"`
    ErrorInstitution float64 `json:"error_institution"`
    RefreshInterval  string  `json:"refresh_interval,omitempty"`
}

type InstitutionsResponse struct {
    Institutions []InstitutionDetail `json:"institutions"`
    RequestID    string              `json:"request_id"`
    Total        int                 `json:"total"`
}