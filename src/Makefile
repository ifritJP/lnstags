LNSC=lnsc

SRCS=
SRCS += $(shell find lns -iname '*.lns')


help:
	@echo make build
	@echo make buildDB
	@echo make check

build: lnstags

lnstags: $(SRCS)
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo --main lns.tags.main $(LNS_OPT)
	$(LNSC) lns/tags/main.lns mkmain entry.go
	go build -tags gopherlua

buildDB: build
	./lnstags init
	./lnstags build lns/tags/main.lns

inq: build
	./lnstags --simpleLog inq-at def test/main.lns 4 6 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 5 12 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 5 21 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 6 15 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 7 7 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 7 13 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 8 7 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 8 22 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 9 10 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 9 16 --log debug
	./lnstags --simpleLog inq-at ref test/main.lns 13 7 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 16 1 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 17 12 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 18 14 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 34 11 --log debug
	./lnstags --simpleLog inq-at def test/main.lns 34 22 --log debug
	./lnstags --simpleLog inq-at def test/mainSub.lns 3 6 --log debug
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
	./lnstags --simpleLog inq-at def test/Sub.lns 21 21 --log debug
	./lnstags --simpleLog inq-at def test/Sub.lns 24 16 --log debug
	./lnstags --simpleLog inq-at def test/Sub.lns 33 14 --log debug
	./lnstags --simpleLog inq-at def test/Sub.lns 41 9 --log debug
	./lnstags --simpleLog inq-at def test/Sub.lns 46 9 --log debug
	./lnstags --simpleLog inq-at def test/Sub.lns 54 9 --log debug
	cat test/Sub2.lns | ./lnstags --simpleLog inq-at ref -i test/Sub.lns 2 5 --log debug
	./lnstags --simpleLog inq def @test.@Sub.Bar.Val1
	./lnstags --simpleLog inq def @test.@Sub.Hoge.func
	./lnstags --simpleLog inq ref @test.@Sub.Hoge.func
	./lnstags --simpleLog inq ref @test.@Sub.HogeHoge.__init

check: build
	./lnstags init
	./lnstags build test/main.lns --log debug
	$(MAKE) inq 2>&1 | grep -v -e '^make\[' > work/output.new.txt 
	diff test/output.txt work/output.new.txt


clean:
	rm -f lnstags
