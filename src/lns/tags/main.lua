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

if not _lune3 then
   _lune3 = _lune
end
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Analyze = _lune.loadModule( 'lns.tags.Analyze' )
local Option = _lune.loadModule( 'lns.tags.Option' )

local function __main( args )

   
   local option = Option.analyzeArgs( args )
   
   do
      local _switchExp = option:get_mode()
      if _switchExp == Option.Mode.Init then
         DBCtrl.initDB(  )
      elseif _switchExp == Option.Mode.Build then
         local db = DBCtrl.open( "lnstags.sqlite3", false )
         if  nil == db then
            local _db = db
         
            print( "error" )
            return -1
         end
         
         Analyze.start( db, option )
         db:dumpAll(  )
      elseif _switchExp == Option.Mode.Test then
         DBCtrl.test(  )
         Analyze.test(  )
      end
   end
   
   
   return 0
end
_moduleObj.__main = __main

return _moduleObj
