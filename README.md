# CTF Starter Tool

## Overview
The CTF Starter Tool is a comprehensive command-line toolkit designed for Capture The Flag (CTF) competitions. Originally built as a microservice-oriented web application with Go (Golang) and Fiber framework, it now features a powerful CLI interface that integrates Python scripts for executing specific tools related to cryptography, steganography, binary analysis, web exploitation, network forensics, OSINT, and miscellaneous tasks.

## Features
- **CLI Interface**: Easy-to-use command-line interface for all CTF tools
- **REST API**: Exposes endpoints for each tool, allowing users to interact with the services
- **File Processing**: Supports file uploads and text inputs for processing
- **Secure Execution**: Utilizes sandboxed Python scripts for safe execution
- **Authentication**: Implements JWT for secure user authentication (API mode)
- **MongoDB Integration**: Stores user data and logs tool usage (API mode)

## Installation

### Quick Install
```bash
# Build and install the CLI tool
make all
make install
```

### Manual Installation
1. **Clone the Repository**: 
   ```bash
   git clone <repository-url>
   cd ctf-starter-tool
   ```

2. **Install Dependencies**: 
   ```bash
   # Go dependencies
   go mod tidy
   
   # Python dependencies
   pip3 install qrcode[pil] requests beautifulsoup4 lxml pillow exifread
   ```

3. **Build the CLI tool**:
   ```bash
   make build
   ```

4. **Install globally** (optional):
   ```bash
   make install
   ```

## CLI Usage Examples

### Cryptography Tools
```bash
# Caesar cipher encryption
cst crypto caesar "Hello World" 3

# Caesar cipher decryption
cst crypto caesar "Khoor Zruog" 3 --decrypt
```

### Miscellaneous Tools
```bash
# Generate QR code (auto-generated filename)
cst misc qr "Hello World"

# Generate QR code with custom filename
cst misc qr "https://example.com" --output myqr.png
```

### File Analysis
```bash
# Extract strings from binary
cst file strings /bin/ls
cst file strings suspicious.exe --min-length 8
```

### Steganography
```bash
# Extract EXIF data from image
cst stego exif photo.jpg
```

### Web Tools
```bash
# HTTP GET request
cst web http https://httpbin.org/get

# HTTP POST request with data
cst web http https://httpbin.org/post --method POST --data "test=data"
```

### Network Analysis
```bash
# Analyze PCAP file
cst network pcap capture.pcap

# Analyze PCAP with filter
cst network pcap traffic.pcapng --filter "http"
```

### OSINT Tools
```bash
# Basic username search
cst osint sherlock testuser

# Username search with custom timeout
cst osint sherlock john_doe --timeout 15
```

### Getting Help
```bash
# General help
cst --help

# Category-specific help
cst crypto --help

# Command-specific help
cst misc qr --help
```

## Project Structure
```
ctf-starter-tool/
├── cmd/
│   ├── cst/                 # CLI tool
│   │   └── main.go
│   └── server/              # API server
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   ├── crypto.go
│   │   │   ├── stego.go
│   │   │   ├── binary.go
│   │   │   ├── web.go
│   │   │   ├── network.go
│   │   │   ├── osint.go
│   │   │   └── misc.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   └── cors.go
│   │   └── routes/
│   │       └── routes.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── mongo.go
│   │   └── models/
│   │       ├── user.go
│   │       └── log.go
│   ├── services/
│   │   ├── executor.go
│   │   └── docker.go
│   └── utils/
│       ├── response.go
│       └── validation.go
├── scripts/
│   ├── crypto/
│   │   └── caesar.py
│   ├── stego/
│   │   └── exif_extractor.py
│   ├── binary/
│   │   └── strings_extractor.py
│   ├── web/
│   │   └── http_simulator.py
│   ├── network/
│   │   └── pcap_analyzer.py
│   ├── osint/
│   │   └── sherlock_search.py
│   └── misc/
│       └── qr_generator.py
├── uploads/                 # Generated files
├── bin/                     # Built binaries
├── docker/
│   ├── Dockerfile.tools
│   └── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## API Server Mode

If you prefer to use the REST API instead of CLI:

1. **Run the API Server**:
   ```bash
   go run cmd/server/main.go
   ```

2. **Access the API**: 
   - The API will be available at `http://localhost:8080`
   - Health check: `GET /api/v1/health`

3. **API Authentication**: 
   - Use authentication endpoints to obtain JWT tokens for protected routes
   - Register: `POST /api/v1/auth/register`
   - Login: `POST /api/v1/auth/login`

## Available Tools

| Category | Tool | Description |
|----------|------|-------------|
| **Crypto** | Caesar Cipher | Encrypt/decrypt text using Caesar cipher |
| **Misc** | QR Generator | Generate QR codes from text |
| **File** | Strings Extractor | Extract printable strings from binary files |
| **Stego** | EXIF Extractor | Extract metadata from images |
| **Web** | HTTP Simulator | Send HTTP requests and analyze responses |
| **Network** | PCAP Analyzer | Analyze network packet captures |
| **OSINT** | Sherlock Search | Search usernames across social media |

## Development

### Adding New Tools

1. Create Python script in appropriate `scripts/` subdirectory
2. Add CLI command in `cmd/cst/main.go`
3. Add API handler in `internal/api/handlers/`
4. Update routes in `internal/api/routes/routes.go`

### Building for Different Platforms
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/cst-linux cmd/cst/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/cst.exe cmd/cst/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o bin/cst-darwin cmd/cst/main.go
```

## Contributing
Contributions are welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Submit a pull request

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Support
For issues and questions:
- Open an issue on GitHub
- Check existing documentation
- Review the help commands: `cst --help`