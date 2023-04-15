// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
import LuneAst "github.com/ifritJP/LuneScript/src/lune/base"
var init_DBCtrl bool
var DBCtrl__mod__ string
var DBCtrl_rootNsId LnsInt
var DBCtrl_userNsId LnsInt
var DBCtrl_systemFileId LnsInt
var DBCtrl_DB_VERSION LnsReal
type DBCtrl_MapFileCallBack func (_env *LnsEnv, arg1 *DBCtrl_ItemFilePath) bool
type DBCtrl_NameSpaceCallback func (_env *LnsEnv, arg1 *DBCtrl_ItemNamespace) bool
type DBCtrl_callbackAsyncMode func (_env *LnsEnv, arg1 *DBCtrl_ItemAsyncMode) bool
type DBCtrl_callbackAllmutDecl func (_env *LnsEnv, arg1 *DBCtrl_ItemAllmutDecl) bool
type DBCtrl_callbackLuavalRef func (_env *LnsEnv, arg1 *DBCtrl_ItemLuavalRef) bool
type DBCtrl_callbackAsyncLock func (_env *LnsEnv, arg1 *DBCtrl_ItemAsyncLock) bool
type DBCtrl_callbackSymbolDecl func (_env *LnsEnv, arg1 *DBCtrl_ItemSymbolDecl) bool
type DBCtrl_callbackSymbolRef func (_env *LnsEnv, arg1 *DBCtrl_ItemSymbolRef) bool
// for 346
func DBCtrl_convExp2_75(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 404
func DBCtrl_convExp2_291(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 428
func DBCtrl_convExp2_431(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 444
func DBCtrl_convExp2_511(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 461
func DBCtrl_convExp2_644(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 489
func DBCtrl_convExp3_100(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 503
func DBCtrl_convExp3_173(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 517
func DBCtrl_convExp3_240(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 525
func DBCtrl_convExp3_302(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 547
func DBCtrl_convExp3_371(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 580
func DBCtrl_convExp3_527(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 599
func DBCtrl_convExp3_605(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 680
func DBCtrl_convExp4_248(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 697
func DBCtrl_convExp4_328(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 722
func DBCtrl_convExp4_444(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 732
func DBCtrl_convExp4_508(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 744
func DBCtrl_convExp4_574(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 764
func DBCtrl_convExp4_667(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 778
func DBCtrl_convExp4_753(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 890
func DBCtrl_convExp0_1555(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 204
func DBCtrl_convExp1_667(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 376
func DBCtrl_convExp2_180(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 414
func DBCtrl_convExp2_336(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 896
func DBCtrl_convExp0_1585(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 908
func DBCtrl_convExp0_1670(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
func DBCtrl_open___anonymous_0_(_env *LnsEnv) string {
    return "open"
}
func DBCtrl_open___anonymous_1_(_env *LnsEnv) string {
    return "unknown version"
}


func DBCtrl_create___anonymous_0_(_env *LnsEnv) string {
    return "create"
}

func DBCtrl_DBCtrl_getMaxId___anonymous_1_(_env *LnsEnv, stmt string,msg string) {
}

// 189: decl @lns.@tags.@DBCtrl.open
func DBCtrl_open(_env *LnsEnv, path string,readonly bool) LnsAny {
    __func__ := "@lns.@tags.@DBCtrl.open"
    Log_log(_env, Log_Level__Log, __func__, 191, Log_CreateMessage(DBCtrl_open___anonymous_0_))
    
    var db *DBAccess_DBAccess
    
    {
        _db := DBAccess_open(_env, path, readonly)
        if _db == nil{
            return nil
        } else {
            db = _db.(*DBAccess_DBAccess)
        }
    }
    db.FP.Exec(_env, "PRAGMA case_sensitive_like=ON;", nil)
    var dbCtrl *DBCtrl_DBCtrl
    dbCtrl = NewDBCtrl_DBCtrl(_env, db, readonly)
    if Lns_op_not(readonly){
        dbCtrl.FP.Begin(_env)
    }
    var item *DBCtrl_ETC
    
    {
        _item := DBCtrl_convExp1_667(Lns_2DDD(DBCtrl_ETC__fromStem_5_(_env, dbCtrl.FP.GetRow(_env, "etc", "keyName = 'version'", nil, nil),nil)))
        if _item == nil{
            Log_log(_env, Log_Level__Err, __func__, 205, Log_CreateMessage(DBCtrl_open___anonymous_1_))
            
            db.FP.Close(_env)
            return nil
        } else {
            item = _item.(*DBCtrl_ETC)
        }
    }
    if Lns_tonumber(item.FP.Get_val(_env), nil) != DBCtrl_DB_VERSION{
        Log_log(_env, Log_Level__Err, __func__, 210, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("not support version. -- %s", Lns_2DDD(item.FP.Get_val(_env)))
        }))
        
        db.FP.Close(_env)
        return nil
    }
    if dbCtrl.FP.IsKilling(_env){
        panic("db is killed now")
    }
    return dbCtrl
}


// 374: decl @lns.@tags.@DBCtrl.getProjDir
func DBCtrl_getProjDir(_env *LnsEnv, path string,mod string) string {
    var workPath string
    workPath = Lns_car(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(mod,"@", "")).(string),"%.", "/")).(string) + ".lns"
    var projDir string
    projDir = DBCtrl_convExp2_180(Lns_2DDD(_env.GetVM().String_gsub(path,workPath + "$", "")))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( len(mod) != 0) &&
        _env.SetStackVal( len(projDir) == 0) ).(bool)){
        projDir = "./"
    }
    return projDir
}

// 383: decl @lns.@tags.@DBCtrl.normalizePath
func DBCtrl_normalizePath_15_(_env *LnsEnv, path string) string {
    return Lns_car(_env.GetVM().String_gsub(path,"^%./", "")).(string)
}












// 647: decl @lns.@tags.@DBCtrl.create
func DBCtrl_create_20_(_env *LnsEnv, dbPath string) LnsAny {
    __func__ := "@lns.@tags.@DBCtrl.create"
    Log_log(_env, Log_Level__Log, __func__, 649, Log_CreateMessage(DBCtrl_create___anonymous_0_))
    
    var db *DBAccess_DBAccess
    
    {
        _db := DBAccess_open(_env, dbPath, false)
        if _db == nil{
            return nil
        } else {
            db = _db.(*DBAccess_DBAccess)
        }
    }
    var dbCtrl *DBCtrl_DBCtrl
    dbCtrl = NewDBCtrl_DBCtrl(_env, db, false)
    dbCtrl.FP.creataTables(_env)
    dbCtrl.FP.Begin(_env)
    return dbCtrl
}

// 663: decl @lns.@tags.@DBCtrl.initDB
func DBCtrl_initDB(_env *LnsEnv, dbPath string) {
    _env.GetVM().OS_remove(dbPath)
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_create_20_(_env, dbPath)
        if _db == nil{
            Lns_print(Lns_2DDD("create error"))
            return 
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    db.FP.Commit(_env)
    db.FP.Close(_env)
}








func DBCtrl_DBCtrl_dumpFile___anonymous_0_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("id"), items.Get("dir")))
    return true
}
func DBCtrl_DBCtrl_dumpFile___anonymous_1_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("id"), items.Get("projId"), items.Get("path"), items.Get("mod")))
    return true
}
func DBCtrl_DBCtrl_dumpFile___anonymous_2_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("mainId"), items.Get("subId")))
    return true
}
func DBCtrl_DBCtrl_dumpAsync___anonymous_0_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("nsId"), items.Get("mode")))
    return true
}
func DBCtrl_DBCtrl_dumpAll___anonymous_0_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("id"), items.Get("name")))
    return true
}
func DBCtrl_DBCtrl_dumpAll___anonymous_1_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("nsId"), items.Get("superNsId")))
    return true
}
func DBCtrl_DBCtrl_dumpAll___anonymous_2_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("nsId"), items.Get("fileId"), items.Get("line"), items.Get("column")))
    return true
}
func DBCtrl_DBCtrl_dumpAll___anonymous_3_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("nsId"), items.Get("fileId"), items.Get("line"), items.Get("column")))
    return true
}
func DBCtrl_DBCtrl_dumpAll___anonymous_4_(_env *LnsEnv, items *LnsMap) bool {
    Lns_print(Lns_2DDD(items.Get("nsId"), items.Get("fileId"), items.Get("line"), items.Get("column")))
    return true
}
// 877: decl @lns.@tags.@DBCtrl.test
func DBCtrl_test(_env *LnsEnv) bool {
    var dbPath string
    dbPath = "lnstags.sqlite3"
    {
        DBCtrl_initDB(_env, dbPath)
    }
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(_env, dbPath, false)
        if _db == nil{
            Lns_print(Lns_2DDD("open error"))
            return false
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    var fileId LnsInt
    fileId = DBCtrl_rootNsId
    for _, _path := range( NewLnsList2_[string](Lns_2DDDGen[string]("aa.lns", "bb.lns", "cc.lns")).Items ) {
        path := _path
        fileId = DBCtrl_convExp0_1555(Lns_2DDD(db.FP.AddFile(_env, path, Lns_car(_env.GetVM().String_gsub(path,"%.lns", "")).(string))))
    }
    var parentId LnsInt
    parentId = DBCtrl_rootNsId
    for _index, _name := range( NewLnsList2_[string](Lns_2DDDGen[string]("@hoge", "@hoge.@foo", "@hoge.@foo.bar")).Items ) {
        index := _index + 1
        name := _name
        var newid LnsInt
        newid = DBCtrl_convExp0_1585(Lns_2DDD(db.FP.AddNamespace(_env, name, parentId)))
        db.FP.AddSymbolDecl(_env, newid, fileId, 100 + index, index * 10)
        db.FP.AddSymbolRef(_env, newid, fileId, 200 + index, index * 20, true)
        db.FP.AddSymbolSet(_env, newid, fileId, 300 + index, index * 30)
        parentId = newid
    }
    {
        var added bool
        _,added = db.FP.AddNamespace(_env, "@hoge", DBCtrl_rootNsId)
        Lns_print(Lns_2DDD("added", added))
    }
    db.FP.Commit(_env)
    db.FP.DumpAll(_env)
    return true
}

// 25: decl @lns.@tags.@DBCtrl.IdMgr.getIdNext
func (self *DBCtrl_IdMgr) getIdNext(_env *LnsEnv) LnsInt {
    var idNum LnsInt
    idNum = self.idNum
    self.idNum = self.idNum + 1
    return idNum
}
// 42: decl @lns.@tags.@DBCtrl.DBCtrl.getMaxId
func DBCtrl_DBCtrl_getMaxId_1_(_env *LnsEnv, access *DBAccess_DBAccess,tableName string,defId LnsInt) LnsInt {
    var id LnsAny
    id = nil
    access.FP.MapRowList(_env, tableName, nil, 1, "MAX(id)", nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        id = items.Get("id")
        return false
    }), base.Base_errHandleForm(DBCtrl_DBCtrl_getMaxId___anonymous_1_))
    if id != nil{
        id_12 := id
        return Lns_forceCastInt(id_12)
    }
    return defId
}
// 73: decl @lns.@tags.@DBCtrl.DBCtrl.close
func (self *DBCtrl_DBCtrl) Close(_env *LnsEnv) {
    self.access.FP.Close(_env)
}
// 77: decl @lns.@tags.@DBCtrl.DBCtrl.begin
func (self *DBCtrl_DBCtrl) Begin(_env *LnsEnv) {
    self.access.FP.Begin(_env)
}
// 80: decl @lns.@tags.@DBCtrl.DBCtrl.commit
func (self *DBCtrl_DBCtrl) Commit(_env *LnsEnv) {
    self.access.FP.Commit(_env)
}
// 84: decl @lns.@tags.@DBCtrl.DBCtrl.exec
func (self *DBCtrl_DBCtrl) Exec(_env *LnsEnv, stmt string,errHandle LnsAny) {
    self.access.FP.Exec(_env, stmt, errHandle)
}
// 88: decl @lns.@tags.@DBCtrl.DBCtrl.insert
func (self *DBCtrl_DBCtrl) Insert(_env *LnsEnv, tableName string,values string) {
    self.access.FP.Insert(_env, tableName, values)
}
// 92: decl @lns.@tags.@DBCtrl.DBCtrl.update
func (self *DBCtrl_DBCtrl) Update(_env *LnsEnv, tableName string,set string,condition LnsAny) {
    var sql string
    sql = _env.GetVM().String_format("UPDATE %s SET %s", Lns_2DDD(tableName, set))
    if Lns_isCondTrue( condition){
        sql = _env.GetVM().String_format("%s WHERE %s", Lns_2DDD(sql, condition))
    }
    self.FP.Exec(_env, sql, nil)
}
// 102: decl @lns.@tags.@DBCtrl.DBCtrl.mapRowList
func (self *DBCtrl_DBCtrl) MapRowList(_env *LnsEnv, tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapRowList(_env, tableName, condition, limit, attrib, nil, _func, errHandle)
}
// 110: decl @lns.@tags.@DBCtrl.DBCtrl.mapRowListSort
func (self *DBCtrl_DBCtrl) MapRowListSort(_env *LnsEnv, tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,order LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapRowList(_env, tableName, condition, limit, attrib, order, _func, errHandle)
}
// 118: decl @lns.@tags.@DBCtrl.DBCtrl.getRowList
func (self *DBCtrl_DBCtrl) getRowList(_env *LnsEnv, tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,errHandle LnsAny) *LnsList2_[*LnsMap] {
    var rows *LnsList2_[*LnsMap]
    rows = NewLnsList2_[*LnsMap]([]*LnsMap{})
    self.FP.MapRowList(_env, tableName, condition, limit, attrib, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        rows.Insert(items)
        return true
    }), errHandle)
    return rows
}
// 131: decl @lns.@tags.@DBCtrl.DBCtrl.mapJoin
func (self *DBCtrl_DBCtrl) mapJoin(_env *LnsEnv, tableName string,otherTable string,on string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapJoin(_env, tableName, otherTable, on, condition, limit, attrib, _func, errHandle)
}
// 140: decl @lns.@tags.@DBCtrl.DBCtrl.getRow
func (self *DBCtrl_DBCtrl) GetRow(_env *LnsEnv, tableName string,condition LnsAny,attrib LnsAny,errHandle LnsAny) LnsAny {
    var row *LnsList2_[*LnsMap]
    row = self.FP.getRowList(_env, tableName, condition, 1, attrib, errHandle)
    if row.Len() == 0{
        return nil
    }
    return row.GetAt(1)
}
// 150: decl @lns.@tags.@DBCtrl.DBCtrl.getEtc
func (self *DBCtrl_DBCtrl) GetEtc(_env *LnsEnv, key string) LnsAny {
    {
        _etc := self.FP.GetRow(_env, "etc", "keyName = '" + key + "'", nil, nil)
        if !Lns_IsNil( _etc ) {
            etc := _etc.(*LnsMap)
            {
                _val := etc.Get("val")
                if !Lns_IsNil( _val ) {
                    val := _val
                    return val.(string)
                }
            }
        }
    }
    return nil
}
// 159: decl @lns.@tags.@DBCtrl.DBCtrl.setEtc
func (self *DBCtrl_DBCtrl) SetEtc(_env *LnsEnv, key string,val string) {
    var keyTxt string
    keyTxt = _env.GetVM().String_format("keyName = '%s'", Lns_2DDD(key))
    var valTxt string
    valTxt = _env.GetVM().String_format("val = '%s'", Lns_2DDD(val))
    if Lns_op_not(self.FP.GetEtc(_env, key)){
        self.FP.Insert(_env, "etc", _env.GetVM().String_format("'%s', '%s'", Lns_2DDD(key, val)))
    } else { 
        self.FP.Update(_env, "etc", valTxt, keyTxt)
    }
}
// 169: decl @lns.@tags.@DBCtrl.DBCtrl.equalsEtc
func (self *DBCtrl_DBCtrl) EqualsEtc(_env *LnsEnv, key string,val string) bool {
    if self.FP.GetEtc(_env, key) == val{
        return true
    }
    return false
}
// 176: decl @lns.@tags.@DBCtrl.DBCtrl.isKilling
func (self *DBCtrl_DBCtrl) IsKilling(_env *LnsEnv) bool {
    if self.FP.EqualsEtc(_env, "killFlag", "1"){
        return true
    }
    return false
}
// 224: decl @lns.@tags.@DBCtrl.DBCtrl.creataTables
func (self *DBCtrl_DBCtrl) creataTables(_env *LnsEnv) {
    self.FP.Exec(_env, _env.GetVM().String_format("BEGIN;\nCREATE TABLE etc ( keyName VARCHAR UNIQUE COLLATE binary PRIMARY KEY, val VARCHAR);\nINSERT INTO etc VALUES( 'version', '%d' );\nINSERT INTO etc VALUES( 'projDir', '' );\nINSERT INTO etc VALUES( 'killFlag', '0' );\nCREATE TABLE namespace ( id INTEGER PRIMARY KEY, snameId INTEGER, parentId INTEGER, digest CHAR(32), name VARCHAR UNIQUE COLLATE binary, otherName VARCHAR COLLATE binary, virtual INTEGER);\nINSERT INTO namespace VALUES( NULL, 1, 0, '', '', '', 0 );\n\nCREATE TABLE simpleName ( id INTEGER PRIMARY KEY, name VARCHAR UNIQUE COLLATE binary);\nCREATE TABLE projInfo ( id INTEGER PRIMARY KEY, dir VARCHAR UNIQUE COLLATE binary);\nINSERT INTO projInfo VALUES( 1, '' );\nCREATE TABLE filePath ( id INTEGER PRIMARY KEY, path VARCHAR UNIQUE COLLATE binary, mod VARCHAR COLLATE binary, projId INTEGER );\nINSERT INTO filePath VALUES( 1, '', '', 1 );\nCREATE TABLE subfile ( mainId INTEGER, subId INTEGER PRIMARY KEY );\n\nCREATE TABLE override (nsId INTEGER, superNsId INTEGER, PRIMARY KEY (nsId, superNsId));\nCREATE TABLE symbolDecl ( nsId INTEGER, snameId INTEGER, parentId INTEGER, type INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, comment VARCHAR COLLATE binary, hasBodyFlag INTEGER, PRIMARY KEY( nsId, fileId, line ) );\nINSERT INTO symbolDecl VALUES( 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, '', 0 );\n\nCREATE TABLE symbolRef ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, setOp INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );\nCREATE TABLE symbolSet ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );\n\nCREATE TABLE funcCall ( nsId INTEGER, snameId INTEGER, belongNsId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, PRIMARY KEY( nsId, belongNsId ) );\n\nCREATE TABLE allmutDecl ( nsId INTEGER PRIMARY KEY );\nCREATE TABLE asyncMode ( nsId INTEGER PRIMARY KEY, mode INTEGER );\nCREATE TABLE luavalRef ( fileId INTEGER, line INTEGER, column INTEGER, PRIMARY KEY( fileId, line, column ) );\nCREATE TABLE asyncLock ( fileId INTEGER, line INTEGER, column INTEGER, kind INTEGER, PRIMARY KEY( fileId, line, column, kind ) );\n\nCREATE TABLE incRef ( id INTEGER, baseFileId INTEGER, line INTEGER );\nCREATE TABLE incCache ( id INTEGER, baseFileId INTEGER, incFlag INTEGER, PRIMARY KEY( id, baseFileId ) );\nCREATE TABLE incBelong ( id INTEGER, baseFileId INTEGER, nsId INTEGER, PRIMARY KEY ( id, nsId ) );\nCREATE INDEX index_ns ON namespace ( id, snameId, parentId, name, otherName );\nCREATE INDEX index_sName ON simpleName ( id, name );\nCREATE INDEX index_filePath ON filePath ( id, path, mod );\nCREATE INDEX index_subfile ON subfile (subId );\nCREATE INDEX index_override ON override (nsId, superNsId);\nCREATE INDEX index_symDecl ON symbolDecl ( nsId, parentId, snameId, fileId );\nCREATE INDEX index_symRef ON symbolRef ( nsId, snameId, fileId, belongNsId );\nCREATE INDEX index_incRef ON incRef ( id, baseFileId );\nCREATE INDEX index_incCache ON incCache ( id, baseFileId, incFlag );\nCREATE INDEX index_incBelong ON incBelong ( id, baseFileId );\nCOMMIT;\n", Lns_2DDD((LnsInt)(DBCtrl_DB_VERSION))), nil)
}
// 340: decl @lns.@tags.@DBCtrl.DBCtrl.getProjId
func (self *DBCtrl_DBCtrl) GetProjId(_env *LnsEnv, path string) LnsAny {
    var projId LnsAny
    projId = nil
    self.FP.MapRowList(_env, "projInfo", _env.GetVM().String_format("dir = '%s'", Lns_2DDD(path)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _projInfo := DBCtrl_convExp2_75(Lns_2DDD(DBCtrl_ItemProjInfo__fromStem_5_(_env, items,nil)))
            if !Lns_IsNil( _projInfo ) {
                projInfo := _projInfo.(*DBCtrl_ItemProjInfo)
                projId = projInfo.FP.Get_id(_env)
            }
        }
        return false
    }), nil)
    return projId
}
// 363: decl @lns.@tags.@DBCtrl.DBCtrl.addProj
func (self *DBCtrl_DBCtrl) AddProj(_env *LnsEnv, path string)(LnsInt, bool) {
    {
        _projId := self.FP.GetProjId(_env, path)
        if !Lns_IsNil( _projId ) {
            projId := _projId.(LnsInt)
            return projId, false
        }
    }
    var id LnsInt
    id = self.idMgrProjInfo.FP.getIdNext(_env)
    self.FP.Insert(_env, "projInfo", _env.GetVM().String_format("%d,'%s'", Lns_2DDD(id, path)))
    return id, true
}
// 396: decl @lns.@tags.@DBCtrl.DBCtrl.addFile
func (self *DBCtrl_DBCtrl) AddFile(_env *LnsEnv, path string,mod string)(LnsInt, bool) {
    path = DBCtrl_normalizePath_15_(_env, path)
    var fileId LnsAny
    fileId = nil
    self.FP.MapRowList(_env, "filePath", _env.GetVM().String_format("path = '%s'", Lns_2DDD(path)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp2_291(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(_env, items,nil)))
            if !Lns_IsNil( _filePath ) {
                filePath := _filePath.(*DBCtrl_ItemFilePath)
                fileId = filePath.FP.Get_id(_env)
            }
        }
        return false
    }), nil)
    if fileId != nil{
        fileId_25 := fileId.(LnsInt)
        return fileId_25, false
    }
    var projDir string
    projDir = DBCtrl_getProjDir(_env, path, mod)
    var projId LnsInt
    projId = DBCtrl_convExp2_336(Lns_2DDD(self.FP.AddProj(_env, projDir)))
    var id LnsInt
    id = self.idMgrFilePath.FP.getIdNext(_env)
    self.FP.Insert(_env, "filePath", _env.GetVM().String_format("%d,'%s','%s', %d", Lns_2DDD(id, path, Lns_car(_env.GetVM().String_gsub(mod,"@", "")).(string), projId)))
    return id, true
}
// 424: decl @lns.@tags.@DBCtrl.DBCtrl.mapFilePath
func (self *DBCtrl_DBCtrl) MapFilePath(_env *LnsEnv, callback DBCtrl_MapFileCallBack) {
    self.FP.MapRowList(_env, "filePath", nil, nil, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp2_431(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(_env, items,nil)))
            if !Lns_IsNil( _filePath ) {
                filePath := _filePath.(*DBCtrl_ItemFilePath)
                if Lns_op_not(callback(_env, filePath)){
                    return false
                }
            }
        }
        return true
    }), nil)
}
// 438: decl @lns.@tags.@DBCtrl.DBCtrl.getFileIdFromPath
func (self *DBCtrl_DBCtrl) GetFileIdFromPath(_env *LnsEnv, path string) LnsInt {
    __func__ := "@lns.@tags.@DBCtrl.DBCtrl.getFileIdFromPath"
    path = DBCtrl_normalizePath_15_(_env, path)
    var fileId LnsAny
    fileId = nil
    self.FP.MapRowList(_env, "filePath", _env.GetVM().String_format("path = '%s'", Lns_2DDD(path)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp2_511(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(_env, items,nil)))
            if !Lns_IsNil( _filePath ) {
                filePath := _filePath.(*DBCtrl_ItemFilePath)
                fileId = filePath.FP.Get_id(_env)
            }
        }
        return false
    }), nil)
    if fileId != nil{
        fileId_44 := fileId.(LnsInt)
        return fileId_44
    }
    Log_log(_env, Log_Level__Err, __func__, 452, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("not found file -- %s", Lns_2DDD(path))
    }))
    
    return DBCtrl_rootNsId
}
// 456: decl @lns.@tags.@DBCtrl.DBCtrl.getFilePath
func (self *DBCtrl_DBCtrl) GetFilePath(_env *LnsEnv, fileId LnsInt) LnsAny {
    var path LnsAny
    path = nil
    self.FP.MapRowList(_env, "filePath", _env.GetVM().String_format("id = %d", Lns_2DDD(fileId)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp2_644(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(_env, items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemFilePath)
                path = obj.FP.Get_path(_env)
            }
        }
        return false
    }), nil)
    return path
}
// 469: decl @lns.@tags.@DBCtrl.DBCtrl.addSubfile
func (self *DBCtrl_DBCtrl) AddSubfile(_env *LnsEnv, mainId LnsInt,subId LnsInt) {
    self.FP.Insert(_env, "subfile", _env.GetVM().String_format("%d, %d", Lns_2DDD(mainId, subId)))
}
// 474: decl @lns.@tags.@DBCtrl.DBCtrl.getMainFilePath
func (self *DBCtrl_DBCtrl) GetMainFilePath(_env *LnsEnv, subId LnsInt) LnsAny {
    {
        __map := self.FP.GetRow(_env, "subfile", _env.GetVM().String_format("subId = %d", Lns_2DDD(subId)), "mainId", nil)
        if !Lns_IsNil( __map ) {
            _map := __map.(*LnsMap)
            {
                _mainId := _map.Get("mainId")
                if !Lns_IsNil( _mainId ) {
                    mainId := _mainId
                    return self.FP.GetFilePath(_env, Lns_forceCastInt(mainId))
                }
            }
        }
    }
    return nil
}
// 484: decl @lns.@tags.@DBCtrl.DBCtrl.getName
func (self *DBCtrl_DBCtrl) GetName(_env *LnsEnv, nsId LnsInt) LnsAny {
    var name LnsAny
    name = nil
    self.FP.MapRowList(_env, "namespace", _env.GetVM().String_format("id = %d", Lns_2DDD(nsId)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp3_100(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(_env, items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemNamespace)
                name = obj.FP.Get_name(_env)
            }
        }
        return false
    }), nil)
    return name
}
// 498: decl @lns.@tags.@DBCtrl.DBCtrl.getNsId
func (self *DBCtrl_DBCtrl) getNsId(_env *LnsEnv, name string) LnsAny {
    var nsId LnsAny
    nsId = nil
    self.FP.MapRowList(_env, "namespace", _env.GetVM().String_format("name = '%s'", Lns_2DDD(name)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp3_173(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(_env, items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemNamespace)
                nsId = obj.FP.Get_id(_env)
            }
        }
        return false
    }), nil)
    return nsId
}
// 512: decl @lns.@tags.@DBCtrl.DBCtrl.mapNamespaceSuffix
func (self *DBCtrl_DBCtrl) MapNamespaceSuffix(_env *LnsEnv, suffix string,callback DBCtrl_NameSpaceCallback) {
    self.FP.MapRowList(_env, "namespace", _env.GetVM().String_format("name like '%%.%s' escape '$'", Lns_2DDD(suffix)), nil, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3_240(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemNamespace)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
    self.FP.MapRowList(_env, "namespace", _env.GetVM().String_format("name = '%s'", Lns_2DDD(suffix)), nil, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3_302(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemNamespace)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
}
// 541: decl @lns.@tags.@DBCtrl.DBCtrl.addNamespace
func (self *DBCtrl_DBCtrl) AddNamespace(_env *LnsEnv, fullName string,parentId LnsInt)(LnsInt, bool) {
    var id LnsAny
    id = nil
    self.FP.MapRowList(_env, "namespace", _env.GetVM().String_format("name = '%s'", Lns_2DDD(fullName)), 1, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp3_371(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(_env, items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemNamespace)
                id = obj.FP.Get_id(_env)
            }
        }
        return false
    }), nil)
    if id != nil{
        id_38 := id.(LnsInt)
        return id_38, false
    }
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var newId LnsInt
    newId = self.idMgrNamespace.FP.getIdNext(_env)
    self.FP.Insert(_env, "namespace", _env.GetVM().String_format("%d, %d, %d, '', '%s', '', 1", Lns_2DDD(newId, snid, parentId, fullName)))
    return newId, true
}
// 563: decl @lns.@tags.@DBCtrl.DBCtrl.addOverride
func (self *DBCtrl_DBCtrl) AddOverride(_env *LnsEnv, nsId LnsInt,superNsId LnsInt) {
    self.FP.Insert(_env, "override", _env.GetVM().String_format("%d, %d", Lns_2DDD(nsId, superNsId)))
}
// 568: decl @lns.@tags.@DBCtrl.DBCtrl.addAsyncMode
func (self *DBCtrl_DBCtrl) AddAsyncMode(_env *LnsEnv, nsId LnsInt,asyncMode LnsInt) {
    self.FP.Insert(_env, "asyncMode", _env.GetVM().String_format("%d, %d", Lns_2DDD(nsId, asyncMode)))
}
// 573: decl @lns.@tags.@DBCtrl.DBCtrl.mapAsyncMode
func (self *DBCtrl_DBCtrl) MapAsyncMode(_env *LnsEnv, asyncMode LnsInt,callback DBCtrl_callbackAsyncMode) {
    self.FP.mapJoin(_env, "asyncMode", "symbolDecl", _env.GetVM().String_format("asyncMode.nsId = symbolDecl.nsId AND asyncMode.mode = %d", Lns_2DDD(asyncMode)), nil, nil, "asyncMode.nsId, asyncMode.mode, symbolDecl.fileId, symbolDecl.line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3_527(Lns_2DDD(DBCtrl_ItemAsyncMode__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemAsyncMode)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
}
// 588: decl @lns.@tags.@DBCtrl.DBCtrl.addAllmutDecl
func (self *DBCtrl_DBCtrl) AddAllmutDecl(_env *LnsEnv, nsId LnsInt) {
    self.FP.Insert(_env, "allmutDecl", _env.GetVM().String_format("%d", Lns_2DDD(nsId)))
}
// 593: decl @lns.@tags.@DBCtrl.DBCtrl.mapAllmutDecl
func (self *DBCtrl_DBCtrl) MapAllmutDecl(_env *LnsEnv, callback DBCtrl_callbackAllmutDecl) {
    self.FP.mapJoin(_env, "allmutDecl", "symbolDecl", "allmutDecl.nsId = symbolDecl.nsId", nil, nil, "allmutDecl.nsId, symbolDecl.fileId, symbolDecl.line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3_605(Lns_2DDD(DBCtrl_ItemAllmutDecl__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemAllmutDecl)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
}
// 607: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolDecl
func (self *DBCtrl_DBCtrl) AddSymbolDecl(_env *LnsEnv, nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var kind LnsInt
    kind = 0
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var parentId LnsAny
    
    {
        _parentId := _env.NilAccFin( _env.NilAccPush( self.FP.GetRow(_env, "namespace", _env.GetVM().String_format("id = %d", Lns_2DDD(nsId)), "parentId", nil)) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsMap).Get("parentId")))
        if _parentId == nil{
            return 
        } else {
            parentId = _parentId
        }
    }
    self.FP.Insert(_env, "symbolDecl", _env.GetVM().String_format("%d, %d, %d, %d, %d, %d, %d, %d, %d, 0, '', 0", Lns_2DDD(nsId, snid, Lns_forceCastInt(parentId), kind, fileId, lineNo, column, lineNo, column)))
}
// 625: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolRef
func (self *DBCtrl_DBCtrl) AddSymbolRef(_env *LnsEnv, nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt,setOp bool) {
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var belongNsId LnsInt
    belongNsId = DBCtrl_rootNsId
    self.FP.Insert(_env, "symbolRef", _env.GetVM().String_format("%d, %d, %d, %d, %d, %d, %d", Lns_2DDD(nsId, snid, fileId, lineNo, column, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( setOp) &&
        _env.SetStackVal( 1) ||
        _env.SetStackVal( 0) ).(LnsInt), belongNsId)))
}
// 635: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolSet
func (self *DBCtrl_DBCtrl) AddSymbolSet(_env *LnsEnv, nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var belongNsId LnsInt
    belongNsId = DBCtrl_rootNsId
    self.FP.Insert(_env, "symbolSet", _env.GetVM().String_format("%d, %d, %d, %d, %d, %d", Lns_2DDD(nsId, snid, fileId, lineNo, column, belongNsId)))
}
// 675: decl @lns.@tags.@DBCtrl.DBCtrl.mapLuavalRef
func (self *DBCtrl_DBCtrl) MapLuavalRef(_env *LnsEnv, callback DBCtrl_callbackLuavalRef) {
    self.FP.MapRowListSort(_env, "luavalRef", nil, nil, nil, "fileId, line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_248(Lns_2DDD(DBCtrl_ItemLuavalRef__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemLuavalRef)
                return callback(_env, item)
            }
        }
        return false
    }), nil)
}
// 687: decl @lns.@tags.@DBCtrl.DBCtrl.addLuavalRef
func (self *DBCtrl_DBCtrl) AddLuavalRef(_env *LnsEnv, fileId LnsInt,lineNo LnsInt,column LnsInt) {
    self.FP.Insert(_env, "luavalRef", _env.GetVM().String_format("%d, %d, %d", Lns_2DDD(fileId, lineNo, column)))
}
// 692: decl @lns.@tags.@DBCtrl.DBCtrl.mapAsyncLock
func (self *DBCtrl_DBCtrl) MapAsyncLock(_env *LnsEnv, callback DBCtrl_callbackAsyncLock) {
    self.FP.MapRowListSort(_env, "asyncLock", nil, nil, nil, "fileId, line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_328(Lns_2DDD(DBCtrl_ItemAsyncLock__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemAsyncLock)
                return callback(_env, item)
            }
        }
        return false
    }), nil)
}
// 704: decl @lns.@tags.@DBCtrl.DBCtrl.addAsyncLock
func (self *DBCtrl_DBCtrl) AddAsyncLock(_env *LnsEnv, fileId LnsInt,lineNo LnsInt,column LnsInt,kind LnsInt) {
    self.FP.Insert(_env, "asyncLock", _env.GetVM().String_format("%d, %d, %d, %d", Lns_2DDD(fileId, lineNo, column, kind)))
}
// 711: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolDecl
func (self *DBCtrl_DBCtrl) MapSymbolDecl(_env *LnsEnv, name string,callback DBCtrl_callbackSymbolDecl) {
    var nsId LnsInt
    
    {
        _nsId := self.FP.getNsId(_env, name)
        if _nsId == nil{
            return 
        } else {
            nsId = _nsId.(LnsInt)
        }
    }
    var overrideStr string
    overrideStr = _env.GetVM().String_format("%d", Lns_2DDD(nsId))
    self.FP.MapRowList(_env, "override", _env.GetVM().String_format("superNsId = %d", Lns_2DDD(nsId)), nil, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_444(Lns_2DDD(DBCtrl_ItemOverride__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemOverride)
                overrideStr = _env.GetVM().String_format("%s, %d", Lns_2DDD(overrideStr, item.FP.Get_nsId(_env)))
            }
        }
        return true
    }), nil)
    self.FP.MapRowListSort(_env, "symbolDecl", _env.GetVM().String_format("nsId IN (%s)", Lns_2DDD(overrideStr)), nil, nil, "fileId, line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_508(Lns_2DDD(DBCtrl_ItemSymbolDecl__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemSymbolDecl)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
}
// 739: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolDeclForNsId
func (self *DBCtrl_DBCtrl) MapSymbolDeclForNsId(_env *LnsEnv, nsId LnsInt,callback DBCtrl_callbackSymbolDecl) {
    self.FP.MapRowListSort(_env, "symbolDecl", _env.GetVM().String_format("nsId IN (%d)", Lns_2DDD(nsId)), nil, nil, "fileId, line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_574(Lns_2DDD(DBCtrl_ItemSymbolDecl__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemSymbolDecl)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
}
// 753: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolRef
func (self *DBCtrl_DBCtrl) MapSymbolRef(_env *LnsEnv, name string,onlySet bool,callback DBCtrl_callbackSymbolRef) {
    var nsId LnsInt
    
    {
        _nsId := self.FP.getNsId(_env, name)
        if _nsId == nil{
            return 
        } else {
            nsId = _nsId.(LnsInt)
        }
    }
    var overrideStr string
    overrideStr = _env.GetVM().String_format("%d", Lns_2DDD(nsId))
    self.FP.MapRowList(_env, "override", _env.GetVM().String_format("nsId = %d", Lns_2DDD(nsId)), nil, nil, base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_667(Lns_2DDD(DBCtrl_ItemOverride__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemOverride)
                overrideStr = _env.GetVM().String_format("%s, %d", Lns_2DDD(overrideStr, item.FP.Get_superNsId(_env)))
            }
        }
        return true
    }), nil)
    var cond string
    cond = _env.GetVM().String_format("nsId IN (%s)", Lns_2DDD(overrideStr))
    if onlySet{
        cond = _env.GetVM().String_format("(%s) AND setOp = 1", Lns_2DDD(cond))
    }
    self.FP.MapRowListSort(_env, "symbolRef", cond, nil, nil, "fileId, line", base.Base_queryMapForm(func(_env *LnsEnv, items *LnsMap) bool {
        {
            _item := DBCtrl_convExp4_753(Lns_2DDD(DBCtrl_ItemSymbolRef__fromStem(_env, items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemSymbolRef)
                return callback(_env, item)
            }
        }
        return true
    }), nil)
}
// 785: decl @lns.@tags.@DBCtrl.DBCtrl.dumpFile
func (self *DBCtrl_DBCtrl) DumpFile(_env *LnsEnv) {
    Lns_print(Lns_2DDD("projId"))
    self.FP.MapRowList(_env, "projInfo", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpFile___anonymous_0_), nil)
    Lns_print(Lns_2DDD("filePath"))
    self.FP.MapRowList(_env, "filePath", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpFile___anonymous_1_), nil)
    Lns_print(Lns_2DDD("subfile"))
    self.FP.MapRowList(_env, "subfile", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpFile___anonymous_2_), nil)
}
// 814: decl @lns.@tags.@DBCtrl.DBCtrl.dumpAsync
func (self *DBCtrl_DBCtrl) DumpAsync(_env *LnsEnv) {
    Lns_print(Lns_2DDD("async"))
    self.FP.MapRowList(_env, "asyncMode", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpAsync___anonymous_0_), nil)
}
// 827: decl @lns.@tags.@DBCtrl.DBCtrl.dumpAll
func (self *DBCtrl_DBCtrl) DumpAll(_env *LnsEnv) {
    self.FP.DumpFile(_env)
    Lns_print(Lns_2DDD("namespace"))
    self.FP.MapRowList(_env, "namespace", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpAll___anonymous_0_), nil)
    Lns_print(Lns_2DDD("override"))
    self.FP.MapRowList(_env, "override", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpAll___anonymous_1_), nil)
    Lns_print(Lns_2DDD("symbolDecl"))
    self.FP.MapRowList(_env, "symbolDecl", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpAll___anonymous_2_), nil)
    Lns_print(Lns_2DDD("symbolRef"))
    self.FP.MapRowList(_env, "symbolRef", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpAll___anonymous_3_), nil)
    Lns_print(Lns_2DDD("symbolSet"))
    self.FP.MapRowList(_env, "symbolSet", nil, nil, nil, base.Base_queryMapForm(DBCtrl_DBCtrl_dumpAll___anonymous_4_), nil)
}
// declaration Class -- IdMgr
type DBCtrl_IdMgrMtd interface {
    getIdNext(_env *LnsEnv) LnsInt
}
type DBCtrl_IdMgr struct {
    idNum LnsInt
    FP DBCtrl_IdMgrMtd
}
func DBCtrl_IdMgr2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_IdMgr).FP
}
func DBCtrl_IdMgr_toSlice(slice []LnsAny) []*DBCtrl_IdMgr {
    ret := make([]*DBCtrl_IdMgr, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_IdMgrDownCast).ToDBCtrl_IdMgr()
    }
    return ret
}
type DBCtrl_IdMgrDownCast interface {
    ToDBCtrl_IdMgr() *DBCtrl_IdMgr
}
func DBCtrl_IdMgrDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_IdMgrDownCast)
    if ok { return work.ToDBCtrl_IdMgr() }
    return nil
}
func (obj *DBCtrl_IdMgr) ToDBCtrl_IdMgr() *DBCtrl_IdMgr {
    return obj
}
func NewDBCtrl_IdMgr(_env *LnsEnv, arg1 LnsInt) *DBCtrl_IdMgr {
    obj := &DBCtrl_IdMgr{}
    obj.FP = obj
    obj.InitDBCtrl_IdMgr(_env, arg1)
    return obj
}
// 21: DeclConstr
func (self *DBCtrl_IdMgr) InitDBCtrl_IdMgr(_env *LnsEnv, idNum LnsInt) {
    self.idNum = idNum
}


// declaration Class -- DBCtrl
type DBCtrl_DBCtrlMtd interface {
    AddAllmutDecl(_env *LnsEnv, arg1 LnsInt)
    AddAsyncLock(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    AddAsyncMode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt)
    AddFile(_env *LnsEnv, arg1 string, arg2 string)(LnsInt, bool)
    AddLuavalRef(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt)
    AddNamespace(_env *LnsEnv, arg1 string, arg2 LnsInt)(LnsInt, bool)
    AddOverride(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt)
    AddProj(_env *LnsEnv, arg1 string)(LnsInt, bool)
    AddSubfile(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt)
    AddSymbolDecl(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    AddSymbolRef(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 bool)
    AddSymbolSet(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    Begin(_env *LnsEnv)
    Close(_env *LnsEnv)
    Commit(_env *LnsEnv)
    creataTables(_env *LnsEnv)
    DumpAll(_env *LnsEnv)
    DumpAsync(_env *LnsEnv)
    DumpFile(_env *LnsEnv)
    EqualsEtc(_env *LnsEnv, arg1 string, arg2 string) bool
    Exec(_env *LnsEnv, arg1 string, arg2 LnsAny)
    GetEtc(_env *LnsEnv, arg1 string) LnsAny
    GetFileIdFromPath(_env *LnsEnv, arg1 string) LnsInt
    GetFilePath(_env *LnsEnv, arg1 LnsInt) LnsAny
    GetMainFilePath(_env *LnsEnv, arg1 LnsInt) LnsAny
    GetName(_env *LnsEnv, arg1 LnsInt) LnsAny
    getNsId(_env *LnsEnv, arg1 string) LnsAny
    GetProjId(_env *LnsEnv, arg1 string) LnsAny
    GetRow(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) LnsAny
    getRowList(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny) *LnsList2_[*LnsMap]
    Get_projDir(_env *LnsEnv) string
    Insert(_env *LnsEnv, arg1 string, arg2 string)
    IsKilling(_env *LnsEnv) bool
    MapAllmutDecl(_env *LnsEnv, arg1 DBCtrl_callbackAllmutDecl)
    MapAsyncLock(_env *LnsEnv, arg1 DBCtrl_callbackAsyncLock)
    MapAsyncMode(_env *LnsEnv, arg1 LnsInt, arg2 DBCtrl_callbackAsyncMode)
    MapFilePath(_env *LnsEnv, arg1 DBCtrl_MapFileCallBack)
    mapJoin(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny, arg7 base.Base_queryMapForm, arg8 LnsAny) bool
    MapLuavalRef(_env *LnsEnv, arg1 DBCtrl_callbackLuavalRef)
    MapNamespaceSuffix(_env *LnsEnv, arg1 string, arg2 DBCtrl_NameSpaceCallback)
    MapRowList(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 base.Base_queryMapForm, arg6 LnsAny) bool
    MapRowListSort(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 base.Base_queryMapForm, arg7 LnsAny) bool
    MapSymbolDecl(_env *LnsEnv, arg1 string, arg2 DBCtrl_callbackSymbolDecl)
    MapSymbolDeclForNsId(_env *LnsEnv, arg1 LnsInt, arg2 DBCtrl_callbackSymbolDecl)
    MapSymbolRef(_env *LnsEnv, arg1 string, arg2 bool, arg3 DBCtrl_callbackSymbolRef)
    SetEtc(_env *LnsEnv, arg1 string, arg2 string)
    Update(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny)
}
type DBCtrl_DBCtrl struct {
    access *DBAccess_DBAccess
    projDir string
    idMgrNamespace *DBCtrl_IdMgr
    idMgrSimpleName *DBCtrl_IdMgr
    idMgrFilePath *DBCtrl_IdMgr
    idMgrProjInfo *DBCtrl_IdMgr
    FP DBCtrl_DBCtrlMtd
}
func DBCtrl_DBCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_DBCtrl).FP
}
func DBCtrl_DBCtrl_toSlice(slice []LnsAny) []*DBCtrl_DBCtrl {
    ret := make([]*DBCtrl_DBCtrl, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_DBCtrlDownCast).ToDBCtrl_DBCtrl()
    }
    return ret
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
func NewDBCtrl_DBCtrl(_env *LnsEnv, arg1 *DBAccess_DBAccess, arg2 bool) *DBCtrl_DBCtrl {
    obj := &DBCtrl_DBCtrl{}
    obj.FP = obj
    obj.InitDBCtrl_DBCtrl(_env, arg1, arg2)
    return obj
}
func (self *DBCtrl_DBCtrl) Get_projDir(_env *LnsEnv) string{ return self.projDir }
// 59: DeclConstr
func (self *DBCtrl_DBCtrl) InitDBCtrl_DBCtrl(_env *LnsEnv, access *DBAccess_DBAccess,readonly bool) {
    self.access = access
    self.projDir = Depend_getCurDir(_env)
    self.idMgrNamespace = NewDBCtrl_IdMgr(_env, DBCtrl_DBCtrl_getMaxId_1_(_env, access, "namespace", DBCtrl_userNsId))
    self.idMgrSimpleName = NewDBCtrl_IdMgr(_env, DBCtrl_DBCtrl_getMaxId_1_(_env, access, "simpleName", DBCtrl_userNsId))
    self.idMgrFilePath = NewDBCtrl_IdMgr(_env, DBCtrl_DBCtrl_getMaxId_1_(_env, access, "filePath", DBCtrl_userNsId))
    self.idMgrProjInfo = NewDBCtrl_IdMgr(_env, DBCtrl_DBCtrl_getMaxId_1_(_env, access, "projInfo", DBCtrl_userNsId))
}


// declaration Class -- ETC
type DBCtrl_ETCMtd interface {
    ToMap() *LnsMap
    Get_keyName(_env *LnsEnv) string
    Get_val(_env *LnsEnv) string
}
type DBCtrl_ETC struct {
    keyName string
    val string
    FP DBCtrl_ETCMtd
}
func DBCtrl_ETC2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ETC).FP
}
func DBCtrl_ETC_toSlice(slice []LnsAny) []*DBCtrl_ETC {
    ret := make([]*DBCtrl_ETC, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ETCDownCast).ToDBCtrl_ETC()
    }
    return ret
}
type DBCtrl_ETCDownCast interface {
    ToDBCtrl_ETC() *DBCtrl_ETC
}
func DBCtrl_ETCDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ETCDownCast)
    if ok { return work.ToDBCtrl_ETC() }
    return nil
}
func (obj *DBCtrl_ETC) ToDBCtrl_ETC() *DBCtrl_ETC {
    return obj
}
func NewDBCtrl_ETC(_env *LnsEnv, arg1 string, arg2 string) *DBCtrl_ETC {
    obj := &DBCtrl_ETC{}
    obj.FP = obj
    obj.InitDBCtrl_ETC(_env, arg1, arg2)
    return obj
}
func (self *DBCtrl_ETC) InitDBCtrl_ETC(_env *LnsEnv, arg1 string, arg2 string) {
    self.keyName = arg1
    self.val = arg2
}
func (self *DBCtrl_ETC) Get_keyName(_env *LnsEnv) string{ return self.keyName }
func (self *DBCtrl_ETC) Get_val(_env *LnsEnv) string{ return self.val }
func (self *DBCtrl_ETC) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["keyName"] = Lns_ToCollection( self.keyName )
    obj.Items["val"] = Lns_ToCollection( self.val )
    return obj
}
func (self *DBCtrl_ETC) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ETC__fromMap_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ETC_FromMap( arg1, paramList )
}
func DBCtrl_ETC__fromStem_5_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ETC_FromMap( arg1, paramList )
}
func DBCtrl_ETC_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ETC_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ETC_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ETC{}
    newObj.FP = newObj
    return DBCtrl_ETC_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ETC_FromMapMain( newObj *DBCtrl_ETC, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["keyName"], false, nil); !ok {
       return false,nil,"keyName:" + mess.(string)
    } else {
       newObj.keyName = conv.(string)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["val"], false, nil); !ok {
       return false,nil,"val:" + mess.(string)
    } else {
       newObj.val = conv.(string)
    }
    return true, newObj, nil
}

// declaration Class -- ItemProjInfo
type DBCtrl_ItemProjInfoMtd interface {
    ToMap() *LnsMap
    Get_dir(_env *LnsEnv) string
    Get_id(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemProjInfo struct {
    id LnsInt
    dir string
    FP DBCtrl_ItemProjInfoMtd
}
func DBCtrl_ItemProjInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemProjInfo).FP
}
func DBCtrl_ItemProjInfo_toSlice(slice []LnsAny) []*DBCtrl_ItemProjInfo {
    ret := make([]*DBCtrl_ItemProjInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemProjInfoDownCast).ToDBCtrl_ItemProjInfo()
    }
    return ret
}
type DBCtrl_ItemProjInfoDownCast interface {
    ToDBCtrl_ItemProjInfo() *DBCtrl_ItemProjInfo
}
func DBCtrl_ItemProjInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemProjInfoDownCast)
    if ok { return work.ToDBCtrl_ItemProjInfo() }
    return nil
}
func (obj *DBCtrl_ItemProjInfo) ToDBCtrl_ItemProjInfo() *DBCtrl_ItemProjInfo {
    return obj
}
func NewDBCtrl_ItemProjInfo(_env *LnsEnv, arg1 LnsInt, arg2 string) *DBCtrl_ItemProjInfo {
    obj := &DBCtrl_ItemProjInfo{}
    obj.FP = obj
    obj.InitDBCtrl_ItemProjInfo(_env, arg1, arg2)
    return obj
}
func (self *DBCtrl_ItemProjInfo) InitDBCtrl_ItemProjInfo(_env *LnsEnv, arg1 LnsInt, arg2 string) {
    self.id = arg1
    self.dir = arg2
}
func (self *DBCtrl_ItemProjInfo) Get_id(_env *LnsEnv) LnsInt{ return self.id }
func (self *DBCtrl_ItemProjInfo) Get_dir(_env *LnsEnv) string{ return self.dir }
func (self *DBCtrl_ItemProjInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["id"] = Lns_ToCollection( self.id )
    obj.Items["dir"] = Lns_ToCollection( self.dir )
    return obj
}
func (self *DBCtrl_ItemProjInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemProjInfo__fromMap_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemProjInfo_FromMap( arg1, paramList )
}
func DBCtrl_ItemProjInfo__fromStem_5_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemProjInfo_FromMap( arg1, paramList )
}
func DBCtrl_ItemProjInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemProjInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemProjInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemProjInfo{}
    newObj.FP = newObj
    return DBCtrl_ItemProjInfo_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemProjInfo_FromMapMain( newObj *DBCtrl_ItemProjInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["id"], false, nil); !ok {
       return false,nil,"id:" + mess.(string)
    } else {
       newObj.id = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["dir"], false, nil); !ok {
       return false,nil,"dir:" + mess.(string)
    } else {
       newObj.dir = conv.(string)
    }
    return true, newObj, nil
}

// declaration Class -- ItemFilePath
type DBCtrl_ItemFilePathMtd interface {
    ToMap() *LnsMap
    Get_id(_env *LnsEnv) LnsInt
    Get_mod(_env *LnsEnv) string
    Get_path(_env *LnsEnv) string
    Get_projId(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemFilePath struct {
    id LnsInt
    path string
    mod string
    projId LnsInt
    FP DBCtrl_ItemFilePathMtd
}
func DBCtrl_ItemFilePath2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemFilePath).FP
}
func DBCtrl_ItemFilePath_toSlice(slice []LnsAny) []*DBCtrl_ItemFilePath {
    ret := make([]*DBCtrl_ItemFilePath, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemFilePathDownCast).ToDBCtrl_ItemFilePath()
    }
    return ret
}
type DBCtrl_ItemFilePathDownCast interface {
    ToDBCtrl_ItemFilePath() *DBCtrl_ItemFilePath
}
func DBCtrl_ItemFilePathDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemFilePathDownCast)
    if ok { return work.ToDBCtrl_ItemFilePath() }
    return nil
}
func (obj *DBCtrl_ItemFilePath) ToDBCtrl_ItemFilePath() *DBCtrl_ItemFilePath {
    return obj
}
func NewDBCtrl_ItemFilePath(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 string, arg4 LnsInt) *DBCtrl_ItemFilePath {
    obj := &DBCtrl_ItemFilePath{}
    obj.FP = obj
    obj.InitDBCtrl_ItemFilePath(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemFilePath) InitDBCtrl_ItemFilePath(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 string, arg4 LnsInt) {
    self.id = arg1
    self.path = arg2
    self.mod = arg3
    self.projId = arg4
}
func (self *DBCtrl_ItemFilePath) Get_id(_env *LnsEnv) LnsInt{ return self.id }
func (self *DBCtrl_ItemFilePath) Get_path(_env *LnsEnv) string{ return self.path }
func (self *DBCtrl_ItemFilePath) Get_mod(_env *LnsEnv) string{ return self.mod }
func (self *DBCtrl_ItemFilePath) Get_projId(_env *LnsEnv) LnsInt{ return self.projId }
func (self *DBCtrl_ItemFilePath) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["id"] = Lns_ToCollection( self.id )
    obj.Items["path"] = Lns_ToCollection( self.path )
    obj.Items["mod"] = Lns_ToCollection( self.mod )
    obj.Items["projId"] = Lns_ToCollection( self.projId )
    return obj
}
func (self *DBCtrl_ItemFilePath) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemFilePath__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemFilePath_FromMap( arg1, paramList )
}
func DBCtrl_ItemFilePath__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemFilePath_FromMap( arg1, paramList )
}
func DBCtrl_ItemFilePath_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemFilePath_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemFilePath_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemFilePath{}
    newObj.FP = newObj
    return DBCtrl_ItemFilePath_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemFilePath_FromMapMain( newObj *DBCtrl_ItemFilePath, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["id"], false, nil); !ok {
       return false,nil,"id:" + mess.(string)
    } else {
       newObj.id = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["path"], false, nil); !ok {
       return false,nil,"path:" + mess.(string)
    } else {
       newObj.path = conv.(string)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["mod"], false, nil); !ok {
       return false,nil,"mod:" + mess.(string)
    } else {
       newObj.mod = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["projId"], false, nil); !ok {
       return false,nil,"projId:" + mess.(string)
    } else {
       newObj.projId = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemNamespace
type DBCtrl_ItemNamespaceMtd interface {
    ToMap() *LnsMap
    Get_id(_env *LnsEnv) LnsInt
    Get_name(_env *LnsEnv) string
    Get_parentId(_env *LnsEnv) LnsInt
    Get_snameId(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemNamespace struct {
    id LnsInt
    snameId LnsInt
    parentId LnsInt
    name string
    FP DBCtrl_ItemNamespaceMtd
}
func DBCtrl_ItemNamespace2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemNamespace).FP
}
func DBCtrl_ItemNamespace_toSlice(slice []LnsAny) []*DBCtrl_ItemNamespace {
    ret := make([]*DBCtrl_ItemNamespace, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemNamespaceDownCast).ToDBCtrl_ItemNamespace()
    }
    return ret
}
type DBCtrl_ItemNamespaceDownCast interface {
    ToDBCtrl_ItemNamespace() *DBCtrl_ItemNamespace
}
func DBCtrl_ItemNamespaceDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemNamespaceDownCast)
    if ok { return work.ToDBCtrl_ItemNamespace() }
    return nil
}
func (obj *DBCtrl_ItemNamespace) ToDBCtrl_ItemNamespace() *DBCtrl_ItemNamespace {
    return obj
}
func NewDBCtrl_ItemNamespace(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 string) *DBCtrl_ItemNamespace {
    obj := &DBCtrl_ItemNamespace{}
    obj.FP = obj
    obj.InitDBCtrl_ItemNamespace(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemNamespace) InitDBCtrl_ItemNamespace(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 string) {
    self.id = arg1
    self.snameId = arg2
    self.parentId = arg3
    self.name = arg4
}
func (self *DBCtrl_ItemNamespace) Get_id(_env *LnsEnv) LnsInt{ return self.id }
func (self *DBCtrl_ItemNamespace) Get_snameId(_env *LnsEnv) LnsInt{ return self.snameId }
func (self *DBCtrl_ItemNamespace) Get_parentId(_env *LnsEnv) LnsInt{ return self.parentId }
func (self *DBCtrl_ItemNamespace) Get_name(_env *LnsEnv) string{ return self.name }
func (self *DBCtrl_ItemNamespace) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["id"] = Lns_ToCollection( self.id )
    obj.Items["snameId"] = Lns_ToCollection( self.snameId )
    obj.Items["parentId"] = Lns_ToCollection( self.parentId )
    obj.Items["name"] = Lns_ToCollection( self.name )
    return obj
}
func (self *DBCtrl_ItemNamespace) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemNamespace__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemNamespace_FromMap( arg1, paramList )
}
func DBCtrl_ItemNamespace__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemNamespace_FromMap( arg1, paramList )
}
func DBCtrl_ItemNamespace_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemNamespace_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemNamespace_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemNamespace{}
    newObj.FP = newObj
    return DBCtrl_ItemNamespace_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemNamespace_FromMapMain( newObj *DBCtrl_ItemNamespace, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["id"], false, nil); !ok {
       return false,nil,"id:" + mess.(string)
    } else {
       newObj.id = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["snameId"], false, nil); !ok {
       return false,nil,"snameId:" + mess.(string)
    } else {
       newObj.snameId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.parentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.name = conv.(string)
    }
    return true, newObj, nil
}

// declaration Class -- ItemAllmutDecl
type DBCtrl_ItemAllmutDeclMtd interface {
    ToMap() *LnsMap
    Get_fileId(_env *LnsEnv) LnsInt
    Get_line(_env *LnsEnv) LnsInt
    Get_nsId(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemAllmutDecl struct {
    nsId LnsInt
    fileId LnsInt
    line LnsInt
    FP DBCtrl_ItemAllmutDeclMtd
}
func DBCtrl_ItemAllmutDecl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemAllmutDecl).FP
}
func DBCtrl_ItemAllmutDecl_toSlice(slice []LnsAny) []*DBCtrl_ItemAllmutDecl {
    ret := make([]*DBCtrl_ItemAllmutDecl, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemAllmutDeclDownCast).ToDBCtrl_ItemAllmutDecl()
    }
    return ret
}
type DBCtrl_ItemAllmutDeclDownCast interface {
    ToDBCtrl_ItemAllmutDecl() *DBCtrl_ItemAllmutDecl
}
func DBCtrl_ItemAllmutDeclDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemAllmutDeclDownCast)
    if ok { return work.ToDBCtrl_ItemAllmutDecl() }
    return nil
}
func (obj *DBCtrl_ItemAllmutDecl) ToDBCtrl_ItemAllmutDecl() *DBCtrl_ItemAllmutDecl {
    return obj
}
func NewDBCtrl_ItemAllmutDecl(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt) *DBCtrl_ItemAllmutDecl {
    obj := &DBCtrl_ItemAllmutDecl{}
    obj.FP = obj
    obj.InitDBCtrl_ItemAllmutDecl(_env, arg1, arg2, arg3)
    return obj
}
func (self *DBCtrl_ItemAllmutDecl) InitDBCtrl_ItemAllmutDecl(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt) {
    self.nsId = arg1
    self.fileId = arg2
    self.line = arg3
}
func (self *DBCtrl_ItemAllmutDecl) Get_nsId(_env *LnsEnv) LnsInt{ return self.nsId }
func (self *DBCtrl_ItemAllmutDecl) Get_fileId(_env *LnsEnv) LnsInt{ return self.fileId }
func (self *DBCtrl_ItemAllmutDecl) Get_line(_env *LnsEnv) LnsInt{ return self.line }
func (self *DBCtrl_ItemAllmutDecl) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    return obj
}
func (self *DBCtrl_ItemAllmutDecl) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemAllmutDecl__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemAllmutDecl_FromMap( arg1, paramList )
}
func DBCtrl_ItemAllmutDecl__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemAllmutDecl_FromMap( arg1, paramList )
}
func DBCtrl_ItemAllmutDecl_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemAllmutDecl_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemAllmutDecl_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemAllmutDecl{}
    newObj.FP = newObj
    return DBCtrl_ItemAllmutDecl_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemAllmutDecl_FromMapMain( newObj *DBCtrl_ItemAllmutDecl, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["nsId"], false, nil); !ok {
       return false,nil,"nsId:" + mess.(string)
    } else {
       newObj.nsId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["fileId"], false, nil); !ok {
       return false,nil,"fileId:" + mess.(string)
    } else {
       newObj.fileId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["line"], false, nil); !ok {
       return false,nil,"line:" + mess.(string)
    } else {
       newObj.line = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemAsyncMode
type DBCtrl_ItemAsyncModeMtd interface {
    ToMap() *LnsMap
    Get_fileId(_env *LnsEnv) LnsInt
    Get_line(_env *LnsEnv) LnsInt
    Get_mode(_env *LnsEnv) LnsInt
    Get_nsId(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemAsyncMode struct {
    nsId LnsInt
    mode LnsInt
    fileId LnsInt
    line LnsInt
    FP DBCtrl_ItemAsyncModeMtd
}
func DBCtrl_ItemAsyncMode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemAsyncMode).FP
}
func DBCtrl_ItemAsyncMode_toSlice(slice []LnsAny) []*DBCtrl_ItemAsyncMode {
    ret := make([]*DBCtrl_ItemAsyncMode, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemAsyncModeDownCast).ToDBCtrl_ItemAsyncMode()
    }
    return ret
}
type DBCtrl_ItemAsyncModeDownCast interface {
    ToDBCtrl_ItemAsyncMode() *DBCtrl_ItemAsyncMode
}
func DBCtrl_ItemAsyncModeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemAsyncModeDownCast)
    if ok { return work.ToDBCtrl_ItemAsyncMode() }
    return nil
}
func (obj *DBCtrl_ItemAsyncMode) ToDBCtrl_ItemAsyncMode() *DBCtrl_ItemAsyncMode {
    return obj
}
func NewDBCtrl_ItemAsyncMode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) *DBCtrl_ItemAsyncMode {
    obj := &DBCtrl_ItemAsyncMode{}
    obj.FP = obj
    obj.InitDBCtrl_ItemAsyncMode(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemAsyncMode) InitDBCtrl_ItemAsyncMode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) {
    self.nsId = arg1
    self.mode = arg2
    self.fileId = arg3
    self.line = arg4
}
func (self *DBCtrl_ItemAsyncMode) Get_nsId(_env *LnsEnv) LnsInt{ return self.nsId }
func (self *DBCtrl_ItemAsyncMode) Get_mode(_env *LnsEnv) LnsInt{ return self.mode }
func (self *DBCtrl_ItemAsyncMode) Get_fileId(_env *LnsEnv) LnsInt{ return self.fileId }
func (self *DBCtrl_ItemAsyncMode) Get_line(_env *LnsEnv) LnsInt{ return self.line }
func (self *DBCtrl_ItemAsyncMode) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["mode"] = Lns_ToCollection( self.mode )
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    return obj
}
func (self *DBCtrl_ItemAsyncMode) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemAsyncMode__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemAsyncMode_FromMap( arg1, paramList )
}
func DBCtrl_ItemAsyncMode__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemAsyncMode_FromMap( arg1, paramList )
}
func DBCtrl_ItemAsyncMode_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemAsyncMode_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemAsyncMode_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemAsyncMode{}
    newObj.FP = newObj
    return DBCtrl_ItemAsyncMode_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemAsyncMode_FromMapMain( newObj *DBCtrl_ItemAsyncMode, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["nsId"], false, nil); !ok {
       return false,nil,"nsId:" + mess.(string)
    } else {
       newObj.nsId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["mode"], false, nil); !ok {
       return false,nil,"mode:" + mess.(string)
    } else {
       newObj.mode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["fileId"], false, nil); !ok {
       return false,nil,"fileId:" + mess.(string)
    } else {
       newObj.fileId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["line"], false, nil); !ok {
       return false,nil,"line:" + mess.(string)
    } else {
       newObj.line = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemSymbolDecl
type DBCtrl_ItemSymbolDeclMtd interface {
    ToMap() *LnsMap
    Get_column(_env *LnsEnv) LnsInt
    Get_fileId(_env *LnsEnv) LnsInt
    Get_line(_env *LnsEnv) LnsInt
    Get_nsId(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemSymbolDecl struct {
    nsId LnsInt
    fileId LnsInt
    line LnsInt
    column LnsInt
    FP DBCtrl_ItemSymbolDeclMtd
}
func DBCtrl_ItemSymbolDecl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemSymbolDecl).FP
}
func DBCtrl_ItemSymbolDecl_toSlice(slice []LnsAny) []*DBCtrl_ItemSymbolDecl {
    ret := make([]*DBCtrl_ItemSymbolDecl, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemSymbolDeclDownCast).ToDBCtrl_ItemSymbolDecl()
    }
    return ret
}
type DBCtrl_ItemSymbolDeclDownCast interface {
    ToDBCtrl_ItemSymbolDecl() *DBCtrl_ItemSymbolDecl
}
func DBCtrl_ItemSymbolDeclDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemSymbolDeclDownCast)
    if ok { return work.ToDBCtrl_ItemSymbolDecl() }
    return nil
}
func (obj *DBCtrl_ItemSymbolDecl) ToDBCtrl_ItemSymbolDecl() *DBCtrl_ItemSymbolDecl {
    return obj
}
func NewDBCtrl_ItemSymbolDecl(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) *DBCtrl_ItemSymbolDecl {
    obj := &DBCtrl_ItemSymbolDecl{}
    obj.FP = obj
    obj.InitDBCtrl_ItemSymbolDecl(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemSymbolDecl) InitDBCtrl_ItemSymbolDecl(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) {
    self.nsId = arg1
    self.fileId = arg2
    self.line = arg3
    self.column = arg4
}
func (self *DBCtrl_ItemSymbolDecl) Get_nsId(_env *LnsEnv) LnsInt{ return self.nsId }
func (self *DBCtrl_ItemSymbolDecl) Get_fileId(_env *LnsEnv) LnsInt{ return self.fileId }
func (self *DBCtrl_ItemSymbolDecl) Get_line(_env *LnsEnv) LnsInt{ return self.line }
func (self *DBCtrl_ItemSymbolDecl) Get_column(_env *LnsEnv) LnsInt{ return self.column }
func (self *DBCtrl_ItemSymbolDecl) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    obj.Items["column"] = Lns_ToCollection( self.column )
    return obj
}
func (self *DBCtrl_ItemSymbolDecl) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemSymbolDecl__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemSymbolDecl_FromMap( arg1, paramList )
}
func DBCtrl_ItemSymbolDecl__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemSymbolDecl_FromMap( arg1, paramList )
}
func DBCtrl_ItemSymbolDecl_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemSymbolDecl_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemSymbolDecl_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemSymbolDecl{}
    newObj.FP = newObj
    return DBCtrl_ItemSymbolDecl_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemSymbolDecl_FromMapMain( newObj *DBCtrl_ItemSymbolDecl, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["nsId"], false, nil); !ok {
       return false,nil,"nsId:" + mess.(string)
    } else {
       newObj.nsId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["fileId"], false, nil); !ok {
       return false,nil,"fileId:" + mess.(string)
    } else {
       newObj.fileId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["line"], false, nil); !ok {
       return false,nil,"line:" + mess.(string)
    } else {
       newObj.line = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["column"], false, nil); !ok {
       return false,nil,"column:" + mess.(string)
    } else {
       newObj.column = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemSymbolRef
type DBCtrl_ItemSymbolRefMtd interface {
    ToMap() *LnsMap
    Get_column(_env *LnsEnv) LnsInt
    Get_fileId(_env *LnsEnv) LnsInt
    Get_line(_env *LnsEnv) LnsInt
    Get_nsId(_env *LnsEnv) LnsInt
    Get_setOp(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemSymbolRef struct {
    nsId LnsInt
    fileId LnsInt
    line LnsInt
    column LnsInt
    setOp LnsInt
    FP DBCtrl_ItemSymbolRefMtd
}
func DBCtrl_ItemSymbolRef2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemSymbolRef).FP
}
func DBCtrl_ItemSymbolRef_toSlice(slice []LnsAny) []*DBCtrl_ItemSymbolRef {
    ret := make([]*DBCtrl_ItemSymbolRef, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemSymbolRefDownCast).ToDBCtrl_ItemSymbolRef()
    }
    return ret
}
type DBCtrl_ItemSymbolRefDownCast interface {
    ToDBCtrl_ItemSymbolRef() *DBCtrl_ItemSymbolRef
}
func DBCtrl_ItemSymbolRefDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemSymbolRefDownCast)
    if ok { return work.ToDBCtrl_ItemSymbolRef() }
    return nil
}
func (obj *DBCtrl_ItemSymbolRef) ToDBCtrl_ItemSymbolRef() *DBCtrl_ItemSymbolRef {
    return obj
}
func NewDBCtrl_ItemSymbolRef(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 LnsInt) *DBCtrl_ItemSymbolRef {
    obj := &DBCtrl_ItemSymbolRef{}
    obj.FP = obj
    obj.InitDBCtrl_ItemSymbolRef(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *DBCtrl_ItemSymbolRef) InitDBCtrl_ItemSymbolRef(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 LnsInt) {
    self.nsId = arg1
    self.fileId = arg2
    self.line = arg3
    self.column = arg4
    self.setOp = arg5
}
func (self *DBCtrl_ItemSymbolRef) Get_nsId(_env *LnsEnv) LnsInt{ return self.nsId }
func (self *DBCtrl_ItemSymbolRef) Get_fileId(_env *LnsEnv) LnsInt{ return self.fileId }
func (self *DBCtrl_ItemSymbolRef) Get_line(_env *LnsEnv) LnsInt{ return self.line }
func (self *DBCtrl_ItemSymbolRef) Get_column(_env *LnsEnv) LnsInt{ return self.column }
func (self *DBCtrl_ItemSymbolRef) Get_setOp(_env *LnsEnv) LnsInt{ return self.setOp }
func (self *DBCtrl_ItemSymbolRef) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    obj.Items["column"] = Lns_ToCollection( self.column )
    obj.Items["setOp"] = Lns_ToCollection( self.setOp )
    return obj
}
func (self *DBCtrl_ItemSymbolRef) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemSymbolRef__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemSymbolRef_FromMap( arg1, paramList )
}
func DBCtrl_ItemSymbolRef__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemSymbolRef_FromMap( arg1, paramList )
}
func DBCtrl_ItemSymbolRef_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemSymbolRef_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemSymbolRef_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemSymbolRef{}
    newObj.FP = newObj
    return DBCtrl_ItemSymbolRef_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemSymbolRef_FromMapMain( newObj *DBCtrl_ItemSymbolRef, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["nsId"], false, nil); !ok {
       return false,nil,"nsId:" + mess.(string)
    } else {
       newObj.nsId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["fileId"], false, nil); !ok {
       return false,nil,"fileId:" + mess.(string)
    } else {
       newObj.fileId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["line"], false, nil); !ok {
       return false,nil,"line:" + mess.(string)
    } else {
       newObj.line = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["column"], false, nil); !ok {
       return false,nil,"column:" + mess.(string)
    } else {
       newObj.column = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["setOp"], false, nil); !ok {
       return false,nil,"setOp:" + mess.(string)
    } else {
       newObj.setOp = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemOverride
type DBCtrl_ItemOverrideMtd interface {
    ToMap() *LnsMap
    Get_nsId(_env *LnsEnv) LnsInt
    Get_superNsId(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemOverride struct {
    nsId LnsInt
    superNsId LnsInt
    FP DBCtrl_ItemOverrideMtd
}
func DBCtrl_ItemOverride2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemOverride).FP
}
func DBCtrl_ItemOverride_toSlice(slice []LnsAny) []*DBCtrl_ItemOverride {
    ret := make([]*DBCtrl_ItemOverride, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemOverrideDownCast).ToDBCtrl_ItemOverride()
    }
    return ret
}
type DBCtrl_ItemOverrideDownCast interface {
    ToDBCtrl_ItemOverride() *DBCtrl_ItemOverride
}
func DBCtrl_ItemOverrideDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemOverrideDownCast)
    if ok { return work.ToDBCtrl_ItemOverride() }
    return nil
}
func (obj *DBCtrl_ItemOverride) ToDBCtrl_ItemOverride() *DBCtrl_ItemOverride {
    return obj
}
func NewDBCtrl_ItemOverride(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *DBCtrl_ItemOverride {
    obj := &DBCtrl_ItemOverride{}
    obj.FP = obj
    obj.InitDBCtrl_ItemOverride(_env, arg1, arg2)
    return obj
}
func (self *DBCtrl_ItemOverride) InitDBCtrl_ItemOverride(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) {
    self.nsId = arg1
    self.superNsId = arg2
}
func (self *DBCtrl_ItemOverride) Get_nsId(_env *LnsEnv) LnsInt{ return self.nsId }
func (self *DBCtrl_ItemOverride) Get_superNsId(_env *LnsEnv) LnsInt{ return self.superNsId }
func (self *DBCtrl_ItemOverride) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["superNsId"] = Lns_ToCollection( self.superNsId )
    return obj
}
func (self *DBCtrl_ItemOverride) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemOverride__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemOverride_FromMap( arg1, paramList )
}
func DBCtrl_ItemOverride__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemOverride_FromMap( arg1, paramList )
}
func DBCtrl_ItemOverride_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemOverride_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemOverride_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemOverride{}
    newObj.FP = newObj
    return DBCtrl_ItemOverride_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemOverride_FromMapMain( newObj *DBCtrl_ItemOverride, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["nsId"], false, nil); !ok {
       return false,nil,"nsId:" + mess.(string)
    } else {
       newObj.nsId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["superNsId"], false, nil); !ok {
       return false,nil,"superNsId:" + mess.(string)
    } else {
       newObj.superNsId = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemLuavalRef
type DBCtrl_ItemLuavalRefMtd interface {
    ToMap() *LnsMap
    Get_fileId(_env *LnsEnv) LnsInt
    Get_line(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemLuavalRef struct {
    fileId LnsInt
    line LnsInt
    FP DBCtrl_ItemLuavalRefMtd
}
func DBCtrl_ItemLuavalRef2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemLuavalRef).FP
}
func DBCtrl_ItemLuavalRef_toSlice(slice []LnsAny) []*DBCtrl_ItemLuavalRef {
    ret := make([]*DBCtrl_ItemLuavalRef, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemLuavalRefDownCast).ToDBCtrl_ItemLuavalRef()
    }
    return ret
}
type DBCtrl_ItemLuavalRefDownCast interface {
    ToDBCtrl_ItemLuavalRef() *DBCtrl_ItemLuavalRef
}
func DBCtrl_ItemLuavalRefDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemLuavalRefDownCast)
    if ok { return work.ToDBCtrl_ItemLuavalRef() }
    return nil
}
func (obj *DBCtrl_ItemLuavalRef) ToDBCtrl_ItemLuavalRef() *DBCtrl_ItemLuavalRef {
    return obj
}
func NewDBCtrl_ItemLuavalRef(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *DBCtrl_ItemLuavalRef {
    obj := &DBCtrl_ItemLuavalRef{}
    obj.FP = obj
    obj.InitDBCtrl_ItemLuavalRef(_env, arg1, arg2)
    return obj
}
func (self *DBCtrl_ItemLuavalRef) InitDBCtrl_ItemLuavalRef(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) {
    self.fileId = arg1
    self.line = arg2
}
func (self *DBCtrl_ItemLuavalRef) Get_fileId(_env *LnsEnv) LnsInt{ return self.fileId }
func (self *DBCtrl_ItemLuavalRef) Get_line(_env *LnsEnv) LnsInt{ return self.line }
func (self *DBCtrl_ItemLuavalRef) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    return obj
}
func (self *DBCtrl_ItemLuavalRef) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemLuavalRef__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemLuavalRef_FromMap( arg1, paramList )
}
func DBCtrl_ItemLuavalRef__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemLuavalRef_FromMap( arg1, paramList )
}
func DBCtrl_ItemLuavalRef_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemLuavalRef_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemLuavalRef_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemLuavalRef{}
    newObj.FP = newObj
    return DBCtrl_ItemLuavalRef_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemLuavalRef_FromMapMain( newObj *DBCtrl_ItemLuavalRef, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["fileId"], false, nil); !ok {
       return false,nil,"fileId:" + mess.(string)
    } else {
       newObj.fileId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["line"], false, nil); !ok {
       return false,nil,"line:" + mess.(string)
    } else {
       newObj.line = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- ItemAsyncLock
type DBCtrl_ItemAsyncLockMtd interface {
    ToMap() *LnsMap
    Get_fileId(_env *LnsEnv) LnsInt
    Get_kind(_env *LnsEnv) LnsInt
    Get_line(_env *LnsEnv) LnsInt
}
type DBCtrl_ItemAsyncLock struct {
    fileId LnsInt
    line LnsInt
    kind LnsInt
    FP DBCtrl_ItemAsyncLockMtd
}
func DBCtrl_ItemAsyncLock2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DBCtrl_ItemAsyncLock).FP
}
func DBCtrl_ItemAsyncLock_toSlice(slice []LnsAny) []*DBCtrl_ItemAsyncLock {
    ret := make([]*DBCtrl_ItemAsyncLock, len(slice))
    for index, val := range slice {
        ret[index] = val.(DBCtrl_ItemAsyncLockDownCast).ToDBCtrl_ItemAsyncLock()
    }
    return ret
}
type DBCtrl_ItemAsyncLockDownCast interface {
    ToDBCtrl_ItemAsyncLock() *DBCtrl_ItemAsyncLock
}
func DBCtrl_ItemAsyncLockDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DBCtrl_ItemAsyncLockDownCast)
    if ok { return work.ToDBCtrl_ItemAsyncLock() }
    return nil
}
func (obj *DBCtrl_ItemAsyncLock) ToDBCtrl_ItemAsyncLock() *DBCtrl_ItemAsyncLock {
    return obj
}
func NewDBCtrl_ItemAsyncLock(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt) *DBCtrl_ItemAsyncLock {
    obj := &DBCtrl_ItemAsyncLock{}
    obj.FP = obj
    obj.InitDBCtrl_ItemAsyncLock(_env, arg1, arg2, arg3)
    return obj
}
func (self *DBCtrl_ItemAsyncLock) InitDBCtrl_ItemAsyncLock(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsInt) {
    self.fileId = arg1
    self.line = arg2
    self.kind = arg3
}
func (self *DBCtrl_ItemAsyncLock) Get_fileId(_env *LnsEnv) LnsInt{ return self.fileId }
func (self *DBCtrl_ItemAsyncLock) Get_line(_env *LnsEnv) LnsInt{ return self.line }
func (self *DBCtrl_ItemAsyncLock) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *DBCtrl_ItemAsyncLock) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    obj.Items["kind"] = Lns_ToCollection( self.kind )
    return obj
}
func (self *DBCtrl_ItemAsyncLock) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemAsyncLock__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemAsyncLock_FromMap( arg1, paramList )
}
func DBCtrl_ItemAsyncLock__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemAsyncLock_FromMap( arg1, paramList )
}
func DBCtrl_ItemAsyncLock_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := DBCtrl_ItemAsyncLock_FromMapSub(obj,false, paramList);
    return conv,mess
}
func DBCtrl_ItemAsyncLock_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &DBCtrl_ItemAsyncLock{}
    newObj.FP = newObj
    return DBCtrl_ItemAsyncLock_FromMapMain( newObj, objMap, paramList )
}
func DBCtrl_ItemAsyncLock_FromMapMain( newObj *DBCtrl_ItemAsyncLock, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["fileId"], false, nil); !ok {
       return false,nil,"fileId:" + mess.(string)
    } else {
       newObj.fileId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["line"], false, nil); !ok {
       return false,nil,"line:" + mess.(string)
    } else {
       newObj.line = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["kind"], false, nil); !ok {
       return false,nil,"kind:" + mess.(string)
    } else {
       newObj.kind = conv.(LnsInt)
    }
    return true, newObj, nil
}

func Lns_DBCtrl_init(_env *LnsEnv) {
    if init_DBCtrl { return }
    init_DBCtrl = true
    DBCtrl__mod__ = "@lns.@tags.@DBCtrl"
    Lns_InitMod()
    base.Lns_base_init(_env)
    LuneAst.Lns_Ast_init(_env)
    Lns_DBAccess_init(_env)
    Lns_Log_init(_env)
    Lns_Depend_init(_env)
    DBCtrl_rootNsId = 1
    DBCtrl_userNsId = 2
    DBCtrl_systemFileId = 1
    DBCtrl_DB_VERSION = 10.0
}
func init() {
    init_DBCtrl = false
}
