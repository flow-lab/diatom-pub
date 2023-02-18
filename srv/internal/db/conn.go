package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

// ConnectTCPSocket initializes a TCP connection pool for a Cloud SQL
// instance of Postgres.
func ConnectTCPSocket() (*sql.DB, error) {
	mustGetEnv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
		}
		return v
	}

	var (
		dbUser    = mustGetEnv("DB_USER") // e.g. 'my-db-user'
		dbPwd     = mustGetEnv("DB_PASS") // e.g. 'my-db-password'
		dbTCPHost = mustGetEnv("DB_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
		dbPort    = mustGetEnv("DB_PORT") // e.g. '5432'
		dbName    = mustGetEnv("DB_NAME") // e.g. 'my-database'
	)

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	// For deployments that connect directly to a Cloud SQL instance without
	// using the Cloud SQL Proxy, configuring SSL certificates will ensure the
	// connection is encrypted.
	if dbCertPath, ok := os.LookupEnv("DB_CERT_PATH"); ok { // e.g., '/path/to/db/certs'
		if dbCertPath[len(dbCertPath)-1:] == "/" {
			dbCertPath = dbCertPath[:len(dbCertPath)-1]
		}
		dbURI += fmt.Sprintf(" sslmode=require sslrootcert=%s/server-ca.pem sslcert=%s/client-cert.pem sslkey=%s/client-key.pem",
			dbCertPath, dbCertPath, dbCertPath)
	}

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return dbPool, nil
}
