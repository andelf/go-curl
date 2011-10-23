
include $(GOROOT)/src/Make.inc

CGO_CFLAGS=$(shell curl-config --cflags)
CGO_LDFLAGS+=$(shell curl-config --libs)


TARG=curl

CGOFILES=\
	core.go \
	easy.go \
	multi.go \
	share.go \
	const.go \
	callback.go \


CGO_OFILES=\
	c-callback.o \


include $(GOROOT)/src/Make.pkg
