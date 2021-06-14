--lns/tags/Inq.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Inq'
local _lune = {}
if _lune4 then
   _lune = _lune4
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

if not _lune4 then
   _lune4 = _lune
end
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Util = _lune.loadModule( 'lns.tags.Util' )
local LuneAst = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )

local function InqDef( db, pattern )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapSymbolDecl( pattern, function ( item )
   
      local path = db:getFilePath( item:get_fileId() )
      if  nil == path then
         local _path = path
      
         error( string.format( "file id is illegal -- %d", item:get_fileId()) )
      end
      
      Util.outputLocate( io.stdout, pattern, path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqDef = InqDef

local function InqRef( db, pattern, onlySet )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapSymbolRef( pattern, onlySet, function ( item )
   
      local path = db:getFilePath( item:get_fileId() )
      if  nil == path then
         local _path = path
      
         error( string.format( "file id is illegal -- %d", item:get_fileId()) )
      end
      
      Util.outputLocate( io.stdout, pattern, path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqRef = InqRef

local function InqAllmut( db )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapAllmutDecl( function ( item )
   
      local path = db:getFilePath( item:get_fileId() )
      if  nil == path then
         local _path = path
      
         error( string.format( "file id is illegal -- %d", item:get_fileId()) )
      end
      
      Util.outputLocate( io.stdout, "allmut", path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqAllmut = InqAllmut

local function InqAsync( db, asyncMode )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapAsyncMode( asyncMode, function ( item )
   
      local path = db:getFilePath( item:get_fileId() )
      if  nil == path then
         local _path = path
      
         error( string.format( "file id is illegal -- %d", item:get_fileId()) )
      end
      
      Util.outputLocate( io.stdout, string.format( "%d", asyncMode), path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqAsync = InqAsync

local function InqLuaval( db )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapLuavalRef( function ( item )
   
      local path = db:getFilePath( item:get_fileId() )
      if  nil == path then
         local _path = path
      
         error( string.format( "file id is illegal -- %d", item:get_fileId()) )
      end
      
      Util.outputLocate( io.stdout, "luaval", path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqLuaval = InqLuaval

local function InqAsyncLock( db )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapAsyncLock( function ( item )
   
      local path = db:getFilePath( item:get_fileId() )
      if  nil == path then
         local _path = path
      
         error( string.format( "file id is illegal -- %d", item:get_fileId()) )
      end
      
      Util.outputLocate( io.stdout, "luaval", path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqAsyncLock = InqAsyncLock

return _moduleObj
