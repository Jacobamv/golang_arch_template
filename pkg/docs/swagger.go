// Package response Application Bridge
//
// ## Сервис Конвеера заявок Хумо для работы с внешними платформами
//
// Данный сервис был создан чтобы сделать Application не зависимым от других платформ, единственное что можно будет заменять это его Bridge (от англ. Мост)
//
// Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// BasePath: /api
//
// securityDefinitions:
//	Bearer:
//	  type: apiKey
//	  name: Authorization
//	  in: header
//
// swagger:meta
package response

// swagger:response success
type swaggerSuccessResponse struct {
	// in: body
	Data struct {
		// example: success
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response conflict
type swaggerConflictResponse struct {
	// in: body
	Data struct {
		// example: conflict
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response badRequest
type swaggerBadRequestResponse struct {
	// in: body
	Data struct {
		// example: bad request
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response accessDenied
type swaggerAccessDeniedResponse struct {
	// in: body
	Data struct {
		// example: access is denied
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response notFound
type swaggerNotFoundResponse struct {
	// in: body
	Data struct {
		// example: notFound
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response retryLimitExceeded
type swaggerRetryLimitExceededResponse struct {
	// in: body
	Data struct {
		// example: retry limit exceeded
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response notImplementation
type swaggerNotImplementationResponse struct {
	// in: body
	Data struct {
		// example: not implementation
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response unauthorized
type swaggerUnauthorizedResponse struct {
	// in: body
	Data struct {
		// example: unauthorized
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response somethingWentWrong
type swaggerSomethingWentWrongResponse struct {
	// in: body
	Data struct {
		// example: something went wrong
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response internalServer
type swaggerInternalServerResponse struct {
	// in: body
	Data struct {
		// example: internal server error
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}
