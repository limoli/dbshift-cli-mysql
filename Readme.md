# DbShift Client for MySQL

It provides simple and light logic for the management of **database-schema migrations** through the implementation of [Dbshit Core](https://github.com/limoli/dbshift-core). 

You will be able to create migrations, check the current db status, decide to upgrade or downgrade easily.

[![GoDoc](https://godoc.org/limoli/dbshift-cli-mysql?status.svg)](https://godoc.org/github.com/limoli/dbshift-cli-mysql)
[![Build Status](https://travis-ci.org/limoli/dbshift-cli-mysql.svg?branch=master)](https://travis-ci.org/limoli/dbshift-cli-mysql)
[![Go Report Card](https://goreportcard.com/badge/github.com/limoli/dbshift-cli-mysql)](https://goreportcard.com/report/github.com/limoli/dbshift-cli-mysql)
[![Maintainability](https://api.codeclimate.com/v1/badges/xxx/maintainability)](https://codeclimate.com/github/limoli/dbshift-cli-mysql/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/xxx/test_coverage)](https://codeclimate.com/github/limoli/dbshift-cli-mysql/test_coverage)
[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/github.com/limoli/dbshift-cli-mysql/LICENSE)

## Install

`go get github.com/limoli/dbshift-cli-mysql`

## Commands

See [Dbshit Core](https://github.com/limoli/dbshift-core).

## Configuration

| Key                                           | Description                                                   | Value example              |
|---                                            |---                                                            |---                         |
| Core configuration                            | See [Dbshit Core](https://github.com/limoli/dbshift-core).    |                            |
|`DBSHIFT_CLI_MYSQL_TABLE`                      | Table used by dbshift to manage migrations logics.            | `dbshift`                  |
|`DBSHIFT_CLI_MYSQL_USERNAME`                   | Database username                                             | See [mysql drive](https://github.com/go-sql-driver/mysql) |
|`DBSHIFT_CLI_MYSQL_PASSWORD`                   | Database password                                             | See [mysql drive](https://github.com/go-sql-driver/mysql) |
|`DBSHIFT_CLI_MYSQL_DATABASE`                   | Database name                                                 | See [mysql drive](https://github.com/go-sql-driver/mysql) |
|`DBSHIFT_CLI_MYSQL_ADDRESS`                    | Database address                                              | See [mysql drive](https://github.com/go-sql-driver/mysql) |
|`DBSHIFT_CLI_MYSQL_OPTION_IS_MULTI_STATEMENT`  | Table used by dbshift to manage migrations logics.            | `true`/`false` See [mysql drive](https://github.com/go-sql-driver/mysql) |

## Exit codes

| Flag       | Description                                                       |
| ---        | ---                                                               |
| `110`      | When username is not provided via environment.                    |
| `111`      | When password is not provided via environment.                    |
| `112`      | When database name is not provided via environment.               |
| `115`      | When the option isMultiStatement is badly set.                    |
| `130`      | When table-name value is missing (env var is not provided).       |
| `150`      | When db connection gets an error.                                 |
| `160`      | When core gets an error on initialisation.                        |
