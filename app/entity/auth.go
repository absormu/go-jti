package entity

type Auth struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type OAuthMessage struct {
	Authorization string `json:"-"`
	ErrorCode     int    `json:"-"`
	AccessToken   string `json:"access_token,omitempty" validate:"max=255,min=1"`
	TokenType     string `json:"token_type,omitempty" validate:"max=10,min=1"`
	Scope         string `json:"scope,omitempty" validate:"max=255,min=1"`
	ExpiresIn     int64  `json:"expires_in,omitempty"`
}

type ExtractToken struct {
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}

type ResetPassword struct {
	ID       string `json:"uid"`
	Password string `json:"password,omitempty"`
}
