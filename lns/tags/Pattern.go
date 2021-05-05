// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import Parser "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
import LnsAst "github.com/ifritJP/LuneScript/src/lune/base"
import LuaVer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsLog "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Pattern bool
var Pattern__mod__ string
// for 287
func Pattern_convExp1979(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}








// 270: decl @lns.@tags.@Pattern.getPatterAt
func Pattern_getPatterAt(db *DBCtrl_DBCtrl,analyzeFileInfo *Option_AnalyzeFileInfo,inqMod string) LnsAny {
    var fileId LnsInt
    fileId = db.FP.GetFileIdFromPath(analyzeFileInfo.FP.Get_path())
    var path string
    
    {
        _path := db.FP.GetMainFilePath(fileId)
        if _path == nil{
            path = analyzeFileInfo.FP.Get_path()
            
        } else {
            path = _path.(string)
        }
    }
    var pattern LnsAny
    pattern = nil
    var useStdInMod LnsAny
    var projDir LnsAny
    projDir = nil
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(path,"^%.%.", nil, nil))) ||
        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(path,"^/", nil, nil))) )){
        var dir string
        dir = Pattern_convExp1979(Lns_2DDD(Lns_getVM().String_gsub(path,"/[^/]+$", "")))
        projDir = LnsUtil.Util_searchProjDir(dir)
        
    }
    if analyzeFileInfo.FP.Get_stdinFlag(){
        useStdInMod = LnsUtil.Util_scriptPath2ModuleFromProjDir(analyzeFileInfo.FP.Get_path(), projDir)
        
    } else { 
        useStdInMod = nil
        
    }
    Ast_buildAst(LnsLog.Log_Level__Err, NewLnsList([]LnsAny{path}), projDir, useStdInMod, false, front.Front_AstCallback(func(ast *TransUnit.TransUnit_ASTInfo) {
        __func__ := "@lns.@tags.@Pattern.getPatterAt.<anonymous>"
        if ast.FP.Get_streamName() == path{
            var filter *Pattern_SyntaxFilter
            filter = NewPattern_SyntaxFilter(ast)
            pattern = filter.FP.GetPattern(path, analyzeFileInfo, inqMod)
            
            Log_log(Log_Level__Log, __func__, 303, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("pattern -- %s", []LnsAny{pattern})
            }))
            
        }
    }))
    return pattern
}

// declaration Class -- SyntaxFilter
type Pattern_SyntaxFilterMtd interface {
    DefaultProcess(arg1 *Nodes.Nodes_Node, arg2 LnsAny)
    GetFull(arg1 *LnsAst.Ast_TypeInfo, arg2 bool) string
    GetPattern(arg1 string, arg2 *Option_AnalyzeFileInfo, arg3 string) LnsAny
    getPatternFromNode(arg1 *Option_AnalyzeFileInfo, arg2 string, arg3 *Nodes.Nodes_Node) LnsAny
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
}
type Pattern_SyntaxFilter struct {
    Nodes.Nodes_Filter
    ast *TransUnit.TransUnit_ASTInfo
    FP Pattern_SyntaxFilterMtd
}
func Pattern_SyntaxFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Pattern_SyntaxFilter).FP
}
type Pattern_SyntaxFilterDownCast interface {
    ToPattern_SyntaxFilter() *Pattern_SyntaxFilter
}
func Pattern_SyntaxFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Pattern_SyntaxFilterDownCast)
    if ok { return work.ToPattern_SyntaxFilter() }
    return nil
}
func (obj *Pattern_SyntaxFilter) ToPattern_SyntaxFilter() *Pattern_SyntaxFilter {
    return obj
}
func NewPattern_SyntaxFilter(arg1 *TransUnit.TransUnit_ASTInfo) *Pattern_SyntaxFilter {
    obj := &Pattern_SyntaxFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitPattern_SyntaxFilter(arg1)
    return obj
}
// 27: DeclConstr
func (self *Pattern_SyntaxFilter) InitPattern_SyntaxFilter(ast *TransUnit.TransUnit_ASTInfo) {
    self.InitNodes_Filter(true, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.ast = ast
    
}

// 33: decl @lns.@tags.@Pattern.SyntaxFilter.getPatternFromNode
func (self *Pattern_SyntaxFilter) getPatternFromNode(analyzeFileInfo *Option_AnalyzeFileInfo,inqMod string,nearest *Nodes.Nodes_Node) LnsAny {
    __func__ := "@lns.@tags.@Pattern.SyntaxFilter.getPatternFromNode"
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
    Log_log(Log_Level__Log, __func__, 19, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s %s:%4d:%3d -- %s", []LnsAny{"nearestNode -- ", nearest.FP.Get_effectivePos().StreamName, nearest.FP.Get_effectivePos().LineNo, nearest.FP.Get_effectivePos().Column, Nodes.Nodes_getNodeKindName(nearest.FP.Get_kind())})
    }))
    
    
    {
        _workNode := Nodes.Nodes_ImportNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ImportNode)
            if _switch297 := inqMod; _switch297 == Option_InqMode__Def {
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            } else if _switch297 == Option_InqMode__Ref {
                return Ast_getFullNameSym(&self.Nodes_Filter, workNode.FP.Get_symbolInfo())
            } else if _switch297 == Option_InqMode__Set {
            }
            return nil
        }
    }
    {
        _workNode := Nodes.Nodes_ExpRefNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpRefNode)
            return Ast_getFullNameSym(&self.Nodes_Filter, workNode.FP.Get_symbolInfo())
        }
    }
    {
        _workNode := Nodes.Nodes_RefFieldNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_RefFieldNode)
            if isInner(workNode.FP.Get_field().Pos, workNode.FP.Get_field().Txt){
                {
                    _symbolInfo := workNode.FP.Get_symbolInfo()
                    if !Lns_IsNil( _symbolInfo ) {
                        symbolInfo := _symbolInfo.(*LnsAst.Ast_SymbolInfo)
                        return Ast_getFullNameSym(&self.Nodes_Filter, symbolInfo)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclVarNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclVarNode)
            for _, _varSym := range( workNode.FP.Get_symbolInfoList().Items ) {
                varSym := _varSym.(LnsAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if isInner(Lns_unwrap( varSym.FP.Get_pos()).(*Types.Types_Position), varSym.FP.Get_name()){
                    return Ast_getFullNameSym(&self.Nodes_Filter, varSym)
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpOmitEnumNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpOmitEnumNode)
            if isInner(workNode.FP.Get_effectivePos(), workNode.FP.Get_valInfo().FP.Get_name()){
                return Ast_getFullNameSym(&self.Nodes_Filter, workNode.FP.Get_valInfo().FP.Get_symbolInfo())
            }
        }
    }
    {
        _workNode := Nodes.Nodes_NewAlgeValNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_NewAlgeValNode)
            if isInner(workNode.FP.Get_effectivePos(), workNode.FP.Get_valInfo().FP.Get_name()){
                return Ast_getFullNameSym(&self.Nodes_Filter, workNode.FP.Get_valInfo().FP.Get_symbolInfo())
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpMacroExpNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpMacroExpNode)
            if isInner(workNode.FP.Get_effectivePos(), workNode.FP.Get_expType().FP.Get_rawTxt()){
                return self.FP.GetFull(workNode.FP.Get_macroType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclFuncNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclFuncNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if !Lns_IsNil( _name ) {
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
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclEnumNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_name()
            if isInner(name.Pos, name.Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
            for _, _valInfo := range( workNode.FP.Get_enumType().FP.Get_name2EnumValInfo().Items ) {
                valInfo := _valInfo.(LnsAst.Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                if isInner(Lns_unwrap( valInfo.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position), valInfo.FP.Get_symbolInfo().FP.Get_name()){
                    return Ast_getFullNameSym(&self.Nodes_Filter, valInfo.FP.Get_symbolInfo())
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclAlgeNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclAlgeNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_name()
            if isInner(name.Pos, name.Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
            for _, _valInfo := range( workNode.FP.Get_algeType().FP.Get_valInfoMap().Items ) {
                valInfo := _valInfo.(LnsAst.Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                if isInner(Lns_unwrap( valInfo.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position), valInfo.FP.Get_symbolInfo().FP.Get_name()){
                    return Ast_getFullNameSym(&self.Nodes_Filter, valInfo.FP.Get_symbolInfo())
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclClassNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclClassNode)
            if isInner(workNode.FP.Get_name().Pos, workNode.FP.Get_name().Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMethodNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclMethodNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if !Lns_IsNil( _name ) {
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
        if !Lns_IsNil( _workNode ) {
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
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclMemberNode)
            if isInner(Lns_unwrap( workNode.FP.Get_symbolInfo().FP.Get_pos()).(*Types.Types_Position), workNode.FP.Get_symbolInfo().FP.Get_name()){
                return Ast_getFullNameSym(&self.Nodes_Filter, workNode.FP.Get_symbolInfo())
            }
            
            {
                _token, _symbol := workNode.FP.Get_getterToken(), workNode.FP.GetGetterSym()
                if !Lns_IsNil( _token ) && !Lns_IsNil( _symbol ) {
                    token := _token.(*Types.Types_Token)
                    symbol := _symbol.(*LnsAst.Ast_SymbolInfo)
                    if isInner(token.Pos, token.Txt){
                        return Ast_getFullNameSym(&self.Nodes_Filter, symbol)
                    }
                }
            }
            {
                _token, _symbol := workNode.FP.Get_setterToken(), workNode.FP.GetSetterSym()
                if !Lns_IsNil( _token ) && !Lns_IsNil( _symbol ) {
                    token := _token.(*Types.Types_Token)
                    symbol := _symbol.(*LnsAst.Ast_SymbolInfo)
                    if isInner(token.Pos, token.Txt){
                        return Ast_getFullNameSym(&self.Nodes_Filter, symbol)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclConstrNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclConstrNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if !Lns_IsNil( _name ) {
                    name := _name.(*Types.Types_Token)
                    if isInner(name.Pos, name.Txt){
                        return self.FP.GetFull(workNode.FP.Get_expType(), false)
                    }
                }
            }
            
        }
    }
    {
        _workNode := Nodes.Nodes_ExpCallSuperCtorNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpCallSuperCtorNode)
            return self.FP.GetFull(workNode.FP.Get_methodType(), false)
        }
    }
    {
        _workNode := Nodes.Nodes_ExpCallSuperNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpCallSuperNode)
            return self.FP.GetFull(workNode.FP.Get_methodType(), false)
        }
    }
    {
        _workNode := Nodes.Nodes_ExpNewNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpNewNode)
            if isInner(workNode.FP.Get_pos(), "new"){
                return self.FP.GetFull(workNode.FP.Get_ctorTypeInfo(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMacroNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclMacroNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_declInfo().FP.Get_name()
            if isInner(name.Pos, name.Txt){
                return self.FP.GetFull(workNode.FP.Get_expType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpMacroExpNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpMacroExpNode)
            if isInner(workNode.FP.Get_pos(), workNode.FP.Get_macroType().FP.Get_rawTxt()){
                return self.FP.GetFull(workNode.FP.Get_macroType(), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_GetFieldNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_GetFieldNode)
            {
                _symbolInfo := workNode.FP.Get_symbolInfo()
                if !Lns_IsNil( _symbolInfo ) {
                    symbolInfo := _symbolInfo.(*LnsAst.Ast_SymbolInfo)
                    if isInner(workNode.FP.Get_field().Pos, workNode.FP.Get_field().Txt){
                        return Ast_getFullNameSym(&self.Nodes_Filter, symbolInfo)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_AliasNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_AliasNode)
            if isInner(Lns_unwrap( workNode.FP.Get_newSymbol().FP.Get_pos()).(*Types.Types_Position), workNode.FP.Get_newSymbol().FP.Get_name()){
                return Ast_getFullNameSym(&self.Nodes_Filter, workNode.FP.Get_newSymbol())
            }
            
        }
    }
    {
        _workNode := Nodes.Nodes_DeclFormNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclFormNode)
            {
                _name := workNode.FP.Get_declInfo().FP.Get_name()
                if !Lns_IsNil( _name ) {
                    name := _name.(*Types.Types_Token)
                    if isInner(name.Pos, name.Txt){
                        return self.FP.GetFull(workNode.FP.Get_expType(), false)
                    }
                }
            }
            
        }
    }
    Log_log(Log_Level__Err, __func__, 191, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("unknown pattern -- %s", []LnsAny{Nodes.Nodes_getNodeKindName(nearest.FP.Get_kind())})
    }))
    
    return nil
}

// 195: decl @lns.@tags.@Pattern.SyntaxFilter.getPattern
func (self *Pattern_SyntaxFilter) GetPattern(path string,analyzeFileInfo *Option_AnalyzeFileInfo,inqMod string) LnsAny {
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
    var pattern LnsAny
    pattern = nil
    if self.ast.FP.Get_streamName() == path{
        var nearestNode LnsAny
        nearestNode = nil
        self.ast.FP.Get_node().FP.Visit(Nodes.Nodes_NodeVisitor(func(node *Nodes.Nodes_Node,parent *Nodes.Nodes_Node,relation string,depth LnsInt) LnsInt {
            __func__ := "@lns.@tags.@Pattern.SyntaxFilter.getPattern.<anonymous>"
            if analyzeFileInfo.FP.Get_path() == node.FP.Get_effectivePos().StreamName{
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( analyzeFileInfo.FP.Get_lineNo() == node.FP.Get_effectivePos().LineNo) &&
                    Lns_GetEnv().SetStackVal( analyzeFileInfo.FP.Get_column() >= node.FP.Get_effectivePos().Column) ).(bool)){
                    {
                        _declMemberNode := Nodes.Nodes_DeclMemberNodeDownCastF(node.FP)
                        if !Lns_IsNil( _declMemberNode ) {
                            declMemberNode := _declMemberNode.(*Nodes.Nodes_DeclMemberNode)
                            {
                                _token := declMemberNode.FP.Get_getterToken()
                                if !Lns_IsNil( _token ) {
                                    token := _token.(*Types.Types_Token)
                                    if isInner(token.Pos, token.Txt){
                                        nearestNode = node
                                        
                                        return Nodes.Nodes_NodeVisitMode__End
                                    }
                                }
                            }
                            {
                                _token := declMemberNode.FP.Get_setterToken()
                                if !Lns_IsNil( _token ) {
                                    token := _token.(*Types.Types_Token)
                                    if isInner(token.Pos, token.Txt){
                                        nearestNode = node
                                        
                                        return Nodes.Nodes_NodeVisitMode__End
                                    }
                                }
                            }
                        }
                    }
                    {
                        _nearest := nearestNode
                        if !Lns_IsNil( _nearest ) {
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
                        if !Lns_IsNil( _nearest ) {
                            nearest := _nearest.(*Nodes.Nodes_Node)
                            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( nearest.FP.Get_effectivePos().LineNo < node.FP.Get_effectivePos().LineNo) &&
                                Lns_GetEnv().SetStackVal( node.FP.Get_effectivePos().LineNo < analyzeFileInfo.FP.Get_lineNo()) ).(bool)){
                                nearestNode = node
                                
                            }
                        } else {
                            nearestNode = node
                            
                        }
                    }
                }
                Log_log(Log_Level__Trace, __func__, 19, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("%s %s:%4d:%3d -- %s", []LnsAny{"visit:", node.FP.Get_effectivePos().StreamName, node.FP.Get_effectivePos().LineNo, node.FP.Get_effectivePos().Column, Nodes.Nodes_getNodeKindName(node.FP.Get_kind())})
                }))
                
                
                return Nodes.Nodes_NodeVisitMode__Child
            }
            return Nodes.Nodes_NodeVisitMode__Next
        }), 0, NewLnsSet([]LnsAny{}))
        {
            _nearest := nearestNode
            if !Lns_IsNil( _nearest ) {
                nearest := _nearest.(*Nodes.Nodes_Node)
                pattern = self.FP.getPatternFromNode(analyzeFileInfo, inqMod, nearest)
                
            }
        }
    }
    return pattern
}


func Lns_Pattern_init() {
    if init_Pattern { return }
    init_Pattern = true
    Pattern__mod__ = "@lns.@tags.@Pattern"
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
    Types.Lns_Types_init()
    LnsAst.Lns_Ast_init()
    LuaVer.Lns_LuaVer_init()
    LnsLog.Lns_Log_init()
    LnsUtil.Lns_Util_init()
}
func init() {
    init_Pattern = false
}
