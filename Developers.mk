# go-sisimai/Developers.mk
#  ____                 _                                       _    
# |  _ \  _____   _____| | ___  _ __   ___ _ __ ___   _ __ ___ | | __
# | | | |/ _ \ \ / / _ \ |/ _ \| '_ \ / _ \ '__/ __| | '_ ` _ \| |/ /
# | |_| |  __/\ V /  __/ | (_) | |_) |  __/ |  \__ \_| | | | | |   < 
# |____/ \___| \_/ \___|_|\___/| .__/ \___|_|  |___(_)_| |_| |_|_|\_\
#                              |_|                                   
# -------------------------------------------------------------------------------------------------
SHELL := /bin/sh
HERE  := $(shell pwd)
NAME  := sisimai
MKDIR := mkdir -p
LS    := ls -1
CP    := cp

# -------------------------------------------------------------------------------------------------
.PHONY: clean

count-test-cases:
	@ go test -v `find sisimai -type f -name '*_test.go' | xargs dirname | sort | uniq` | \
		grep 'The number of ' | awk '{ cx += $$7 } END { print cx }'

loc:
	@ find . -type f -name '*.go' | grep -vE '/(tmp|sbin|internal)/' | grep -v '_test.go' | \
		xargs grep -vE '(^$$|^//|/[*]|[*]/|^ |^--)' | grep -vE "\t+//" | wc -l

clean:

