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
MAILSUFFIX := eml
PUBLICFILE := set-of-emails
PRIVATESET := $(PUBLICFILE)/private
ASSEMBLEIN := tmp/assembled-in-here
PROFILESET := tmp/all-the-emails
EXECUTABLE := bin/sisid
BUILDFLAGS := -ldflags="-s -w" -trimpath
GOLANGLINT := golangci-lint
GO_SYSNAME := $(shell echo $$GOOS   || $(GO) env GOOS  )
GO_CPUARCH := $(shell echo $$GOARCH || $(GO) env GOARCH)
LISTENADDR := 127.0.0.1:5321
K          := neko

# -------------------------------------------------------------------------------------------------
.PHONY: clean
$(EXECUTABLE):
	CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o $@ $@.go

$(EXECUTABLE).$(GO_SYSNAME)-$(GO_CPUARCH):
	GOOS=$(GO_SYSNAME) GOARCH=$(GO_CPUARCH) CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o $@ $(EXECUTABLE).go

build:
	$(RM) $(EXECUTABLE)
	$(MAKE) -f $(FILE) $(EXECUTABLE)

cross-build:
	# https://go.dev/doc/install/source#environment
	test -n "$(GO_SYSNAME)"
	test -n "$(GO_CPUARCH)"
	$(RM) $(EXECUTABLE).$(GO_SYSNAME)-$(GO_CPUARCH)
	$(MAKE) -f $(FILE) $(EXECUTABLE).$(GO_SYSNAME)-$(GO_CPUARCH)

test:
	@ $(GO) test ./ $(addprefix ./, $(SISIMAIDIR))
	@ $(foreach v, $(shell make -f ./Developers.mk lhost-files), grep -Fq "$(v)" ./lhost/*_test.go || echo '❌ **** $(v) not registered' 1>&2;)
	@ $(foreach v, $(shell make -f ./Developers.mk other-files), grep -Fq "$(v)" ./rfc3*/*_test.go || echo '❌ **** $(v) not registered' 1>&2;)

lhost-files:
	@ $(LS) $(PUBLICFILE)/maildir/bsd/lhost-*.eml \
		| grep -vE 'lhost-(amavis|amazon|barracuda|bigfoot|domino|mailru|mcafee|mfilter|mxlogic|office365)' \
		| grep -vE 'lhost-(outlook|powermta|receivingses|sendgrid|surfcontrol|x4|x5|yahoo|yandex)'          \
		| xargs basename | sed -e 's/.eml//g'

other-files:
	@ $(LS) $(PUBLICFILE)/maildir/bsd/rfc*.eml | xargs basename | sed -e 's/.eml//g'

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

lint:
	test -x `which $(GOLANGLINT)`
	NO_COLOR=2 $(GOLANGLINT) run $(SISIMAIDIR)

samples:
	$(MKDIR) $(PROFILESET)
	$(CP) -p $(PUBLICFILE)/mailbox/mbox-* $(PROFILESET)/
	$(CP) -p $(PUBLICFILE)/maildir/bsd/*.$(MAILSUFFIX) $(PROFILESET)/
	find $(PRICATESET)/ -type f -name '*.$(MAILSUFFIX)' | xargs -I__EEF__ $(CP) -p __EEF__ $(PROFILESET)

private-sample:
	@test -n "$(E)" || ( echo 'Usage: make -f Developers.mk $@ E=/path/to/email' && exit 1 )
	@test -f $(EXECUTABLE).go
	@test -f $(E)
	$(GO) run $(EXECUTABLE).go $(E)
	@echo
	@while true; do \
		d=`$(GO) run $(EXECUTABLE).go -format json $(E) | jq -M '.decodedby' | head -1 \
			| tr '[A-Z]' '[a-z]' | tr -d '-' | sed -e 's/"//g' -e 's/^/lhost-/g'`; \
		if [ -d "$(PRIVATESET)/$$d" ]; then \
			thelatest=`ls -1 $(PRIVATESET)/$$d/*.$(MAILSUFFIX) | tail -1`; \
			currindex=`basename $$thelatest | cut -d'-' -f1`; \
			nextindex=`echo $$currindex + 1 | bc`; \
		else \
			$(MKDIR) $(PRIVATESET)/$$d; \
			nextindex=1001; \
		fi; \
		hashvalue=`md5 -q $(E) | cut -c 1-8`; \
		if [ -n "`ls -1 $(PRIVATESET)/$$d/ | grep $$hashvalue`" ]; then \
			echo 'Already exists:' `ls -1 $(PRIVATESET)/$$d/*$$hashvalue.$(MAILSUFFX)`; \
		else \
			printf "[%04d] %s %s\n" $$nextindex $$hashvalue; \
			mv -v $(E) $(PRIVATESET)/$$d/$${nextindex}-$${hashvalue}.$(MAILSUFFIX); \
		fi; \
		break; \
	done

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
		grep -Eh '^import ' `ls -1 ./smtp/$$v/*.go | grep -v _test.go` | sort | uniq       >> $(ASSEMBLEIN)/smtp/$$v/lib.go; \
		ls -1 smtp/$$v/*.go | grep -v _test.go | xargs cat | grep -vE '^(import|package) ' >> $(ASSEMBLEIN)/smtp/$$v/lib.go; \
	done

	for v in $(SISIMAIDIR); do \
		test -n "`echo $$v | grep 'smtp/'`" && continue; \
		$(MKDIR) $(ASSEMBLEIN)/$$v; \
		printf "package %s\n" $$v                                                      > $(ASSEMBLEIN)/$$v/lib.go; \
		grep -Eh '^import ' `ls -1 ./$$v/*.go | grep -v _test.go` | sort | uniq       >> $(ASSEMBLEIN)/$$v/lib.go; \
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

