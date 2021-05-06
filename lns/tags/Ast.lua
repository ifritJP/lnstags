--lns/tags/Ast.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Ast'
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

if not _lune3 then
   _lune3 = _lune
end
local LnsOpt = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Option' )
local Types = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )
local Parser = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Parser' )
local front = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.front' )
local Nodes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Nodes' )
local LnsAst = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )
local LuaVer = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.LuaVer' )
local LnsLog = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Log' )
local LnsUtil = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Util' )

local function getFullNameSym( filter, symbolInfo )

   if symbolInfo:get_namespaceTypeInfo() == LnsAst.headTypeInfo then
      return symbolInfo:get_name()
   end
   
   local name = string.format( "%s.%s", filter:getFull( symbolInfo:get_namespaceTypeInfo(), false ), symbolInfo:get_name())
   return name
end
_moduleObj.getFullNameSym = getFullNameSym

local function buildAst( logLevel, pathList, projDir, useStdInMod, forceAll, transCtrlInfo, astCallback )

   if #pathList == 0 then
      return 
   end
   
   
   LnsLog.setLevel( logLevel )
   LnsUtil.setDebugFlag( false )
   
   if useStdInMod ~= nil then
      Parser.StreamParser.setStdinStream( useStdInMod )
   end
   
   
   local lnsOpt = LnsOpt.createDefaultOption( pathList, projDir )
   lnsOpt.targetLuaVer = LuaVer.ver53
   lnsOpt.transCtrlInfo.legacyMutableControl = transCtrlInfo.legacyMutableControl
   if forceAll then
      lnsOpt.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.ForceAll)
   else
    
      lnsOpt.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.Force1, {LnsUtil.scriptPath2Module( pathList[1] )})
   end
   
   
   front.build( lnsOpt, astCallback )
end
_moduleObj.buildAst = buildAst

return _moduleObj
