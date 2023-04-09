package main

const (
	secret = "12312321IDFOIDUFAP123"

	ErrorRedisCannotSetAccessKey  = "cannot set new access_token key"
	ErrorRedisCannotSetRefreshKey = "cannot set new refresh_token key"
	ErrorRedisTokenIsInvalid      = "token is invalid"
	ErrorRedisCannotDeleteKey     = "cannot delete key from redis"
)
