--lns/tags/DBCtrl.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@DBCtrl'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
end

function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
end
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune3 then
   _lune3 = _lune
end
local base = _lune.loadModule( 'go/github:com.ifritJP.lnssqlite3.src.lns.sqlite3.base' )
local DBAccess = _lune.loadModule( 'lns.tags.DBAccess' )
local Log = _lune.loadModule( 'lns.tags.Log' )
local Depend = _lune.loadModule( 'lns.tags.Depend' )

local rootNsId = 1
_moduleObj.rootNsId = rootNsId

local userNsId = 2
local systemFileId = 1
local DB_VERSION = 9.0

local IdMgr = {}
function IdMgr.new( idNum )
   local obj = {}
   IdMgr.setmeta( obj )
   if obj.__init then obj:__init( idNum ); end
   return obj
end
function IdMgr:__init(idNum) 
   self.idNum = idNum
end
function IdMgr:getIdNext(  )

   local idNum = self.idNum
   self.idNum = self.idNum + 1
   return idNum
end
function IdMgr.setmeta( obj )
  setmetatable( obj, { __index = IdMgr  } )
end


local DBCtrl = {}
_moduleObj.DBCtrl = DBCtrl
function DBCtrl.getMaxId( access, tableName, defId )

   local id = nil
   access:mapRowList( tableName, nil, 1, "MAX(id)", nil, function ( items )
   
      id = items['id']
      return false
   end, function ( stmt, msg )
   
   end )
   if id ~= nil then
      return math.floor(id)
   end
   
   return defId
end
function DBCtrl.new( access, readonly )
   local obj = {}
   DBCtrl.setmeta( obj )
   if obj.__init then obj:__init( access, readonly ); end
   return obj
end
function DBCtrl:__init(access, readonly) 
   self.access = access
   self.individualTypeFlag = false
   self.individualStructFlag = false
   self.individualMacroFlag = false
   self.projDir = Depend.getCurDir(  )
   
   self.idMgrNamespace = IdMgr.new(DBCtrl.getMaxId( access, "namespace", userNsId ))
   self.idMgrSimpleName = IdMgr.new(DBCtrl.getMaxId( access, "simpleName", userNsId ))
   self.idMgrFilePath = IdMgr.new(DBCtrl.getMaxId( access, "filePath", userNsId ))
end
function DBCtrl:close(  )

   self.access:close(  )
end
function DBCtrl:begin(  )

   self.access:begin(  )
end
function DBCtrl:commit(  )

   self.access:commit(  )
end
function DBCtrl:exec( stmt, errHandle )

   self.access:exec( stmt, errHandle )
end
function DBCtrl:insert( tableName, values )

   self.access:insert( tableName, values )
end
function DBCtrl:update( tableName, set, condition )

   local sql = string.format( "UPDATE %s SET %s", tableName, set )
   if condition then
      sql = string.format( "%s WHERE %s", sql, condition )
   end
   
   self:exec( sql )
end
function DBCtrl:mapRowList( tableName, condition, limit, attrib, func, errHandle )

   return self.access:mapRowList( tableName, condition, limit, attrib, nil, func, errHandle )
end
function DBCtrl:mapRowListSort( tableName, condition, limit, attrib, order, func, errHandle )

   return self.access:mapRowList( tableName, condition, limit, attrib, order, func, errHandle )
end
function DBCtrl:getRowList( tableName, condition, limit, attrib, errHandle )

   local rows = {}
   self:mapRowList( tableName, condition, limit, attrib, function ( items )
   
      table.insert( rows, items )
      return true
   end, errHandle )
   return rows
end
function DBCtrl:getRow( tableName, condition, attrib, errHandle )

   local row = self:getRowList( tableName, condition, 1, attrib, errHandle )
   if #row == 0 then
      return nil
   end
   
   
   return row[1]
end
function DBCtrl:getEtc( key )

   do
      local etc = self:getRow( "etc", "keyName = '" .. key .. "'" )
      if etc ~= nil then
         do
            local val = etc['val']
            if val ~= nil then
               return val
            end
         end
         
      end
   end
   
   return nil
end
function DBCtrl:setEtc( key, val )

   local keyTxt = string.format( "keyName = '%s'", key)
   local valTxt = string.format( "val = '%s'", val)
   if not self:getEtc( key ) then
      self:insert( "etc", string.format( "'%s', '%s'", key, val) )
   else
    
      self:update( "etc", valTxt, keyTxt )
   end
   
end
function DBCtrl:equalsEtc( key, val )

   if self:getEtc( key ) == val then
      return true
   end
   
   return false
end
function DBCtrl:isKilling(  )

   if self:equalsEtc( "killFlag", "1" ) then
      return true
   end
   
   return false
end
function DBCtrl.setmeta( obj )
  setmetatable( obj, { __index = DBCtrl  } )
end
function DBCtrl:get_individualTypeFlag()
   return self.individualTypeFlag
end
function DBCtrl:get_individualStructFlag()
   return self.individualStructFlag
end
function DBCtrl:get_individualMacroFlag()
   return self.individualMacroFlag
end
function DBCtrl:get_projDir()
   return self.projDir
end


local ETC = {}
setmetatable( ETC, { ifList = {Mapping,} } )
function ETC.setmeta( obj )
  setmetatable( obj, { __index = ETC  } )
end
function ETC.new( keyName, val )
   local obj = {}
   ETC.setmeta( obj )
   if obj.__init then
      obj:__init( keyName, val )
   end
   return obj
end
function ETC:__init( keyName, val )

   self.keyName = keyName
   self.val = val
end
function ETC:get_keyName()
   return self.keyName
end
function ETC:get_val()
   return self.val
end
function ETC:_toMap()
  return self
end
function ETC._fromMap( val )
  local obj, mes = ETC._fromMapSub( {}, val )
  if obj then
     ETC.setmeta( obj )
  end
  return obj, mes
end
function ETC._fromStem( val )
  return ETC._fromMap( val )
end

function ETC._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "keyName", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "val", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local function open( path, readonly )
   local __func__ = '@lns.@tags.@DBCtrl.open'

   Log.log( Log.Level.Log, __func__, 179, function (  )
   
      return "open"
   end )
   
   local db = DBAccess.open( path, readonly )
   if  nil == db then
      local _db = db
   
      return nil
   end
   
   
   local dbCtrl = DBCtrl.new(db, readonly)
   
   if not readonly then
      dbCtrl:begin(  )
   end
   
   
   local item = ETC._fromStem( dbCtrl:getRow( "etc", "keyName = 'version'" ) )
   if  nil == item then
      local _item = item
   
      Log.log( Log.Level.Err, __func__, 191, function (  )
      
         return "unknown version"
      end )
      
      db:close(  )
      return nil
   end
   
   if tonumber( item:get_val() ) ~= DB_VERSION then
      Log.log( Log.Level.Err, __func__, 196, function (  )
      
         return string.format( "not support version. -- %s", item:get_val())
      end )
      
      db:close(  )
      return nil
   end
   
   
   if dbCtrl:isKilling(  ) then
      error( "db is killed now" )
   end
   
   
   dbCtrl.individualTypeFlag = dbCtrl:equalsEtc( "individualTypeFlag", "1" )
   dbCtrl.individualStructFlag = dbCtrl:equalsEtc( "individualStructFlag", "1" )
   dbCtrl.individualMacroFlag = dbCtrl:equalsEtc( "individualMacroFlag", "1" )
   
   return dbCtrl
end
_moduleObj.open = open

function DBCtrl:creataTables(  )

   self:exec( string.format( [==[
BEGIN;
CREATE TABLE etc ( keyName VARCHAR UNIQUE COLLATE binary PRIMARY KEY, val VARCHAR);
INSERT INTO etc VALUES( 'version', '%d' );
INSERT INTO etc VALUES( 'projDir', '' );
INSERT INTO etc VALUES( 'individualStructFlag', '0' );
INSERT INTO etc VALUES( 'individualTypeFlag', '0' );
INSERT INTO etc VALUES( 'individualMacroFlag', '0' );
INSERT INTO etc VALUES( 'killFlag', '0' );
CREATE TABLE namespace ( id INTEGER PRIMARY KEY, snameId INTEGER, parentId INTEGER, digest CHAR(32), name VARCHAR UNIQUE COLLATE binary, otherName VARCHAR COLLATE binary, virtual INTEGER);
INSERT INTO namespace VALUES( NULL, 1, 0, '', '', '', 0 );

CREATE TABLE simpleName ( id INTEGER PRIMARY KEY, name VARCHAR UNIQUE COLLATE binary);
CREATE TABLE filePath ( id INTEGER PRIMARY KEY, path VARCHAR UNIQUE COLLATE binary, incFlag INTEGER, digest CHAR(32), currentDir VARCHAR COLLATE binary, invalidSkip INTEGER);
INSERT INTO filePath VALUES( NULL, '', 0, '', '', 1 );

CREATE TABLE override (nsId INTEGER, superNsId INTEGER, PRIMARY KEY (nsId, superNsId));
CREATE TABLE symbolDecl ( nsId INTEGER, snameId INTEGER, parentId INTEGER, type INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, comment VARCHAR COLLATE binary, hasBodyFlag INTEGER, PRIMARY KEY( nsId, fileId, line ) );
INSERT INTO symbolDecl VALUES( 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, '', 0 );

CREATE TABLE symbolRef ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );
CREATE TABLE symbolSet ( nsId INTEGER, snameId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, belongNsId INTEGER, PRIMARY KEY( nsId, fileId, line, column ) );

CREATE TABLE funcCall ( nsId INTEGER, snameId INTEGER, belongNsId INTEGER, fileId INTEGER, line INTEGER, column INTEGER, endLine INTEGER, endColumn INTEGER, charSize INTEGER, PRIMARY KEY( nsId, belongNsId ) );
CREATE TABLE incRef ( id INTEGER, baseFileId INTEGER, line INTEGER );
CREATE TABLE incCache ( id INTEGER, baseFileId INTEGER, incFlag INTEGER, PRIMARY KEY( id, baseFileId ) );
CREATE TABLE incBelong ( id INTEGER, baseFileId INTEGER, nsId INTEGER, PRIMARY KEY ( id, nsId ) );
CREATE INDEX index_ns ON namespace ( id, snameId, parentId, name, otherName );
CREATE INDEX index_sName ON simpleName ( id, name );
CREATE INDEX index_filePath ON filePath ( id, path );
CREATE INDEX index_override ON override (nsId, superNsId);
CREATE INDEX index_symDecl ON symbolDecl ( nsId, parentId, snameId, fileId );
CREATE INDEX index_symRef ON symbolRef ( nsId, snameId, fileId, belongNsId );
CREATE INDEX index_incRef ON incRef ( id, baseFileId );
CREATE INDEX index_incCache ON incCache ( id, baseFileId, incFlag );
CREATE INDEX index_incBelong ON incBelong ( id, baseFileId );
COMMIT;
]==], math.floor(DB_VERSION)) )
end


local ItemFilePath = {}
setmetatable( ItemFilePath, { ifList = {Mapping,} } )
function ItemFilePath.setmeta( obj )
  setmetatable( obj, { __index = ItemFilePath  } )
end
function ItemFilePath.new( id, path, incFlag, digest, currentDir, invalidSkip )
   local obj = {}
   ItemFilePath.setmeta( obj )
   if obj.__init then
      obj:__init( id, path, incFlag, digest, currentDir, invalidSkip )
   end
   return obj
end
function ItemFilePath:__init( id, path, incFlag, digest, currentDir, invalidSkip )

   self.id = id
   self.path = path
   self.incFlag = incFlag
   self.digest = digest
   self.currentDir = currentDir
   self.invalidSkip = invalidSkip
end
function ItemFilePath:get_id()
   return self.id
end
function ItemFilePath:get_path()
   return self.path
end
function ItemFilePath:_toMap()
  return self
end
function ItemFilePath._fromMap( val )
  local obj, mes = ItemFilePath._fromMapSub( {}, val )
  if obj then
     ItemFilePath.setmeta( obj )
  end
  return obj, mes
end
function ItemFilePath._fromStem( val )
  return ItemFilePath._fromMap( val )
end

function ItemFilePath._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "id", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "path", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "incFlag", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "digest", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "currentDir", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "invalidSkip", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local ItemNamespace = {}
setmetatable( ItemNamespace, { ifList = {Mapping,} } )
function ItemNamespace.setmeta( obj )
  setmetatable( obj, { __index = ItemNamespace  } )
end
function ItemNamespace.new( id, snameId, parentId, name )
   local obj = {}
   ItemNamespace.setmeta( obj )
   if obj.__init then
      obj:__init( id, snameId, parentId, name )
   end
   return obj
end
function ItemNamespace:__init( id, snameId, parentId, name )

   self.id = id
   self.snameId = snameId
   self.parentId = parentId
   self.name = name
end
function ItemNamespace:get_id()
   return self.id
end
function ItemNamespace:get_snameId()
   return self.snameId
end
function ItemNamespace:get_parentId()
   return self.parentId
end
function ItemNamespace:get_name()
   return self.name
end
function ItemNamespace:_toMap()
  return self
end
function ItemNamespace._fromMap( val )
  local obj, mes = ItemNamespace._fromMapSub( {}, val )
  if obj then
     ItemNamespace.setmeta( obj )
  end
  return obj, mes
end
function ItemNamespace._fromStem( val )
  return ItemNamespace._fromMap( val )
end

function ItemNamespace._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "id", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "snameId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "name", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local ItemSymbolDecl = {}
setmetatable( ItemSymbolDecl, { ifList = {Mapping,} } )
_moduleObj.ItemSymbolDecl = ItemSymbolDecl
function ItemSymbolDecl.setmeta( obj )
  setmetatable( obj, { __index = ItemSymbolDecl  } )
end
function ItemSymbolDecl.new( nsId, fileId, line, column )
   local obj = {}
   ItemSymbolDecl.setmeta( obj )
   if obj.__init then
      obj:__init( nsId, fileId, line, column )
   end
   return obj
end
function ItemSymbolDecl:__init( nsId, fileId, line, column )

   self.nsId = nsId
   self.fileId = fileId
   self.line = line
   self.column = column
end
function ItemSymbolDecl:get_nsId()
   return self.nsId
end
function ItemSymbolDecl:get_fileId()
   return self.fileId
end
function ItemSymbolDecl:get_line()
   return self.line
end
function ItemSymbolDecl:get_column()
   return self.column
end
function ItemSymbolDecl:_toMap()
  return self
end
function ItemSymbolDecl._fromMap( val )
  local obj, mes = ItemSymbolDecl._fromMapSub( {}, val )
  if obj then
     ItemSymbolDecl.setmeta( obj )
  end
  return obj, mes
end
function ItemSymbolDecl._fromStem( val )
  return ItemSymbolDecl._fromMap( val )
end

function ItemSymbolDecl._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "nsId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "fileId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "line", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "column", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local ItemSymbolRef = {}
setmetatable( ItemSymbolRef, { ifList = {Mapping,} } )
_moduleObj.ItemSymbolRef = ItemSymbolRef
function ItemSymbolRef.setmeta( obj )
  setmetatable( obj, { __index = ItemSymbolRef  } )
end
function ItemSymbolRef.new( nsId, fileId, line, column )
   local obj = {}
   ItemSymbolRef.setmeta( obj )
   if obj.__init then
      obj:__init( nsId, fileId, line, column )
   end
   return obj
end
function ItemSymbolRef:__init( nsId, fileId, line, column )

   self.nsId = nsId
   self.fileId = fileId
   self.line = line
   self.column = column
end
function ItemSymbolRef:get_nsId()
   return self.nsId
end
function ItemSymbolRef:get_fileId()
   return self.fileId
end
function ItemSymbolRef:get_line()
   return self.line
end
function ItemSymbolRef:get_column()
   return self.column
end
function ItemSymbolRef:_toMap()
  return self
end
function ItemSymbolRef._fromMap( val )
  local obj, mes = ItemSymbolRef._fromMapSub( {}, val )
  if obj then
     ItemSymbolRef.setmeta( obj )
  end
  return obj, mes
end
function ItemSymbolRef._fromStem( val )
  return ItemSymbolRef._fromMap( val )
end

function ItemSymbolRef._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "nsId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "fileId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "line", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "column", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local ItemOverride = {}
setmetatable( ItemOverride, { ifList = {Mapping,} } )
_moduleObj.ItemOverride = ItemOverride
function ItemOverride.setmeta( obj )
  setmetatable( obj, { __index = ItemOverride  } )
end
function ItemOverride.new( nsId, superNsId )
   local obj = {}
   ItemOverride.setmeta( obj )
   if obj.__init then
      obj:__init( nsId, superNsId )
   end
   return obj
end
function ItemOverride:__init( nsId, superNsId )

   self.nsId = nsId
   self.superNsId = superNsId
end
function ItemOverride:get_nsId()
   return self.nsId
end
function ItemOverride:get_superNsId()
   return self.superNsId
end
function ItemOverride:_toMap()
  return self
end
function ItemOverride._fromMap( val )
  local obj, mes = ItemOverride._fromMapSub( {}, val )
  if obj then
     ItemOverride.setmeta( obj )
  end
  return obj, mes
end
function ItemOverride._fromStem( val )
  return ItemOverride._fromMap( val )
end

function ItemOverride._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "nsId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "superNsId", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


function DBCtrl:addFile( path )

   if path:find( self.projDir, 1, true ) == 1 then
      path = string.format( "|%s", path:sub( #path + 1 ))
   end
   
   
   local fileId = nil
   self:mapRowList( "filePath", "path = '%s'", 1, nil, function ( items )
   
      do
         local filePath = ItemFilePath._fromStem( items )
         if filePath ~= nil then
            fileId = filePath:get_id()
         end
      end
      
      return false
   end )
   if fileId ~= nil then
      return fileId
   end
   
   local id = self.idMgrFilePath:getIdNext(  )
   self:insert( "filePath", string.format( "%d,'%s',0,'','|',0", id, path) )
   
   return id
end


function DBCtrl:getName( nsId )

   local name = nil
   self:mapRowList( "namespace", string.format( "id = %d", nsId), 1, nil, function ( items )
   
      do
         local obj = ItemNamespace._fromStem( items )
         if obj ~= nil then
            name = obj:get_name()
         end
      end
      
      return false
   end )
   return name
end


function DBCtrl:getNsId( name )

   local nsId = nil
   self:mapRowList( "namespace", string.format( "name = '%s'", name), 1, nil, function ( items )
   
      do
         local obj = ItemNamespace._fromStem( items )
         if obj ~= nil then
            nsId = obj:get_id()
         end
      end
      
      return false
   end )
   return nsId
end


function DBCtrl:addNamespace( fullName, parentId )

   local id = nil
   self:mapRowList( "namespace", string.format( "name = '%s'", fullName), 1, nil, function ( items )
   
      do
         local obj = ItemNamespace._fromStem( items )
         if obj ~= nil then
            id = obj:get_id()
         end
      end
      
      return false
   end )
   if id ~= nil then
      return id, false
   end
   
   local snid = _moduleObj.rootNsId
   local newId = self.idMgrNamespace:getIdNext(  )
   self:insert( "namespace", string.format( "%d, %d, %d, '', '%s', '', 1", newId, snid, parentId, fullName) )
   
   return newId, true
end


function DBCtrl:getFilePath( fileId )

   local path = nil
   self:mapRowList( "filePath", string.format( "id = %d", fileId), 1, nil, function ( items )
   
      do
         local obj = ItemFilePath._fromStem( items )
         if obj ~= nil then
            path = obj:get_path()
         end
      end
      
      return false
   end )
   return path
end


function DBCtrl:addOverride( nsId, superNsId )

   self:insert( "override", string.format( "%d, %d", nsId, superNsId) )
end


function DBCtrl:addSymbolDecl( nsId, fileId, lineNo, column )

   local kind = 0
   local snid = _moduleObj.rootNsId
   
   local parentId = _lune.nilacc( self:getRow( "namespace", string.format( "id = %d", nsId), "parentId" ), nil, 'item', "parentId")
   if  nil == parentId then
      local _parentId = parentId
   
      return 
   end
   
   
   self:insert( "symbolDecl", string.format( "%d, %d, %d, %d, %d, %d, %d, %d, %d, 0, '', 0", nsId, snid, math.floor(parentId), kind, fileId, lineNo, column, lineNo, column) )
end


function DBCtrl:addSymbolRef( nsId, fileId, lineNo, column )

   local snid = _moduleObj.rootNsId
   local belongNsId = _moduleObj.rootNsId
   self:insert( "symbolRef", string.format( "%d, %d, %d, %d, %d, %d, %d, 0, %d", nsId, snid, fileId, lineNo, column, lineNo, column, belongNsId) )
end


function DBCtrl:addSymbolSet( nsId, fileId, lineNo, column )

   local snid = _moduleObj.rootNsId
   local belongNsId = _moduleObj.rootNsId
   self:insert( "symbolSet", string.format( "%d, %d, %d, %d, %d, %d", nsId, snid, fileId, lineNo, column, belongNsId) )
end


local function create( dbPath )
   local __func__ = '@lns.@tags.@DBCtrl.create'

   Log.log( Log.Level.Log, __func__, 446, function (  )
   
      return "create"
   end )
   
   local db = DBAccess.open( dbPath, false )
   if  nil == db then
      local _db = db
   
      return nil
   end
   
   
   local dbCtrl = DBCtrl.new(db, false)
   
   dbCtrl:creataTables(  )
   
   dbCtrl:begin(  )
   
   dbCtrl:setEtc( "individualTypeFlag", "1" )
   dbCtrl:setEtc( "individualStructFlag", "1" )
   dbCtrl:setEtc( "individualMacroFlag", "1" )
   
   return dbCtrl
end

local function initDB( dbPath )

   os.remove( dbPath )
   local db = create( dbPath )
   if  nil == db then
      local _db = db
   
      print( "create error" )
      return 
   end
   
   db:commit(  )
   db:close(  )
end
_moduleObj.initDB = initDB


function DBCtrl:mapSymbolDecl( name, callback )

   local nsId = self:getNsId( name )
   if  nil == nsId then
      local _nsId = nsId
   
      return 
   end
   
   local overrideStr = string.format( "%d", nsId)
   self:mapRowList( "override", string.format( "superNsId = %d", nsId), nil, nil, function ( items )
   
      do
         local item = ItemOverride._fromStem( items )
         if item ~= nil then
            overrideStr = string.format( "%s, %d", overrideStr, item:get_nsId())
         end
      end
      
      return true
   end )
   
   self:mapRowListSort( "symbolDecl", string.format( "nsId IN (%s)", overrideStr), nil, nil, "fileId, line", function ( items )
   
      do
         local item = ItemSymbolDecl._fromStem( items )
         if item ~= nil then
            return callback( item )
         end
      end
      
      return true
   end )
end



function DBCtrl:mapSymbolRef( name, callback )

   local nsId = self:getNsId( name )
   if  nil == nsId then
      local _nsId = nsId
   
      return 
   end
   
   local overrideStr = string.format( "%d", nsId)
   self:mapRowList( "override", string.format( "nsId = %d", nsId), nil, nil, function ( items )
   
      do
         local item = ItemOverride._fromStem( items )
         if item ~= nil then
            overrideStr = string.format( "%s, %d", overrideStr, item:get_superNsId())
         end
      end
      
      return true
   end )
   
   self:mapRowListSort( "symbolRef", string.format( "nsId IN (%s)", overrideStr), nil, nil, "fileId, line", function ( items )
   
      do
         local item = ItemSymbolRef._fromStem( items )
         if item ~= nil then
            return callback( item )
         end
      end
      
      return true
   end )
end


function DBCtrl:dumpAll(  )

   print( "filePath" )
   self:mapRowList( "filePath", nil, nil, nil, function ( items )
   
      print( items['id'], items['path'] )
      return true
   end )
   
   print( "namespace" )
   self:mapRowList( "namespace", nil, nil, nil, function ( items )
   
      print( items['id'], items['name'] )
      return true
   end )
   
   print( "override" )
   self:mapRowList( "override", nil, nil, nil, function ( items )
   
      print( items['nsId'], items['superNsId'] )
      return true
   end )
   
   print( "symbolDecl" )
   self:mapRowList( "symbolDecl", nil, nil, nil, function ( items )
   
      print( items['nsId'], items['fileId'], items['line'], items['column'] )
      return true
   end )
   
   print( "symbolRef" )
   self:mapRowList( "symbolRef", nil, nil, nil, function ( items )
   
      print( items['nsId'], items['fileId'], items['line'], items['column'] )
      return true
   end )
   
   print( "symbolSet" )
   self:mapRowList( "symbolSet", nil, nil, nil, function ( items )
   
      print( items['nsId'], items['fileId'], items['line'], items['column'] )
      return true
   end )
end


local function test(  )

   local dbPath = "lnstags.sqlite3"
   do
      initDB( dbPath )
   end
   
   
   local db = open( dbPath, false )
   if  nil == db then
      local _db = db
   
      print( "open error" )
      return false
   end
   
   
   local fileId = _moduleObj.rootNsId
   for __index, path in pairs( {"aa.lns", "bb.lns", "cc.lns"} ) do
      fileId = db:addFile( path )
   end
   
   
   local parentId = _moduleObj.rootNsId
   for index, name in pairs( {"@hoge", "@hoge.@foo", "@hoge.@foo.bar"} ) do
      local newid = db:addNamespace( name, parentId )
      db:addSymbolDecl( newid, fileId, 100 + index, index * 10 )
      db:addSymbolRef( newid, fileId, 200 + index, index * 20 )
      db:addSymbolSet( newid, fileId, 300 + index, index * 30 )
      
      parentId = newid
   end
   
   
   do
      local _
      local _539, added = db:addNamespace( "@hoge", _moduleObj.rootNsId )
      print( "added", added )
   end
   
   
   db:commit(  )
   
   db:dumpAll(  )
   
   return true
end
_moduleObj.test = test

return _moduleObj
