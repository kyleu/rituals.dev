// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/vreport/Detail.html:2
package vreport

//line views/vstandup/vreport/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vreport/Detail.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vreport/Detail.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vreport/Detail.html:13
type Detail struct {
	layout.Basic
	Model              *report.Report
	StandupByStandupID *standup.Standup
	UserByUserID       *user.User
}

//line views/vstandup/vreport/Detail.html:20
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vreport/Detail.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-report"><button type="button">JSON</button></a>
      <a href="`)
//line views/vstandup/vreport/Detail.html:24
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/vreport/Detail.html:24
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/vreport/Detail.html:24
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vstandup/vreport/Detail.html:24
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/vreport/Detail.html:26
	components.StreamSVGRefIcon(qw422016, `file-alt`, ps)
//line views/vstandup/vreport/Detail.html:26
	qw422016.N().S(` `)
//line views/vstandup/vreport/Detail.html:26
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/vreport/Detail.html:26
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup/report"><em>Report</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vstandup/vreport/Detail.html:32
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vstandup/vreport/Detail.html:32
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Standup ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/vreport/Detail.html:37
	view.StreamUUID(qw422016, &p.Model.StandupID)
//line views/vstandup/vreport/Detail.html:37
	if p.StandupByStandupID != nil {
//line views/vstandup/vreport/Detail.html:37
		qw422016.N().S(` (`)
//line views/vstandup/vreport/Detail.html:37
		qw422016.E().S(p.StandupByStandupID.TitleString())
//line views/vstandup/vreport/Detail.html:37
		qw422016.N().S(`)`)
//line views/vstandup/vreport/Detail.html:37
	}
//line views/vstandup/vreport/Detail.html:37
	qw422016.N().S(`
            <a title="Standup" href="`)
//line views/vstandup/vreport/Detail.html:38
	qw422016.E().S(`/admin/db/standup` + `/` + p.Model.StandupID.String())
//line views/vstandup/vreport/Detail.html:38
	qw422016.N().S(`">`)
//line views/vstandup/vreport/Detail.html:38
	components.StreamSVGRef(qw422016, "standup", 18, 18, "", ps)
//line views/vstandup/vreport/Detail.html:38
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Calendar date">Day</th>
          <td>`)
//line views/vstandup/vreport/Detail.html:43
	view.StreamTimestampDay(qw422016, &p.Model.Day)
//line views/vstandup/vreport/Detail.html:43
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/vreport/Detail.html:48
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vstandup/vreport/Detail.html:48
	if p.UserByUserID != nil {
//line views/vstandup/vreport/Detail.html:48
		qw422016.N().S(` (`)
//line views/vstandup/vreport/Detail.html:48
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vstandup/vreport/Detail.html:48
		qw422016.N().S(`)`)
//line views/vstandup/vreport/Detail.html:48
	}
//line views/vstandup/vreport/Detail.html:48
	qw422016.N().S(`
            <a title="User" href="`)
//line views/vstandup/vreport/Detail.html:49
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vstandup/vreport/Detail.html:49
	qw422016.N().S(`">`)
//line views/vstandup/vreport/Detail.html:49
	components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vstandup/vreport/Detail.html:49
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Content</th>
          <td>`)
//line views/vstandup/vreport/Detail.html:54
	view.StreamString(qw422016, p.Model.Content)
//line views/vstandup/vreport/Detail.html:54
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">HTML</th>
          <td>`)
//line views/vstandup/vreport/Detail.html:58
	view.StreamFormat(qw422016, p.Model.HTML, "html")
//line views/vstandup/vreport/Detail.html:58
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vstandup/vreport/Detail.html:62
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/vreport/Detail.html:62
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vstandup/vreport/Detail.html:66
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vstandup/vreport/Detail.html:66
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vstandup/vreport/Detail.html:72
	qw422016.N().S(`  `)
//line views/vstandup/vreport/Detail.html:73
	components.StreamJSONModal(qw422016, "report", "Report JSON", p.Model, 1)
//line views/vstandup/vreport/Detail.html:73
	qw422016.N().S(`
`)
//line views/vstandup/vreport/Detail.html:74
}

//line views/vstandup/vreport/Detail.html:74
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vreport/Detail.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vreport/Detail.html:74
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vreport/Detail.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vreport/Detail.html:74
}

//line views/vstandup/vreport/Detail.html:74
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vreport/Detail.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vreport/Detail.html:74
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vreport/Detail.html:74
	qs422016 := string(qb422016.B)
//line views/vstandup/vreport/Detail.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vreport/Detail.html:74
	return qs422016
//line views/vstandup/vreport/Detail.html:74
}
