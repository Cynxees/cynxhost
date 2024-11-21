run:
	swag init --dir ./server --output ./www/docs
	go run ./server/ 