format:
	@gofmt -l -s -w . && go mod tidy

release:
	@bash ./scripts/build.bash parrot ${PWD}/cmd/parrot/main.go