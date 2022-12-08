package constants

import "time"

const (
	RedisExpireDuration time.Duration = time.Duration(168) * time.Hour
	SuccessResponse                   = "Get Data Success"
	NotFoundResponse                  = "Get Data Not Found"
	FailedCreateJwt                   = "Failed Create JWT"
)
