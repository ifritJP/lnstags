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
func Log_Level_get__allList(_env *LnsEnv) *LnsList{
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
func Log_Level__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Log_LevelMap_[arg1]; ok { return arg1 }
    return nil
}

func Log_Level_getTxt(arg1 LnsInt) string {
    return Log_LevelMap_[arg1];
}
var Log_name2levelMap *LnsMap2_[string,LnsInt]
var Log_outputLevel LnsInt
var Log_detail bool
type Log_CreateMessage func (_env *LnsEnv) string
// 24: decl @lns.@tags.@Log.str2level
func Log_str2level(_env *LnsEnv, txt string) LnsAny {
    return Log_name2levelMap.Get(txt)
}

// 29: decl @lns.@tags.@Log.setLevel
func Log_setLevel(_env *LnsEnv, level LnsInt) {
    Log_outputLevel = level
}

// 34: decl @lns.@tags.@Log.enableDetail
func Log_enableDetail(_env *LnsEnv, flag bool) {
    Log_detail = flag
}

// 42: decl @lns.@tags.@Log.log
func Log_log(_env *LnsEnv, level LnsInt,funcName string,lineNo LnsInt,callback Log_CreateMessage) {
    var logStream Lns_oStream
    logStream = Lns_io_stderr
    if level <= Log_outputLevel{
        if Log_detail{
            var nowClock LnsReal
            nowClock = _env.GetVM().OS_clock()
            logStream.Write(_env, _env.GetVM().String_format("%6d:%s:%s:%d:", Lns_2DDD((LnsInt)((nowClock * LnsReal(1000))), Log_Level_getTxt( level), funcName, lineNo)))
        } else { 
            logStream.Write(_env, _env.GetVM().String_format("%s:%s:", Lns_2DDD(Log_Level_getTxt( level), funcName)))
        }
        logStream.Write(_env, callback(_env))
        logStream.Write(_env, "\n")
    }
}

func Lns_Log_init(_env *LnsEnv) {
    if init_Log { return }
    init_Log = true
    Log__mod__ = "@lns.@tags.@Log"
    Lns_InitMod()
    Log_name2levelMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"fatal":Log_Level__Fatal,"error":Log_Level__Err,"warn":Log_Level__Warn,"log":Log_Level__Log,"info":Log_Level__Info,"debug":Log_Level__Debug,"trace":Log_Level__Trace,})
    Log_outputLevel = Log_Level__Warn
    Log_detail = true
}
func init() {
    init_Log = false
}
