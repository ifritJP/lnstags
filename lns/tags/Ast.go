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
// 15: decl @lns.@tags.@Ast.getFullNameSym
func Ast_getFullNameSym(_env *LnsEnv, filter *Nodes.Nodes_Filter,symbolInfo *LnsAst.Ast_SymbolInfo) string {
    if symbolInfo.FP.Get_namespaceTypeInfo(_env) == LnsAst.Ast_headTypeInfo{
        return symbolInfo.FP.Get_name(_env)
    }
    var name string
    name = _env.LuaVM.String_format("%s.%s", []LnsAny{filter.FP.GetFull(_env, symbolInfo.FP.Get_namespaceTypeInfo(_env), false), symbolInfo.FP.Get_name(_env)})
    return name
}

// 24: decl @lns.@tags.@Ast.buildAst
func Ast_buildAst(_env *LnsEnv, logLevel LnsInt,pathList *LnsList,projDir LnsAny,useStdInMod LnsAny,forceAll bool,transCtrlInfo *Types.Types_TransCtrlInfo,astCallback front.Front_AstCallback) {
    if pathList.Len() == 0{
        return 
    }
    LnsLog.Log_setLevel(_env, logLevel)
    LnsUtil.Util_setDebugFlag(_env, false)
    if useStdInMod != nil{
        useStdInMod_34 := useStdInMod.(string)
        Parser.Parser_StreamParser_setStdinStream(_env, useStdInMod_34)
    }
    var lnsOpt *LnsOpt.Option_Option
    lnsOpt = LnsOpt.Option_createDefaultOption(_env, pathList, projDir)
    lnsOpt.TargetLuaVer = LuaVer.LuaVer_ver53
    
    lnsOpt.TransCtrlInfo.LegacyMutableControl = transCtrlInfo.LegacyMutableControl
    
    if forceAll{
        lnsOpt.TransCtrlInfo.UptodateMode = Types.Types_CheckingUptodateMode__ForceAll_Obj
        
    } else { 
        lnsOpt.TransCtrlInfo.UptodateMode = &Types.Types_CheckingUptodateMode__Force1{LnsUtil.Util_scriptPath2Module(_env, pathList.GetAt(1).(string))}
        
    }
    front.Front_build(_env, lnsOpt, astCallback)
}

func Lns_Ast_init(_env *LnsEnv) {
    if init_Ast { return }
    init_Ast = true
    Ast__mod__ = "@lns.@tags.@Ast"
    Lns_InitMod()
    LnsOpt.Lns_Option_init(_env)
    Types.Lns_Types_init(_env)
    Parser.Lns_Parser_init(_env)
    front.Lns_front_init(_env)
    Nodes.Lns_Nodes_init(_env)
    LnsAst.Lns_Ast_init(_env)
    LuaVer.Lns_LuaVer_init(_env)
    LnsLog.Lns_Log_init(_env)
    LnsUtil.Lns_Util_init(_env)
}
func init() {
    init_Ast = false
}
