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
            
            # Remove some exiftool metadata that's not useful
            metadata_to_remove = ['SourceFile', 'ExifTool:ExifToolVersion']
            for key in metadata_to_remove:
                exif_data.pop(key, None)
            
            return {
                "success": True,
                "tool": "exiftool",
                "file_path": image_path,
                "data": exif_data,
                "data_count": len(exif_data),
                "message": f"Found {len(exif_data)} metadata fields using exiftool"
            }
        else:
            # Fallback to PIL/Pillow
            return extract_with_pillow(image_path)
            
    except subprocess.TimeoutExpired:
        return {
            "success": False, 
            "error": "exiftool timeout",
            "fallback": "Trying PIL..."
        }
    except FileNotFoundError:
        # exiftool not found, use PIL
        return extract_with_pillow(image_path)
    except json.JSONDecodeError:
        return extract_with_pillow(image_path)
    except Exception as e:
        return {
            "success": False, 
            "error": f"exiftool failed: {str(e)}",
            "fallback": "Trying PIL..."
        }

def extract_with_pillow(image_path):
    try:
        image = Image.open(image_path)
        exif_data = {}
        basic_info = {}
        
        # Get basic image information
        basic_info['filename'] = os.path.basename(image_path)
        basic_info['format'] = image.format
        basic_info['mode'] = image.mode
        basic_info['size'] = f"{image.size[0]}x{image.size[1]}"
        basic_info['file_size'] = os.path.getsize(image_path)
        
        # Try to get EXIF data
        if hasattr(image, '_getexif') and image._getexif():
            exif = image._getexif()
            for tag_id, value in exif.items():
                tag = TAGS.get(tag_id, tag_id)
                exif_data[tag] = str(value)
        
        # Try to get other metadata
        if hasattr(image, 'info') and image.info:
            for key, value in image.info.items():
                if key not in exif_data:
                    exif_data[key] = str(value)
        
        message = "No EXIF data found" if not exif_data else f"Found {len(exif_data)} EXIF fields"
        
        return {
            "success": True,
            "tool": "pillow",
            "file_path": image_path,
            "basic_info": basic_info,
            "exif_data": exif_data,
            "data_count": len(exif_data),
            "message": message,
            "note": "PNG files typically don't contain EXIF data. Try JPEG files for more metadata."
        }
        
    except Exception as e:
        return {
            "success": False, 
            "error": f"PIL extraction failed: {str(e)}"
        }

def get_file_info(file_path):
    """Get basic file information"""
    try:
        stat = os.stat(file_path)
        return {
            "file_size": stat.st_size,
            "created": stat.st_ctime,
            "modified": stat.st_mtime,
            "file_extension": os.path.splitext(file_path)[1].lower()
        }
    except:
        return {}

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({
            "success": False, 
            "error": "Usage: python exif_extractor.py <image_path>",
            "example": "python exif_extractor.py photo.jpg"
        }))
        sys.exit(1)

    image_path = sys.argv[1]
    
    if not os.path.exists(image_path):
        print(json.dumps({
            "success": False, 
            "error": f"File not found: {image_path}"
        }))
        sys.exit(1)
    
    # Check if it's likely an image file
    valid_extensions = ['.jpg', '.jpeg', '.png', '.tiff', '.bmp', '.gif', '.webp']
    file_ext = os.path.splitext(image_path)[1].lower()
    
    if file_ext not in valid_extensions:
        print(json.dumps({
            "success": False,
            "error": f"File extension '{file_ext}' may not be a supported image format",
            "supported_formats": valid_extensions
        }))
        sys.exit(1)
    
    result = extract_exif_data(image_path)
    print(json.dumps(result, indent=2))