#!/usr/bin/env sh

set -eu

APP=rss-generator

usage() {
    echo "Usage: $0 [-d] [-m] [-q] [-w]" 1>&2
    echo "Options: " 1>&2
    echo "-d: Run as development mode" 1>&2
    echo "-j: Run as job mode" 1>&2
    echo "-m: Run migration" 1>&2
    echo "-q: Quit without running server" 1>&2
    echo "-w: Wait for database to start" 1>&2
    exit 1
}

WAIT=0
QUIT=0
MIGRATION=0
ENVIRONMENT=prod
JOB=0

while getopts :djwmqh OPT
do
    case $OPT in
    d)  ENVIRONMENT=dev
        ;;
    j)  JOB=1
        ;;
    w)  WAIT=1
        ;;
    m)  MIGRATION=1
        ;;
    q)  QUIT=1
        ;;
    h)  usage
        ;;
    \?) usage
        ;;
    esac
done

if [ "$WAIT" = "1" ]; then
    echo "Waiting for db..."
    dockerize -wait "tcp://${DATABASE_HOST}:${DATABASE_PORT:-3306}" -timeout 480s
    echo "Connected to db!"
fi

if [ "$MIGRATION" = "1" ]; then
    echo "Running migration..."
    sh /usr/local/bin/goose.sh up
    echo "Migration completed!"
fi

if [ "$QUIT" = "1" ]; then
    exit 0
fi

if [ "$JOB" = "1" ]; then
    /usr/local/bin/${APP} jobs update_feeds
    /usr/local/bin/${APP} jobs delete_expired_test_feeds
    exit 0
fi

if [ "${ENVIRONMENT:-}" = "dev" ]; then
    air -c .air.toml
elif [ "${ENVIRONMENT:-}" = "prod" ]; then
    /usr/local/bin/${APP} serve
fi
