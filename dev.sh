#!/bin/bash

# This file is for local development only!
# It configures and starts sts slotlist for local development.

export STS_SLOTLIST_LOGLEVEL=debug
export STS_SLOTLIST_WATCH_BUILD_JS=true
export STS_SLOTLIST_ALLOWED_ORIGINS=*
export STS_SLOTLIST_HOST=localhost:8080

go run main.go
