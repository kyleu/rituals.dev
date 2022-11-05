// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrmember/Detail.html:2
package vrmember

//line views/vretro/vrmember/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrmember/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrmember/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrmember/Detail.html:12
type Detail struct {
	layout.Basic
	Model  *rmember.RetroMember
	Retros retro.Retros
	Users  user.Users
}

//line views/vretro/vrmember/Detail.html:19
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Detail.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-retroMember"><button type="button">JSON</button></a>
      <a href="`)
//line views/vretro/vrmember/Detail.html:23
	qw422016.E().S(p.Model.WebPath())
//line views/vretro/vrmember/Detail.html:23
	qw422016.N().S(`/edit"><button>`)
//line views/vretro/vrmember/Detail.html:23
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vretro/vrmember/Detail.html:23
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vretro/vrmember/Detail.html:25
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vretro/vrmember/Detail.html:25
	qw422016.N().S(` `)
//line views/vretro/vrmember/Detail.html:25
	qw422016.E().S(p.Model.TitleString())
//line views/vretro/vrmember/Detail.html:25
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/retro/member"><em>Member</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Retro ID</th>
          <td>
            <div class="icon">`)
//line views/vretro/vrmember/Detail.html:32
	components.StreamDisplayUUID(qw422016, &p.Model.RetroID)
//line views/vretro/vrmember/Detail.html:32
	if x := p.Retros.Get(p.Model.RetroID); x != nil {
//line views/vretro/vrmember/Detail.html:32
		qw422016.N().S(` (`)
//line views/vretro/vrmember/Detail.html:32
		qw422016.E().S(x.TitleString())
//line views/vretro/vrmember/Detail.html:32
		qw422016.N().S(`)`)
//line views/vretro/vrmember/Detail.html:32
	}
//line views/vretro/vrmember/Detail.html:32
	qw422016.N().S(`</div>
            <a title="Retro" href="`)
//line views/vretro/vrmember/Detail.html:33
	qw422016.E().S(`/retro` + `/` + p.Model.RetroID.String())
//line views/vretro/vrmember/Detail.html:33
	qw422016.N().S(`">`)
//line views/vretro/vrmember/Detail.html:33
	components.StreamSVGRefIcon(qw422016, "retro", ps)
//line views/vretro/vrmember/Detail.html:33
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td>
            <div class="icon">`)
//line views/vretro/vrmember/Detail.html:39
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vretro/vrmember/Detail.html:39
	if x := p.Users.Get(p.Model.UserID); x != nil {
//line views/vretro/vrmember/Detail.html:39
		qw422016.N().S(` (`)
//line views/vretro/vrmember/Detail.html:39
		qw422016.E().S(x.TitleString())
//line views/vretro/vrmember/Detail.html:39
		qw422016.N().S(`)`)
//line views/vretro/vrmember/Detail.html:39
	}
//line views/vretro/vrmember/Detail.html:39
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vretro/vrmember/Detail.html:40
	qw422016.E().S(`/user` + `/` + p.Model.UserID.String())
//line views/vretro/vrmember/Detail.html:40
	qw422016.N().S(`">`)
//line views/vretro/vrmember/Detail.html:40
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vretro/vrmember/Detail.html:40
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Name</th>
          <td><strong>`)
//line views/vretro/vrmember/Detail.html:45
	qw422016.E().S(p.Model.Name)
//line views/vretro/vrmember/Detail.html:45
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="URL in string form">Picture</th>
          <td><a href="`)
//line views/vretro/vrmember/Detail.html:49
	qw422016.E().S(p.Model.Picture)
//line views/vretro/vrmember/Detail.html:49
	qw422016.N().S(`" target="_blank">`)
//line views/vretro/vrmember/Detail.html:49
	qw422016.E().S(p.Model.Picture)
//line views/vretro/vrmember/Detail.html:49
	qw422016.N().S(`</a></td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [owner, member, observer]">Role</th>
          <td>`)
//line views/vretro/vrmember/Detail.html:53
	qw422016.E().V(p.Model.Role)
//line views/vretro/vrmember/Detail.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vretro/vrmember/Detail.html:57
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vretro/vrmember/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vretro/vrmember/Detail.html:61
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vretro/vrmember/Detail.html:61
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vretro/vrmember/Detail.html:67
	qw422016.N().S(`  `)
//line views/vretro/vrmember/Detail.html:68
	components.StreamJSONModal(qw422016, "retroMember", "Member JSON", p.Model, 1)
//line views/vretro/vrmember/Detail.html:68
	qw422016.N().S(`
`)
//line views/vretro/vrmember/Detail.html:69
}

//line views/vretro/vrmember/Detail.html:69
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Detail.html:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrmember/Detail.html:69
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrmember/Detail.html:69
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrmember/Detail.html:69
}

//line views/vretro/vrmember/Detail.html:69
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrmember/Detail.html:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrmember/Detail.html:69
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrmember/Detail.html:69
	qs422016 := string(qb422016.B)
//line views/vretro/vrmember/Detail.html:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrmember/Detail.html:69
	return qs422016
//line views/vretro/vrmember/Detail.html:69
}
