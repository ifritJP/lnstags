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

   do
      local path = OS.getenv( "PWD" )
      if path ~= nil then
         return path
      end
   end
   
   return "./"
end
_moduleObj.getCurDir = getCurDir

return _moduleObj
