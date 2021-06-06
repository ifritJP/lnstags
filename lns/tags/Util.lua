--lns/tags/Util.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Util'
local _lune = {}
if _lune4 then
   _lune = _lune4
end
function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
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
local SourceCodeLineAccessor = {}
_moduleObj.SourceCodeLineAccessor = SourceCodeLineAccessor
function SourceCodeLineAccessor:getLine( lineNo )

   if lineNo < 0 or #self.lineList < lineNo then
      return nil
   end
   
   return self.lineList[lineNo]
end
function SourceCodeLineAccessor.setmeta( obj )
  setmetatable( obj, { __index = SourceCodeLineAccessor  } )
end
function SourceCodeLineAccessor.new( path, lineList )
   local obj = {}
   SourceCodeLineAccessor.setmeta( obj )
   if obj.__init then
      obj:__init( path, lineList )
   end
   return obj
end
function SourceCodeLineAccessor:__init( path, lineList )

   self.path = path
   self.lineList = lineList
end
function SourceCodeLineAccessor:get_path()
   return self.path
end


local SourceCodeLineAccessorFactory = {}
_moduleObj.SourceCodeLineAccessorFactory = SourceCodeLineAccessorFactory
function SourceCodeLineAccessorFactory.new(  )
   local obj = {}
   SourceCodeLineAccessorFactory.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function SourceCodeLineAccessorFactory:__init() 
   self.path2accessor = {}
end
function SourceCodeLineAccessorFactory:create( filePath, fileContents )

   do
      local _exp = self.path2accessor[filePath]
      if _exp ~= nil then
         return _exp
      end
   end
   
   local lineList = {}
   if fileContents ~= nil then
      for line in string.gmatch( fileContents, "[^\n]*\n" ) do
         table.insert( lineList, string.sub( line, 1, #line - 1 ) )
      end
      
   else
      local handle = io.open( filePath, "r" )
      if  nil == handle then
         local _handle = handle
      
         return nil
      end
      
      while true do
         local text = handle:read( '*l' )
         if  nil == text then
            local _text = text
         
            break
         end
         
         table.insert( lineList, text )
      end
      
   end
   
   local accessor = SourceCodeLineAccessor.new(filePath, lineList)
   self.path2accessor[filePath] = accessor
   return accessor
end
function SourceCodeLineAccessorFactory.setmeta( obj )
  setmetatable( obj, { __index = SourceCodeLineAccessorFactory  } )
end


local function outputLocate( stream, symbol, path, lineAccessor, lineNo )

   local line = _lune.nilacc( lineAccessor, 'getLine', 'callmtd' , lineNo )
   if  nil == line then
      local _line = line
   
      line = ""
   end
   
   stream:write( string.format( "%-16s %4d %-16s %s\n", symbol, lineNo, path, line) )
end
_moduleObj.outputLocate = outputLocate

return _moduleObj
