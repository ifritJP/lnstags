--lns/tags/Analyze.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Analyze'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
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
local LnsOpt = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Option' )
local Nodes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Nodes' )
local TransUnit = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.TransUnit' )
local front = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.front' )
local Ast = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )
local Types = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )
local LuaVer = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.LuaVer' )
local LnsLog = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Log' )
local LnsUtil = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Util' )

local Opt = {}
function Opt.setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end
function Opt.new(  )
   local obj = {}
   Opt.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Opt:__init(  )

end


local tagFilter = {}
setmetatable( tagFilter, { __index = Nodes.Filter } )
function tagFilter.new( rootNode, option, db, streamName )
   local obj = {}
   tagFilter.setmeta( obj )
   if obj.__init then obj:__init( rootNode, option, db, streamName ); end
   return obj
end
function tagFilter:__init(rootNode, option, db, streamName) 
   Nodes.Filter.__init( self,true, rootNode:get_moduleTypeInfo(), rootNode:get_moduleTypeInfo():get_scope())
   
   self.option = option
   self.db = db
   self.streamName = streamName
   self.type2nsid = {}
   self.sym2nsid = {}
end
function tagFilter.setmeta( obj )
  setmetatable( obj, { __index = tagFilter  } )
end


function tagFilter:registerType( typeInfo )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registerType'

   typeInfo = typeInfo:get_nonnilableType():get_srcTypeInfo():get_genSrcTypeInfo()
   do
      local _exp = self.type2nsid[typeInfo]
      if _exp ~= nil then
         return _exp, false
      end
   end
   
   local parentNsId = self:registerType( typeInfo:get_parentInfo() )
   local name = self:getFull( typeInfo, false )
   local nsId, added = self.db:addNamespace( name, parentNsId )
   Log.log( Log.Level.Debug, __func__, 47, function (  )
   
      return string.format( "%s %s %d", name, added, nsId)
   end )
   
   self.type2nsid[typeInfo] = nsId
   return nsId, added
end


local function getFullNameSym( filter, symbolInfo )

   if symbolInfo:get_namespaceTypeInfo() == Ast.headTypeInfo then
      return symbolInfo:get_name()
   end
   
   local name = string.format( "%s.%s", filter:getFull( symbolInfo:get_namespaceTypeInfo(), false ), symbolInfo:get_name())
   return name
end

function tagFilter:registerSymbol( symbolInfo )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registerSymbol'

   symbolInfo = symbolInfo:getOrg(  )
   do
      local _exp = self.sym2nsid[symbolInfo]
      if _exp ~= nil then
         return _exp, false
      end
   end
   
   local parentNsId = self:registerType( symbolInfo:get_namespaceTypeInfo() )
   local name = getFullNameSym( self, symbolInfo )
   local nsId, added = self.db:addNamespace( name, parentNsId )
   Log.log( Log.Level.Debug, __func__, 70, function (  )
   
      return string.format( "%s %s %d", name, added, nsId)
   end )
   
   self.sym2nsid[symbolInfo] = nsId
   return nsId, added
end


function tagFilter:registerDecl( nodeManager, fileId, processInfo )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registerDecl'

   local function registDeclTypeOp( typeInfo, pos )
   
      local nsId = self:registerType( typeInfo )
      self.db:addSymbolDecl( nsId, fileId, pos.lineNo, pos.column )
      return nsId
   end
   local function registDeclType( workNode )
   
      return registDeclTypeOp( workNode:get_expType(), workNode:get_pos() )
   end
   
   
   local function registerAutoMethod( parentInfo, pos )
   
      for __index, symbolInfo in pairs( _lune.unwrap( _lune.nilacc( parentInfo:get_scope(), 'get_symbol2SymbolInfoMap', 'callmtd' )) ) do
         local typeInfo = symbolInfo:get_typeInfo()
         if typeInfo:get_autoFlag() then
            registDeclTypeOp( typeInfo, pos )
         end
         
      end
      
   end
   for __index, workNode in pairs( nodeManager:getDeclFormNodeList(  ) ) do
      registDeclType( workNode )
   end
   
   
   for __index, workNode in pairs( nodeManager:getProtoClassNodeList(  ) ) do
      registDeclType( workNode )
   end
   
   
   for __index, workNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
      registDeclType( workNode )
      registerAutoMethod( workNode:get_expType(), workNode:get_pos() )
   end
   
   
   for __index, workNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
      registDeclType( workNode )
   end
   
   for __index, workNode in pairs( nodeManager:getProtoMethodNodeList(  ) ) do
      registDeclType( workNode )
   end
   
   for __index, workNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
      local nsId = registDeclType( workNode )
      local methodType = workNode:get_expType()
      if not methodType:get_staticFlag() then
         do
            local scope = methodType:get_parentInfo():get_scope()
            if scope ~= nil then
               scope:filterTypeInfoFieldAndIF( false, scope, Ast.ScopeAccess.Normal, function ( symbolInfo )
               
                  if symbolInfo:get_name() == methodType:get_rawTxt() then
                     local superNsId = self:registerType( symbolInfo:get_typeInfo() )
                     self.db:addOverride( nsId, superNsId )
                  end
                  
                  return true
               end )
            end
         end
         
      end
      
   end
   
   for __index, workNode in pairs( nodeManager:getDeclEnumNodeList(  ) ) do
      registDeclType( workNode )
      registerAutoMethod( workNode:get_expType(), workNode:get_pos() )
      for __index, name in pairs( workNode:get_valueNameList() ) do
         do
            local _exp = workNode:get_scope():getSymbolInfoChild( name.txt )
            if _exp ~= nil then
               
               local symNsId = self:registerSymbol( _exp )
               local pos = _lune.unwrap( _exp:get_pos())
               self.db:addSymbolDecl( symNsId, fileId, pos.lineNo, pos.column )
               Log.log( Log.Level.Debug, __func__, 90, function (  )
               
                  return _exp:get_name()
               end )
               
               
            end
         end
         
      end
      
   end
   
   for __index, workNode in pairs( nodeManager:getDeclAlgeNodeList(  ) ) do
      registDeclType( workNode )
      do
         local __sorted = {}
         local __map = workNode:get_algeType():get_valInfoMap()
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, name in ipairs( __sorted ) do
            local _5945 = __map[ name ]
            do
               do
                  local _exp = _lune.nilacc( workNode:get_algeType():get_scope(), 'getSymbolInfoChild', 'callmtd' , name )
                  if _exp ~= nil then
                     
                     local symNsId = self:registerSymbol( _exp )
                     local pos = _lune.unwrap( _exp:get_pos())
                     self.db:addSymbolDecl( symNsId, fileId, pos.lineNo, pos.column )
                     Log.log( Log.Level.Debug, __func__, 90, function (  )
                     
                        return _exp:get_name()
                     end )
                     
                     
                  end
               end
               
            end
         end
      end
      
   end
   
   for __index, workNode in pairs( nodeManager:getDeclMacroNodeList(  ) ) do
      registDeclType( workNode )
   end
   
   
   for __index, workNode in pairs( nodeManager:getDeclVarNodeList(  ) ) do
      if Ast.isPubToExternal( workNode:get_accessMode() ) then
         for __index, symbolInfo in pairs( workNode:get_symbolInfoList() ) do
            
            local symNsId = self:registerSymbol( symbolInfo )
            local pos = _lune.unwrap( symbolInfo:get_pos())
            self.db:addSymbolDecl( symNsId, fileId, pos.lineNo, pos.column )
            Log.log( Log.Level.Debug, __func__, 90, function (  )
            
               return symbolInfo:get_name()
            end )
            
            
         end
         
      end
      
   end
   
   for __index, workNode in pairs( nodeManager:getDeclMemberNodeList(  ) ) do
      
      local symNsId = self:registerSymbol( workNode:get_symbolInfo() )
      local pos = _lune.unwrap( workNode:get_symbolInfo():get_pos())
      self.db:addSymbolDecl( symNsId, fileId, pos.lineNo, pos.column )
      Log.log( Log.Level.Debug, __func__, 90, function (  )
      
         return workNode:get_symbolInfo():get_name()
      end )
      
      
      if workNode:get_getterMode() ~= Ast.AccessMode.None then
         local name = string.format( "get_%s", workNode:get_name().txt)
         registDeclTypeOp( _lune.unwrap( _lune.nilacc( workNode:get_classType():get_scope(), 'getTypeInfoChild', 'callmtd' , name )), workNode:get_pos() )
      end
      
      if workNode:get_setterMode() ~= Ast.AccessMode.None then
         local name = string.format( "set_%s", workNode:get_name().txt)
         registDeclTypeOp( _lune.unwrap( _lune.nilacc( workNode:get_classType():get_scope(), 'getTypeInfoChild', 'callmtd' , name )), workNode:get_pos() )
      end
      
   end
   
end


function tagFilter:registerRefs( nodeManager, fileId )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registerRefs'

   
   local function addSymbolRef( symbolInfo, pos )
      local __func__ = '@lns.@tags.@Analyze.tagFilter.registerRefs.addSymbolRef'
   
      do
         local _switchExp = symbolInfo:get_name()
         if _switchExp == "__func__" or _switchExp == "__mod__" or _switchExp == "__line__" then
            return 
         end
      end
      
      local nsId, added = self:registerSymbol( symbolInfo )
      if added and not Ast.isBuiltin( symbolInfo:get_namespaceTypeInfo():get_typeId() ) then
         Log.log( Log.Level.Err, __func__, 193, function (  )
         
            return string.format( "no register sym -- %d:%d:%s", pos.lineNo, pos.column, getFullNameSym( self, symbolInfo ))
         end )
         
      end
      
      self.db:addSymbolRef( nsId, fileId, pos.lineNo, pos.column )
   end
   
   local function registerRefType( typeInfo, pos )
      local __func__ = '@lns.@tags.@Analyze.tagFilter.registerRefs.registerRefType'
   
      local nsId, added = self:registerType( typeInfo )
      if added and not Ast.isBuiltin( typeInfo:get_typeId() ) then
         Log.log( Log.Level.Err, __func__, 203, function (  )
         
            return string.format( "no register type -- %d:%d:%s", pos.lineNo, pos.column, self:getFull( typeInfo, false ))
         end )
         
      end
      
      self.db:addSymbolRef( nsId, fileId, pos.lineNo, pos.column )
   end
   
   local function registerRefSym( symbolInfo, pos )
   
      do
         local _switchExp = symbolInfo:get_namespaceTypeInfo():get_kind()
         if _switchExp == Ast.TypeInfoKind.Enum or _switchExp == Ast.TypeInfoKind.Alge then
            addSymbolRef( symbolInfo, pos )
         else 
            
               do
                  local _switchExp = symbolInfo:get_kind()
                  if _switchExp == Ast.SymbolKind.Fun or _switchExp == Ast.SymbolKind.Mac or _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Typ then
                     registerRefType( symbolInfo:get_typeInfo(), pos )
                  elseif _switchExp == Ast.SymbolKind.Mbr then
                     addSymbolRef( symbolInfo, pos )
                  elseif _switchExp == Ast.SymbolKind.Var then
                     if Ast.isPubToExternal( symbolInfo:get_accessMode() ) then
                        addSymbolRef( symbolInfo, pos )
                     end
                     
                  elseif _switchExp == Ast.SymbolKind.Arg then
                     
                  end
               end
               
         end
      end
      
   end
   for __index, workNode in pairs( nodeManager:getExpNewNodeList(  ) ) do
      registerRefType( workNode:get_ctorTypeInfo(), workNode:get_pos() )
   end
   
   
   for __index, workNode in pairs( nodeManager:getExpCallSuperCtorNodeList(  ) ) do
      registerRefType( workNode:get_methodType(), workNode:get_pos() )
   end
   
   
   for __index, workNode in pairs( nodeManager:getExpCallSuperNodeList(  ) ) do
      registerRefType( workNode:get_methodType(), workNode:get_pos() )
   end
   
   
   for __index, workNode in pairs( nodeManager:getExpRefNodeList(  ) ) do
      registerRefSym( workNode:get_symbolInfo(), workNode:get_pos() )
   end
   
   for __index, workNode in pairs( nodeManager:getRefFieldNodeList(  ) ) do
      do
         local _exp = workNode:get_symbolInfo()
         if _exp ~= nil then
            registerRefSym( _exp, workNode:get_pos() )
         else
            Log.log( Log.Level.Warn, __func__, 257, function (  )
            
               return string.format( "no symbolInfo -- %s", workNode:get_field().txt)
            end )
            
         end
      end
      
   end
   
   for __index, workNode in pairs( nodeManager:getExpOmitEnumNodeList(  ) ) do
      do
         local _exp = _lune.nilacc( workNode:get_enumTypeInfo():get_scope(), 'getSymbolInfoChild', 'callmtd' , workNode:get_valInfo():get_name() )
         if _exp ~= nil then
            registerRefSym( _exp, workNode:get_pos() )
         end
      end
      
   end
   
   
   for __index, workNode in pairs( nodeManager:getMatchNodeList(  ) ) do
      for __index, matchCase in pairs( workNode:get_caseList() ) do
         local valInfo = matchCase:get_valInfo()
         do
            local _exp = _lune.nilacc( valInfo:get_algeTpye():get_scope(), 'getSymbolInfoChild', 'callmtd' , valInfo:get_name() )
            if _exp ~= nil then
               registerRefSym( _exp, matchCase:get_block():get_pos() )
            end
         end
         
      end
      
   end
   
end


function tagFilter:processRoot( node, opt )

   local fileId = self.db:addFile( self.streamName )
   
   self.type2nsid[Ast.headTypeInfo] = DBCtrl.rootNsId
   self:registerType( node:get_moduleTypeInfo() )
   
   self:registerDecl( node:get_nodeManager(), fileId, node:get_processInfo() )
   self:registerRefs( node:get_nodeManager(), fileId )
end



local function dumpRoot( rootNode, db, option, streamName )
   local __func__ = '@lns.@tags.@Analyze.dumpRoot'

   Log.log( Log.Level.Log, __func__, 292, function (  )
   
      return streamName
   end )
   
   local filter = tagFilter.new(rootNode, option, db, streamName)
   db:begin(  )
   rootNode:processFilter( filter, Opt.new() )
   db:commit(  )
end

local function buildAst( logLevel, path, astCallback )

   LnsLog.setLevel( logLevel )
   LnsUtil.setDebugFlag( false )
   
   local lnsOpt = LnsOpt.createDefaultOption( path )
   lnsOpt.targetLuaVer = LuaVer.ver53
   lnsOpt.transCtrlInfo.uptodateMode = Types.CheckingUptodateMode.Force
   front.build( lnsOpt, astCallback )
end

local function start( db, option )

   for __index, path in pairs( option:get_pathList() ) do
      buildAst( LnsLog.Level.Log, path, function ( ast )
      
         do
            local rootNode = _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )
            if rootNode ~= nil then
               dumpRoot( rootNode, db, option, ast:get_streamName() )
            end
         end
         
      end )
   end
   
end
_moduleObj.start = start



local SyntaxFilter = {}
setmetatable( SyntaxFilter, { __index = Nodes.Filter } )
function SyntaxFilter.new( ast )
   local obj = {}
   SyntaxFilter.setmeta( obj )
   if obj.__init then obj:__init( ast ); end
   return obj
end
function SyntaxFilter:__init(ast) 
   Nodes.Filter.__init( self,true, ast:get_moduleTypeInfo(), ast:get_moduleTypeInfo():get_scope())
   
   self.ast = ast
end
function SyntaxFilter:getPatternFromNode( analyzeFileInfo, nearest )
   local __func__ = '@lns.@tags.@Analyze.SyntaxFilter.getPatternFromNode'

   local function isInner( pos, name )
   
      if pos.lineNo == analyzeFileInfo:get_lineNo() and pos.column <= analyzeFileInfo:get_column() and pos.column + #name >= analyzeFileInfo:get_column() then
         return true
      end
      
      return false
   end
   
   
   
   Log.log( Log.Level.Log, __func__, 326, function (  )
   
      return string.format( "%s %s:%4d:%3d -- %s", "nearestNode -- ", nearest:get_effectivePos().streamName, nearest:get_effectivePos().lineNo, nearest:get_effectivePos().column, Nodes.getNodeKindName( nearest:get_kind() ))
   end )
   
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpRefNode )
      if workNode ~= nil then
         return getFullNameSym( self, workNode:get_symbolInfo() )
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.RefFieldNode )
      if workNode ~= nil then
         if isInner( workNode:get_field().pos, workNode:get_field().txt ) then
            do
               local symbolInfo = workNode:get_symbolInfo()
               if symbolInfo ~= nil then
                  return getFullNameSym( self, symbolInfo )
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
               return getFullNameSym( self, varSym )
            end
            
            
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.ExpOmitEnumNode )
      if workNode ~= nil then
         if isInner( workNode:get_effectivePos(), workNode:get_valInfo():get_name() ) then
            return getFullNameSym( self, workNode:get_valInfo():get_symbolInfo() )
         end
         
      end
   end
   
   do
      local workNode = _lune.__Cast( nearest, 3, Nodes.NewAlgeValNode )
      if workNode ~= nil then
         if isInner( workNode:get_effectivePos(), workNode:get_valInfo():get_name() ) then
            return getFullNameSym( self, workNode:get_valInfo():get_symbolInfo() )
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
               return getFullNameSym( self, valInfo:get_symbolInfo() )
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
               return getFullNameSym( self, valInfo:get_symbolInfo() )
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
            return getFullNameSym( self, workNode:get_symbolInfo() )
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
   
   Log.log( Log.Level.Err, __func__, 442, function (  )
   
      return string.format( "unknown pattern -- %s", Nodes.getNodeKindName( nearest:get_kind() ))
   end )
   
   return nil
end
function SyntaxFilter:getPattern( analyzeFileInfo )

   local pattern = nil
   if self.ast:get_streamName() == analyzeFileInfo:get_path() then
      local nearestNode = nil
      self.ast:get_node():visit( function ( node, parent, relation, depth )
         local __func__ = '@lns.@tags.@Analyze.SyntaxFilter.getPattern.<anonymous>'
      
         if analyzeFileInfo:get_path() == node:get_effectivePos().streamName then
            if analyzeFileInfo:get_lineNo() == node:get_effectivePos().lineNo and analyzeFileInfo:get_column() >= node:get_effectivePos().column then
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
                     if nearest:get_effectivePos().lineNo < node:get_pos().lineNo and node:get_pos().lineNo < analyzeFileInfo:get_lineNo() then
                        nearestNode = node
                     end
                     
                  else
                     nearestNode = node
                  end
               end
               
            end
            
            
            Log.log( Log.Level.Trace, __func__, 326, function (  )
            
               return string.format( "%s %s:%4d:%3d -- %s", "visit:", node:get_effectivePos().streamName, node:get_effectivePos().lineNo, node:get_effectivePos().column, Nodes.getNodeKindName( node:get_kind() ))
            end )
            
            
            return Nodes.NodeVisitMode.Child
         end
         
         return Nodes.NodeVisitMode.Next
      end, 0 )
      do
         local nearest = nearestNode
         if nearest ~= nil then
            pattern = self:getPatternFromNode( analyzeFileInfo, nearest )
         end
      end
      
   end
   
   return pattern
end
function SyntaxFilter.setmeta( obj )
  setmetatable( obj, { __index = SyntaxFilter  } )
end


local function getPatterAt( analyzeFileInfo )

   
   local pattern = nil
   buildAst( LnsLog.Level.Err, analyzeFileInfo:get_path(), function ( ast )
      local __func__ = '@lns.@tags.@Analyze.getPatterAt.<anonymous>'
   
      if ast:get_streamName() == analyzeFileInfo:get_path() then
         local filter = SyntaxFilter.new(ast)
         pattern = filter:getPattern( analyzeFileInfo )
         Log.log( Log.Level.Log, __func__, 501, function (  )
         
            return string.format( "pattern -- %s", pattern)
         end )
         
      end
      
   end )
   
   return pattern
end
_moduleObj.getPatterAt = getPatterAt

return _moduleObj
