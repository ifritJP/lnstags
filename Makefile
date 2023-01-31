LNSC=lnsc

SRCS=
SRCS += $(shell find lns -iname '*.lns')

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif


.PHONY: lnstags


help:
	@echo make build [LNS_OPT=opt]
	@echo make buildDB
	@echo make check
	@echo make update_lns

build: lnstags

lnstags: $(SRCS)
ifndef ONLY_GO
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo --main lns.tags.main $(LNS_OPT)
	$(LNSC) lns/tags/main.lns mkmain entry.go
endif
	go build -tags gopherlua -o lnstags

update_lns:
	go get github.com/ifritJP/LuneScript/src@latest


buildDB: build
	./lnstags init
	./lnstags build lns/tags/main.lns

define inq-test
	(cd test; ../lnstags --simpleLog --log debug $1)
endef

inq: build
	$(call inq-test,inq-at def main.lns 3 9)
	$(call inq-test,inq-at ref main.lns 3 9)
	$(call inq-test,inq-at def main.lns 4 1)
	$(call inq-test,inq-at def main.lns 4 6)
	$(call inq-test,inq-at def main.lns 5 12)
	$(call inq-test,inq-at def main.lns 5 21)
	$(call inq-test,inq-at def main.lns 6 16)
	$(call inq-test,inq-at def main.lns 7 7)
	$(call inq-test,inq-at def main.lns 7 13)
	$(call inq-test,inq-at def main.lns 8 7)
	$(call inq-test,inq-at def main.lns 8 22)
	$(call inq-test,inq-at def main.lns 9 10)
	$(call inq-test,inq-at def main.lns 9 16)
	$(call inq-test,inq-at ref main.lns 13 7)
	$(call inq-test,inq-at def main.lns 16 1)
	$(call inq-test,inq-at def main.lns 17 12)
	$(call inq-test,inq-at def main.lns 18 14)
	$(call inq-test,inq-at def main.lns 34 11)
	$(call inq-test,inq-at def main.lns 34 22)
	$(call inq-test,inq-at def main.lns 41 23)
	$(call inq-test,inq-at def main.lns 52 15)
	$(call inq-test,inq-at def main.lns 52 28)
	$(call inq-test,inq-at def main.lns 57 17)

	$(call inq-test,inq def @main.Al2)
	$(call inq-test,inq ref @main.Al2)
	$(call inq-test,inq def @main.Al2.Val)
	$(call inq-test,inq ref @main.Al2.Val)


	$(call inq-test,inq-at def mainSub.lns 3 6)
	(cd test; cat mainSub.lns | ../lnstags --simpleLog inq-at def -i mainSub.lns 3 5 --log debug)

	$(call inq-test,inq-at ref Sub.lns 1 10)
	$(call inq-test,inq-at ref Sub.lns 2 6)
	$(call inq-test,inq-at ref Sub.lns 4 11)
	$(call inq-test,inq-at ref Sub.lns 5 5)
	$(call inq-test,inq-at ref Sub.lns 8 10)
	$(call inq-test,inq-at ref Sub.lns 9 5)
	$(call inq-test,inq-at ref Sub.lns 12 15)
	$(call inq-test,inq-at ref Sub.lns 13 12)
	$(call inq-test,inq-at ref Sub.lns 15 17)
	$(call inq-test,inq-at ref Sub.lns 15 31)
	$(call inq-test,inq-at ref Sub.lns 16 12)
	$(call inq-test,inq-at ref Sub.lns 16 33)
	$(call inq-test,inq-at ref Sub.lns 17 16)
	$(call inq-test,inq-at set Sub.lns 17 16)
	$(call inq-test,inq-at ref Sub.lns 18 16)
	$(call inq-test,inq-at def Sub.lns 21 21)
	$(call inq-test,inq-at def Sub.lns 24 16)
	$(call inq-test,inq-at def Sub.lns 33 14)
	$(call inq-test,inq-at def Sub.lns 41 9)
	$(call inq-test,inq-at def Sub.lns 46 9)
	$(call inq-test,inq-at def Sub.lns 54 9)
	$(call inq-test,inq-at ref Sub.lns 74 13)
	$(call inq-test,inq-at ref Sub.lns 75 12)
	(cd test; cat Sub2.lns | ../lnstags --simpleLog inq-at ref -i Sub.lns 2 5 --log debug)
	$(call inq-test,inq def @Sub.Bar.Val1)
	$(call inq-test,inq def @Sub.Hoge.func)
	$(call inq-test,inq ref @Sub.Hoge.func)
	$(call inq-test,inq ref @Sub.HogeHoge.__init)

	$(call inq-test,inq allmut)
	$(call inq-test,inq async)
	$(call inq-test,inq noasync)

	(cd test/; ../lnstags --simpleLog suffix val --log debug)


check: build
	(cd test;../lnstags init)
	(cd test;../lnstags build main.lns --log debug)
	mkdir -p work
	$(MAKE) inq 2>&1 | grep -v -e '^make\[' -e ' -o lnstags' | \
			sed 's/: done .* msec//g' > work/output.new.txt 
	diff -I ': done' -I ': waiting...'  test/output.txt work/output.new.txt


clean:
	rm -f lnstags
