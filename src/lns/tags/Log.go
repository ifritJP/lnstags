// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Log bool
var Log__mod__ string
// decl enum -- Level 
type Log_Level = LnsInt
const Log_Level__Debug = 5
const Log_Level__Err = 1
const Log_Level__Fatal = 0
const Log_Level__Info = 4
const Log_Level__Log = 3
const Log_Level__Trace = 6
const Log_Level__Warn = 2
var Log_LevelList_ = NewLnsList( []LnsAny {
  Log_Level__Fatal,
  Log_Level__Err,
  Log_Level__Warn,
  Log_Level__Log,
  Log_Level__Info,
  Log_Level__Debug,
  Log_Level__Trace,
})
func Log_Level_get__allList() *LnsList{
    return Log_LevelList_
}
var Log_LevelMap_ = map[LnsInt]string {
  Log_Level__Debug: "Level.Debug",
  Log_Level__Err: "Level.Err",
  Log_Level__Fatal: "Level.Fatal",
  Log_Level__Info: "Level.Info",
  Log_Level__Log: "Level.Log",
  Log_Level__Trace: "Level.Trace",
  Log_Level__Warn: "Level.Warn",
}
func Log_Level__from(arg1 LnsInt) LnsAny{
    if _, ok := Log_LevelMap_[arg1]; ok { return arg1 }
    return nil
}

func Log_Level_getTxt(arg1 LnsInt) string {
    return Log_LevelMap_[arg1];
}
var Log_name2levelMap *LnsMap
var Log_outputLevel LnsInt
var Log_detail bool
var Log_logStream Lns_oStream
type Log_CreateMessage func () string
// 20: decl @lns.@tags.@Log.str2level
func Log_str2level(txt string) LnsAny {
    return Log_name2levelMap.Items[txt]
}

// 25: decl @lns.@tags.@Log.setLevel
func Log_setLevel(level LnsInt) {
    Log_outputLevel = level
    
}

// 30: decl @lns.@tags.@Log.enableDetail
func Log_enableDetail(flag bool) {
    Log_detail = flag
    
}

// 39: decl @lns.@tags.@Log.log
func Log_log(level LnsInt,funcName string,lineNo LnsInt,callback Log_CreateMessage) {
    if level <= Log_outputLevel{
        if Log_detail{
            var nowClock LnsReal
            nowClock = Lns_getVM().OS_clock()
            Log_logStream.Write(Lns_getVM().String_format("%6d:%s:%s:%d:", []LnsAny{(LnsInt)((nowClock * LnsReal(1000))), Log_Level_getTxt( level), funcName, lineNo}))
        } else { 
            Log_logStream.Write(Lns_getVM().String_format("%s:%s:", []LnsAny{Log_Level_getTxt( level), funcName}))
        }
        Log_logStream.Write(callback())
        Log_logStream.Write("\n")
    }
}

func Lns_Log_init() {
    if init_Log { return }
    init_Log = true
    Log__mod__ = "@lns.@tags.@Log"
    Lns_InitMod()
    Log_name2levelMap = NewLnsMap( map[LnsAny]LnsAny{})
    Log_name2levelMap.Set("fatal",Log_Level__Fatal)
    Log_name2levelMap.Set("error",Log_Level__Err)
    Log_name2levelMap.Set("warn",Log_Level__Warn)
    Log_name2levelMap.Set("log",Log_Level__Log)
    Log_name2levelMap.Set("info",Log_Level__Info)
    Log_name2levelMap.Set("debug",Log_Level__Debug)
    Log_name2levelMap.Set("trace",Log_Level__Trace)
    Log_outputLevel = Log_Level__Log
    Log_detail = true
    Log_logStream = Lns_io_stderr
}
func init() {
    init_Log = false
}
