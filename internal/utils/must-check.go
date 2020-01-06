package utils

import "github.com/rs/zerolog/log"

func MustCheck(err error) {
	if err != nil {
		log.Fatal().Err(err)
	}
}

func MustCheckWithLog(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
	}
}
