// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
// for 35
func Util_convExp0_181(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 51: decl @lns.@tags.@Util.outputLocate
func Util_outputLocate(_env *LnsEnv, stream Lns_oStream,symbol string,path string,lineAccessor LnsAny,lineNo LnsInt) {
    var line string
    
    {
        _line := _env.NilAccFin(_env.NilAccPush(lineAccessor) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Util_SourceCodeLineAccessor).FP.GetLine(_env, lineNo)})/* 55:16 */)
        if _line == nil{
            line = ""
        } else {
            line = _line.(string)
        }
    }
    stream.Write(_env, _env.GetVM().String_format("%-16s %4d %-16s %s\n", Lns_2DDD(symbol, lineNo, path, line)))
}

// 7: decl @lns.@tags.@Util.SourceCodeLineAccessor.getLine
func (self *Util_SourceCodeLineAccessor) GetLine(_env *LnsEnv, lineNo LnsInt) LnsAny {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( lineNo < 0) ||
        _env.SetStackVal( self.lineList.Len() < lineNo) ).(bool){
        return nil
    }
    return self.lineList.GetAt(lineNo)
}
// 22: decl @lns.@tags.@Util.SourceCodeLineAccessorFactory.create
func (self *Util_SourceCodeLineAccessorFactory) Create(_env *LnsEnv, filePath string,fileContents LnsAny) LnsAny {
    {
        __exp := self.path2accessor.Get(filePath)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Util_SourceCodeLineAccessor)
            return _exp
        }
    }
    var lineList *LnsList2_[string]
    lineList = NewLnsList2_[string]([]string{})
    if fileContents != nil{
        fileContents_39 := fileContents.(string)
            {
                _applyForm0, _applyParam0, _applyPrev0 := _env.GetVM().String_gmatch(fileContents_39, "[^\n]*\n")
                for {
                    _applyWork0 := _applyForm0.(*Lns_luaValue).Call( Lns_2DDD( _applyParam0, _applyPrev0 ) )
                    _applyPrev0 = Lns_getFromMulti(_applyWork0,0)
                    if Lns_IsNil( _applyPrev0 ) { break }
                    line := _applyPrev0.(string)
                    lineList.Insert(_env.GetVM().String_sub(line, 1, len(line) - 1))
                }
            }
    } else {
        var handle Lns_luaStream
        
        {
            _handle := Util_convExp0_181(Lns_2DDD(Lns_io_open(filePath, "r")))
            if _handle == nil{
                return nil
            } else {
                handle = _handle.(Lns_luaStream)
            }
        }
        for  {
            var text string
            
            {
                _text := handle.Read(_env, "*l")
                if _text == nil{
                    break
                } else {
                    text = _text.(string)
                }
            }
            lineList.Insert(text)
        }
    }
    var accessor *Util_SourceCodeLineAccessor
    accessor = NewUtil_SourceCodeLineAccessor(_env, filePath, lineList)
    self.path2accessor.Set(filePath,accessor)
    return accessor
}
// declaration Class -- SourceCodeLineAccessor
type Util_SourceCodeLineAccessorMtd interface {
    GetLine(_env *LnsEnv, arg1 LnsInt) LnsAny
    Get_path(_env *LnsEnv) string
}
type Util_SourceCodeLineAccessor struct {
    path string
    lineList *LnsList2_[string]
    FP Util_SourceCodeLineAccessorMtd
}
func Util_SourceCodeLineAccessor2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_SourceCodeLineAccessor).FP
}
func Util_SourceCodeLineAccessor_toSlice(slice []LnsAny) []*Util_SourceCodeLineAccessor {
    ret := make([]*Util_SourceCodeLineAccessor, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_SourceCodeLineAccessorDownCast).ToUtil_SourceCodeLineAccessor()
    }
    return ret
}
type Util_SourceCodeLineAccessorDownCast interface {
    ToUtil_SourceCodeLineAccessor() *Util_SourceCodeLineAccessor
}
func Util_SourceCodeLineAccessorDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_SourceCodeLineAccessorDownCast)
    if ok { return work.ToUtil_SourceCodeLineAccessor() }
    return nil
}
func (obj *Util_SourceCodeLineAccessor) ToUtil_SourceCodeLineAccessor() *Util_SourceCodeLineAccessor {
    return obj
}
func NewUtil_SourceCodeLineAccessor(_env *LnsEnv, arg1 string, arg2 *LnsList2_[string]) *Util_SourceCodeLineAccessor {
    obj := &Util_SourceCodeLineAccessor{}
    obj.FP = obj
    obj.InitUtil_SourceCodeLineAccessor(_env, arg1, arg2)
    return obj
}
func (self *Util_SourceCodeLineAccessor) InitUtil_SourceCodeLineAccessor(_env *LnsEnv, arg1 string, arg2 *LnsList2_[string]) {
    self.path = arg1
    self.lineList = arg2
}
func (self *Util_SourceCodeLineAccessor) Get_path(_env *LnsEnv) string{ return self.path }

// declaration Class -- SourceCodeLineAccessorFactory
type Util_SourceCodeLineAccessorFactoryMtd interface {
    Create(_env *LnsEnv, arg1 string, arg2 LnsAny) LnsAny
}
type Util_SourceCodeLineAccessorFactory struct {
    path2accessor *LnsMap2_[string,*Util_SourceCodeLineAccessor]
    FP Util_SourceCodeLineAccessorFactoryMtd
}
func Util_SourceCodeLineAccessorFactory2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_SourceCodeLineAccessorFactory).FP
}
func Util_SourceCodeLineAccessorFactory_toSlice(slice []LnsAny) []*Util_SourceCodeLineAccessorFactory {
    ret := make([]*Util_SourceCodeLineAccessorFactory, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_SourceCodeLineAccessorFactoryDownCast).ToUtil_SourceCodeLineAccessorFactory()
    }
    return ret
}
type Util_SourceCodeLineAccessorFactoryDownCast interface {
    ToUtil_SourceCodeLineAccessorFactory() *Util_SourceCodeLineAccessorFactory
}
func Util_SourceCodeLineAccessorFactoryDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_SourceCodeLineAccessorFactoryDownCast)
    if ok { return work.ToUtil_SourceCodeLineAccessorFactory() }
    return nil
}
func (obj *Util_SourceCodeLineAccessorFactory) ToUtil_SourceCodeLineAccessorFactory() *Util_SourceCodeLineAccessorFactory {
    return obj
}
func NewUtil_SourceCodeLineAccessorFactory(_env *LnsEnv) *Util_SourceCodeLineAccessorFactory {
    obj := &Util_SourceCodeLineAccessorFactory{}
    obj.FP = obj
    obj.InitUtil_SourceCodeLineAccessorFactory(_env)
    return obj
}
// 18: DeclConstr
func (self *Util_SourceCodeLineAccessorFactory) InitUtil_SourceCodeLineAccessorFactory(_env *LnsEnv) {
    self.path2accessor = NewLnsMap2_[string,*Util_SourceCodeLineAccessor]( map[string]*Util_SourceCodeLineAccessor{})
}


func Lns_Util_init(_env *LnsEnv) {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@lns.@tags.@Util"
    Lns_InitMod()
}
func init() {
    init_Util = false
}
