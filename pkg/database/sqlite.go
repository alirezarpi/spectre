package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "spectre/pkg/cve"
)

func InitSQLite(filepath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        return nil, err
    }

    schema := `
    CREATE TABLE IF NOT EXISTS cve_results (
        id INTEGER PRIMARY KEY,
        cve_id TEXT,
        description TEXT,
        severity TEXT,
        image_name TEXT
    );
    `
    _, err = db.Exec(schema)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func SaveResults(db *sql.DB, results []cve.CVE) error {
    for _, result := range results {
        _, err := db.Exec("INSERT INTO cve_results (cve_id, description, severity, image_name) VALUES (?, ?, ?, ?)",
            result.ID, result.Description, result.Severity, "docker-image-name")
        if err != nil {
            return err
        }
    }
    return nil
}

