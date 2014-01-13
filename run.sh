#!/bin/bash
go install github.com/goDB/single-server/server
go install github.com/goDB/single-server/client
server &
client localhost
