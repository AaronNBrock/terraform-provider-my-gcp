build:
	go build -o terraform-provider-my-gcp
	terraform init


plan: build
	terraform plan

apply: build
	terraform apply -auto-approve