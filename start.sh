#!/bin/bash
SERVER_ADDRESS=api \
SERVER_PORT=8080 \
AUTH_ADDRESS=auth \
AUTH_PORT=8181 \
DB_USER=root \
DB_PASSWD=codecamp \
DB_ADDR=mysql \
DB_PORT=3307 \
DB_NAME=cwh \
AMQP_SERVER_URL=message-broker \
AMQP_USER=rabbitmq \
AMQP_PASSWD=rabbitmq \
AMQP_PORT=5672 \
go run main.go