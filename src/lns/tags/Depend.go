// This code is transcompiled by LuneScript.
package tags
import "os"
var Depend__mod__ string
// 6: decl @lns.@tags.@Depend.getCurDir
func Depend_getCurDir() string {
    dir, _ := os.Getwd()
    return dir
}
