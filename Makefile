test:
	go test ./cotisationCalculator

test_adapter:
	go test -v ./cotisationCalculator/adapter/ -tags=adapter
