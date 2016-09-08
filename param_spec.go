//GParamSpec — Metadata for parameter specifications
package glib

/*
#include <glib-object.h>
*/
import "C"

type ParamSpec C.GParamSpec

func (p *ParamSpec) g() *C.GParamSpec {
	return (*C.GParamSpec)(p)
}

func (p *ParamSpec) Type() Type {
	return TypeFromName("GParam")
}
