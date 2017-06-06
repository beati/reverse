#!/bin/sh
go install -v -ldflags "-X github.com/beati/reverse/cmd.version=$(git describe --tags)" github.com/beati/reverse
err=$?
if [ $err -ne 0 ]; then
	exit $err
fi
