package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnector struct {
	Host           string
	Port           int
	Username       string
	Password       string
	DBName         string
	ConnectTimeout int
	SSLCert        string
	SSLKey         string
	SSLRootCert    string
	SSLMode        string
}

func (dbConnector DatabaseConnector) Connect() (*sql.DB, error) {
	var sslMode string
	if dbConnector.SSLMode != "" && dbConnector.SSLMode != "disabled" {
		sslMode = fmt.Sprintf("sslmode=%s&sslrootcert=%s&sslcert=%s&sslkey=%s",
			dbConnector.SSLMode,
			dbConnector.SSLRootCert,
			dbConnector.SSLCert,
			dbConnector.SSLKey)
	}

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		dbConnector.Username,
		dbConnector.Password,
		dbConnector.Host,
		dbConnector.Port,
		dbConnector.DBName,
		sslMode)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
