test:
	bash env.sh
	go test ./client
	go test ./face