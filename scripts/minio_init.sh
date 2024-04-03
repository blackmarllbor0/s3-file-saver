#!/bin/bash

source .env

/usr/bin/minio server /data &

sleep 5
mc alias set myminio http://localhost:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb myminio/$MINIO_BUCKET_NAME

tail -f /dev/null