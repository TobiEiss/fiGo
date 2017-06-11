package fastconnect

import "errors"

var (
	// ErrRecoveryPassword cant find a "recovery_password" in JSON
	ErrRecoveryPassword = errors.New("figo-fastconnect: extract recovery password failed")
)

// FigoUser represent a user
type FigoUser struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

// BankCredentials represent the BankCredentials
type BankCredentials struct {
	BankCode    string   `json:"bankCode,omitempty" bson:"bankCode,omitempty"`
	Country     string   `json:"country,omitempty" bson:"country,omitempty"`
	Credentials []string `json:"credentials,omitempty" bson:"credentials,omitempty"`
}

type Task struct {
	AccountID            string `json:"account_id,omitempty"`
	IsEnded              bool   `json:"is_ended,omitempty"`
	IsErrpneous          bool   `json:"is_erroneous,omitempty"`
	IsWaitingForPin      bool   `json:"is_waiting_for_pin,omitempty"`
	IsWaitingForResponse bool   `json:"is_waiting_for_response,omitempty"`
	Message              string `json:"message,omitempty"`
}

type CatalogEntry struct {
	Name            string `json:"name"`
	BankCode        string `json:"bank_code"`
	Icon            string `json:"icon"`
	AdditionalIcons struct {
		Four8X48  string `json:"48x48"`
		Six0X60   string `json:"60x60"`
		Seven2X72 string `json:"72x72"`
		Eight4X84 string `json:"84x84"`
		Nine6X96  string `json:"96x96"`
		One20X120 string `json:"120x120"`
		One44X144 string `json:"144x144"`
		One92X192 string `json:"192x192"`
		Two56X256 string `json:"256x256"`
	} `json:"additional_icons"`
	Credentials []struct {
		Label  string `json:"label"`
		Masked bool   `json:"masked"`
	} `json:"credentials"`
	Advice string `json:"advice"`
}
