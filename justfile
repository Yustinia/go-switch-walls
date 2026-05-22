entryPoint := "./cmd"
exeName := "wally"

# list recipes
default:
    just --list

run:
    go run {{ entryPoint }}

build:
    go build -o {{ exeName }} {{ entryPoint }}

clean:
    rm -rv {{ exeName }}

install: build
    mv -v {{ exeName }}
