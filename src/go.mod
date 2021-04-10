module lnstags

go 1.16

require (
	github.com/ifritJP/LuneScript/src/lune v0.0.0-20210410094853-f0ef42d377ca
	github.com/ifritJP/lnssqlite3/src/lns/sqlite3 v0.0.0-20210410083954-26a3782ad99f
)

replace (
	github.com/ifritJP/LuneScript/src/lune => ../../LuneScript/src/lune
	github.com/ifritJP/lnssqlite3/src/lns/sqlite3 => ../../lnssqlite3/src/lns/sqlite3
)
