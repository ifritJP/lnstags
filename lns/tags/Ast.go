// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LuneOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import LuneAst "github.com/ifritJP/LuneScript/src/lune/base"
import LuaVer "github.com/ifritJP/LuneScript/src/lune/base"
import LuneLog "github.com/ifritJP/LuneScript/src/lune/base"
import LuneUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Ast bool
var Ast__mod__ string
// 15: decl @lns.@tags.@Ast.getFullNameSym
func Ast_getFullNameSym(_env *LnsEnv, filter *Nodes.Nodes_Filter,symbolInfo *LuneAst.Ast_SymbolInfo) string {
    if symbolInfo.FP.Get_namespaceTypeInfo(_env) == LuneAst.Ast_headTypeInfo{
        return symbolInfo.FP.Get_name(_env)
    }
    var name string
    name = _env.GetVM().String_format("%s.%s", Lns_2DDD(filter.FP.GetFull(_env, symbolInfo.FP.Get_namespaceTypeInfo(_env), false), symbolInfo.FP.Get_name(_env)))
    return name
}

// 24: decl @lns.@tags.@Ast.buildAst
func Ast_buildAst(_env *LnsEnv, logLevel LnsInt,pathList *LnsList2_[string],projDir LnsAny,stdinFile LnsAny,forceAll bool,transCtrlInfo *Types.Types_TransCtrlInfo,astCallback front.Front_AstCallback) {
    if pathList.Len() == 0{
        return 
    }
    LuneLog.Log_setLevel(_env, logLevel)
    LuneUtil.Util_setDebugFlag(_env, false)
    var lnsOpt *LuneOpt.Option_Option
    lnsOpt = LuneOpt.Option_createDefaultOption(_env, pathList, projDir)
    lnsOpt.FP.Set_stdinFile(_env, stdinFile)
    lnsOpt.TargetLuaVer = LuaVer.LuaVer_ver53
    lnsOpt.TransCtrlInfo.LegacyMutableControl = transCtrlInfo.LegacyMutableControl
    if forceAll{
        lnsOpt.TransCtrlInfo.UptodateMode = Types.Types_CheckingUptodateMode__ForceAll_Obj
    } else { 
        lnsOpt.TransCtrlInfo.UptodateMode = &Types.Types_CheckingUptodateMode__Force1{LuneUtil.Util_scriptPath2Module(_env, pathList.GetAt(1))}
    }
    front.Front_build(_env, lnsOpt, astCallback)
}

func Lns_Ast_init(_env *LnsEnv) {
    if init_Ast { return }
    init_Ast = true
    Ast__mod__ = "@lns.@tags.@Ast"
    Lns_InitMod()
    LuneOpt.Lns_Option_init(_env)
    Types.Lns_Types_init(_env)
    front.Lns_front_init(_env)
    Nodes.Lns_Nodes_init(_env)
    LuneAst.Lns_Ast_init(_env)
    LuaVer.Lns_LuaVer_init(_env)
    LuneLog.Lns_Log_init(_env)
    LuneUtil.Lns_Util_init(_env)
}
func init() {
    init_Ast = false
}
