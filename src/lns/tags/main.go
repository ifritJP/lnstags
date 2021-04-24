// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_main bool
var main__mod__ string
var main_dbPath string
// 13: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    var option *Option_Option
    option = Option_analyzeArgs(args)
    if _switch189 := option.FP.Get_mode(); _switch189 == Option_Mode__Init {
        DBCtrl_initDB(main_dbPath)
    } else if _switch189 == Option_Mode__Build {
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
        db.FP.Commit()
        Analyze_start(db, option)
        db.FP.DumpAll()
        db.FP.Close()
    } else if _switch189 == Option_Mode__InqDef || _switch189 == Option_Mode__InqRef {
        Log_setLevel(Log_Level__Err)
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
        if _switch173 := option.FP.Get_mode(); _switch173 == Option_Mode__InqDef {
            Inq_InqDef(db, option.FP.Get_pattern())
        } else if _switch173 == Option_Mode__InqRef {
            Inq_InqRef(db, option.FP.Get_pattern())
        }
        db.FP.Close()
    } else if _switch189 == Option_Mode__Test {
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
