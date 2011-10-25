

# include $(GOROOT)/src/Make.inc


# TARG=prog

# GOFILES=\
# 	prog.go



# include $(GOROOT)/src/Make.cmd


PROGS = prog sepheaders channal_callback ftpget progress simple \
	post-callback sendrecv

ifeq ($(GOARCH), 386)
COMPILE.go = 8g
LINK.go = 8l

%: %.8
	$(LINK.go) -o $@ $^

%.8: %.go
	$(COMPILE.go) $<

else
ifeq ($(GOARCH), amd64)

COMPILE.go = 6g
LINK.go = 6l

%: %.6
	$(LINK.go) -o $@ $^

%.6: %.go
	$(COMPILE.go) $<

endif
endif

all: $(PROGS)

clean::
	rm -fv $(PROGS)
	rm -fv *.8
	rm -fv *.6
