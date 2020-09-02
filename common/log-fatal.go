package common

import "log"

//LogFatal - Checks if the given argument is an error
func LogFatal(err error) {
	if err != nil {
		log.Fatal("Fatal error: " + err.Error())
	}
}
