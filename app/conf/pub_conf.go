package conf

import "time"

const (
	IS_RELEASE_ENV_KEY = "is_release"

	AUTHORIZATION_KEY = "Authorization"
	BEARER            = "Bearer "

	// ttl
	USER_INFO_TTL = 10 * time.Second
)
