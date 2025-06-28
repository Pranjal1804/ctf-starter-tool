# CTF Toolkit Backend

## Overview
The CTF Toolkit Backend is a microservice-oriented web application designed to provide various tools for Capture The Flag (CTF) competitions. The backend is built using Go (Golang) with the Fiber or Gin framework, and it integrates Python scripts for executing specific tools related to cryptography, steganography, binary analysis, web exploitation, network forensics, OSINT, and miscellaneous tasks.

## Features
- **REST API**: Exposes endpoints for each tool, allowing users to interact with the services.
- **File Uploads**: Supports uploading files or providing text inputs for processing.
- **Secure Execution**: Utilizes Docker for sandboxing Python scripts to ensure safe execution of untrusted inputs.
- **Authentication**: Implements JWT for secure user authentication.
- **MongoDB Integration**: Stores user data and logs tool usage.

## Project Structure
```
ctf-toolkit-backend
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   ├── auth.go
│   │   │   ├── crypto.go
│   │   │   ├── stego.go
│   │   │   ├── binary.go
│   │   │   ├── web.go
│   │   │   ├── network.go
│   │   │   ├── osint.go
│   │   │   └── misc.go
│   │   ├── middleware
│   │   │   ├── auth.go
│   │   │   └── cors.go
│   │   └── routes
│   │       └── routes.go
│   ├── config
│   │   └── config.go
│   ├── database
│   │   ├── mongo.go
│   │   └── models
│   │       ├── user.go
│   │       └── log.go
│   ├── services
│   │   ├── executor.go
│   │   └── docker.go
│   └── utils
│       ├── response.go
│       └── validation.go
├── scripts
│   ├── crypto
│   │   └── caesar.py
│   ├── stego
│   │   └── exif_extractor.py
│   ├── binary
│   │   └── strings_extractor.py
│   ├── web
│   │   └── http_simulator.py
│   ├── network
│   │   └── pcap_analyzer.py
│   ├── osint
│   │   └── sherlock_search.py
│   └── misc
│       └── qr_generator.py
├── uploads
├── docker
│   ├── Dockerfile.tools
│   └── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Setup Instructions
1. **Clone the Repository**: 
   ```
   git clone <repository-url>
   cd ctf-toolkit-backend
   ```

2. **Install Dependencies**: 
   - For Go:
     ```
     go mod tidy
     ```
   - For Python scripts, ensure you have the necessary libraries installed (e.g., `exiftool`, `requests`, etc.).

3. **Configure Database**: 
   - Update the MongoDB connection settings in `internal/config/config.go`.

4. **Run the Application**: 
   ```
   go run cmd/server/main.go
   ```

5. **Access the API**: 
   - The API will be available at `http://localhost:8080` (or the configured port).

## Usage Guidelines
- **Authentication**: Use the authentication endpoints to obtain a JWT token for accessing protected routes.
- **Tool Endpoints**: Each tool has its own endpoint. Refer to the API documentation for specific usage instructions.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.