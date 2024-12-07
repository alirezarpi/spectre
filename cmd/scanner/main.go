package main

import (
    "os"

    "spectre/pkg/docker"
    "spectre/pkg/cve"
    "spectre/pkg/database"
    "spectre/pkg/log"
)

func main() {
    image := os.Args[1]
    if image == "" {
        log.Fatal("Please specify a Docker image")
    }

    log.Info("Scanning image:" + image)
    
    db, err := database.InitSQLite("data/cve_scanner.db")
    if err != nil {
        log.Fatal("failed to initialize database: " + err.Error())
    }
    
    defer db.Close()

    details, err := docker.InspectImage(image)
    if err != nil {
        log.Fatal("failed to inspect image: " + err.Error())
    }

    results, err := cve.ScanImage(details)
    if err != nil {
        log.Fatal("failed to scan image: " + err.Error())
    }

    err = database.SaveResults(db, results)
    if err != nil {
        log.Fatal("failed to save results: " + err.Error())
    }

    log.Info("Scan complete! Results saved to the database.")
}
