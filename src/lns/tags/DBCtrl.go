// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
var init_DBCtrl bool
var DBCtrl__mod__ string
// 22: decl @lns.@tags.@DBCtrl.open
func DBCtrl_open(path string,readonly bool) LnsAny {
    var db *DBAccess_DBAccess
    
    {
        _db := DBAccess_open(path, readonly)
        if _db == nil{
            return nil
        } else {
            db = _db.(*DBAccess_DBAccess)
        }
    }
    var dbCtrl *DBCtrl_DBCtrl
    dbCtrl = NewDBCtrl_DBCtrl(db)
    if Lns_op_not(readonly){
        dbCtrl.FP.Begin()
    }
    return dbCtrl
}

func DBCtrl_test___anonymous_1030_(row *LnsList) bool {
    Lns_print([]LnsAny{Lns_forceCastInt(row.GetAt(1)) + 10, row.GetAt(2).(string) + "hoge"})
    return true
}
func DBCtrl_test___anonymous_1037_(row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1)})
    return false
}
func DBCtrl_test___anonymous_1044_(row *LnsList) bool {
    Lns_print([]LnsAny{row.GetAt(1), row.GetAt(2)})
    return true
}
// 121: decl @lns.@tags.@DBCtrl.test
func DBCtrl_test() bool {
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open("hoge.sqlite3", false)
        if _db == nil{
            Lns_print([]LnsAny{"open error"})
            return false
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    db.FP.Test()
    return true
}

// declaration Class -- DBCtrl
type DBCtrl_DBCtrlMtd interface {
    Begin()
    Test()
}
type DBCtrl_DBCtrl struct {
    access *DBAccess_DBAccess
    FP DBCtrl_DBCtrlMtd
}
func DBCtrl_DBCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_DBCtrl).FP
}
type DBCtrl_DBCtrlDownCast interface {
    ToDBCtrl_DBCtrl() *DBCtrl_DBCtrl
}
func DBCtrl_DBCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_DBCtrlDownCast)
    if ok { return work.ToDBCtrl_DBCtrl() }
    return nil
}
func (obj *DBCtrl_DBCtrl) ToDBCtrl_DBCtrl() *DBCtrl_DBCtrl {
    return obj
}
func NewDBCtrl_DBCtrl(arg1 *DBAccess_DBAccess) *DBCtrl_DBCtrl {
    obj := &DBCtrl_DBCtrl{}
    obj.FP = obj
    obj.InitDBCtrl_DBCtrl(arg1)
    return obj
}
// 8: DeclConstr
func (self *DBCtrl_DBCtrl) InitDBCtrl_DBCtrl(access *DBAccess_DBAccess) {
    self.access = access
    
}

// 12: decl @lns.@tags.@DBCtrl.DBCtrl.begin
func (self *DBCtrl_DBCtrl) Begin() {
    self.access.FP.Begin()
}

// 78: decl @lns.@tags.@DBCtrl.DBCtrl.test
func (self *DBCtrl_DBCtrl) Test() {
    var stmt string
    stmt = "      create table foo (id integer not null primary key, name text);\n   delete from foo;\n"
    self.access.FP.Exec(stmt, nil)
    {
        var _from132 LnsInt = 0
        var _to132 LnsInt = 10
        for _work132 := _from132; _work132 <= _to132; _work132++ {
            index := _work132
            var sql string
            sql = Lns_getVM().String_format("insert into foo(id, name) values(%d, 'テスト%03d')", []LnsAny{index, index})
            self.access.FP.Exec(sql, nil)
        }
    }
    self.access.FP.Commit()
    self.access.FP.MapRowList("foo", nil, nil, "id, name", base.Base_queryForm(DBCtrl_test___anonymous_1030_))
    self.access.FP.MapRowList("foo", "id = 3", nil, "name", base.Base_queryForm(DBCtrl_test___anonymous_1037_))
    self.access.FP.Exec("delete from foo", nil)
    self.access.FP.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil)
    self.access.FP.MapRowList("foo", nil, nil, "id, name", base.Base_queryForm(DBCtrl_test___anonymous_1044_))
    self.access.FP.Close()
}


func Lns_DBCtrl_init() {
    if init_DBCtrl { return }
    init_DBCtrl = true
    DBCtrl__mod__ = "@lns.@tags.@DBCtrl"
    Lns_InitMod()
    base.Lns_base_init()
    Lns_DBAccess_init()
}
func init() {
    init_DBCtrl = false
}
