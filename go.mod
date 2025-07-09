module github.com/Nikitarsis/goTokens

go 1.22.2

require (
	github.com/Nikitarsis/goTokens/common v0.0.0-20250708183542-89b868deca8c
	github.com/Nikitarsis/goTokens/controller v0.0.0-20250707224304-8578dc0a6a5f
	github.com/Nikitarsis/goTokens/iptracer v0.0.0-00010101000000-000000000000
	github.com/Nikitarsis/goTokens/repository v0.0.0-20250707224304-8578dc0a6a5f
	github.com/Nikitarsis/goTokens/tokenProducer v0.0.0-20250707224304-8578dc0a6a5f
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace github.com/Nikitarsis/goTokens/common => ./common

replace github.com/Nikitarsis/goTokens/repository => ./repository

replace github.com/Nikitarsis/goTokens/tokenProducer => ./tokenProducer

replace github.com/Nikitarsis/goTokens/controller => ./controller

replace github.com/Nikitarsis/goTokens/iptracer => ./iptracer
