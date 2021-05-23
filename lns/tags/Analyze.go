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
// for 35
func Analyze_convExp95(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 71
func Analyze_convExp368(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 73
func Analyze_convExp393(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 85
func Analyze_convExp519(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 87
func Analyze_convExp546(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 94
func Analyze_convExp644(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 115
func Analyze_convExp801(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 176
func Analyze_convExp1118(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 258
func Analyze_convExp1533(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 267
func Analyze_convExp1666(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 383
func Analyze_convExp2302(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
















// 404: decl @lns.@tags.@Analyze.dumpRoot
func Analyze_dumpRoot_1090_(_env *LnsEnv, rootNode *Nodes.Nodes_RootNode,db *DBCtrl_DBCtrl,streamName string) {
    __func__ := "@lns.@tags.@Analyze.dumpRoot"
    Log_log(_env, Log_Level__Log, __func__, 406, Log_CreateMessage(func(_env *LnsEnv) string {
        return streamName
    }))
    
    var filter *Analyze_tagFilter
    filter = NewAnalyze_tagFilter(_env, rootNode, db, streamName)
    rootNode.FP.ProcessFilter(_env, &filter.Nodes_Filter, Analyze_Opt2Stem(NewAnalyze_Opt(_env)))
}


// 411: decl @lns.@tags.@Analyze.start
func Analyze_start(_env *LnsEnv, db *DBCtrl_DBCtrl,pathList *LnsList,transCtrlInfo *Types.Types_TransCtrlInfo) {
    var set *LnsSet
    set = NewLnsSet([]LnsAny{})
    Ast_buildAst(_env, LnsLog.Log_Level__Log, pathList, nil, nil, true, transCtrlInfo, front.Front_AstCallback(func(_env *LnsEnv, ast *TransUnit.TransUnit_ASTInfo) {
        if Lns_op_not(set.Has(ast.FP.Get_streamName(_env))){
            set.Add(ast.FP.Get_streamName(_env))
            {
                _rootNode := Nodes.Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)
                if !Lns_IsNil( _rootNode ) {
                    rootNode := _rootNode.(*Nodes.Nodes_RootNode)
                    db.FP.Begin(_env)
                    Analyze_dumpRoot_1090_(_env, rootNode, db, ast.FP.Get_streamName(_env))
                    db.FP.Commit(_env)
                }
            }
        }
    }))
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
func NewAnalyze_Opt(_env *LnsEnv) *Analyze_Opt {
    obj := &Analyze_Opt{}
    obj.FP = obj
    obj.InitAnalyze_Opt(_env)
    return obj
}
func (self *Analyze_Opt) InitAnalyze_Opt(_env *LnsEnv) {
}

// declaration Class -- tagFilter
type Analyze_tagFilterMtd interface {
    addFileId(_env *LnsEnv, arg1 string, arg2 string) LnsInt
    addSymbolRef(_env *LnsEnv, arg1 LnsInt, arg2 *Types.Types_Position, arg3 bool)
    DefaultProcess(_env *LnsEnv, arg1 *Nodes.Nodes_Node, arg2 LnsAny)
    getFileId(_env *LnsEnv, arg1 string) LnsInt
    GetFull(_env *LnsEnv, arg1 *LnsAst.Ast_TypeInfo, arg2 bool) string
    Get_moduleInfoManager(_env *LnsEnv) LnsAst.Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *LnsAst.Ast_TypeNameCtrl
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes.Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes.Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(_env *LnsEnv, arg1 *Nodes.Nodes_ApplyNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes.Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes.Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes.Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes.Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes.Nodes_BreakNode, arg2 LnsAny)
    ProcessConvStat(_env *LnsEnv, arg1 *Nodes.Nodes_ConvStatNode, arg2 LnsAny)
    ProcessDeclAdvertise(_env *LnsEnv, arg1 *Nodes.Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(_env *LnsEnv, arg1 *Nodes.Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(_env *LnsEnv, arg1 *Nodes.Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(_env *LnsEnv, arg1 *Nodes.Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(_env *LnsEnv, arg1 *Nodes.Nodes_DeclClassNode, arg2 LnsAny)
    ProcessDeclConstr(_env *LnsEnv, arg1 *Nodes.Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(_env *LnsEnv, arg1 *Nodes.Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(_env *LnsEnv, arg1 *Nodes.Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(_env *LnsEnv, arg1 *Nodes.Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(_env *LnsEnv, arg1 *Nodes.Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(_env *LnsEnv, arg1 *Nodes.Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(_env *LnsEnv, arg1 *Nodes.Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(_env *LnsEnv, arg1 *Nodes.Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes.Nodes_DeclVarNode, arg2 LnsAny)
    ProcessEnv(_env *LnsEnv, arg1 *Nodes.Nodes_EnvNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes.Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(_env *LnsEnv, arg1 *Nodes.Nodes_ExpListNode, arg2 LnsAny)
    ProcessExpMRet(_env *LnsEnv, arg1 *Nodes.Nodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMacroArgExp(_env *LnsEnv, arg1 *Nodes.Nodes_ExpMacroArgExpNode, arg2 LnsAny)
    ProcessExpMacroExp(_env *LnsEnv, arg1 *Nodes.Nodes_ExpMacroExpNode, arg2 LnsAny)
    ProcessExpMacroStat(_env *LnsEnv, arg1 *Nodes.Nodes_ExpMacroStatNode, arg2 LnsAny)
    ProcessExpMacroStatList(_env *LnsEnv, arg1 *Nodes.Nodes_ExpMacroStatListNode, arg2 LnsAny)
    ProcessExpMultiTo1(_env *LnsEnv, arg1 *Nodes.Nodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(_env *LnsEnv, arg1 *Nodes.Nodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(_env *LnsEnv, arg1 *Nodes.Nodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(_env *LnsEnv, arg1 *Nodes.Nodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(_env *LnsEnv, arg1 *Nodes.Nodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(_env *LnsEnv, arg1 *Nodes.Nodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(_env *LnsEnv, arg1 *Nodes.Nodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(_env *LnsEnv, arg1 *Nodes.Nodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSetItem(_env *LnsEnv, arg1 *Nodes.Nodes_ExpSetItemNode, arg2 LnsAny)
    ProcessExpSetVal(_env *LnsEnv, arg1 *Nodes.Nodes_ExpSetValNode, arg2 LnsAny)
    ProcessExpSubDDD(_env *LnsEnv, arg1 *Nodes.Nodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(_env *LnsEnv, arg1 *Nodes.Nodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(_env *LnsEnv, arg1 *Nodes.Nodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessFor(_env *LnsEnv, arg1 *Nodes.Nodes_ForNode, arg2 LnsAny)
    ProcessForeach(_env *LnsEnv, arg1 *Nodes.Nodes_ForeachNode, arg2 LnsAny)
    ProcessForsort(_env *LnsEnv, arg1 *Nodes.Nodes_ForsortNode, arg2 LnsAny)
    ProcessGetField(_env *LnsEnv, arg1 *Nodes.Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(_env *LnsEnv, arg1 *Nodes.Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(_env *LnsEnv, arg1 *Nodes.Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(_env *LnsEnv, arg1 *Nodes.Nodes_ImportNode, arg2 LnsAny)
    ProcessLiteralArray(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralNil(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralNilNode, arg2 LnsAny)
    ProcessLiteralReal(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(_env *LnsEnv, arg1 *Nodes.Nodes_LiteralSymbolNode, arg2 LnsAny)
    ProcessLuneControl(_env *LnsEnv, arg1 *Nodes.Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(_env *LnsEnv, arg1 *Nodes.Nodes_LuneKindNode, arg2 LnsAny)
    ProcessMatch(_env *LnsEnv, arg1 *Nodes.Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(_env *LnsEnv, arg1 *Nodes.Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(_env *LnsEnv, arg1 *Nodes.Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(_env *LnsEnv, arg1 *Nodes.Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(_env *LnsEnv, arg1 *Nodes.Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(_env *LnsEnv, arg1 *Nodes.Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(_env *LnsEnv, arg1 *Nodes.Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(_env *LnsEnv, arg1 *Nodes.Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(_env *LnsEnv, arg1 *Nodes.Nodes_RepeatNode, arg2 LnsAny)
    ProcessReturn(_env *LnsEnv, arg1 *Nodes.Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(_env *LnsEnv, arg1 *Nodes.Nodes_RootNode, arg2 LnsAny)
    ProcessScope(_env *LnsEnv, arg1 *Nodes.Nodes_ScopeNode, arg2 LnsAny)
    ProcessShebang(_env *LnsEnv, arg1 *Nodes.Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(_env *LnsEnv, arg1 *Nodes.Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(_env *LnsEnv, arg1 *Nodes.Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(_env *LnsEnv, arg1 *Nodes.Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(_env *LnsEnv, arg1 *Nodes.Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(_env *LnsEnv, arg1 *Nodes.Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes.Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes.Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes.Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes.Nodes_WhileNode, arg2 LnsAny)
    registDeclSym(_env *LnsEnv, arg1 *LnsAst.Ast_SymbolInfo) LnsInt
    registerDecl(_env *LnsEnv, arg1 *Nodes.Nodes_NodeManager)
    registerRefs(_env *LnsEnv, arg1 *Nodes.Nodes_NodeManager)
    registerSymbol(_env *LnsEnv, arg1 *LnsAst.Ast_SymbolInfo)(LnsInt, bool)
    registerType(_env *LnsEnv, arg1 *LnsAst.Ast_TypeInfo)(LnsInt, bool)
}
type Analyze_tagFilter struct {
    Nodes.Nodes_Filter
    db *DBCtrl_DBCtrl
    streamName string
    type2nsid *LnsMap
    sym2nsid *LnsMap
    file2id *LnsMap
    moduleTypeInfo *LnsAst.Ast_TypeInfo
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
func NewAnalyze_tagFilter(_env *LnsEnv, arg1 *Nodes.Nodes_RootNode, arg2 *DBCtrl_DBCtrl, arg3 string) *Analyze_tagFilter {
    obj := &Analyze_tagFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitAnalyze_tagFilter(_env, arg1, arg2, arg3)
    return obj
}
// 31: decl @lns.@tags.@Analyze.tagFilter.addFileId
func (self *Analyze_tagFilter) addFileId(_env *LnsEnv, path string,mod string) LnsInt {
    __func__ := "@lns.@tags.@Analyze.tagFilter.addFileId"
    {
        __exp := self.file2id.Get(path)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            return _exp
        }
    }
    var fileId LnsInt
    fileId = Analyze_convExp95(Lns_2DDD(self.db.FP.AddFile(_env, path, mod)))
    Log_log(_env, Log_Level__Debug, __func__, 36, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("add file -- %d, %s", []LnsAny{fileId, path})
    }))
    
    self.file2id.Set(path,fileId)
    return fileId
}

// 41: decl @lns.@tags.@Analyze.tagFilter.getFileId
func (self *Analyze_tagFilter) getFileId(_env *LnsEnv, path string) LnsInt {
    {
        __exp := self.file2id.Get(path)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            return _exp
        }
    }
    var fileId LnsInt
    fileId = self.db.FP.GetFileIdFromPath(_env, path)
    self.file2id.Set(path,fileId)
    return fileId
}

// 51: DeclConstr
func (self *Analyze_tagFilter) InitAnalyze_tagFilter(_env *LnsEnv, rootNode *Nodes.Nodes_RootNode,db *DBCtrl_DBCtrl,streamName string) {
    self.InitNodes_Filter(_env, true, rootNode.FP.Get_moduleTypeInfo(_env), rootNode.FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env))
    self.moduleTypeInfo = rootNode.FP.Get_moduleTypeInfo(_env)
    
    self.file2id = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.db = db
    
    self.streamName = streamName
    
    self.type2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.sym2nsid = NewLnsMap( map[LnsAny]LnsAny{})
    
    var mod string
    mod = self.FP.GetFull(_env, rootNode.FP.Get_moduleTypeInfo(_env), false)
    self.FP.addFileId(_env, streamName, mod)
}

// 66: decl @lns.@tags.@Analyze.tagFilter.registerType
func (self *Analyze_tagFilter) registerType(_env *LnsEnv, typeInfo *LnsAst.Ast_TypeInfo)(LnsInt, bool) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerType"
    typeInfo = typeInfo.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env).FP.Get_genSrcTypeInfo(_env)
    
    {
        __exp := self.type2nsid.Get(typeInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            return _exp, false
        }
    }
    var parentNsId LnsInt
    parentNsId = Analyze_convExp368(Lns_2DDD(self.FP.registerType(_env, typeInfo.FP.Get_parentInfo(_env))))
    var name string
    name = self.FP.GetFull(_env, typeInfo, false)
    var nsId LnsInt
    var added bool
    nsId,added = self.db.FP.AddNamespace(_env, name, parentNsId)
    Log_log(_env, Log_Level__Debug, __func__, 74, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("%s %s %d", []LnsAny{name, added, nsId})
    }))
    
    self.type2nsid.Set(typeInfo,nsId)
    return nsId, added
}

// 79: decl @lns.@tags.@Analyze.tagFilter.registerSymbol
func (self *Analyze_tagFilter) registerSymbol(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo)(LnsInt, bool) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerSymbol"
    symbolInfo = symbolInfo.FP.GetOrg(_env)
    
    {
        __exp := self.sym2nsid.Get(symbolInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            return _exp, false
        }
    }
    var parentNsId LnsInt
    parentNsId = Analyze_convExp519(Lns_2DDD(self.FP.registerType(_env, symbolInfo.FP.Get_namespaceTypeInfo(_env))))
    var name string
    name = Ast_getFullNameSym(_env, &self.Nodes_Filter, symbolInfo)
    var nsId LnsInt
    var added bool
    nsId,added = self.db.FP.AddNamespace(_env, name, parentNsId)
    Log_log(_env, Log_Level__Debug, __func__, 88, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("%s %s %d", []LnsAny{name, added, nsId})
    }))
    
    self.sym2nsid.Set(symbolInfo,nsId)
    return nsId, added
}

// 93: decl @lns.@tags.@Analyze.tagFilter.registDeclSym
func (self *Analyze_tagFilter) registDeclSym(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo) LnsInt {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registDeclSym"
    var symNsId LnsInt
    symNsId = Analyze_convExp644(Lns_2DDD(self.FP.registerSymbol(_env, symbolInfo)))
    var pos *Types.Types_Position
    pos = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(symbolInfo.FP.Get_pos(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Types.Types_Position).FP.Get_orgPos(_env)}))).(*Types.Types_Position)
    var fileId LnsInt
    fileId = self.FP.getFileId(_env, pos.StreamName)
    self.db.FP.AddSymbolDecl(_env, symNsId, fileId, pos.LineNo, pos.Column)
    Log_log(_env, Log_Level__Debug, __func__, 98, Log_CreateMessage(func(_env *LnsEnv) string {
        return symbolInfo.FP.Get_name(_env)
    }))
    
    return symNsId
}

// 102: decl @lns.@tags.@Analyze.tagFilter.addSymbolRef
func (self *Analyze_tagFilter) addSymbolRef(_env *LnsEnv, mbrNsId LnsInt,pos *Types.Types_Position,setOp bool) {
    pos = pos.FP.Get_orgPos(_env)
    
    self.db.FP.AddSymbolRef(_env, mbrNsId, self.FP.getFileId(_env, pos.StreamName), pos.LineNo, pos.Column, setOp)
}

// 111: decl @lns.@tags.@Analyze.tagFilter.registerDecl
func (self *Analyze_tagFilter) registerDecl(_env *LnsEnv, nodeManager *Nodes.Nodes_NodeManager) {
    var registDeclTypeOp func(_env *LnsEnv, typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt
    registDeclTypeOp = func(_env *LnsEnv, typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt {
        var nsId LnsInt
        nsId = Analyze_convExp801(Lns_2DDD(self.FP.registerType(_env, typeInfo)))
        pos = pos.FP.Get_orgPos(_env)
        
        self.db.FP.AddSymbolDecl(_env, nsId, self.FP.getFileId(_env, pos.StreamName), pos.LineNo, pos.Column)
        return nsId
    }
    var registDeclType func(_env *LnsEnv, workNode *Nodes.Nodes_Node) LnsInt
    registDeclType = func(_env *LnsEnv, workNode *Nodes.Nodes_Node) LnsInt {
        return registDeclTypeOp(_env, workNode.FP.Get_expType(_env), workNode.FP.Get_pos(_env))
    }
    var registerAutoMethod func(_env *LnsEnv, parentInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position)
    registerAutoMethod = func(_env *LnsEnv, parentInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) {
        for _, _symbolInfo := range( Lns_unwrap( _env.NilAccFin(_env.NilAccPush(parentInfo.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*LnsAst.Ast_Scope).FP.Get_symbol2SymbolInfoMap(_env)}))).(*LnsMap).Items ) {
            symbolInfo := _symbolInfo.(LnsAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var typeInfo *LnsAst.Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo(_env)
            if typeInfo.FP.Get_autoFlag(_env){
                registDeclTypeOp(_env, typeInfo, pos)
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclFormNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
        registDeclType(_env, &workNode.Nodes_Node)
        for _, _altType := range( workNode.FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).Items ) {
            altType := _altType.(LnsAst.Ast_TypeInfoDownCast).ToAst_TypeInfo()
            registDeclTypeOp(_env, altType, workNode.FP.Get_pos(_env))
        }
    }
    for _, _workNode := range( nodeManager.FP.GetProtoClassNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoClassNodeDownCast).ToNodes_ProtoClassNode()
        registDeclType(_env, &workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetDeclClassNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        registDeclType(_env, &workNode.Nodes_Node)
        for _, _altType := range( workNode.FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).Items ) {
            altType := _altType.(LnsAst.Ast_TypeInfoDownCast).ToAst_TypeInfo()
            registDeclTypeOp(_env, altType, workNode.FP.Get_pos(_env))
        }
        registerAutoMethod(_env, workNode.FP.Get_expType(_env), workNode.FP.Get_pos(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetDeclConstrNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclConstrNodeDownCast).ToNodes_DeclConstrNode()
        registDeclType(_env, &workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetDeclFuncNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
        registDeclType(_env, &workNode.Nodes_Node)
        for _, _altType := range( workNode.FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).Items ) {
            altType := _altType.(LnsAst.Ast_TypeInfoDownCast).ToAst_TypeInfo()
            registDeclTypeOp(_env, altType, workNode.FP.Get_pos(_env))
        }
    }
    for _, _workNode := range( nodeManager.FP.GetProtoMethodNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ProtoMethodNodeDownCast).ToNodes_ProtoMethodNode()
        registDeclType(_env, &workNode.Nodes_Node)
    }
    var addMethod func(_env *LnsEnv, methodType *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt
    addMethod = func(_env *LnsEnv, methodType *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) LnsInt {
        var nsId LnsInt
        nsId = registDeclTypeOp(_env, methodType, pos)
        if Lns_op_not(methodType.FP.Get_staticFlag(_env)){
            {
                _scope := methodType.FP.Get_parentInfo(_env).FP.Get_scope(_env)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*LnsAst.Ast_Scope)
                    scope.FP.FilterTypeInfoFieldAndIF(_env, false, scope, LnsAst.Ast_ScopeAccess__Normal, LnsAst.Ast_filterForm(func(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo) bool {
                        if symbolInfo.FP.Get_name(_env) == methodType.FP.Get_rawTxt(_env){
                            var superNsId LnsInt
                            superNsId = Analyze_convExp1118(Lns_2DDD(self.FP.registerType(_env, symbolInfo.FP.Get_typeInfo(_env))))
                            self.db.FP.AddOverride(_env, nsId, superNsId)
                        }
                        return true
                    }))
                }
            }
        }
        for _, _altType := range( methodType.FP.Get_itemTypeInfoList(_env).Items ) {
            altType := _altType.(LnsAst.Ast_TypeInfoDownCast).ToAst_TypeInfo()
            registDeclTypeOp(_env, altType, pos)
        }
        return nsId
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMethodNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
        addMethod(_env, workNode.FP.Get_expType(_env), workNode.FP.Get_pos(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetDeclEnumNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
        registDeclType(_env, &workNode.Nodes_Node)
        registerAutoMethod(_env, workNode.FP.Get_expType(_env), workNode.FP.Get_pos(_env))
        for _, _name := range( workNode.FP.Get_valueNameList(_env).Items ) {
            name := _name.(Types.Types_TokenDownCast).ToTypes_Token()
            {
                __exp := workNode.FP.Get_scope(_env).FP.GetSymbolInfoChild(_env, name.Txt)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                    self.FP.registDeclSym(_env, _exp)
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclAlgeNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
        registDeclType(_env, &workNode.Nodes_Node)
        {
            __collection1274 := workNode.FP.Get_algeType(_env).FP.Get_valInfoMap(_env)
            __sorted1274 := __collection1274.CreateKeyListStr()
            __sorted1274.Sort( LnsItemKindStr, nil )
            for _, _name := range( __sorted1274.Items ) {
                name := _name.(string)
                {
                    __exp := _env.NilAccFin(_env.NilAccPush(workNode.FP.Get_algeType(_env).FP.Get_scope(_env)) && 
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*LnsAst.Ast_Scope).FP.GetSymbolInfoChild(_env, name)})/* 203:14 */)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                        self.FP.registDeclSym(_env, _exp)
                    }
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMacroNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMacroNodeDownCast).ToNodes_DeclMacroNode()
        registDeclType(_env, &workNode.Nodes_Node)
    }
    for _, _workNode := range( nodeManager.FP.GetAliasNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        registDeclTypeOp(_env, workNode.FP.Get_expType(_env), Lns_unwrap( workNode.FP.Get_newSymbol(_env).FP.Get_pos(_env)).(*Types.Types_Position))
    }
    for _, _workNode := range( nodeManager.FP.GetImportNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        self.FP.registDeclSym(_env, workNode.FP.Get_symbolInfo(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetDeclVarNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
        for _, _symbolInfo := range( workNode.FP.Get_symbolInfoList(_env).Items ) {
            symbolInfo := _symbolInfo.(LnsAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbolInfo.FP.Get_scope(_env) == self.moduleTypeInfo.FP.Get_scope(_env){
                self.FP.registDeclSym(_env, symbolInfo)
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetDeclMemberNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var mbrNsId LnsInt
        mbrNsId = self.FP.registDeclSym(_env, workNode.FP.Get_symbolInfo(_env))
        if workNode.FP.Get_getterMode(_env) != LnsAst.Ast_AccessMode__None{
            var name string
            name = _env.LuaVM.String_format("get_%s", []LnsAny{workNode.FP.Get_name(_env).Txt})
            var pos *Types.Types_Position
            pos = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(workNode.FP.Get_getterToken(_env)) && 
            _env.NilAccPush(_env.NilAccPop().(*Types.Types_Token).Pos))).(*Types.Types_Position)
            addMethod(_env, Lns_unwrap( _env.NilAccFin(_env.NilAccPush(workNode.FP.Get_classType(_env).FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*LnsAst.Ast_Scope).FP.GetTypeInfoChild(_env, name)})/* 235:20 */)).(*LnsAst.Ast_TypeInfo), pos)
            self.FP.addSymbolRef(_env, mbrNsId, pos, true)
        }
        if workNode.FP.Get_setterMode(_env) != LnsAst.Ast_AccessMode__None{
            var name string
            name = _env.LuaVM.String_format("set_%s", []LnsAny{workNode.FP.Get_name(_env).Txt})
            var pos *Types.Types_Position
            pos = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(workNode.FP.Get_setterToken(_env)) && 
            _env.NilAccPush(_env.NilAccPop().(*Types.Types_Token).Pos))).(*Types.Types_Position)
            addMethod(_env, Lns_unwrap( _env.NilAccFin(_env.NilAccPush(workNode.FP.Get_classType(_env).FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*LnsAst.Ast_Scope).FP.GetTypeInfoChild(_env, name)})/* 243:20 */)).(*LnsAst.Ast_TypeInfo), pos)
            self.FP.addSymbolRef(_env, mbrNsId, pos, false)
        }
    }
}

// 250: decl @lns.@tags.@Analyze.tagFilter.registerRefs
func (self *Analyze_tagFilter) registerRefs(_env *LnsEnv, nodeManager *Nodes.Nodes_NodeManager) {
    __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs"
    var addSymbolRef func(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position,setOp bool)
    addSymbolRef = func(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position,setOp bool) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.addSymbolRef"
        if _switch1522 := symbolInfo.FP.Get_name(_env); _switch1522 == "__func__" || _switch1522 == "__mod__" || _switch1522 == "__line__" {
            return 
        }
        var nsId LnsInt
        var added bool
        nsId,added = self.FP.registerSymbol(_env, symbolInfo)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( added) &&
            _env.SetStackVal( Lns_op_not(LnsAst.Ast_isBuiltin(_env, symbolInfo.FP.Get_namespaceTypeInfo(_env).FP.Get_typeId(_env).Id))) ).(bool)){
            Log_log(_env, Log_Level__Err, __func__, 260, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("no register sym -- %d:%d:%s", []LnsAny{pos.LineNo, pos.Column, Ast_getFullNameSym(_env, &self.Nodes_Filter, symbolInfo)})
            }))
            
        }
        self.FP.addSymbolRef(_env, nsId, pos, setOp)
    }
    var registerRefType func(_env *LnsEnv, typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position)
    registerRefType = func(_env *LnsEnv, typeInfo *LnsAst.Ast_TypeInfo,pos *Types.Types_Position) {
        __func__ := "@lns.@tags.@Analyze.tagFilter.registerRefs.registerRefType"
        var nsId LnsInt
        var added bool
        nsId,added = self.FP.registerType(_env, typeInfo)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( added) &&
            _env.SetStackVal( Lns_op_not(LnsAst.Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id))) ).(bool)){
            Log_log(_env, Log_Level__Err, __func__, 269, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("no register type -- %d:%d:%s", []LnsAny{pos.LineNo, pos.Column, self.FP.GetFull(_env, typeInfo, false)})
            }))
            
        }
        self.FP.addSymbolRef(_env, nsId, pos, true)
    }
    var registerRefSym func(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position,setOp bool)
    registerRefSym = func(_env *LnsEnv, symbolInfo *LnsAst.Ast_SymbolInfo,pos *Types.Types_Position,setOp bool) {
        if _switch1919 := symbolInfo.FP.Get_namespaceTypeInfo(_env).FP.Get_kind(_env); _switch1919 == LnsAst.Ast_TypeInfoKind__Enum || _switch1919 == LnsAst.Ast_TypeInfoKind__Alge {
            addSymbolRef(_env, symbolInfo, pos, true)
        } else {
            if _switch1917 := symbolInfo.FP.Get_kind(_env); _switch1917 == LnsAst.Ast_SymbolKind__Fun || _switch1917 == LnsAst.Ast_SymbolKind__Mac || _switch1917 == LnsAst.Ast_SymbolKind__Mtd {
                registerRefType(_env, symbolInfo.FP.Get_typeInfo(_env), pos)
            } else if _switch1917 == LnsAst.Ast_SymbolKind__Typ {
                if symbolInfo.FP.Get_typeInfo(_env).FP.Get_kind(_env) == LnsAst.Ast_TypeInfoKind__Module{
                    addSymbolRef(_env, symbolInfo, pos, true)
                }
                registerRefType(_env, symbolInfo.FP.Get_typeInfo(_env), pos)
            } else if _switch1917 == LnsAst.Ast_SymbolKind__Mbr {
                addSymbolRef(_env, symbolInfo, pos, setOp)
            } else if _switch1917 == LnsAst.Ast_SymbolKind__Var {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( LnsAst.Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env))) ||
                    _env.SetStackVal( symbolInfo.FP.Get_scope(_env) == self.moduleTypeInfo.FP.Get_scope(_env)) ).(bool){
                    addSymbolRef(_env, symbolInfo, pos, setOp)
                }
            } else if _switch1917 == LnsAst.Ast_SymbolKind__Arg {
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetExpSetValNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpSetValNodeDownCast).ToNodes_ExpSetValNode()
        for _, _leftSym := range( workNode.FP.Get_LeftSymList(_env).Items ) {
            leftSym := _leftSym.(LnsAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if _switch1984 := leftSym.FP.Get_kind(_env); _switch1984 == LnsAst.Ast_SymbolKind__Mbr {
                registerRefSym(_env, leftSym, workNode.FP.Get_pos(_env), true)
            } else if _switch1984 == LnsAst.Ast_SymbolKind__Var {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( LnsAst.Ast_isPubToExternal(_env, leftSym.FP.Get_accessMode(_env))) ||
                    _env.SetStackVal( leftSym.FP.Get_scope(_env) == self.moduleTypeInfo.FP.Get_scope(_env)) ).(bool){
                    registerRefSym(_env, leftSym, workNode.FP.Get_pos(_env), true)
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetExpNewNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpNewNodeDownCast).ToNodes_ExpNewNode()
        registerRefType(_env, workNode.FP.Get_ctorTypeInfo(_env), workNode.FP.Get_pos(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetExpCallSuperCtorNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpCallSuperCtorNodeDownCast).ToNodes_ExpCallSuperCtorNode()
        registerRefType(_env, workNode.FP.Get_methodType(_env), workNode.FP.Get_pos(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetExpCallSuperNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpCallSuperNodeDownCast).ToNodes_ExpCallSuperNode()
        registerRefType(_env, workNode.FP.Get_methodType(_env), workNode.FP.Get_pos(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetExpRefNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpRefNodeDownCast).ToNodes_ExpRefNode()
        registerRefSym(_env, workNode.FP.Get_symbolInfo(_env), workNode.FP.Get_pos(_env), false)
    }
    for _, _workNode := range( nodeManager.FP.GetRefFieldNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_RefFieldNodeDownCast).ToNodes_RefFieldNode()
        {
            __exp := workNode.FP.Get_symbolInfo(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                registerRefSym(_env, _exp, workNode.FP.Get_pos(_env), false)
            } else {
                if Lns_op_not(workNode.FP.Get_macroArgFlag(_env)){
                    Log_log(_env, Log_Level__Warn, __func__, 354, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.LuaVM.String_format("no symbolInfo -- %s, %s:%d:%d", []LnsAny{workNode.FP.Get_field(_env).Txt, workNode.FP.Get_pos(_env).StreamName, workNode.FP.Get_pos(_env).LineNo, workNode.FP.Get_pos(_env).Column})
                    }))
                    
                }
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetGetFieldNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_GetFieldNodeDownCast).ToNodes_GetFieldNode()
        registerRefType(_env, workNode.FP.Get_getterTypeInfo(_env), workNode.FP.Get_pos(_env))
    }
    for _, _workNode := range( nodeManager.FP.GetExpOmitEnumNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_ExpOmitEnumNodeDownCast).ToNodes_ExpOmitEnumNode()
        {
            __exp := _env.NilAccFin(_env.NilAccPush(workNode.FP.Get_enumTypeInfo(_env).FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*LnsAst.Ast_Scope).FP.GetSymbolInfoChild(_env, workNode.FP.Get_valInfo(_env).FP.Get_name(_env))})/* 364:11 */)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                registerRefSym(_env, _exp, workNode.FP.Get_pos(_env), false)
            }
        }
    }
    for _, _workNode := range( nodeManager.FP.GetMatchNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_MatchNodeDownCast).ToNodes_MatchNode()
        for _, _matchCase := range( workNode.FP.Get_caseList(_env).Items ) {
            matchCase := _matchCase.(Nodes.Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            var valInfo *LnsAst.Ast_AlgeValInfo
            valInfo = matchCase.FP.Get_valInfo(_env)
            {
                __exp := _env.NilAccFin(_env.NilAccPush(valInfo.FP.Get_algeTpye(_env).FP.Get_scope(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*LnsAst.Ast_Scope).FP.GetSymbolInfoChild(_env, valInfo.FP.Get_name(_env))})/* 372:14 */)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*LnsAst.Ast_SymbolInfo)
                    registerRefSym(_env, _exp, matchCase.FP.Get_block(_env).FP.Get_pos(_env), false)
                }
            }
        }
    }
}

// 379: decl @lns.@tags.@Analyze.tagFilter.processRoot
func (self *Analyze_tagFilter) ProcessRoot(_env *LnsEnv, node *Nodes.Nodes_RootNode,_opt LnsAny) {
    self.type2nsid.Set(LnsAst.Ast_headTypeInfo,DBCtrl_rootNsId)
    var modNsId LnsInt
    modNsId = Analyze_convExp2302(Lns_2DDD(self.FP.registerType(_env, node.FP.Get_moduleTypeInfo(_env))))
    var fileId LnsInt
    fileId = self.FP.getFileId(_env, self.streamName)
    self.db.FP.AddSymbolDecl(_env, modNsId, fileId, 1, 1)
    var projDir string
    projDir = DBCtrl_getProjDir(_env, self.streamName, self.FP.GetFull(_env, node.FP.Get_moduleTypeInfo(_env), false))
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetSubfileNodeList(_env).Items ) {
        workNode := _workNode.(Nodes.Nodes_SubfileNodeDownCast).ToNodes_SubfileNode()
        {
            _usePath := workNode.FP.Get_usePath(_env)
            if !Lns_IsNil( _usePath ) {
                usePath := _usePath.(string)
                var subPath string
                subPath = LnsUtil.Util_pathJoin(_env, projDir, Lns_car(_env.LuaVM.String_gsub(usePath,"%.", "/")).(string) + ".lns")
                var subId LnsInt
                subId = self.FP.addFileId(_env, subPath, usePath)
                self.db.FP.AddSubfile(_env, fileId, subId)
            }
        }
    }
    self.FP.registerDecl(_env, node.FP.Get_nodeManager(_env))
    self.FP.registerRefs(_env, node.FP.Get_nodeManager(_env))
}


func Lns_Analyze_init(_env *LnsEnv) {
    if init_Analyze { return }
    init_Analyze = true
    Analyze__mod__ = "@lns.@tags.@Analyze"
    Lns_InitMod()
    Lns_DBCtrl_init(_env)
    Lns_Option_init(_env)
    Lns_Log_init(_env)
    Lns_Ast_init(_env)
    LnsOpt.Lns_Option_init(_env)
    Nodes.Lns_Nodes_init(_env)
    Parser.Lns_Parser_init(_env)
    TransUnit.Lns_TransUnit_init(_env)
    front.Lns_front_init(_env)
    LnsAst.Lns_Ast_init(_env)
    Types.Lns_Types_init(_env)
    LuaVer.Lns_LuaVer_init(_env)
    LnsLog.Lns_Log_init(_env)
    LnsUtil.Lns_Util_init(_env)
}
func init() {
    init_Analyze = false
}
