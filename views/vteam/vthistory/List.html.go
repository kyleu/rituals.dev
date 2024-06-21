// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vteam/vthistory/List.html:1
package vthistory

//line views/vteam/vthistory/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vteam/vthistory/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/vthistory/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/vthistory/List.html:11
type List struct {
	layout.Basic
	Models        thistory.TeamHistories
	TeamsByTeamID team.Teams
	Params        filter.ParamSet
}

//line views/vteam/vthistory/List.html:18
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/vthistory/List.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right mrs large-buttons">
`)
//line views/vteam/vthistory/List.html:21
	if len(p.Models) > 1 {
//line views/vteam/vthistory/List.html:21
		qw422016.N().S(`<a href="/admin/db/team/history/_random"><button>`)
//line views/vteam/vthistory/List.html:21
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vteam/vthistory/List.html:21
		qw422016.N().S(` Random</button></a>`)
//line views/vteam/vthistory/List.html:21
	}
//line views/vteam/vthistory/List.html:21
	qw422016.N().S(`      <a href="/admin/db/team/history/_new"><button>`)
//line views/vteam/vthistory/List.html:22
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vteam/vthistory/List.html:22
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vteam/vthistory/List.html:24
	components.StreamSVGIcon(qw422016, `history`, ps)
//line views/vteam/vthistory/List.html:24
	qw422016.N().S(` `)
//line views/vteam/vthistory/List.html:24
	qw422016.E().S(ps.Title)
//line views/vteam/vthistory/List.html:24
	qw422016.N().S(`</h3>
`)
//line views/vteam/vthistory/List.html:25
	if len(p.Models) == 0 {
//line views/vteam/vthistory/List.html:25
		qw422016.N().S(`    <div class="mt"><em>No histories available</em></div>
`)
//line views/vteam/vthistory/List.html:27
	} else {
//line views/vteam/vthistory/List.html:27
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vteam/vthistory/List.html:29
		StreamTable(qw422016, p.Models, p.TeamsByTeamID, p.Params, as, ps)
//line views/vteam/vthistory/List.html:29
		qw422016.N().S(`
    </div>
`)
//line views/vteam/vthistory/List.html:31
	}
//line views/vteam/vthistory/List.html:31
	qw422016.N().S(`  </div>
`)
//line views/vteam/vthistory/List.html:33
}

//line views/vteam/vthistory/List.html:33
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/vthistory/List.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/vthistory/List.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vteam/vthistory/List.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/vthistory/List.html:33
}

//line views/vteam/vthistory/List.html:33
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vteam/vthistory/List.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/vthistory/List.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vteam/vthistory/List.html:33
	qs422016 := string(qb422016.B)
//line views/vteam/vthistory/List.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/vthistory/List.html:33
	return qs422016
//line views/vteam/vthistory/List.html:33
}
