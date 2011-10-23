

# include $(GOROOT)/src/Make.inc


# TARG=prog

# GOFILES=\
# 	prog.go



# include $(GOROOT)/src/Make.cmd


PROGS = prog sepheaders channal_callback ftpget progress simple \
	post-callback

COMPILE.go = 6g
LINK.go = 6l

# .go.6:
# 	$(COMPILE.go) $(OUTPUT_OPTION) $<

# .6.6:
# 	$(LINK.go) $(OUTPUT_OPTION) $<


%: %.6
	$(LINK.go) -o $@ $^

%.6: %.go
	$(COMPILE.go) $<


all: $(PROGS)


clean::
	rm -fv $(PROGS)
	rm -fv *.6
