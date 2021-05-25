// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
var main_dbPath string
// 15: decl @lns.@tags.@main.inq
func main_inq_1008_(_env *LnsEnv, inqMode string,pattern string) LnsInt {
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(_env, main_dbPath, false)
        if _db == nil{
            Lns_print([]LnsAny{"error"})
            return -1
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    if _switch98 := inqMode; _switch98 == Option_InqMode__Def {
        Inq_InqDef(_env, db, pattern)
    } else if _switch98 == Option_InqMode__Ref {
        Inq_InqRef(_env, db, pattern, false)
    } else if _switch98 == Option_InqMode__Set {
        Inq_InqRef(_env, db, pattern, true)
    }
    db.FP.Close(_env)
    return 0
}

// 35: decl @lns.@tags.@main.build
func main_build_1024_(_env *LnsEnv, pathList *LnsList,transCtrlInfo *LnsTypes.Types_TransCtrlInfo) LnsInt {
    DBCtrl_initDB(_env, main_dbPath)
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(_env, main_dbPath, false)
        if _db == nil{
            Lns_print([]LnsAny{"error"})
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


func __main___anonymous_1076_(_env *LnsEnv, item *DBCtrl_ItemNamespace) bool {
    Lns_print([]LnsAny{item.FP.Get_name(_env)})
    return true
}
// 49: decl @lns.@tags.@main.__main
func Main___main(_env *LnsEnv, args *LnsList) LnsInt {
    Lns_main_init( _env )
    var option *Option_Option
    option = Option_analyzeArgs(_env, args)
    if _switch594 := option.FP.Get_mode(_env); _switch594 == Option_Mode__Init {
        DBCtrl_initDB(_env, main_dbPath)
    } else if _switch594 == Option_Mode__Build {
        return main_build_1024_(_env, option.FP.Get_pathList(_env), option.FP.Get_transCtrlInfo(_env))
    } else if _switch594 == Option_Mode__Update {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        var projId LnsAny
        projId = db.FP.GetProjId(_env, "./")
        var pathList *LnsList
        pathList = NewLnsList([]LnsAny{})
        db.FP.MapFilePath(_env, DBCtrl_MapFileCallBack(func(_env *LnsEnv, item *DBCtrl_ItemFilePath) bool {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( item.FP.Get_projId(_env) == projId) &&
                _env.SetStackVal( Lns_op_not(db.FP.GetMainFilePath(_env, item.FP.Get_id(_env)))) ).(bool)){
                pathList.Insert(item.FP.Get_path(_env))
            }
            return true
        }))
        db.FP.Close(_env)
        return main_build_1024_(_env, pathList, option.FP.Get_transCtrlInfo(_env))
    } else if _switch594 == Option_Mode__Suffix {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        db.FP.MapNamespaceSuffix(_env, option.FP.Get_pattern(_env), DBCtrl_NameSpaceCallback(__main___anonymous_1076_))
        db.FP.Close(_env)
    } else if _switch594 == Option_Mode__Inq {
        main_inq_1008_(_env, option.FP.Get_inqMode(_env), option.FP.Get_pattern(_env))
    } else if _switch594 == Option_Mode__InqAt {
        var analyzeFileInfo *Option_AnalyzeFileInfo
        analyzeFileInfo = option.FP.Get_analyzeFileInfo(_env)
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
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
                Lns_print([]LnsAny{_env.LuaVM.String_format("illegal pos -- %s:%d:%d", []LnsAny{analyzeFileInfo.FP.Get_path(_env), analyzeFileInfo.FP.Get_lineNo(_env), analyzeFileInfo.FP.Get_column(_env)})})
                return 1
            } else {
                pattern = _pattern.(string)
            }
        }
        db.FP.Close(_env)
        main_inq_1008_(_env, option.FP.Get_inqMode(_env), pattern)
    } else if _switch594 == Option_Mode__Dump {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(_env, main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        db.FP.DumpAll(_env)
        db.FP.Close(_env)
    } else if _switch594 == Option_Mode__Test {
        DBCtrl_test(_env)
    }
    return 0
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
    LnsTypes.Lns_Types_init(_env)
    main_dbPath = "lnstags.sqlite3"
}
func init() {
    init_main = false
}
