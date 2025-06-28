#!/usr/bin/env python3

import sys
import json
import qrcode
import base64
from io import BytesIO

def generate_qr_code(data, output_file=None):
    try:
        # Create a QR Code instance
        qr = qrcode.QRCode(
            version=1,
            error_correction=qrcode.constants.ERROR_CORRECT_L,
            box_size=10,
            border=4,
        )
        
        # Add data to the QR Code
        qr.add_data(data)
        qr.make(fit=True)
        
        # Create QR code image
        img = qr.make_image(fill_color="black", back_color="white")
        
        if output_file:
            img.save(output_file)
            return {
                "success": True,
                "message": f"QR code saved to {output_file}",
                "data": data,
                "tool": "qr_generator"
            }
        else:
            # Return base64 encoded image
            buffer = BytesIO()
            img.save(buffer, format='PNG')
            img_str = base64.b64encode(buffer.getvalue()).decode()
            
            return {
                "success": True,
                "message": "QR code generated successfully",
                "data": data,
                "image": img_str,
                "format": "base64_png",
                "tool": "qr_generator"
            }
            
    except Exception as e:
        return {
            "success": False,
            "error": f"QR code generation failed: {str(e)}"
        }

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"success": False, "error": "Usage: python qr_generator.py <data> [output_file]"}))
        sys.exit(1)
    
    data = sys.argv[1]
    output_file = sys.argv[2] if len(sys.argv) > 2 else None
    
    result = generate_qr_code(data, output_file)
    print(json.dumps(result, indent=2))