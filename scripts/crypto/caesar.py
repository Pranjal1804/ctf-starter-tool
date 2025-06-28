#!/usr/bin/env python3

import sys
import json

def caesar_cipher(text, shift, mode="encrypt"):
    """
    Perform Caesar cipher encryption/decryption
    """
    if mode == "decrypt":
        shift = -shift
    
    result = ""
    
    for char in text:
        if char.isalpha():
            # Handle uppercase
            if char.isupper():
                result += chr((ord(char) - ord('A') + shift) % 26 + ord('A'))
            # Handle lowercase
            else:
                result += chr((ord(char) - ord('a') + shift) % 26 + ord('a'))
        else:
            # Keep non-alphabetic characters unchanged
            result += char
    
    return result

def brute_force_caesar(text):
    """
    Try all possible Caesar cipher shifts
    """
    results = []
    for shift in range(26):
        decrypted = caesar_cipher(text, shift, "decrypt")
        results.append({
            "shift": shift,
            "result": decrypted
        })
    return results

def main():
    if len(sys.argv) < 3:
        print(json.dumps({"error": "Usage: python caesar.py <text> <shift> [mode]"}))
        sys.exit(1)
    
    text = sys.argv[1]
    
    try:
        shift = int(sys.argv[2])
        mode = sys.argv[3] if len(sys.argv) > 3 else "encrypt"
        
        if mode not in ["encrypt", "decrypt"]:
            mode = "encrypt"
        
        result = caesar_cipher(text, shift, mode)
        
        output = {
            "mode": mode,
            "shift": shift,
            "original_text": text,
            "result": result,
            "tool": "caesar_cipher"
        }
        
        print(json.dumps(output))
        
    except ValueError:
        print(json.dumps({"error": "Shift must be a valid integer"}))
        sys.exit(1)

if __name__ == "__main__":
    main()