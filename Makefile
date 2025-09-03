.PHONY: build install clean deps

# Build the CLI tool
build:
	go build -o bin/cst cmd/cst/main.go

# Install globally
install: build
	sudo cp bin/cst /usr/local/bin/cst
	sudo chmod +x /usr/local/bin/cst

# Clean build files
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod init ctf-toolkit-cli 2>/dev/null || true
	go get github.com/spf13/cobra@latest
	pip3 install qrcode[pil] requests beautifulsoup4 lxml pillow exifread

# Create necessary directories
setup:
	mkdir -p uploads
	mkdir -p bin
	mkdir -p scripts/{crypto,misc,stego,binary,web,network,osint}

# Full setup
all: deps setup build

# Test the CLI
test: build
	./bin/cst --help