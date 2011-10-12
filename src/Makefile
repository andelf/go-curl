

include $(GOROOT)/src/Make.inc

CGO_CFLAGS=
CGO_LDFLAGS=`curl-config --libs`

TARG=curl

# GOFILES=
CGOFILES=\
	core.go \
	easy.go \
	multi.go \
	share.go \
	const.go \
	callback.go \


CGO_OFILES=\
	c-callback.o


include $(GOROOT)/src/Make.pkg
