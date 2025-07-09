module github.com/Nikitarsis/goTokens/repository

go 1.22.2

require (
	github.com/Nikitarsis/goTokens/common v0.0.0-20250706230719-f00337b49d23
	github.com/lib/pq v1.10.9
)

replace github.com/Nikitarsis/goTokens/common => ../common
