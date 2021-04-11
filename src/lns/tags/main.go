// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import Option "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
func __main____anonymous_1017_(node *Nodes.Nodes_Node,parent *Nodes.Nodes_Node,relation string,depth LnsInt) LnsInt {
    Lns_print([]LnsAny{Nodes.Nodes_getNodeKindName(node.FP.Get_kind())})
    return Nodes.Nodes_NodeVisitMode__Child
}
func __main___anonymous_1014_(ast *TransUnit.TransUnit_ASTInfo) {
    ast.FP.Get_node().FP.Visit(Nodes.Nodes_NodeVisitor(__main____anonymous_1017_), 0)
}
// 9: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    DBCtrl_test()
    var option *Option.Option_Option
    option = Option.Option_createDefaultOption("test/main.lns")
    front.Front_build(option, front.Front_AstCallback(__main___anonymous_1014_))
    return 0
}

func Lns_main_init() {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@tags.@main"
    Lns_InitMod()
    Option.Lns_Option_init()
    Nodes.Lns_Nodes_init()
    TransUnit.Lns_TransUnit_init()
    front.Lns_front_init()
    Lns_DBCtrl_init()
}
func init() {
    init_main = false
}
