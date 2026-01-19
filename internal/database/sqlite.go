package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func New(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(1) // SQLite works best with single connection
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	// Enable WAL mode for better concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, fmt.Errorf("failed to set WAL mode: %w", err)
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys=ON"); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	database := &DB{db}

	// Run migrations
	if err := database.migrate(); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return database, nil
}

func (db *DB) migrate() error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS settings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key TEXT UNIQUE NOT NULL,
			value TEXT NOT NULL,
			created_at INTEGER DEFAULT (strftime('%s', 'now')),
			updated_at INTEGER DEFAULT (strftime('%s', 'now'))
		)`,
		`CREATE TABLE IF NOT EXISTS scan_order (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			file_name TEXT UNIQUE NOT NULL,
			position INTEGER NOT NULL,
			created_at INTEGER DEFAULT (strftime('%s', 'now'))
		)`,
		`CREATE INDEX IF NOT EXISTS idx_scan_order_position ON scan_order(position)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	// Insert default settings if not exist
	defaultSettings := map[string]string{
		"resolution":     "300",
		"quality":        "80",
		"format":         "jpeg",
		"default_device": "",
		"theme":          "light",
	}

	for key, value := range defaultSettings {
		_, err := db.Exec(
			`INSERT OR IGNORE INTO settings (key, value) VALUES (?, ?)`,
			key, value,
		)
		if err != nil {
			return fmt.Errorf("failed to insert default setting %s: %w", key, err)
		}
	}

	return nil
}

// GetSetting retrieves a setting by key
func (db *DB) GetSetting(key string) (string, error) {
	var value string
	err := db.QueryRow("SELECT value FROM settings WHERE key = ?", key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return value, err
}

// SetSetting sets or updates a setting
func (db *DB) SetSetting(key, value string) error {
	_, err := db.Exec(
		`INSERT INTO settings (key, value, updated_at) VALUES (?, ?, strftime('%s', 'now'))
		ON CONFLICT(key) DO UPDATE SET value = excluded.value, updated_at = strftime('%s', 'now')`,
		key, value,
	)
	return err
}

// GetAllSettings retrieves all settings
func (db *DB) GetAllSettings() (map[string]string, error) {
	rows, err := db.Query("SELECT key, value FROM settings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	settings := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		settings[key] = value
	}

	return settings, rows.Err()
}

// GetFileOrder retrieves the order position for a file
func (db *DB) GetFileOrder(fileName string) (int, error) {
	var position int
	err := db.QueryRow("SELECT position FROM scan_order WHERE file_name = ?", fileName).Scan(&position)
	if err == sql.ErrNoRows {
		return -1, nil
	}
	return position, err
}

// SetFileOrder sets the order position for a file
func (db *DB) SetFileOrder(fileName string, position int) error {
	_, err := db.Exec(
		`INSERT INTO scan_order (file_name, position) VALUES (?, ?)
		ON CONFLICT(file_name) DO UPDATE SET position = excluded.position`,
		fileName, position,
	)
	return err
}

// GetAllFileOrders retrieves all file orders
func (db *DB) GetAllFileOrders() (map[string]int, error) {
	rows, err := db.Query("SELECT file_name, position FROM scan_order ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make(map[string]int)
	for rows.Next() {
		var fileName string
		var position int
		if err := rows.Scan(&fileName, &position); err != nil {
			return nil, err
		}
		orders[fileName] = position
	}

	return orders, rows.Err()
}

// UpdateFileOrders updates multiple file orders in a transaction
func (db *DB) UpdateFileOrders(orders map[string]int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		`INSERT INTO scan_order (file_name, position) VALUES (?, ?)
		ON CONFLICT(file_name) DO UPDATE SET position = excluded.position`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for fileName, position := range orders {
		if _, err := stmt.Exec(fileName, position); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteFileOrder removes the order for a file
func (db *DB) DeleteFileOrder(fileName string) error {
	_, err := db.Exec("DELETE FROM scan_order WHERE file_name = ?", fileName)
	return err
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}
