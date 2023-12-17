#!/bin/bash

# Build the Go executable
go build main.go

# Rename the executable if needed
mv main mycurl

# Move the executable to a directory in the PATH
sudo mv mycurl /usr/local/bin/

echo "Build completed. You can now run 'myprogram' from the command line."