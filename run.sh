#!/bin/bash
go install github.com/goDB/single-server/server
go install github.com/goDB/single-server/client
server &
sleep 4
client localhost
