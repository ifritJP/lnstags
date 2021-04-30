// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
import Parser "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import LnsAst "github.com/ifritJP/LuneScript/src/lune/base"
import LuaVer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsLog "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Ast bool
var Ast__mod__ string
// 13: decl @lns.@tags.@Ast.getFullNameSym
func Ast_getFullNameSym(filter *Nodes.Nodes_Filter,symbolInfo *LnsAst.Ast_SymbolInfo) string {
    if symbolInfo.FP.Get_namespaceTypeInfo() == LnsAst.Ast_headTypeInfo{
        return symbolInfo.FP.Get_name()
    }
    var name string
    name = Lns_getVM().String_format("%s.%s", []LnsAny{filter.FP.GetFull(symbolInfo.FP.Get_namespaceTypeInfo(), false), symbolInfo.FP.Get_name()})
    return name
}

// 22: decl @lns.@tags.@Ast.buildAst
func Ast_buildAst(logLevel LnsInt,path string,projDir LnsAny,useStdInMod LnsAny,forceAll bool,astCallback front.Front_AstCallback) {
    LnsLog.Log_setLevel(logLevel)
    LnsUtil.Util_setDebugFlag(false)
    if useStdInMod != nil{
        useStdInMod_5720 := useStdInMod.(string)
        Parser.Parser_StreamParser_setStdinStream(useStdInMod_5720)
    }
    var lnsOpt *LnsOpt.Option_Option
    lnsOpt = LnsOpt.Option_createDefaultOption(path, projDir)
    lnsOpt.TargetLuaVer = LuaVer.LuaVer_ver53
    
    if forceAll{
        lnsOpt.TransCtrlInfo.UptodateMode = Types.Types_CheckingUptodateMode__ForceAll_Obj
        
    } else { 
        lnsOpt.TransCtrlInfo.UptodateMode = &Types.Types_CheckingUptodateMode__Force1{LnsUtil.Util_scriptPath2Module(path)}
        
    }
    front.Front_build(lnsOpt, astCallback)
}

func Lns_Ast_init() {
    if init_Ast { return }
    init_Ast = true
    Ast__mod__ = "@lns.@tags.@Ast"
    Lns_InitMod()
    LnsOpt.Lns_Option_init()
    Types.Lns_Types_init()
    Parser.Lns_Parser_init()
    front.Lns_front_init()
    Nodes.Lns_Nodes_init()
    LnsAst.Lns_Ast_init()
    LuaVer.Lns_LuaVer_init()
    LnsLog.Lns_Log_init()
    LnsUtil.Lns_Util_init()
}
func init() {
    init_Ast = false
}
