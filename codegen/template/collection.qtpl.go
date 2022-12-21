// Code generated by qtc from "collection.qtpl". DO NOT EDIT.
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

func StreamCollection(qw422016 *qt422016.Writer, id, coll string, funcs []*polylang.Function) {
	qw422016.N().S(`type I`)
	qw422016.E().S(strcase.ToCamel(coll))
	qw422016.N().S(` interface {
`)
	for _, fc := range funcs {
		qw422016.N().S(`	`)
		qw422016.E().S(strcase.ToCamel(fc.Name))
		qw422016.N().S(`(ctx context.Context `)
		if fc.Name != "constructor" {
			qw422016.N().S(`, id string`)
		}
		qw422016.N().S(`,input *`)
		qw422016.E().S(strcase.ToCamel(coll) + strcase.ToCamel(fc.Name))
		qw422016.N().S(`Input`)
		qw422016.N().S(`)
`)
	}
	qw422016.N().S(`}

type `)
	qw422016.E().S(strcase.ToLowerCamel(coll))
	qw422016.N().S(` struct{ coll polybase.Collection }

func New`)
	qw422016.E().S(strcase.ToCamel(coll))
	qw422016.N().S(`(db polybase.Polybase) I`)
	qw422016.E().S(strcase.ToCamel(coll))
	qw422016.N().S(` {
	return &`)
	qw422016.E().S(strcase.ToLowerCamel(coll))
	qw422016.N().S(`{coll: db.Collection("`)
	qw422016.E().S(id)
	qw422016.N().S(`")}
}
`)
}

func WriteCollection(qq422016 qtio422016.Writer, id, coll string, funcs []*polylang.Function) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamCollection(qw422016, id, coll, funcs)
	qt422016.ReleaseWriter(qw422016)
}

func Collection(id, coll string, funcs []*polylang.Function) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteCollection(qb422016, id, coll, funcs)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
