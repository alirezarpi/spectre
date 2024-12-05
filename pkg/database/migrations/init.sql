CREATE TABLE IF NOT EXISTS cve_results (
    id INTEGER PRIMARY KEY,
    cve_id TEXT,
    description TEXT,
    severity TEXT,
    image_name TEXT
);
