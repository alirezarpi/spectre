#!/bin/bash -e
echo "[*] building go spectre cve scanner..."
go build -o spectre ./cmd/scanner
echo "[*] build complete!"

