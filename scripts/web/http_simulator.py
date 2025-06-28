#!/usr/bin/env python3

import sys
import json
import requests
from urllib.parse import urlparse

def simulate_http_request(url, method='GET', headers=None, data=None):
    try:
        # Validate URL
        parsed_url = urlparse(url)
        if not parsed_url.scheme or not parsed_url.netloc:
            return {
                "success": False,
                "error": "Invalid URL format"
            }
        
        # Set default headers
        if headers is None:
            headers = {
                'User-Agent': 'CTF-Toolkit-HTTP-Simulator/1.0'
            }
        
        # Make the request based on method
        method = method.upper()
        
        if method == 'GET':
            response = requests.get(url, headers=headers, timeout=10)
        elif method == 'POST':
            response = requests.post(url, headers=headers, data=data, timeout=10)
        elif method == 'PUT':
            response = requests.put(url, headers=headers, data=data, timeout=10)
        elif method == 'DELETE':
            response = requests.delete(url, headers=headers, timeout=10)
        else:
            return {
                "success": False,
                "error": f"Unsupported HTTP method: {method}"
            }
        
        # Extract response data
        result = {
            "success": True,
            "request": {
                "url": url,
                "method": method,
                "headers": dict(headers) if headers else {},
                "data": data
            },
            "response": {
                "status_code": response.status_code,
                "headers": dict(response.headers),
                "content": response.text[:1000],  # Limit content to 1000 chars
                "content_length": len(response.text),
                "encoding": response.encoding
            },
            "tool": "http_simulator"
        }
        
        return result
        
    except requests.exceptions.Timeout:
        return {"success": False, "error": "Request timed out"}
    except requests.exceptions.ConnectionError:
        return {"success": False, "error": "Connection error"}
    except requests.exceptions.RequestException as e:
        return {"success": False, "error": f"Request failed: {str(e)}"}
    except Exception as e:
        return {"success": False, "error": f"Unexpected error: {str(e)}"}

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print(json.dumps({
            "success": False, 
            "error": "Usage: python http_simulator.py <url> [method] [data]"
        }))
        sys.exit(1)
    
    url = sys.argv[1]
    method = sys.argv[2] if len(sys.argv) > 2 else 'GET'
    data = sys.argv[3] if len(sys.argv) > 3 else None
    
    result = simulate_http_request(url, method, data=data)
    print(json.dumps(result, indent=2))