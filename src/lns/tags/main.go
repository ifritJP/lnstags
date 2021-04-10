// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
import Option "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
// for 21
func main_convExp123(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func __main____anonymous_1017_(node *Nodes.Nodes_Node,parent *Nodes.Nodes_Node,relation string,depth LnsInt) LnsInt {
    Lns_print([]LnsAny{Nodes.Nodes_getNodeKindName(node.FP.Get_kind())})
    return Nodes.Nodes_NodeVisitMode__Child
}
func __main___anonymous_1014_(ast *TransUnit.TransUnit_ASTInfo) {
    ast.FP.Get_node().FP.Visit(Nodes.Nodes_NodeVisitor(__main____anonymous_1017_), 0)
}
func __main___anonymous_1024_(row *LnsList) bool {
    Lns_print([]LnsAny{Lns_forceCastInt(row.GetAt(1)) + 10, row.GetAt(2).(string) + "hoge"})
    return true
}
func __main___anonymous_1031_(row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1)})
    return false
}
func __main___anonymous_1038_(row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1), row.GetAt(2)})
    return true
}
// 7: decl @lns.@tags.@main.__main
func Main___main(args *LnsList) LnsInt {
    Lns_main_init()
    var option *Option.Option_Option
    option = Option.Option_createDefaultOption("test/main.lns")
    front.Front_build(option, front.Front_AstCallback(__main___anonymous_1014_))
    var db base.Base_DB
    
    {
        _db := main_convExp123(Lns_2DDD(base.Base_Open("hoge.sqlite3", false, false)))
        if _db == nil{
            Lns_print([]LnsAny{"open error"})
            return 1
        } else {
            db = _db.(base.Base_DB)
        }
    }
    var stmt string
    stmt = "      create table foo (id integer not null primary key, name text);\n   delete from foo;\n"
    db.Exec(stmt, nil)
    db.Begin()
    {
        var _from173 LnsInt = 0
        var _to173 LnsInt = 10
        for _work173 := _from173; _work173 <= _to173; _work173++ {
            index := _work173
            var sql string
            sql = Lns_getVM().String_format("insert into foo(id, name) values(%d, 'こんにちわ世界%03d')", []LnsAny{index, index})
            db.Exec(sql, nil)
        }
    }
    db.Commit()
    db.MapQuery("select id, name from foo", base.Base_queryForm(__main___anonymous_1024_))
    db.MapQuery("select name from foo where id = 3", base.Base_queryForm(__main___anonymous_1031_))
    db.Exec("delete from foo", nil)
    db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil)
    db.MapQuery("select id, name from foo", base.Base_queryForm(__main___anonymous_1038_))
    db.Close()
    return 0
}

func Lns_main_init() {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@tags.@main"
    Lns_InitMod()
    base.Lns_base_init()
    Option.Lns_Option_init()
    Nodes.Lns_Nodes_init()
    TransUnit.Lns_TransUnit_init()
    front.Lns_front_init()
}
func init() {
    init_main = false
}
