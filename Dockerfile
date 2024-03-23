FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
