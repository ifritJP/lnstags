// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LuneOpt "github.com/ifritJP/LuneScript/src/lune/base"
import Nodes "github.com/ifritJP/LuneScript/src/lune/base"
import Parser "github.com/ifritJP/LuneScript/src/lune/base"
import TransUnit "github.com/ifritJP/LuneScript/src/lune/base"
import front "github.com/ifritJP/LuneScript/src/lune/base"
import Types "github.com/ifritJP/LuneScript/src/lune/base"
import LuneAst "github.com/ifritJP/LuneScript/src/lune/base"
import AstInfo "github.com/ifritJP/LuneScript/src/lune/base"
import LuaVer "github.com/ifritJP/LuneScript/src/lune/base"
import LuneLog "github.com/ifritJP/LuneScript/src/lune/base"
import LuneUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_Pattern bool
var Pattern__mod__ string
// for 288
func Pattern_convExp0_1986(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}



// 272: decl @lns.@tags.@Pattern.getPatterAt
func Pattern_getPatterAt(_env *LnsEnv, db *DBCtrl_DBCtrl,analyzeFileInfo *Option_AnalyzeFileInfo,inqMod string,transCtrlInfo *Types.Types_TransCtrlInfo) LnsAny {
    var fileId LnsInt
    fileId = db.FP.GetFileIdFromPath(_env, analyzeFileInfo.FP.Get_path(_env))
    var path string
    
    {
        _path := db.FP.GetMainFilePath(_env, fileId)
        if _path == nil{
            path = analyzeFileInfo.FP.Get_path(_env)
        } else {
            path = _path.(string)
        }
    }
    var pattern LnsAny
    pattern = nil
    var projDir LnsAny
    projDir = nil
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_car(_env.GetVM().String_find(path,"^%.%.", nil, nil))) ||
        _env.SetStackVal( Lns_car(_env.GetVM().String_find(path,"^/", nil, nil))) )){
        var dir string
        dir = Pattern_convExp0_1986(Lns_2DDD(_env.GetVM().String_gsub(path,"/[^/]+$", "")))
        projDir = LuneUtil.Util_searchProjDir(_env, dir)
    }
    Ast_buildAst(_env, LuneLog.Log_Level__Err, NewLnsList([]LnsAny{path}), projDir, analyzeFileInfo.FP.Get_stdinFile(_env), false, transCtrlInfo, front.Front_AstCallback(func(_env *LnsEnv, ast *AstInfo.AstInfo_ASTInfo) {
        __func__ := "@lns.@tags.@Pattern.getPatterAt.<anonymous>"
        if ast.FP.Get_streamName(_env) == path{
            var filter *Pattern_SyntaxFilter
            filter = NewPattern_SyntaxFilter(_env, ast)
            pattern = filter.FP.GetPattern(_env, path, analyzeFileInfo, inqMod)
            Log_log(_env, Log_Level__Log, __func__, 298, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("pattern -- %s", []LnsAny{pattern})
            }))
            
        }
    }))
    return pattern
}






// 35: decl @lns.@tags.@Pattern.SyntaxFilter.getPatternFromNode
func (self *Pattern_SyntaxFilter) getPatternFromNode(_env *LnsEnv, analyzeFileInfo *Option_AnalyzeFileInfo,inqMod string,nearest *Nodes.Nodes_Node) LnsAny {
    __func__ := "@lns.@tags.@Pattern.SyntaxFilter.getPatternFromNode"
    var Pattern_isInner func(_env *LnsEnv, pos Types.Types_Position,name string) bool
    Pattern_isInner = func(_env *LnsEnv, pos Types.Types_Position,name string) bool {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( pos.LineNo == analyzeFileInfo.FP.Get_lineNo(_env)) &&
            _env.SetStackVal( pos.Column <= analyzeFileInfo.FP.Get_column(_env)) &&
            _env.SetStackVal( pos.Column + len(name) >= analyzeFileInfo.FP.Get_column(_env)) ).(bool)){
            return true
        }
        return false
    }
    Log_log(_env, Log_Level__Log, __func__, 20, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("%s %s:%4d:%3d -- %s", []LnsAny{"nearestNode -- ", nearest.FP.Get_effectivePos(_env).StreamName, nearest.FP.Get_effectivePos(_env).LineNo, nearest.FP.Get_effectivePos(_env).Column, Nodes.Nodes_getNodeKindName(_env, nearest.FP.Get_kind(_env))})
    }))
    
    
    {
        _workNode := Nodes.Nodes_ImportNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ImportNode)
            if _switch0 := inqMod; _switch0 == Option_InqMode__Def {
                return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
            } else if _switch0 == Option_InqMode__Ref {
                return Ast_getFullNameSym(_env, &self.Nodes_Filter, workNode.FP.Get_info(_env).FP.Get_symbolInfo(_env))
            } else if _switch0 == Option_InqMode__Set || _switch0 == Option_InqMode__AllMut || _switch0 == Option_InqMode__Async || _switch0 == Option_InqMode__Noasync || _switch0 == Option_InqMode__Luaval || _switch0 == Option_InqMode__AsyncLock {
            }
            return nil
        }
    }
    {
        _workNode := Nodes.Nodes_ExpRefNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpRefNode)
            return Ast_getFullNameSym(_env, &self.Nodes_Filter, workNode.FP.Get_symbolInfo(_env))
        }
    }
    {
        _workNode := Nodes.Nodes_RefFieldNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_RefFieldNode)
            if Pattern_isInner(_env, workNode.FP.Get_field(_env).Pos, workNode.FP.Get_field(_env).Txt){
                {
                    _symbolInfo := workNode.FP.Get_symbolInfo(_env)
                    if !Lns_IsNil( _symbolInfo ) {
                        symbolInfo := _symbolInfo.(*LuneAst.Ast_SymbolInfo)
                        return Ast_getFullNameSym(_env, &self.Nodes_Filter, symbolInfo)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclVarNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclVarNode)
            for _, _varSym := range( workNode.FP.Get_symbolInfoList(_env).Items ) {
                varSym := _varSym.(LuneAst.Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Pattern_isInner(_env, Lns_unwrap( varSym.FP.Get_pos(_env)).(Types.Types_Position), varSym.FP.Get_name(_env)){
                    return Ast_getFullNameSym(_env, &self.Nodes_Filter, varSym)
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpOmitEnumNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpOmitEnumNode)
            if Pattern_isInner(_env, workNode.FP.Get_effectivePos(_env), workNode.FP.Get_valInfo(_env).FP.Get_name(_env)){
                return Ast_getFullNameSym(_env, &self.Nodes_Filter, workNode.FP.Get_valInfo(_env).FP.Get_symbolInfo(_env))
            }
        }
    }
    {
        _workNode := Nodes.Nodes_NewAlgeValNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_NewAlgeValNode)
            if Pattern_isInner(_env, workNode.FP.Get_effectivePos(_env), workNode.FP.Get_valInfo(_env).FP.Get_name(_env)){
                return Ast_getFullNameSym(_env, &self.Nodes_Filter, workNode.FP.Get_valInfo(_env).FP.Get_symbolInfo(_env))
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpMacroExpNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpMacroExpNode)
            if Pattern_isInner(_env, workNode.FP.Get_effectivePos(_env), workNode.FP.Get_expType(_env).FP.Get_rawTxt(_env)){
                return self.FP.GetFull(_env, workNode.FP.Get_macroType(_env), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclFuncNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclFuncNode)
            {
                _name := workNode.FP.Get_declInfo(_env).FP.Get_name(_env)
                if !Lns_IsNil( _name ) {
                    name := _name.(*Types.Types_Token)
                    if Pattern_isInner(_env, name.Pos, name.Txt){
                        return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
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
            name = workNode.FP.Get_name(_env)
            if Pattern_isInner(_env, name.Pos, name.Txt){
                return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
            }
            for _, _valInfo := range( workNode.FP.Get_enumType(_env).FP.Get_name2EnumValInfo(_env).Items ) {
                valInfo := _valInfo.(LuneAst.Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                if Pattern_isInner(_env, Lns_unwrap( valInfo.FP.Get_symbolInfo(_env).FP.Get_pos(_env)).(Types.Types_Position), valInfo.FP.Get_symbolInfo(_env).FP.Get_name(_env)){
                    return Ast_getFullNameSym(_env, &self.Nodes_Filter, valInfo.FP.Get_symbolInfo(_env))
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclAlgeNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclAlgeNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_name(_env)
            if Pattern_isInner(_env, name.Pos, name.Txt){
                return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
            }
            for _, _valInfo := range( workNode.FP.Get_algeType(_env).FP.Get_valInfoMap(_env).Items ) {
                valInfo := _valInfo.(LuneAst.Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                if Pattern_isInner(_env, Lns_unwrap( valInfo.FP.Get_symbolInfo(_env).FP.Get_pos(_env)).(Types.Types_Position), valInfo.FP.Get_symbolInfo(_env).FP.Get_name(_env)){
                    return Ast_getFullNameSym(_env, &self.Nodes_Filter, valInfo.FP.Get_symbolInfo(_env))
                }
                
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclClassNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclClassNode)
            if Pattern_isInner(_env, workNode.FP.Get_name(_env).Pos, workNode.FP.Get_name(_env).Txt){
                return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMethodNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclMethodNode)
            {
                _name := workNode.FP.Get_declInfo(_env).FP.Get_name(_env)
                if !Lns_IsNil( _name ) {
                    name := _name.(*Types.Types_Token)
                    if Pattern_isInner(_env, name.Pos, name.Txt){
                        return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
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
            name = workNode.FP.Get_name(_env)
            if Pattern_isInner(_env, name.Pos, name.Txt){
                return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMemberNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclMemberNode)
            if Pattern_isInner(_env, Lns_unwrap( workNode.FP.Get_symbolInfo(_env).FP.Get_pos(_env)).(Types.Types_Position), workNode.FP.Get_symbolInfo(_env).FP.Get_name(_env)){
                return Ast_getFullNameSym(_env, &self.Nodes_Filter, workNode.FP.Get_symbolInfo(_env))
            }
            
            {
                _token, _symbol := workNode.FP.Get_getterToken(_env), workNode.FP.GetGetterSym(_env)
                if !Lns_IsNil( _token ) && !Lns_IsNil( _symbol ) {
                    token := _token.(*Types.Types_Token)
                    symbol := _symbol.(*LuneAst.Ast_SymbolInfo)
                    if Pattern_isInner(_env, token.Pos, token.Txt){
                        return Ast_getFullNameSym(_env, &self.Nodes_Filter, symbol)
                    }
                }
            }
            {
                _token, _symbol := workNode.FP.Get_setterToken(_env), workNode.FP.GetSetterSym(_env)
                if !Lns_IsNil( _token ) && !Lns_IsNil( _symbol ) {
                    token := _token.(*Types.Types_Token)
                    symbol := _symbol.(*LuneAst.Ast_SymbolInfo)
                    if Pattern_isInner(_env, token.Pos, token.Txt){
                        return Ast_getFullNameSym(_env, &self.Nodes_Filter, symbol)
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
                _name := workNode.FP.Get_declInfo(_env).FP.Get_name(_env)
                if !Lns_IsNil( _name ) {
                    name := _name.(*Types.Types_Token)
                    if Pattern_isInner(_env, name.Pos, name.Txt){
                        return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
                    }
                }
            }
            
        }
    }
    {
        _workNode := Nodes.Nodes_ExpCallSuperCtorNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpCallSuperCtorNode)
            return self.FP.GetFull(_env, workNode.FP.Get_methodType(_env), false)
        }
    }
    {
        _workNode := Nodes.Nodes_ExpCallSuperNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpCallSuperNode)
            return self.FP.GetFull(_env, workNode.FP.Get_methodType(_env), false)
        }
    }
    {
        _workNode := Nodes.Nodes_ExpNewNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpNewNode)
            if Pattern_isInner(_env, workNode.FP.Get_pos(_env), "new"){
                return self.FP.GetFull(_env, workNode.FP.Get_ctorTypeInfo(_env), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_DeclMacroNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclMacroNode)
            var name *Types.Types_Token
            name = workNode.FP.Get_declInfo(_env).FP.Get_name(_env)
            if Pattern_isInner(_env, name.Pos, name.Txt){
                return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_ExpMacroExpNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_ExpMacroExpNode)
            if Pattern_isInner(_env, workNode.FP.Get_pos(_env), workNode.FP.Get_macroType(_env).FP.Get_rawTxt(_env)){
                return self.FP.GetFull(_env, workNode.FP.Get_macroType(_env), false)
            }
        }
    }
    {
        _workNode := Nodes.Nodes_GetFieldNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_GetFieldNode)
            {
                _symbolInfo := workNode.FP.Get_symbolInfo(_env)
                if !Lns_IsNil( _symbolInfo ) {
                    symbolInfo := _symbolInfo.(*LuneAst.Ast_SymbolInfo)
                    if Pattern_isInner(_env, workNode.FP.Get_field(_env).Pos, workNode.FP.Get_field(_env).Txt){
                        return Ast_getFullNameSym(_env, &self.Nodes_Filter, symbolInfo)
                    }
                }
            }
        }
    }
    {
        _workNode := Nodes.Nodes_AliasNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_AliasNode)
            if Pattern_isInner(_env, Lns_unwrap( workNode.FP.Get_newSymbol(_env).FP.Get_pos(_env)).(Types.Types_Position), workNode.FP.Get_newSymbol(_env).FP.Get_name(_env)){
                return Ast_getFullNameSym(_env, &self.Nodes_Filter, workNode.FP.Get_newSymbol(_env))
            }
            
        }
    }
    {
        _workNode := Nodes.Nodes_DeclFormNodeDownCastF(nearest.FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes.Nodes_DeclFormNode)
            {
                _name := workNode.FP.Get_declInfo(_env).FP.Get_name(_env)
                if !Lns_IsNil( _name ) {
                    name := _name.(*Types.Types_Token)
                    if Pattern_isInner(_env, name.Pos, name.Txt){
                        return self.FP.GetFull(_env, workNode.FP.Get_expType(_env), false)
                    }
                }
            }
            
        }
    }
    Log_log(_env, Log_Level__Err, __func__, 193, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("unknown pattern -- %s", []LnsAny{Nodes.Nodes_getNodeKindName(_env, nearest.FP.Get_kind(_env))})
    }))
    
    return nil
}
// 197: decl @lns.@tags.@Pattern.SyntaxFilter.getPattern
func (self *Pattern_SyntaxFilter) GetPattern(_env *LnsEnv, path string,analyzeFileInfo *Option_AnalyzeFileInfo,inqMod string) LnsAny {
    var Pattern_isInner func(_env *LnsEnv, pos Types.Types_Position,name string) bool
    Pattern_isInner = func(_env *LnsEnv, pos Types.Types_Position,name string) bool {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( pos.LineNo == analyzeFileInfo.FP.Get_lineNo(_env)) &&
            _env.SetStackVal( pos.Column <= analyzeFileInfo.FP.Get_column(_env)) &&
            _env.SetStackVal( pos.Column + len(name) >= analyzeFileInfo.FP.Get_column(_env)) ).(bool)){
            return true
        }
        return false
    }
    var pattern LnsAny
    pattern = nil
    if self.ast.FP.Get_streamName(_env) == path{
        var nearestNode LnsAny
        nearestNode = nil
        self.ast.FP.Get_node(_env).FP.Visit(_env, Nodes.Nodes_NodeVisitor(func(_env *LnsEnv, node *Nodes.Nodes_Node,parent *Nodes.Nodes_Node,relation string,depth LnsInt) LnsInt {
            __func__ := "@lns.@tags.@Pattern.SyntaxFilter.getPattern.<anonymous>"
            if analyzeFileInfo.FP.Get_path(_env) == node.FP.Get_effectivePos(_env).StreamName{
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( analyzeFileInfo.FP.Get_lineNo(_env) == node.FP.Get_effectivePos(_env).LineNo) &&
                    _env.SetStackVal( analyzeFileInfo.FP.Get_column(_env) >= node.FP.Get_effectivePos(_env).Column) ).(bool)){
                    {
                        _declMemberNode := Nodes.Nodes_DeclMemberNodeDownCastF(node.FP)
                        if !Lns_IsNil( _declMemberNode ) {
                            declMemberNode := _declMemberNode.(*Nodes.Nodes_DeclMemberNode)
                            {
                                _token := declMemberNode.FP.Get_getterToken(_env)
                                if !Lns_IsNil( _token ) {
                                    token := _token.(*Types.Types_Token)
                                    if Pattern_isInner(_env, token.Pos, token.Txt){
                                        nearestNode = node
                                        return Nodes.Nodes_NodeVisitMode__End
                                    }
                                }
                            }
                            {
                                _token := declMemberNode.FP.Get_setterToken(_env)
                                if !Lns_IsNil( _token ) {
                                    token := _token.(*Types.Types_Token)
                                    if Pattern_isInner(_env, token.Pos, token.Txt){
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
                            if nearest.FP.Get_effectivePos(_env).LineNo != analyzeFileInfo.FP.Get_lineNo(_env){
                                nearestNode = node
                            }
                            if nearest.FP.Get_effectivePos(_env).Column < node.FP.Get_effectivePos(_env).Column{
                                nearestNode = node
                            } else if nearest.FP.Get_effectivePos(_env).Column == node.FP.Get_effectivePos(_env).Column{
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
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( nearest.FP.Get_effectivePos(_env).LineNo < node.FP.Get_effectivePos(_env).LineNo) &&
                                _env.SetStackVal( node.FP.Get_effectivePos(_env).LineNo < analyzeFileInfo.FP.Get_lineNo(_env)) ).(bool)){
                                nearestNode = node
                            }
                        } else {
                            nearestNode = node
                        }
                    }
                }
                Log_log(_env, Log_Level__Trace, __func__, 20, Log_CreateMessage(func(_env *LnsEnv) string {
                    return _env.GetVM().String_format("%s %s:%4d:%3d -- %s", []LnsAny{"visit:", node.FP.Get_effectivePos(_env).StreamName, node.FP.Get_effectivePos(_env).LineNo, node.FP.Get_effectivePos(_env).Column, Nodes.Nodes_getNodeKindName(_env, node.FP.Get_kind(_env))})
                }))
                
                
                return Nodes.Nodes_NodeVisitMode__Child
            }
            return Nodes.Nodes_NodeVisitMode__Next
        }), 0, NewLnsSet([]LnsAny{}))
        {
            _nearest := nearestNode
            if !Lns_IsNil( _nearest ) {
                nearest := _nearest.(*Nodes.Nodes_Node)
                pattern = self.FP.getPatternFromNode(_env, analyzeFileInfo, inqMod, nearest)
            }
        }
    }
    return pattern
}
// declaration Class -- SyntaxFilter
type Pattern_SyntaxFilterMtd interface {
    DefaultProcess(_env *LnsEnv, arg1 *Nodes.Nodes_Node, arg2 LnsAny)
    GetFull(_env *LnsEnv, arg1 *LuneAst.Ast_TypeInfo, arg2 bool) string
    GetPattern(_env *LnsEnv, arg1 string, arg2 *Option_AnalyzeFileInfo, arg3 string) LnsAny
    getPatternFromNode(_env *LnsEnv, arg1 *Option_AnalyzeFileInfo, arg2 string, arg3 *Nodes.Nodes_Node) LnsAny
    Get_moduleInfoManager(_env *LnsEnv) LuneAst.Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *LuneAst.Ast_TypeNameCtrl
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes.Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes.Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(_env *LnsEnv, arg1 *Nodes.Nodes_ApplyNode, arg2 LnsAny)
    ProcessAsyncLock(_env *LnsEnv, arg1 *Nodes.Nodes_AsyncLockNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes.Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes.Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes.Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes.Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes.Nodes_BreakNode, arg2 LnsAny)
    ProcessCondRet(_env *LnsEnv, arg1 *Nodes.Nodes_CondRetNode, arg2 LnsAny)
    ProcessCondRetList(_env *LnsEnv, arg1 *Nodes.Nodes_CondRetListNode, arg2 LnsAny)
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
    ProcessDeclTuple(_env *LnsEnv, arg1 *Nodes.Nodes_DeclTupleNode, arg2 LnsAny)
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes.Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes.Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes.Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpExpandTuple(_env *LnsEnv, arg1 *Nodes.Nodes_ExpExpandTupleNode, arg2 LnsAny)
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
    ProcessLetExpandTuple(_env *LnsEnv, arg1 *Nodes.Nodes_LetExpandTupleNode, arg2 LnsAny)
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
    ProcessRequest(_env *LnsEnv, arg1 *Nodes.Nodes_RequestNode, arg2 LnsAny)
    ProcessReturn(_env *LnsEnv, arg1 *Nodes.Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(_env *LnsEnv, arg1 *Nodes.Nodes_RootNode, arg2 LnsAny)
    ProcessScope(_env *LnsEnv, arg1 *Nodes.Nodes_ScopeNode, arg2 LnsAny)
    ProcessShebang(_env *LnsEnv, arg1 *Nodes.Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(_env *LnsEnv, arg1 *Nodes.Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(_env *LnsEnv, arg1 *Nodes.Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(_env *LnsEnv, arg1 *Nodes.Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(_env *LnsEnv, arg1 *Nodes.Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(_env *LnsEnv, arg1 *Nodes.Nodes_TestCaseNode, arg2 LnsAny)
    ProcessTupleConst(_env *LnsEnv, arg1 *Nodes.Nodes_TupleConstNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes.Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes.Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes.Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes.Nodes_WhileNode, arg2 LnsAny)
}
type Pattern_SyntaxFilter struct {
    Nodes.Nodes_Filter
    ast *AstInfo.AstInfo_ASTInfo
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
func NewPattern_SyntaxFilter(_env *LnsEnv, arg1 *AstInfo.AstInfo_ASTInfo) *Pattern_SyntaxFilter {
    obj := &Pattern_SyntaxFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitPattern_SyntaxFilter(_env, arg1)
    return obj
}
// 28: DeclConstr
func (self *Pattern_SyntaxFilter) InitPattern_SyntaxFilter(_env *LnsEnv, ast *AstInfo.AstInfo_ASTInfo) {
    self.InitNodes_Filter(_env, true, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env))
    self.ast = ast
}


func Lns_Pattern_init(_env *LnsEnv) {
    if init_Pattern { return }
    init_Pattern = true
    Pattern__mod__ = "@lns.@tags.@Pattern"
    Lns_InitMod()
    Lns_DBCtrl_init(_env)
    Lns_Option_init(_env)
    Lns_Log_init(_env)
    Lns_Ast_init(_env)
    LuneOpt.Lns_Option_init(_env)
    Nodes.Lns_Nodes_init(_env)
    Parser.Lns_Parser_init(_env)
    TransUnit.Lns_TransUnit_init(_env)
    front.Lns_front_init(_env)
    Types.Lns_Types_init(_env)
    LuneAst.Lns_Ast_init(_env)
    AstInfo.Lns_AstInfo_init(_env)
    LuaVer.Lns_LuaVer_init(_env)
    LuneLog.Lns_Log_init(_env)
    LuneUtil.Lns_Util_init(_env)
}
func init() {
    init_Pattern = false
}
