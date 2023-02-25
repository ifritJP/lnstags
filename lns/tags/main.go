// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LuneTypes "github.com/ifritJP/LuneScript/src/lune/base"
import LuneAst "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
var main_dbPath string
// 18: decl @lns.@tags.@main.inq
func main_inq_0_(_env *LnsEnv, inqMode string,pattern string) LnsInt {
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(_env, main_dbPath, false)
        if _db == nil{
            Lns_print(Lns_2DDD("error"))
            return -1
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    if _switch0 := inqMode; _switch0 == Option_InqMode__Def {
        Inq_InqDef(_env, db, pattern)
    } else if _switch0 == Option_InqMode__Ref {
        Inq_InqRef(_env, db, pattern, false)
    } else if _switch0 == Option_InqMode__Set {
        Inq_InqRef(_env, db, pattern, true)
    } else if _switch0 == Option_InqMode__AllMut {
        Inq_InqAllmut(_env, db)
    } else if _switch0 == Option_InqMode__Async {
        Inq_InqAsync(_env, db, LuneAst.Ast_Async__Async)
    } else if _switch0 == Option_InqMode__Noasync {
        Inq_InqAsync(_env, db, LuneAst.Ast_Async__Noasync)
    } else if _switch0 == Option_InqMode__Luaval {
        Inq_InqLuaval(_env, db)
    } else if _switch0 == Option_InqMode__AsyncLock {
        Inq_InqAsyncLock(_env, db)
    }
    db.FP.Close(_env)
    return 0
}

// 53: decl @lns.@tags.@main.build
func main_build_1_(_env *LnsEnv, pathList *LnsList2_[string],transCtrlInfo *LuneTypes.Types_TransCtrlInfo) LnsInt {
    DBCtrl_initDB(_env, main_dbPath)
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(_env, main_dbPath, false)
        if _db == nil{
            Lns_print(Lns_2DDD("error"))
            return 1
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    db.FP.Commit(_env)
    Analyze_start(_env, db, pathList, transCtrlInfo)
    db.FP.Close(_env)
    return 0
}

// 67: decl @lns.@tags.@main.__main
func Main___main(_env *LnsEnv, args *LnsList) LnsInt {
    Lns_main_init( _env )
    var option *Option_Option
    option = Option_analyzeArgs(_env, args)
    if _switch0 := option.FP.Get_mode(_env); _switch0 == Option_Mode__Init {
        DBCtrl_initDB(_env, main_dbPath)
    } else if _switch0 == Option_Mode__Build {
        return main_build_1_(_env, option.FP.Get_pathList(_env), option.FP.Get_transCtrlInfo(_env))
    } else if _switch0 == Option_Mode__Update {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print(Lns_2DDD("error"))
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        var projId LnsAny
        projId = db.FP.GetProjId(_env, "./")
        var pathList *LnsList2_[string]
        pathList = NewLnsList2_[string]([]string{})
        db.FP.MapFilePath(_env, DBCtrl_MapFileCallBack(func(_env *LnsEnv, item *DBCtrl_ItemFilePath) bool {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( item.FP.Get_projId(_env) == projId) &&
                _env.SetStackVal( Lns_op_not(db.FP.GetMainFilePath(_env, item.FP.Get_id(_env)))) ).(bool)){
                pathList.Insert(item.FP.Get_path(_env))
            }
            return true
        }))
        db.FP.Close(_env)
        return main_build_1_(_env, pathList, option.FP.Get_transCtrlInfo(_env))
    } else if _switch0 == Option_Mode__Suffix {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print(Lns_2DDD("error"))
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        db.FP.MapNamespaceSuffix(_env, option.FP.Get_pattern(_env), DBCtrl_NameSpaceCallback(main___main___anonymous_1_))
        db.FP.Close(_env)
    } else if _switch0 == Option_Mode__Inq {
        main_inq_0_(_env, option.FP.Get_inqMode(_env), option.FP.Get_pattern(_env))
    } else if _switch0 == Option_Mode__InqAt {
        var analyzeFileInfo *Option_AnalyzeFileInfo
        analyzeFileInfo = option.FP.Get_analyzeFileInfo(_env)
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print(Lns_2DDD("error"))
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        var pattern string
        
        {
            _pattern := Pattern_getPatterAt(_env, db, analyzeFileInfo, option.FP.Get_inqMode(_env), option.FP.Get_transCtrlInfo(_env))
            if _pattern == nil{
                db.FP.Close(_env)
                Lns_print(Lns_2DDD(_env.GetVM().String_format("illegal pos -- %s:%d:%d", Lns_2DDD(analyzeFileInfo.FP.Get_path(_env), analyzeFileInfo.FP.Get_lineNo(_env), analyzeFileInfo.FP.Get_column(_env)))))
                return 1
            } else {
                pattern = _pattern.(string)
            }
        }
        db.FP.Close(_env)
        main_inq_0_(_env, option.FP.Get_inqMode(_env), pattern)
    } else if _switch0 == Option_Mode__Dump {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print(Lns_2DDD("error"))
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        db.FP.DumpAll(_env)
        db.FP.Close(_env)
    } else if _switch0 == Option_Mode__Test {
        DBCtrl_test(_env)
    }
    return 0
}


func main___main___anonymous_1_(_env *LnsEnv, item *DBCtrl_ItemNamespace) bool {
    Lns_print(Lns_2DDD(item.FP.Get_name(_env)))
    return true
}
func Lns_main_init(_env *LnsEnv) {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@tags.@main"
    Lns_InitMod()
    Lns_DBCtrl_init(_env)
    Lns_Analyze_init(_env)
    Lns_Option_init(_env)
    Lns_Util_init(_env)
    Lns_Inq_init(_env)
    Lns_Log_init(_env)
    Lns_Pattern_init(_env)
    LuneTypes.Lns_Types_init(_env)
    LuneAst.Lns_Ast_init(_env)
    main_dbPath = "lnstags.sqlite3"
}
func init() {
    init_main = false
}
