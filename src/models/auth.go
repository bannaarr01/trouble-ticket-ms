package models

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthJwtPayload struct {
	AccessToken      string `json:"access_token"`
	IDToken          string `json:"id_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type SignUpDTO struct {
	Auth
	Email      string   `json:"email"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	RealmRoles []string `json:"RealmRoles"`
}
