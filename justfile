entryPoint := "./cmd"

# list recipes
default:
    just --list

run:
    go run {{ entryPoint }}

build:
    go build -o "gsw" {{ entryPoint }}

clean:
    rm -rv "gsw"
