--lns/tags/DBCtrl.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@DBCtrl'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

local DBCtrl = {}
_moduleObj.DBCtrl = DBCtrl
function DBCtrl.new( access )
   local obj = {}
   DBCtrl.setmeta( obj )
   if obj.__init then obj:__init( access ); end
   return obj
end
function DBCtrl:__init(access) 
   self.access = access
end
function DBCtrl:begin(  )

   self.access:begin(  )
end
function DBCtrl.setmeta( obj )
  setmetatable( obj, { __index = DBCtrl  } )
end


local function open( path, readonly )

   local db = DBAccess.open( path, readonly )
   if  nil == db then
      local _db = db
   
      return nil
   end
   
   
   local dbCtrl = DBCtrl.new(db)
   
   if not readonly then
      dbCtrl:begin(  )
   end
   
   
   return dbCtrl
end
_moduleObj.open = open

function DBCtrl:test(  )

   local stmt = [==[
      create table foo (id integer not null primary key, name text);
   delete from foo;
]==]
   self.access:exec( stmt, nil )
   
   for index = 0, 10 do
      local sql = string.format( "insert into foo(id, name) values(%d, 'テスト%03d')", index, index)
      self.access:exec( sql, nil )
   end
   
   self.access:commit(  )
   
   self.access:mapRowList( "foo", nil, nil, "id, name", function ( row )
   
      print( math.floor(row[1]) + 10, row[2] .. "hoge" )
      return true
   end )
   
   self.access:mapRowList( "foo", "id = 3", nil, "name", function ( row )
   
      print( row[1] )
      return false
   end )
   
   self.access:exec( "delete from foo", nil )
   
   self.access:exec( "insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')", nil )
   
   self.access:mapRowList( "foo", nil, nil, "id, name", function ( row )
   
      print( row[1], row[2] )
      return true
   end )
   
   self.access:close(  )
end


local function test(  )

   local db = open( "hoge.sqlite3", false )
   if  nil == db then
      local _db = db
   
      print( "open error" )
      return false
   end
   
   
   db:test(  )
   
   return true
end
_moduleObj.test = test

return _moduleObj
