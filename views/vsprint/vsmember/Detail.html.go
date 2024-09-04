// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsprint/vsmember/Detail.html:1
package vsmember

//line views/vsprint/vsmember/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vsmember/Detail.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/Detail.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/Detail.html:13
type Detail struct {
	layout.Basic
	Model            *smember.SprintMember
	SprintBySprintID *sprint.Sprint
	UserByUserID     *user.User
	Paths            []string
}

//line views/vsprint/vsmember/Detail.html:21
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Detail.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprintMember"><button type="button">`)
//line views/vsprint/vsmember/Detail.html:24
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vsprint/vsmember/Detail.html:24
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vsprint/vsmember/Detail.html:25
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vsprint/vsmember/Detail.html:25
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/vsmember/Detail.html:25
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vsprint/vsmember/Detail.html:25
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/vsmember/Detail.html:27
	components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vsprint/vsmember/Detail.html:27
	qw422016.N().S(` `)
//line views/vsprint/vsmember/Detail.html:27
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/vsmember/Detail.html:27
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vsprint/vsmember/Detail.html:28
	qw422016.E().S(smember.Route(p.Paths...))
//line views/vsprint/vsmember/Detail.html:28
	qw422016.N().S(`"><em>Member</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Sprint ID</th>
            <td class="nowrap">
              `)
//line views/vsprint/vsmember/Detail.html:35
	view.StreamUUID(qw422016, &p.Model.SprintID)
//line views/vsprint/vsmember/Detail.html:35
	if p.SprintBySprintID != nil {
//line views/vsprint/vsmember/Detail.html:35
		qw422016.N().S(` (`)
//line views/vsprint/vsmember/Detail.html:35
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vsprint/vsmember/Detail.html:35
		qw422016.N().S(`)`)
//line views/vsprint/vsmember/Detail.html:35
	}
//line views/vsprint/vsmember/Detail.html:35
	qw422016.N().S(`
              <a title="Sprint" href="`)
//line views/vsprint/vsmember/Detail.html:36
	if x := p.SprintBySprintID; x != nil {
//line views/vsprint/vsmember/Detail.html:36
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vsprint/vsmember/Detail.html:36
	}
//line views/vsprint/vsmember/Detail.html:36
	qw422016.N().S(`">`)
//line views/vsprint/vsmember/Detail.html:36
	components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vsprint/vsmember/Detail.html:36
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
            <td class="nowrap">
              `)
//line views/vsprint/vsmember/Detail.html:42
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vsprint/vsmember/Detail.html:42
	if p.UserByUserID != nil {
//line views/vsprint/vsmember/Detail.html:42
		qw422016.N().S(` (`)
//line views/vsprint/vsmember/Detail.html:42
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vsprint/vsmember/Detail.html:42
		qw422016.N().S(`)`)
//line views/vsprint/vsmember/Detail.html:42
	}
//line views/vsprint/vsmember/Detail.html:42
	qw422016.N().S(`
              <a title="User" href="`)
//line views/vsprint/vsmember/Detail.html:43
	if x := p.UserByUserID; x != nil {
//line views/vsprint/vsmember/Detail.html:43
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vsprint/vsmember/Detail.html:43
	}
//line views/vsprint/vsmember/Detail.html:43
	qw422016.N().S(`">`)
//line views/vsprint/vsmember/Detail.html:43
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vsprint/vsmember/Detail.html:43
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Name</th>
            <td><strong>`)
//line views/vsprint/vsmember/Detail.html:48
	view.StreamString(qw422016, p.Model.Name)
//line views/vsprint/vsmember/Detail.html:48
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="URL in string form">Picture</th>
            <td><a href="`)
//line views/vsprint/vsmember/Detail.html:52
	qw422016.E().S(p.Model.Picture)
//line views/vsprint/vsmember/Detail.html:52
	qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vsprint/vsmember/Detail.html:52
	qw422016.E().S(p.Model.Picture)
//line views/vsprint/vsmember/Detail.html:52
	qw422016.N().S(`</a></td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vsprint/vsmember/Detail.html:55
	qw422016.E().S(enum.AllMemberStatuses.Help())
//line views/vsprint/vsmember/Detail.html:55
	qw422016.N().S(`">Role</th>
            <td>`)
//line views/vsprint/vsmember/Detail.html:56
	qw422016.E().S(p.Model.Role.String())
//line views/vsprint/vsmember/Detail.html:56
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vsprint/vsmember/Detail.html:60
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/vsmember/Detail.html:60
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vsprint/vsmember/Detail.html:64
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vsprint/vsmember/Detail.html:64
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vsprint/vsmember/Detail.html:71
	qw422016.N().S(`  `)
//line views/vsprint/vsmember/Detail.html:72
	components.StreamJSONModal(qw422016, "sprintMember", "Member JSON", p.Model, 1)
//line views/vsprint/vsmember/Detail.html:72
	qw422016.N().S(`
`)
//line views/vsprint/vsmember/Detail.html:73
}

//line views/vsprint/vsmember/Detail.html:73
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Detail.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/Detail.html:73
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vsmember/Detail.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/Detail.html:73
}

//line views/vsprint/vsmember/Detail.html:73
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vsmember/Detail.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/Detail.html:73
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vsmember/Detail.html:73
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/Detail.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/Detail.html:73
	return qs422016
//line views/vsprint/vsmember/Detail.html:73
}
