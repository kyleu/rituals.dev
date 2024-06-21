// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vuser/List.html:1
package vuser

//line views/vuser/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vuser/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vuser/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vuser/List.html:11
type List struct {
	layout.Basic
	Models      user.Users
	Params      filter.ParamSet
	SearchQuery string
}

//line views/vuser/List.html:18
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/List.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vuser/List.html:20
	edit.StreamSearchForm(qw422016, "", "q", "Search Users", p.SearchQuery, ps)
//line views/vuser/List.html:20
	qw422016.N().S(`</div>
    <div class="right mrs large-buttons">
`)
//line views/vuser/List.html:22
	if len(p.Models) > 1 {
//line views/vuser/List.html:22
		qw422016.N().S(`<a href="/admin/db/user/_random"><button>`)
//line views/vuser/List.html:22
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vuser/List.html:22
		qw422016.N().S(` Random</button></a>`)
//line views/vuser/List.html:22
	}
//line views/vuser/List.html:22
	qw422016.N().S(`      <a href="/admin/db/user/_new"><button>`)
//line views/vuser/List.html:23
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vuser/List.html:23
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vuser/List.html:25
	components.StreamSVGIcon(qw422016, `profile`, ps)
//line views/vuser/List.html:25
	qw422016.N().S(` `)
//line views/vuser/List.html:25
	qw422016.E().S(ps.Title)
//line views/vuser/List.html:25
	qw422016.N().S(`</h3>
    <div class="clear"></div>
`)
//line views/vuser/List.html:27
	if p.SearchQuery != "" {
//line views/vuser/List.html:27
		qw422016.N().S(`    <hr />
    <em>Search results for [`)
//line views/vuser/List.html:29
		qw422016.E().S(p.SearchQuery)
//line views/vuser/List.html:29
		qw422016.N().S(`]</em> (<a href="?">clear</a>)
`)
//line views/vuser/List.html:30
	}
//line views/vuser/List.html:31
	if len(p.Models) == 0 {
//line views/vuser/List.html:31
		qw422016.N().S(`    <div class="mt"><em>No users available</em></div>
`)
//line views/vuser/List.html:33
	} else {
//line views/vuser/List.html:33
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vuser/List.html:35
		StreamTable(qw422016, p.Models, p.Params, as, ps)
//line views/vuser/List.html:35
		qw422016.N().S(`
    </div>
`)
//line views/vuser/List.html:37
	}
//line views/vuser/List.html:37
	qw422016.N().S(`  </div>
`)
//line views/vuser/List.html:39
}

//line views/vuser/List.html:39
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/List.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vuser/List.html:39
	p.StreamBody(qw422016, as, ps)
//line views/vuser/List.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vuser/List.html:39
}

//line views/vuser/List.html:39
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vuser/List.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vuser/List.html:39
	p.WriteBody(qb422016, as, ps)
//line views/vuser/List.html:39
	qs422016 := string(qb422016.B)
//line views/vuser/List.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vuser/List.html:39
	return qs422016
//line views/vuser/List.html:39
}
