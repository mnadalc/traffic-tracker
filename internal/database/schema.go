package database

const createRoutesTable = `
	CREATE TABLE IF NOT EXISTS routes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		origin TEXT NOT NULL,
		destination TEXT NOT NULL
	);
`

const createTrafficDataTable = `
	CREATE TABLE IF NOT EXISTS traffic_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		route_id INTEGER NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		duration_seconds INTEGER NOT NULL,
		distance_meters INTEGER NOT NULL,
		traffic_condition TEXT,
		FOREIGN KEY (route_id) REFERENCES routes(id) ON DELETE CASCADE
	);
`