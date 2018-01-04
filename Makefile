.PHONY: all clean

all: bin/passphrase

clean:
	rm -Rv bin _build/ 2>/dev/null || true

bin/passphrase:
	mkdir bin 2>/dev/null || true
	jbuilder build cmd.bc
	cp -v _build/default/cmd.bc bin/passphrase
