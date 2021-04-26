// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import Ast "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
import LuaVer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsLog "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Analyze bool
var Analyze__mod__ string
// for 44
func Analyze_convExp170(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 67
func Analyze_convExp376(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 79
func Analyze_convExp516(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 131
func Analyze_convExp757(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 87
func Analyze_convExp843(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 87
func Analyze_convExp963(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 87
func Analyze_convExp1086(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 87
func Analyze_convExp1186(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}

// 52: decl @lns.@tags.@Analyze.getFullNameSym
func Analyze_getFullNameSym_1054_(filter *Nodes.Nodes_Filter,symbolInfo *Ast.Ast_SymbolInfo) string {
    if symbolInfo.FP.Get_namespaceTypeInfo() == Ast.Ast_headTypeInfo{
        return symbolInfo.FP.Get_name()
    }
    var name string
    name = Lns_getVM().String_format("%s.%s", []LnsAny{filter.FP.GetFull(symbolInfo.FP.Get_namespaceTypeInfo(), false), symbolInfo.FP.Get_name()})
    return name
}

















// 289: decl @lns.@tags.@Analyze.dumpRoot
func Analyze_dumpRoot_1117_(rootNode *Nodes.Nodes_RootNode,db *DBCtrl_DBCtrl,option *Option_Option,streamName string) {
    __func__ := "@lns.@tags.@Analyze.dumpRoot"
    Log_log(Log_Level__Log, __func__, 292, Log_CreateMessage(func() string {
        return streamName
    }))
    
    var filter *Analyze_tagFilter
    filter = NewAnalyze_tagFilter(rootNode, option, db, streamName)
    db.FP.Begin()
    rootNode.FP.ProcessFilter(&filter.Nodes_Filter, Analyze_Opt2Stem(NewAnalyze_Opt()))
    db.FP.Commit()
}

// 299: decl @lns.@tags.@Analyze.buildAst
func Analyze_buildAst_1123_(logLevel LnsInt,path string,astCallback front.Front_AstCallback) {
    LnsLog.Log_setLevel(logLevel)
    LnsUtil.Util_setDebugFlag(false)
    var lnsOpt *LnsOpt.Option_Option
    lnsOpt = LnsOpt.Option_createDefaultOption(path)
    lnsOpt.TargetLuaVer = LuaVer.LuaVer_ver53
    
    lnsOpt.TransCtrlInfo.UptodateMode = Types.Types_CheckingUptodateMode__Force
    
    front.Front_build(lnsOpt, astCallback)
}


// 314: decl @lns.@tags.@Analyze.start
func Analyze_start(db *DBCtrl_DBCtrl,option *Option_Option) {
    for _, _path := range( option.FP.Get_pathList().Items ) {
        path := _path.(string)
        Analyze_buildAst_1123_(LnsLog.Log_Level__Log, path, front.Front_AstCallback(func(ast *TransUnit.TransUnit_ASTInfo) {
            {
                _rootNode := Nodes.Nodes_RootNodeDownCastF(ast.FP.Get_node().FP)
                if _rootNode != nil {
                    rootNode := _rootNode.(*Nodes.Nodes_RootNode)
                    Analyze_dumpRoot_1117_(rootNode, db, option, ast.FP.Get_streamName())
                }
            }
        }))
    }
}








// 494: decl @lns.@tags.@Analyze.getPatterAt
func Analyze_getPatterAt(analyzeFileInfo *Option_AnalyzeFileInfo) LnsAny {
    var pattern LnsAny
    pattern = nil
    Analyze_buildAst_1123_(LnsLog.Log_Level__Err, analyzeFileInfo.FP.Get_path(), front.Front_AstCallback(func(ast *TransUnit.TransUnit_ASTInfo) {
        __func__ := "@lns.@tags.@Analyze.getPatterAt.<anonymous>"
        if ast.FP.Get_streamName() == analyzeFileInfo.FP.Get_path(){
            var filter *Analyze_SyntaxFilter
            filter = NewAnalyze_SyntaxFilter(ast)
            pattern = filter.FP.GetPattern(analyzeFileInfo)
            
            Log_log(Log_Level__Log, __func__, 501, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("pattern -- %s", []LnsAny{pattern})
            }))
            
        }
    }))
    return pattern
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
    DefaultProcess(arg1 *Nodes.Nodes_Node, arg2 LnsAny)
    GetFull(arg1 *Ast.Ast_TypeInfo, arg2 bool) string
    Get_moduleInfoManager() Ast.Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast.Ast_TypeNameCtrl
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
    registerDecl(arg1 *Nodes.Nodes_NodeManager, arg2 LnsInt, arg3 *Ast.Ast_ProcessInfo)
    registerRefs(arg1 *Nodes.Nodes_NodeManager, arg2 LnsInt)
    registerSymbol(arg1 *Ast.Ast_SymbolInfo)(LnsInt, bool)
    registerType(arg1 *Ast.Ast_TypeInfo)(LnsInt, bool)
}
type Analyze_tagFilter struct {
    Nodes.Nodes_Filter
    option *Option_Option
    db *DBCtrl_DBCtrl
    streamName string
    type2nsid *LnsMap
    sym2nsid *LnsMap
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
// 26: DeclConstr
func (self *Analyze_tagFilter) InitAnalyze_tagFilter(rootNode *Nodes.Nodes_RootNode,option *Option_Option,db *DBCtrl_DBCtrl,streamName string) {
    self.InitNodes_Filter(true, rootNode.FP.Get_moduleTypeInfo(), rootNode.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.option = option
    
    self.db = db
    
    self.streamName = streamName
    
    self.type2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.sym2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 39: decl @lns.@tags.@Analyze.tagFilter.registerType
func (self *Analyze_tagFilter) registerType(typeInfo *Ast.Ast_TypeInfo)(LnsInt, bool) {
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
    parentNsId = Analyze_convExp170(Lns_2DDD(self.FP.registerType(typeInfo.FP.Get_parentInfo())))
    var name string
    name = self.FP.GetFull(typeInfo, false)
    var nsId LnsInt
    var added bool
    nsId,added = self.db.FP.AddNamespace(name, parentNsId)
    Log_log(Log_Level__Debug, __func__, 47, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s %s %d", []LnsAny{name, added, nsId})
    }))
    
    self.type2nsid.Set(typeInfo,nsId)
    return nsId, added
}

// 61: decl @lns.@tags.@Analyze.tagFilter.registerSymbol
func (self *Analyze_tagFilter) registerSymbol(symbolInfo *Ast.Ast_SymbolInfo)(LnsInt, bool) {
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
    parentNsId = Analyze_convExp376(Lns_2DDD(self.FP.registerType(symbolInfo.FP.Get_namespaceTypeInfo())))
    var name string
    name = Analyze_getFullNameSym_1054_(&self.Nodes_Filter, symbolInfo)
    var nsId LnsInt
    var added bool
    nsId,added = self.db.FP.AddNamespace(name, parentNsId)
    Log_log(Log_Level__Debug, __func__, 70, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s %s %d", []LnsAny{name, added, nsId})
    }))
    
    self.sym2nsid.Set(symbolInfo,nsId)
    return nsId, added
}

// 75: decl @lns.@tags.@Analyze.tagFilter.registerDecl
func (self *Analyze_tagFilter) registerDecl(nodeManager *Nodes.Nodes_NodeManager,fileId LnsInt,processInfo *Ast.Ast_ProcessInfo) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerDecl"
    var registDeclTypeOp func(typeInfo *Ast.Ast_TypeInfo,pos *Types.Types_Position) LnsInt
    registDeclTypeOp = func(typeInfo *Ast.Ast_TypeInfo,pos *Types.Types_Position) LnsInt {
        var nsId LnsInt
        nsId = Analyze_convExp516(Lns_2DDD(self.FP.registerType(typeInfo)))
        self.db.FP.AddSymbolDecl(nsId, fileId, pos.LineNo, pos.Column)
        return nsId
    }
    var registDeclType func(workNode *Nodes.Nodes_Node) LnsInt
    registDeclType = func(workNode *Nodes.Nodes_Node) LnsInt {
        return registDeclTypeOp(workNode.FP.Get_expType(), workNode.FP.Get_pos())
    }
    var registerAutoMethod func(parentInfo *Ast.Ast_TypeInfo,pos *Types.Types_Position)
    registerAutoMethod = func(parentInfo *Ast.Ast_TypeInfo,pos *Types.Types_Position) {
        for _, _symbolInfo := range( Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentInfo.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.Get_symbol2SymbolInfoMap()}))).(*LnsMap).Items ) {
            symbolInfo := _symbolInfo.(Ast.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var typeInfo *Ast.Ast_TypeInfo
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
    for _, _workNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetProtoMethodNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoMethodNodeDownCast).ToNodes_ProtoMethodNode()
        registDeclType(&workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
        var nsId LnsInt
        nsId = registDeclType(&workNode.Nodes_Node)
        var methodType *Ast.Ast_TypeInfo
        methodType = workNode.FP.Get_expType()
        if Lns_op_not(methodType.FP.Get_staticFlag()){
            {
                _scope := methodType.FP.Get_parentInfo().FP.Get_scope()
                if _scope != nil {
                    scope := _scope.(*Ast.Ast_Scope)
                    scope.FP.FilterTypeInfoFieldAndIF(false, scope, Ast.Ast_ScopeAccess__Normal, Ast.Ast_filterForm(func(symbolInfo *Ast.Ast_SymbolInfo) bool {
                        if symbolInfo.FP.Get_name() == methodType.FP.Get_rawTxt(){
                            var superNsId LnsInt
                            superNsId = Analyze_convExp757(Lns_2DDD(self.FP.registerType(symbolInfo.FP.Get_typeInfo())))
                            self.db.FP.AddOverride(nsId, superNsId)
                        }
                        return true
                    }))
                }
            }
        }
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
                    _exp := __exp.(*Ast.Ast_SymbolInfo)
                    var symNsId LnsInt
                    symNsId = Analyze_convExp843(Lns_2DDD(self.FP.registerSymbol(_exp)))
                    var pos *Types.Types_Position
                    pos = Lns_unwrap( _exp.FP.Get_pos()).(*Types.Types_Position)
                    self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
                    Log_log(Log_Level__Debug, __func__, 90, Log_CreateMessage(func() string {
                        return _exp.FP.Get_name()
                    }))
                    
                    
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclAlgeNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
        registDeclType(&workNode.Nodes_Node)
        {
            __collection1035 := workNode.FP.Get_algeType().FP.Get_valInfoMap()
            __sorted1035 := __collection1035.CreateKeyListStr()
            __sorted1035.Sort( LnsItemKindStr, nil )
            for _, _name := range( __sorted1035.Items ) {
                name := _name.(string)
                {
                    __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_algeType().FP.Get_scope()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetSymbolInfoChild(name)})/* 151:14 */)
                    if __exp != nil {
                        _exp := __exp.(*Ast.Ast_SymbolInfo)
                        var symNsId LnsInt
                        symNsId = Analyze_convExp963(Lns_2DDD(self.FP.registerSymbol(_exp)))
                        var pos *Types.Types_Position
                        pos = Lns_unwrap( _exp.FP.Get_pos()).(*Types.Types_Position)
                        self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
                        Log_log(Log_Level__Debug, __func__, 90, Log_CreateMessage(func() string {
                            return _exp.FP.Get_name()
                        }))
                        
                        
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
        if Ast.Ast_isPubToExternal(workNode.FP.Get_accessMode()){
            for _, _symbolInfo := range( workNode.FP.Get_symbolInfoList().Items ) {
                symbolInfo := _symbolInfo.(Ast.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                var symNsId LnsInt
                symNsId = Analyze_convExp1086(Lns_2DDD(self.FP.registerSymbol(symbolInfo)))
                var pos *Types.Types_Position
                pos = Lns_unwrap( symbolInfo.FP.Get_pos()).(*Types.Types_Position)
                self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
                Log_log(Log_Level__Debug, __func__, 90, Log_CreateMessage(func() string {
                    return symbolInfo.FP.Get_name()
                }))
                
                
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMemberNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var symNsId LnsInt
        symNsId = Analyze_convExp1186(Lns_2DDD(self.FP.registerSymbol(workNode.FP.Get_symbolInfo())))
        var pos *Types.Types_Position
        pos = Lns_unwrap( workNode.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position)
        self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
        Log_log(Log_Level__Debug, __func__, 90, Log_CreateMessage(func() string {
            return workNode.FP.Get_symbolInfo().FP.Get_name()
        }))
        
        
        if workNode.FP.Get_getterMode() != Ast.Ast_AccessMode__None{
            var name string
            name = Lns_getVM().String_format("get_%s", []LnsAny{workNode.FP.Get_name().Txt})
            registDeclTypeOp(Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_classType().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetTypeInfoChild(name)})/* 173:20 */)).(*Ast.Ast_TypeInfo), workNode.FP.Get_pos())
        }
        if workNode.FP.Get_setterMode() != Ast.Ast_AccessMode__None{
            var name string
            name = Lns_getVM().String_format("set_%s", []LnsAny{workNode.FP.Get_name().Txt})
            registDeclTypeOp(Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_classType().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetTypeInfoChild(name)})/* 178:20 */)).(*Ast.Ast_TypeInfo), workNode.FP.Get_pos())
        }
    }
}

// 183: decl @lns.@tags.@Analyze.tagFilter.registerRefs
func (self *Analyze_tagFilter) registerRefs(nodeManager *Nodes.Nodes_NodeManager,fileId LnsInt) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs"
    var addSymbolRef func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position)
    addSymbolRef = func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.addSymbolRef"
        if _switch1370 := symbolInfo.FP.Get_name(); _switch1370 == "__func__" || _switch1370 == "__mod__" || _switch1370 == "__line__" {
            return 
        }
        var nsId LnsInt
        var added bool
        nsId,added = self.FP.registerSymbol(symbolInfo)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( added) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast.Ast_isBuiltin(symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_typeId()))) ).(bool)){
            Log_log(Log_Level__Err, __func__, 193, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("no register sym -- %d:%d:%s", []LnsAny{pos.LineNo, pos.Column, Analyze_getFullNameSym_1054_(&self.Nodes_Filter, symbolInfo)})
            }))
            
        }
        self.db.FP.AddSymbolRef(nsId, fileId, pos.LineNo, pos.Column)
    }
    var registerRefType func(typeInfo *Ast.Ast_TypeInfo,pos *Types.Types_Position)
    registerRefType = func(typeInfo *Ast.Ast_TypeInfo,pos *Types.Types_Position) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.registerRefType"
        var nsId LnsInt
        var added bool
        nsId,added = self.FP.registerType(typeInfo)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( added) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast.Ast_isBuiltin(typeInfo.FP.Get_typeId()))) ).(bool)){
            Log_log(Log_Level__Err, __func__, 203, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("no register type -- %d:%d:%s", []LnsAny{pos.LineNo, pos.Column, self.FP.GetFull(typeInfo, false)})
            }))
            
        }
        self.db.FP.AddSymbolRef(nsId, fileId, pos.LineNo, pos.Column)
    }
    var registerRefSym func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position)
    registerRefSym = func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position) {
        if _switch1723 := symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_kind(); _switch1723 == Ast.Ast_TypeInfoKind__Enum || _switch1723 == Ast.Ast_TypeInfoKind__Alge {
            addSymbolRef(symbolInfo, pos)
        } else {
            if _switch1721 := symbolInfo.FP.Get_kind(); _switch1721 == Ast.Ast_SymbolKind__Fun || _switch1721 == Ast.Ast_SymbolKind__Mac || _switch1721 == Ast.Ast_SymbolKind__Mtd || _switch1721 == Ast.Ast_SymbolKind__Typ {
                registerRefType(symbolInfo.FP.Get_typeInfo(), pos)
            } else if _switch1721 == Ast.Ast_SymbolKind__Mbr {
                addSymbolRef(symbolInfo, pos)
            } else if _switch1721 == Ast.Ast_SymbolKind__Var {
                if Ast.Ast_isPubToExternal(symbolInfo.FP.Get_accessMode()){
                    addSymbolRef(symbolInfo, pos)
                }
            } else if _switch1721 == Ast.Ast_SymbolKind__Arg {
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
                _exp := __exp.(*Ast.Ast_SymbolInfo)
                registerRefSym(_exp, workNode.FP.Get_pos())
            } else {
                Log_log(Log_Level__Warn, __func__, 257, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("no symbolInfo -- %s", []LnsAny{workNode.FP.Get_field().Txt})
                }))
                
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetExpOmitEnumNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpOmitEnumNodeDownCast).ToNodes_ExpOmitEnumNode()
        {
            __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_enumTypeInfo().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetSymbolInfoChild(workNode.FP.Get_valInfo().FP.Get_name())})/* 261:11 */)
            if __exp != nil {
                _exp := __exp.(*Ast.Ast_SymbolInfo)
                registerRefSym(_exp, workNode.FP.Get_pos())
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetMatchNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_MatchNodeDownCast).ToNodes_MatchNode()
        for _, _matchCase := range( workNode.FP.Get_caseList().Items ) {
            matchCase := _matchCase.(Nodes.Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            var valInfo *Ast.Ast_AlgeValInfo
            valInfo = matchCase.FP.Get_valInfo()
            {
                __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(valInfo.FP.Get_algeTpye().FP.Get_scope()) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetSymbolInfoChild(valInfo.FP.Get_name())})/* 269:14 */)
                if __exp != nil {
                    _exp := __exp.(*Ast.Ast_SymbolInfo)
                    registerRefSym(_exp, matchCase.FP.Get_block().FP.Get_pos())
                }
            }
        }
    }
}

// 276: decl @lns.@tags.@Analyze.tagFilter.processRoot
func (self *Analyze_tagFilter) ProcessRoot(node *Nodes.Nodes_RootNode,_opt LnsAny) {
    var fileId LnsInt
    fileId = self.db.FP.AddFile(self.streamName)
    self.type2nsid.Set(Ast.Ast_headTypeInfo,DBCtrl_rootNsId)
    self.FP.registerType(node.FP.Get_moduleTypeInfo())
    self.FP.registerDecl(node.FP.Get_nodeManager(), fileId, node.FP.Get_processInfo())
    self.FP.registerRefs(node.FP.Get_nodeManager(), fileId)
}


// declaration Class -- SyntaxFilter
type Analyze_SyntaxFilterMtd interface {
    DefaultProcess(arg1 *Nodes.Nodes_Node, arg2 LnsAny)
    GetFull(arg1 *Ast.Ast_TypeInfo, arg2 bool) string
    GetPattern(arg1 *Option_AnalyzeFileInfo) LnsAny
    getPatternFromNode(arg1 *Option_AnalyzeFileInfo, arg2 *Nodes.Nodes_Node) LnsAny
    Get_moduleInfoManager() Ast.Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast.Ast_TypeNameCtrl
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
}
type Analyze_SyntaxFilter struct {
    Nodes.Nodes_Filter
    ast *TransUnit.TransUnit_ASTInfo
    FP Analyze_SyntaxFilterMtd
}
func Analyze_SyntaxFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Analyze_SyntaxFilter).FP
}
type Analyze_SyntaxFilterDownCast interface {
    ToAnalyze_SyntaxFilter() *Analyze_SyntaxFilter
}
func Analyze_SyntaxFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Analyze_SyntaxFilterDownCast)
    if ok { return work.ToAnalyze_SyntaxFilter() }
    return nil
}
func (obj *Analyze_SyntaxFilter) ToAnalyze_SyntaxFilter() *Analyze_SyntaxFilter {
    return obj
}
func NewAnalyze_SyntaxFilter(arg1 *TransUnit.TransUnit_ASTInfo) *Analyze_SyntaxFilter {
    obj := &Analyze_SyntaxFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitAnalyze_SyntaxFilter(arg1)
    return obj
}
// 334: DeclConstr
func (self *Analyze_SyntaxFilter) InitAnalyze_SyntaxFilter(ast *TransUnit.TransUnit_ASTInfo) {
    self.InitNodes_Filter(true, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.ast = ast
    
}

// 340: decl @lns.@tags.@Analyze.SyntaxFilter.getPatternFromNode
func (self *Analyze_SyntaxFilter) getPatternFromNode(analyzeFileInfo *Option_AnalyzeFileInfo,nearest *Nodes.Nodes_Node) LnsAny {
    __func__ := "@lns.@tags.@Analyze.SyntaxFilter.getPatternFromNode"
    var isInner func(pos *Types.Types_Position,name string) bool
    isInner = func(pos *Types.Types_Position,name string) bool {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( pos.LineNo == analyzeFileInfo.FP.Get_lineNo()) &&
            Lns_GetEnv().SetStackVal( pos.Column <= analyzeFileInfo.FP.Get_column()) &&
            Lns_GetEnv().SetStackVal( pos.Column + len(name) >= analyzeFileInfo.FP.Get_column()) ).(bool)){
            return true
        }
        return false
    }
    Log_log(Log_Level__Log, __func__, 326, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s %s:%4d:%3d -- %s", []LnsAny{"nearestNode -- ", nearest.FP.Get_effectivePos().StreamName, nearest.FP.Get_effectivePos().LineNo, nearest.FP.Get_effectivePos().Column, Nodes.Nodes_getNodeKindName(nearest.FP.Get_kind())})
    }))
    
    
    {
        _workNode := Nodes.Nodes_ExpRefNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_ExpRefNode)
            return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, workNode.FP.Get_symbolInfo())
        }
    }
    {
        _workNode := Nodes.Nodes_RefFieldNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_RefFieldNode)
            if isInner(workNode.FP.Get_field().Pos, workNode.FP.Get_field().Txt){
                {
                    _symbolInfo := workNode.FP.Get_symbolInfo()
                    if _symbolInfo != nil {
                        symbolInfo := _symbolInfo.(*Ast.Ast_SymbolInfo)
                        return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, symbolInfo)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclVarNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclVarNode)
            for _, _varSym := range( workNode.FP.Get_symbolInfoList().Items ) {
                varSym := _varSym.(Ast.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if isInner(Lns_unwrap( varSym.FP.Get_pos()).(*Types.Types_Position), varSym.FP.Get_name()){
                    return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, varSym)
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpOmitEnumNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_ExpOmitEnumNode)
            if isInner(workNode.FP.Get_effectivePos(), workNode.FP.Get_valInfo().FP.Get_name()){
                return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, workNode.FP.Get_valInfo().FP.Get_symbolInfo())
            }
        }
    }
    {
        _workNode := Nodes.Nodes_NewAlgeValNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_NewAlgeValNode)
            if isInner(workNode.FP.Get_effectivePos(), workNode.FP.Get_valInfo().FP.Get_name()){
                return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, workNode.FP.Get_valInfo().FP.Get_symbolInfo())
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpMacroExpNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_ExpMacroExpNode)
            if isInner(workNode.FP.Get_effectivePos(), workNode.FP.Get_expType().FP.Get_rawTxt()){
                return self.FP.GetFull(workNode.FP.Get_macroType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclFuncNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclFuncNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if _name != nil {
                    name := _name.(*Types.Types_Token)
                    if isInner(name.Pos, name.Txt){
                        return self.FP.GetFull(workNode.FP.Get_expType(), false)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclEnumNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclEnumNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_name()
            if isInner(name.Pos, name.Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
            for _, _valInfo := range( workNode.FP.Get_enumType().FP.Get_name2EnumValInfo().Items ) {
                valInfo := _valInfo.(Ast.Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                if isInner(Lns_unwrap( valInfo.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position), valInfo.FP.Get_symbolInfo().FP.Get_name()){
                    return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, valInfo.FP.Get_symbolInfo())
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclAlgeNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclAlgeNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_name()
            if isInner(name.Pos, name.Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
            for _, _valInfo := range( workNode.FP.Get_algeType().FP.Get_valInfoMap().Items ) {
                valInfo := _valInfo.(Ast.Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                if isInner(Lns_unwrap( valInfo.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position), valInfo.FP.Get_symbolInfo().FP.Get_name()){
                    return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, valInfo.FP.Get_symbolInfo())
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclClassNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclClassNode)
            if isInner(workNode.FP.Get_name().Pos, workNode.FP.Get_name().Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMethodNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclMethodNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if _name != nil {
                    name := _name.(*Types.Types_Token)
                    if isInner(name.Pos, name.Txt){
                        return self.FP.GetFull(workNode.FP.Get_expType(), false)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ProtoClassNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_ProtoClassNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_name()
            if isInner(name.Pos, name.Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMemberNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclMemberNode)
            if isInner(Lns_unwrap( workNode.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position), workNode.FP.Get_symbolInfo().FP.Get_name()){
                return Analyze_getFullNameSym_1054_(&self.Nodes_Filter, workNode.FP.Get_symbolInfo())
            }
            
        }
    }
    {
        _workNode := Nodes.Nodes_DeclConstrNodeDownCastF(nearest.FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes.Nodes_DeclConstrNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if _name != nil {
                    name := _name.(*Types.Types_Token)
                    if isInner(name.Pos, name.Txt){
                        return self.FP.GetFull(workNode.FP.Get_expType(), false)
                    }
                }
            }
        }
    }
    Log_log(Log_Level__Err, __func__, 442, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("unknown pattern -- %s", []LnsAny{Nodes.Nodes_getNodeKindName(nearest.FP.Get_kind())})
    }))
    
    return nil
}

// 446: decl @lns.@tags.@Analyze.SyntaxFilter.getPattern
func (self *Analyze_SyntaxFilter) GetPattern(analyzeFileInfo *Option_AnalyzeFileInfo) LnsAny {
    var pattern LnsAny
    pattern = nil
    if self.ast.FP.Get_streamName() == analyzeFileInfo.FP.Get_path(){
        var nearestNode LnsAny
        nearestNode = nil
        self.ast.FP.Get_node().FP.Visit(Nodes.Nodes_NodeVisitor(func(node *Nodes.Nodes_Node,parent *Nodes.Nodes_Node,relation string,depth LnsInt) LnsInt {
            __func__ := "@lns.@tags.@Analyze.SyntaxFilter.getPattern.<anonymous>"
            if analyzeFileInfo.FP.Get_path() == node.FP.Get_effectivePos().StreamName{
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( analyzeFileInfo.FP.Get_lineNo() == node.FP.Get_effectivePos().LineNo) &&
                    Lns_GetEnv().SetStackVal( analyzeFileInfo.FP.Get_column() >= node.FP.Get_effectivePos().Column) ).(bool)){
                    {
                        _nearest := nearestNode
                        if _nearest != nil {
                            nearest := _nearest.(*Nodes.Nodes_Node)
                            if nearest.FP.Get_effectivePos().LineNo != analyzeFileInfo.FP.Get_lineNo(){
                                nearestNode = node
                                
                            }
                            if nearest.FP.Get_effectivePos().Column < node.FP.Get_effectivePos().Column{
                                nearestNode = node
                                
                            } else if nearest.FP.Get_effectivePos().Column == node.FP.Get_effectivePos().Column{
                                nearestNode = node
                                
                            }
                        } else {
                            nearestNode = node
                            
                        }
                    }
                } else { 
                    {
                        _nearest := nearestNode
                        if _nearest != nil {
                            nearest := _nearest.(*Nodes.Nodes_Node)
                            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( nearest.FP.Get_effectivePos().LineNo < node.FP.Get_pos().LineNo) &&
                                Lns_GetEnv().SetStackVal( node.FP.Get_pos().LineNo < analyzeFileInfo.FP.Get_lineNo()) ).(bool)){
                                nearestNode = node
                                
                            }
                        } else {
                            nearestNode = node
                            
                        }
                    }
                }
                Log_log(Log_Level__Trace, __func__, 326, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("%s %s:%4d:%3d -- %s", []LnsAny{"visit:", node.FP.Get_effectivePos().StreamName, node.FP.Get_effectivePos().LineNo, node.FP.Get_effectivePos().Column, Nodes.Nodes_getNodeKindName(node.FP.Get_kind())})
                }))
                
                
                return Nodes.Nodes_NodeVisitMode__Child
            }
            return Nodes.Nodes_NodeVisitMode__Next
        }), 0)
        {
            _nearest := nearestNode
            if _nearest != nil {
                nearest := _nearest.(*Nodes.Nodes_Node)
                pattern = self.FP.getPatternFromNode(analyzeFileInfo, nearest)
                
            }
        }
    }
    return pattern
}


func Lns_Analyze_init() {
    if init_Analyze { return }
    init_Analyze = true
    Analyze__mod__ = "@lns.@tags.@Analyze"
    Lns_InitMod()
    Lns_DBCtrl_init()
    Lns_Option_init()
    Lns_Log_init()
    LnsOpt.Lns_Option_init()
    Nodes.Lns_Nodes_init()
    TransUnit.Lns_TransUnit_init()
    front.Lns_front_init()
    Ast.Lns_Ast_init()
    Types.Lns_Types_init()
    LuaVer.Lns_LuaVer_init()
    LnsLog.Lns_Log_init()
    LnsUtil.Lns_Util_init()
}
func init() {
    init_Analyze = false
}
