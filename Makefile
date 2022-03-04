run:
	go test -v ./... -tags=test
clean-test-cache:
	go clean -testcache
