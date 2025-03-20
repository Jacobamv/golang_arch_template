package utils

import "fmt"

func CibtTimeFormater(cibtTime string) (resTime string) {
	if cibtTime == "" {
		return
	}

	day := cibtTime[:len(cibtTime)-6]
	month := cibtTime[len(cibtTime)-6 : len(cibtTime)-4]
	year := cibtTime[len(cibtTime)-4:]

	resTime = fmt.Sprintf("%v.%v.%v", day, month, year)
	return
}
