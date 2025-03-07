VERSION := 0.1.0
BINDIR := bin
BINNAME := templater

.PHONY: build install clean release

build:
	go build -o $(BINDIR)/$(BINNAME) ./cmd/review

install:
	go install ./cmd/review

clean:
	rm -rf $(BINDIR)

build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/$(BINNAME)-$(VERSION)-linux-amd64 ./cmd/review
	GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/$(BINNAME)-$(VERSION)-darwin-amd64 ./cmd/review
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/$(BINNAME)-$(VERSION)-windows-amd64.exe ./cmd/review

release: build-all
	tar -czf $(BINDIR)/$(BINNAME)-$(VERSION)-linux-amd64.tar.gz -C $(BINDIR) $(BINNAME)-)-$(VERSION)-linux-amd64
	tar -czf $(BINDIR)/$(BINNAME)-$(VERSION)-darwin-amd64.tar.gz -C $(BINDIR) $(BINNAME)-)-$(VERSION)-darwin-amd64
	zip -j $(BINDIR)/$(BINNAME)-$(VERSION)-windows-amd64.zip $(BINDIR)/$(BINNAME)-)-$(VERSION)-windows-amd64.exe
