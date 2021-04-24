--lns/tags/Inq.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Inq'
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
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Util = _lune.loadModule( 'lns.tags.Util' )

local function InqDef( db, pattern )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapSymbolDecl( pattern, function ( item )
   
      local path = _lune.unwrap( db:getFilePath( item:get_fileId() ))
      Util.outputLocate( io.stdout, pattern, path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqDef = InqDef

local function InqRef( db, pattern )

   local factory = Util.SourceCodeLineAccessorFactory.new()
   db:mapSymbolRef( pattern, function ( item )
   
      local path = _lune.unwrap( db:getFilePath( item:get_fileId() ))
      Util.outputLocate( io.stdout, pattern, path, factory:create( path ), item:get_line() )
      return true
   end )
end
_moduleObj.InqRef = InqRef

return _moduleObj
