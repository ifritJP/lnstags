// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_Option bool
var Option__mod__ string
// decl enum -- InqMode 
type Option_InqMode = string
const Option_InqMode__Def = "def"
const Option_InqMode__Ref = "ref"
const Option_InqMode__Set = "set"
var Option_InqModeList_ = NewLnsList( []LnsAny {
  Option_InqMode__Def,
  Option_InqMode__Ref,
  Option_InqMode__Set,
})
func Option_InqMode_get__allList(_env *LnsEnv) *LnsList{
    return Option_InqModeList_
}
var Option_InqModeMap_ = map[string]string {
  Option_InqMode__Def: "InqMode.Def",
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
type analyzeArgs__getNextOpNonNilFunc_1076_ func (_env *LnsEnv, arg1 string) string
// 54: decl @lns.@tags.@Option.printUsage
func Option_printUsage_1068_(_env *LnsEnv, messages LnsAny) {
    if messages != nil{
        messages_63 := messages.(string)
        Lns_io_stderr.Write(_env, _env.LuaVM.String_format("%s\n", []LnsAny{messages_63}))
    }
    Lns_print([]LnsAny{"usage: lnstags init [option]"})
    Lns_print([]LnsAny{"usage: lnstags build [option] filepath"})
    Lns_print([]LnsAny{"usage: lnstags inq <def|ref> pattern"})
    Lns_print([]LnsAny{"usage: lnstags inq-at <def|ref> filepath lineno column"})
    Lns_print([]LnsAny{"usage: lnstags test [option]"})
    _env.LuaVM.OS_exit(1)
}

func analyzeArgs___anonymous_1078_(_env *LnsEnv, mess string) string {
    Option_printUsage_1068_(_env, "")
// insert a dummy
    return ""
}




// 66: decl @lns.@tags.@Option.analyzeArgs
func Option_analyzeArgs(_env *LnsEnv, argList *LnsList) *Option_Option {
    var option *Option_Option
    option = NewOption_Option(_env)
    var index LnsInt
    index = 1
    var getNextOpNonNil analyzeArgs__getNextOpNonNilFunc_1076_
    getNextOpNonNil = analyzeArgs__getNextOpNonNilFunc_1076_(analyzeArgs___anonymous_1078_)
    
    var getNextOp func(_env *LnsEnv) LnsAny
    getNextOp = func(_env *LnsEnv) LnsAny {
        for  {
            if argList.Len() <= index{
                return nil
            }
            index = index + 1
            
            var arg string
            arg = argList.GetAt(index).(string)
            if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(arg,"^-", nil, nil))){
                if _switch414 := arg; _switch414 == "-i" {
                    option.analyzeFileInfo.stdinFlag = true
                    
                } else if _switch414 == "--log" {
                    option.logLevel = Log_str2level(_env, getNextOpNonNil(_env, "logLevel"))
                    
                } else if _switch414 == "--simpleLog" {
                    Log_enableDetail(_env, false)
                } else if _switch414 == "--legacy-mutable-control" {
                    option.transCtrlInfo.LegacyMutableControl = true
                    
                }
            } else { 
                return arg
            }
        }
    // insert a dummy
        return nil
    }
    getNextOpNonNil = analyzeArgs__getNextOpNonNilFunc_1076_(func(_env *LnsEnv, mess string) string {
        {
            _arg := getNextOp(_env)
            if !Lns_IsNil( _arg ) {
                arg := _arg.(string)
                return arg
            }
        }
        Option_printUsage_1068_(_env, mess)
    // insert a dummy
        return ""
    })
    
    var getNextOpInt func(_env *LnsEnv, mess string) LnsInt
    getNextOpInt = func(_env *LnsEnv, mess string) LnsInt {
        {
            _arg := getNextOp(_env)
            if !Lns_IsNil( _arg ) {
                arg := _arg.(string)
                {
                    _num := Lns_tonumber(arg, nil)
                    if !Lns_IsNil( _num ) {
                        num := _num.(LnsReal)
                        return (LnsInt)(num)
                    }
                }
                Option_printUsage_1068_(_env, _env.LuaVM.String_format("illegal num -- %s", []LnsAny{arg}))
            }
        }
        Option_printUsage_1068_(_env, mess)
    // insert a dummy
        return 0
    }
    var getInqMode func(_env *LnsEnv) string
    getInqMode = func(_env *LnsEnv) string {
        var nextToken string
        
        {
            _nextToken := getNextOp(_env)
            if _nextToken == nil{
                Option_printUsage_1068_(_env, "illegal inqMod")
            } else {
                nextToken = _nextToken.(string)
            }
        }
        var inqMode string
        
        {
            _inqMode := Option_InqMode__from(_env, nextToken)
            if _inqMode == nil{
                Option_printUsage_1068_(_env, _env.LuaVM.String_format("illegal inqMod -- %s", []LnsAny{nextToken}))
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
            _arg := getNextOp(_env)
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
                    
                    if _switch697 := mode; _switch697 == Option_Mode__Inq {
                        option.inqMode = getInqMode(_env)
                        
                        option.pattern = getNextOpNonNil(_env, "none pattern")
                        
                        Log_setLevel(_env, Log_Level__Warn)
                    } else if _switch697 == Option_Mode__InqAt {
                        option.inqMode = getInqMode(_env)
                        
                        option.analyzeFileInfo.path = getNextOpNonNil(_env, "none path")
                        
                        option.analyzeFileInfo.lineNo = getNextOpInt(_env, "none lineno")
                        
                        option.analyzeFileInfo.column = getNextOpInt(_env, "none column")
                        
                    } else if _switch697 == Option_Mode__Suffix {
                        option.pattern = getNextOpNonNil(_env, "none pattern")
                        
                    }
                } else {
                    Option_printUsage_1068_(_env, _env.LuaVM.String_format("illegal option -- %s", []LnsAny{arg}))
                }
            }
        } else { 
            if _switch780 := mode; _switch780 == Option_Mode__Build {
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
    if mode != nil{
        mode_134 := mode.(string)
        option.mode = mode_134
        
        return option
    }
    Option_printUsage_1068_(_env, "none mode")
// insert a dummy
    return nil
}

// declaration Class -- AnalyzeFileInfo
type Option_AnalyzeFileInfoMtd interface {
    Get_column(_env *LnsEnv) LnsInt
    Get_lineNo(_env *LnsEnv) LnsInt
    Get_path(_env *LnsEnv) string
    Get_stdinFlag(_env *LnsEnv) bool
}
type Option_AnalyzeFileInfo struct {
    path string
    lineNo LnsInt
    column LnsInt
    stdinFlag bool
    FP Option_AnalyzeFileInfoMtd
}
func Option_AnalyzeFileInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_AnalyzeFileInfo).FP
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
func (self *Option_AnalyzeFileInfo) Get_stdinFlag(_env *LnsEnv) bool{ return self.stdinFlag }
// 26: DeclConstr
func (self *Option_AnalyzeFileInfo) InitOption_AnalyzeFileInfo(_env *LnsEnv) {
    self.path = ""
    
    self.lineNo = 0
    
    self.column = 0
    
    self.stdinFlag = false
    
}


// declaration Class -- Option
type Option_OptionMtd interface {
    Get_analyzeFileInfo(_env *LnsEnv) *Option_AnalyzeFileInfo
    Get_inqMode(_env *LnsEnv) string
    Get_logLevel(_env *LnsEnv) LnsAny
    Get_mode(_env *LnsEnv) string
    Get_pathList(_env *LnsEnv) *LnsList
    Get_pattern(_env *LnsEnv) string
    Get_transCtrlInfo(_env *LnsEnv) *LnsTypes.Types_TransCtrlInfo
}
type Option_Option struct {
    pathList *LnsList
    mode string
    inqMode string
    pattern string
    analyzeFileInfo *Option_AnalyzeFileInfo
    logLevel LnsAny
    transCtrlInfo *LnsTypes.Types_TransCtrlInfo
    FP Option_OptionMtd
}
func Option_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_Option).FP
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
func (self *Option_Option) Get_pathList(_env *LnsEnv) *LnsList{ return self.pathList }
func (self *Option_Option) Get_mode(_env *LnsEnv) string{ return self.mode }
func (self *Option_Option) Get_inqMode(_env *LnsEnv) string{ return self.inqMode }
func (self *Option_Option) Get_pattern(_env *LnsEnv) string{ return self.pattern }
func (self *Option_Option) Get_analyzeFileInfo(_env *LnsEnv) *Option_AnalyzeFileInfo{ return self.analyzeFileInfo }
func (self *Option_Option) Get_logLevel(_env *LnsEnv) LnsAny{ return self.logLevel }
func (self *Option_Option) Get_transCtrlInfo(_env *LnsEnv) *LnsTypes.Types_TransCtrlInfo{ return self.transCtrlInfo }
// 43: DeclConstr
func (self *Option_Option) InitOption_Option(_env *LnsEnv) {
    self.logLevel = nil
    
    self.pathList = NewLnsList([]LnsAny{})
    
    self.mode = Option_Mode__Build
    
    self.inqMode = Option_InqMode__Def
    
    self.pattern = ""
    
    self.analyzeFileInfo = NewOption_AnalyzeFileInfo(_env)
    
    self.transCtrlInfo = LnsTypes.Types_TransCtrlInfo_create_normal(_env)
    
}


func Lns_Option_init(_env *LnsEnv) {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lns.@tags.@Option"
    Lns_InitMod()
    Lns_Log_init(_env)
    LnsTypes.Lns_Types_init(_env)
}
func init() {
    init_Option = false
}
