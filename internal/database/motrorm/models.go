// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package motrorm

import (
	"database/sql"
	"time"
)

type Account struct {
	ID            string       `json:"id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
	Number        int64        `json:"number"`
	Sequence      int64        `json:"sequence"`
	Address       string       `json:"address"`
	PublicKey     string       `json:"public_key"`
	ChainID       string       `json:"chain_id"`
	Controller    string       `json:"controller"`
	IsSubsidiary  bool         `json:"is_subsidiary"`
	IsValidator   bool         `json:"is_validator"`
	IsDelegator   bool         `json:"is_delegator"`
	IsAccountable bool         `json:"is_accountable"`
}

type Asset struct {
	ID          string         `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
	Name        string         `json:"name"`
	Symbol      string         `json:"symbol"`
	Decimals    int64          `json:"decimals"`
	ChainID     string         `json:"chain_id"`
	Channel     string         `json:"channel"`
	AssetType   string         `json:"asset_type"`
	CoingeckoID sql.NullString `json:"coingecko_id"`
}

type Credential struct {
	ID                      string       `json:"id"`
	CreatedAt               time.Time    `json:"created_at"`
	UpdatedAt               time.Time    `json:"updated_at"`
	DeletedAt               sql.NullTime `json:"deleted_at"`
	Handle                  string       `json:"handle"`
	CredentialID            string       `json:"credential_id"`
	AuthenticatorAttachment string       `json:"authenticator_attachment"`
	Origin                  string       `json:"origin"`
	Type                    string       `json:"type"`
	Transports              string       `json:"transports"`
}

type Profile struct {
	ID        string       `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Address   string       `json:"address"`
	Handle    string       `json:"handle"`
	Origin    string       `json:"origin"`
	Name      string       `json:"name"`
}

type Session struct {
	ID             string       `json:"id"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	DeletedAt      sql.NullTime `json:"deleted_at"`
	BrowserName    string       `json:"browser_name"`
	BrowserVersion string       `json:"browser_version"`
	ClientIpaddr   string       `json:"client_ipaddr"`
	Platform       string       `json:"platform"`
	IsDesktop      bool         `json:"is_desktop"`
	IsMobile       bool         `json:"is_mobile"`
	IsTablet       bool         `json:"is_tablet"`
	IsTv           bool         `json:"is_tv"`
	IsBot          bool         `json:"is_bot"`
	Challenge      string       `json:"challenge"`
	IsHumanFirst   bool         `json:"is_human_first"`
	IsHumanLast    bool         `json:"is_human_last"`
	ProfileID      int64        `json:"profile_id"`
}

type Vault struct {
	ID          string       `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Handle      string       `json:"handle"`
	Origin      string       `json:"origin"`
	Address     string       `json:"address"`
	Cid         string       `json:"cid"`
	Config      string       `json:"config"`
	SessionID   string       `json:"session_id"`
	RedirectUri string       `json:"redirect_uri"`
}
