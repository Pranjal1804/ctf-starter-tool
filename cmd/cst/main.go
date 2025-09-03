package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const version = "v1.0.0"
const banner = `
 ██████╗███████╗████████╗    ████████╗ ██████╗  ██████╗ ██╗     ██╗  ██╗██╗████████╗
██╔════╝██╔════╝╚══██╔══╝    ╚══██╔══╝██╔═══██╗██╔═══██╗██║     ██║ ██╔╝██║╚══██╔══╝
██║     ███████╗   ██║          ██║   ██║   ██║██║   ██║██║     █████╔╝ ██║   ██║   
██║     ╚════██║   ██║          ██║   ██║   ██║██║   ██║██║     ██╔═██╗ ██║   ██║   
╚██████╗███████║   ██║          ██║   ╚██████╔╝╚██████╔╝███████╗██║  ██╗██║   ██║   
 ╚═════╝╚══════╝   ╚═╝          ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝   ╚═╝   
																					  
						  CTF Starter Toolkit %s 
					  The Ultimate CTF Challenge Solver 
						
	┌─────────────────────────────────────────────────────────────────┐
	│   Crypto  │    Stego  │   Files  │   Web  │   Network │
	│    OSINT  │    Misc   │   Pwn    │   Rev  │    Fast   │
	└─────────────────────────────────────────────────────────────────┘
	
					Made for CTF enthusiasts
					Type 'cst --help' to get started
`

var rootCmd = &cobra.Command{
	Use:     "cst",
	Version: version,
	Short:   "CTF Starter Tool - Complete CTF toolkit in CLI",
	Long:    fmt.Sprintf(banner, version),
	Example: `  cst crypto caesar "Hello World" 3
  cst misc qr "Hello World"
  cst file strings binary.exe
  cst stego exif image.jpg`,
}

// Add a version command with banner
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and banner",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(banner, version)
		fmt.Printf("\n🔧 Build Info:\n")
		fmt.Printf("   • Version: %s\n", version)
		fmt.Printf("   • Built with: Go + Python\n")
		fmt.Printf("   • Platform: Cross-platform\n")
		fmt.Printf("   • License: MIT\n\n")
	},
}

// Crypto commands
var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Cryptography tools",
	Long:  "Various cryptography utilities for encryption, decryption, and analysis",
}

var caesarCmd = &cobra.Command{
	Use:   "caesar [text] [key]",
	Short: "Caesar cipher encryption/decryption",
	Long:  "Encrypt or decrypt text using Caesar cipher with specified shift key",
	Args:  cobra.ExactArgs(2),
	Example: `  cst crypto caesar "Hello World" 3
  cst crypto caesar "Khoor Zruog" 3 --decrypt`,
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		key, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: Key must be a number\n")
			return
		}

		decrypt, _ := cmd.Flags().GetBool("decrypt")
		mode := "encrypt"
		if decrypt {
			mode = "decrypt"
		}

		output, err := exec.Command("python3", "scripts/crypto/caesar.py", text, strconv.Itoa(key), mode).Output()
		if err != nil {
			fmt.Printf("Error executing Caesar cipher: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing result: %v\n", err)
			return
		}

		printResult("Caesar Cipher", result)
	},
}

// Misc commands
var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Miscellaneous tools",
	Long:  "Various utility tools including QR code generation, encoding/decoding, etc.",
}

var qrCmd = &cobra.Command{
	Use:   "qr [text]",
	Short: "Generate QR code",
	Long:  "Generate QR code from text and save as PNG image",
	Args:  cobra.ExactArgs(1),
	Example: `  cst misc qr "Hello World"
  cst misc qr "https://example.com" --output custom.png`,
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		outputFile, _ := cmd.Flags().GetString("output")

		// Generate filename if not provided
		if outputFile == "" {
			timestamp := time.Now().Format("20060102150405")
			outputFile = fmt.Sprintf("qr_%s.png", timestamp)
		}

		// Ensure uploads directory exists
		os.MkdirAll("uploads", 0755)
		fullPath := filepath.Join("uploads", outputFile)

		output, err := exec.Command("python3", "scripts/misc/qr_generator.py", text, fullPath).Output()
		if err != nil {
			fmt.Printf("Error generating QR code: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing result: %v\n", err)
			return
		}

		result["file_path"] = fullPath
		printResult("QR Code Generator", result)
	},
}

// File analysis commands
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File analysis tools",
	Long:  "Various file analysis tools for binary analysis, metadata extraction, etc.",
}

var stringsCmd = &cobra.Command{
	Use:   "strings [file]",
	Short: "Extract strings from binary file",
	Long:  "Extract printable strings from binary files for analysis",
	Args:  cobra.ExactArgs(1),
	Example: `  cst file strings binary.exe
  cst file strings suspicious_file --min-length 10`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		minLength, _ := cmd.Flags().GetInt("min-length")

		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Error: File '%s' does not exist\n", filePath)
			return
		}

		var cmdArgs []string
		cmdArgs = append(cmdArgs, "scripts/binary/strings_extractor.py", filePath)
		if minLength > 0 {
			cmdArgs = append(cmdArgs, "--min-length", strconv.Itoa(minLength))
		}

		output, err := exec.Command("python3", cmdArgs...).Output()
		if err != nil {
			fmt.Printf("Error extracting strings: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing result: %v\n", err)
			return
		}

		printResult("Strings Extraction", result)
	},
}

// Steganography commands
var stegoCmd = &cobra.Command{
	Use:   "stego",
	Short: "Steganography tools",
	Long:  "Steganography and metadata analysis tools",
}

var exifCmd = &cobra.Command{
	Use:   "exif [image]",
	Short: "Extract EXIF data from image",
	Long:  "Extract EXIF metadata from image files",
	Args:  cobra.ExactArgs(1),
	Example: `  cst stego exif photo.jpg
  cst stego exif image.png --verbose`,
	Run: func(cmd *cobra.Command, args []string) {
		imagePath := args[0]

		// Check if file exists
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			fmt.Printf("Error: Image file '%s' does not exist\n", imagePath)
			return
		}

		output, err := exec.Command("python3", "scripts/stego/exif_extractor.py", imagePath).Output()
		if err != nil {
			fmt.Printf("Error extracting EXIF data: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing result: %v\n", err)
			return
		}

		printResult("EXIF Data Extraction", result)
	},
}

// Web tools
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Web exploitation tools",
	Long:  "Web exploitation and analysis tools",
}

var httpCmd = &cobra.Command{
	Use:   "http [url]",
	Short: "HTTP request simulator",
	Long:  "Send HTTP requests and analyze responses",
	Args:  cobra.ExactArgs(1),
	Example: `  cst web http https://httpbin.org/get
  cst web http https://httpbin.org/post --method POST --data "key=value"`,
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		method, _ := cmd.Flags().GetString("method")
		data, _ := cmd.Flags().GetString("data")

		var cmdArgs []string
		cmdArgs = append(cmdArgs, "scripts/web/http_simulator.py", url, method)
		if data != "" {
			cmdArgs = append(cmdArgs, data)
		}

		output, err := exec.Command("python3", cmdArgs...).Output()
		if err != nil {
			fmt.Printf("Error executing HTTP request: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing result: %v\n", err)
			return
		}

		printResult("HTTP Request", result)
	},
}

// Network tools
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Network analysis tools",
	Long:  "Network forensics and analysis tools",
}

var pcapCmd = &cobra.Command{
	Use:   "pcap [file]",
	Short: "Analyze PCAP file",
	Long:  "Analyze network packet capture files",
	Args:  cobra.ExactArgs(1),
	Example: `  cst network pcap capture.pcap
  cst network pcap traffic.pcapng --filter "http"`,
	Run: func(cmd *cobra.Command, args []string) {
		pcapFile := args[0]
		filter, _ := cmd.Flags().GetString("filter")

		// Check if file exists
		if _, err := os.Stat(pcapFile); os.IsNotExist(err) {
			fmt.Printf("Error: PCAP file '%s' does not exist\n", pcapFile)
			return
		}

		var cmdArgs []string
		cmdArgs = append(cmdArgs, "scripts/network/pcap_analyzer.py", pcapFile)
		if filter != "" {
			cmdArgs = append(cmdArgs, "--filter", filter)
		}

		output, err := exec.Command("python3", cmdArgs...).Output()
		if err != nil {
			fmt.Printf("Error analyzing PCAP: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing result: %v\n", err)
			return
		}

		printResult("PCAP Analysis", result)
	},
}

// OSINT tools
var osintCmd = &cobra.Command{
	Use:   "osint",
	Short: "OSINT tools",
	Long:  "Open Source Intelligence gathering tools",
}

var sherlockCmd = &cobra.Command{
	Use:   "sherlock [username]",
	Short: "Search username across social media",
	Long:  "Search for username across multiple social media platforms",
	Args:  cobra.ExactArgs(1),
	Example: `  cst osint sherlock john_doe
  cst osint sherlock testuser --timeout 10`,
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		timeout, _ := cmd.Flags().GetInt("timeout")

		var cmdArgs []string
		cmdArgs = append(cmdArgs, "scripts/osint/sherlock_search.py", username)
		if timeout != 5 { // Only add timeout if different from default
			cmdArgs = append(cmdArgs, "--timeout", strconv.Itoa(timeout))
		}

		fmt.Printf("Searching for username '%s' across social media platforms...\n", username)
		output, err := exec.Command("python3", cmdArgs...).Output()
		if err != nil {
			fmt.Printf("Error executing Sherlock search: %v\n", err)
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			// If JSON parsing fails, just print raw output
			fmt.Println(string(output))
			return
		}

		printResult("Sherlock Search", result)
	},
}

// Helper function to print results nicely
func printResult(tool string, result map[string]interface{}) {
	fmt.Printf("\n=== %s ===\n", tool)
	
	if success, ok := result["success"].(bool); ok && !success {
		if errorMsg, ok := result["error"].(string); ok {
			fmt.Printf("❌ Error: %s\n", errorMsg)
			return
		}
	}

	// Pretty print the result
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("Result: %v\n", result)
	} else {
		fmt.Println(string(jsonBytes))
	}
	fmt.Println()
}

func init() {
	// Add version command
	rootCmd.AddCommand(versionCmd)
	
	// Add flags for commands
	caesarCmd.Flags().Bool("decrypt", false, "Decrypt instead of encrypt")
	
	qrCmd.Flags().StringP("output", "o", "", "Output filename (default: auto-generated)")
	
	stringsCmd.Flags().Int("min-length", 4, "Minimum string length to extract")
	
	exifCmd.Flags().Bool("verbose", false, "Show verbose output")
	
	httpCmd.Flags().StringP("method", "m", "GET", "HTTP method")
	httpCmd.Flags().StringP("data", "d", "", "Request data")
	
	pcapCmd.Flags().String("filter", "", "Filter packets (e.g., 'http', 'tcp')")
	
	sherlockCmd.Flags().Int("timeout", 5, "Request timeout in seconds")

	// Build command tree
	cryptoCmd.AddCommand(caesarCmd)
	miscCmd.AddCommand(qrCmd)
	fileCmd.AddCommand(stringsCmd)
	stegoCmd.AddCommand(exifCmd)
	webCmd.AddCommand(httpCmd)
	networkCmd.AddCommand(pcapCmd)
	osintCmd.AddCommand(sherlockCmd)
	
	rootCmd.AddCommand(cryptoCmd, miscCmd, fileCmd, stegoCmd, webCmd, networkCmd, osintCmd)
}

func main() {
	// Show banner on first run or when no arguments
	if len(os.Args) == 1 {
		fmt.Printf(banner, version)
		fmt.Printf("💡 Quick start: cst --help\n\n")
	}
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}