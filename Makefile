#!/usr/bin/make -f

test:
	go fmt ./...
	go mod tidy
	go test -v -cover -timeout=1s -race ./...
	go test    -cover -timeout=1s -race ./...

doc:
	printf '# ' > README.md && \
		head -n 1 go.mod | sed 's/^module //' >> README.md && \
		echo >> README.md && \
		echo >> README.md && \
		go doc -all assert | sed 's/^/\t/' >> README.md && \
		echo >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md && \
		go doc -all suite  | sed 's/^/\t/' >> README.md
