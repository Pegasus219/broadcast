package common

import (
	"log"
)

func PanicRecover() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}
