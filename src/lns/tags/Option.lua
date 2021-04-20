--lns/tags/Option.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Option'
local _lune = {}
if _lune3 then
   _lune = _lune3
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
local Mode = {}
_moduleObj.Mode = Mode
Mode._val2NameMap = {}
function Mode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Mode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Mode._from( val )
   if Mode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Mode.__allList = {}
function Mode.get__allList()
   return Mode.__allList
end

Mode.Init = 'init'
Mode._val2NameMap['init'] = 'Init'
Mode.__allList[1] = Mode.Init
Mode.Build = 'build'
Mode._val2NameMap['build'] = 'Build'
Mode.__allList[2] = Mode.Build
Mode.Test = 'test'
Mode._val2NameMap['test'] = 'Test'
Mode.__allList[3] = Mode.Test


local Option = {}
_moduleObj.Option = Option
function Option.new(  )
   local obj = {}
   Option.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Option:__init() 
   self.pathList = {}
   self.mode = Mode.Build
end
function Option.setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end
function Option:get_pathList()
   return self.pathList
end
function Option:get_mode()
   return self.mode
end


local function printUsage( messages )

   if messages ~= nil then
      io.stderr:write( string.format( "%s\n", messages) )
   end
   
   print( "usage: lnstags init [option]" )
   print( "usage: lnstags build [option] filepath" )
   print( "usage: lnstags test [option]" )
   os.exit( 1 )
end

local function analyzeArgs( argList )

   
   local index = 1
   local function getNextOp(  )
   
      if #argList <= index then
         return nil
      end
      
      index = index + 1
      return argList[index]
   end
   
   local mode = nil
   local option = Option.new()
   
   while true do
      local arg = getNextOp(  )
      if  nil == arg then
         local _arg = arg
      
         break
      end
      
      
      if arg:find( "^-" ) then
      else
       
         if not mode then
            do
               local work = Mode._from( arg )
               if work ~= nil then
                  mode = work
               else
                  printUsage( string.format( "illegal option -- %s", arg) )
               end
            end
            
         else
          
            table.insert( option.pathList, arg )
         end
         
      end
      
   end
   
   
   if mode ~= nil then
      option.mode = mode
      
      return option
   end
   
   
   printUsage( "none mode" )
end
_moduleObj.analyzeArgs = analyzeArgs

return _moduleObj
