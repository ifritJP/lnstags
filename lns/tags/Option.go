// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LuneTypes "github.com/ifritJP/LuneScript/src/lune/base"
import LuneUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Option bool
var Option__mod__ string
// decl enum -- InqMode 
type Option_InqMode = string
const Option_InqMode__AllMut = "allmut"
const Option_InqMode__Async = "async"
const Option_InqMode__AsyncLock = "asyncLock"
const Option_InqMode__Def = "def"
const Option_InqMode__Luaval = "luaval"
const Option_InqMode__Noasync = "noasync"
const Option_InqMode__Ref = "ref"
const Option_InqMode__Set = "set"
var Option_InqModeList_ = NewLnsList( []LnsAny {
  Option_InqMode__Def,
  Option_InqMode__Ref,
  Option_InqMode__Set,
  Option_InqMode__AllMut,
  Option_InqMode__Async,
  Option_InqMode__Noasync,
  Option_InqMode__Luaval,
  Option_InqMode__AsyncLock,
})
func Option_InqMode_get__allList(_env *LnsEnv) *LnsList{
    return Option_InqModeList_
}
var Option_InqModeMap_ = map[string]string {
  Option_InqMode__AllMut: "InqMode.AllMut",
  Option_InqMode__Async: "InqMode.Async",
  Option_InqMode__AsyncLock: "InqMode.AsyncLock",
  Option_InqMode__Def: "InqMode.Def",
  Option_InqMode__Luaval: "InqMode.Luaval",
  Option_InqMode__Noasync: "InqMode.Noasync",
  Option_InqMode__Ref: "InqMode.Ref",
  Option_InqMode__Set: "InqMode.Set",
}
func Option_InqMode__from(_env *LnsEnv, arg1 string) LnsAny{
    if _, ok := Option_InqModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_InqMode_getTxt(arg1 string) string {
    return Option_InqModeMap_[arg1];
}
// decl enum -- Mode 
type Option_Mode = string
const Option_Mode__Build = "build"
const Option_Mode__Dump = "dump"
const Option_Mode__Init = "init"
const Option_Mode__Inq = "inq"
const Option_Mode__InqAt = "inq-at"
const Option_Mode__Suffix = "suffix"
const Option_Mode__Test = "test"
const Option_Mode__Update = "update"
var Option_ModeList_ = NewLnsList( []LnsAny {
  Option_Mode__Init,
  Option_Mode__Build,
  Option_Mode__Update,
  Option_Mode__Inq,
  Option_Mode__InqAt,
  Option_Mode__Suffix,
  Option_Mode__Dump,
  Option_Mode__Test,
})
func Option_Mode_get__allList(_env *LnsEnv) *LnsList{
    return Option_ModeList_
}
var Option_ModeMap_ = map[string]string {
  Option_Mode__Build: "Mode.Build",
  Option_Mode__Dump: "Mode.Dump",
  Option_Mode__Init: "Mode.Init",
  Option_Mode__Inq: "Mode.Inq",
  Option_Mode__InqAt: "Mode.InqAt",
  Option_Mode__Suffix: "Mode.Suffix",
  Option_Mode__Test: "Mode.Test",
  Option_Mode__Update: "Mode.Update",
}
func Option_Mode__from(_env *LnsEnv, arg1 string) LnsAny{
    if _, ok := Option_ModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_Mode_getTxt(arg1 string) string {
    return Option_ModeMap_[arg1];
}
type Option_analyzeArgs__getNextOpNonNilFunc_0_ func (_env *LnsEnv, arg1 string) string
// 62: decl @lns.@tags.@Option.printUsage
func Option_printUsage_4_(_env *LnsEnv, messages LnsAny) {
    if messages != nil{
        messages_72 := messages.(string)
        Lns_io_stderr.Write(_env, _env.GetVM().String_format("%s\n", Lns_2DDD(messages_72)))
    }
    Lns_print(Lns_2DDD("usage: lnstags init [option]"))
    Lns_print(Lns_2DDD("usage: lnstags build [option] filepath"))
    Lns_print(Lns_2DDD("usage: lnstags inq <def|ref|set|allmut> pattern"))
    Lns_print(Lns_2DDD("usage: lnstags inq-at <def|ref|set> filepath lineno column"))
    Lns_print(Lns_2DDD("usage: lnstags suffix pattern"))
    Lns_print(Lns_2DDD("usage: lnstags test [option]"))
    _env.GetVM().OS_exit(1)
}

// 75: decl @lns.@tags.@Option.analyzeArgs
func Option_analyzeArgs(_env *LnsEnv, argList *LnsList) *Option_Option {
    var option *Option_Option
    option = NewOption_Option(_env)
    var index LnsInt
    index = 1
    var getNextOpNonNil Option_analyzeArgs__getNextOpNonNilFunc_0_
    getNextOpNonNil = Option_analyzeArgs__getNextOpNonNilFunc_0_(Option_analyzeArgs___anonymous_1_)
    var stdinFlag bool
    stdinFlag = false
    var Option_getNextOpRaw func(_env *LnsEnv) LnsAny
    Option_getNextOpRaw = func(_env *LnsEnv) LnsAny {
        if argList.Len() <= index{
            return nil
        }
        index = index + 1
        return argList.GetAt(index).(string)
    }
    var Option_getNextOp func(_env *LnsEnv) LnsAny
    Option_getNextOp = func(_env *LnsEnv) LnsAny {
        for  {
            var arg string
            
            {
                _arg := Option_getNextOpRaw(_env)
                if _arg == nil{
                    return nil
                } else {
                    arg = _arg.(string)
                }
            }
            if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(arg,"^-", nil, nil))){
                if _switch0 := arg; _switch0 == "-i" {
                    stdinFlag = true
                } else if _switch0 == "--log" {
                    option.logLevel = Log_str2level(_env, getNextOpNonNil(_env, "logLevel"))
                } else if _switch0 == "--simpleLog" {
                    Log_enableDetail(_env, false)
                } else if _switch0 == "--legacy-mutable-control" {
                    option.transCtrlInfo.LegacyMutableControl = true
                }
            } else { 
                return arg
            }
        }
    // insert a dummy
        return nil
    }
    getNextOpNonNil = Option_analyzeArgs__getNextOpNonNilFunc_0_(func(_env *LnsEnv, mess string) string {
        {
            _arg := Option_getNextOp(_env)
            if !Lns_IsNil( _arg ) {
                arg := _arg.(string)
                return arg
            }
        }
        Option_printUsage_4_(_env, mess)
    // insert a dummy
        return ""
    })
    var Option_getNextOpInt func(_env *LnsEnv, mess string) LnsInt
    Option_getNextOpInt = func(_env *LnsEnv, mess string) LnsInt {
        {
            _arg := Option_getNextOp(_env)
            if !Lns_IsNil( _arg ) {
                arg := _arg.(string)
                {
                    _num := Lns_tonumber(arg, nil)
                    if !Lns_IsNil( _num ) {
                        num := _num.(LnsReal)
                        return (LnsInt)(num)
                    }
                }
                Option_printUsage_4_(_env, _env.GetVM().String_format("illegal num -- %s", Lns_2DDD(arg)))
            }
        }
        Option_printUsage_4_(_env, mess)
    // insert a dummy
        return 0
    }
    var Option_getInqMode func(_env *LnsEnv) string
    Option_getInqMode = func(_env *LnsEnv) string {
        var nextToken string
        
        {
            _nextToken := Option_getNextOp(_env)
            if _nextToken == nil{
                Option_printUsage_4_(_env, "illegal inqMod")
            } else {
                nextToken = _nextToken.(string)
            }
        }
        var inqMode string
        
        {
            _inqMode := Option_InqMode__from(_env, nextToken)
            if _inqMode == nil{
                Option_printUsage_4_(_env, _env.GetVM().String_format("illegal inqMod -- %s", Lns_2DDD(nextToken)))
            } else {
                inqMode = _inqMode.(string)
            }
        }
        return inqMode
    }
    var mode LnsAny
    mode = nil
    for  {
        var arg string
        
        {
            _arg := Option_getNextOp(_env)
            if _arg == nil{
                break
            } else {
                arg = _arg.(string)
            }
        }
        if Lns_op_not(mode){
            {
                _work := Option_Mode__from(_env, arg)
                if !Lns_IsNil( _work ) {
                    work := _work.(string)
                    mode = work
                    if _switch1 := mode; _switch1 == Option_Mode__Inq {
                        option.inqMode = Option_getInqMode(_env)
                        if _switch0 := option.inqMode; _switch0 == Option_InqMode__AllMut || _switch0 == Option_InqMode__Async || _switch0 == Option_InqMode__Noasync || _switch0 == Option_InqMode__Luaval || _switch0 == Option_InqMode__AsyncLock {
                        } else {
                            option.pattern = getNextOpNonNil(_env, "none pattern")
                        }
                        Log_setLevel(_env, Log_Level__Warn)
                    } else if _switch1 == Option_Mode__InqAt {
                        option.inqMode = Option_getInqMode(_env)
                        option.analyzeFileInfo.path = getNextOpNonNil(_env, "none path")
                        option.analyzeFileInfo.lineNo = Option_getNextOpInt(_env, "none lineno")
                        option.analyzeFileInfo.column = Option_getNextOpInt(_env, "none column")
                    } else if _switch1 == Option_Mode__Suffix {
                        option.pattern = getNextOpNonNil(_env, "none pattern")
                    }
                } else {
                    Option_printUsage_4_(_env, _env.GetVM().String_format("illegal option -- %s", Lns_2DDD(arg)))
                }
            }
        } else { 
            if _switch2 := mode; _switch2 == Option_Mode__Build {
                if arg == "@-"{
                    for  {
                        var line string
                        
                        {
                            _line := Lns_io_stdin.Read(_env, "*l")
                            if _line == nil{
                                break
                            } else {
                                line = _line.(string)
                            }
                        }
                        if len(line) > 0{
                            option.pathList.Insert(line)
                        }
                    }
                } else { 
                    option.pathList.Insert(arg)
                }
            }
        }
    }
    {
        _logLevel := option.logLevel
        if !Lns_IsNil( _logLevel ) {
            logLevel := _logLevel.(LnsInt)
            Log_setLevel(_env, logLevel)
        }
    }
    if stdinFlag{
        option.analyzeFileInfo.stdinFile = LuneTypes.NewTypes_StdinFile(_env, LuneUtil.Util_scriptPath2Module(_env, option.analyzeFileInfo.FP.Get_path(_env)), Lns_unwrap( Lns_io_stdin.Read(_env, "*a")).(string))
    }
    if mode != nil{
        mode_149 := mode.(string)
        option.mode = mode_149
        return option
    }
    Option_printUsage_4_(_env, "none mode")
// insert a dummy
    return nil
}

func Option_analyzeArgs___anonymous_1_(_env *LnsEnv, mess string) string {
    Option_printUsage_4_(_env, "")
// insert a dummy
    return ""
}





// declaration Class -- AnalyzeFileInfo
type Option_AnalyzeFileInfoMtd interface {
    Get_column(_env *LnsEnv) LnsInt
    Get_lineNo(_env *LnsEnv) LnsInt
    Get_path(_env *LnsEnv) string
    Get_stdinFile(_env *LnsEnv) LnsAny
}
type Option_AnalyzeFileInfo struct {
    path string
    lineNo LnsInt
    column LnsInt
    stdinFile LnsAny
    FP Option_AnalyzeFileInfoMtd
}
func Option_AnalyzeFileInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_AnalyzeFileInfo).FP
}
func Option_AnalyzeFileInfo_toSlice(slice []LnsAny) []*Option_AnalyzeFileInfo {
    ret := make([]*Option_AnalyzeFileInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(Option_AnalyzeFileInfoDownCast).ToOption_AnalyzeFileInfo()
    }
    return ret
}
type Option_AnalyzeFileInfoDownCast interface {
    ToOption_AnalyzeFileInfo() *Option_AnalyzeFileInfo
}
func Option_AnalyzeFileInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_AnalyzeFileInfoDownCast)
    if ok { return work.ToOption_AnalyzeFileInfo() }
    return nil
}
func (obj *Option_AnalyzeFileInfo) ToOption_AnalyzeFileInfo() *Option_AnalyzeFileInfo {
    return obj
}
func NewOption_AnalyzeFileInfo(_env *LnsEnv) *Option_AnalyzeFileInfo {
    obj := &Option_AnalyzeFileInfo{}
    obj.FP = obj
    obj.InitOption_AnalyzeFileInfo(_env)
    return obj
}
func (self *Option_AnalyzeFileInfo) Get_path(_env *LnsEnv) string{ return self.path }
func (self *Option_AnalyzeFileInfo) Get_lineNo(_env *LnsEnv) LnsInt{ return self.lineNo }
func (self *Option_AnalyzeFileInfo) Get_column(_env *LnsEnv) LnsInt{ return self.column }
func (self *Option_AnalyzeFileInfo) Get_stdinFile(_env *LnsEnv) LnsAny{ return self.stdinFile }
// 34: DeclConstr
func (self *Option_AnalyzeFileInfo) InitOption_AnalyzeFileInfo(_env *LnsEnv) {
    self.path = ""
    self.lineNo = 0
    self.column = 0
    self.stdinFile = nil
}


// declaration Class -- Option
type Option_OptionMtd interface {
    Get_analyzeFileInfo(_env *LnsEnv) *Option_AnalyzeFileInfo
    Get_inqMode(_env *LnsEnv) string
    Get_logLevel(_env *LnsEnv) LnsAny
    Get_mode(_env *LnsEnv) string
    Get_pathList(_env *LnsEnv) *LnsList2_[string]
    Get_pattern(_env *LnsEnv) string
    Get_transCtrlInfo(_env *LnsEnv) *LuneTypes.Types_TransCtrlInfo
}
type Option_Option struct {
    pathList *LnsList2_[string]
    mode string
    inqMode string
    pattern string
    analyzeFileInfo *Option_AnalyzeFileInfo
    logLevel LnsAny
    transCtrlInfo *LuneTypes.Types_TransCtrlInfo
    FP Option_OptionMtd
}
func Option_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_Option).FP
}
func Option_Option_toSlice(slice []LnsAny) []*Option_Option {
    ret := make([]*Option_Option, len(slice))
    for index, val := range slice {
        ret[index] = val.(Option_OptionDownCast).ToOption_Option()
    }
    return ret
}
type Option_OptionDownCast interface {
    ToOption_Option() *Option_Option
}
func Option_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_OptionDownCast)
    if ok { return work.ToOption_Option() }
    return nil
}
func (obj *Option_Option) ToOption_Option() *Option_Option {
    return obj
}
func NewOption_Option(_env *LnsEnv) *Option_Option {
    obj := &Option_Option{}
    obj.FP = obj
    obj.InitOption_Option(_env)
    return obj
}
func (self *Option_Option) Get_pathList(_env *LnsEnv) *LnsList2_[string]{ return self.pathList }
func (self *Option_Option) Get_mode(_env *LnsEnv) string{ return self.mode }
func (self *Option_Option) Get_inqMode(_env *LnsEnv) string{ return self.inqMode }
func (self *Option_Option) Get_pattern(_env *LnsEnv) string{ return self.pattern }
func (self *Option_Option) Get_analyzeFileInfo(_env *LnsEnv) *Option_AnalyzeFileInfo{ return self.analyzeFileInfo }
func (self *Option_Option) Get_logLevel(_env *LnsEnv) LnsAny{ return self.logLevel }
func (self *Option_Option) Get_transCtrlInfo(_env *LnsEnv) *LuneTypes.Types_TransCtrlInfo{ return self.transCtrlInfo }
// 51: DeclConstr
func (self *Option_Option) InitOption_Option(_env *LnsEnv) {
    self.logLevel = nil
    self.pathList = NewLnsList2_[string]([]string{})
    self.mode = Option_Mode__Build
    self.inqMode = Option_InqMode__Def
    self.pattern = ""
    self.analyzeFileInfo = NewOption_AnalyzeFileInfo(_env)
    self.transCtrlInfo = LuneTypes.Types_TransCtrlInfo_create_normal(_env)
}


func Lns_Option_init(_env *LnsEnv) {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lns.@tags.@Option"
    Lns_InitMod()
    Lns_Log_init(_env)
    LuneTypes.Lns_Types_init(_env)
    LuneUtil.Lns_Util_init(_env)
}
func init() {
    init_Option = false
}
