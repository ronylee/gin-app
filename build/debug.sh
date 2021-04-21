#!/bin/sh

reset

cd ..

rm debug

go build -o debug

# export GIN_MODE=release

./debug