#!/bin/bash

# Download and extract the Go 1.22.9 binary distribution
wget https://golang.org/dl/go1.22.9.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.9.linux-amd64.tar.gz

# Add the Go binary directory to the PATH environment variable
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

# Apply the changes to the current shell session
source ~/.bashrc

# Verify the installation
go version
