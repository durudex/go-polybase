// Code generated by qtc from "model.qtpl". DO NOT EDIT.
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

func StreamModel(qw422016 *qt422016.Writer, name string, fields []*polylang.Field) {
	qw422016.N().S(`
type `)
	qw422016.E().S(strcase.ToCamel(name))
	qw422016.N().S(` struct {
`)
	for _, field := range fields {
		if field.Optional {
			qw422016.N().S(`	`)
			qw422016.E().S(strcase.ToCamel(field.Name))
			qw422016.N().S(` *`)
			qw422016.E().S(field.Type)
			qw422016.N().S(` `)
			qw422016.N().S("`")
			qw422016.N().S(`json:"`)
			qw422016.E().S(field.Name)
			qw422016.N().S(`"`)
			qw422016.N().S("`")
			qw422016.N().S(`
`)
			continue
		}
		qw422016.N().S(`	`)
		qw422016.E().S(strcase.ToCamel(field.Name))
		qw422016.N().S(` `)
		qw422016.E().S(field.Type)
		qw422016.N().S(` `)
		qw422016.N().S("`")
		qw422016.N().S(`json:"`)
		qw422016.E().S(field.Name)
		qw422016.N().S(`"`)
		qw422016.N().S("`")
		qw422016.N().S(`
`)
	}
	qw422016.N().S(`}
`)
}

func WriteModel(qq422016 qtio422016.Writer, name string, fields []*polylang.Field) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamModel(qw422016, name, fields)
	qt422016.ReleaseWriter(qw422016)
}

func Model(name string, fields []*polylang.Field) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteModel(qb422016, name, fields)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
