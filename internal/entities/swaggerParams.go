package entities

// Authorization
//
// Строение роутов для авторизации

// swagger:parameters authLogin
type swaggerAuthLogin struct {
	// in:body
	Body struct {
		swaggerOtpCode
		swaggerPassword
		swaggerPhone
	}
}
