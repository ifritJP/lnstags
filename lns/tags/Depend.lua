--lns/tags/Depend.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Depend'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
if not _lune3 then
   _lune3 = _lune
end



local OS = require( "os" )

local function getCurDir(  )

   local path
   
   do
      do
         local _exp = OS.getenv( "PWD" )
         if _exp ~= nil then
            path = _exp
         else
            path = "./"
         end
      end
      
   end
   
   return path
end
_moduleObj.getCurDir = getCurDir

return _moduleObj
