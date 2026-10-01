#!/bin/bash

set -e

echo "Downloading Flutter SDK..."
if [ ! -d "flutter" ]; then
  git clone https://github.com/flutter/flutter.git -b stable --depth 1
fi

export PATH="$PATH:`pwd`/flutter/bin"

echo "Checking Flutter installation..."
flutter doctor -v

echo "Getting dependencies..."
flutter pub get

echo "Building for web..."
flutter build web --release
