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
ASSEMBLEIN := tmp/assembled-in-here
PROFILESET := set-of-emails/maildir/bsd
EXECUTABLE := bin/sisid
BUILDFLAGS := -ldflags="-s -w" -trimpath
LISTENADDR := 127.0.0.1:5321
K          := neko

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
	@ find $(SISIMAIDIR) -type f -not -path '*/tmp/*' -name '*_test.go'

count-test-cases:
	@ $(GO) test -v ./ $(addprefix ./, $(SISIMAIDIR)) | grep 'The number of ' | awk '{ cx += $$7 } END { print cx }'

loc:
	@ find ./*.go $(SISIMAIDIR) -type f -name '*.go' -not -name '*_test.go' | \
		xargs grep -vE '(^$$|^//|/[*]|[*]/|^ |^--)' | grep -vE "\t+//" | wc -l

how-many-engines:
	@ echo `$(LS) lhost/via-* rhost/for-* | wc -l | tr -d ' '` + 4 | bc

coverage:
	@ $(GO) test -v ./ $(addprefix ./, $(SISIMAIDIR)) -coverprofile=$(COVERAGETO)

profile:
	test -f bin/cpu-prof.go && CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o cpu-sisid ./bin/cpu-prof.go
	test -f ./cpu-sisid     && ./cpu-sisid ./$(PROFILESET)
	go tool pprof --top ./cpu.pprof > usage-of-cpu-x

	test -f bin/mem-prof.go && CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o mem-sisid ./bin/mem-prof.go
	test -f ./mem-sisid     && ./mem-sisid ./$(PROFILESET)
	go tool pprof --top ./mem.pprof > usage-of-mem-x

	ls -laF ./usage-of-*

benchmark:
	test -f bin/benchmark.go && CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o min-sisid ./bin/benchmark.go
	uptime
	while true; do zsh -c 'time ./min-sisid $(PROFILESET)'; sleep 10; done

find:
	find . -type f -name '*.go' -not -name '*_test.go' -not -path '*/bin/*' -not -path '*/sbin/*' \
		-not -path '*/stash/*' -not -path '*/tmp/*' -exec grep '$(K)' {} +

assemble:
	$(MKDIR) $(ASSEMBLEIN)/smtp
	printf "package sisimai\n"                                      > $(ASSEMBLEIN)/libsisimai.go
	grep -Eh  '^import ' libsisimai.go interfaces.go | sort | uniq >> $(ASSEMBLEIN)/libsisimai.go
	grep -Ehv '^(import|package) ' interfaces.go libsisimai.go     >> $(ASSEMBLEIN)/libsisimai.go

	for v in `ls -1 smtp/`; do \
		$(MKDIR) $(ASSEMBLEIN)/smtp/$$v; \
		printf "package %s\n" $$v                                                           > $(ASSEMBLEIN)/smtp/$$v/lib.go; \
		grep '^import ' `ls -1 ./smtp/$$v/*.go | grep -v _test.go` | sort | uniq           >> $(ASSEMBLEIN)/smtp/$$v/lib.go; \
		ls -1 smtp/$$v/*.go | grep -v _test.go | xargs cat | grep -vE '^(import|package) ' >> $(ASSEMBLEIN)/smtp/$$v/lib.go; \
	done

	for v in $(SISIMAIDIR); do \
		test -n "`echo $$v | grep 'smtp/'`" && continue; \
		$(MKDIR) $(ASSEMBLEIN)/$$v; \
		printf "package %s\n" $$v                                                      > $(ASSEMBLEIN)/$$v/lib.go; \
		grep '^import ' `ls -1 ./$$v/*.go | grep -v _test.go` | sort | uniq           >> $(ASSEMBLEIN)/$$v/lib.go; \
		ls -1 $$v/*.go | grep -v _test.go | xargs cat | grep -vE '^(import|package) ' >> $(ASSEMBLEIN)/$$v/lib.go; \
	done

init:
	test -e ./go.mod || $(GO) mod init $(LIBSISIMAI)/$(NAME)

update-go-mod:
	@ $(GO) mod tidy

start-godoc-server:
	open http://$(LISTENADDR)
	godoc -http=$(LISTENADDR)

clean:
	$(RM)    ./$(EXECUTABLE)
	$(RM)    ./$(COVERAGETO)
	$(RM) -r ./$(ASSEMBLEIN)

