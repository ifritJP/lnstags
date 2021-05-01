module lnstags

go 1.16

require (
	github.com/ifritJP/LuneScript/src v0.0.0-20210501015200-4593906b745d
	github.com/ifritJP/lnssqlite3/src v0.0.0-20210501053119-1b59ffd8b301
)

replace (
	github.com/ifritJP/LuneScript/src => ../../LuneScript/src
	github.com/ifritJP/lnssqlite3/src => ../../lnssqlite3/src
)
