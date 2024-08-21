package error

import "log"

func CheckError(e error) {
	if e != nil {
		log.Fatal("error occurred reading csv file: ", e)
	}
}
