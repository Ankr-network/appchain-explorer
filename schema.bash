#!/bin/bash

PG_URL="postgres://postgres:@rpc.dev-01.bas.ankr.com:7432/blockscout?sslmode=disable"

#mkdir -p shared/entity
#xo schema "$PG_URL" -o ./shared/entity

xo query "$PG_URL" --append --trim --strip -o shared/entity -T Block <<ENDSQL
SELECT * FROM "blocks"
ORDER BY "timestamp" DESC
LIMIT %%limit uint64%%
ENDSQL

xo query "$PG_URL" --append --trim --strip -o shared/entity -T Block <<ENDSQL
SELECT * FROM "blocks"
WHERE %%number uint64%% >= number
ORDER BY "timestamp" DESC
LIMIT %%limit uint64%%
ENDSQL