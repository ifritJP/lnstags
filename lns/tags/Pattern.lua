--lns/tags/Pattern.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Pattern'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
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

if not _lune3 then
   _lune3 = _lune
end
local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )
local Option = _lune.loadModule( 'lns.tags.Option' )
local Log = _lune.loadModule( 'lns.tags.Log' )
local Ast = _lune.loadModule( 'lns.tags.Ast' )
local LnsOpt = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Option' )
local Nodes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Nodes' )
local Parser = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Parser' )
local TransUnit = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.TransUnit' )
local front = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.front' )
local Types = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )
local LnsAst = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )
local LuaVer = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.LuaVer' )
local LnsLog = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Log' )
local LnsUtil = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Util' )



local SyntaxFilter = {}
setmetatable( SyntaxFilter, { __index = Nodes.Filter } )
function SyntaxFilter.new( ast )
   local obj = {}
   SyntaxFilter.setmeta( obj )
   if obj.__init then obj:__init( ast ); end
   return obj
end
function SyntaxFilter:__init(ast) 
   Nodes.Filter.__init( self,true, ast:get_exportInfo():get_moduleTypeInfo(), ast:get_exportInfo():get_moduleTypeInfo():get_scope())
   
   self.ast = ast
end
function SyntaxFilter:getPatternFromNode( analyzeFileInfo, inqMod, nearest )
   local __func__ = '@lns.@tags.@Pattern.SyntaxFilter.getPatternFromNode'

   local function isInner( pos, name )
   
      if pos.lineNo == analyzeFileInfo:get_lineNo() and pos.column <= analyzeFileInfo:get_column() and pos.column + #name >= analyzeFileInfo:get_column() then
         return true
      end
      
      return false
   end
   
   
   
   
   Log.log( Log.Level.Log, __func__, 19, function (  )
   
      return string.format( "%s %s:%4d:%3d -- %s", "nearestNode -- ", nearest:get_effectivePos().streamName, nearest:get_effectivePos().lineNo, nearest:get_effectivePos().column, Nodes.getNodeKindName( nearest:get_kind() ))
   end )
   
   
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ImportNode )
      if workNode ~= nil then
         do
            local _switchExp = inqMod
            if _switchExp == Option.InqMode.Def then
               return self:getFull( workNode:get_expType(), false )
            elseif _switchExp == Option.InqMode.Ref then
               return Ast.getFullNameSym( self, workNode:get_symbolInfo() )
            elseif _switchExp == Option.InqMode.Set then
            end
         end
         
         return nil
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpRefNode )
      if workNode ~= nil then
         return Ast.getFullNameSym( self, workNode:get_symbolInfo() )
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.RefFieldNode )
      if workNode ~= nil then
         if isInner( workNode:get_field().pos, workNode:get_field().txt ) then
            do
               local symbolInfo = workNode:get_symbolInfo()
               if symbolInfo ~= nil then
                  return Ast.getFullNameSym( self, symbolInfo )
               end
            end
            
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclVarNode )
      if workNode ~= nil then
         for __index, varSym in pairs( workNode:get_symbolInfoList() ) do
            if isInner( _lune.unwrap( varSym:get_pos()), varSym:get_name() ) then
               return Ast.getFullNameSym( self, varSym )
            end
            
            
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpOmitEnumNode )
      if workNode ~= nil then
         if isInner( workNode:get_effectivePos(), workNode:get_valInfo():get_name() ) then
            return Ast.getFullNameSym( self, workNode:get_valInfo():get_symbolInfo() )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.NewAlgeValNode )
      if workNode ~= nil then
         if isInner( workNode:get_effectivePos(), workNode:get_valInfo():get_name() ) then
            return Ast.getFullNameSym( self, workNode:get_valInfo():get_symbolInfo() )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpMacroExpNode )
      if workNode ~= nil then
         if isInner( workNode:get_effectivePos(), workNode:get_expType():get_rawTxt() ) then
            return self:getFull( workNode:get_macroType(), false )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclFuncNode )
      if workNode ~= nil then
         do
            local name = workNode:get_declInfo():get_name()
            if name ~= nil then
               if isInner( name.pos, name.txt ) then
                  return self:getFull( workNode:get_expType(), false )
               end
               
            end
         end
         
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclEnumNode )
      if workNode ~= nil then
         local name = workNode:get_name()
         if isInner( name.pos, name.txt ) then
            return self:getFull( workNode:get_expType(), false )
         end
         
         for __index, valInfo in pairs( workNode:get_enumType():get_name2EnumValInfo() ) do
            if isInner( _lune.unwrap( valInfo:get_symbolInfo():get_pos()), valInfo:get_symbolInfo():get_name() ) then
               return Ast.getFullNameSym( self, valInfo:get_symbolInfo() )
            end
            
            
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclAlgeNode )
      if workNode ~= nil then
         local name = workNode:get_name()
         if isInner( name.pos, name.txt ) then
            return self:getFull( workNode:get_expType(), false )
         end
         
         for __index, valInfo in pairs( workNode:get_algeType():get_valInfoMap() ) do
            if isInner( _lune.unwrap( valInfo:get_symbolInfo():get_pos()), valInfo:get_symbolInfo():get_name() ) then
               return Ast.getFullNameSym( self, valInfo:get_symbolInfo() )
            end
            
            
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclClassNode )
      if workNode ~= nil then
         if isInner( workNode:get_name().pos, workNode:get_name().txt ) then
            return self:getFull( workNode:get_expType(), false )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclMethodNode )
      if workNode ~= nil then
         do
            local name = workNode:get_declInfo():get_name()
            if name ~= nil then
               if isInner( name.pos, name.txt ) then
                  return self:getFull( workNode:get_expType(), false )
               end
               
            end
         end
         
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ProtoClassNode )
      if workNode ~= nil then
         local name = workNode:get_name()
         if isInner( name.pos, name.txt ) then
            return self:getFull( workNode:get_expType(), false )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclMemberNode )
      if workNode ~= nil then
         if isInner( _lune.unwrap( workNode:get_symbolInfo():get_pos()), workNode:get_symbolInfo():get_name() ) then
            return Ast.getFullNameSym( self, workNode:get_symbolInfo() )
         end
         
         
         do
            local token, symbol = workNode:get_getterToken(), workNode:getGetterSym(  )
            if token ~= nil and  symbol ~= nil then
               if isInner( token.pos, token.txt ) then
                  return Ast.getFullNameSym( self, symbol )
               end
               
            end
         end
         
         do
            local token, symbol = workNode:get_setterToken(), workNode:getSetterSym(  )
            if token ~= nil and  symbol ~= nil then
               if isInner( token.pos, token.txt ) then
                  return Ast.getFullNameSym( self, symbol )
               end
               
            end
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclConstrNode )
      if workNode ~= nil then
         do
            local name = workNode:get_declInfo():get_name()
            if name ~= nil then
               if isInner( name.pos, name.txt ) then
                  return self:getFull( workNode:get_expType(), false )
               end
               
            end
         end
         
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpCallSuperCtorNode )
      if workNode ~= nil then
         return self:getFull( workNode:get_methodType(), false )
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpCallSuperNode )
      if workNode ~= nil then
         return self:getFull( workNode:get_methodType(), false )
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpNewNode )
      if workNode ~= nil then
         if isInner( workNode:get_pos(), "new" ) then
            return self:getFull( workNode:get_ctorTypeInfo(), false )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclMacroNode )
      if workNode ~= nil then
         local name = workNode:get_declInfo():get_name()
         if isInner( name.pos, name.txt ) then
            return self:getFull( workNode:get_expType(), false )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpMacroExpNode )
      if workNode ~= nil then
         if isInner( workNode:get_pos(), workNode:get_macroType():get_rawTxt() ) then
            return self:getFull( workNode:get_macroType(), false )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.GetFieldNode )
      if workNode ~= nil then
         do
            local symbolInfo = workNode:get_symbolInfo()
            if symbolInfo ~= nil then
               if isInner( workNode:get_field().pos, workNode:get_field().txt ) then
                  return Ast.getFullNameSym( self, symbolInfo )
               end
               
            end
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.AliasNode )
      if workNode ~= nil then
         if isInner( _lune.unwrap( workNode:get_newSymbol():get_pos()), workNode:get_newSymbol():get_name() ) then
            return Ast.getFullNameSym( self, workNode:get_newSymbol() )
         end
         
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.DeclFormNode )
      if workNode ~= nil then
         do
            local name = workNode:get_declInfo():get_name()
            if name ~= nil then
               if isInner( name.pos, name.txt ) then
                  return self:getFull( workNode:get_expType(), false )
               end
               
            end
         end
         
         
      end
   end
   
   
   Log.log( Log.Level.Err, __func__, 192, function (  )
   
      return string.format( "unknown pattern -- %s", Nodes.getNodeKindName( nearest:get_kind() ))
   end )
   
   return nil
end
function SyntaxFilter:getPattern( path, analyzeFileInfo, inqMod )

   local function isInner( pos, name )
   
      if pos.lineNo == analyzeFileInfo:get_lineNo() and pos.column <= analyzeFileInfo:get_column() and pos.column + #name >= analyzeFileInfo:get_column() then
         return true
      end
      
      return false
   end
   
   local pattern = nil
   if self.ast:get_streamName() == path then
      local nearestNode = nil
      self.ast:get_node():visit( function ( node, parent, relation, depth )
         local __func__ = '@lns.@tags.@Pattern.SyntaxFilter.getPattern.<anonymous>'
      
         if analyzeFileInfo:get_path() == node:get_effectivePos().streamName then
            if analyzeFileInfo:get_lineNo() == node:get_effectivePos().lineNo and analyzeFileInfo:get_column() >= node:get_effectivePos().column then
               do
                  local declMemberNode = _lune.__Cast( node, 3, Nodes.DeclMemberNode )
                  if declMemberNode ~= nil then
                     do
                        local token = declMemberNode:get_getterToken()
                        if token ~= nil then
                           if isInner( token.pos, token.txt ) then
                              nearestNode = node
                              return Nodes.NodeVisitMode.End
                           end
                           
                        end
                     end
                     
                     do
                        local token = declMemberNode:get_setterToken()
                        if token ~= nil then
                           if isInner( token.pos, token.txt ) then
                              nearestNode = node
                              return Nodes.NodeVisitMode.End
                           end
                           
                        end
                     end
                     
                  end
               end
               
               
               do
                  local nearest = nearestNode
                  if nearest ~= nil then
                     if nearest:get_effectivePos().lineNo ~= analyzeFileInfo:get_lineNo() then
                        nearestNode = node
                     end
                     
                     if nearest:get_effectivePos().column < node:get_effectivePos().column then
                        nearestNode = node
                     elseif nearest:get_effectivePos().column == node:get_effectivePos().column then
                        nearestNode = node
                     end
                     
                  else
                     nearestNode = node
                  end
               end
               
            else
             
               do
                  local nearest = nearestNode
                  if nearest ~= nil then
                     if nearest:get_effectivePos().lineNo < node:get_effectivePos().lineNo and node:get_effectivePos().lineNo < analyzeFileInfo:get_lineNo() then
                        nearestNode = node
                     end
                     
                  else
                     nearestNode = node
                  end
               end
               
            end
            
            Log.log( Log.Level.Trace, __func__, 19, function (  )
            
               return string.format( "%s %s:%4d:%3d -- %s", "visit:", node:get_effectivePos().streamName, node:get_effectivePos().lineNo, node:get_effectivePos().column, Nodes.getNodeKindName( node:get_kind() ))
            end )
            
            
            return Nodes.NodeVisitMode.Child
         end
         
         return Nodes.NodeVisitMode.Next
      end, 0, {} )
      do
         local nearest = nearestNode
         if nearest ~= nil then
            pattern = self:getPatternFromNode( analyzeFileInfo, inqMod, nearest )
         end
      end
      
   end
   
   return pattern
end
function SyntaxFilter.setmeta( obj )
  setmetatable( obj, { __index = SyntaxFilter  } )
end


local function getPatterAt( db, analyzeFileInfo, inqMod, transCtrlInfo )

   
   local fileId = db:getFileIdFromPath( analyzeFileInfo:get_path() )
   local path = db:getMainFilePath( fileId )
   if  nil == path then
      local _path = path
   
      path = analyzeFileInfo:get_path()
   end
   
   
   local pattern = nil
   local useStdInMod
   
   
   local projDir = nil
   if path:find( "^%.%." ) or path:find( "^/" ) then
      local dir = path:gsub( "/[^/]+$", "" )
      projDir = LnsUtil.searchProjDir( dir )
   end
   
   
   if analyzeFileInfo:get_stdinFlag() then
      useStdInMod = LnsUtil.scriptPath2ModuleFromProjDir( analyzeFileInfo:get_path(), projDir )
   else
    
      useStdInMod = nil
   end
   
   
   Ast.buildAst( LnsLog.Level.Err, {path}, projDir, useStdInMod, false, transCtrlInfo, function ( ast )
      local __func__ = '@lns.@tags.@Pattern.getPatterAt.<anonymous>'
   
      if ast:get_streamName() == path then
         local filter = SyntaxFilter.new(ast)
         pattern = filter:getPattern( path, analyzeFileInfo, inqMod )
         Log.log( Log.Level.Log, __func__, 305, function (  )
         
            return string.format( "pattern -- %s", pattern)
         end )
         
      end
      
   end )
   
   return pattern
end
_moduleObj.getPatterAt = getPatterAt

return _moduleObj
