.POSIX:
.SUFFIXES:

PKGNAME=qol-assist
MODULE=github.com/getsolus/qol-assist

VERSION="0.5.0"

PREFIX?=/usr/local
BINDIR?=$(DESTDIR)$(PREFIX)/bin
SYSDIR?=$(DESTDIR)/etc/$(PKGNAME).d
USRDIR?=$(DESTDIR)$(PREFIX)/share/default/$(PKGNAME).d
TRACKDIR?=$(DESTDIR)/var/lib/$(PKGNAME)
GO?=go
GOFLAGS?=

GOSRC!=find . -name '*.go'
GOSRC+=go.mod go.sum

qol-assist: $(GOSRC)
	$(GO) build $(GOFLAGS) \
		-ldflags " \
		-X $(MODULE)/cli.VersionNumber=$(VERSION) \
		-X $(MODULE)/core.TrackDir=$(TRACKDIR) \
		-X $(MODULE)/core.SysDir=$(SYSDIR) \
		-X $(MODULE)/core.UsrDir=$(USRDIR)" \
		-o $@

all: qol-assist

# Exists in GNUMake but not in NetBSD make and others.
RM?=rm -f

clean:
	$(GO) mod tidy
	$(RM) $(DOCS) $(PKGNAME) *.tar.gz
	$(RM) -r vendor

install: all
	mkdir -m755 -p $(BINDIR) $(USRDIR) $(SYSDIR) $(LOGDIR)
	install -m755 $(PKGNAME) $(BINDIR)/$(PKGNAME)

RMDIR_IF_EMPTY:=sh -c '\
if test -d $$0 && ! ls -1qA $$0 | grep -q . ; then \
	rmdir $$0; \
fi'

uninstall:
	$(RM) $(BINDIR)/$(PKGNAME)
	$(RM) -r $(LOGDIR)
	$(RM) -r $(SYSDIR)
	$(RM) -r $(USRDIR)
	$(RMDIR_IF_EMPTY) $(BINDIR)

check:
	$(GO) get -u golang.org/x/lint/golint
	$(GO) get -u github.com/securego/gosec/cmd/gosec
	$(GO) get -u honnef.co/go/tools/cmd/staticcheck
	$(GO) get -u gitlab.com/opennota/check/cmd/aligncheck
	$(GO) fmt -x ./...
	$(GO) vet ./...
	golint -set_exit_status `go list ./... | grep -v vendor`
	gosec -exclude=G204 ./...
	staticcheck ./...
	aligncheck ./...
	$(GO) test -cover ./...

vendor: check clean
	$(GO) mod vendor

package: vendor
	tar --exclude='.git' \
		--exclude='*.tar.gz' \
	       	--exclude='examples' \
	       	--exclude="tags" \
	       	--exclude=".vscode" \
	       	--exclude=".idea"
		--exclude="*~" \
		-zcvf $(PKGNAME)-v$(VERSION).tar.gz ../$(PKGNAME)

.DEFAULT_GOAL := all

.PHONY: all clean install uninstall check vendor package
