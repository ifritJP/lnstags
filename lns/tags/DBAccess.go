// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
var init_DBAccess bool
var DBAccess__mod__ string
// for 32
func DBAccess_convExp0_334(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// 30: decl @lns.@tags.@DBAccess.open
func DBAccess_open(_env *LnsEnv, path string,readonly bool) LnsAny {
    var db base.Base_DB
    var err string
    
    {
        _db, _err := base.Base_Open(_env, path, readonly, false)
        if _db == nil{
            Lns_print(Lns_2DDD(1, _env.GetVM().String_format("DBAccess:open error -- %s", Lns_2DDD(err))))
            return nil
        } else {
            db = _db.(base.Base_DB)
            err = _err
        }
    }
    return NewDBAccess_DBAccess(_env, db, path, readonly)
}

func DBAccess_DBAccess_createTables___anonymous_0_(_env *LnsEnv, stmt string,msg string) {
    if Lns_op_not(Lns_car(_env.GetVM().String_find(msg,"already exists", 1, true))){
        Lns_print(Lns_2DDD(msg))
    }
}

func DBAccess_DBAccess_begin___anonymous_0_(_env *LnsEnv) string {
    return "start"
}
func DBAccess_DBAccess_begin___anonymous_1_(_env *LnsEnv) string {
    return "db mode is read only"
}
func DBAccess_DBAccess_commit___anonymous_0_(_env *LnsEnv) string {
    return "commit: start"
}
func DBAccess_DBAccess_commit___anonymous_1_(_env *LnsEnv) string {
    return "commit: end"
}
// 24: decl @lns.@tags.@DBAccess.DBAccess.errorExit
func (self *DBAccess_DBAccess) ErrorExit(_env *LnsEnv, mess string) {
    Lns_io_stderr.Write(_env, mess + "\n")
    _env.GetVM().OS_exit(1)
}
// 39: decl @lns.@tags.@DBAccess.DBAccess.close
func (self *DBAccess_DBAccess) Close(_env *LnsEnv) {
    self.db.Close(_env)
}
// 43: decl @lns.@tags.@DBAccess.DBAccess.outputLog
func (self *DBAccess_DBAccess) outputLog(_env *LnsEnv, message string) {
}
// 52: decl @lns.@tags.@DBAccess.DBAccess.begin
func (self *DBAccess_DBAccess) Begin(_env *LnsEnv) {
    __func__ := "@lns.@tags.@DBAccess.DBAccess.begin"
    Log_log(_env, Log_Level__Log, __func__, 54, Log_CreateMessage(DBAccess_DBAccess_begin___anonymous_0_))
    
    if self.readonlyFlag{
        Log_log(_env, Log_Level__Err, __func__, 57, Log_CreateMessage(DBAccess_DBAccess_begin___anonymous_1_))
        
        _env.GetVM().OS_exit(1)
    }
    self.beginFlag = true
    self.db.Begin(_env)
}
// 72: decl @lns.@tags.@DBAccess.DBAccess.commit
func (self *DBAccess_DBAccess) Commit(_env *LnsEnv) {
    __func__ := "@lns.@tags.@DBAccess.DBAccess.commit"
    if self.readonlyFlag{
        return 
    }
    if Lns_op_not(self.beginFlag){
        return 
    }
    self.beginFlag = false
    Log_log(_env, Log_Level__Log, __func__, 82, Log_CreateMessage(DBAccess_DBAccess_commit___anonymous_0_))
    
    self.db.Commit(_env)
    Log_log(_env, Log_Level__Log, __func__, 86, Log_CreateMessage(DBAccess_DBAccess_commit___anonymous_1_))
    
}
// 89: decl @lns.@tags.@DBAccess.DBAccess.exec
func (self *DBAccess_DBAccess) Exec(_env *LnsEnv, stmt string,errHandle LnsAny) {
    self.db.Exec(_env, stmt, errHandle)
}
// 93: decl @lns.@tags.@DBAccess.DBAccess.mapJoin
func (self *DBAccess_DBAccess) MapJoin(_env *LnsEnv, tableName string,otherTable string,on string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    var query string
    query = _env.GetVM().String_format("SELECT DISTINCT %s FROM %s INNER JOIN %s ON %s", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( attrib) ||
        _env.SetStackVal( "*") ).(string), tableName, otherTable, on))
    if condition != nil{
        condition_137 := condition.(string)
        query = _env.GetVM().String_format("%s WHERE %s", Lns_2DDD(query, condition_137))
    }
    if limit != nil{
        limit_139 := limit.(LnsInt)
        query = _env.GetVM().String_format("%s LIMIT %d", Lns_2DDD(query, limit_139))
    }
    return self.db.MapQueryAsMap(_env, query, _func, errHandle)
}
// 109: decl @lns.@tags.@DBAccess.DBAccess.mapJoin2
func (self *DBAccess_DBAccess) MapJoin2(_env *LnsEnv, tableName string,otherTable string,on string,otherTable2 string,on2 string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    var query string
    query = _env.GetVM().String_format("SELECT DISTINCT %s FROM %s INNER JOIN %s ON %s INNER JOIN %s ON %s", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( attrib) ||
        _env.SetStackVal( "*") ).(string), tableName, otherTable, on, otherTable2, on2))
    if condition != nil{
        condition_143 := condition.(string)
        query = _env.GetVM().String_format("%s WHERE %s", Lns_2DDD(query, condition_143))
    }
    if limit != nil{
        limit_145 := limit.(LnsInt)
        query = _env.GetVM().String_format("%s LIMIT %d", Lns_2DDD(query, limit_145))
    }
    return self.db.MapQueryAsMap(_env, query, _func, errHandle)
}
// 125: decl @lns.@tags.@DBAccess.DBAccess.mapJoin3
func (self *DBAccess_DBAccess) MapJoin3(_env *LnsEnv, tableName string,otherTable string,on string,otherTable2 string,on2 string,otherTable3 string,on3 string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    var query string
    query = _env.GetVM().String_format("SELECT DISTINCT %s FROM %s INNER JOIN %s ON %s INNER JOIN %s ON %s INNER JOIN %s ON %s", Lns_2DDD(_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( attrib) ||
        _env.SetStackVal( "*") ).(string), tableName, otherTable, on, otherTable2, on2, otherTable3, on3))
    if condition != nil{
        condition_149 := condition.(string)
        query = _env.GetVM().String_format("%s WHERE %s", Lns_2DDD(query, condition_149))
    }
    if limit != nil{
        limit_151 := limit.(LnsInt)
        query = _env.GetVM().String_format("%s LIMIT %d", Lns_2DDD(query, limit_151))
    }
    return self.db.MapQueryAsMap(_env, query, _func, errHandle)
}
// 143: decl @lns.@tags.@DBAccess.DBAccess.mapRowList
func (self *DBAccess_DBAccess) MapRowList(_env *LnsEnv, tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,order LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    var query string
    var ATTRIB string
    ATTRIB = Lns_unwrapDefault( attrib, "*").(string)
    if condition != nil{
        condition_156 := condition.(string)
        query = _env.GetVM().String_format("SELECT %s FROM %s WHERE %s", Lns_2DDD(ATTRIB, tableName, condition_156))
    } else {
        query = _env.GetVM().String_format("SELECT %s FROM %s", Lns_2DDD(ATTRIB, tableName))
    }
    if order != nil{
        order_159 := order.(string)
        query = _env.GetVM().String_format("%s ORDER BY %s", Lns_2DDD(query, order_159))
    }
    if limit != nil{
        limit_161 := limit.(LnsInt)
        query = _env.GetVM().String_format("%s LIMIT %d", Lns_2DDD(query, limit_161))
    }
    return self.db.MapQueryAsMap(_env, query, _func, errHandle)
}
// 163: decl @lns.@tags.@DBAccess.DBAccess.createTables
func (self *DBAccess_DBAccess) CreateTables(_env *LnsEnv, sqlTxt string) {
    self.FP.Exec(_env, sqlTxt, base.Base_errHandleForm(DBAccess_DBAccess_createTables___anonymous_0_))
}
// 173: decl @lns.@tags.@DBAccess.DBAccess.insert
func (self *DBAccess_DBAccess) Insert(_env *LnsEnv, tableName string,values string) {
    self.FP.Exec(_env, _env.GetVM().String_format("INSERT INTO %s VALUES ( %s );", Lns_2DDD(tableName, values)), base.Base_errHandleForm(func(_env *LnsEnv, stmt string,message string) {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(Lns_car(_env.GetVM().String_find(message,"UNIQUE constraint failed", 1, true)))) &&
            _env.SetStackVal( Lns_op_not(Lns_car(_env.GetVM().String_find(message," not unique", 1, true)))) ).(bool)){
            self.FP.ErrorExit(_env, _env.GetVM().String_format("%s\n%s", Lns_2DDD(message, stmt)))
        }
    }))
}
// 187: decl @lns.@tags.@DBAccess.DBAccess.update
func (self *DBAccess_DBAccess) Update(_env *LnsEnv, tableName string,set string,condition LnsAny) {
    var sql string
    sql = _env.GetVM().String_format("UPDATE %s SET %s", Lns_2DDD(tableName, set))
    if Lns_isCondTrue( condition){
        sql = _env.GetVM().String_format("%s WHERE %s", Lns_2DDD(sql, condition))
    }
    self.FP.Exec(_env, sql, nil)
}
// declaration Class -- DBAccess
type DBAccess_DBAccessMtd interface {
    Begin(_env *LnsEnv)
    Close(_env *LnsEnv)
    Commit(_env *LnsEnv)
    CreateTables(_env *LnsEnv, arg1 string)
    ErrorExit(_env *LnsEnv, arg1 string)
    Exec(_env *LnsEnv, arg1 string, arg2 LnsAny)
    Get_beginFlag(_env *LnsEnv) bool
    Get_readonlyFlag(_env *LnsEnv) bool
    Insert(_env *LnsEnv, arg1 string, arg2 string)
    MapJoin(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny, arg7 base.Base_queryMapForm, arg8 LnsAny) bool
    MapJoin2(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 base.Base_queryMapForm, arg10 LnsAny) bool
    MapJoin3(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 string, arg8 LnsAny, arg9 LnsAny, arg10 LnsAny, arg11 base.Base_queryMapForm, arg12 LnsAny) bool
    MapRowList(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 base.Base_queryMapForm, arg7 LnsAny) bool
    outputLog(_env *LnsEnv, arg1 string)
    Update(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny)
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
func DBAccess_DBAccess_toSlice(slice []LnsAny) []*DBAccess_DBAccess {
    ret := make([]*DBAccess_DBAccess, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBAccess_DBAccessDownCast).ToDBAccess_DBAccess()
    }
    return ret
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
func NewDBAccess_DBAccess(_env *LnsEnv, arg1 base.Base_DB, arg2 string, arg3 bool) *DBAccess_DBAccess {
    obj := &DBAccess_DBAccess{}
    obj.FP = obj
    obj.InitDBAccess_DBAccess(_env, arg1, arg2, arg3)
    return obj
}
func (self *DBAccess_DBAccess) Get_readonlyFlag(_env *LnsEnv) bool{ return self.readonlyFlag }
func (self *DBAccess_DBAccess) Get_beginFlag(_env *LnsEnv) bool{ return self.beginFlag }
// 17: DeclConstr
func (self *DBAccess_DBAccess) InitDBAccess_DBAccess(_env *LnsEnv, db base.Base_DB,path string,readonlyFlag bool) {
    self.db = db
    self.path = path
    self.beginFlag = false
    self.readonlyFlag = readonlyFlag
}


func Lns_DBAccess_init(_env *LnsEnv) {
    if init_DBAccess { return }
    init_DBAccess = true
    DBAccess__mod__ = "@lns.@tags.@DBAccess"
    Lns_InitMod()
    base.Lns_base_init(_env)
    Lns_Log_init(_env)
}
func init() {
    init_DBAccess = false
}
