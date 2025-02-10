# libsisimai.org/sisimai/Makefile
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

LIBSISIMAI := libsisimai.org
SISIMAIDIR := address arf fact lda lhost mail message reason rfc1123 rfc1894 rfc2045 rfc3464 \
			  rfc3834 rfc5322 rfc5965 rfc791 rhost sis smtp/command smtp/failure smtp/reply  \
			  smtp/status smtp/transcript string

.DEFAULT_GOAL = git-status
REPOS_TARGETS = git-status git-push git-commit-amend git-tag-list git-diff git-reset-soft \
				git-rm-cached git-branch
# -------------------------------------------------------------------------------------------------
.PHONY: clean
build:
	$(GO) build lib$(NAME).go

test:
	$(MAKE) -f Developers.mk $@

# -------------------------------------------------------------------------------------------------
$(REPOS_TARGETS):
	$(MAKE) -f Repository.mk $@

fix-commit-message:       git-commit-amend
cancel-the-latest-commit: git-reset-soft
remove-added-file:        git-rm-cached
diff push branch:
	@$(MAKE) git-$@

# -------------------------------------------------------------------------------------------------
clean:
	go clean -testcache

