# libsisimai.org/sisimai/Developers.mk
#  ____                 _                                       _    
# |  _ \  _____   _____| | ___  _ __   ___ _ __ ___   _ __ ___ | | __
# | | | |/ _ \ \ / / _ \ |/ _ \| '_ \ / _ \ '__/ __| | '_ ` _ \| |/ /
# | |_| |  __/\ V /  __/ | (_) | |_) |  __/ |  \__ \_| | | | | |   < 
# |____/ \___| \_/ \___|_|\___/| .__/ \___|_|  |___(_)_| |_| |_|_|\_\
#                              |_|                                   
# -------------------------------------------------------------------------------------------------
SHELL := /bin/sh
HERE  := $(shell pwd)
FILE  := $(firstword $(MAKEFILE_LIST))
NAME  := sisimai
MKDIR := mkdir -p
LS    := ls -1
RM    := rm -f
CP    := cp
GO    := go

GOROOT := $(shell echo $$GOROOT)
GOPATH := $(shell echo $$GOPATH)

LIBSISIMAI := libsisimai.org
SISIMAIDIR := address arf fact lda lhost mail message moji reason rfc1123 rfc1894 rfc2045 rfc3464 \
			  rfc3834 rfc5322 rfc5965 rfc791 rhost sis smtp/*/
COVERAGETO := coverage.txt
EXECUTABLE := bin/sisid
BUILDFLAGS := -ldflags="-s -w" -trimpath
LISTENADDR := 127.0.0.1:5321

# -------------------------------------------------------------------------------------------------
.PHONY: clean
$(EXECUTABLE):
	CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o $@ $@.go

build:
	$(RM) $(EXECUTABLE)
	$(MAKE) -f $(FILE) $(EXECUTABLE)

test:
	@ $(GO) test ./ $(addprefix ./, $(SISIMAIDIR))

list-test-files:
	@ find $(SISIMAIDIR) -type f -name '*_test.go'

count-test-cases:
	@ $(GO) test -v ./ $(addprefix ./, $(SISIMAIDIR)) | grep 'The number of ' | awk '{ cx += $$7 } END { print cx }'

loc:
	@ find ./*.go $(SISIMAIDIR) -type f -name '*.go' -not -name '*_test.go' | \
		xargs grep -vE '(^$$|^//|/[*]|[*]/|^ |^--)' | grep -vE "\t+//" | wc -l

coverage:
	@ $(GO) test -v ./ $(addprefix ./, $(SISIMAIDIR)) -coverprofile=$(COVERAGETO)

init:
	test -e ./go.mod || $(GO) mod init $(LIBSISIMAI)/$(NAME)

update-go-mod:
	@ $(GO) mod tidy

start-godoc-server:
	open http://$(LISTENADDR)
	godoc -http=$(LISTENADDR)

clean:
	$(RM) ./$(EXECUTABLE)
	$(RM) ./$(COVERAGETO)

