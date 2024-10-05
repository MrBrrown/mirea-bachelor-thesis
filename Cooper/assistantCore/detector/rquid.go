package detector

type Token struct {
	Token     string `json:"access_token"`
	ExpiresAt int64  `json:"expires_at"`
}

func (u *Token) IsActive(curTime int64) bool {
	return curTime < u.ExpiresAt
}
