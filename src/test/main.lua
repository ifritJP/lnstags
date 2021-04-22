--test/main.lns
local _moduleObj = {}
local __mod__ = '@test.@main'
local _lune = require( "lune.base.runtime3" )
if not _lune3 then
   _lune3 = _lune
end
local Sub = _lune.loadModule( 'test.Sub' )
Sub.func(  )
local hoge = Sub.Hoge.new()
hoge:func( Sub.Foo.Val1 )
Sub.Hoge.sub( _lune.newAlge( Sub.Bar.Val2) )



Sub.func(  )


return _moduleObj
