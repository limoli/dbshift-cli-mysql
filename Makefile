
.PHONY: build
build:
	go build -o dbshift-mysql .

.PHONY: test
test: build
	./dbshift-mysql status
	./dbshift-mysql upgrade
	./dbshift-mysql status
	./dbshift-mysql downgrade
	./dbshift-mysql status
	./dbshift-mysql upgrade 20190926154408
	./dbshift-mysql status
	./dbshift-mysql downgrade 20190926154408
	./dbshift-mysql status