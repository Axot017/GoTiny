run:
	go run cmd/gotiny/main.go

test:
	go test ./...

dev_infra:
	cd deployments/infrastructure/dev && terraform init && terraform apply

spec:
	swagger generate spec -o ./api/swagger.yaml --scan-models
