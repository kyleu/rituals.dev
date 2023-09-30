// Code generated by qtc from "Statement.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vdatabase/Statement.html:2
package vdatabase

//line views/vdatabase/Statement.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vdatabase/Statement.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vdatabase/Statement.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vdatabase/Statement.html:11
type Statement struct {
	layout.Basic
	Statement *database.DebugStatement
}

//line views/vdatabase/Statement.html:16
func (p *Statement) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Statement.html:16
	qw422016.N().S(`
`)
//line views/vdatabase/Statement.html:17
	s := p.Statement

//line views/vdatabase/Statement.html:17
	qw422016.N().S(`  <div class="card">
    <div class="right">`)
//line views/vdatabase/Statement.html:19
	qw422016.E().S(util.MicrosToMillis(s.Timing))
//line views/vdatabase/Statement.html:19
	qw422016.N().S(` elapsed</div>
    <h3>`)
//line views/vdatabase/Statement.html:20
	components.StreamSVGRefIcon(qw422016, `database`, ps)
//line views/vdatabase/Statement.html:20
	qw422016.N().S(`Statement [`)
//line views/vdatabase/Statement.html:20
	qw422016.N().D(p.Statement.Index)
//line views/vdatabase/Statement.html:20
	qw422016.N().S(`]</h3>
    <div class="right">`)
//line views/vdatabase/Statement.html:21
	qw422016.N().D(s.Count)
//line views/vdatabase/Statement.html:21
	qw422016.N().S(` rows returned</div>
    <em>`)
//line views/vdatabase/Statement.html:22
	qw422016.E().S(s.Message)
//line views/vdatabase/Statement.html:22
	qw422016.N().S(`</em>
`)
//line views/vdatabase/Statement.html:23
	if s.Error != "" {
//line views/vdatabase/Statement.html:23
		qw422016.N().S(`    <div class="mt error">`)
//line views/vdatabase/Statement.html:24
		qw422016.E().S(s.Error)
//line views/vdatabase/Statement.html:24
		qw422016.N().S(`</div>
`)
//line views/vdatabase/Statement.html:25
	}
//line views/vdatabase/Statement.html:25
	qw422016.N().S(`    <div class="mt">
      <ul class="accordion">
        <li>
          <input id="accordion-sql" type="checkbox" hidden />
          <label for="accordion-sql">`)
//line views/vdatabase/Statement.html:30
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vdatabase/Statement.html:30
	qw422016.N().S(` SQL</label>
          <div class="bd"><div><div>
            `)
//line views/vdatabase/Statement.html:32
	streamstatementSQL(qw422016, s)
//line views/vdatabase/Statement.html:32
	qw422016.N().S(`
          </div></div></div>
        </li>
`)
//line views/vdatabase/Statement.html:35
	if len(s.Values) > 0 {
//line views/vdatabase/Statement.html:35
		qw422016.N().S(`        <li>
          <input id="accordion-values" type="checkbox" hidden />
          <label for="accordion-values">`)
//line views/vdatabase/Statement.html:38
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vdatabase/Statement.html:38
		qw422016.N().S(` Values</label>
          <div class="bd"><div><div>
            `)
//line views/vdatabase/Statement.html:40
		streamstatementValues(qw422016, s)
//line views/vdatabase/Statement.html:40
		qw422016.N().S(`
          </div></div></div>
        </li>
`)
//line views/vdatabase/Statement.html:43
	}
//line views/vdatabase/Statement.html:43
	qw422016.N().S(`        <li>
          <input id="accordion-out" type="checkbox" hidden />
          <label for="accordion-out">`)
//line views/vdatabase/Statement.html:46
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vdatabase/Statement.html:46
	qw422016.N().S(` Result</label>
          <div class="bd"><div><div>
            `)
//line views/vdatabase/Statement.html:48
	streamstatementOut(qw422016, s)
//line views/vdatabase/Statement.html:48
	qw422016.N().S(`
          </div></div></div>
        </li>
`)
//line views/vdatabase/Statement.html:51
	if len(s.Extra) > 0 {
//line views/vdatabase/Statement.html:51
		qw422016.N().S(`        <li>
          <input id="accordion-extra" type="checkbox" hidden />
          <label for="accordion-extra">`)
//line views/vdatabase/Statement.html:54
		components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vdatabase/Statement.html:54
		qw422016.N().S(` Query Plan</label>
          <div class="bd"><div><div>
            `)
//line views/vdatabase/Statement.html:56
		components.StreamDisplayMaps(qw422016, s.Extra, nil, true, ps)
//line views/vdatabase/Statement.html:56
		qw422016.N().S(`
          </div></div></div>
        </li>
`)
//line views/vdatabase/Statement.html:59
	}
//line views/vdatabase/Statement.html:59
	qw422016.N().S(`      </ul>
    </div>
  </div>
`)
//line views/vdatabase/Statement.html:63
}

//line views/vdatabase/Statement.html:63
func (p *Statement) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Statement.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Statement.html:63
	p.StreamBody(qw422016, as, ps)
//line views/vdatabase/Statement.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Statement.html:63
}

//line views/vdatabase/Statement.html:63
func (p *Statement) Body(as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Statement.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Statement.html:63
	p.WriteBody(qb422016, as, ps)
//line views/vdatabase/Statement.html:63
	qs422016 := string(qb422016.B)
//line views/vdatabase/Statement.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Statement.html:63
	return qs422016
//line views/vdatabase/Statement.html:63
}

//line views/vdatabase/Statement.html:65
func streamstatementSQL(qw422016 *qt422016.Writer, s *database.DebugStatement) {
//line views/vdatabase/Statement.html:65
	qw422016.N().S(`
`)
//line views/vdatabase/Statement.html:66
	out, _ := cutil.FormatLang(s.SQL, "sql")

//line views/vdatabase/Statement.html:66
	qw422016.N().S(`  `)
//line views/vdatabase/Statement.html:67
	qw422016.N().S(out)
//line views/vdatabase/Statement.html:67
	qw422016.N().S(`
`)
//line views/vdatabase/Statement.html:68
}

//line views/vdatabase/Statement.html:68
func writestatementSQL(qq422016 qtio422016.Writer, s *database.DebugStatement) {
//line views/vdatabase/Statement.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Statement.html:68
	streamstatementSQL(qw422016, s)
//line views/vdatabase/Statement.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Statement.html:68
}

//line views/vdatabase/Statement.html:68
func statementSQL(s *database.DebugStatement) string {
//line views/vdatabase/Statement.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Statement.html:68
	writestatementSQL(qb422016, s)
//line views/vdatabase/Statement.html:68
	qs422016 := string(qb422016.B)
//line views/vdatabase/Statement.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Statement.html:68
	return qs422016
//line views/vdatabase/Statement.html:68
}

//line views/vdatabase/Statement.html:70
func streamstatementValues(qw422016 *qt422016.Writer, s *database.DebugStatement) {
//line views/vdatabase/Statement.html:70
	qw422016.N().S(`
  <div class="overflow full-width">
    <table>
      <thead>
        <tr>
          <th>#</th>
          <th>Values</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vdatabase/Statement.html:80
	for idx, v := range s.Values {
//line views/vdatabase/Statement.html:80
		qw422016.N().S(`        <tr>
          <td>`)
//line views/vdatabase/Statement.html:82
		qw422016.N().D(idx + 1)
//line views/vdatabase/Statement.html:82
		qw422016.N().S(`</td>
          <td>`)
//line views/vdatabase/Statement.html:83
		components.StreamJSON(qw422016, v)
//line views/vdatabase/Statement.html:83
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vdatabase/Statement.html:85
	}
//line views/vdatabase/Statement.html:85
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vdatabase/Statement.html:89
}

//line views/vdatabase/Statement.html:89
func writestatementValues(qq422016 qtio422016.Writer, s *database.DebugStatement) {
//line views/vdatabase/Statement.html:89
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Statement.html:89
	streamstatementValues(qw422016, s)
//line views/vdatabase/Statement.html:89
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Statement.html:89
}

//line views/vdatabase/Statement.html:89
func statementValues(s *database.DebugStatement) string {
//line views/vdatabase/Statement.html:89
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Statement.html:89
	writestatementValues(qb422016, s)
//line views/vdatabase/Statement.html:89
	qs422016 := string(qb422016.B)
//line views/vdatabase/Statement.html:89
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Statement.html:89
	return qs422016
//line views/vdatabase/Statement.html:89
}

//line views/vdatabase/Statement.html:91
func streamstatementOut(qw422016 *qt422016.Writer, s *database.DebugStatement) {
//line views/vdatabase/Statement.html:91
	qw422016.N().S(`
`)
//line views/vdatabase/Statement.html:92
	if len(s.Out) == 0 {
//line views/vdatabase/Statement.html:92
		qw422016.N().S(`  <em>no results</em>
`)
//line views/vdatabase/Statement.html:94
	} else {
//line views/vdatabase/Statement.html:94
		qw422016.N().S(`  <div class="overflow full-width">
    <table>
      <thead>
        <tr>
          <th class="shrink">Idx</th>
          <th>Result</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vdatabase/Statement.html:104
		for idx, v := range s.Out {
//line views/vdatabase/Statement.html:104
			qw422016.N().S(`        <tr>
          <td>`)
//line views/vdatabase/Statement.html:106
			qw422016.N().D(idx + 1)
//line views/vdatabase/Statement.html:106
			qw422016.N().S(`</td>
          <td>`)
//line views/vdatabase/Statement.html:107
			components.StreamJSON(qw422016, v)
//line views/vdatabase/Statement.html:107
			qw422016.N().S(`</td>
        </tr>
`)
//line views/vdatabase/Statement.html:109
		}
//line views/vdatabase/Statement.html:110
		if s.Count > len(s.Out) {
//line views/vdatabase/Statement.html:110
			qw422016.N().S(`        <tr>
          <td colspan="2">...and `)
//line views/vdatabase/Statement.html:112
			qw422016.N().D(s.Count - len(s.Out))
//line views/vdatabase/Statement.html:112
			qw422016.N().S(` additional rows</td>
        </tr>
`)
//line views/vdatabase/Statement.html:114
		}
//line views/vdatabase/Statement.html:114
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vdatabase/Statement.html:118
	}
//line views/vdatabase/Statement.html:119
}

//line views/vdatabase/Statement.html:119
func writestatementOut(qq422016 qtio422016.Writer, s *database.DebugStatement) {
//line views/vdatabase/Statement.html:119
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Statement.html:119
	streamstatementOut(qw422016, s)
//line views/vdatabase/Statement.html:119
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Statement.html:119
}

//line views/vdatabase/Statement.html:119
func statementOut(s *database.DebugStatement) string {
//line views/vdatabase/Statement.html:119
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Statement.html:119
	writestatementOut(qb422016, s)
//line views/vdatabase/Statement.html:119
	qs422016 := string(qb422016.B)
//line views/vdatabase/Statement.html:119
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Statement.html:119
	return qs422016
//line views/vdatabase/Statement.html:119
}
