package models

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
)

// Claims is a struct to hold the JWT claims
type Claims struct {
	Acr               string         `json:"acr"`
	AllowedOrigins    []string       `json:"allowed-origins"`
	Aud               string         `json:"aud"`
	Azp               string         `json:"azp"`
	Email             string         `json:"email"`
	EmailVerified     bool           `json:"email_verified"`
	Exp               float64        `json:"exp"`
	FamilyName        string         `json:"family_name"`
	GivenName         string         `json:"given_name"`
	Iat               float64        `json:"iat"`
	Iss               string         `json:"iss"`
	Jti               string         `json:"jti"`
	Name              string         `json:"name"`
	PreferredUsername string         `json:"preferred_username"`
	RealmAccess       RealmAccess    `json:"realm_access"`
	ResourceAccess    ResourceAccess `json:"resource_access"`
	Scope             string         `json:"scope"`
	Sid               string         `json:"sid"`
	Sub               string         `json:"sub"`
	Typ               string         `json:"typ"`
}

// RealmAccess struct to hold the realm access claims
type RealmAccess struct {
	Roles []string `json:"roles"`
}

type ResourceAccess struct {
	Account Account `json:"account"`
}

type Account struct {
	Roles []string `json:"roles"`
}

// GetClaims returns a Claims object from a jwt.MapClaims
func GetClaims(claims jwt.MapClaims) (*Claims, error) {
	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	var c Claims
	err = json.Unmarshal(jsonClaims, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
