#!/bin/bash

# Simple icon generator for BONK Flatpak
# This creates a basic 256x256 PNG icon

set -e

echo "🎨 Creating BONK icon..."

# Check if ImageMagick is available
if command -v convert &> /dev/null; then
    echo "Using ImageMagick to create icon..."
    
    # Create a simple icon with text
    convert -size 256x256 xc:transparent \
            -gravity center \
            -pointsize 48 \
            -fill "#FF6B35" \
            -font "DejaVu-Sans-Bold" \
            -annotate +0+0 "BONK" \
            -stroke "#000000" \
            -strokewidth 2 \
            -annotate +0+0 "BONK" \
            icon.png
    
    echo "✅ Icon created: icon.png"
    
elif command -v inkscape &> /dev/null; then
    echo "Using Inkscape to create icon..."
    
    # Create SVG first
    cat > temp_icon.svg << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<svg width="256" height="256" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
  <rect width="256" height="256" fill="#f0f0f0" stroke="#333" stroke-width="4" rx="32"/>
  <text x="128" y="140" font-family="DejaVu Sans, sans-serif" font-size="48" font-weight="bold" 
        text-anchor="middle" fill="#FF6B35" stroke="#000" stroke-width="1">BONK</text>
  <circle cx="64" cy="80" r="8" fill="#FF6B35"/>
  <circle cx="192" cy="80" r="8" fill="#FF6B35"/>
  <path d="M80 200 Q128 220 176 200" stroke="#FF6B35" stroke-width="4" fill="none"/>
</svg>
EOF
    
    # Convert SVG to PNG
    inkscape --export-type=png --export-width=256 --export-height=256 temp_icon.svg -o icon.png
    rm temp_icon.svg
    
    echo "✅ Icon created: icon.png"
    
else
    echo "⚠️  Neither ImageMagick nor Inkscape found."
    echo "Creating a placeholder text file instead..."
    echo "To create an icon manually:"
    echo "1. Create a 256x256 PNG file"
    echo "2. Design it to represent BONK (hammer, text, etc.)"
    echo "3. Save it as 'icon.png' in this directory"
    
    # Create a placeholder
    echo "BONK ICON PLACEHOLDER" > icon_placeholder.txt
    echo "📝 Created placeholder: icon_placeholder.txt"
    echo "Please replace with actual icon.png file"
fi