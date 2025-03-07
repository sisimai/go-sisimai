# go-sisimai/set-of-emails/Makefile
#           _               __                            _ _        ____  __       _         __ _ _      
#  ___  ___| |_       ___  / _|       ___ _ __ ___   __ _(_) |___   / /  \/  | __ _| | _____ / _(_) | ___ 
# / __|/ _ \ __|____ / _ \| |_ _____ / _ \ '_ ` _ \ / _` | | / __| / /| |\/| |/ _` | |/ / _ \ |_| | |/ _ \
# \__ \  __/ ||_____| (_) |  _|_____|  __/ | | | | | (_| | | \__ \/ / | |  | | (_| |   <  __/  _| | |  __/
# |___/\___|\__|     \___/|_|        \___|_| |_| |_|\__,_|_|_|___/_/  |_|  |_|\__,_|_|\_\___|_| |_|_|\___|
# -------------------------------------------------------------------------------------------------
SHELL := /bin/sh
TIME  := $(shell date '+%F')
COPY  := cp -vRp

EMAILDIR := set-of-emails
MAILDIRS := $(EMAILDIR)/maildir

sync:
	test -d ../p5-sisimai
	$(COPY) ./$(MAILDIRS)/bsd/*.eml ../p5-sisimai/$(MAILDIRS)/bsd/
	$(COPY) ./$(MAILDIRS)/mac/*.eml ../p5-sisimai/$(MAILDIRS)/mac/
	$(COPY) ./$(MAILDIRS)/dos/*.eml ../p5-sisimai/$(MAILDIRS)/dos/
	$(COPY) ./$(EMAILDIR)/private/* ../p5-sisimai/$(EMAILDIR)/private/

	test -d ../rb-sisimai
	$(COPY) ./$(MAILDIRS)/bsd/*.eml ../rb-sisimai/$(MAILDIRS)/bsd/
	$(COPY) ./$(MAILDIRS)/mac/*.eml ../rb-sisimai/$(MAILDIRS)/mac/
	$(COPY) ./$(MAILDIRS)/dos/*.eml ../rb-sisimai/$(MAILDIRS)/dos/
	$(COPY) ./$(EMAILDIR)/private/* ../rb-sisimai/$(EMAILDIR)/private/

