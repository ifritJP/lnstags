--lns/tags/Analyze.lns
local _moduleObj = {}
local __mod__ = '@lns.@tags.@Analyze'
local _lune = {}
if _lune8 then
   _lune = _lune8
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
   if __luneScript and not package.preload[ mod ] then
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

if not _lune8 then
   _lune8 = _lune
end



local DBCtrl = _lune.loadModule( 'lns.tags.DBCtrl' )

local Option = _lune.loadModule( 'lns.tags.Option' )

local Log = _lune.loadModule( 'lns.tags.Log' )

local Ast = _lune.loadModule( 'lns.tags.Ast' )

local LuneOpt = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Option' )

local Nodes = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Nodes' )

local TransUnit = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.TransUnit' )

local front = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.front' )

local LuneAst = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Ast' )

local AstInfo = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.AstInfo' )

local Types = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Types' )

local LuaVer = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.LuaVer' )

local LuneLog = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Log' )

local LuneUtil = _lune.loadModule( 'go/github:com.ifritJP.LuneScript.src.lune.base.Util' )


local Opt = {}
function Opt._setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end
function Opt._new(  )
   local obj = {}
   Opt._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Opt:__init(  )

end


local tagFilter = {}
setmetatable( tagFilter, { __index = Nodes.Filter } )
function tagFilter:addFileId( path, mod )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.addFileId'

   do
      local _exp = self.file2id[path]
      if _exp ~= nil then
         return _exp
      end
   end
   local fileId = self.db:addFile( path, mod )
   Log.log( Log.Level.Debug, __func__, 39, function (  )
   
      return string.format( "add file -- %d, %s", fileId, path)
   end
    )
   
   
   self.file2id[path] = fileId
   return fileId
end
function tagFilter:getFileId( path )

   do
      local _exp = self.file2id[path]
      if _exp ~= nil then
         return _exp
      end
   end
   local fileId = self.db:getFileIdFromPath( path )
   self.file2id[path] = fileId
   return fileId
end
function tagFilter._new( rootNode, db, streamName, noDefNsIdList )
   local obj = {}
   tagFilter._setmeta( obj )
   if obj.__init then obj:__init( rootNode, db, streamName, noDefNsIdList ); end
   return obj
end
function tagFilter:__init(rootNode, db, streamName, noDefNsIdList) 
   Nodes.Filter.__init( self,true, rootNode:get_moduleTypeInfo(), rootNode:get_moduleTypeInfo():get_scope())
   self.noDefNsIdList = noDefNsIdList
   self.moduleTypeInfo = rootNode:get_moduleTypeInfo()
   self.file2id = {}
   self.db = db
   self.streamName = streamName
   self.type2nsid = {}
   self.sym2nsid = {}
   local mod = self:getFull( rootNode:get_moduleTypeInfo(), false )
   self:addFileId( streamName, mod )
end
function tagFilter._setmeta( obj )
  setmetatable( obj, { __index = tagFilter  } )
end


function tagFilter:registerType( typeInfo )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registerType'

   typeInfo = typeInfo:get_nonnilableType():get_srcTypeInfo():get_genSrcTypeInfo()
   if not LuneAst.TypeInfo.hasParent( typeInfo ) then
      return DBCtrl.rootNsId, false
   end
   do
      local _exp = self.type2nsid[typeInfo]
      if _exp ~= nil then
         return _exp, false
      end
   end
   local parentNsId = self:registerType( typeInfo:get_parentInfo() )
   local name = self:getFull( typeInfo, false )
   local nsId, added = self.db:addNamespace( name, parentNsId )
   Log.log( Log.Level.Debug, __func__, 82, function (  )
   
      return string.format( "%s %s %d", name, added, nsId)
   end
    )
   
   
   self.type2nsid[typeInfo] = nsId
   return nsId, added
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
   local name = Ast.getFullNameSym( self, symbolInfo )
   local nsId, added = self.db:addNamespace( name, parentNsId )
   Log.log( Log.Level.Debug, __func__, 96, function (  )
   
      return string.format( "%s %s %d", name, added, nsId)
   end
    )
   
   
   self.sym2nsid[symbolInfo] = nsId
   return nsId, added
end


function tagFilter:registDeclSym( symbolInfo )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registDeclSym'

   local symNsId = self:registerSymbol( symbolInfo )
   local pos = _lune.unwrap( _lune.nilacc( symbolInfo:get_pos(), 'get_orgPos', 'callmtd' ))
   local fileId = self:getFileId( pos.streamName )
   self.db:addSymbolDecl( symNsId, fileId, pos.lineNo, pos.column )
   Log.log( Log.Level.Debug, __func__, 106, function (  )
   
      return symbolInfo:get_name()
   end
    )
   
   
   return symNsId
end


function tagFilter:addSymbolRef( mbrNsId, pos, setOp )

   pos = pos:get_orgPos()
   self.db:addSymbolRef( mbrNsId, self:getFileId( pos.streamName ), pos.lineNo, pos.column, setOp )
end


function tagFilter:registerDecl( nodeManager )

   local function registDeclTypeOp( typeInfo, pos )
   
      local nsId = self:registerType( typeInfo )
      pos = pos:get_orgPos()
      self.db:addSymbolDecl( nsId, self:getFileId( pos.streamName ), pos.lineNo, pos.column )
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
      for __index, altType in pairs( workNode:get_expType():get_itemTypeInfoList() ) do
         registDeclTypeOp( altType, workNode:get_pos() )
      end
   end
   for __index, workNode in pairs( nodeManager:getProtoClassNodeList(  ) ) do
      registDeclType( workNode )
   end
   for __index, workNode in pairs( nodeManager:getDeclClassNodeList(  ) ) do
      registDeclType( workNode )
      for __index, altType in pairs( workNode:get_expType():get_itemTypeInfoList() ) do
         registDeclTypeOp( altType, workNode:get_pos() )
      end
      registerAutoMethod( workNode:get_expType(), workNode:get_pos() )
   end
   for __index, workNode in pairs( nodeManager:getDeclConstrNodeList(  ) ) do
      registDeclType( workNode )
   end
   for __index, workNode in pairs( nodeManager:getDeclFuncNodeList(  ) ) do
      registDeclType( workNode )
      for __index, altType in pairs( workNode:get_expType():get_itemTypeInfoList() ) do
         registDeclTypeOp( altType, workNode:get_pos() )
      end
   end
   for __index, workNode in pairs( nodeManager:getProtoMethodNodeList(  ) ) do
      registDeclType( workNode )
   end
   local function addMethod( methodType, pos )
   
      local nsId = registDeclTypeOp( methodType, pos )
      if not methodType:get_staticFlag() then
         do
            local scope = methodType:get_parentInfo():get_scope()
            if scope ~= nil then
               scope:filterTypeInfoFieldAndIF( false, scope, LuneAst.ScopeAccess.Normal, function ( symbolInfo )
               
                  if symbolInfo:get_name() == methodType:get_rawTxt() then
                     local superNsId = self:registerType( symbolInfo:get_typeInfo() )
                     self.db:addOverride( nsId, superNsId )
                  end
                  return true
               end
                )
            end
         end
      end
      for __index, altType in pairs( methodType:get_itemTypeInfoList() ) do
         registDeclTypeOp( altType, pos )
      end
      self.db:addAsyncMode( nsId, methodType:get_asyncMode() )
      return nsId
   end
   for __index, workNode in pairs( nodeManager:getDeclMethodNodeList(  ) ) do
      addMethod( workNode:get_expType(), workNode:get_pos() )
   end
   for __index, workNode in pairs( nodeManager:getDeclEnumNodeList(  ) ) do
      registDeclType( workNode )
      registerAutoMethod( workNode:get_expType(), workNode:get_pos() )
      for __index, name in pairs( workNode:get_valueNameList() ) do
         do
            local _exp = workNode:get_scope():getSymbolInfoChild( name.txt )
            if _exp ~= nil then
               self:registDeclSym( _exp )
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
            local _1 = __map[ name ]
            do
               do
                  local _exp = _lune.nilacc( workNode:get_algeType():get_scope(), 'getSymbolInfoChild', 'callmtd' , name )
                  if _exp ~= nil then
                     self:registDeclSym( _exp )
                  end
               end
            end
         end
      end
   end
   for __index, workNode in pairs( nodeManager:getDeclMacroNodeList(  ) ) do
      registDeclType( workNode )
   end
   for __index, workNode in pairs( nodeManager:getAliasNodeList(  ) ) do
      registDeclTypeOp( workNode:get_expType(), _lune.unwrap( workNode:get_newSymbol():get_pos()) )
   end
   for __index, workNode in pairs( nodeManager:getImportNodeList(  ) ) do
      self:registDeclSym( workNode:get_info():get_symbolInfo() )
   end
   for __index, workNode in pairs( nodeManager:getDeclVarNodeList(  ) ) do
      for __index, symbolInfo in pairs( workNode:get_symbolInfoList() ) do
         if symbolInfo:get_scope() == self.moduleTypeInfo:get_scope() then
            self:registDeclSym( symbolInfo )
         end
      end
   end
   for __index, workNode in pairs( nodeManager:getDeclMemberNodeList(  ) ) do
      local mbrNsId = self:registDeclSym( workNode:get_symbolInfo() )
      if workNode:get_getterMode() ~= LuneAst.AccessMode.None then
         local name = string.format( "get_%s", workNode:get_name().txt)
         local pos = _lune.unwrap( _lune.nilacc( workNode:get_getterToken(), "pos" ))
         addMethod( _lune.unwrap( _lune.nilacc( workNode:get_classType():get_scope(), 'getTypeInfoChild', 'callmtd' , name )), pos )
         self:addSymbolRef( mbrNsId, pos, true )
      end
      if workNode:get_setterMode() ~= LuneAst.AccessMode.None then
         local name = string.format( "set_%s", workNode:get_name().txt)
         local pos = _lune.unwrap( _lune.nilacc( workNode:get_setterToken(), "pos" ))
         addMethod( _lune.unwrap( _lune.nilacc( workNode:get_classType():get_scope(), 'getTypeInfoChild', 'callmtd' , name )), pos )
         self:addSymbolRef( mbrNsId, pos, false )
      end
      if workNode:get_refType():get_mutMode() == LuneAst.MutMode.AllMut then
         self.db:addAllmutDecl( mbrNsId )
      end
   end
end


function tagFilter:registerRefs( nodeManager )
   local __func__ = '@lns.@tags.@Analyze.tagFilter.registerRefs'

   local function addSymbolRef( symbolInfo, pos, setOp )
   
      do
         local _switchExp = symbolInfo:get_name()
         if _switchExp == "__func__" or _switchExp == "__mod__" or _switchExp == "__line__" then
            return 
         end
      end
      local nsId, added = self:registerSymbol( symbolInfo )
      if added and not LuneAst.isBuiltin( symbolInfo:get_namespaceTypeInfo():get_typeId().id ) then
         table.insert( self.noDefNsIdList, nsId )
      end
      self:addSymbolRef( nsId, pos, setOp )
   end
   local function registerRefType( typeInfo, pos )
   
      local nsId, added = self:registerType( typeInfo )
      if added and not LuneAst.isBuiltin( typeInfo:get_typeId().id ) then
         table.insert( self.noDefNsIdList, nsId )
      end
      self:addSymbolRef( nsId, pos, true )
   end
   local function registerRefSym( symbolInfo, pos, setOp )
   
      do
         local _switchExp = symbolInfo:get_namespaceTypeInfo():get_kind()
         if _switchExp == LuneAst.TypeInfoKind.Enum or _switchExp == LuneAst.TypeInfoKind.Alge then
            addSymbolRef( symbolInfo, pos, true )
         else 
            
               do
                  local _switchExp = symbolInfo:get_kind()
                  if _switchExp == LuneAst.SymbolKind.Fun or _switchExp == LuneAst.SymbolKind.Mac or _switchExp == LuneAst.SymbolKind.Mtd then
                     registerRefType( symbolInfo:get_typeInfo(), pos )
                  elseif _switchExp == LuneAst.SymbolKind.Typ then
                     if symbolInfo:get_typeInfo():get_kind() == LuneAst.TypeInfoKind.Module then
                        addSymbolRef( symbolInfo, pos, true )
                     end
                     registerRefType( symbolInfo:get_typeInfo(), pos )
                  elseif _switchExp == LuneAst.SymbolKind.Mbr then
                     addSymbolRef( symbolInfo, pos, setOp )
                  elseif _switchExp == LuneAst.SymbolKind.Var then
                     if LuneAst.isPubToExternal( symbolInfo:get_accessMode() ) or symbolInfo:get_scope() == self.moduleTypeInfo:get_scope() then
                        addSymbolRef( symbolInfo, pos, setOp )
                     end
                  elseif _switchExp == LuneAst.SymbolKind.Arg then
                  end
               end
         end
      end
   end
   local function addRefLuaval( node )
   
      if node:get_expType():get_kind() == LuneAst.TypeInfoKind.Ext then
         self.db:addLuavalRef( self:getFileId( node:get_pos().streamName ), node:get_pos().lineNo, node:get_pos().column )
      end
   end
   for __index, workNode in pairs( nodeManager:getExpSetValNodeList(  ) ) do
      for __index, leftSym in pairs( workNode:get_LeftSymList() ) do
         do
            local _switchExp = leftSym:get_kind()
            if _switchExp == LuneAst.SymbolKind.Mbr then
               registerRefSym( leftSym, workNode:get_pos(), true )
            elseif _switchExp == LuneAst.SymbolKind.Var then
               if LuneAst.isPubToExternal( leftSym:get_accessMode() ) or leftSym:get_scope() == self.moduleTypeInfo:get_scope() then
                  registerRefSym( leftSym, workNode:get_pos(), true )
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
      registerRefSym( workNode:get_symbolInfo(), workNode:get_pos(), false )
      addRefLuaval( workNode )
   end
   for __index, workNode in pairs( nodeManager:getRefFieldNodeList(  ) ) do
      do
         local _exp = workNode:get_symbolInfo()
         if _exp ~= nil then
            registerRefSym( _exp, workNode:get_pos(), false )
         else
            if not workNode:get_macroArgFlag() then
               Log.log( Log.Level.Warn, __func__, 378, function (  )
               
                  return string.format( "no symbolInfo -- %s, %s:%d:%d", workNode:get_field().txt, workNode:get_pos().streamName, workNode:get_pos().lineNo, workNode:get_pos().column)
               end
                )
               
               
            end
         end
      end
      addRefLuaval( workNode )
   end
   for __index, workNode in pairs( nodeManager:getExpCallNodeList(  ) ) do
      addRefLuaval( workNode )
   end
   for __index, workNode in pairs( nodeManager:getGetFieldNodeList(  ) ) do
      registerRefType( workNode:get_getterTypeInfo(), workNode:get_pos() )
   end
   for __index, workNode in pairs( nodeManager:getExpOmitEnumNodeList(  ) ) do
      do
         local _exp = _lune.nilacc( workNode:get_enumTypeInfo():get_scope(), 'getSymbolInfoChild', 'callmtd' , workNode:get_valInfo():get_name() )
         if _exp ~= nil then
            registerRefSym( _exp, workNode:get_pos(), false )
         end
      end
   end
   for __index, workNode in pairs( nodeManager:getAsyncLockNodeList(  ) ) do
      self.db:addAsyncLock( self:getFileId( workNode:get_pos().streamName ), workNode:get_pos().lineNo, workNode:get_pos().column, workNode:get_lockKind() )
   end
   for __index, workNode in pairs( nodeManager:getMatchNodeList(  ) ) do
      for __index, matchCase in pairs( workNode:get_caseList() ) do
         local valInfo = matchCase:get_valInfo()
         do
            local _exp = _lune.nilacc( valInfo:get_algeTpye():get_scope(), 'getSymbolInfoChild', 'callmtd' , valInfo:get_name() )
            if _exp ~= nil then
               registerRefSym( _exp, matchCase:get_block():get_pos(), false )
            end
         end
      end
   end
   for __index, workNode in pairs( nodeManager:getNewAlgeValNodeList(  ) ) do
      local valInfo = workNode:get_valInfo()
      do
         local _exp = _lune.nilacc( valInfo:get_algeTpye():get_scope(), 'getSymbolInfoChild', 'callmtd' , valInfo:get_name() )
         if _exp ~= nil then
            registerRefSym( _exp, workNode:get_pos(), false )
         end
      end
   end
end


function tagFilter:processRoot( node, opt )

   self.type2nsid[LuneAst.headTypeInfo] = DBCtrl.rootNsId
   local modNsId = self:registerType( node:get_moduleTypeInfo() )
   local fileId = self:getFileId( self.streamName )
   self.db:addSymbolDecl( modNsId, fileId, 1, 1 )
   local projDir = DBCtrl.getProjDir( self.streamName, self:getFull( node:get_moduleTypeInfo(), false ) )
   for __index, workNode in pairs( node:get_nodeManager():getSubfileNodeList(  ) ) do
      do
         local usePath = workNode:get_usePath()
         if usePath ~= nil then
            local subPath = LuneUtil.pathJoin( projDir, usePath:gsub( "%.", "/" ) .. ".lns" )
            local subId = self:addFileId( subPath, usePath )
            self.db:addSubfile( fileId, subId )
         end
      end
   end
   self:registerDecl( node:get_nodeManager() )
   self:registerRefs( node:get_nodeManager() )
end



local function dumpRoot( rootNode, db, streamName, noDefNsIdList )
   local __func__ = '@lns.@tags.@Analyze.dumpRoot'

   Log.log( Log.Level.Log, __func__, 448, function (  )
   
      return streamName
   end
    )
   
   
   local filter = tagFilter._new(rootNode, db, streamName, noDefNsIdList)
   rootNode:processFilter( filter, Opt._new() )
end


local function start( db, pathList, transCtrlInfo )
   local __func__ = '@lns.@tags.@Analyze.start'

   local set = {}
   local noDefNsIdList = {}
   Ast.buildAst( LuneLog.Level.Log, pathList, nil, nil, true, transCtrlInfo, function ( ast )
   
      if not _lune._Set_has(set, ast:get_streamName() ) then
         set[ast:get_streamName()]= true
         do
            local rootNode = _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )
            if rootNode ~= nil then
               db:begin(  )
               dumpRoot( rootNode, db, ast:get_streamName(), noDefNsIdList )
               db:commit(  )
            end
         end
      end
   end
    )
   for __index, nsId in pairs( noDefNsIdList ) do
      local find = false
      db:mapSymbolDeclForNsId( nsId, function ( item )
      
         find = true
         return false
      end
       )
      if not find then
         Log.log( Log.Level.Err, __func__, 482, function (  )
         
            return string.format( "no register the define for the sym -- %s", db:getName( nsId ))
         end
          )
         
         
      end
   end
end
_moduleObj.start = start


return _moduleObj
