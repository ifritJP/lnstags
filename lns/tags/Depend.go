// This code is transcompiled by LuneScript.
package tags

import (
	"os"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

var Depend__mod__ string

// 6: decl @lns.@tags.@Depend.getCurDir
func Depend_getCurDir(_env *LnsEnv) string {
	dir, _ := os.Getwd()
	return dir
}

func Lns_Depend_init(_env *LnsEnv) {
}
