module github.com/satoshi1975/smartChat/services/auth-service

go 1.13

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.10.9
	github.com/satoshi1975/smartChat/common v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.19.0
	golang.org/x/crypto v0.21.0
)

replace github.com/satoshi1975/smartChat/common => ../../common
