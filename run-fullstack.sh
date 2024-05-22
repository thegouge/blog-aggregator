#!/bin/bash

(cd server; go build -o out && ./out) & (cd client; npm run dev) && fg
