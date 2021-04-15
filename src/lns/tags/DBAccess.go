// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
var init_DBAccess bool
var DBAccess__mod__ string
// 21: decl @lns.@tags.@DBAccess.open
func DBAccess_open(path string,readonly bool) LnsAny {
    var db base.Base_DB
    var err string
    
    {
        _db, _err := base.Base_Open(path, readonly, false)
        if _db == nil{
            Lns_print([]LnsAny{1, Lns_getVM().String_format("DBAccess:open error -- %s", []LnsAny{err})})
            return nil
        } else {
            db = _db.(base.Base_DB)
            err = _err
        }
    }
    return NewDBAccess_DBAccess(db, path, readonly)
}

func DBAccess_begin___anonymous_1036_() string {
    return "start"
}
func DBAccess_begin___anonymous_1039_() string {
    return "db mode is read only"
}
func DBAccess_commit___anonymous_1045_() string {
    return "commit: start"
}
func DBAccess_commit___anonymous_1048_() string {
    return "commit: end"
}
// declaration Class -- DBAccess
type DBAccess_DBAccessMtd interface {
    Begin()
    Close()
    Commit()
    Exec(arg1 string, arg2 LnsAny)
    Get_beginFlag() bool
    Get_readonlyFlag() bool
    MapJoin(arg1 string, arg2 string, arg3 string, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny, arg7 base.Base_queryForm) bool
    MapJoin2(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 base.Base_queryForm) bool
    MapJoin3(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 string, arg8 LnsAny, arg9 LnsAny, arg10 LnsAny, arg11 base.Base_queryForm) bool
    MapRowList(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 base.Base_queryForm) bool
    outputLog(arg1 string)
}
type DBAccess_DBAccess struct {
    db base.Base_DB
    readonlyFlag bool
    path string
    beginFlag bool
    FP DBAccess_DBAccessMtd
}
func DBAccess_DBAccess2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBAccess_DBAccess).FP
}
type DBAccess_DBAccessDownCast interface {
    ToDBAccess_DBAccess() *DBAccess_DBAccess
}
func DBAccess_DBAccessDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBAccess_DBAccessDownCast)
    if ok { return work.ToDBAccess_DBAccess() }
    return nil
}
func (obj *DBAccess_DBAccess) ToDBAccess_DBAccess() *DBAccess_DBAccess {
    return obj
}
func NewDBAccess_DBAccess(arg1 base.Base_DB, arg2 string, arg3 bool) *DBAccess_DBAccess {
    obj := &DBAccess_DBAccess{}
    obj.FP = obj
    obj.InitDBAccess_DBAccess(arg1, arg2, arg3)
    return obj
}
func (self *DBAccess_DBAccess) Get_readonlyFlag() bool{ return self.readonlyFlag }
func (self *DBAccess_DBAccess) Get_beginFlag() bool{ return self.beginFlag }
// 13: DeclConstr
func (self *DBAccess_DBAccess) InitDBAccess_DBAccess(db base.Base_DB,path string,readonlyFlag bool) {
    self.db = db
    
    self.path = path
    
    self.beginFlag = false
    
    self.readonlyFlag = readonlyFlag
    
}

// 30: decl @lns.@tags.@DBAccess.DBAccess.close
func (self *DBAccess_DBAccess) Close() {
    self.db.Close()
}

// 34: decl @lns.@tags.@DBAccess.DBAccess.outputLog
func (self *DBAccess_DBAccess) outputLog(message string) {
}

// 43: decl @lns.@tags.@DBAccess.DBAccess.begin
func (self *DBAccess_DBAccess) Begin() {
    __func__ := "@lns.@tags.@DBAccess.DBAccess.begin"
    Log_log(Log_Level__Log, __func__, 45, Log_CreateMessage(DBAccess_begin___anonymous_1036_))
    
    if self.readonlyFlag{
        Log_log(Log_Level__Err, __func__, 48, Log_CreateMessage(DBAccess_begin___anonymous_1039_))
        
        Lns_getVM().OS_exit(1)
    }
    self.beginFlag = true
    
    self.db.Begin()
}

// 63: decl @lns.@tags.@DBAccess.DBAccess.commit
func (self *DBAccess_DBAccess) Commit() {
    __func__ := "@lns.@tags.@DBAccess.DBAccess.commit"
    if self.readonlyFlag{
        return 
    }
    if Lns_op_not(self.beginFlag){
        return 
    }
    self.beginFlag = false
    
    Log_log(Log_Level__Log, __func__, 73, Log_CreateMessage(DBAccess_commit___anonymous_1045_))
    
    self.db.Commit()
    Log_log(Log_Level__Log, __func__, 77, Log_CreateMessage(DBAccess_commit___anonymous_1048_))
    
}

// 80: decl @lns.@tags.@DBAccess.DBAccess.exec
func (self *DBAccess_DBAccess) Exec(stmt string,errHandle LnsAny) {
    self.db.Exec(stmt, errHandle)
}

// 84: decl @lns.@tags.@DBAccess.DBAccess.mapJoin
func (self *DBAccess_DBAccess) MapJoin(tableName string,otherTable string,on string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryForm) bool {
    var query string
    query = Lns_getVM().String_format("SELECT %s FROM %s INNER JOIN %s ON %s", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( attrib) ||
        Lns_GetEnv().SetStackVal( "*") ).(string), tableName, otherTable, on})
    if condition != nil{
        condition_113 := condition.(string)
        query = Lns_getVM().String_format("%s WHERE %s", []LnsAny{query, condition_113})
        
    }
    if limit != nil{
        limit_115 := limit.(LnsInt)
        query = Lns_getVM().String_format("%s LIMIT %d", []LnsAny{query, limit_115})
        
    }
    return self.db.MapQuery(query, _func)
}

// 99: decl @lns.@tags.@DBAccess.DBAccess.mapJoin2
func (self *DBAccess_DBAccess) MapJoin2(tableName string,otherTable string,on string,otherTable2 string,on2 string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryForm) bool {
    var query string
    query = Lns_getVM().String_format("SELECT %s FROM %s INNER JOIN %s ON %s INNER JOIN %s ON %s", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( attrib) ||
        Lns_GetEnv().SetStackVal( "*") ).(string), tableName, otherTable, on, otherTable2, on2})
    if condition != nil{
        condition_131 := condition.(string)
        query = Lns_getVM().String_format("%s WHERE %s", []LnsAny{query, condition_131})
        
    }
    if limit != nil{
        limit_133 := limit.(LnsInt)
        query = Lns_getVM().String_format("%s LIMIT %d", []LnsAny{query, limit_133})
        
    }
    return self.db.MapQuery(query, _func)
}

// 114: decl @lns.@tags.@DBAccess.DBAccess.mapJoin3
func (self *DBAccess_DBAccess) MapJoin3(tableName string,otherTable string,on string,otherTable2 string,on2 string,otherTable3 string,on3 string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryForm) bool {
    var query string
    query = Lns_getVM().String_format("SELECT %s FROM %s INNER JOIN %s ON %s INNER JOIN %s ON %s INNER JOIN %s ON %s", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( attrib) ||
        Lns_GetEnv().SetStackVal( "*") ).(string), tableName, otherTable, on, otherTable2, on2, otherTable3, on3})
    if condition != nil{
        condition_151 := condition.(string)
        query = Lns_getVM().String_format("%s WHERE %s", []LnsAny{query, condition_151})
        
    }
    if limit != nil{
        limit_153 := limit.(LnsInt)
        query = Lns_getVM().String_format("%s LIMIT %d", []LnsAny{query, limit_153})
        
    }
    return self.db.MapQuery(query, _func)
}

// 131: decl @lns.@tags.@DBAccess.DBAccess.mapRowList
func (self *DBAccess_DBAccess) MapRowList(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryForm) bool {
    var query string
    var ATTRIB string
    ATTRIB = Lns_unwrapDefault( attrib, "*").(string)
    if condition != nil{
        condition_166 := condition.(string)
        query = Lns_getVM().String_format("SELECT %s FROM %s WHERE %s", []LnsAny{ATTRIB, tableName, condition_166})
        
    } else {
        query = Lns_getVM().String_format("SELECT %s FROM %s", []LnsAny{ATTRIB, tableName})
        
    }
    if limit != nil{
        limit_169 := limit.(LnsInt)
        query = Lns_getVM().String_format("%s LIMIT %d", []LnsAny{query, limit_169})
        
    }
    return self.db.MapQuery(query, _func)
}


func Lns_DBAccess_init() {
    if init_DBAccess { return }
    init_DBAccess = true
    DBAccess__mod__ = "@lns.@tags.@DBAccess"
    Lns_InitMod()
    base.Lns_base_init()
    Lns_Log_init()
}
func init() {
    init_DBAccess = false
}
