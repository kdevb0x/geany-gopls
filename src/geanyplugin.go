// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package geanygopls

/*
#cgo pkgconfig: geany

#include <geany/geanyplugin.h>
#include <geany/plugindata.h>

extern void geany_load_module(GeanyPlugin * plugin);
*/
import "C"

type GeanyPlugin C.GeanyPlugin

//export geany_load_module
func geany_load_module(plugin *C.GeanyPlugin) {

}
