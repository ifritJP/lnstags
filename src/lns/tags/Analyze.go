// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import Parser "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import LnsAst "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
import LuaVer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsLog "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Analyze bool
var Analyze__mod__ string
// for 33
func Analyze_convExp93(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 69
func Analyze_convExp369(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 83
func Analyze_convExp520(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 92
func Analyze_convExp645(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 113
func Analyze_convExp801(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 164
func Analyze_convExp1062(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
















// 342: decl @lns.@tags.@Analyze.dumpRoot
func Analyze_dumpRoot_1121_(rootNode *Nodes.Nodes_RootNode,db *DBCtrl_DBCtrl,option *Option_Option,streamName string) {
    __func__ := "@lns.@tags.@Analyze.dumpRoot"
    Log_log(Log_Level__Log, __func__, 345, Log_CreateMessage(func() string {
        return streamName
    }))
    
    var filter *Analyze_tagFilter
    filter = NewAnalyze_tagFilter(rootNode, option, db, streamName)
    rootNode.FP.ProcessFilter(&filter.Nodes_Filter, Analyze_Opt2Stem(NewAnalyze_Opt()))
}


// 350: decl @lns.@tags.@Analyze.start
func Analyze_start(db *DBCtrl_DBCtrl,option *Option_Option) {
    for _, _path := range( option.FP.Get_pathList().Items ) {
        path := _path.(string)
        Ast_buildAst(LnsLog.Log_Level__Log, path, nil, nil, true, front.Front_AstCallback(func(ast *TransUnit.TransUnit_ASTInfo) {
            {
                _rootNode := Nodes.Nodes_RootNodeDownCastF(ast.FP.Get_node().FP)
                if _rootNode != nil {
                    rootNode := _rootNode.(*Nodes.Nodes_RootNode)
                    db.FP.Begin()
                    Analyze_dumpRoot_1121_(rootNode, db, option, ast.FP.Get_streamName())
                    db.FP.Commit()
                }
            }
        }))
    }
}

// declaration Class -- Opt
type Analyze_OptMtd interface {
}
type Analyze_Opt struct {
    FP Analyze_OptMtd
}
func Analyze_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Analyze_Opt).FP
}
type Analyze_OptDownCast interface {
    ToAnalyze_Opt() *Analyze_Opt
}
func Analyze_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Analyze_OptDownCast)
    if ok { return work.ToAnalyze_Opt() }
    return nil
}
func (obj *Analyze_Opt) ToAnalyze_Opt() *Analyze_Opt {
    return obj
}
func NewAnalyze_Opt() *Analyze_Opt {
    obj := &Analyze_Opt{}
    obj.FP = obj
    obj.InitAnalyze_Opt()
    return obj
}
func (self *Analyze_Opt) InitAnalyze_Opt() {
}

// declaration Class -- tagFilter
type Analyze_tagFilterMtd interface {
    addFileId(arg1 string, arg2 string) LnsInt
    addSymbolRef(arg1 LnsInt, arg2 *Types.Types_Position)
    DefaultProcess(arg1 *Nodes.Nodes_Node, arg2 LnsAny)
    getFileId(arg1 string) LnsInt
    GetFull(arg1 *LnsAst.Ast_TypeInfo, arg2 bool) string
    Get_moduleInfoManager() LnsAst.Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *LnsAst.Ast_TypeNameCtrl
    ProcessAbbr(arg1 *Nodes.Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(arg1 *Nodes.Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(arg1 *Nodes.Nodes_ApplyNode, arg2 LnsAny)
    ProcessBlankLine(arg1 *Nodes.Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(arg1 *Nodes.Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(arg1 *Nodes.Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(arg1 *Nodes.Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(arg1 *Nodes.Nodes_BreakNode, arg2 LnsAny)
    ProcessConvStat(arg1 *Nodes.Nodes_ConvStatNode, arg2 LnsAny)
    ProcessDeclAdvertise(arg1 *Nodes.Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(arg1 *Nodes.Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(arg1 *Nodes.Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(arg1 *Nodes.Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(arg1 *Nodes.Nodes_DeclClassNode, arg2 LnsAny)
    ProcessDeclConstr(arg1 *Nodes.Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(arg1 *Nodes.Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(arg1 *Nodes.Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(arg1 *Nodes.Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(arg1 *Nodes.Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(arg1 *Nodes.Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(arg1 *Nodes.Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(arg1 *Nodes.Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclVar(arg1 *Nodes.Nodes_DeclVarNode, arg2 LnsAny)
    ProcessEnv(arg1 *Nodes.Nodes_EnvNode, arg2 LnsAny)
    ProcessExpAccessMRet(arg1 *Nodes.Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(arg1 *Nodes.Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(arg1 *Nodes.Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(arg1 *Nodes.Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(arg1 *Nodes.Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(arg1 *Nodes.Nodes_ExpListNode, arg2 LnsAny)
    ProcessExpMRet(arg1 *Nodes.Nodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMacroArgExp(arg1 *Nodes.Nodes_ExpMacroArgExpNode, arg2 LnsAny)
    ProcessExpMacroExp(arg1 *Nodes.Nodes_ExpMacroExpNode, arg2 LnsAny)
    ProcessExpMacroStat(arg1 *Nodes.Nodes_ExpMacroStatNode, arg2 LnsAny)
    ProcessExpMacroStatList(arg1 *Nodes.Nodes_ExpMacroStatListNode, arg2 LnsAny)
    ProcessExpMultiTo1(arg1 *Nodes.Nodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(arg1 *Nodes.Nodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(arg1 *Nodes.Nodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(arg1 *Nodes.Nodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(arg1 *Nodes.Nodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(arg1 *Nodes.Nodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(arg1 *Nodes.Nodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(arg1 *Nodes.Nodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSetItem(arg1 *Nodes.Nodes_ExpSetItemNode, arg2 LnsAny)
    ProcessExpSetVal(arg1 *Nodes.Nodes_ExpSetValNode, arg2 LnsAny)
    ProcessExpSubDDD(arg1 *Nodes.Nodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(arg1 *Nodes.Nodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(arg1 *Nodes.Nodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessFor(arg1 *Nodes.Nodes_ForNode, arg2 LnsAny)
    ProcessForeach(arg1 *Nodes.Nodes_ForeachNode, arg2 LnsAny)
    ProcessForsort(arg1 *Nodes.Nodes_ForsortNode, arg2 LnsAny)
    ProcessGetField(arg1 *Nodes.Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(arg1 *Nodes.Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(arg1 *Nodes.Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(arg1 *Nodes.Nodes_ImportNode, arg2 LnsAny)
    ProcessLiteralArray(arg1 *Nodes.Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(arg1 *Nodes.Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(arg1 *Nodes.Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(arg1 *Nodes.Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(arg1 *Nodes.Nodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(arg1 *Nodes.Nodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralNil(arg1 *Nodes.Nodes_LiteralNilNode, arg2 LnsAny)
    ProcessLiteralReal(arg1 *Nodes.Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(arg1 *Nodes.Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(arg1 *Nodes.Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(arg1 *Nodes.Nodes_LiteralSymbolNode, arg2 LnsAny)
    ProcessLuneControl(arg1 *Nodes.Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(arg1 *Nodes.Nodes_LuneKindNode, arg2 LnsAny)
    ProcessMatch(arg1 *Nodes.Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(arg1 *Nodes.Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(arg1 *Nodes.Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(arg1 *Nodes.Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(arg1 *Nodes.Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(arg1 *Nodes.Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(arg1 *Nodes.Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(arg1 *Nodes.Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(arg1 *Nodes.Nodes_RepeatNode, arg2 LnsAny)
    ProcessReturn(arg1 *Nodes.Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(arg1 *Nodes.Nodes_RootNode, arg2 LnsAny)
    ProcessScope(arg1 *Nodes.Nodes_ScopeNode, arg2 LnsAny)
    ProcessShebang(arg1 *Nodes.Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(arg1 *Nodes.Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(arg1 *Nodes.Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(arg1 *Nodes.Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(arg1 *Nodes.Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(arg1 *Nodes.Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(arg1 *Nodes.Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(arg1 *Nodes.Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(arg1 *Nodes.Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(arg1 *Nodes.Nodes_WhileNode, arg2 LnsAny)
    registDeclSym(arg1 *LnsAst.Ast_SymbolInfo) LnsInt
    registerDecl(arg1 *Nodes.Nodes_NodeManager, arg2 *LnsAst.Ast_ProcessInfo)
    registerRefs(arg1 *Nodes.Nodes_NodeManager)
    registerSymbol(arg1 *LnsAst.Ast_SymbolInfo)(LnsInt, bool)
    registerType(arg1 *LnsAst.Ast_TypeInfo)(LnsInt, bool)
}
type Analyze_tagFilter struct {
    Nodes.Nodes_Filter
    option *Option_Option
    db *DBCtrl_DBCtrl
    streamName string
    type2nsid *LnsMap
    sym2nsid *LnsMap
    file2id *LnsMap
    FP Analyze_tagFilterMtd
}
func Analyze_tagFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Analyze_tagFilter).FP
}
type Analyze_tagFilterDownCast interface {
    ToAnalyze_tagFilter() *Analyze_tagFilter
}
func Analyze_tagFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Analyze_tagFilterDownCast)
    if ok { return work.ToAnalyze_tagFilter() }
    return nil
}
func (obj *Analyze_tagFilter) ToAnalyze_tagFilter() *Analyze_tagFilter {
    return obj
}
func NewAnalyze_tagFilter(arg1 *Nodes.Nodes_RootNode, arg2 *Option_Option, arg3 *DBCtrl_DBCtrl, arg4 string) *Analyze_tagFilter {
    obj := &Analyze_tagFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitAnalyze_tagFilter(arg1, arg2, arg3, arg4)
    return obj
}
// 29: decl @lns.@tags.@Analyze.tagFilter.addFileId
func (self *Analyze_tagFilter) addFileId(path string,mod string) LnsInt {
    __func__ := "@lns.@tags.@Analyze.tagFilter.addFileId"
    {
        __exp := self.file2id.Items[path]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            return _exp
        }
    }
    var fileId LnsInt
    fileId = Analyze_convExp93(Lns_2DDD(self.db.FP.AddFile(path, mod)))
    Log_log(Log_Level__Debug, __func__, 34, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("add file -- %d, %s", []LnsAny{fileId, path})
    }))
    
    self.file2id.Set(path,fileId)
    return fileId
}

// 39: decl @lns.@tags.@Analyze.tagFilter.getFileId
func (self *Analyze_tagFilter) getFileId(path string) LnsInt {
    {
        __exp := self.file2id.Items[path]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            return _exp
        }
    }
    var fileId LnsInt
    fileId = self.db.FP.GetFileIdFromPath(path)
    self.file2id.Set(path,fileId)
    return fileId
}

// 49: DeclConstr
func (self *Analyze_tagFilter) InitAnalyze_tagFilter(rootNode *Nodes.Nodes_RootNode,option *Option_Option,db *DBCtrl_DBCtrl,streamName string) {
    self.InitNodes_Filter(true, rootNode.FP.Get_moduleTypeInfo(), rootNode.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.file2id = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.option = option
    
    self.db = db
    
    self.streamName = streamName
    
    self.type2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.sym2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
    var mod string
    mod = self.FP.GetFull(rootNode.FP.Get_moduleTypeInfo(), false)
    self.FP.addFileId(streamName, mod)
}

// 64: decl @lns.@tags.@Analyze.tagFilter.registerType
func (self *Analyze_tagFilter) registerType(typeInfo *LnsAst.Ast_TypeInfo)(LnsInt, bool) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerType"
    typeInfo = typeInfo.FP.Get_nonnilableType().FP.Get_srcTypeInfo().FP.Get_genSrcTypeInfo()
    
    {
        __exp := self.type2nsid.Items[typeInfo]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            return _exp, false
        }
    }
    var parentNsId LnsInt
    parentNsId = Analyze_convExp369(Lns_2DDD(self.FP.registerType(typeInfo.FP.Get_parentInfo())))
    var name string
    name = self.FP.GetFull(typeInfo, false)
    var nsId LnsInt
    var added bool
    nsId,added = self.db.FP.AddNamespace(name, parentNsId)
    Log_log(Log_Level__Debug, __func__, 72, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s %s %d", []LnsAny{name, added, nsId})
    }))
    
    self.type2nsid.Set(typeInfo,nsId)
    return nsId, added
}

// 77: decl @lns.@tags.@Analyze.tagFilter.registerSymbol
func (self *Analyze_tagFilter) registerSymbol(symbolInfo *LnsAst.Ast_SymbolInfo)(LnsInt, bool) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerSymbol"
    symbolInfo = symbolInfo.FP.GetOrg()
    
    {
        __exp := self.sym2nsid.Items[symbolInfo]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            return _exp, false
        }
    }
    var parentNsId LnsInt
    parentNsId = Analyze_convExp520(Lns_2DDD(self.FP.registerType(symbolInfo.FP.Get_namespaceTypeInfo())))
    var name string
    name = Ast_getFullNameSym(&self.Nodes_Filter, symbolInfo)
    var nsId LnsInt
    var added bool
    nsId,added = self.db.FP.AddNamespace(name, parentNsId)
    Log_log(Log_Level__Debug, __func__, 86, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s %s %d", []LnsAny{name, added, nsId})
    }))
    
    self.sym2nsid.Set(symbolInfo,nsId)
    return nsId, added
}

// 91: decl @lns.@tags.@Analyze.tagFilter.registDeclSym
func (self *Analyze_tagFilter) registDeclSym(symbolInfo *LnsAst.Ast_SymbolInfo) LnsInt {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registDeclSym"
    var symNsId LnsInt
    symNsId = Analyze_convExp645(Lns_2DDD(self.FP.registerSymbol(symbolInfo)))
    var pos *Types.Types_Position
    pos = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(symbolInfo.FP.Get_pos()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Types.Types_Position).FP.Get_orgPos()}))).(*Types.Types_Position)
    var fileId LnsInt
    fileId = self.FP.getFileId(pos.StreamName)
    self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
    Log_log(Log_Level__Debug, __func__, 96, Log_CreateMessage(func() string {
        return symbolInfo.FP.Get_name()
    }))
    
    return symNsId
}

// 100: decl @lns.@tags.@Analyze.tagFilter.addSymbolRef
func (self *Analyze_tagFilter) addSymbolRef(mbrNsId LnsInt,pos *Types.Types_Position) {
    pos = pos.FP.Get_orgPos()
    
    self.db.FP.AddSymbolRef(mbrNsId, self.FP.getFileId(pos.StreamName), pos.LineNo, pos.Column)
}

// 109: decl @lns.@tags.@Analyze.tagFilter.registerDecl
func (self *Analyze_tagFilter) registerDecl(nodeManager *Nodes.Nodes_NodeManager,processInfo *LnsAst.Ast_ProcessInfo) {
    var registDeclTypeOp func(typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt
    registDeclTypeOp = func(typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt {
        var nsId LnsInt
        nsId = Analyze_convExp801(Lns_2DDD(self.FP.registerType(typeInfo)))
        self.db.FP.AddSymbolDecl(nsId, self.FP.getFileId(pos.StreamName), pos.LineNo, pos.Column)
        return nsId
    }
    var registDeclType func(workNode *Nodes.Nodes_Node) LnsInt
    registDeclType = func(workNode *Nodes.Nodes_Node) LnsInt {
        return registDeclTypeOp(workNode.FP.Get_expType(), workNode.FP.Get_pos())
    }
    var registerAutoMethod func(parentInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position)
    registerAutoMethod = func(parentInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) {
        for _, _symbolInfo := range( Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentInfo.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*LnsAst.Ast_Scope).FP.Get_symbol2SymbolInfoMap()}))).(*LnsMap).Items ) {
            symbolInfo := _symbolInfo.(LnsAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var typeInfo *LnsAst.Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo()
            if typeInfo.FP.Get_autoFlag(){
                registDeclTypeOp(typeInfo, pos)
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclFormNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetProtoClassNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoClassNodeDownCast).ToNodes_ProtoClassNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetDeclClassNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        registDeclType(&workNode.Nodes_Node)
        registerAutoMethod(workNode.FP.Get_expType(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetDeclConstrNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclConstrNodeDownCast).ToNodes_DeclConstrNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetProtoMethodNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoMethodNodeDownCast).ToNodes_ProtoMethodNode()
        registDeclType(&workNode.Nodes_Node)
    }
    var addMethod func(methodType *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt
    addMethod = func(methodType *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt {
        var nsId LnsInt
        nsId = registDeclTypeOp(methodType, pos)
        if Lns_op_not(methodType.FP.Get_staticFlag()){
            {
                _scope := methodType.FP.Get_parentInfo().FP.Get_scope()
                if _scope != nil {
                    scope := _scope.(*LnsAst.Ast_Scope)
                    scope.FP.FilterTypeInfoFieldAndIF(false, scope, LnsAst.Ast_ScopeAccess__Normal, LnsAst.Ast_filterForm(func(symbolInfo *LnsAst.Ast_SymbolInfo) bool {
                        if symbolInfo.FP.Get_name() == methodType.FP.Get_rawTxt(){
                            var superNsId LnsInt
                            superNsId = Analyze_convExp1062(Lns_2DDD(self.FP.registerType(symbolInfo.FP.Get_typeInfo())))
                            self.db.FP.AddOverride(nsId, superNsId)
                        }
                        return true
                    }))
                }
            }
        }
        return nsId
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
        addMethod(workNode.FP.Get_expType(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetDeclEnumNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
        registDeclType(&workNode.Nodes_Node)
        registerAutoMethod(workNode.FP.Get_expType(), workNode.FP.Get_pos())
        for _, _name := range( workNode.FP.Get_valueNameList().Items ) {
            name := _name.(Types.Types_TokenDownCast).ToTypes_Token()
            {
                __exp := workNode.FP.Get_scope().FP.GetSymbolInfoChild(name.Txt)
                if __exp != nil {
                    _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                    self.FP.registDeclSym(_exp)
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclAlgeNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
        registDeclType(&workNode.Nodes_Node)
        {
            __collection1204 := workNode.FP.Get_algeType().FP.Get_valInfoMap()
            __sorted1204 := __collection1204.CreateKeyListStr()
            __sorted1204.Sort( LnsItemKindStr, nil )
            for _, _name := range( __sorted1204.Items ) {
                name := _name.(string)
                {
                    __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_algeType().FP.Get_scope()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*LnsAst.Ast_Scope).FP.GetSymbolInfoChild(name)})/* 188:14 */)
                    if __exp != nil {
                        _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                        self.FP.registDeclSym(_exp)
                    }
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMacroNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMacroNodeDownCast).ToNodes_DeclMacroNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetDeclVarNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
        if LnsAst.Ast_isPubToExternal(workNode.FP.Get_accessMode()){
            for _, _symbolInfo := range( workNode.FP.Get_symbolInfoList().Items ) {
                symbolInfo := _symbolInfo.(LnsAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.registDeclSym(symbolInfo)
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMemberNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var mbrNsId LnsInt
        mbrNsId = self.FP.registDeclSym(workNode.FP.Get_symbolInfo())
        if workNode.FP.Get_getterMode() != LnsAst.Ast_AccessMode__None{
            var name string
            name = Lns_getVM().String_format("get_%s", []LnsAny{workNode.FP.Get_name().Txt})
            var pos *Types.Types_Position
            pos = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_getterToken()) && 
            Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types.Types_Token).Pos))).(*Types.Types_Position)
            addMethod(Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_classType().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*LnsAst.Ast_Scope).FP.GetTypeInfoChild(name)})/* 211:20 */)).(*LnsAst.Ast_TypeInfo), pos)
            self.FP.addSymbolRef(mbrNsId, pos)
        }
        if workNode.FP.Get_setterMode() != LnsAst.Ast_AccessMode__None{
            var name string
            name = Lns_getVM().String_format("set_%s", []LnsAny{workNode.FP.Get_name().Txt})
            var pos *Types.Types_Position
            pos = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_setterToken()) && 
            Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types.Types_Token).Pos))).(*Types.Types_Position)
            addMethod(Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_classType().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*LnsAst.Ast_Scope).FP.GetTypeInfoChild(name)})/* 219:20 */)).(*LnsAst.Ast_TypeInfo), pos)
            self.FP.addSymbolRef(mbrNsId, pos)
        }
    }
}

// 226: decl @lns.@tags.@Analyze.tagFilter.registerRefs
func (self *Analyze_tagFilter) registerRefs(nodeManager *Nodes.Nodes_NodeManager) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs"
    var addSymbolRef func(symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position)
    addSymbolRef = func(symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.addSymbolRef"
        if _switch1408 := symbolInfo.FP.Get_name(); _switch1408 == "__func__" || _switch1408 == "__mod__" || _switch1408 == "__line__" {
            return 
        }
        var nsId LnsInt
        var added bool
        nsId,added = self.FP.registerSymbol(symbolInfo)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( added) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(LnsAst.Ast_isBuiltin(symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_typeId()))) ).(bool)){
            Log_log(Log_Level__Err, __func__, 236, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("no register sym -- %d:%d:%s", []LnsAny{pos.LineNo, pos.Column, Ast_getFullNameSym(&self.Nodes_Filter, symbolInfo)})
            }))
            
        }
        self.FP.addSymbolRef(nsId, pos)
    }
    var registerRefType func(typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position)
    registerRefType = func(typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.registerRefType"
        var nsId LnsInt
        var added bool
        nsId,added = self.FP.registerType(typeInfo)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( added) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(LnsAst.Ast_isBuiltin(typeInfo.FP.Get_typeId()))) ).(bool)){
            Log_log(Log_Level__Err, __func__, 245, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("no register type -- %d:%d:%s", []LnsAny{pos.LineNo, pos.Column, self.FP.GetFull(typeInfo, false)})
            }))
            
        }
        self.FP.addSymbolRef(nsId, pos)
    }
    var registerRefSym func(symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position)
    registerRefSym = func(symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position) {
        if _switch1749 := symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_kind(); _switch1749 == LnsAst.Ast_TypeInfoKind__Enum || _switch1749 == LnsAst.Ast_TypeInfoKind__Alge {
            addSymbolRef(symbolInfo, pos)
        } else {
            if _switch1747 := symbolInfo.FP.Get_kind(); _switch1747 == LnsAst.Ast_SymbolKind__Fun || _switch1747 == LnsAst.Ast_SymbolKind__Mac || _switch1747 == LnsAst.Ast_SymbolKind__Mtd || _switch1747 == LnsAst.Ast_SymbolKind__Typ {
                registerRefType(symbolInfo.FP.Get_typeInfo(), pos)
            } else if _switch1747 == LnsAst.Ast_SymbolKind__Mbr {
                addSymbolRef(symbolInfo, pos)
            } else if _switch1747 == LnsAst.Ast_SymbolKind__Var {
                if LnsAst.Ast_isPubToExternal(symbolInfo.FP.Get_accessMode()){
                    addSymbolRef(symbolInfo, pos)
                }
            } else if _switch1747 == LnsAst.Ast_SymbolKind__Arg {
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetExpNewNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpNewNodeDownCast).ToNodes_ExpNewNode()
        registerRefType(workNode.FP.Get_ctorTypeInfo(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetExpCallSuperCtorNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpCallSuperCtorNodeDownCast).ToNodes_ExpCallSuperCtorNode()
        registerRefType(workNode.FP.Get_methodType(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetExpCallSuperNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpCallSuperNodeDownCast).ToNodes_ExpCallSuperNode()
        registerRefType(workNode.FP.Get_methodType(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetExpRefNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpRefNodeDownCast).ToNodes_ExpRefNode()
        registerRefSym(workNode.FP.Get_symbolInfo(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetRefFieldNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_RefFieldNodeDownCast).ToNodes_RefFieldNode()
        {
            __exp := workNode.FP.Get_symbolInfo()
            if __exp != nil {
                _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                registerRefSym(_exp, workNode.FP.Get_pos())
            } else {
                Log_log(Log_Level__Warn, __func__, 298, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("no symbolInfo -- %s", []LnsAny{workNode.FP.Get_field().Txt})
                }))
                
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetGetFieldNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_GetFieldNodeDownCast).ToNodes_GetFieldNode()
        registerRefType(workNode.FP.Get_getterTypeInfo(), workNode.FP.Get_pos())
    }
    for _, _workNode := range( nodeManager.FP.GetExpOmitEnumNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpOmitEnumNodeDownCast).ToNodes_ExpOmitEnumNode()
        {
            __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_enumTypeInfo().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*LnsAst.Ast_Scope).FP.GetSymbolInfoChild(workNode.FP.Get_valInfo().FP.Get_name())})/* 305:11 */)
            if __exp != nil {
                _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                registerRefSym(_exp, workNode.FP.Get_pos())
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetMatchNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_MatchNodeDownCast).ToNodes_MatchNode()
        for _, _matchCase := range( workNode.FP.Get_caseList().Items ) {
            matchCase := _matchCase.(Nodes.Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            var valInfo *LnsAst.Ast_AlgeValInfo
            valInfo = matchCase.FP.Get_valInfo()
            {
                __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(valInfo.FP.Get_algeTpye().FP.Get_scope()) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*LnsAst.Ast_Scope).FP.GetSymbolInfoChild(valInfo.FP.Get_name())})/* 313:14 */)
                if __exp != nil {
                    _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                    registerRefSym(_exp, matchCase.FP.Get_block().FP.Get_pos())
                }
            }
        }
    }
}

// 320: decl @lns.@tags.@Analyze.tagFilter.processRoot
func (self *Analyze_tagFilter) ProcessRoot(node *Nodes.Nodes_RootNode,_opt LnsAny) {
    self.type2nsid.Set(LnsAst.Ast_headTypeInfo,DBCtrl_rootNsId)
    self.FP.registerType(node.FP.Get_moduleTypeInfo())
    var projDir string
    projDir = DBCtrl_getProjDir(self.streamName, self.FP.GetFull(node.FP.Get_moduleTypeInfo(), false))
    for _, _workNode := range( node.FP.Get_nodeManager().FP.GetSubfileNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_SubfileNodeDownCast).ToNodes_SubfileNode()
        {
            _usePath := workNode.FP.Get_usePath()
            if _usePath != nil {
                usePath := _usePath.(string)
                var subPath string
                subPath = LnsUtil.Util_pathJoin(projDir, Lns_car(Lns_getVM().String_gsub(usePath,"%.", "/")).(string) + ".lns")
                var subId LnsInt
                subId = self.FP.addFileId(subPath, usePath)
                self.db.FP.AddSubfile(self.FP.getFileId(self.streamName), subId)
            }
        }
    }
    self.FP.registerDecl(node.FP.Get_nodeManager(), node.FP.Get_processInfo())
    self.FP.registerRefs(node.FP.Get_nodeManager())
}


func Lns_Analyze_init() {
    if init_Analyze { return }
    init_Analyze = true
    Analyze__mod__ = "@lns.@tags.@Analyze"
    Lns_InitMod()
    Lns_DBCtrl_init()
    Lns_Option_init()
    Lns_Log_init()
    Lns_Ast_init()
    LnsOpt.Lns_Option_init()
    Nodes.Lns_Nodes_init()
    Parser.Lns_Parser_init()
    TransUnit.Lns_TransUnit_init()
    front.Lns_front_init()
    LnsAst.Lns_Ast_init()
    Types.Lns_Types_init()
    LuaVer.Lns_LuaVer_init()
    LnsLog.Lns_Log_init()
    LnsUtil.Lns_Util_init()
}
func init() {
    init_Analyze = false
}
