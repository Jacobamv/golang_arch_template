package entities

// SwaggerInn ...
type SwaggerInn struct {
	// ИНН клиента
	//
	// example: 123456789
	// required: true
	Inn string `json:"inn"`
}

type swaggerOtpCode struct {
	// Одноразовый код
	//
	// required: true
	// max length: 4
	// default: 1234
	Code string `json:"code"`
}

type swaggerPassword struct {
	// Пароль
	//
	// required: true
	// min length: 8
	// max length: 100
	// default: !qwerty123
	Password string `json:"password"`
}

type swaggerPhone struct {
	// Телефонный номер
	//
	// required: true
	// default: 992000331341
	// pattern: 992\d{9}$
	Phone string `json:"phone"`
}
