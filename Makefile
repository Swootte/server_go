init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen generate
	
generate_:
	go run generation/generation.go

tidy: 
	go mod tidy
	
run:
	go run server.go

test:
	go test -v ./...


compile_contract:
	abigen --bin=./contracts/fcfa.bin --abi=./contracts/fcfa.abi --pkg=fcfa --out=./fcfa/fcfa.go

test_finance:
	go test -v ./finance/finance_test.go