LNSC=lnsc

SRCS=
SRCS += $(shell find lns -iname '*.lns')


help:
	@echo make build
	@echo make buildDB
	@echo make inq

build: lnstags

lnstags: $(SRCS)
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo --main lns.tags.main $(LNS_OPT)
	$(LNSC) lns/tags/main.lns mkmain entry.go
	go build -tags gopherlua

buildDB: build
#	rm -f lnstags.sqlite3
	./lnstags init
	./lnstags build test/main.lns
	$(MAKE) inq

inq: build
	./lnstags inq-def @test.@Sub.Bar.Val1
	./lnstags inq-def @test.@Sub.Hoge.func
	./lnstags inq-ref @test.@Sub.Hoge.func
	./lnstags inq-ref @test.@Sub.HogeHoge.__init

check: build
	rm -f lnstags.sqlite3
	./lnstags test


clean:
	rm -f lnstags.sqlite3 lnstags
