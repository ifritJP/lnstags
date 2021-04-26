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
	$(MAKE) check

inq: build
	./lnstags --simpleLog inq-at def test/main.lns 2 6 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 3 21 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 4 15 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 5 7 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 5 13 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 6 7 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 6 22 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 7 10 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 7 16 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 14 1 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 15 12 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 1 10 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 2 6 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 4 11 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 5 5 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 8 10 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 9 5 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 12 15 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 13 12 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 15 17 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 15 31 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 16 12 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 16 33 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 17 12 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 18 12 --log debug
	./lnstags --simpleLog inq-at ref test/Sub.lns 23 13 --log debug
	./lnstags --simpleLog inq def @test.@Sub.Bar.Val1
	./lnstags --simpleLog inq def @test.@Sub.Hoge.func
	./lnstags --simpleLog inq ref @test.@Sub.Hoge.func
	./lnstags --simpleLog inq ref @test.@Sub.HogeHoge.__init

check: build
	$(MAKE) inq > work/output.new.txt 2>&1
	diff test/output.txt work/output.new.txt


clean:
	rm -f lnstags.sqlite3 lnstags
