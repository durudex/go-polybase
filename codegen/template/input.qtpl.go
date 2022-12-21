// Code generated by qtc from "input.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package template

import (
	"github.com/v1def/go-polybase/polylang"

	"github.com/iancoleman/strcase"
)

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamInput(qw422016 *qt422016.Writer, coll, name string, params []*polylang.Parameter) {
	qw422016.N().S(`
type `)
	qw422016.E().S(strcase.ToCamel(coll))
	qw422016.E().S(strcase.ToCamel(name))
	qw422016.N().S(`Input struct {
`)
	for _, param := range params {
		qw422016.N().S(`	`)
		qw422016.E().S(strcase.ToCamel(param.Name))
		qw422016.N().S(` `)
		if param.Optional {
			qw422016.N().S(`*`)
		}
		qw422016.N().S(` `)
		qw422016.E().S(param.Type)
		qw422016.N().S(` `)
		qw422016.N().S("`")
		qw422016.N().S(`json:"`)
		qw422016.E().S(param.Name)
		qw422016.N().S(`"`)
		qw422016.N().S("`")
		qw422016.N().S(`
`)
	}
	qw422016.N().S(`}
`)
}

func WriteInput(qq422016 qtio422016.Writer, coll, name string, params []*polylang.Parameter) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamInput(qw422016, coll, name, params)
	qt422016.ReleaseWriter(qw422016)
}

func Input(coll, name string, params []*polylang.Parameter) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteInput(qb422016, coll, name, params)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
