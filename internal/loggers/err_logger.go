package loggers

import (
	"fmt"
	"log"
	"net/http"
)

// TODO: не используется
func ErrLog(err error, w http.ResponseWriter, op string) {
	if err != nil {
		log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", err, op))
		w.Write([]byte(err.Error()))
		return
	}
}

func ErrLogS(err string, w http.ResponseWriter, op string) {
	log.Println(fmt.Errorf("ERROR: %s. ERROR PATH: %s", err, op))
	w.Write([]byte(err))
}
