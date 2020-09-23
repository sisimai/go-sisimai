# go-sisimai/Makefile
#  __  __       _         __ _ _      
# |  \/  | __ _| | _____ / _(_) | ___ 
# | |\/| |/ _` | |/ / _ \ |_| | |/ _ \
# | |  | | (_| |   <  __/  _| | |  __/
# |_|  |_|\__,_|_|\_\___|_| |_|_|\___|
# -----------------------------------------------------------------------------
SHELL := /bin/sh
TIME  := $(shell date '+%F')
NAME  := sisimai
WGET  := wget -c
CURL  := curl -L
CHMOD := chmod
GO    := go
CP    := cp
RM    := rm -f

GOROOT := $(shell echo $$GOROOT)
GOPATH := $(shell echo $$GOPATH)
DOMAIN := libsisimai.org

# -----------------------------------------------------------------------------
.PHONY: clean

git-status:
	git status

format:
	@ for v in `find . -type f -name '*.go' -not -path '*/tmp/*'`; do \
		echo $$v; \
		$(GOROOT)/bin/gofmt -w $$v; \
	done

init:
	test -e $(NAME)/go.mod || cd ./$(NAME) && $(GO) mod init $(NAME)

build:
	$(GO) build $(NAME).go

clean:
	:
