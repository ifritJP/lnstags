--lns/tags/DBAccess.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@DBAccess'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
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
local Log = _lune.loadModule( 'lns.tags.Log' )

local DBAccess = {}
_moduleObj.DBAccess = DBAccess
function DBAccess.new( db, path, readonlyFlag )
   local obj = {}
   DBAccess.setmeta( obj )
   if obj.__init then obj:__init( db, path, readonlyFlag ); end
   return obj
end
function DBAccess:__init(db, path, readonlyFlag) 
   self.db = db
   self.path = path
   self.beginFlag = false
   self.readonlyFlag = readonlyFlag
end
function DBAccess.setmeta( obj )
  setmetatable( obj, { __index = DBAccess  } )
end
function DBAccess:get_readonlyFlag()
   return self.readonlyFlag
end
function DBAccess:get_beginFlag()
   return self.beginFlag
end


local function open( path, readonly )

   local db, err = base.Open( path, readonly, false )
   if  nil == db or  nil == err then
      local _db = db
      local _err = err
   
      print( 1, string.format( "DBAccess:open error -- %s", err) )
      return nil
   end
   
   return DBAccess.new(db, path, readonly)
end
_moduleObj.open = open

function DBAccess:close(  )

   self.db:Close(  )
end


function DBAccess:outputLog( message )

   
end


function DBAccess:begin(  )
   local __func__ = '@lns.@tags.@DBAccess.DBAccess.begin'

   Log.log( Log.Level.Log, __func__, 45, function (  )
   
      return "start"
   end )
   
   
   if self.readonlyFlag then
      Log.log( Log.Level.Err, __func__, 48, function (  )
      
         return "db mode is read only"
      end )
      
      os.exit( 1 )
   end
   
   
   self.beginFlag = true
   self.db:Begin(  )
   
end


function DBAccess:commit(  )
   local __func__ = '@lns.@tags.@DBAccess.DBAccess.commit'

   if self.readonlyFlag then
      return 
   end
   
   if not self.beginFlag then
      return 
   end
   
   self.beginFlag = false
   
   Log.log( Log.Level.Log, __func__, 73, function (  )
   
      return "commit: start"
   end )
   
   
   self.db:Commit(  )
   
   Log.log( Log.Level.Log, __func__, 77, function (  )
   
      return "commit: end"
   end )
   
end


function DBAccess:exec( stmt, errHandle )

   self.db:Exec( stmt, errHandle )
end


function DBAccess:mapJoin( tableName, otherTable, on, condition, limit, attrib, func )

   local query = string.format( "SELECT %s FROM %s INNER JOIN %s ON %s", attrib or "*", tableName, otherTable, on)
   if condition ~= nil then
      query = string.format( "%s WHERE %s", query, condition)
   end
   
   if limit ~= nil then
      query = string.format( "%s LIMIT %d", query, limit)
   end
   
   return self.db:MapQuery( query, func )
end


function DBAccess:mapJoin2( tableName, otherTable, on, otherTable2, on2, condition, limit, attrib, func )

   local query = string.format( "SELECT %s FROM %s INNER JOIN %s ON %s INNER JOIN %s ON %s", attrib or "*", tableName, otherTable, on, otherTable2, on2)
   if condition ~= nil then
      query = string.format( "%s WHERE %s", query, condition)
   end
   
   if limit ~= nil then
      query = string.format( "%s LIMIT %d", query, limit)
   end
   
   return self.db:MapQuery( query, func )
end


function DBAccess:mapJoin3( tableName, otherTable, on, otherTable2, on2, otherTable3, on3, condition, limit, attrib, func )

   local query = string.format( "SELECT %s FROM %s INNER JOIN %s ON %s INNER JOIN %s ON %s INNER JOIN %s ON %s", attrib or "*", tableName, otherTable, on, otherTable2, on2, otherTable3, on3)
   if condition ~= nil then
      query = string.format( "%s WHERE %s", query, condition)
   end
   
   if limit ~= nil then
      query = string.format( "%s LIMIT %d", query, limit)
   end
   
   return self.db:MapQuery( query, func )
end


function DBAccess:mapRowList( tableName, condition, limit, attrib, func )

   local query
   
   local ATTRIB = _lune.unwrapDefault( attrib, "*")
   if condition ~= nil then
      query = string.format( "SELECT %s FROM %s WHERE %s", ATTRIB, tableName, condition)
   else
      query = string.format( "SELECT %s FROM %s", ATTRIB, tableName)
   end
   
   if limit ~= nil then
      query = string.format( "%s LIMIT %d", query, limit)
   end
   
   return self.db:MapQuery( query, func )
end


return _moduleObj
