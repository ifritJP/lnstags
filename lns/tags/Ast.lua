--lns/tags/Ast.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Ast'
local _lune = {}
if _lune7 then
   _lune = _lune7
end
function _lune.loadModule( mod )
   if __luneScript and not package.preload[ mod ] then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune7 then
   _lune7 = _lune
end


local LuneOpt = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Option' )
local Types = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )
local Parser = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Parser' )
local front = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.front' )
local Nodes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Nodes' )
local LuneAst = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )
local LuaVer = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.LuaVer' )
local LuneLog = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Log' )
local LuneUtil = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Util' )

local function getFullNameSym( filter, symbolInfo )

   if symbolInfo:get_namespaceTypeInfo() == LuneAst.headTypeInfo then
      return symbolInfo:get_name()
   end
   
   local name = string.format( "%s.%s", filter:getFull( symbolInfo:get_namespaceTypeInfo(), false ), symbolInfo:get_name())
   return name
end
_moduleObj.getFullNameSym = getFullNameSym

local function buildAst( logLevel, pathList, projDir, stdinFile, forceAll, transCtrlInfo, astCallback )

   if #pathList == 0 then
      return 
   end
   
   
   LuneLog.setLevel( logLevel )
   LuneUtil.setDebugFlag( false )
   
   local lnsOpt = LuneOpt.createDefaultOption( pathList, projDir )
   lnsOpt:set_stdinFile( stdinFile )
   
   lnsOpt.targetLuaVer = LuaVer.ver53
   lnsOpt.transCtrlInfo.legacyMutableControl = transCtrlInfo.legacyMutableControl
   
   if forceAll then
      lnsOpt.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.ForceAll)
   else
    
      lnsOpt.transCtrlInfo.uptodateMode = _lune.newAlge( Types.CheckingUptodateMode.Force1, {LuneUtil.scriptPath2Module( pathList[1] )})
   end
   
   
   front.build( lnsOpt, astCallback )
end
_moduleObj.buildAst = buildAst

return _moduleObj
