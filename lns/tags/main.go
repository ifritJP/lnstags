// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
var main_dbPath string
// 15: decl @lns.@tags.@main.inq
func main_inq_1004_(inqMode string,pattern string) LnsInt {
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(main_dbPath, false)
        if _db == nil{
            Lns_print([]LnsAny{"error"})
            return -1
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    if _switch98 := inqMode; _switch98 == Option_InqMode__Def {
        Inq_InqDef(db, pattern)
    } else if _switch98 == Option_InqMode__Ref {
        Inq_InqRef(db, pattern, false)
    } else if _switch98 == Option_InqMode__Set {
        Inq_InqRef(db, pattern, true)
    }
    db.FP.Close()
    return 0
}

// 35: decl @lns.@tags.@main.build
func main_build_1009_(pathList *LnsList,transCtrlInfo *LnsTypes.Types_TransCtrlInfo) LnsInt {
    DBCtrl_initDB(main_dbPath)
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(main_dbPath, false)
        if _db == nil{
            Lns_print([]LnsAny{"error"})
            return 1
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    db.FP.Commit()
    Analyze_start(db, pathList, transCtrlInfo)
    db.FP.Close()
    return 0
}


func __main___anonymous_1022_(item *DBCtrl_ItemNamespace) bool {
    Lns_print([]LnsAny{item.FP.Get_name()})
    return true
}
// 49: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    var option *Option_Option
    option = Option_analyzeArgs(args)
    if _switch594 := option.FP.Get_mode(); _switch594 == Option_Mode__Init {
        DBCtrl_initDB(main_dbPath)
    } else if _switch594 == Option_Mode__Build {
        return main_build_1009_(option.FP.Get_pathList(), option.FP.Get_transCtrlInfo())
    } else if _switch594 == Option_Mode__Update {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        var projId LnsAny
        projId = db.FP.GetProjId("./")
        var pathList *LnsList
        pathList = NewLnsList([]LnsAny{})
        db.FP.MapFilePath(DBCtrl_MapFileCallBack(func(item *DBCtrl_ItemFilePath) bool {
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( item.FP.Get_projId() == projId) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(db.FP.GetMainFilePath(item.FP.Get_id()))) ).(bool)){
                pathList.Insert(item.FP.Get_path())
            }
            return true
        }))
        db.FP.Close()
        return main_build_1009_(pathList, option.FP.Get_transCtrlInfo())
    } else if _switch594 == Option_Mode__Suffix {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        db.FP.MapNamespaceSuffix(option.FP.Get_pattern(), DBCtrl_NameSpaceCallback(__main___anonymous_1022_))
        db.FP.Close()
    } else if _switch594 == Option_Mode__Inq {
        main_inq_1004_(option.FP.Get_inqMode(), option.FP.Get_pattern())
    } else if _switch594 == Option_Mode__InqAt {
        var analyzeFileInfo *Option_AnalyzeFileInfo
        analyzeFileInfo = option.FP.Get_analyzeFileInfo()
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        var pattern string
        
        {
            _pattern := Pattern_getPatterAt(db, analyzeFileInfo, option.FP.Get_inqMode(), option.FP.Get_transCtrlInfo())
            if _pattern == nil{
                db.FP.Close()
                Lns_print([]LnsAny{Lns_getVM().String_format("illegal pos -- %s:%d:%d", []LnsAny{analyzeFileInfo.FP.Get_path(), analyzeFileInfo.FP.Get_lineNo(), analyzeFileInfo.FP.Get_column()})})
                return 1
            } else {
                pattern = _pattern.(string)
            }
        }
        db.FP.Close()
        main_inq_1004_(option.FP.Get_inqMode(), pattern)
    } else if _switch594 == Option_Mode__Dump {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open(main_dbPath, true)
            if _db == nil{
                Lns_print([]LnsAny{"error"})
                return 1
            } else {
                db = _db.(*DBCtrl_DBCtrl)
            }
        }
        db.FP.DumpAll()
        db.FP.Close()
    } else if _switch594 == Option_Mode__Test {
        DBCtrl_test()
    }
    return 0
}

func Lns_main_init() {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@tags.@main"
    Lns_InitMod()
    Lns_DBCtrl_init()
    Lns_Analyze_init()
    Lns_Option_init()
    Lns_Util_init()
    Lns_Inq_init()
    Lns_Log_init()
    Lns_Pattern_init()
    LnsTypes.Lns_Types_init()
    main_dbPath = "lnstags.sqlite3"
}
func init() {
    init_main = false
}
