#!/bin/sh
docker run --name secretify --rm \
-p 8800:8800 -p 3000:3000 \
-e FOO='bar' \
dariocalovic/secretify:0.2.3-beta