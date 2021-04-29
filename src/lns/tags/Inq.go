// This code is transcompiled by LuneScript.
package tags
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Inq bool
var Inq__mod__ string

// 4: decl @lns.@tags.@Inq.InqDef
func Inq_InqDef(db *DBCtrl_DBCtrl,pattern string) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory()
    db.FP.MapSymbolDecl(pattern, DBCtrl_callbackSymbolDecl(func(item *DBCtrl_ItemSymbolDecl) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(item.FP.Get_fileId())
            if _path == nil{
                panic(Lns_getVM().String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId()}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(Lns_io_stdout, pattern, path, factory.FP.Create(path, nil), item.FP.Get_line())
        return true
    }))
}


// 18: decl @lns.@tags.@Inq.InqRef
func Inq_InqRef(db *DBCtrl_DBCtrl,pattern string) {
    var factory *Util_SourceCodeLineAccessorFactory
    factory = NewUtil_SourceCodeLineAccessorFactory()
    db.FP.MapSymbolRef(pattern, DBCtrl_callbackSymbolRef(func(item *DBCtrl_ItemSymbolRef) bool {
        var path string
        
        {
            _path := db.FP.GetFilePath(item.FP.Get_fileId())
            if _path == nil{
                panic(Lns_getVM().String_format("file id is illegal -- %d", []LnsAny{item.FP.Get_fileId()}))
            } else {
                path = _path.(string)
            }
        }
        Util_outputLocate(Lns_io_stdout, pattern, path, factory.FP.Create(path, nil), item.FP.Get_line())
        return true
    }))
}

func Lns_Inq_init() {
    if init_Inq { return }
    init_Inq = true
    Inq__mod__ = "@lns.@tags.@Inq"
    Lns_InitMod()
    Lns_DBCtrl_init()
    Lns_Util_init()
}
func init() {
    init_Inq = false
}
