LNSC=lnsc

SRCS=
SRCS += lns/tags/Log.lns
SRCS += lns/tags/DBAccess.lns
SRCS += lns/tags/DBCtrl.lns
SRCS += lns/tags/Analyze.lns
SRCS += lns/tags/main.lns


help:
	@echo make build

build: lnstags

lnstags: $(SRCS)
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo --main lns.tags.main $(LNS_OPT)
	$(LNSC) lns/tags/main.lns mkmain entry.go
	go build -tags gopherlua
	$(MAKE) exec

exec: build
	rm -f hoge.sqlite3
	./lnstags

clean:
	rm -f hoge.sqlite3 lnstags
