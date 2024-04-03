#!/bin/bash

source .env

/usr/bin/minio server /data &

export MINIO_ROOT_USER=$(echo "$MINIO_ROOT_USER" | tr -d '\r')
export MINIO_ROOT_PASSWORD=$(echo "$MINIO_ROOT_PASSWORD" | tr -d '\r')
export MINIO_HOST=$(echo "$MINIO_HOST" | tr -d '\r')
export MINIO_PORT=$(echo "$MINIO_PORT" | tr -d '\r')
export MINIO_BUCKET_NAME=$(echo "$MINIO_BUCKET_NAME" | tr -d '\r')

sleep 5
mc alias set myminio "$MINIO_HOST:$MINIO_PORT" "$MINIO_ROOT_USER" "$MINIO_ROOT_PASSWORD"

mc admin info myminio

mc mb myminio/"$MINIO_BUCKET_NAME"

tail -f /dev/null