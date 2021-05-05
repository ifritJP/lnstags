// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
// for 31
func Util_convExp160(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 47: decl @lns.@tags.@Util.outputLocate
func Util_outputLocate(stream Lns_oStream,symbol string,path string,lineAccessor LnsAny,lineNo LnsInt) {
    var line string
    
    {
        _line := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(lineAccessor) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Util_SourceCodeLineAccessor).FP.GetLine(lineNo)})/* 51:16 */)
        if _line == nil{
            line = ""
            
        } else {
            line = _line.(string)
        }
    }
    stream.Write(Lns_getVM().String_format("%-16s %4d %-16s %s\n", []LnsAny{symbol, lineNo, path, line}))
}

// declaration Class -- SourceCodeLineAccessor
type Util_SourceCodeLineAccessorMtd interface {
    GetLine(arg1 LnsInt) LnsAny
    Get_path() string
}
type Util_SourceCodeLineAccessor struct {
    path string
    lineList *LnsList
    FP Util_SourceCodeLineAccessorMtd
}
func Util_SourceCodeLineAccessor2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_SourceCodeLineAccessor).FP
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
func NewUtil_SourceCodeLineAccessor(arg1 string, arg2 *LnsList) *Util_SourceCodeLineAccessor {
    obj := &Util_SourceCodeLineAccessor{}
    obj.FP = obj
    obj.InitUtil_SourceCodeLineAccessor(arg1, arg2)
    return obj
}
func (self *Util_SourceCodeLineAccessor) InitUtil_SourceCodeLineAccessor(arg1 string, arg2 *LnsList) {
    self.path = arg1
    self.lineList = arg2
}
func (self *Util_SourceCodeLineAccessor) Get_path() string{ return self.path }
// 5: decl @lns.@tags.@Util.SourceCodeLineAccessor.getLine
func (self *Util_SourceCodeLineAccessor) GetLine(lineNo LnsInt) LnsAny {
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( lineNo < 0) ||
        Lns_GetEnv().SetStackVal( self.lineList.Len() < lineNo) ).(bool){
        return nil
    }
    return self.lineList.GetAt(lineNo).(string)
}


// declaration Class -- SourceCodeLineAccessorFactory
type Util_SourceCodeLineAccessorFactoryMtd interface {
    Create(arg1 string, arg2 LnsAny) LnsAny
}
type Util_SourceCodeLineAccessorFactory struct {
    path2accessor *LnsMap
    FP Util_SourceCodeLineAccessorFactoryMtd
}
func Util_SourceCodeLineAccessorFactory2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_SourceCodeLineAccessorFactory).FP
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
func NewUtil_SourceCodeLineAccessorFactory() *Util_SourceCodeLineAccessorFactory {
    obj := &Util_SourceCodeLineAccessorFactory{}
    obj.FP = obj
    obj.InitUtil_SourceCodeLineAccessorFactory()
    return obj
}
// 16: DeclConstr
func (self *Util_SourceCodeLineAccessorFactory) InitUtil_SourceCodeLineAccessorFactory() {
    self.path2accessor = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 20: decl @lns.@tags.@Util.SourceCodeLineAccessorFactory.create
func (self *Util_SourceCodeLineAccessorFactory) Create(filePath string,fileContents LnsAny) LnsAny {
    {
        __exp := self.path2accessor.Get(filePath)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Util_SourceCodeLineAccessor)
            return _exp
        }
    }
    var lineList *LnsList
    lineList = NewLnsList([]LnsAny{})
    if fileContents != nil{
        fileContents_32 := fileContents.(string)
        {
            _form141, _param141, _prev141 := Lns_getVM().String_gmatch(fileContents_32, "[^\n]*\n")
            for {
                _work141 := _form141.(*Lns_luaValue).Call( Lns_2DDD( _param141, _prev141 ) )
                _prev141 = Lns_getFromMulti(_work141,0)
                if Lns_IsNil( _prev141 ) { break }
                line := _prev141.(string)
                lineList.Insert(Lns_getVM().String_sub(line, 1, len(line) - 1))
            }
        }
    } else {
        var handle Lns_luaStream
        
        {
            _handle := Util_convExp160(Lns_2DDD(Lns_io_open(filePath, "r")))
            if _handle == nil{
                return nil
            } else {
                handle = _handle.(Lns_luaStream)
            }
        }
        for  {
            var text string
            
            {
                _text := handle.Read("*l")
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
    accessor = NewUtil_SourceCodeLineAccessor(filePath, lineList)
    self.path2accessor.Set(filePath,accessor)
    return accessor
}


func Lns_Util_init() {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@lns.@tags.@Util"
    Lns_InitMod()
}
func init() {
    init_Util = false
}
