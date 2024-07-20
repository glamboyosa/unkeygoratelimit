fmt:
	gofmt -s -w .	
run-examples:
	go run ./examples/basic/basic_usage.go && go run ./examples/options/with_options.go && go run ./examples/timeout/with_timeout.go