#!/bin/bash

sleep 2 && goose -dir "./migrations" postgres "host=${PG_HOST} port=${PG_PORT} dbname=${PG_DB} user=${PG_USER} password=${PG_PASSWORD} sslmode=disable" up -v