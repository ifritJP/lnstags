// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_main bool
var main__mod__ string
// 6: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    DBCtrl_test()
    Analyze_test()
    return 0
}

func Lns_main_init() {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@tags.@main"
    Lns_InitMod()
    Lns_DBCtrl_init()
    Lns_Analyze_init()
}
func init() {
    init_main = false
}
