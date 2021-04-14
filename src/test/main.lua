--test/main.lns
local _moduleObj = {}
local __mod__ = '@test.@main'
local _lune = require( "lune.base.runtime3" )
if not _lune3 then
   _lune3 = _lune
end
local function func(  )

   print( "hoge" )
end
local Hoge = {}
function Hoge:func(  )

end
function Hoge.sub(  )

end
function Hoge.setmeta( obj )
  setmetatable( obj, { __index = Hoge  } )
end
function Hoge.new(  )
   local obj = {}
   Hoge.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Hoge:__init(  )

end

func(  )
local hoge = Hoge.new()



return _moduleObj
