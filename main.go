package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	dbshiftcore "github.com/limoli/dbshift-core"
	"os"
)

const (
	envPrefix                      = "DBSHIFT_CLI_MYSQL"
	envMySqlTable                  = envPrefix + "_TABLE"
	envMySqlUsername               = envPrefix + "_USERNAME"
	envMySqlPassword               = envPrefix + "_PASSWORD"
	envMySqlDatabase               = envPrefix + "_DATABASE"
	envMySqlAddress                = envPrefix + "_ADDRESS"
	envMySqlOptionIsMultiStatement = envPrefix + "_OPTION_IS_MULTI_STATEMENT"
)

func main() {

	cfg, err := initConfiguration()
	if err != nil {
		dbshiftcore.PrintFailure(err.Error())
		if errWithCodeObj, ok := err.(*errorWithCode); ok {
			os.Exit(errWithCodeObj.code)
		}
	}

	db, err := newMysqlDatabase(*cfg)
	if err != nil {
		dbshiftcore.PrintFailure(err.Error())
		os.Exit(150)
	}

	cmd, err := dbshiftcore.NewCmd(db)
	if err != nil {
		dbshiftcore.PrintFailure(err.Error())
		os.Exit(160)
	}

	cmd.Run()
}

func initConfiguration() (*config, error) {

	const errMsg = "missing or unset %s environment variable"

	username := os.Getenv(envMySqlUsername)
	if username == "" {
		return nil, &errorWithCode{110, fmt.Sprintf(errMsg, envMySqlUsername)}
	}

	password := os.Getenv(envMySqlPassword)
	if password == "" {
		return nil, &errorWithCode{111, fmt.Sprintf(errMsg, envMySqlUsername)}
	}

	databaseName := os.Getenv(envMySqlDatabase)
	if databaseName == "" {
		return nil, &errorWithCode{112, fmt.Sprintf(errMsg, envMySqlUsername)}
	}

	address := os.Getenv(envMySqlAddress)

	optionIsMultiStatement := os.Getenv(envMySqlOptionIsMultiStatement)
	if optionIsMultiStatement != "" && !(optionIsMultiStatement == "true" || optionIsMultiStatement == "false") {
		return nil, &errorWithCode{115, fmt.Sprintf("bad boolean value for %s environment variable", envMySqlOptionIsMultiStatement)}
	}

	tableName := os.Getenv(envMySqlTable)
	if tableName == "" {
		return nil, &errorWithCode{code: 130, message: fmt.Sprintf("missing or unset %s environment variable", envMySqlTable)}
	}

	return &config{
		username:  username,
		password:  password,
		address:   address,
		dbName:    databaseName,
		tableName: tableName,
		options: configOptions{
			isMultiStatement: optionIsMultiStatement,
		},
	}, nil
}
