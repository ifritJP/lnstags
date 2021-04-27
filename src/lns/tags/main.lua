--lns/tags/main.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@main'
local _lune = require( "lune.base.runtime3" )
if not _lune3 then
   _lune3 = _lune
end
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Analyze = _lune.loadModule( 'lns.tags.Analyze' )
local Option = _lune.loadModule( 'lns.tags.Option' )
local Util = _lune.loadModule( 'lns.tags.Util' )
local Inq = _lune.loadModule( 'lns.tags.Inq' )
local Log = _lune.loadModule( 'lns.tags.Log' )

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
         Inq.InqRef( db, pattern )
      end
   end
   
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
         local db = DBCtrl.open( dbPath, false )
         if  nil == db then
            local _db = db
         
            print( "error" )
            return 1
         end
         
         db:commit(  )
         
         Analyze.start( db, option )
         db:close(  )
      elseif _switchExp == Option.Mode.Inq then
         inq( option:get_inqMode(), option:get_pattern() )
      elseif _switchExp == Option.Mode.InqAt then
         local analyzeFileInfo = option:get_analyzeFileInfo()
         local pattern = Analyze.getPatterAt( analyzeFileInfo )
         if  nil == pattern then
            local _pattern = pattern
         
            print( string.format( "illegal pos -- %s:%d:%d", analyzeFileInfo:get_path(), analyzeFileInfo:get_lineNo(), analyzeFileInfo:get_column()) )
            return 1
         end
         
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

return _moduleObj
