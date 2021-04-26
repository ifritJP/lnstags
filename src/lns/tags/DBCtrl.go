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
type DBCtrl_callbackSymbolDecl func (arg1 *DBCtrl_ItemSymbolDecl) bool
type DBCtrl_callbackSymbolRef func (arg1 *DBCtrl_ItemSymbolRef) bool
// for 315
func DBCtrl_convExp1362(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 334
func DBCtrl_convExp1478(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 348
func DBCtrl_convExp1559(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 370
func DBCtrl_convExp1645(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 391
func DBCtrl_convExp1773(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 486
func DBCtrl_convExp2273(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 496
func DBCtrl_convExp2337(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 516
func DBCtrl_convExp2445(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 526
func DBCtrl_convExp2509(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 190
func DBCtrl_convExp1017(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 609
func DBCtrl_convExp2995(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}

func DBCtrl_getMaxId___anonymous_1050_(stmt string,msg string) {
}

func open___anonymous_1157_() string {
    return "open"
}
func open___anonymous_1161_() string {
    return "unknown version"
}

// 177: decl @lns.@tags.@DBCtrl.open
func DBCtrl_open(path string,readonly bool) LnsAny {
    __func__ := "@lns.@tags.@DBCtrl.open"
    Log_log(Log_Level__Log, __func__, 179, Log_CreateMessage(open___anonymous_1157_))
    
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
    dbCtrl = NewDBCtrl_DBCtrl(db, readonly)
    if Lns_op_not(readonly){
        dbCtrl.FP.Begin()
    }
    var item *DBCtrl_ETC
    
    {
        _item := DBCtrl_convExp1017(Lns_2DDD(DBCtrl_ETC__fromStem_1151_(dbCtrl.FP.GetRow("etc", "keyName = 'version'", nil, nil),nil)))
        if _item == nil{
            Log_log(Log_Level__Err, __func__, 191, Log_CreateMessage(open___anonymous_1161_))
            
            db.FP.Close()
            return nil
        } else {
            item = _item.(*DBCtrl_ETC)
        }
    }
    if Lns_tonumber(item.FP.Get_val(), nil) != DBCtrl_DB_VERSION{
        Log_log(Log_Level__Err, __func__, 196, Log_CreateMessage(func() string {
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






func create___anonymous_1385_() string {
    return "create"
}
// 444: decl @lns.@tags.@DBCtrl.create
func DBCtrl_create_1382_(dbPath string) LnsAny {
    __func__ := "@lns.@tags.@DBCtrl.create"
    Log_log(Log_Level__Log, __func__, 446, Log_CreateMessage(create___anonymous_1385_))
    
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

// 464: decl @lns.@tags.@DBCtrl.initDB
func DBCtrl_initDB(dbPath string) {
    Lns_getVM().OS_remove(dbPath)
    var db *DBCtrl_DBCtrl
    
    {
        _db := DBCtrl_create_1382_(dbPath)
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





func DBCtrl_dumpAll___anonymous_1441_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Items["id"], items.Items["path"]})
    return true
}
func DBCtrl_dumpAll___anonymous_1448_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Items["id"], items.Items["name"]})
    return true
}
func DBCtrl_dumpAll___anonymous_1455_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Items["nsId"], items.Items["superNsId"]})
    return true
}
func DBCtrl_dumpAll___anonymous_1462_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Items["nsId"], items.Items["fileId"], items.Items["line"], items.Items["column"]})
    return true
}
func DBCtrl_dumpAll___anonymous_1469_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Items["nsId"], items.Items["fileId"], items.Items["line"], items.Items["column"]})
    return true
}
func DBCtrl_dumpAll___anonymous_1476_(items *LnsMap) bool {
    Lns_print([]LnsAny{items.Items["nsId"], items.Items["fileId"], items.Items["line"], items.Items["column"]})
    return true
}
// 590: decl @lns.@tags.@DBCtrl.test
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
        fileId = db.FP.AddFile(path)
        
    }
    var parentId LnsInt
    parentId = DBCtrl_rootNsId
    for _index, _name := range( NewLnsList([]LnsAny{"@hoge", "@hoge.@foo", "@hoge.@foo.bar"}).Items ) {
        index := _index + 1
        name := _name.(string)
        var newid LnsInt
        newid = DBCtrl_convExp2995(Lns_2DDD(db.FP.AddNamespace(name, parentId)))
        db.FP.AddSymbolDecl(newid, fileId, 100 + index, index * 10)
        db.FP.AddSymbolRef(newid, fileId, 200 + index, index * 20)
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
    AddFile(arg1 string) LnsInt
    AddNamespace(arg1 string, arg2 LnsInt)(LnsInt, bool)
    AddOverride(arg1 LnsInt, arg2 LnsInt)
    AddSymbolDecl(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    AddSymbolRef(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    AddSymbolSet(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt)
    Begin()
    Close()
    Commit()
    creataTables()
    DumpAll()
    EqualsEtc(arg1 string, arg2 string) bool
    Exec(arg1 string, arg2 LnsAny)
    GetEtc(arg1 string) LnsAny
    GetFilePath(arg1 LnsInt) LnsAny
    getName(arg1 LnsInt) LnsAny
    getNsId(arg1 string) LnsAny
    GetRow(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) LnsAny
    getRowList(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny) *LnsList
    Get_individualMacroFlag() bool
    Get_individualStructFlag() bool
    Get_individualTypeFlag() bool
    Get_projDir() string
    Insert(arg1 string, arg2 string)
    IsKilling() bool
    MapRowList(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 base.Base_queryMapForm, arg6 LnsAny) bool
    MapRowListSort(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 base.Base_queryMapForm, arg7 LnsAny) bool
    MapSymbolDecl(arg1 string, arg2 DBCtrl_callbackSymbolDecl)
    MapSymbolRef(arg1 string, arg2 DBCtrl_callbackSymbolRef)
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
// 38: decl @lns.@tags.@DBCtrl.DBCtrl.getMaxId
func DBCtrl_DBCtrl_getMaxId_1040_(access *DBAccess_DBAccess,tableName string,defId LnsInt) LnsInt {
    var id LnsAny
    id = nil
    access.FP.MapRowList(tableName, nil, 1, "MAX(id)", nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        id = items.Items["id"]
        
        return false
    }), base.Base_errHandleForm(DBCtrl_getMaxId___anonymous_1050_))
    if id != nil{
        id_125 := id
        return Lns_forceCastInt(id_125)
    }
    return defId
}

// 55: DeclConstr
func (self *DBCtrl_DBCtrl) InitDBCtrl_DBCtrl(access *DBAccess_DBAccess,readonly bool) {
    self.access = access
    
    self.individualTypeFlag = false
    
    self.individualStructFlag = false
    
    self.individualMacroFlag = false
    
    self.projDir = Depend_getCurDir()
    
    self.idMgrNamespace = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1040_(access, "namespace", DBCtrl_userNsId))
    
    self.idMgrSimpleName = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1040_(access, "simpleName", DBCtrl_userNsId))
    
    self.idMgrFilePath = NewDBCtrl_IdMgr(DBCtrl_DBCtrl_getMaxId_1040_(access, "filePath", DBCtrl_userNsId))
    
}

// 70: decl @lns.@tags.@DBCtrl.DBCtrl.close
func (self *DBCtrl_DBCtrl) Close() {
    self.access.FP.Close()
}

// 74: decl @lns.@tags.@DBCtrl.DBCtrl.begin
func (self *DBCtrl_DBCtrl) Begin() {
    self.access.FP.Begin()
}

// 77: decl @lns.@tags.@DBCtrl.DBCtrl.commit
func (self *DBCtrl_DBCtrl) Commit() {
    self.access.FP.Commit()
}

// 81: decl @lns.@tags.@DBCtrl.DBCtrl.exec
func (self *DBCtrl_DBCtrl) Exec(stmt string,errHandle LnsAny) {
    self.access.FP.Exec(stmt, errHandle)
}

// 85: decl @lns.@tags.@DBCtrl.DBCtrl.insert
func (self *DBCtrl_DBCtrl) Insert(tableName string,values string) {
    self.access.FP.Insert(tableName, values)
}

// 89: decl @lns.@tags.@DBCtrl.DBCtrl.update
func (self *DBCtrl_DBCtrl) Update(tableName string,set string,condition LnsAny) {
    var sql string
    sql = Lns_getVM().String_format("UPDATE %s SET %s", []LnsAny{tableName, set})
    if Lns_isCondTrue( condition){
        sql = Lns_getVM().String_format("%s WHERE %s", []LnsAny{sql, condition})
        
    }
    self.FP.Exec(sql, nil)
}

// 99: decl @lns.@tags.@DBCtrl.DBCtrl.mapRowList
func (self *DBCtrl_DBCtrl) MapRowList(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapRowList(tableName, condition, limit, attrib, nil, _func, errHandle)
}

// 107: decl @lns.@tags.@DBCtrl.DBCtrl.mapRowListSort
func (self *DBCtrl_DBCtrl) MapRowListSort(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,order LnsAny,_func base.Base_queryMapForm,errHandle LnsAny) bool {
    return self.access.FP.MapRowList(tableName, condition, limit, attrib, order, _func, errHandle)
}

// 115: decl @lns.@tags.@DBCtrl.DBCtrl.getRowList
func (self *DBCtrl_DBCtrl) getRowList(tableName string,condition LnsAny,limit LnsAny,attrib LnsAny,errHandle LnsAny) *LnsList {
    var rows *LnsList
    rows = NewLnsList([]LnsAny{})
    self.FP.MapRowList(tableName, condition, limit, attrib, base.Base_queryMapForm(func(items *LnsMap) bool {
        rows.Insert(items)
        return true
    }), errHandle)
    return rows
}

// 128: decl @lns.@tags.@DBCtrl.DBCtrl.getRow
func (self *DBCtrl_DBCtrl) GetRow(tableName string,condition LnsAny,attrib LnsAny,errHandle LnsAny) LnsAny {
    var row *LnsList
    row = self.FP.getRowList(tableName, condition, 1, attrib, errHandle)
    if row.Len() == 0{
        return nil
    }
    return row.GetAt(1).(*LnsMap)
}

// 138: decl @lns.@tags.@DBCtrl.DBCtrl.getEtc
func (self *DBCtrl_DBCtrl) GetEtc(key string) LnsAny {
    {
        _etc := self.FP.GetRow("etc", "keyName = '" + key + "'", nil, nil)
        if _etc != nil {
            etc := _etc.(*LnsMap)
            {
                _val := etc.Items["val"]
                if _val != nil {
                    val := _val
                    return val.(string)
                }
            }
        }
    }
    return nil
}

// 147: decl @lns.@tags.@DBCtrl.DBCtrl.setEtc
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

// 157: decl @lns.@tags.@DBCtrl.DBCtrl.equalsEtc
func (self *DBCtrl_DBCtrl) EqualsEtc(key string,val string) bool {
    if self.FP.GetEtc(key) == val{
        return true
    }
    return false
}

// 164: decl @lns.@tags.@DBCtrl.DBCtrl.isKilling
func (self *DBCtrl_DBCtrl) IsKilling() bool {
    if self.FP.EqualsEtc("killFlag", "1"){
        return true
    }
    return false
}

// 213: decl @lns.@tags.@DBCtrl.DBCtrl.creataTables
func (self *DBCtrl_DBCtrl) creataTables() {
    self.FP.Exec(Lns_getVM().String_format("BEGIN;\nCREATE TABLE etc ( keyName VARCHAR UNIQUE COLLATE binary PRIMARY KEY, val VARCHAR);\nINSERT INTO etc VALUES( 'version', '%d' );\nINSERT INTO etc VALUES( 'projDir', '' );\nINSERT INTO etc VALUES( 'individualStructFlag', '0' );\nINSERT INTO etc VALUES( 'individualTypeFlag', '0' );\nINSERT INTO etc VALUES( 'individualMacroFlag', '0' );\nINSERT INTO etc VALUES( 'killFlag', '0' );\nCREATE TABLE namespace ( id INTEGER PRIMARY KEY, snameId INTEGER, parentId INTEGER, digest CHAR(32), name VARCHAR UNIQUE COLLATE binary, otherName VARCHAR COLLATE binary, virtual INTEGER);\nINSERT INTO namespace VALUES( NULL, 1, 0, '', '', '', 0 );\n\nCREATE TABLE simpleName ( id INTEGER PRIMARY KEY, name VARCHAR UNIQUE COLLATE binary);\nCREATE TABLE filePath ( id INTEGER PRIMARY KEY, path VARCHAR UNIQUE COLLATE binary, incFlag INTEGER, digest CHAR(32), currentDir VARCHAR COLLATE binary, invalidSkip INTEGER);\nINSERT INTO filePath VALUES( NULL, '', 0, '', '', 1 );\n\nCREATE TABLE override (nsId INTEGER, superNsId INTEGER, PRIMARY KEY (nsId, superNsId));\nCREATE TABLE symbolDecl ( nsId INTEGER, snameId INTEGER, parentId INTEGER, type INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, comment VARCHAR COLLATE binary, hasBodyFlag INTEGER, PRIMARY KEY( nsId, fileId, line ) );\nINSERT INTO symbolDecl VALUES( 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, '', 0 );\n\nCREATE TABLE symbolRef ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );\nCREATE TABLE symbolSet ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );\n\nCREATE TABLE funcCall ( nsId INTEGER, snameId INTEGER, belongNsId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, PRIMARY KEY( nsId, belongNsId ) );\nCREATE TABLE incRef ( id INTEGER, baseFileId INTEGER, line INTEGER );\nCREATE TABLE incCache ( id INTEGER, baseFileId INTEGER, incFlag INTEGER, PRIMARY KEY( id, baseFileId ) );\nCREATE TABLE incBelong ( id INTEGER, baseFileId INTEGER, nsId INTEGER, PRIMARY KEY ( id, nsId ) );\nCREATE INDEX index_ns ON namespace ( id, snameId, parentId, name, otherName );\nCREATE INDEX index_sName ON simpleName ( id, name );\nCREATE INDEX index_filePath ON filePath ( id, path );\nCREATE INDEX index_override ON override (nsId, superNsId);\nCREATE INDEX index_symDecl ON symbolDecl ( nsId, parentId, snameId, fileId );\nCREATE INDEX index_symRef ON symbolRef ( nsId, snameId, fileId, belongNsId );\nCREATE INDEX index_incRef ON incRef ( id, baseFileId );\nCREATE INDEX index_incCache ON incCache ( id, baseFileId, incFlag );\nCREATE INDEX index_incBelong ON incBelong ( id, baseFileId );\nCOMMIT;\n", []LnsAny{(LnsInt)(DBCtrl_DB_VERSION)}), nil)
}

// 305: decl @lns.@tags.@DBCtrl.DBCtrl.addFile
func (self *DBCtrl_DBCtrl) AddFile(path string) LnsInt {
    if Lns_car(Lns_getVM().String_find(path,self.projDir, 1, true)) == 1{
        path = Lns_getVM().String_format("|%s", []LnsAny{Lns_getVM().String_sub(path,len(path) + 1, nil)})
        
    }
    var fileId LnsAny
    fileId = nil
    self.FP.MapRowList("filePath", "path = '%s'", 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _filePath := DBCtrl_convExp1362(Lns_2DDD(DBCtrl_ItemFilePath__fromStem_1193_(items,nil)))
            if _filePath != nil {
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
    var id LnsInt
    id = self.idMgrFilePath.FP.getIdNext()
    self.FP.Insert("filePath", Lns_getVM().String_format("%d,'%s',0,'','|',0", []LnsAny{id, path}))
    return id
}

// 329: decl @lns.@tags.@DBCtrl.DBCtrl.getName
func (self *DBCtrl_DBCtrl) getName(nsId LnsInt) LnsAny {
    var name LnsAny
    name = nil
    self.FP.MapRowList("namespace", Lns_getVM().String_format("id = %d", []LnsAny{nsId}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp1478(Lns_2DDD(DBCtrl_ItemNamespace__fromStem_1225_(items,nil)))
            if _obj != nil {
                obj := _obj.(*DBCtrl_ItemNamespace)
                name = obj.FP.Get_name()
                
            }
        }
        return false
    }), nil)
    return name
}

// 343: decl @lns.@tags.@DBCtrl.DBCtrl.getNsId
func (self *DBCtrl_DBCtrl) getNsId(name string) LnsAny {
    var nsId LnsAny
    nsId = nil
    self.FP.MapRowList("namespace", Lns_getVM().String_format("name = '%s'", []LnsAny{name}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp1559(Lns_2DDD(DBCtrl_ItemNamespace__fromStem_1225_(items,nil)))
            if _obj != nil {
                obj := _obj.(*DBCtrl_ItemNamespace)
                nsId = obj.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    return nsId
}

// 364: decl @lns.@tags.@DBCtrl.DBCtrl.addNamespace
func (self *DBCtrl_DBCtrl) AddNamespace(fullName string,parentId LnsInt)(LnsInt, bool) {
    var id LnsAny
    id = nil
    self.FP.MapRowList("namespace", Lns_getVM().String_format("name = '%s'", []LnsAny{fullName}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp1645(Lns_2DDD(DBCtrl_ItemNamespace__fromStem_1225_(items,nil)))
            if _obj != nil {
                obj := _obj.(*DBCtrl_ItemNamespace)
                id = obj.FP.Get_id()
                
            }
        }
        return false
    }), nil)
    if id != nil{
        id_385 := id.(LnsInt)
        return id_385, false
    }
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var newId LnsInt
    newId = self.idMgrNamespace.FP.getIdNext()
    self.FP.Insert("namespace", Lns_getVM().String_format("%d, %d, %d, '', '%s', '', 1", []LnsAny{newId, snid, parentId, fullName}))
    return newId, true
}

// 386: decl @lns.@tags.@DBCtrl.DBCtrl.getFilePath
func (self *DBCtrl_DBCtrl) GetFilePath(fileId LnsInt) LnsAny {
    var path LnsAny
    path = nil
    self.FP.MapRowList("filePath", Lns_getVM().String_format("id = %d", []LnsAny{fileId}), 1, nil, base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _obj := DBCtrl_convExp1773(Lns_2DDD(DBCtrl_ItemFilePath__fromStem_1193_(items,nil)))
            if _obj != nil {
                obj := _obj.(*DBCtrl_ItemFilePath)
                path = obj.FP.Get_path()
                
            }
        }
        return false
    }), nil)
    return path
}

// 399: decl @lns.@tags.@DBCtrl.DBCtrl.addOverride
func (self *DBCtrl_DBCtrl) AddOverride(nsId LnsInt,superNsId LnsInt) {
    self.FP.Insert("override", Lns_getVM().String_format("%d, %d", []LnsAny{nsId, superNsId}))
}

// 404: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolDecl
func (self *DBCtrl_DBCtrl) AddSymbolDecl(nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var kind LnsInt
    kind = 0
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var parentId LnsAny
    
    {
        _parentId := Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
        self.FP.GetRow("namespace", Lns_getVM().String_format("id = %d", []LnsAny{nsId}), "parentId", nil)) && 
        Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsMap).Items["parentId"]))
        if _parentId == nil{
            return 
        } else {
            parentId = _parentId
        }
    }
    self.FP.Insert("symbolDecl", Lns_getVM().String_format("%d, %d, %d, %d, %d, %d, %d, %d, %d, 0, '', 0", []LnsAny{nsId, snid, Lns_forceCastInt(parentId), kind, fileId, lineNo, column, lineNo, column}))
}

// 422: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolRef
func (self *DBCtrl_DBCtrl) AddSymbolRef(nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var belongNsId LnsInt
    belongNsId = DBCtrl_rootNsId
    self.FP.Insert("symbolRef", Lns_getVM().String_format("%d, %d, %d, %d, %d, %d, %d, 0, %d", []LnsAny{nsId, snid, fileId, lineNo, column, lineNo, column, belongNsId}))
}

// 432: decl @lns.@tags.@DBCtrl.DBCtrl.addSymbolSet
func (self *DBCtrl_DBCtrl) AddSymbolSet(nsId LnsInt,fileId LnsInt,lineNo LnsInt,column LnsInt) {
    var snid LnsInt
    snid = DBCtrl_rootNsId
    var belongNsId LnsInt
    belongNsId = DBCtrl_rootNsId
    self.FP.Insert("symbolSet", Lns_getVM().String_format("%d, %d, %d, %d, %d, %d", []LnsAny{nsId, snid, fileId, lineNo, column, belongNsId}))
}

// 475: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolDecl
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
            _item := DBCtrl_convExp2273(Lns_2DDD(DBCtrl_ItemOverride__fromStem(items,nil)))
            if _item != nil {
                item := _item.(*DBCtrl_ItemOverride)
                overrideStr = Lns_getVM().String_format("%s, %d", []LnsAny{overrideStr, item.FP.Get_nsId()})
                
            }
        }
        return true
    }), nil)
    self.FP.MapRowListSort("symbolDecl", Lns_getVM().String_format("nsId IN (%s)", []LnsAny{overrideStr}), nil, nil, "fileId, line", base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp2337(Lns_2DDD(DBCtrl_ItemSymbolDecl__fromStem(items,nil)))
            if _item != nil {
                item := _item.(*DBCtrl_ItemSymbolDecl)
                return callback(item)
            }
        }
        return true
    }), nil)
}

// 505: decl @lns.@tags.@DBCtrl.DBCtrl.mapSymbolRef
func (self *DBCtrl_DBCtrl) MapSymbolRef(name string,callback DBCtrl_callbackSymbolRef) {
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
            _item := DBCtrl_convExp2445(Lns_2DDD(DBCtrl_ItemOverride__fromStem(items,nil)))
            if _item != nil {
                item := _item.(*DBCtrl_ItemOverride)
                overrideStr = Lns_getVM().String_format("%s, %d", []LnsAny{overrideStr, item.FP.Get_superNsId()})
                
            }
        }
        return true
    }), nil)
    self.FP.MapRowListSort("symbolRef", Lns_getVM().String_format("nsId IN (%s)", []LnsAny{overrideStr}), nil, nil, "fileId, line", base.Base_queryMapForm(func(items *LnsMap) bool {
        {
            _item := DBCtrl_convExp2509(Lns_2DDD(DBCtrl_ItemSymbolRef__fromStem(items,nil)))
            if _item != nil {
                item := _item.(*DBCtrl_ItemSymbolRef)
                return callback(item)
            }
        }
        return true
    }), nil)
}

// 534: decl @lns.@tags.@DBCtrl.DBCtrl.dumpAll
func (self *DBCtrl_DBCtrl) DumpAll() {
    Lns_print([]LnsAny{"filePath"})
    self.FP.MapRowList("filePath", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1441_), nil)
    Lns_print([]LnsAny{"namespace"})
    self.FP.MapRowList("namespace", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1448_), nil)
    Lns_print([]LnsAny{"override"})
    self.FP.MapRowList("override", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1455_), nil)
    Lns_print([]LnsAny{"symbolDecl"})
    self.FP.MapRowList("symbolDecl", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1462_), nil)
    Lns_print([]LnsAny{"symbolRef"})
    self.FP.MapRowList("symbolRef", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1469_), nil)
    Lns_print([]LnsAny{"symbolSet"})
    self.FP.MapRowList("symbolSet", nil, nil, nil, base.Base_queryMapForm(DBCtrl_dumpAll___anonymous_1476_), nil)
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
func DBCtrl_ETC__fromMap_1148_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ETC_FromMap( arg1, paramList )
}
func DBCtrl_ETC__fromStem_1151_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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

// declaration Class -- ItemFilePath
type DBCtrl_ItemFilePathMtd interface {
    ToMap() *LnsMap
    Get_id() LnsInt
    Get_path() string
}
type DBCtrl_ItemFilePath struct {
    id LnsInt
    path string
    incFlag LnsInt
    digest string
    currentDir string
    invalidSkip LnsInt
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
func NewDBCtrl_ItemFilePath(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 string, arg5 string, arg6 LnsInt) *DBCtrl_ItemFilePath {
    obj := &DBCtrl_ItemFilePath{}
    obj.FP = obj
    obj.InitDBCtrl_ItemFilePath(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *DBCtrl_ItemFilePath) InitDBCtrl_ItemFilePath(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 string, arg5 string, arg6 LnsInt) {
    self.id = arg1
    self.path = arg2
    self.incFlag = arg3
    self.digest = arg4
    self.currentDir = arg5
    self.invalidSkip = arg6
}
func (self *DBCtrl_ItemFilePath) Get_id() LnsInt{ return self.id }
func (self *DBCtrl_ItemFilePath) Get_path() string{ return self.path }
func (self *DBCtrl_ItemFilePath) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["id"] = Lns_ToCollection( self.id )
    obj.Items["path"] = Lns_ToCollection( self.path )
    obj.Items["incFlag"] = Lns_ToCollection( self.incFlag )
    obj.Items["digest"] = Lns_ToCollection( self.digest )
    obj.Items["currentDir"] = Lns_ToCollection( self.currentDir )
    obj.Items["invalidSkip"] = Lns_ToCollection( self.invalidSkip )
    return obj
}
func (self *DBCtrl_ItemFilePath) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func DBCtrl_ItemFilePath__fromMap_1190_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemFilePath_FromMap( arg1, paramList )
}
func DBCtrl_ItemFilePath__fromStem_1193_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["incFlag"], false, nil); !ok {
       return false,nil,"incFlag:" + mess.(string)
    } else {
       newObj.incFlag = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["digest"], false, nil); !ok {
       return false,nil,"digest:" + mess.(string)
    } else {
       newObj.digest = conv.(string)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["currentDir"], false, nil); !ok {
       return false,nil,"currentDir:" + mess.(string)
    } else {
       newObj.currentDir = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["invalidSkip"], false, nil); !ok {
       return false,nil,"invalidSkip:" + mess.(string)
    } else {
       newObj.invalidSkip = conv.(LnsInt)
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
func DBCtrl_ItemNamespace__fromMap_1222_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return DBCtrl_ItemNamespace_FromMap( arg1, paramList )
}
func DBCtrl_ItemNamespace__fromStem_1225_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
}
type DBCtrl_ItemSymbolRef struct {
    nsId LnsInt
    fileId LnsInt
    line LnsInt
    column LnsInt
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
func NewDBCtrl_ItemSymbolRef(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) *DBCtrl_ItemSymbolRef {
    obj := &DBCtrl_ItemSymbolRef{}
    obj.FP = obj
    obj.InitDBCtrl_ItemSymbolRef(arg1, arg2, arg3, arg4)
    return obj
}
func (self *DBCtrl_ItemSymbolRef) InitDBCtrl_ItemSymbolRef(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) {
    self.nsId = arg1
    self.fileId = arg2
    self.line = arg3
    self.column = arg4
}
func (self *DBCtrl_ItemSymbolRef) Get_nsId() LnsInt{ return self.nsId }
func (self *DBCtrl_ItemSymbolRef) Get_fileId() LnsInt{ return self.fileId }
func (self *DBCtrl_ItemSymbolRef) Get_line() LnsInt{ return self.line }
func (self *DBCtrl_ItemSymbolRef) Get_column() LnsInt{ return self.column }
func (self *DBCtrl_ItemSymbolRef) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["nsId"] = Lns_ToCollection( self.nsId )
    obj.Items["fileId"] = Lns_ToCollection( self.fileId )
    obj.Items["line"] = Lns_ToCollection( self.line )
    obj.Items["column"] = Lns_ToCollection( self.column )
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
