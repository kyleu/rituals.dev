// Code generated by qtc from "EstimateList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/EstimateList.html:1
package vworkspace

//line views/vworkspace/EstimateList.html:1
import (
	"fmt"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vworkspace/EstimateList.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/EstimateList.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/EstimateList.html:14
type EstimateList struct {
	layout.Basic
	Estimates estimate.Estimates
	Sprints   sprint.Sprints
	Teams     team.Teams
}

//line views/vworkspace/EstimateList.html:21
func (p *EstimateList) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/EstimateList.html:21
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/EstimateList.html:23
	components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vworkspace/EstimateList.html:23
	qw422016.N().D(len(p.Estimates))
//line views/vworkspace/EstimateList.html:23
	qw422016.N().S(` `)
//line views/vworkspace/EstimateList.html:23
	qw422016.E().S(util.StringPluralMaybe("Estimate", len(p.Estimates)))
//line views/vworkspace/EstimateList.html:23
	qw422016.N().S(`</h3>
    <em>Planning poker for any stories you need to work on</em>
    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/EstimateList.html:27
	for _, e := range p.Estimates {
//line views/vworkspace/EstimateList.html:27
		qw422016.N().S(`        <tr>
          <td><a href="/estimate/`)
//line views/vworkspace/EstimateList.html:29
		qw422016.E().S(e.Slug)
//line views/vworkspace/EstimateList.html:29
		qw422016.N().S(`">`)
//line views/vworkspace/EstimateList.html:29
		qw422016.E().S(e.TitleString())
//line views/vworkspace/EstimateList.html:29
		qw422016.N().S(`</a></td>
          <td style="text-align: right;">`)
//line views/vworkspace/EstimateList.html:30
		components.StreamDisplayTimestamp(qw422016, &e.Created)
//line views/vworkspace/EstimateList.html:30
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vworkspace/EstimateList.html:32
	}
//line views/vworkspace/EstimateList.html:32
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkspace/EstimateList.html:37
	components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vworkspace/EstimateList.html:37
	qw422016.N().S(`New Estimate</h3>
    `)
//line views/vworkspace/EstimateList.html:38
	StreamEstimateForm(qw422016, &estimate.Estimate{}, p.Teams, p.Sprints, as, ps)
//line views/vworkspace/EstimateList.html:38
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/EstimateList.html:40
}

//line views/vworkspace/EstimateList.html:40
func (p *EstimateList) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/EstimateList.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/EstimateList.html:40
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/EstimateList.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/EstimateList.html:40
}

//line views/vworkspace/EstimateList.html:40
func (p *EstimateList) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/EstimateList.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/EstimateList.html:40
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/EstimateList.html:40
	qs422016 := string(qb422016.B)
//line views/vworkspace/EstimateList.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/EstimateList.html:40
	return qs422016
//line views/vworkspace/EstimateList.html:40
}

//line views/vworkspace/EstimateList.html:42
func StreamEstimateForm(qw422016 *qt422016.Writer, e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/EstimateList.html:42
	qw422016.N().S(`
  <form action="" method="post">
    <table class="mt expanded">
      <tbody>
        `)
//line views/vworkspace/EstimateList.html:46
	components.StreamTableInput(qw422016, "title", "Team Title", e.Title, 5, "The name of your estimate")
//line views/vworkspace/EstimateList.html:46
	qw422016.N().S(`
        `)
//line views/vworkspace/EstimateList.html:47
	components.StreamTableInput(qw422016, "name", "Your Name", ps.Profile.Name, 5, "Whatever you prefer to be called")
//line views/vworkspace/EstimateList.html:47
	qw422016.N().S(`
        `)
//line views/vworkspace/EstimateList.html:48
	components.StreamTableSelect(qw422016, "team", "Team", fmt.Sprint(e.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5, "The team associated to this estimate")
//line views/vworkspace/EstimateList.html:48
	qw422016.N().S(`
        `)
//line views/vworkspace/EstimateList.html:49
	components.StreamTableSelect(qw422016, "sprint", "Sprint", fmt.Sprint(e.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5, "The sprint associated to this estimate")
//line views/vworkspace/EstimateList.html:49
	qw422016.N().S(`
        <tr><td colspan="2"><button type="submit">Add Estimate</button></td></tr>
      </tbody>
    </table>
  </form>
`)
//line views/vworkspace/EstimateList.html:54
}

//line views/vworkspace/EstimateList.html:54
func WriteEstimateForm(qq422016 qtio422016.Writer, e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/EstimateList.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/EstimateList.html:54
	StreamEstimateForm(qw422016, e, teams, sprints, as, ps)
//line views/vworkspace/EstimateList.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/EstimateList.html:54
}

//line views/vworkspace/EstimateList.html:54
func EstimateForm(e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/EstimateList.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/EstimateList.html:54
	WriteEstimateForm(qb422016, e, teams, sprints, as, ps)
//line views/vworkspace/EstimateList.html:54
	qs422016 := string(qb422016.B)
//line views/vworkspace/EstimateList.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/EstimateList.html:54
	return qs422016
//line views/vworkspace/EstimateList.html:54
}
