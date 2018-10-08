#!/usr/bin/env bash

PUBLIC_KEY=$( grep -i public public.toml  | sed -e 's/ *Public *= *//' -e 's/"//g' | tr '[:upper:]' '[:lower:]' )
DB_FILE_NAME=$( echo ${PUBLIC_KEY} | xxd -r -p | gsha256sum | sed -e 's/[ -]//g' )
DB_FILE_NAME=${DB_FILE_NAME}.db
echo $DB_FILE_NAME
