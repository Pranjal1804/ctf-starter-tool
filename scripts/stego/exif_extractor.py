#!/usr/bin/env python3

import sys
import json
import subprocess
import os
from PIL import Image
from PIL.ExifTags import TAGS

def extract_exif_data(image_path):
    try:
        # Try using exiftool first
        result = subprocess.run([
            'exiftool', '-json', '-all', image_path
        ], capture_output=True, text=True, timeout=30)
        
        if result.returncode == 0:
            exif_data = json.loads(result.stdout)[0]
            return {
                "success": True,
                "tool": "exiftool",
                "data": exif_data
            }
        else:
            # Fallback to PIL/Pillow
            return extract_with_pillow(image_path)
            
    except subprocess.TimeoutExpired:
        return {"success": False, "error": "exiftool timeout"}
    except FileNotFoundError:
        # exiftool not found, use PIL
        return extract_with_pillow(image_path)
    except Exception as e:
        return {"success": False, "error": f"exiftool failed: {str(e)}"}

def extract_with_pillow(image_path):
    try:
        image = Image.open(image_path)
        exif_data = {}
        
        if hasattr(image, '_getexif'):
            exif = image._getexif()
            if exif:
                for tag_id, value in exif.items():
                    tag = TAGS.get(tag_id, tag_id)
                    exif_data[tag] = str(value)
        
        return {
            "success": True,
            "tool": "pillow",
            "data": exif_data
        }
    except Exception as e:
        return {"success": False, "error": f"PIL extraction failed: {str(e)}"}

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({"success": False, "error": "Usage: python exif_extractor.py <image_path>"}))
        sys.exit(1)

    image_path = sys.argv[1]
    
    if not os.path.exists(image_path):
        print(json.dumps({"success": False, "error": "File not found"}))
        sys.exit(1)
    
    result = extract_exif_data(image_path)
    print(json.dumps(result, indent=2))