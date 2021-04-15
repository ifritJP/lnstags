module lnstags

go 1.16

require (
	github.com/ifritJP/LuneScript/src v0.0.0-20210415092331-6ecbb8565ddc
	github.com/ifritJP/lnssqlite3/src v0.0.0-20210415104505-1975737f9acd
)

replace (
	github.com/ifritJP/LuneScript/src => ../../LuneScript/src
	github.com/ifritJP/lnssqlite3/src => ../../lnssqlite3/src
)
