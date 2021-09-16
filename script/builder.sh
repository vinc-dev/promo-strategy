#!/bin/bash

go test -p 1 -cover -coverprofile=c.out ./internal/domain/*
go build -o app ./cmd
