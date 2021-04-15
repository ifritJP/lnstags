--lns/tags/Analyze.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Analyze'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
end

function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
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
local Option = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Option' )
local Nodes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Nodes' )
local TransUnit = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.TransUnit' )
local front = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.front' )
local Ast = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )

local Opt = {}
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end
function Opt.new(  )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Opt:__init(  )

end


local tagFilter = {}
setmetatable( tagFilter, { __index = Nodes.Filter } )
function tagFilter.setmeta( obj )
  setmetatable( obj, { __index = tagFilter  } )
end
function tagFilter.new( __superarg1, __superarg2, __superarg3 )
   local obj = {}
   tagFilter.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2, __superarg3 )
   end
   return obj
end
function tagFilter:__init( __superarg1, __superarg2, __superarg3 )

   Nodes.Filter.__init( self, __superarg1, __superarg2, __superarg3 )
end


function tagFilter:processRoot( node, opt )

   local nodeManager = node:get_nodeManager()
   for __index, workNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
      print( "declFunc:", workNode:get_pos().lineNo, self:getFull( workNode:get_expType(), false ) )
   end
   
   for __index, workNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
      print( "declMethod:", workNode:get_pos().lineNo, self:getFull( workNode:get_expType(), false ) )
   end
   
   for __index, workNode in pairs( nodeManager:getExpRefNodeList(  ) ) do
      print( "ref:", workNode:get_pos().lineNo, workNode:get_symbolInfo():get_name() )
   end
   
end



local function dumpRoot( rootNode )

   local filter = tagFilter.new(true, rootNode:get_moduleTypeInfo(), rootNode:get_moduleTypeInfo():get_scope())
   rootNode:processFilter( filter, Opt.new() )
end

local function test(  )

   local option = Option.createDefaultOption( "test/main.lns" )
   front.build( option, function ( ast )
   
      
      do
         local rootNode = _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )
         if rootNode ~= nil then
            dumpRoot( rootNode )
         end
      end
      
   end )
end
_moduleObj.test = test

return _moduleObj