#!/bin/bash
set -e
set -x

# TODO ikke prodklart enda, men det kommer!

java \
-server \
-classpath "${APP_DIR}/WEB-INF/lib/*:${APP_DIR}/WEB-INF/classes" \
Main