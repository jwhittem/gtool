APP=gtool

default: deps install

install:
	@echo "[ Install binary ]"
	@go install

build: deps
	@echo "[ Build binary ]"
	@go build -o ${APP}

build-all: deps
	@echo "[ Build binaries for common platforms ]"
	@mkdir -p bin/mac bin/windows bin/linux
	@GOOS=windows GOARCH=amd64 go build -o bin/widows/${APP}
	@GOOS=darwin GOARCH=amd64 go build -o bin/mac/${APP}
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/${APP}

clean:
	rm -rf ${APP} bin/

deps:
	@echo "[ Get dependencies ]"
	@if [ ! -f "Gopkg.toml" ]; then \
		dep init; \
	 	fi
	@dep ensure

test:
	@echo "[ Running all tests ]"
	go test -v ./...
