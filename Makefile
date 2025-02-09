# go-sisimai/Makefile
#  __  __       _         __ _ _      
# |  \/  | __ _| | _____ / _(_) | ___ 
# | |\/| |/ _` | |/ / _ \ |_| | |/ _ \
# | |  | | (_| |   <  __/  _| | |  __/
# |_|  |_|\__,_|_|\_\___|_| |_|_|\___|
# -------------------------------------------------------------------------------------------------
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
COVERS := coverage.txt

.DEFAULT_GOAL = git-status
REPOS_TARGETS = git-status git-push git-commit-amend git-tag-list git-diff git-reset-soft \
				git-rm-cached git-branch
# -------------------------------------------------------------------------------------------------
.PHONY: clean
init:
	test -e $(NAME)/go.mod || cd ./$(NAME) && $(GO) mod init $(NAME)

build:
	$(GO) build lib$(NAME).go

test:
	go test `find sisimai -type f -name '*_test.go' | xargs dirname | sort | uniq`

coverage:
	go test `find sisimai -type f -name '*_test.go' | xargs dirname | sort | uniq` -coverprofile=$(COVERS)

$(REPOS_TARGETS):
	$(MAKE) -f Repository.mk $@

fix-commit-message:       git-commit-amend
cancel-the-latest-commit: git-reset-soft
remove-added-file:        git-rm-cached
diff push branch:
	@$(MAKE) git-$@

clean:
	go clean -testcache

