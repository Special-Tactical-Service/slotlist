Rem This file is for local development only!
Rem It configures and starts sts slotlist for local development.

set STS_SLOTLIST_LOGLEVEL=debug
set STS_SLOTLIST_WATCH_BUILD_JS=true
set STS_SLOTLIST_ALLOWED_ORIGINS=*
set STS_SLOTLIST_HOST=localhost:8888

go run main.go
