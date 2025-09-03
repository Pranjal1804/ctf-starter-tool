.PHONY: build install clean deps banner

# Colors for output (Fixed)
GREEN=\033[0;32m
BLUE=\033[0;34m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m
BOLD=\033[1m

# Show banner with printf instead of echo
banner:
	@printf "$(BLUE)\n"
	@printf " ██████╗███████╗████████╗    ████████╗ ██████╗  ██████╗ ██╗     ██╗  ██╗██╗████████╗\n"
	@printf "██╔════╝██╔════╝╚══██╔══╝    ╚══██╔══╝██╔═══██╗██╔═══██╗██║     ██║ ██╔╝██║╚══██╔══╝\n"
	@printf "██║     ███████╗   ██║          ██║   ██║   ██║██║   ██║██║     █████╔╝ ██║   ██║   \n"
	@printf "██║     ╚════██║   ██║          ██║   ██║   ██║██║   ██║██║     ██╔═██╗ ██║   ██║   \n"
	@printf "╚██████╗███████║   ██║          ██║   ╚██████╔╝╚██████╔╝███████╗██║  ██╗██║   ██║   \n"
	@printf " ╚═════╝╚══════╝   ╚═╝          ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝   ╚═╝   \n"
	@printf "$(NC)\n"
	@printf "                        $(YELLOW)  CTF Starter Toolkit v1.0.0 $(NC)\n"
	@printf "                     $(GREEN) The Ultimate CTF Challenge Solver $(NC)\n"
	@printf "\n"

# Build the CLI tool
build: banner
	@printf "$(YELLOW) Building CTF Starter Tool...$(NC)\n"
	@mkdir -p bin
	@go build -o bin/cst cmd/cst/main.go
	@printf "$(GREEN) Build completed successfully!$(NC)\n"
	@printf "$(BLUE) Binary location: ./bin/cst$(NC)\n"
	@printf "$(YELLOW) Test it: ./bin/cst --help$(NC)\n"
	@printf "\n"

# Install globally
install: build
	@printf "$(YELLOW)📦 Installing CST globally...$(NC)\n"
	@sudo cp bin/cst /usr/local/bin/cst
	@sudo chmod +x /usr/local/bin/cst
	@printf "$(GREEN)🎉 Installation completed successfully!$(NC)\n"
	@printf "$(BLUE)📍 Installed to: /usr/local/bin/cst$(NC)\n"
	@printf "$(GREEN)✨ You can now use 'cst' from anywhere!$(NC)\n"
	@printf "\n"
	@printf "$(BOLD)🚀 Quick Start:$(NC)\n"
	@printf "   cst --help              $(BLUE)# Show help$(NC)\n"
	@printf "   cst version             $(BLUE)# Show version$(NC)\n"
	@printf "   cst crypto caesar \"Hi\" 3 $(BLUE)# Encrypt text$(NC)\n"
	@printf "   cst misc qr \"Hello\"      $(BLUE)# Generate QR$(NC)\n"
	@printf "\n"

# Clean build files
clean:
	@printf "$(YELLOW)🧹 Cleaning build files...$(NC)\n"
	@rm -rf bin/
	@printf "$(GREEN)✅ Clean completed!$(NC)\n"

# Install dependencies
deps: banner
	@printf "$(YELLOW)📥 Installing dependencies...$(NC)\n"
	@printf "$(BLUE)  → Installing Go dependencies...$(NC)\n"
	@go mod init ctf-toolkit-cli 2>/dev/null || true
	@go get github.com/spf13/cobra@latest
	@printf "$(BLUE)  → Installing Python dependencies...$(NC)\n"
	@pip3 install qrcode[pil] requests beautifulsoup4 lxml pillow exifread
	@printf "$(GREEN)✅ Dependencies installed successfully!$(NC)\n"

# Create necessary directories
setup:
	@printf "$(YELLOW)📁 Setting up directories...$(NC)\n"
	@mkdir -p uploads bin scripts/{crypto,misc,stego,binary,web,network,osint}
	@printf "$(GREEN)✅ Directory setup completed!$(NC)\n"

# Full setup
all: deps setup build
	@printf "$(GREEN)$(BOLD)🎊 CTF Starter Tool is ready to use! 🎊$(NC)\n"
	@printf "$(YELLOW)Next steps:$(NC)\n"
	@printf "  1. Run: $(BLUE)make install$(NC) to install globally\n"
	@printf "  2. Or use: $(BLUE)./bin/cst --help$(NC) to test locally\n"
	@printf "\n"

# Test the CLI
test: build
	@printf "$(YELLOW)🧪 Testing CTF Starter Tool...$(NC)\n"
	@./bin/cst --help
	@printf "$(GREEN)✅ Test completed successfully!$(NC)\n"

# Show usage help
help:
	@printf "$(BOLD)CTF Starter Tool - Makefile Commands:$(NC)\n"
	@printf "\n"
	@printf "$(YELLOW)Build Commands:$(NC)\n"
	@printf "  $(GREEN)make build$(NC)     - Build the CLI tool\n"
	@printf "  $(GREEN)make install$(NC)   - Build and install globally\n"
	@printf "  $(GREEN)make all$(NC)       - Complete setup (deps + build)\n"
	@printf "\n"
	@printf "$(YELLOW)Utility Commands:$(NC)\n"
	@printf "  $(GREEN)make deps$(NC)      - Install dependencies\n"
	@printf "  $(GREEN)make setup$(NC)     - Create directories\n"
	@printf "  $(GREEN)make clean$(NC)     - Clean build files\n"
	@printf "  $(GREEN)make test$(NC)      - Test the built tool\n"
	@printf "  $(GREEN)make banner$(NC)    - Show the banner\n"
	@printf "\n"