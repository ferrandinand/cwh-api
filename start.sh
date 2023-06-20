#!/bin/bash
SERVER_ADDRESS=0.0.0.0 \
SERVER_PORT=8080 \
DB_USER=root \
DB_PASSWD=codecamp \
DB_ADDR=mysql \
DB_PORT=3307 \
DB_NAME=cwh \
STATUS_API="https://status.core.un-icc.cloud" \
go run main.go