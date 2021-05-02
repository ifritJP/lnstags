--lns/tags/main.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@main'
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
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Analyze = _lune.loadModule( 'lns.tags.Analyze' )
local Option = _lune.loadModule( 'lns.tags.Option' )
local Util = _lune.loadModule( 'lns.tags.Util' )
local Inq = _lune.loadModule( 'lns.tags.Inq' )
local Log = _lune.loadModule( 'lns.tags.Log' )
local Pattern = _lune.loadModule( 'lns.tags.Pattern' )

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
      end
   end
   
   db:close(  )
   return 0
end

local function build( pathList )

   DBCtrl.initDB( dbPath )
   local db = DBCtrl.open( dbPath, false )
   if  nil == db then
      local _db = db
   
      print( "error" )
      return 1
   end
   
   db:commit(  )
   
   Analyze.start( db, pathList )
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
         return build( option:get_pathList() )
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
         
         return build( pathList )
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
         
         
         local pattern = Pattern.getPatterAt( db, analyzeFileInfo, option:get_inqMode() )
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

return _moduleObj