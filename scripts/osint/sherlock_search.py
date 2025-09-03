import sys
import json
import requests
import concurrent.futures
import argparse

def search_username(username, timeout=5):
    try:
        # Popular social media platforms to check
        platforms = {
            "GitHub": f"https://github.com/{username}",
            "Twitter": f"https://twitter.com/{username}",
            "Instagram": f"https://instagram.com/{username}",
            "Reddit": f"https://reddit.com/user/{username}",
            "LinkedIn": f"https://linkedin.com/in/{username}",
            "YouTube": f"https://youtube.com/@{username}",
            "Facebook": f"https://facebook.com/{username}"
        }
        
        results = {}
        
        def check_platform(platform_data):
            platform, url = platform_data
            try:
                response = requests.get(url, timeout=timeout, allow_redirects=True)
                
                # Different platforms have different indicators for existing profiles
                if platform == "GitHub":
                    exists = response.status_code == 200 and "Not Found" not in response.text
                elif platform == "Twitter":
                    exists = response.status_code == 200 and "This account doesn't exist" not in response.text
                elif platform == "Instagram":
                    exists = response.status_code == 200 and "Sorry, this page isn't available" not in response.text
                elif platform == "Reddit":
                    exists = response.status_code == 200 and "page not found" not in response.text.lower()
                else:
                    exists = response.status_code == 200
                
                return platform, {
                    "url": url,
                    "exists": exists,
                    "status_code": response.status_code,
                    "response_time": round(response.elapsed.total_seconds(), 2)
                }
            except requests.exceptions.RequestException as e:
                return platform, {
                    "url": url,
                    "exists": False,
                    "status_code": None,
                    "error": str(e)
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
        
        # Separate found and not found for better display
        found_profiles = {k: v for k, v in results.items() if v.get("exists", False)}
        not_found_profiles = {k: v for k, v in results.items() if not v.get("exists", False)}
        
        return {
            "success": True,
            "username": username,
            "found_count": found_count,
            "total_checked": len(platforms),
            "timeout": timeout,
            "found_profiles": found_profiles,
            "not_found_profiles": not_found_profiles,
            "summary": f"Found {found_count} out of {len(platforms)} platforms",
            "tool": "sherlock_search"
        }
        
    except Exception as e:
        return {
            "success": False,
            "error": f"Username search failed: {str(e)}"
        }

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Search username across social media platforms')
    parser.add_argument('username', help='Username to search for')
    parser.add_argument('--timeout', type=int, default=5, help='Request timeout in seconds (default: 5)')
    
    try:
        args = parser.parse_args()
    except SystemExit:
        # Handle argparse errors gracefully
        print(json.dumps({
            "success": False,
            "error": "Usage: python sherlock_search.py <username> [--timeout SECONDS]"
        }))
        sys.exit(1)
    
    username = args.username
    timeout = args.timeout
    
    if not username or len(username) < 2:
        print(json.dumps({
            "success": False,
            "error": "Username must be at least 2 characters long"
        }))
        sys.exit(1)
    
    if timeout < 1 or timeout > 30:
        print(json.dumps({
            "success": False,
            "error": "Timeout must be between 1 and 30 seconds"
        }))
        sys.exit(1)
    
    response = search_username(username, timeout)
    print(json.dumps(response, indent=2))