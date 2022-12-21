// Code generated by qtc from "common.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package template

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamHeader(qw422016 *qt422016.Writer, name string) {
	qw422016.N().S(`// Code generated by github.com/v1def/go-polybase. DO NOT EDIT.

package `)
	qw422016.E().S(name)
	qw422016.N().S(`

import "context"
`)
}

func WriteHeader(qq422016 qtio422016.Writer, name string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamHeader(qw422016, name)
	qt422016.ReleaseWriter(qw422016)
}

func Header(name string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteHeader(qb422016, name)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
