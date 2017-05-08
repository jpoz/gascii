BINARY=gascii

VERSION=1.1.0

REPO=github.com/jpoz/gascii

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '**/*.go')
.PHONY: gascii install clean assets

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build -o ${BINARY} cmd/gascii/gascii.go

install:
	go install ${LDFLAGS} ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

assets:
	go-bindata -pkg gascii inpact.ttf
