// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Option bool
var Option__mod__ string
// decl enum -- InqMode 
type Option_InqMode = string
const Option_InqMode__Def = "def"
const Option_InqMode__Ref = "ref"
var Option_InqModeList_ = NewLnsList( []LnsAny {
  Option_InqMode__Def,
  Option_InqMode__Ref,
})
func Option_InqMode_get__allList() *LnsList{
    return Option_InqModeList_
}
var Option_InqModeMap_ = map[string]string {
  Option_InqMode__Def: "InqMode.Def",
  Option_InqMode__Ref: "InqMode.Ref",
}
func Option_InqMode__from(arg1 string) LnsAny{
    if _, ok := Option_InqModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_InqMode_getTxt(arg1 string) string {
    return Option_InqModeMap_[arg1];
}
// decl enum -- Mode 
type Option_Mode = string
const Option_Mode__Build = "build"
const Option_Mode__Init = "init"
const Option_Mode__Inq = "inq"
const Option_Mode__InqAt = "inq-at"
const Option_Mode__Test = "test"
var Option_ModeList_ = NewLnsList( []LnsAny {
  Option_Mode__Init,
  Option_Mode__Build,
  Option_Mode__Inq,
  Option_Mode__InqAt,
  Option_Mode__Test,
})
func Option_Mode_get__allList() *LnsList{
    return Option_ModeList_
}
var Option_ModeMap_ = map[string]string {
  Option_Mode__Build: "Mode.Build",
  Option_Mode__Init: "Mode.Init",
  Option_Mode__Inq: "Mode.Inq",
  Option_Mode__InqAt: "Mode.InqAt",
  Option_Mode__Test: "Mode.Test",
}
func Option_Mode__from(arg1 string) LnsAny{
    if _, ok := Option_ModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_Mode_getTxt(arg1 string) string {
    return Option_ModeMap_[arg1];
}
type analyzeArgs__getNextOpNonNilFunc_1106_ func (arg1 string) string
// 47: decl @lns.@tags.@Option.printUsage
func Option_printUsage_1095_(messages LnsAny) {
    if messages != nil{
        messages_74 := messages.(string)
        Lns_io_stderr.Write(Lns_getVM().String_format("%s\n", []LnsAny{messages_74}))
    }
    Lns_print([]LnsAny{"usage: lnstags init [option]"})
    Lns_print([]LnsAny{"usage: lnstags build [option] filepath"})
    Lns_print([]LnsAny{"usage: lnstags inq <def|ref> pattern"})
    Lns_print([]LnsAny{"usage: lnstags inq-at <def|ref> filepath lineno column"})
    Lns_print([]LnsAny{"usage: lnstags test [option]"})
    Lns_getVM().OS_exit(1)
}

func analyzeArgs___anonymous_1109_(mess string) string {
    Option_printUsage_1095_("")
// insert a dummy
    return ""
}




// 59: decl @lns.@tags.@Option.analyzeArgs
func Option_analyzeArgs(argList *LnsList) *Option_Option {
    var option *Option_Option
    option = NewOption_Option()
    var index LnsInt
    index = 1
    var getNextOpNonNil analyzeArgs__getNextOpNonNilFunc_1106_
    getNextOpNonNil = analyzeArgs__getNextOpNonNilFunc_1106_(analyzeArgs___anonymous_1109_)
    
    var getNextOp func() LnsAny
    getNextOp = func() LnsAny {
        for  {
            if argList.Len() <= index{
                return nil
            }
            index = index + 1
            
            var arg string
            arg = argList.GetAt(index).(string)
            if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(arg,"^-", nil, nil))){
                if _switch377 := arg; _switch377 == "-i" {
                    option.analyzeFileInfo.stdinFlag = true
                    
                } else if _switch377 == "--log" {
                    option.logLevel = Log_str2level(getNextOpNonNil("logLevel"))
                    
                } else if _switch377 == "--simpleLog" {
                    Log_enableDetail(false)
                }
            } else { 
                return arg
            }
        }
    // insert a dummy
        return nil
    }
    getNextOpNonNil = analyzeArgs__getNextOpNonNilFunc_1106_(func(mess string) string {
        {
            _arg := getNextOp()
            if _arg != nil {
                arg := _arg.(string)
                return arg
            }
        }
        Option_printUsage_1095_(mess)
    // insert a dummy
        return ""
    })
    
    var getNextOpInt func(mess string) LnsInt
    getNextOpInt = func(mess string) LnsInt {
        {
            _arg := getNextOp()
            if _arg != nil {
                arg := _arg.(string)
                {
                    _num := Lns_tonumber(arg, nil)
                    if _num != nil {
                        num := _num.(LnsReal)
                        return (LnsInt)(num)
                    }
                }
                Option_printUsage_1095_(Lns_getVM().String_format("illegal num -- %s", []LnsAny{arg}))
            }
        }
        Option_printUsage_1095_(mess)
    // insert a dummy
        return 0
    }
    var getInqMode func() string
    getInqMode = func() string {
        var nextToken string
        
        {
            _nextToken := getNextOp()
            if _nextToken == nil{
                Option_printUsage_1095_("illegal inqMod")
            } else {
                nextToken = _nextToken.(string)
            }
        }
        var inqMode string
        
        {
            _inqMode := Option_InqMode__from(nextToken)
            if _inqMode == nil{
                Option_printUsage_1095_(Lns_getVM().String_format("illegal inqMod -- %s", []LnsAny{nextToken}))
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
            _arg := getNextOp()
            if _arg == nil{
                break
            } else {
                arg = _arg.(string)
            }
        }
        if Lns_op_not(mode){
            {
                _work := Option_Mode__from(arg)
                if _work != nil {
                    work := _work.(string)
                    mode = work
                    
                    if _switch644 := mode; _switch644 == Option_Mode__Inq {
                        option.inqMode = getInqMode()
                        
                        option.pattern = getNextOpNonNil("none pattern")
                        
                        Log_setLevel(Log_Level__Warn)
                    } else if _switch644 == Option_Mode__InqAt {
                        option.inqMode = getInqMode()
                        
                        option.analyzeFileInfo.path = getNextOpNonNil("none path")
                        
                        option.analyzeFileInfo.lineNo = getNextOpInt("none lineno")
                        
                        option.analyzeFileInfo.column = getNextOpInt("none column")
                        
                    }
                } else {
                    Option_printUsage_1095_(Lns_getVM().String_format("illegal option -- %s", []LnsAny{arg}))
                }
            }
        } else { 
            if _switch680 := mode; _switch680 == Option_Mode__Build {
                option.pathList.Insert(arg)
            }
        }
    }
    {
        _logLevel := option.logLevel
        if _logLevel != nil {
            logLevel := _logLevel.(LnsInt)
            Log_setLevel(logLevel)
        }
    }
    if mode != nil{
        mode_136 := mode.(string)
        option.mode = mode_136
        
        return option
    }
    Option_printUsage_1095_("none mode")
// insert a dummy
    return nil
}

// declaration Class -- AnalyzeFileInfo
type Option_AnalyzeFileInfoMtd interface {
    Get_column() LnsInt
    Get_lineNo() LnsInt
    Get_path() string
    Get_stdinFlag() bool
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
func NewOption_AnalyzeFileInfo() *Option_AnalyzeFileInfo {
    obj := &Option_AnalyzeFileInfo{}
    obj.FP = obj
    obj.InitOption_AnalyzeFileInfo()
    return obj
}
func (self *Option_AnalyzeFileInfo) Get_path() string{ return self.path }
func (self *Option_AnalyzeFileInfo) Get_lineNo() LnsInt{ return self.lineNo }
func (self *Option_AnalyzeFileInfo) Get_column() LnsInt{ return self.column }
func (self *Option_AnalyzeFileInfo) Get_stdinFlag() bool{ return self.stdinFlag }
// 21: DeclConstr
func (self *Option_AnalyzeFileInfo) InitOption_AnalyzeFileInfo() {
    self.path = ""
    
    self.lineNo = 0
    
    self.column = 0
    
    self.stdinFlag = false
    
}


// declaration Class -- Option
type Option_OptionMtd interface {
    Get_analyzeFileInfo() *Option_AnalyzeFileInfo
    Get_inqMode() string
    Get_logLevel() LnsAny
    Get_mode() string
    Get_pathList() *LnsList
    Get_pattern() string
}
type Option_Option struct {
    pathList *LnsList
    mode string
    inqMode string
    pattern string
    analyzeFileInfo *Option_AnalyzeFileInfo
    logLevel LnsAny
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
func NewOption_Option() *Option_Option {
    obj := &Option_Option{}
    obj.FP = obj
    obj.InitOption_Option()
    return obj
}
func (self *Option_Option) Get_pathList() *LnsList{ return self.pathList }
func (self *Option_Option) Get_mode() string{ return self.mode }
func (self *Option_Option) Get_inqMode() string{ return self.inqMode }
func (self *Option_Option) Get_pattern() string{ return self.pattern }
func (self *Option_Option) Get_analyzeFileInfo() *Option_AnalyzeFileInfo{ return self.analyzeFileInfo }
func (self *Option_Option) Get_logLevel() LnsAny{ return self.logLevel }
// 37: DeclConstr
func (self *Option_Option) InitOption_Option() {
    self.logLevel = nil
    
    self.pathList = NewLnsList([]LnsAny{})
    
    self.mode = Option_Mode__Build
    
    self.inqMode = Option_InqMode__Def
    
    self.pattern = ""
    
    self.analyzeFileInfo = NewOption_AnalyzeFileInfo()
    
}


func Lns_Option_init() {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lns.@tags.@Option"
    Lns_InitMod()
    Lns_Log_init()
}
func init() {
    init_Option = false
}
