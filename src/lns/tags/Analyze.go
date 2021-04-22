// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import Ast "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
var init_Analyze bool
var Analyze__mod__ string



// 197: decl @lns.@tags.@Analyze.dumpRoot
func Analyze_dumpRoot_1070_(rootNode *Nodes.Nodes_RootNode,db *DBCtrl_DBCtrl,option *Option_Option,streamName string) {
    var filter *Analyze_tagFilter
    filter = NewAnalyze_tagFilter(rootNode, option, db, streamName)
    rootNode.FP.ProcessFilter(&filter.Nodes_Filter, Analyze_Opt2Stem(NewAnalyze_Opt()))
}


// 204: decl @lns.@tags.@Analyze.start
func Analyze_start(db *DBCtrl_DBCtrl,option *Option_Option) {
    for _, _path := range( option.FP.Get_pathList().Items ) {
        path := _path.(string)
        var lnsOpt *LnsOpt.Option_Option
        lnsOpt = LnsOpt.Option_createDefaultOption(path)
        front.Front_build(lnsOpt, front.Front_AstCallback(func(ast *TransUnit.TransUnit_ASTInfo) {
            {
                _rootNode := Nodes.Nodes_RootNodeDownCastF(ast.FP.Get_node().FP)
                if _rootNode != nil {
                    rootNode := _rootNode.(*Nodes.Nodes_RootNode)
                    Analyze_dumpRoot_1070_(rootNode, db, option, ast.FP.Get_streamName())
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
    registerDecl(arg1 *Nodes.Nodes_NodeManager, arg2 LnsInt)
    registerRefs(arg1 *Nodes.Nodes_NodeManager, arg2 LnsInt)
    registerSymbol(arg1 *Ast.Ast_SymbolInfo) LnsInt
    registerType(arg1 *Ast.Ast_TypeInfo) LnsInt
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
// 23: DeclConstr
func (self *Analyze_tagFilter) InitAnalyze_tagFilter(rootNode *Nodes.Nodes_RootNode,option *Option_Option,db *DBCtrl_DBCtrl,streamName string) {
    self.InitNodes_Filter(true, rootNode.FP.Get_moduleTypeInfo(), rootNode.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.option = option
    
    self.db = db
    
    self.streamName = streamName
    
    self.type2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.sym2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 36: decl @lns.@tags.@Analyze.tagFilter.registerType
func (self *Analyze_tagFilter) registerType(typeInfo *Ast.Ast_TypeInfo) LnsInt {
    typeInfo = typeInfo.FP.Get_nonnilableType().FP.Get_srcTypeInfo().FP.Get_genSrcTypeInfo()
    
    {
        __exp := self.type2nsid.Items[typeInfo]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            return _exp
        }
    }
    var parentNsId LnsInt
    parentNsId = self.FP.registerType(typeInfo.FP.Get_parentInfo())
    var nsId LnsInt
    nsId = self.db.FP.AddNamespace(self.FP.GetFull(typeInfo, false), parentNsId)
    self.type2nsid.Set(typeInfo,nsId)
    return nsId
}

// 47: decl @lns.@tags.@Analyze.tagFilter.registerSymbol
func (self *Analyze_tagFilter) registerSymbol(symbolInfo *Ast.Ast_SymbolInfo) LnsInt {
    {
        __exp := self.sym2nsid.Items[symbolInfo]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            return _exp
        }
    }
    var parentNsId LnsInt
    parentNsId = self.FP.registerType(symbolInfo.FP.Get_namespaceTypeInfo())
    var name string
    name = Lns_getVM().String_format("%s.%s", []LnsAny{self.FP.GetFull(symbolInfo.FP.Get_namespaceTypeInfo(), false), symbolInfo.FP.Get_name()})
    var nsId LnsInt
    nsId = self.db.FP.AddNamespace(name, parentNsId)
    self.sym2nsid.Set(symbolInfo,nsId)
    return nsId
}

// 60: decl @lns.@tags.@Analyze.tagFilter.registerDecl
func (self *Analyze_tagFilter) registerDecl(nodeManager *Nodes.Nodes_NodeManager,fileId LnsInt) {
    for _, _workNode := range( nodeManager.FP.GetProtoClassNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoClassNodeDownCast).ToNodes_ProtoClassNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
    }
    for _, _workNode := range( nodeManager.FP.GetDeclClassNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
    }
    for _, _workNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
    }
    for _, _workNode := range( nodeManager.FP.GetProtoMethodNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoMethodNodeDownCast).ToNodes_ProtoMethodNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
    }
    for _, _workNode := range( nodeManager.FP.GetDeclEnumNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
        for _, _name := range( workNode.FP.Get_valueNameList().Items ) {
            name := _name.(Types.Types_TokenDownCast).ToTypes_Token()
            {
                __exp := workNode.FP.Get_scope().FP.GetSymbolInfoChild(name.Txt)
                if __exp != nil {
                    _exp := __exp.(*Ast.Ast_SymbolInfo)
                    var symNsId LnsInt
                    symNsId = self.FP.registerSymbol(_exp)
                    var pos *Types.Types_Position
                    pos = Lns_unwrap( _exp.FP.Get_pos()).(*Types.Types_Position)
                    self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
                    
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclAlgeNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
        {
            __collection744 := workNode.FP.Get_algeType().FP.Get_valInfoMap()
            __sorted744 := __collection744.CreateKeyListStr()
            __sorted744.Sort( LnsItemKindStr, nil )
            for _, _name := range( __sorted744.Items ) {
                name := _name.(string)
                {
                    __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_algeType().FP.Get_scope()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetSymbolInfoChild(name)})/* 101:14 */)
                    if __exp != nil {
                        _exp := __exp.(*Ast.Ast_SymbolInfo)
                        var symNsId LnsInt
                        symNsId = self.FP.registerSymbol(_exp)
                        var pos *Types.Types_Position
                        pos = Lns_unwrap( _exp.FP.Get_pos()).(*Types.Types_Position)
                        self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
                        
                    }
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMacroNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMacroNodeDownCast).ToNodes_DeclMacroNode()
        var nsId LnsInt
        nsId = self.FP.registerType(workNode.FP.Get_expType())
        self.db.FP.AddSymbolDecl(nsId, fileId, workNode.FP.Get_pos().LineNo, workNode.FP.Get_pos().Column)
        
    }
    for _, _workNode := range( nodeManager.FP.GetDeclVarNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
        if Ast.Ast_isPubToExternal(workNode.FP.Get_accessMode()){
            for _, _symbolInfo := range( workNode.FP.Get_symbolInfoList().Items ) {
                symbolInfo := _symbolInfo.(Ast.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                var symNsId LnsInt
                symNsId = self.FP.registerSymbol(symbolInfo)
                var pos *Types.Types_Position
                pos = Lns_unwrap( symbolInfo.FP.Get_pos()).(*Types.Types_Position)
                self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
                
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMemberNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var symNsId LnsInt
        symNsId = self.FP.registerSymbol(workNode.FP.Get_symbolInfo())
        var pos *Types.Types_Position
        pos = Lns_unwrap( workNode.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position)
        self.db.FP.AddSymbolDecl(symNsId, fileId, pos.LineNo, pos.Column)
        
    }
}

// 123: decl @lns.@tags.@Analyze.tagFilter.registerRefs
func (self *Analyze_tagFilter) registerRefs(nodeManager *Nodes.Nodes_NodeManager,fileId LnsInt) {
    var addSymbol func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position)
    addSymbol = func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.addSymbol"
        {
            _nsId := self.sym2nsid.Items[symbolInfo]
            if _nsId != nil {
                nsId := _nsId.(LnsInt)
                self.db.FP.AddSymbolRef(nsId, fileId, pos.LineNo, pos.Column)
            } else {
                Log_log(Log_Level__Err, __func__, 130, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("no register -- %s", []LnsAny{symbolInfo.FP.Get_name()})
                }))
                
            }
        }
    }
    var registerRefSym func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position)
    registerRefSym = func(symbolInfo *Ast.Ast_SymbolInfo,pos *Types.Types_Position) {
        if _switch1116 := symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_kind(); _switch1116 == Ast.Ast_TypeInfoKind__Enum || _switch1116 == Ast.Ast_TypeInfoKind__Alge {
            addSymbol(symbolInfo, pos)
        } else {
            if _switch1114 := symbolInfo.FP.Get_kind(); _switch1114 == Ast.Ast_SymbolKind__Fun || _switch1114 == Ast.Ast_SymbolKind__Mac || _switch1114 == Ast.Ast_SymbolKind__Mbr || _switch1114 == Ast.Ast_SymbolKind__Mtd || _switch1114 == Ast.Ast_SymbolKind__Typ {
                var nsId LnsInt
                nsId = self.FP.registerType(symbolInfo.FP.Get_typeInfo())
                self.db.FP.AddSymbolRef(nsId, fileId, pos.LineNo, pos.Column)
            } else if _switch1114 == Ast.Ast_SymbolKind__Var {
                if Ast.Ast_isPubToExternal(symbolInfo.FP.Get_accessMode()){
                    addSymbol(symbolInfo, pos)
                }
            } else if _switch1114 == Ast.Ast_SymbolKind__Arg {
            }
        }
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
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetExpOmitEnumNodeList().Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpOmitEnumNodeDownCast).ToNodes_ExpOmitEnumNode()
        {
            __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workNode.FP.Get_enumTypeInfo().FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetSymbolInfoChild(workNode.FP.Get_valInfo().FP.Get_name())})/* 169:11 */)
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
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast.Ast_Scope).FP.GetSymbolInfoChild(valInfo.FP.Get_name())})/* 177:14 */)
                if __exp != nil {
                    _exp := __exp.(*Ast.Ast_SymbolInfo)
                    registerRefSym(_exp, matchCase.FP.Get_block().FP.Get_pos())
                }
            }
        }
    }
}

// 184: decl @lns.@tags.@Analyze.tagFilter.processRoot
func (self *Analyze_tagFilter) ProcessRoot(node *Nodes.Nodes_RootNode,_opt LnsAny) {
    var fileId LnsInt
    fileId = self.db.FP.AddFile(self.streamName)
    self.type2nsid.Set(Ast.Ast_headTypeInfo,DBCtrl_rootNsId)
    self.FP.registerType(node.FP.Get_moduleTypeInfo())
    self.FP.registerDecl(node.FP.Get_nodeManager(), fileId)
    self.FP.registerRefs(node.FP.Get_nodeManager(), fileId)
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
}
func init() {
    init_Analyze = false
}
