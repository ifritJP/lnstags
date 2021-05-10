// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import base "github.com/ifritJP/lnssqlite3/src/lns/sqlite3"
var init_DBCtrl bool
var DBCtrl__mod__ string
var DBCtrl_rootNsId LnsInt
var DBCtrl_userNsId LnsInt
var DBCtrl_systemFileId LnsInt
var DBCtrl_DB_VERSION LnsReal
type DBCtrl_MapFileCallBack func (arg1 *DBCtrl_ItemFilePath) bool
type DBCtrl_NameSpaceCallback func (arg1 *DBCtrl_ItemNamespace) bool
type DBCtrl_callbackSymbolDecl func (arg1 *DBCtrl_ItemSymbolDecl) bool
type DBCtrl_callbackSymbolRef func (arg1 *DBCtrl_ItemSymbolRef) bool
// for 312
func DBCtrl_convExp1371(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 370
func DBCtrl_convExp1628(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 394
func DBCtrl_convExp1780(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 410
func DBCtrl_convExp1868(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 427
func DBCtrl_convExp2010(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 455
func DBCtrl_convExp2171(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 469
func DBCtrl_convExp2252(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 483
func DBCtrl_convExp2335(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 491
func DBCtrl_convExp2397(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 513
func DBCtrl_convExp2478(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 616
func DBCtrl_convExp3033(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 626
func DBCtrl_convExp3097(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 646
func DBCtrl_convExp3208(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 660
func DBCtrl_convExp3294(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 760
func DBCtrl_convExp3895(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 195
func DBCtrl_convExp1055(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 342
func DBCtrl_convExp1496(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 380
func DBCtrl_convExp1673(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 766
func DBCtrl_convExp3925(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}

func DBCtrl_getMaxId___anonymous_1034_(stmt string,msg string) {
}

func open___anonymous_1109_() string {
    return "open"
}
func open___anonymous_1112_() string {
    return "unknown version"
}

// 180: decl @lns.@tags.@DBCtrl.open
func DBCtrl_open(path string,readonly bool) LnsAny {
    __func__ := "@lns.@tags.@DBCtrl.open"
    Log_log(Log_Level__Log, __func__, 182, Log_CreateMessage(open___anonymous_1109_))
    
    var db *DBAccess_DBAccess
    
    {
        _db := DBAccess_open(path, readonly)
        if _db == nil{
            return nil
        } else {
            db = _db.(*DBAccess_DBAccess)
        }
    }
    db.FP.Exec("PRAGMA case_sensitive_like=ON;", nil)
    var dbCtrl *DBCtrl_DBCtrl
    dbCtrl = NewDBCtrl_DBCtrl(db, readonly)
    if Lns_op_not(readonly){
        dbCtrl.FP.Begin()
    }
    var item *DBCtrl_ETC
    
    {
        _item := DBCtrl_convExp1055(Lns_2DDD(DBCtrl_ETC__fromStem_1105_(dbCtrl.FP.GetRow("etc", "keyName = 'version'", nil, nil),nil)))
        if _item == nil{
            Log_log(Log_Level__Err, __func__, 196, Log_CreateMessage(open___anonymous_1112_))
            
            db.FP.Close()
            return nil
        } else {
            item = _item.(*DBCtrl_ETC)
        }
    }
    if Lns_tonumber(item.FP.Get_val(), nil) != DBCtrl_DB_VERSION{
        Log_log(Log_Level__Err, __func__, 201, Log_CreateMessage(func() string {
            return Lns_getVM().String_format("not support version. -- %s", []LnsAny{item.FP.Get_val()})
        }))
        
        db.FP.Close()
        return nil
    }
    if dbCtrl.FP.IsKilling(){
        panic("db is killed now")
    }
    dbCtrl.individualTypeFlag = dbCtrl.FP.EqualsEtc("individualTypeFlag", "1")
    
    dbCtrl.individualStructFlag = dbCtrl.FP.EqualsEtc("individualStructFlag", "1")
    
    dbCtrl.individualMacroFlag = dbCtrl.FP.EqualsEtc("individualMacroFlag", "1")
    
    return dbCtrl
}


// 340: decl @lns.@tags.@DBCtrl.getProjDir
func DBCtrl_getProjDir(path string,mod string) string {
    var workPath string
    workPath = Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(mod,"@", "")).(string),"%.", "/")).(string) + ".lns"
    var projDir string
    projDir = DBCtrl_convExp1496(Lns_2DDD(Lns_getVM().String_gsub(path,workPath + "$", "")))
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( len(mod) != 0) &&
        Lns_GetEnv().SetStackVal( len(projDir) == 0) ).(bool)){
        projDir = "./"
        
    }
    return projDir
}

// 349: decl @lns.@tags.@DBCtrl.normalizePath
func DBCtrl_normalizePath_1256_(path string) string {
    return Lns_car(Lns_getVM().String_gsub(path,"^%./", "")).(string)
}











func create___anonymous_1341_() string {
    return "create"
}
// 574: decl @lns.@tags.@DBCtrl.create
func DBCtrl_create_1339_(dbPath string) LnsAny {
    __func__ := "@lns.@tags.@DBCtrl.create"
    Log_log(Log_Level__Log, __func__, 576, Log_CreateMessage(create___anonymous_1341_))
    
    var db *DBAccess_DBAccess
    
    {
        _db := DBAccess_open(dbPath, false)
        if _db == nil{
            return nil
        } else {
            db = _db.(*DBAccess_DBAccess)
        }
    }
    var dbCtrl *DBCtrl_DBCtrl
    dbCtrl = NewDBCtrl_DBCtrl(db, false)
    dbCtrl.FP.creataTables()
    dbCtrl.FP.Begin()
    dbCtrl.FP.SetEtc("individualTypeFlag", "1")
    dbCtrl.FP.SetEtc("individualStructFlag", "1")
    dbCtrl.FP.SetEtc("individualMacroFlag", "1")
    return dbCtrl
}

// 594: decl @lns.@tags.@DBCtrl.initDB
func DBCtrl_initDB(dbPath string) {
    Lns_getVM().OS_remove(dbPath)
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_create_1339_(dbPath)
        if _db == nil{
            Lns_print([]LnsAny{"create error"})
            return 
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    db.FP.Commit()
    db.FP.Close()
}





func DBCtrl_dumpFile___anonymous_1381_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("id"), items.Get("dir")})
    return true
}
func DBCtrl_dumpFile___anonymous_1386_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("id"), items.Get("projId"), items.Get("path"), items.Get("mod")})
    return true
}
func DBCtrl_dumpFile___anonymous_1391_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("mainId"), items.Get("subId")})
    return true
}
func DBCtrl_dumpAll___anonymous_1398_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("id"), items.Get("name")})
    return true
}
func DBCtrl_dumpAll___anonymous_1403_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("nsId"), items.Get("superNsId")})
    return true
}
func DBCtrl_dumpAll___anonymous_1408_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("nsId"), items.Get("fileId"), items.Get("line"), items.Get("column")})
    return true
}
func DBCtrl_dumpAll___anonymous_1413_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("nsId"), items.Get("fileId"), items.Get("line"), items.Get("column")})
    return true
}
func DBCtrl_dumpAll___anonymous_1418_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Get("nsId"), items.Get("fileId"), items.Get("line"), items.Get("column")})
    return true
}
// 747: decl @lns.@tags.@DBCtrl.test
func DBCtrl_test() bool {
    var dbPath string
    dbPath = "lnstags.sqlite3"
    {
        DBCtrl_initDB(dbPath)
    }
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_open(dbPath, false)
        if _db == nil{
            Lns_print([]LnsAny{"open error"})
            return false
        } else {
            db = _db.(*DBCtrl_DBCtrl)
        }
    }
    var fileId LnsInt
    fileId = DBCtrl_rootNsId
    for _, _path := range( NewLnsList([]LnsAny{"aa.lns", "bb.lns", "cc.lns"}).Items ) {
        path := _path.(string)
        fileId = DBCtrl_convExp3895(Lns_2DDD(db.FP.AddFile(path, Lns_car(Lns_getVM().String_gsub(path,"%.lns", "")).(string))))
        
    }
    var parentId LnsInt
    parentId = DBCtrl_rootNsId
    for _index, _name := range( NewLnsList([]LnsAny{"@hoge", "@hoge.@foo", "@hoge.@foo.bar"}).Items ) {
        index := _index + 1
        name := _name.(string)
        var newid LnsInt
        newid = DBCtrl_convExp3925(Lns_2DDD(db.FP.AddNamespace(name, parentId)))
        db.FP.AddSymbolDecl(newid, fileId, 100 + index, index * 10)
        db.FP.AddSymbolRef(newid, fileId, 200 + index, index * 20, true)
        db.FP.AddSymbolSet(newid, fileId, 300 + index, index * 30)
        parentId = newid
        
    }
    {
        var added bool
        _,added = db.FP.AddNamespace("@hoge", DBCtrl_rootNsId)
        Lns_print([]LnsAny{"added", added})
    }
    db.FP.Commit()
    db.FP.DumpAll()
    return true
}

// declaration Class -- IdMgr
type DBCtrl_IdMgrMtd interface {
    getIdNext() LnsInt
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
func NewDBCtrl_IdMgr(arg1 LnsInt) *DBCtrl_IdMgr {
    obj := &DBCtrl_IdMgr{}
    obj.FP = obj
    obj.InitDBCtrl_IdMgr(arg1)
    return obj
}
// 15: DeclConstr
func (self *DBCtrl_IdMgr) InitDBCtrl_IdMgr(idNum LnsInt) {
    self.idNum = idNum
    
}

// 19: decl @lns.@tags.@DBCtrl.IdMgr.getIdNext
func (self *DBCtrl_IdMgr) getIdNext() LnsInt {
    var idNum LnsInt
    idNum = self.idNum
    self.idNum = self.idNum + 1
    
    return idNum
}


// declaration Class -- DBCtrl
type DBCtrl_DBCtrlMtd interface {
    AddFile(arg1 string, arg2 string)(LnsInt, bool)
    AddNamespace(arg1 string, arg2 LnsInt)(LnsInt, bool)
    AddOverride(arg1 LnsInt, arg2 LnsInt)
    AddProj(arg1 string)(LnsInt, bool)
    AddSubfile(arg1 LnsInt, arg2 LnsInt)
    AddSymbolDecl(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    AddSymbolRef(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 bool)
    AddSymbolSet(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    Begin()
    Close()
    Commit()
    creataTables()
    DumpAll()
    DumpFile()
    EqualsEtc(arg1 string, arg2 string) bool
    Exec(arg1 string, arg2 LnsAny)
    GetEtc(arg1 string) LnsAny
    GetFileIdFromPath(arg1 string) LnsInt
    GetFilePath(arg1 LnsInt) LnsAny
    GetMainFilePath(arg1 LnsInt) LnsAny
    getName(arg1 LnsInt) LnsAny
    getNsId(arg1 string) LnsAny
    GetProjId(arg1 string) LnsAny
    GetRow(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) LnsAny
    getRowList(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny) *LnsList
    Get_individualMacroFlag() bool
    Get_individualStructFlag() bool
    Get_individualTypeFlag() bool
    Get_projDir() string
    Insert(arg1 string, arg2 string)
    IsKilling() bool
    MapFilePath(arg1 DBCtrl_MapFileCallBack)
    MapNamespaceSuffix(arg1 string, arg2 DBCtrl_NameSpaceCallback)
    MapRowList(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 base.Base_queryMapForm, arg6 LnsAny) bool
    MapRowListSort(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 base.Base_queryMapForm, arg7 LnsAny) bool
    MapSymbolDecl(arg1 string, arg2 DBCtrl_callbackSymbolDecl)
    MapSymbolRef(arg1 string, arg2 bool, arg3 DBCtrl_callbackSymbolRef)
    SetEtc(arg1 string, arg2 string)
    Update(arg1 string, arg2 string, arg3 LnsAny)
}
type DBCtrl_DBCtrl struct {
    access *DBAccess_DBAccess
    individualTypeFlag bool
    individualStructFlag bool
    individualMacroFlag bool
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
func NewDBCtrl_DBCtrl(arg1 *DBAccess_DBAccess, arg2 bool) *DBCtrl_DBCtrl {
    obj := &DBCtrl_DBCtrl{}
    obj.FP = obj
    obj.InitDBCtrl_DBCtrl(arg1, arg2)
    return obj
}
func (self *DBCtrl_DBCtrl) Get_individualTypeFlag() bool{ return self.individualTypeFlag }
func (self *DBCtrl_DBCtrl) Get_individualStructFlag() bool{ return self.individualStructFlag }
func (self *DBCtrl_DBCtrl) Get_individualMacroFlag() bool{ return self.individualMacroFlag }
func (self *DBCtrl_DBCtrl) Get_projDir() string{ return self.projDir }
// 39: decl @lns.@tags.@DBCtrl.DBCtrl.getMaxId
func DBCtrl_DBCtrl_getMaxId_1027_(access *DBAccess_DBAccess,tableName string,defId LnsInt) LnsInt {
    var id LnsAny
    id = nil
    access.FP.MapRowList(tableName, nil, 1, "MAX(id)", nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        id = items.Get("id")
        
        return false
    }), base.Base_errHandleForm(DBCtrl_getMaxId___anonymous_1034_))
    if id != nil{
        id_54 := id
        return Lns_forceCastInt(id_54)
    }
    return defId
}

// 56: DeclConstr
func (self *DBCtrl_DBCtrl) InitDBCtrl_DBCtrl(access *DBAccess_DBAccess,readonly bool) {
    self.access = access
    
    self.individualTypeFlag = false
    
    self.individualStructFlag = false
    
    self.individualMacroFlag = false
    
    self.projDir = Depend_getCurDir()
    
    self.idMgrNamespace = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1027_(access, "namespace", DBCtrl_userNsId))
    
    self.idMgrSimpleName = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1027_(access, "simpleName", DBCtrl_userNsId))
    
    self.idMgrFilePath = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1027_(access, "filePath", DBCtrl_userNsId))
    
    self.idMgrProjInfo = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1027_(access, "projInfo", DBCtrl_userNsId))
    
}

// 73: decl @lns.@tags.@DBCtrl.DBCtrl.close
func (self *DBCtrl_DBCtrl) Close() {
    self.access.FP.Close()
}

// 77: decl @lns.@tags.@DBCtrl.DBCtrl.begin
func (self *DBCtrl_DBCtrl) Begin() {
    self.access.FP.Begin()
}

// 80: decl @lns.@tags.@DBCtrl.DBCtrl.commit
func (self *DBCtrl_DBCtrl) Commit() {
    self.access.FP.Commit()
}

// 84: decl @lns.@tags.@DBCtrl.DBCtrl.exec
func (self *DBCtrl_DBCtrl) Exec(stmt string,errHandle LnsAny) {
    self.access.FP.Exec(stmt, errHandle)
}

// 88: decl @lns.@tags.@DBCtrl.DBCtrl.insert
func (self *DBCtrl_DBCtrl) Insert(tableName string,values string) {
    self.access.FP.Insert(tableName, values)
}

// 92: decl @lns.@tags.@DBCtrl.DBCtrl.update
func (self *DBCtrl_DBCtrl) Update(tableName string,set string,condition LnsAny) {
    var sql string
    sql = Lns_getVM().String_format("UPDATE %s SET %s", []LnsAny{tableName, set})
    if Lns_isCondTrue( condition){
        sql = Lns_getVM().String_format("%s WHERE %s", []LnsAny{sql, condition})
        
    }
    self.FP.Exec(sql, nil)
}

// 102: decl @lns.@tags.@DBCtrl.DBCtrl.mapRowList
func (self *DBCtrl_DBCtrl) MapRowList(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapRowList(tableName, condition, limit, attrib, nil, _func, errHandle)
}

// 110: decl @lns.@tags.@DBCtrl.DBCtrl.mapRowListSort
func (self *DBCtrl_DBCtrl) MapRowListSort(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,order LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapRowList(tableName, condition, limit, attrib, order, _func, errHandle)
}

// 118: decl @lns.@tags.@DBCtrl.DBCtrl.getRowList
func (self *DBCtrl_DBCtrl) getRowList(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,errHandle LnsAny) *LnsList {
    var rows *LnsList
    rows = NewLnsList([]LnsAny{})
    self.FP.MapRowList(tableName, condition, limit, attrib, base.Base_queryMapForm(func(items *LnsMap) bool {
        rows.Insert(items)
        return true
    }), errHandle)
    return rows
}

// 131: decl @lns.@tags.@DBCtrl.DBCtrl.getRow
func (self *DBCtrl_DBCtrl) GetRow(tableName string,condition LnsAny,attrib LnsAny,errHandle LnsAny) LnsAny {
    var row *LnsList
    row = self.FP.getRowList(tableName, condition, 1, attrib, errHandle)
    if row.Len() == 0{
        return nil
    }
    return row.GetAt(1).(*LnsMap)
}

// 141: decl @lns.@tags.@DBCtrl.DBCtrl.getEtc
func (self *DBCtrl_DBCtrl) GetEtc(key string) LnsAny {
    {
        _etc := self.FP.GetRow("etc", "keyName = '" + key + "'", nil, nil)
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

// 150: decl @lns.@tags.@DBCtrl.DBCtrl.setEtc
func (self *DBCtrl_DBCtrl) SetEtc(key string,val string) {
    var keyTxt string
    keyTxt = Lns_getVM().String_format("keyName = '%s'", []LnsAny{key})
    var valTxt string
    valTxt = Lns_getVM().String_format("val = '%s'", []LnsAny{val})
    if Lns_op_not(self.FP.GetEtc(key)){
        self.FP.Insert("etc", Lns_getVM().String_format("'%s', '%s'", []LnsAny{key, val}))
    } else { 
        self.FP.Update("etc", valTxt, keyTxt)
    }
}

// 160: decl @lns.@tags.@DBCtrl.DBCtrl.equalsEtc
func (self *DBCtrl_DBCtrl) EqualsEtc(key string,val string) bool {
    if self.FP.GetEtc(key) == val{
        return true
    }
    return false
}

// 167: decl @lns.@tags.@DBCtrl.DBCtrl.isKilling
func (self *DBCtrl_DBCtrl) IsKilling() bool {
    if self.FP.EqualsEtc("killFlag", "1"){
        return true
    }
    return false
}

// 218: decl @lns.@tags.@DBCtrl.DBCtrl.creataTables
func (self *DBCtrl_DBCtrl) creataTables() {
    self.FP.Exec(Lns_getVM().String_format("BEGIN;\nCREATE TABLE etc ( keyName VARCHAR UNIQUE COLLATE binary PRIMARY KEY, val VARCHAR);\nINSERT INTO etc VALUES( 'version', '%d' );\nINSERT INTO etc VALUES( 'projDir', '' );\nINSERT INTO etc VALUES( 'individualStructFlag', '0' );\nINSERT INTO etc VALUES( 'individualTypeFlag', '0' );\nINSERT INTO etc VALUES( 'individualMacroFlag', '0' );\nINSERT INTO etc VALUES( 'killFlag', '0' );\nCREATE TABLE namespace ( id INTEGER PRIMARY KEY, snameId INTEGER, parentId INTEGER, digest CHAR(32), name VARCHAR UNIQUE COLLATE binary, otherName VARCHAR COLLATE binary, virtual INTEGER);\nINSERT INTO namespace VALUES( NULL, 1, 0, '', '', '', 0 );\n\nCREATE TABLE simpleName ( id INTEGER PRIMARY KEY, name VARCHAR UNIQUE COLLATE binary);\nCREATE TABLE projInfo ( id INTEGER PRIMARY KEY, dir VARCHAR UNIQUE COLLATE binary);\nINSERT INTO projInfo VALUES( 1, '' );\nCREATE TABLE filePath ( id INTEGER PRIMARY KEY, path VARCHAR UNIQUE COLLATE binary, mod VARCHAR COLLATE binary, projId INTEGER );\nINSERT INTO filePath VALUES( 1, '', '', 1 );\nCREATE TABLE subfile ( mainId INTEGER, subId INTEGER PRIMARY KEY );\n\nCREATE TABLE override (nsId INTEGER, superNsId INTEGER, PRIMARY KEY (nsId, superNsId));\nCREATE TABLE symbolDecl ( nsId INTEGER, snameId INTEGER, parentId INTEGER, type INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, comment VARCHAR COLLATE binary, hasBodyFlag INTEGER, PRIMARY KEY( nsId, fileId, line ) );\nINSERT INTO symbolDecl VALUES( 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, '', 0 );\n\nCREATE TABLE symbolRef ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, setOp INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );\nCREATE TABLE symbolSet ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );\n\nCREATE TABLE funcCall ( nsId INTEGER, snameId INTEGER, belongNsId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, PRIMARY KEY( nsId, belongNsId ) );\nCREATE TABLE incRef ( id INTEGER, baseFileId INTEGER, line INTEGER );\nCREATE TABLE incCache ( id INTEGER, baseFileId INTEGER, incFlag INTEGER, PRIMARY KEY( id, baseFileId ) );\nCREATE TABLE incBelong ( id INTEGER, baseFileId INTEGER, nsId INTEGER, PRIMARY KEY ( id, nsId ) );\nCREATE INDEX index_ns ON namespace ( id, snameId, parentId, name, otherName );\nCREATE INDEX index_sName ON simpleName ( id, name );\nCREATE INDEX index_filePath ON filePath ( id, path, mod );\nCREATE INDEX index_subfile ON subfile (subId );\nCREATE INDEX index_override ON override (nsId, superNsId);\nCREATE INDEX index_symDecl ON symbolDecl ( nsId, parentId, snameId, fileId );\nCREATE INDEX index_symRef ON symbolRef ( nsId, snameId, fileId, belongNsId );\nCREATE INDEX index_incRef ON incRef ( id, baseFileId );\nCREATE INDEX index_incCache ON incCache ( id, baseFileId, incFlag );\nCREATE INDEX index_incBelong ON incBelong ( id, baseFileId );\nCOMMIT;\n", []LnsAny{(LnsInt)(DBCtrl_DB_VERSION)}), nil)
}

// 306: decl @lns.@tags.@DBCtrl.DBCtrl.getProjId
func (self *DBCtrl_DBCtrl) GetProjId(path string) LnsAny {
    var projId LnsAny
    projId = nil
    self.FP.MapRowList("projInfo", Lns_getVM().String_format("dir = '%s'", []LnsAny{path}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _projInfo := DBCtrl_convExp1371(Lns_2DDD(DBCtrl_ItemProjInfo__fromStem_1134_(items,nil)))
            if !Lns_IsNil( _projInfo ) {
                projInfo := _projInfo.(*DBCtrl_ItemProjInfo)
                projId = projInfo.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    return projId
}

// 329: decl @lns.@tags.@DBCtrl.DBCtrl.addProj
func (self *DBCtrl_DBCtrl) AddProj(path string)(LnsInt, bool) {
    {
        _projId := self.FP.GetProjId(path)
        if !Lns_IsNil( _projId ) {
            projId := _projId.(LnsInt)
            return projId, false
        }
    }
    var id LnsInt
    id = self.idMgrProjInfo.FP.getIdNext()
    self.FP.Insert("projInfo", Lns_getVM().String_format("%d,'%s'", []LnsAny{id, path}))
    return id, true
}

// 362: decl @lns.@tags.@DBCtrl.DBCtrl.addFile
func (self *DBCtrl_DBCtrl) AddFile(path string,mod string)(LnsInt, bool) {
    path = DBCtrl_normalizePath_1256_(path)
    
    var fileId LnsAny
    fileId = nil
    self.FP.MapRowList("filePath", Lns_getVM().String_format("path = '%s'", []LnsAny{path}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp1628(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(items,nil)))
            if !Lns_IsNil( _filePath ) {
                filePath := _filePath.(*DBCtrl_ItemFilePath)
                fileId = filePath.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    if fileId != nil{
        fileId_319 := fileId.(LnsInt)
        return fileId_319, false
    }
    var projDir string
    projDir = DBCtrl_getProjDir(path, mod)
    var projId LnsInt
    projId = DBCtrl_convExp1673(Lns_2DDD(self.FP.AddProj(projDir)))
    var id LnsInt
    id = self.idMgrFilePath.FP.getIdNext()
    self.FP.Insert("filePath", Lns_getVM().String_format("%d,'%s','%s', %d", []LnsAny{id, path, Lns_car(Lns_getVM().String_gsub(mod,"@", "")).(string), projId}))
    return id, true
}

// 390: decl @lns.@tags.@DBCtrl.DBCtrl.mapFilePath
func (self *DBCtrl_DBCtrl) MapFilePath(callback DBCtrl_MapFileCallBack) {
    self.FP.MapRowList("filePath", nil, nil, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp1780(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(items,nil)))
            if !Lns_IsNil( _filePath ) {
                filePath := _filePath.(*DBCtrl_ItemFilePath)
                if Lns_op_not(callback(filePath)){
                    return false
                }
            }
        }
        return true
    }), nil)
}

// 404: decl @lns.@tags.@DBCtrl.DBCtrl.getFileIdFromPath
func (self *DBCtrl_DBCtrl) GetFileIdFromPath(path string) LnsInt {
    __func__ := "@lns.@tags.@DBCtrl.DBCtrl.getFileIdFromPath"
    path = DBCtrl_normalizePath_1256_(path)
    
    var fileId LnsAny
    fileId = nil
    self.FP.MapRowList("filePath", Lns_getVM().String_format("path = '%s'", []LnsAny{path}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp1868(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(items,nil)))
            if !Lns_IsNil( _filePath ) {
                filePath := _filePath.(*DBCtrl_ItemFilePath)
                fileId = filePath.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    if fileId != nil{
        fileId_348 := fileId.(LnsInt)
        return fileId_348
    }
    Log_log(Log_Level__Err, __func__, 418, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("not found file -- %s", []LnsAny{path})
    }))
    
    return DBCtrl_rootNsId
}

// 422: decl @lns.@tags.@DBCtrl.DBCtrl.getFilePath
func (self *DBCtrl_DBCtrl) GetFilePath(fileId LnsInt) LnsAny {
    var path LnsAny
    path = nil
    self.FP.MapRowList("filePath", Lns_getVM().String_format("id = %d", []LnsAny{fileId}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp2010(Lns_2DDD(DBCtrl_ItemFilePath__fromStem(items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemFilePath)
                path = obj.FP.Get_path()
                
            }
        }
        return false
    }), nil)
    return path
}

// 435: decl @lns.@tags.@DBCtrl.DBCtrl.addSubfile
func (self *DBCtrl_DBCtrl) AddSubfile(mainId LnsInt,subId LnsInt) {
    self.FP.Insert("subfile", Lns_getVM().String_format("%d, %d", []LnsAny{mainId, subId}))
}

// 440: decl @lns.@tags.@DBCtrl.DBCtrl.getMainFilePath
func (self *DBCtrl_DBCtrl) GetMainFilePath(subId LnsInt) LnsAny {
    {
        __map := self.FP.GetRow("subfile", Lns_getVM().String_format("subId = %d", []LnsAny{subId}), "mainId", nil)
        if !Lns_IsNil( __map ) {
            _map := __map.(*LnsMap)
            {
                _mainId := _map.Get("mainId")
                if !Lns_IsNil( _mainId ) {
                    mainId := _mainId
                    return self.FP.GetFilePath(Lns_forceCastInt(mainId))
                }
            }
        }
    }
    return nil
}

// 450: decl @lns.@tags.@DBCtrl.DBCtrl.getName
func (self *DBCtrl_DBCtrl) getName(nsId LnsInt) LnsAny {
    var name LnsAny
    name = nil
    self.FP.MapRowList("namespace", Lns_getVM().String_format("id = %d", []LnsAny{nsId}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp2171(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemNamespace)
                name = obj.FP.Get_name()
                
            }
        }
        return false
    }), nil)
    return name
}

// 464: decl @lns.@tags.@DBCtrl.DBCtrl.getNsId
func (self *DBCtrl_DBCtrl) getNsId(name string) LnsAny {
    var nsId LnsAny
    nsId = nil
    self.FP.MapRowList("namespace", Lns_getVM().String_format("name = '%s'", []LnsAny{name}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp2252(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemNamespace)
                nsId = obj.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    return nsId
}

// 478: decl @lns.@tags.@DBCtrl.DBCtrl.mapNamespaceSuffix
func (self *DBCtrl_DBCtrl) MapNamespaceSuffix(suffix string,callback DBCtrl_NameSpaceCallback) {
    self.FP.MapRowList("namespace", Lns_getVM().String_format("name like '%%.%s' escape '$'", []LnsAny{suffix}), nil, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp2335(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemNamespace)
                return callback(item)
            }
        }
        return true
    }), nil)
    self.FP.MapRowList("namespace", Lns_getVM().String_format("name = '%s'", []LnsAny{suffix}), nil, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp2397(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemNamespace)
                return callback(item)
            }
        }
        return true
    }), nil)
}

// 507: decl @lns.@tags.@DBCtrl.DBCtrl.addNamespace
func (self *DBCtrl_DBCtrl) AddNamespace(fullName string,parentId LnsInt)(LnsInt, bool) {
    var id LnsAny
    id = nil
    self.FP.MapRowList("namespace", Lns_getVM().String_format("name = '%s'", []LnsAny{fullName}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp2478(Lns_2DDD(DBCtrl_ItemNamespace__fromStem(items,nil)))
            if !Lns_IsNil( _obj ) {
                obj := _obj.(*DBCtrl_ItemNamespace)
                id = obj.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    if id != nil{
        id_430 := id.(LnsInt)
        return id_430, false
    }
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var newId LnsInt
    newId = self.idMgrNamespace.FP.getIdNext()
    self.FP.Insert("namespace", Lns_getVM().String_format("%d, %d, %d, '', '%s', '', 1", []LnsAny{newId, snid, parentId, fullName}))
    return newId, true
}

// 529: decl @lns.@tags.@DBCtrl.DBCtrl.addOverride
func (self *DBCtrl_DBCtrl) AddOverride(nsId LnsInt,superNsId LnsInt) {
    self.FP.Insert("override", Lns_getVM().String_format("%d, %d", []LnsAny{nsId, superNsId}))
}

// 534: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolDecl
func (self *DBCtrl_DBCtrl) AddSymbolDecl(nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var kind LnsInt
    kind = 0
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var parentId LnsAny
    
    {
        _parentId := Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
        self.FP.GetRow("namespace", Lns_getVM().String_format("id = %d", []LnsAny{nsId}), "parentId", nil)) && 
        Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsMap).Get("parentId")))
        if _parentId == nil{
            return 
        } else {
            parentId = _parentId
        }
    }
    self.FP.Insert("symbolDecl", Lns_getVM().String_format("%d, %d, %d, %d, %d, %d, %d, %d, %d, 0, '', 0", []LnsAny{nsId, snid, Lns_forceCastInt(parentId), kind, fileId, lineNo, column, lineNo, column}))
}

// 552: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolRef
func (self *DBCtrl_DBCtrl) AddSymbolRef(nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt,setOp bool) {
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var belongNsId LnsInt
    belongNsId = DBCtrl_rootNsId
    self.FP.Insert("symbolRef", Lns_getVM().String_format("%d, %d, %d, %d, %d, %d, %d", []LnsAny{nsId, snid, fileId, lineNo, column, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( setOp) &&
        Lns_GetEnv().SetStackVal( 1) ||
        Lns_GetEnv().SetStackVal( 0) ).(LnsInt), belongNsId}))
}

// 562: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolSet
func (self *DBCtrl_DBCtrl) AddSymbolSet(nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var belongNsId LnsInt
    belongNsId = DBCtrl_rootNsId
    self.FP.Insert("symbolSet", Lns_getVM().String_format("%d, %d, %d, %d, %d, %d", []LnsAny{nsId, snid, fileId, lineNo, column, belongNsId}))
}

// 605: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolDecl
func (self *DBCtrl_DBCtrl) MapSymbolDecl(name string,callback DBCtrl_callbackSymbolDecl) {
    var nsId LnsInt
    
    {
        _nsId := self.FP.getNsId(name)
        if _nsId == nil{
            return 
        } else {
            nsId = _nsId.(LnsInt)
        }
    }
    var overrideStr string
    overrideStr = Lns_getVM().String_format("%d", []LnsAny{nsId})
    self.FP.MapRowList("override", Lns_getVM().String_format("superNsId = %d", []LnsAny{nsId}), nil, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3033(Lns_2DDD(DBCtrl_ItemOverride__fromStem(items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemOverride)
                overrideStr = Lns_getVM().String_format("%s, %d", []LnsAny{overrideStr, item.FP.Get_nsId()})
                
            }
        }
        return true
    }), nil)
    self.FP.MapRowListSort("symbolDecl", Lns_getVM().String_format("nsId IN (%s)", []LnsAny{overrideStr}), nil, nil, "fileId, line", base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3097(Lns_2DDD(DBCtrl_ItemSymbolDecl__fromStem(items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemSymbolDecl)
                return callback(item)
            }
        }
        return true
    }), nil)
}

// 635: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolRef
func (self *DBCtrl_DBCtrl) MapSymbolRef(name string,onlySet bool,callback DBCtrl_callbackSymbolRef) {
    var nsId LnsInt
    
    {
        _nsId := self.FP.getNsId(name)
        if _nsId == nil{
            return 
        } else {
            nsId = _nsId.(LnsInt)
        }
    }
    var overrideStr string
    overrideStr = Lns_getVM().String_format("%d", []LnsAny{nsId})
    self.FP.MapRowList("override", Lns_getVM().String_format("nsId = %d", []LnsAny{nsId}), nil, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3208(Lns_2DDD(DBCtrl_ItemOverride__fromStem(items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemOverride)
                overrideStr = Lns_getVM().String_format("%s, %d", []LnsAny{overrideStr, item.FP.Get_superNsId()})
                
            }
        }
        return true
    }), nil)
    var cond string
    cond = Lns_getVM().String_format("nsId IN (%s)", []LnsAny{overrideStr})
    if onlySet{
        cond = Lns_getVM().String_format("(%s) AND setOp = 1", []LnsAny{cond})
        
    }
    self.FP.MapRowListSort("symbolRef", cond, nil, nil, "fileId, line", base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp3294(Lns_2DDD(DBCtrl_ItemSymbolRef__fromStem(items,nil)))
            if !Lns_IsNil( _item ) {
                item := _item.(*DBCtrl_ItemSymbolRef)
                return callback(item)
            }
        }
        return true
    }), nil)
}

// 667: decl @lns.@tags.@DBCtrl.DBCtrl.dumpFile
func (self *DBCtrl_DBCtrl) DumpFile() {
    Lns_print([]LnsAny{"projId"})
    self.FP.MapRowList("projInfo", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpFile___anonymous_1381_), nil)
    Lns_print([]LnsAny{"filePath"})
    self.FP.MapRowList("filePath", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpFile___anonymous_1386_), nil)
    Lns_print([]LnsAny{"subfile"})
    self.FP.MapRowList("subfile", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpFile___anonymous_1391_), nil)
}

// 697: decl @lns.@tags.@DBCtrl.DBCtrl.dumpAll
func (self *DBCtrl_DBCtrl) DumpAll() {
    self.FP.DumpFile()
    Lns_print([]LnsAny{"namespace"})
    self.FP.MapRowList("namespace", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1398_), nil)
    Lns_print([]LnsAny{"override"})
    self.FP.MapRowList("override", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1403_), nil)
    Lns_print([]LnsAny{"symbolDecl"})
    self.FP.MapRowList("symbolDecl", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1408_), nil)
    Lns_print([]LnsAny{"symbolRef"})
    self.FP.MapRowList("symbolRef", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1413_), nil)
    Lns_print([]LnsAny{"symbolSet"})
    self.FP.MapRowList("symbolSet", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1418_), nil)
}


// declaration Class -- ETC
type DBCtrl_ETCMtd interface {
    ToMap() *LnsMap
    Get_keyName() string
    Get_val() string
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
func NewDBCtrl_ETC(arg1 string, arg2 string) *DBCtrl_ETC {
    obj := &DBCtrl_ETC{}
    obj.FP = obj
    obj.InitDBCtrl_ETC(arg1, arg2)
    return obj
}
func (self *DBCtrl_ETC) InitDBCtrl_ETC(arg1 string, arg2 string) {
    self.keyName = arg1
    self.val = arg2
}
func (self *DBCtrl_ETC) Get_keyName() string{ return self.keyName }
func (self *DBCtrl_ETC) Get_val() string{ return self.val }
func (self *DBCtrl_ETC) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["keyName"] = Lns_ToCollection( self.keyName )
    obj.Items["val"] = Lns_ToCollection( self.val )
    return obj
}
func (self *DBCtrl_ETC) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ETC__fromMap_1103_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ETC_FromMap( arg1, paramList )
}
func DBCtrl_ETC__fromStem_1105_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    Get_dir() string
    Get_id() LnsInt
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
func NewDBCtrl_ItemProjInfo(arg1 LnsInt, arg2 string) *DBCtrl_ItemProjInfo {
    obj := &DBCtrl_ItemProjInfo{}
    obj.FP = obj
    obj.InitDBCtrl_ItemProjInfo(arg1, arg2)
    return obj
}
func (self *DBCtrl_ItemProjInfo) InitDBCtrl_ItemProjInfo(arg1 LnsInt, arg2 string) {
    self.id = arg1
    self.dir = arg2
}
func (self *DBCtrl_ItemProjInfo) Get_id() LnsInt{ return self.id }
func (self *DBCtrl_ItemProjInfo) Get_dir() string{ return self.dir }
func (self *DBCtrl_ItemProjInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["id"] = Lns_ToCollection( self.id )
    obj.Items["dir"] = Lns_ToCollection( self.dir )
    return obj
}
func (self *DBCtrl_ItemProjInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemProjInfo__fromMap_1132_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemProjInfo_FromMap( arg1, paramList )
}
func DBCtrl_ItemProjInfo__fromStem_1134_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    Get_id() LnsInt
    Get_mod() string
    Get_path() string
    Get_projId() LnsInt
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
func NewDBCtrl_ItemFilePath(arg1 LnsInt, arg2 string, arg3 string, arg4 LnsInt) *DBCtrl_ItemFilePath {
    obj := &DBCtrl_ItemFilePath{}
    obj.FP = obj
    obj.InitDBCtrl_ItemFilePath(arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemFilePath) InitDBCtrl_ItemFilePath(arg1 LnsInt, arg2 string, arg3 string, arg4 LnsInt) {
    self.id = arg1
    self.path = arg2
    self.mod = arg3
    self.projId = arg4
}
func (self *DBCtrl_ItemFilePath) Get_id() LnsInt{ return self.id }
func (self *DBCtrl_ItemFilePath) Get_path() string{ return self.path }
func (self *DBCtrl_ItemFilePath) Get_mod() string{ return self.mod }
func (self *DBCtrl_ItemFilePath) Get_projId() LnsInt{ return self.projId }
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
func DBCtrl_ItemFilePath__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemFilePath_FromMap( arg1, paramList )
}
func DBCtrl_ItemFilePath__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    Get_id() LnsInt
    Get_name() string
    Get_parentId() LnsInt
    Get_snameId() LnsInt
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
func NewDBCtrl_ItemNamespace(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 string) *DBCtrl_ItemNamespace {
    obj := &DBCtrl_ItemNamespace{}
    obj.FP = obj
    obj.InitDBCtrl_ItemNamespace(arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemNamespace) InitDBCtrl_ItemNamespace(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 string) {
    self.id = arg1
    self.snameId = arg2
    self.parentId = arg3
    self.name = arg4
}
func (self *DBCtrl_ItemNamespace) Get_id() LnsInt{ return self.id }
func (self *DBCtrl_ItemNamespace) Get_snameId() LnsInt{ return self.snameId }
func (self *DBCtrl_ItemNamespace) Get_parentId() LnsInt{ return self.parentId }
func (self *DBCtrl_ItemNamespace) Get_name() string{ return self.name }
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
func DBCtrl_ItemNamespace__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemNamespace_FromMap( arg1, paramList )
}
func DBCtrl_ItemNamespace__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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

// declaration Class -- ItemSymbolDecl
type DBCtrl_ItemSymbolDeclMtd interface {
    ToMap() *LnsMap
    Get_column() LnsInt
    Get_fileId() LnsInt
    Get_line() LnsInt
    Get_nsId() LnsInt
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
func NewDBCtrl_ItemSymbolDecl(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) *DBCtrl_ItemSymbolDecl {
    obj := &DBCtrl_ItemSymbolDecl{}
    obj.FP = obj
    obj.InitDBCtrl_ItemSymbolDecl(arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemSymbolDecl) InitDBCtrl_ItemSymbolDecl(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) {
    self.nsId = arg1
    self.fileId = arg2
    self.line = arg3
    self.column = arg4
}
func (self *DBCtrl_ItemSymbolDecl) Get_nsId() LnsInt{ return self.nsId }
func (self *DBCtrl_ItemSymbolDecl) Get_fileId() LnsInt{ return self.fileId }
func (self *DBCtrl_ItemSymbolDecl) Get_line() LnsInt{ return self.line }
func (self *DBCtrl_ItemSymbolDecl) Get_column() LnsInt{ return self.column }
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
func DBCtrl_ItemSymbolDecl__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemSymbolDecl_FromMap( arg1, paramList )
}
func DBCtrl_ItemSymbolDecl__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    Get_column() LnsInt
    Get_fileId() LnsInt
    Get_line() LnsInt
    Get_nsId() LnsInt
    Get_setOp() LnsInt
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
func NewDBCtrl_ItemSymbolRef(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 LnsInt) *DBCtrl_ItemSymbolRef {
    obj := &DBCtrl_ItemSymbolRef{}
    obj.FP = obj
    obj.InitDBCtrl_ItemSymbolRef(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *DBCtrl_ItemSymbolRef) InitDBCtrl_ItemSymbolRef(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 LnsInt) {
    self.nsId = arg1
    self.fileId = arg2
    self.line = arg3
    self.column = arg4
    self.setOp = arg5
}
func (self *DBCtrl_ItemSymbolRef) Get_nsId() LnsInt{ return self.nsId }
func (self *DBCtrl_ItemSymbolRef) Get_fileId() LnsInt{ return self.fileId }
func (self *DBCtrl_ItemSymbolRef) Get_line() LnsInt{ return self.line }
func (self *DBCtrl_ItemSymbolRef) Get_column() LnsInt{ return self.column }
func (self *DBCtrl_ItemSymbolRef) Get_setOp() LnsInt{ return self.setOp }
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
func DBCtrl_ItemSymbolRef__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemSymbolRef_FromMap( arg1, paramList )
}
func DBCtrl_ItemSymbolRef__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    Get_nsId() LnsInt
    Get_superNsId() LnsInt
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
func NewDBCtrl_ItemOverride(arg1 LnsInt, arg2 LnsInt) *DBCtrl_ItemOverride {
    obj := &DBCtrl_ItemOverride{}
    obj.FP = obj
    obj.InitDBCtrl_ItemOverride(arg1, arg2)
    return obj
}
func (self *DBCtrl_ItemOverride) InitDBCtrl_ItemOverride(arg1 LnsInt, arg2 LnsInt) {
    self.nsId = arg1
    self.superNsId = arg2
}
func (self *DBCtrl_ItemOverride) Get_nsId() LnsInt{ return self.nsId }
func (self *DBCtrl_ItemOverride) Get_superNsId() LnsInt{ return self.superNsId }
func (self *DBCtrl_ItemOverride) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["superNsId"] = Lns_ToCollection( self.superNsId )
    return obj
}
func (self *DBCtrl_ItemOverride) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemOverride__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemOverride_FromMap( arg1, paramList )
}
func DBCtrl_ItemOverride__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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

func Lns_DBCtrl_init() {
    if init_DBCtrl { return }
    init_DBCtrl = true
    DBCtrl__mod__ = "@lns.@tags.@DBCtrl"
    Lns_InitMod()
    base.Lns_base_init()
    Lns_DBAccess_init()
    Lns_Log_init()
    Lns_Depend_init()
    DBCtrl_rootNsId = 1
    DBCtrl_userNsId = 2
    _ = 1
    DBCtrl_DB_VERSION = 9.0
}
func init() {
    init_DBCtrl = false
}
