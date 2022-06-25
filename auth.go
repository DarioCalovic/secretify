package secretify

// AuthToken holds authentication token details with refresh token
type AuthToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
