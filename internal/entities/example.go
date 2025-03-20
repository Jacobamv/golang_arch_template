package entities

// OtpSession — structure is designed to form a payload,
// which will later be unloaded into a redis
type OtpSession struct {
	CreatedAt     int64  `json:"created_at"`
	SendAttempts  int    `json:"send_attempts"`
	CheckAttempts int    `json:"check_attempts"`
	Message       string `json:"-"`
	Phone         string `json:"phone"`
	Code          string `json:"code" validate:"required,len=4"`
	Password      string `json:"password" validate:"required,min=8,max=100"`
}
