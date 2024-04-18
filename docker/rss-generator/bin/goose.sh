#!/bin/bash -e

set -e

APP=rss-generator

### CONSTANT VARIABLES ###
MIGRATION_DIR="/usr/local/src/${APP}/migrations"
DRIVER="mysql"
DATABASE_DSN="${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT:-3306})/${DATABASE_NAME}?parseTime=true"
###

shift $((OPTIND - 1))
goose -dir ${MIGRATION_DIR} ${DRIVER} "${DATABASE_DSN}" "$@"

exit
