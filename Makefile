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
GO    := go
CP    := cp
RM    := rm -f

GOROOT := $(shell echo $$GOROOT)
GOPATH := $(shell echo $$GOPATH)

.DEFAULT_GOAL = git-status
REPOS_TARGETS = git-status git-push git-commit-amend git-tag-list git-diff git-reset-soft \
				git-rm-cached git-branch
# -------------------------------------------------------------------------------------------------
.PHONY: clean
build:
	$(MAKE) -f Developers.mk $@

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
	$(GO) clean -testcache
	$(MAKE) -f Developers.mk $@

