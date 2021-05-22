package logging

import "github.com/rs/zerolog/log"

func Error(err error, msg string, args ...interface{}) bool {
	if err == nil {
		return false
	}

	log.Debug().Msgf(msg, args...)
	return true
}
