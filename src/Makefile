LNSC=lnsc

help:
	@echo make test-lns

test-lns:
	$(LNSC) lns/tags/main.lns save -langGo --main lns.tags.main $(LNS_OPT)
	$(LNSC) lns/tags/main.lns mkmain entry.go
	go build
	rm -f hoge.sqlite3
	./lnstags
