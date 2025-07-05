module github.com/Nikitarsis/goTokens/tokenProducer

go 1.22.2

require (
	github.com/Nikitarsis/goTokens/common v0.0.0-20250705114537-39f6b537c866
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.6.0
)

replace github.com/Nikitarsis/goTokens/common => ../common
