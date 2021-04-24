// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Option bool
var Option__mod__ string
// decl enum -- Mode 
type Option_Mode = string
const Option_Mode__Build = "build"
const Option_Mode__Init = "init"
const Option_Mode__InqDef = "inq-def"
const Option_Mode__InqRef = "inq-ref"
const Option_Mode__Test = "test"
var Option_ModeList_ = NewLnsList( []LnsAny {
  Option_Mode__Init,
  Option_Mode__Build,
  Option_Mode__InqDef,
  Option_Mode__InqRef,
  Option_Mode__Test,
})
func Option_Mode_get__allList() *LnsList{
    return Option_ModeList_
}
var Option_ModeMap_ = map[string]string {
  Option_Mode__Build: "Mode.Build",
  Option_Mode__Init: "Mode.Init",
  Option_Mode__InqDef: "Mode.InqDef",
  Option_Mode__InqRef: "Mode.InqRef",
  Option_Mode__Test: "Mode.Test",
}
func Option_Mode__from(arg1 string) LnsAny{
    if _, ok := Option_ModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_Mode_getTxt(arg1 string) string {
    return Option_ModeMap_[arg1];
}
// 21: decl @lns.@tags.@Option.printUsage
func Option_printUsage_1048_(messages LnsAny) {
    if messages != nil{
        messages_30 := messages.(string)
        Lns_io_stderr.Write(Lns_getVM().String_format("%s\n", []LnsAny{messages_30}))
    }
    Lns_print([]LnsAny{"usage: lnstags init [option]"})
    Lns_print([]LnsAny{"usage: lnstags build [option] filepath"})
    Lns_print([]LnsAny{"usage: lnstags inq-def pattern"})
    Lns_print([]LnsAny{"usage: lnstags test [option]"})
    Lns_getVM().OS_exit(1)
}


// 32: decl @lns.@tags.@Option.analyzeArgs
func Option_analyzeArgs(argList *LnsList) *Option_Option {
    var index LnsInt
    index = 1
    var getNextOp func() LnsAny
    getNextOp = func() LnsAny {
        if argList.Len() <= index{
            return nil
        }
        index = index + 1
        
        return argList.GetAt(index).(string)
    }
    var mode LnsAny
    mode = nil
    var option *Option_Option
    option = NewOption_Option()
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
        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(arg,"^-", nil, nil))){
        } else { 
            if Lns_op_not(mode){
                {
                    _work := Option_Mode__from(arg)
                    if _work != nil {
                        work := _work.(string)
                        mode = work
                        
                    } else {
                        Option_printUsage_1048_(Lns_getVM().String_format("illegal option -- %s", []LnsAny{arg}))
                    }
                }
            } else { 
                if _switch280 := mode; _switch280 == Option_Mode__Build {
                    option.pathList.Insert(arg)
                } else if _switch280 == Option_Mode__InqDef || _switch280 == Option_Mode__InqRef {
                    option.pattern = arg
                    
                }
            }
        }
    }
    if mode != nil{
        mode_56 := mode.(string)
        option.mode = mode_56
        
        return option
    }
    Option_printUsage_1048_("none mode")
// insert a dummy
    return nil
}

// declaration Class -- Option
type Option_OptionMtd interface {
    Get_mode() string
    Get_pathList() *LnsList
    Get_pattern() string
}
type Option_Option struct {
    pathList *LnsList
    mode string
    pattern string
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
func (self *Option_Option) Get_pattern() string{ return self.pattern }
// 14: DeclConstr
func (self *Option_Option) InitOption_Option() {
    self.pathList = NewLnsList([]LnsAny{})
    
    self.mode = Option_Mode__Build
    
    self.pattern = ""
    
}


func Lns_Option_init() {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lns.@tags.@Option"
    Lns_InitMod()
}
func init() {
    init_Option = false
}
