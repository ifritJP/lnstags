--lns/tags/main.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@main'
local _lune = {}
if _lune8 then
   _lune = _lune8
end
function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end

function _lune.loadModule( mod )
   if __luneScript and not package.preload[ mod ] then
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

if not _lune8 then
   _lune8 = _lune
end
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Analyze = _lune.loadModule( 'lns.tags.Analyze' )
local Option = _lune.loadModule( 'lns.tags.Option' )
local Util = _lune.loadModule( 'lns.tags.Util' )
local Inq = _lune.loadModule( 'lns.tags.Inq' )
local Log = _lune.loadModule( 'lns.tags.Log' )
local Pattern = _lune.loadModule( 'lns.tags.Pattern' )
local LuneTypes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )
local LuneAst = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )

local dbPath = "lnstags.sqlite3"

local function inq( inqMode, pattern )

   local db = DBCtrl.open( dbPath, false )
   if  nil == db then
      local _db = db
   
      print( "error" )
      return -1
   end
   
   do
      local _switchExp = inqMode
      if _switchExp == Option.InqMode.Def then
         Inq.InqDef( db, pattern )
      elseif _switchExp == Option.InqMode.Ref then
         Inq.InqRef( db, pattern, false )
      elseif _switchExp == Option.InqMode.Set then
         Inq.InqRef( db, pattern, true )
      elseif _switchExp == Option.InqMode.AllMut then
         Inq.InqAllmut( db )
      elseif _switchExp == Option.InqMode.Async then
         Inq.InqAsync( db, LuneAst.Async.Async )
      elseif _switchExp == Option.InqMode.Noasync then
         Inq.InqAsync( db, LuneAst.Async.Noasync )
      elseif _switchExp == Option.InqMode.Luaval then
         Inq.InqLuaval( db )
      elseif _switchExp == Option.InqMode.AsyncLock then
         Inq.InqAsyncLock( db )
      end
   end
   
   db:close(  )
   return 0
end

local function build( pathList, transCtrlInfo )

   
   DBCtrl.initDB( dbPath )
   local db = DBCtrl.open( dbPath, false )
   if  nil == db then
      local _db = db
   
      print( "error" )
      return 1
   end
   
   db:commit(  )
   
   Analyze.start( db, pathList, transCtrlInfo )
   db:close(  )
   return 0
end

local function __main( args )

   
   local option = Option.analyzeArgs( args )
   
   do
      local _switchExp = option:get_mode()
      if _switchExp == Option.Mode.Init then
         DBCtrl.initDB( dbPath )
      elseif _switchExp == Option.Mode.Build then
         return build( option:get_pathList(), option:get_transCtrlInfo() )
      elseif _switchExp == Option.Mode.Update then
         local db = DBCtrl.open( dbPath, true )
         if  nil == db then
            local _db = db
         
            print( "error" )
            return 1
         end
         
         local projId = db:getProjId( "./" )
         local pathList = {}
         db:mapFilePath( function ( item )
         
            if item:get_projId() == projId and not db:getMainFilePath( item:get_id() ) then
               table.insert( pathList, item:get_path() )
            end
            
            return true
         end )
         db:close(  )
         
         return build( pathList, option:get_transCtrlInfo() )
      elseif _switchExp == Option.Mode.Suffix then
         local db = DBCtrl.open( dbPath, true )
         if  nil == db then
            local _db = db
         
            print( "error" )
            return 1
         end
         
         
         db:mapNamespaceSuffix( option:get_pattern(), function ( item )
         
            print( item:get_name() )
            return true
         end )
         
         db:close(  )
      elseif _switchExp == Option.Mode.Inq then
         inq( option:get_inqMode(), option:get_pattern() )
      elseif _switchExp == Option.Mode.InqAt then
         local analyzeFileInfo = option:get_analyzeFileInfo()
         
         local db = DBCtrl.open( dbPath, true )
         if  nil == db then
            local _db = db
         
            print( "error" )
            return 1
         end
         
         
         local pattern = Pattern.getPatterAt( db, analyzeFileInfo, option:get_inqMode(), option:get_transCtrlInfo() )
         if  nil == pattern then
            local _pattern = pattern
         
            db:close(  )
            print( string.format( "illegal pos -- %s:%d:%d", analyzeFileInfo:get_path(), analyzeFileInfo:get_lineNo(), analyzeFileInfo:get_column()) )
            return 1
         end
         
         db:close(  )
         
         inq( option:get_inqMode(), pattern )
      elseif _switchExp == Option.Mode.Dump then
         local db = DBCtrl.open( dbPath, true )
         if  nil == db then
            local _db = db
         
            print( "error" )
            return 1
         end
         
         db:dumpAll(  )
         
         db:close(  )
      elseif _switchExp == Option.Mode.Test then
         DBCtrl.test(  )
      end
   end
   
   
   return 0
end
_moduleObj.__main = __main

do
   local loaded, mess = _lune.loadstring52( [=[
if _lune and _lune._shebang then
  return nil
else
  return arg
end
]=] )
   if loaded ~= nil then
      local args = loaded(  )
      do
         local obj = (args )
         if obj ~= nil then
            local work = obj
            local argList = {""}
            do
               local _exp = work[0]
               if _exp ~= nil then
                  argList[1] = _exp
               end
            end
            for key, val in pairs( work ) do
               if key > 0 then
                  table.insert( argList, val )
               end
            end
            __main( argList )
         else
            -- print( "via lnsc" )
         end
      end
   else
      error( mess )
   end
end

return _moduleObj
