package exceptions

import (
	"log"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
