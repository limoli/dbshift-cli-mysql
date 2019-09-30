package main

import (
	"database/sql"
	"fmt"
	dbshiftcore "github.com/limoli/dbshift-core"
)

type mysqlDatabase struct {
	client *sql.DB
	cfg    config
}

type config struct {
	username  string
	password  string
	address   string
	dbName    string
	tableName string
	options   configOptions
}

type configOptions struct {
	isMultiStatement string
}

func newMysqlDatabase(c config) (*mysqlDatabase, error) {
	connectionString := fmt.Sprintf("%s:%s@%s/%s?multiStatements=%s", c.username, c.password, c.address, c.dbName, c.options.isMultiStatement)
	client, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return &mysqlDatabase{client: client, cfg: c}, err
}

func (db *mysqlDatabase) isInitialised() bool {
	_, err := db.client.Query(fmt.Sprintf("SELECT 1 FROM %s LIMIT 1", db.cfg.tableName))
	return err == nil || err == sql.ErrNoRows
}

func (db *mysqlDatabase) initialise() error {
	_, err := db.client.Exec(fmt.Sprintf(`CREATE TABLE %s
		(
    		id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'migration id',
    		version varchar(150) NOT NULL COMMENT 'migration name',
    		type tinyint(1) NOT NULL COMMENT 'migration type: DOWN(0), UP(1)',
    		name varchar(250) NOT NULL COMMENT 'migration name',
    		executionTime double DEFAULT NULL COMMENT 'execution time in seconds',
    		appliedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
    		PRIMARY KEY (id)
		) ENGINE = MyISAM DEFAULT CHARSET = utf8mb4;`, db.cfg.tableName))
	return err
}

func (db *mysqlDatabase) GetExtension() string {
	return "sql"
}

func (db *mysqlDatabase) GetStatus() (*dbshiftcore.Status, error) {
	if !db.isInitialised() {
		if err := db.initialise(); err != nil {
			return nil, err
		}
	}

	var status dbshiftcore.Status

	row := db.client.QueryRow(fmt.Sprintf("SELECT version, type FROM %s ORDER BY id DESC LIMIT 1", db.cfg.tableName))
	if err := row.Scan(&status.Version, &status.Type); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &status, nil
}

func (db *mysqlDatabase) SetStatus(m dbshiftcore.Migration, executionTimeInSeconds float64) error {
	if !db.isInitialised() {
		if err := db.initialise(); err != nil {
			return err
		}
	}

	_, err := db.client.Exec(
		fmt.Sprintf("INSERT INTO %s (version, type, name, executionTime) VALUES (?,?,?,?)", db.cfg.tableName),
		m.Version,
		m.Type,
		m.Name,
		executionTimeInSeconds)

	return err
}

func (db *mysqlDatabase) ExecuteMigration(queries []byte) error {
	_, err := db.client.Exec(string(queries))
	return err
}
