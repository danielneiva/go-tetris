package exceptions

import (
	"log"
)

func handle(err error) {
	if err != nil {
		log.fatal(err)
	}
}