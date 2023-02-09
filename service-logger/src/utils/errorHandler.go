package utils

import "log"

func HandleErrors(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}
