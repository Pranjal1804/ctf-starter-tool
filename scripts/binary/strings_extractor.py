#!/usr/bin/env python3

import sys
import json
import re
import os

def extract_strings(file_path, min_length=4):
    try:
        with open(file_path, 'rb') as f:
            content = f.read()
        
        # Extract ASCII strings
        ascii_strings = re.findall(b'[ -~]{' + str(min_length).encode() + b',}', content)
        ascii_results = [s.decode('ascii', errors='ignore') for s in ascii_strings]
        
        # Extract Unicode strings
        unicode_strings = re.findall(b'(?:[ -~]\x00){' + str(min_length).encode() + b',}', content)
        unicode_results = [s.decode('utf-16le', errors='ignore') for s in unicode_strings]
        
        # Filter out empty strings
        ascii_results = [s for s in ascii_results if s.strip()]
        unicode_results = [s for s in unicode_results if s.strip()]
        
        result = {
            "success": True,
            "file_info": {
                "filename": os.path.basename(file_path),
                "size": len(content)
            },
            "strings": {
                "ascii": ascii_results[:100],  # Limit to first 100 strings
                "unicode": unicode_results[:100],
                "total_ascii": len(ascii_results),
                "total_unicode": len(unicode_results)
            },
            "tool": "strings_extractor"
        }
        
        return result
        
    except Exception as e:
        return {
            "success": False,
            "error": f"String extraction failed: {str(e)}"
        }

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({
            "success": False,
            "error": "Usage: python strings_extractor.py <file_path>"
        }))
        sys.exit(1)

    file_path = sys.argv[1]
    
    if not os.path.exists(file_path):
        print(json.dumps({
            "success": False,
            "error": "File not found"
        }))
        sys.exit(1)
    
    result = extract_strings(file_path)
    print(json.dumps(result, indent=2))