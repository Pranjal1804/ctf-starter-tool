#!/usr/bin/env python3

import sys
import json
import requests
import concurrent.futures
import time

def search_username(username):
    try:
        # Popular social media platforms to check
        platforms = {
            "GitHub": f"https://github.com/{username}",
            "Twitter": f"https://twitter.com/{username}",
            "Instagram": f"https://instagram.com/{username}",
            "Reddit": f"https://reddit.com/user/{username}",
            "LinkedIn": f"https://linkedin.com/in/{username}",
            "YouTube": f"https://youtube.com/@{username}",
            "Facebook": f"https://facebook.com/{username}",
            "TikTok": f"https://tiktok.com/@{username}"
        }
        
        results = {}
        
        def check_platform(platform_data):
            platform, url = platform_data
            try:
                response = requests.get(url, timeout=5, allow_redirects=True)
                
                # Different platforms have different indicators for existing profiles
                if platform == "GitHub":
                    exists = response.status_code == 200 and "Not Found" not in response.text
                elif platform == "Twitter":
                    exists = response.status_code == 200 and "This account doesn't exist" not in response.text
                else:
                    exists = response.status_code == 200
                
                return platform, {
                    "url": url,
                    "exists": exists,
                    "status_code": response.status_code,
                    "response_time": response.elapsed.total_seconds()
                }
            except requests.exceptions.RequestException:
                return platform, {
                    "url": url,
                    "exists": False,
                    "status_code": None,
                    "error": "Connection failed"
                }
        
        # Use ThreadPoolExecutor for concurrent requests
        with concurrent.futures.ThreadPoolExecutor(max_workers=5) as executor:
            future_to_platform = {
                executor.submit(check_platform, item): item[0] 
                for item in platforms.items()
            }
            
            for future in concurrent.futures.as_completed(future_to_platform):
                platform, result = future.result()
                results[platform] = result
        
        # Count found profiles
        found_count = sum(1 for result in results.values() if result.get("exists", False))
        
        return {
            "success": True,
            "username": username,
            "found_count": found_count,
            "total_checked": len(platforms),
            "results": results,
            "tool": "sherlock_search"
        }
        
    except Exception as e:
        return {
            "success": False,
            "error": f"Username search failed: {str(e)}"
        }

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({
            "success": False,
            "error": "Usage: python sherlock_search.py <username>"
        }))
        sys.exit(1)
    
    username = sys.argv[1]
    
    if not username or len(username) < 2:
        print(json.dumps({
            "success": False,
            "error": "Username must be at least 2 characters long"
        }))
        sys.exit(1)
    
    response = search_username(username)
    print(json.dumps(response, indent=2))