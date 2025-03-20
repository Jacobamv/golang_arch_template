package response

import "errors"

// Errors used in the project. Sorted alphabetically
var (
	// Default errors
	ErrAccessDenied        = errors.New("Отказано в доступе")
	ErrBadRequest          = errors.New("Неверный запрос")
	ErrBindingClientToCard = errors.New("Ошибка при привязки клиента с картой")
	ErrCardNotFound        = errors.New("Карта не существует")
	ErrCheckPhoneInOrzu    = errors.New("Необработанное исключение")
	ErrDataNotFound        = errors.New("Не найдено")

	ErrDublicateClient              = errors.New("Ошибка дублирования клиента")
	ErrEmployee                     = errors.New("Данный клиент является сотрудником банка")
	ErrInternalServer               = errors.New("Внутренняя ошибка сервера")
	ErrLimitExceeded                = errors.New("Превышен лимит попыток. Повторите позже")
	ErrInvalidData                  = errors.New("Ошибка валидации данных")
	ErrMoreThanOneClientFound       = errors.New("Ошибка в БД, найден больше чем один клиент по данному ИНН, обратитесь в ближайший офис ХУМО")
	ErrNoContent                    = errors.New("Нет контента")
	ErrNotFoundInCibt               = errors.New("Данных клиента нет в КИБТе")
	ErrNotFoundInHumoDB             = errors.New("Клиент будет добавлен в базу после того как пройдет всю процедуру заявки")
	ErrNotImplementation            = errors.New("Не реализовано")
	ErrPhoneNumberExists            = errors.New("По этому номеру телефону уже зарегистрирован ОРЗУ")
	ErrSomethingWentWrong           = errors.New("Что-то пошло не так")
	ErrSuccess                      = errors.New("Успешно")
	ErrUnauthorized                 = errors.New("Не авторизованный пользователь")
	ErrWaitingApplicationFromBasket = errors.New("В ожидании получения новой заявки")
)

// Error statuses
var errorCode = map[string]int{
	// Default errors
	ErrSuccess.Error():                      200,
	ErrNoContent.Error():                    201,
	ErrBadRequest.Error():                   400,
	ErrBindingClientToCard.Error():          400,
	ErrWaitingApplicationFromBasket.Error(): 400,
	ErrUnauthorized.Error():                 401,
	ErrAccessDenied.Error():                 403,
	ErrDataNotFound.Error():                 404,
	ErrCardNotFound.Error():                 404,
	ErrNotFoundInCibt.Error():               404,
	ErrNotFoundInHumoDB.Error():             404,
	ErrCheckPhoneInOrzu.Error():             406,
	ErrMoreThanOneClientFound.Error():       409,
	ErrPhoneNumberExists.Error():            409,
	ErrLimitExceeded.Error():                429,
	ErrSomethingWentWrong.Error():           500,
	ErrInternalServer.Error():               500,
	ErrNotImplementation.Error():            501,
}
