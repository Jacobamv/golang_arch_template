package entities

// Ping
//
// swagger:response ping
type swaggerPing struct {
	Data struct {
		// example: Успешно
		Message string `json:"message"`
		// example: Clean arch pong!
		Payload string `json:"payload"`
	}
}
