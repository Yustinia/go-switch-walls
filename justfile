entryPoint := "./cmd"

# list recipes
default:
    just --list

run:
    go run {{ entryPoint }}

build:
    go build {{ entryPoint }}
