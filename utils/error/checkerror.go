package error

import "log"

func HandleError(e error) {
	if e != nil {
		log.Fatalln("Error occurred reading csv file: ", e)
	}
}
