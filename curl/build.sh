#!/bin/bash

green='\033[0;32m'
clear='\033[0m'

go build main.go
# Rename the executable if needed
mv main mycurl
# Move the executable to a directory in the PATH
sudo mv mycurl /usr/local/bin/
echo "Done!🚀"
printf "${green}Build completed. You can now run 'mycurl' from the command line!${clear}🚀"