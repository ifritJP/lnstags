// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_main bool
var main__mod__ string
// 7: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    var option *Option_Option
    option = Option_analyzeArgs(args)
    if _switch99 := option.FP.Get_mode(); _switch99 == Option_Mode__Init {
        DBCtrl_initDB()
    } else if _switch99 == Option_Mode__Build {
        var db *DBCtrl_DBCtrl
        
        {
            _db := DBCtrl_open("lnstags.sqlite3", false)
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
    } else if _switch99 == Option_Mode__Test {
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
}
func init() {
    init_main = false
}
