#!/bin/bash

APP=rss-generator

RELATIVE_PROJECT_DIR_PATH=..
PROJECT_DIR=$(cd "$(dirname "$0")" || exit; cd ${RELATIVE_PROJECT_DIR_PATH} || exit; pwd)

# init mysql log files
mkdir -p "${PROJECT_DIR}"/log/mysql
touch "${PROJECT_DIR}"/log/mysql/mysqld.log

# init log file permission
find "${PROJECT_DIR}"/log -type f -print | xargs chmod 666
