// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsAst "github.com/ifritJP/LuneScript/src/lune/base"
var init_Inq bool
var Inq__mod__ string

// 5: decl @lns.@tags.@Inq.InqDef
func Inq_InqDef(_env *LnsEnv, db *DBCtrl_DBCtrl,pattern string) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory(_env)
    db.FP.MapSymbolDecl(_env, pattern, DBCtrl_callbackSymbolDecl(func(_env *LnsEnv, item *DBCtrl_ItemSymbolDecl) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(_env, item.FP.Get_fileId(_env))
            if _path == nil{
                panic(_env.LuaVM.String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId(_env)}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(_env, Lns_io_stdout, pattern, path, factory.FP.Create(_env, path, nil), item.FP.Get_line(_env))
        return true
    }))
}


// 19: decl @lns.@tags.@Inq.InqRef
func Inq_InqRef(_env *LnsEnv, db *DBCtrl_DBCtrl,pattern string,onlySet bool) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory(_env)
    db.FP.MapSymbolRef(_env, pattern, onlySet, DBCtrl_callbackSymbolRef(func(_env *LnsEnv, item *DBCtrl_ItemSymbolRef) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(_env, item.FP.Get_fileId(_env))
            if _path == nil{
                panic(_env.LuaVM.String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId(_env)}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(_env, Lns_io_stdout, pattern, path, factory.FP.Create(_env, path, nil), item.FP.Get_line(_env))
        return true
    }))
}


// 33: decl @lns.@tags.@Inq.InqAllmut
func Inq_InqAllmut(_env *LnsEnv, db *DBCtrl_DBCtrl) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory(_env)
    db.FP.MapAllmutDecl(_env, DBCtrl_callbackAllmutDecl(func(_env *LnsEnv, item *DBCtrl_ItemAllmutDecl) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(_env, item.FP.Get_fileId(_env))
            if _path == nil{
                panic(_env.LuaVM.String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId(_env)}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(_env, Lns_io_stdout, "allmut", path, factory.FP.Create(_env, path, nil), item.FP.Get_line(_env))
        return true
    }))
}


// 47: decl @lns.@tags.@Inq.InqAsync
func Inq_InqAsync(_env *LnsEnv, db *DBCtrl_DBCtrl,asyncMode LnsInt) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory(_env)
    db.FP.MapAsyncMode(_env, asyncMode, DBCtrl_callbackAsyncMode(func(_env *LnsEnv, item *DBCtrl_ItemAsyncMode) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(_env, item.FP.Get_fileId(_env))
            if _path == nil{
                panic(_env.LuaVM.String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId(_env)}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(_env, Lns_io_stdout, _env.LuaVM.String_format("%d", []LnsAny{asyncMode}), path, factory.FP.Create(_env, path, nil), item.FP.Get_line(_env))
        return true
    }))
}


// 62: decl @lns.@tags.@Inq.InqLuaval
func Inq_InqLuaval(_env *LnsEnv, db *DBCtrl_DBCtrl) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory(_env)
    db.FP.MapLuavalRef(_env, DBCtrl_callbackLuavalRef(func(_env *LnsEnv, item *DBCtrl_ItemLuavalRef) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(_env, item.FP.Get_fileId(_env))
            if _path == nil{
                panic(_env.LuaVM.String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId(_env)}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(_env, Lns_io_stdout, "luaval", path, factory.FP.Create(_env, path, nil), item.FP.Get_line(_env))
        return true
    }))
}


// 75: decl @lns.@tags.@Inq.InqAsyncLock
func Inq_InqAsyncLock(_env *LnsEnv, db *DBCtrl_DBCtrl) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory(_env)
    db.FP.MapAsyncLock(_env, DBCtrl_callbackAsyncLock(func(_env *LnsEnv, item *DBCtrl_ItemAsyncLock) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(_env, item.FP.Get_fileId(_env))
            if _path == nil{
                panic(_env.LuaVM.String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId(_env)}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(_env, Lns_io_stdout, "luaval", path, factory.FP.Create(_env, path, nil), item.FP.Get_line(_env))
        return true
    }))
}

func Lns_Inq_init(_env *LnsEnv) {
    if init_Inq { return }
    init_Inq = true
    Inq__mod__ = "@lns.@tags.@Inq"
    Lns_InitMod()
    Lns_DBCtrl_init(_env)
    Lns_Util_init(_env)
    LnsAst.Lns_Ast_init(_env)
}
func init() {
    init_Inq = false
}
