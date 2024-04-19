#!/bin/bash

cd sql/schema

DIRECTION="$1"

if [ "$DIRECTION" == "" ]; then
	DIRECTION="up"
fi

goose postgres "postgres://postgres:postgres@localhost:5432/blogator" $DIRECTION

cd ../..

