// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/List.html:2
package vretro

//line views/vretro/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/List.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/List.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/List.html:13
type List struct {
	layout.Basic
	Models      retro.Retros
	Teams       team.Teams
	Sprints     sprint.Sprints
	Params      filter.ParamSet
	SearchQuery string
}

//line views/vretro/List.html:22
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/List.html:22
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vretro/List.html:24
	components.StreamSearchForm(qw422016, "", "q", "Search retros", p.SearchQuery, ps)
//line views/vretro/List.html:24
	qw422016.N().S(`</div>
    <h3>`)
//line views/vretro/List.html:25
	components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vretro/List.html:25
	qw422016.E().S(ps.Title)
//line views/vretro/List.html:25
	qw422016.N().S(` <a href="/admin/db/retro/new"><button>New</button></a></h3>
    <div class="clear"></div>
`)
//line views/vretro/List.html:27
	if p.SearchQuery != "" {
//line views/vretro/List.html:27
		qw422016.N().S(`    <em>Search results for [`)
//line views/vretro/List.html:28
		qw422016.E().S(p.SearchQuery)
//line views/vretro/List.html:28
		qw422016.N().S(`]</em>
`)
//line views/vretro/List.html:29
	}
//line views/vretro/List.html:30
	if len(p.Models) == 0 {
//line views/vretro/List.html:30
		qw422016.N().S(`    <div class="mt"><em>No retros available</em></div>
`)
//line views/vretro/List.html:32
	} else {
//line views/vretro/List.html:32
		qw422016.N().S(`    <div class="overflow clear">
      `)
//line views/vretro/List.html:34
		StreamTable(qw422016, p.Models, p.Teams, p.Sprints, p.Params, as, ps)
//line views/vretro/List.html:34
		qw422016.N().S(`
    </div>
`)
//line views/vretro/List.html:36
	}
//line views/vretro/List.html:36
	qw422016.N().S(`  </div>
`)
//line views/vretro/List.html:38
}

//line views/vretro/List.html:38
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/List.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/List.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vretro/List.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/List.html:38
}

//line views/vretro/List.html:38
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/List.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/List.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vretro/List.html:38
	qs422016 := string(qb422016.B)
//line views/vretro/List.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/List.html:38
	return qs422016
//line views/vretro/List.html:38
}
