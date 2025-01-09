include .env
# export

url := http://localhost:$(APP_PORT)/api
header := "Content-type:application/json"


appname := go-employee

run:
	@go run cmd/server/main.go

air:
	@air -c .air.windows.conf

build:
	@cd cmd/server; go build -o ../../bin/$(appname) && chmod +x ../../bin/$(appname)

exec:
	@./bin/$(appname)

startapp: build exec


post-emp:
	curl -H $(header) -d @pkg/v1/mock-data/create-employee.json -X POST $(url)/v1/employee
