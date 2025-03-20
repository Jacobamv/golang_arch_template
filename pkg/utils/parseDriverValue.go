package utils

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"time"

	"github.com/godror/godror"
)

func ParseDriverValue(val driver.Value, typeVal string) (res interface{}, err error) {
	var ok bool
	var number godror.Number

	switch typeVal {
	case "int64":

		number, ok = val.(godror.Number)
		if !ok {
			return 0, errors.New("error on parsing " + typeVal)
		}

		res, err = strconv.ParseInt(string(number), 10, 64)

		if err != nil {
			return 0, errors.New("error on parsing " + typeVal + err.Error())
		}
	case "string":
		res, ok = val.(string)
		if !ok {
			return "", errors.New("error on parsing " + typeVal)
		}

	case "float64":

		number, ok = val.(godror.Number)
		if !ok {
			return 0, errors.New("error on parsing " + typeVal)
		}

		res, err = strconv.ParseFloat(string(number), 64)

		if err != nil {
			return
		}

	case "time":
		res, ok = val.(time.Time)

		if !ok {
			return time.Time{}, errors.New("error on parsing " + typeVal)
		}
	}
	return
}
