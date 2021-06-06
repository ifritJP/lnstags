--lns/tags/Option.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Option'
local _lune = {}
if _lune4 then
   _lune = _lune4
end
function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
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

if not _lune4 then
   _lune4 = _lune
end
local Log = _lune.loadModule( 'lns.tags.Log' )
local LnsTypes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )
local LnsUtil = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Util' )

local InqMode = {}
_moduleObj.InqMode = InqMode
InqMode._val2NameMap = {}
function InqMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "InqMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function InqMode._from( val )
   if InqMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
InqMode.__allList = {}
function InqMode.get__allList()
   return InqMode.__allList
end

InqMode.Def = 'def'
InqMode._val2NameMap['def'] = 'Def'
InqMode.__allList[1] = InqMode.Def
InqMode.Ref = 'ref'
InqMode._val2NameMap['ref'] = 'Ref'
InqMode.__allList[2] = InqMode.Ref
InqMode.Set = 'set'
InqMode._val2NameMap['set'] = 'Set'
InqMode.__allList[3] = InqMode.Set
InqMode.AllMut = 'allmut'
InqMode._val2NameMap['allmut'] = 'AllMut'
InqMode.__allList[4] = InqMode.AllMut
InqMode.Async = 'async'
InqMode._val2NameMap['async'] = 'Async'
InqMode.__allList[5] = InqMode.Async
InqMode.Noasync = 'noasync'
InqMode._val2NameMap['noasync'] = 'Noasync'
InqMode.__allList[6] = InqMode.Noasync
InqMode.Luaval = 'luaval'
InqMode._val2NameMap['luaval'] = 'Luaval'
InqMode.__allList[7] = InqMode.Luaval
InqMode.AsyncLock = 'asyncLock'
InqMode._val2NameMap['asyncLock'] = 'AsyncLock'
InqMode.__allList[8] = InqMode.AsyncLock


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
Mode.Update = 'update'
Mode._val2NameMap['update'] = 'Update'
Mode.__allList[3] = Mode.Update
Mode.Inq = 'inq'
Mode._val2NameMap['inq'] = 'Inq'
Mode.__allList[4] = Mode.Inq
Mode.InqAt = 'inq-at'
Mode._val2NameMap['inq-at'] = 'InqAt'
Mode.__allList[5] = Mode.InqAt
Mode.Suffix = 'suffix'
Mode._val2NameMap['suffix'] = 'Suffix'
Mode.__allList[6] = Mode.Suffix
Mode.Dump = 'dump'
Mode._val2NameMap['dump'] = 'Dump'
Mode.__allList[7] = Mode.Dump
Mode.Test = 'test'
Mode._val2NameMap['test'] = 'Test'
Mode.__allList[8] = Mode.Test


local AnalyzeFileInfo = {}
_moduleObj.AnalyzeFileInfo = AnalyzeFileInfo
function AnalyzeFileInfo.new(  )
   local obj = {}
   AnalyzeFileInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function AnalyzeFileInfo:__init() 
   self.path = ""
   self.lineNo = 0
   self.column = 0
   self.stdinFile = nil
end
function AnalyzeFileInfo.setmeta( obj )
  setmetatable( obj, { __index = AnalyzeFileInfo  } )
end
function AnalyzeFileInfo:get_path()
   return self.path
end
function AnalyzeFileInfo:get_lineNo()
   return self.lineNo
end
function AnalyzeFileInfo:get_column()
   return self.column
end
function AnalyzeFileInfo:get_stdinFile()
   return self.stdinFile
end


local Option = {}
_moduleObj.Option = Option
function Option.new(  )
   local obj = {}
   Option.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Option:__init() 
   self.logLevel = nil
   self.pathList = {}
   self.mode = Mode.Build
   self.inqMode = InqMode.Def
   self.pattern = ""
   self.analyzeFileInfo = AnalyzeFileInfo.new()
   self.transCtrlInfo = LnsTypes.TransCtrlInfo.create_normal(  )
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
function Option:get_inqMode()
   return self.inqMode
end
function Option:get_pattern()
   return self.pattern
end
function Option:get_analyzeFileInfo()
   return self.analyzeFileInfo
end
function Option:get_logLevel()
   return self.logLevel
end
function Option:get_transCtrlInfo()
   return self.transCtrlInfo
end


local function printUsage( messages )

   if messages ~= nil then
      io.stderr:write( string.format( "%s\n", messages) )
   end
   
   print( "usage: lnstags init [option]" )
   print( "usage: lnstags build [option] filepath" )
   print( "usage: lnstags inq <def|ref|set|allmut> pattern" )
   print( "usage: lnstags inq-at <def|ref|set> filepath lineno column" )
   print( "usage: lnstags test [option]" )
   os.exit( 1 )
end

local function analyzeArgs( argList )

   local option = Option.new()
   
   local index = 1
   
   local getNextOpNonNil
   
   getNextOpNonNil = function ( mess )
   
      printUsage( "" )
   end
   
   local stdinFlag = false
   
   local function getNextOpRaw(  )
   
      if #argList <= index then
         return nil
      end
      
      index = index + 1
      return argList[index]
   end
   local function getNextOp(  )
   
      while true do
         local arg = getNextOpRaw(  )
         if  nil == arg then
            local _arg = arg
         
            return nil
         end
         
         if arg:find( "^-" ) then
            do
               local _switchExp = arg
               if _switchExp == "-i" then
                  stdinFlag = true
               elseif _switchExp == "--log" then
                  option.logLevel = Log.str2level( getNextOpNonNil( "logLevel" ) )
               elseif _switchExp == "--simpleLog" then
                  Log.enableDetail( false )
               elseif _switchExp == "--legacy-mutable-control" then
                  option.transCtrlInfo.legacyMutableControl = true
               end
            end
            
         else
          
            return arg
         end
         
      end
      
   end
   
   getNextOpNonNil = function ( mess )
   
      do
         local arg = getNextOp(  )
         if arg ~= nil then
            return arg
         end
      end
      
      printUsage( mess )
   end
   local function getNextOpInt( mess )
   
      do
         local arg = getNextOp(  )
         if arg ~= nil then
            do
               local num = tonumber( arg )
               if num ~= nil then
                  return math.floor(num)
               end
            end
            
            printUsage( string.format( "illegal num -- %s", arg) )
         end
      end
      
      printUsage( mess )
   end
   
   local function getInqMode(  )
   
      local nextToken = getNextOp(  )
      if  nil == nextToken then
         local _nextToken = nextToken
      
         printUsage( "illegal inqMod" )
      end
      
      local inqMode = InqMode._from( nextToken )
      if  nil == inqMode then
         local _inqMode = inqMode
      
         printUsage( string.format( "illegal inqMod -- %s", nextToken) )
      end
      
      return inqMode
   end
   
   local mode = nil
   
   while true do
      local arg = getNextOp(  )
      if  nil == arg then
         local _arg = arg
      
         break
      end
      
      
      if not mode then
         do
            local work = Mode._from( arg )
            if work ~= nil then
               mode = work
               do
                  local _switchExp = mode
                  if _switchExp == Mode.Inq then
                     option.inqMode = getInqMode(  )
                     do
                        local _switchExp = option.inqMode
                        if _switchExp == InqMode.AllMut or _switchExp == InqMode.Async or _switchExp == InqMode.Noasync or _switchExp == InqMode.Luaval or _switchExp == InqMode.AsyncLock then
                        else 
                           
                              option.pattern = getNextOpNonNil( "none pattern" )
                        end
                     end
                     
                     Log.setLevel( Log.Level.Warn )
                  elseif _switchExp == Mode.InqAt then
                     option.inqMode = getInqMode(  )
                     option.analyzeFileInfo.path = getNextOpNonNil( "none path" )
                     option.analyzeFileInfo.lineNo = getNextOpInt( "none lineno" )
                     option.analyzeFileInfo.column = getNextOpInt( "none column" )
                  elseif _switchExp == Mode.Suffix then
                     option.pattern = getNextOpNonNil( "none pattern" )
                  end
               end
               
            else
               printUsage( string.format( "illegal option -- %s", arg) )
            end
         end
         
      else
       
         do
            local _switchExp = mode
            if _switchExp == Mode.Build then
               if arg == "@-" then
                  while true do
                     local line = io.stdin:read( "*l" )
                     if  nil == line then
                        local _line = line
                     
                        break
                     end
                     
                     if #line > 0 then
                        table.insert( option.pathList, line )
                     end
                     
                  end
                  
               else
                
                  table.insert( option.pathList, arg )
               end
               
            end
         end
         
      end
      
   end
   
   do
      local logLevel = option.logLevel
      if logLevel ~= nil then
         Log.setLevel( logLevel )
      end
   end
   
   
   if stdinFlag then
      option.analyzeFileInfo.stdinFile = LnsTypes.StdinFile.new(LnsUtil.scriptPath2Module( option.analyzeFileInfo:get_path() ), _lune.unwrap( io.stdin:read( "*a" )))
   end
   
   
   if mode ~= nil then
      option.mode = mode
      
      return option
   end
   
   
   printUsage( "none mode" )
end
_moduleObj.analyzeArgs = analyzeArgs

return _moduleObj
