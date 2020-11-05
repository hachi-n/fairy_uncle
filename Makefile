# private tasks.
go-clean:
	rm -fr ./pkg

go-build:
	go build -o ./pkg/fairy_uncle ./cmd/fairy_uncle

# cmd interfaces.
build:
	$(MAKE) go-clean
	$(MAKE) go-build

clean:
	$(MAKE) go-clean

install:
	$(MAKE) build
	cp -pr ./pkg/fairy_uncle /usr/local/bin/

