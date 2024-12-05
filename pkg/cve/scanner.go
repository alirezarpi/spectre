package cve

import "spectre/pkg/docker"

type CVE struct {
    ID          string
    Description string
    Severity    string
}

func ScanImage(image *docker.ImageDetails) ([]CVE, error) {
    // Example: Mock scan returning static results
    results := []CVE{
        {ID: "CVE-2023-1234", Description: "Example vulnerability", Severity: "High"},
    }
    return results, nil
}

