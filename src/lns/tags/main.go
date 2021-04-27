// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_main bool
var main__mod__ string
var main_dbPath string
// 13: decl @lns.@tags.@main.inq
func main_inq_1007_(inqMode string,pattern string) LnsInt {
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
    if _switch78 := inqMode; _switch78 == Option_InqMode__Def {
        Inq_InqDef(db, pattern)
    } else if _switch78 == Option_InqMode__Ref {
        Inq_InqRef(db, pattern)
    }
    db.FP.Close()
    return 0
}

// 30: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    var option *Option_Option
    option = Option_analyzeArgs(args)
    if _switch296 := option.FP.Get_mode(); _switch296 == Option_Mode__Init {
        DBCtrl_initDB(main_dbPath)
    } else if _switch296 == Option_Mode__Build {
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
        Analyze_start(db, option)
        db.FP.Close()
    } else if _switch296 == Option_Mode__Inq {
        main_inq_1007_(option.FP.Get_inqMode(), option.FP.Get_pattern())
    } else if _switch296 == Option_Mode__InqAt {
        var analyzeFileInfo *Option_AnalyzeFileInfo
        analyzeFileInfo = option.FP.Get_analyzeFileInfo()
        var pattern string
        
        {
            _pattern := Analyze_getPatterAt(analyzeFileInfo)
            if _pattern == nil{
                Lns_print([]LnsAny{Lns_getVM().String_format("illegal pos -- %s:%d:%d", []LnsAny{analyzeFileInfo.FP.Get_path(), analyzeFileInfo.FP.Get_lineNo(), analyzeFileInfo.FP.Get_column()})})
                return 1
            } else {
                pattern = _pattern.(string)
            }
        }
        main_inq_1007_(option.FP.Get_inqMode(), pattern)
    } else if _switch296 == Option_Mode__Dump {
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
    } else if _switch296 == Option_Mode__Test {
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
    main_dbPath = "lnstags.sqlite3"
}
func init() {
    init_main = false
}
