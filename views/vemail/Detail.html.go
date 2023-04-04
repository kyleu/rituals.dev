// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vemail/Detail.html:2
package vemail

//line views/vemail/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vemail/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vemail/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vemail/Detail.html:11
type Detail struct {
	layout.Basic
	Model         *email.Email
	UsersByUserID user.Users
}

//line views/vemail/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-email"><button type="button">JSON</button></a>
      <a href="`)
//line views/vemail/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vemail/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vemail/Detail.html:21
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vemail/Detail.html:21
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vemail/Detail.html:23
	components.StreamSVGRefIcon(qw422016, `email`, ps)
//line views/vemail/Detail.html:23
	qw422016.N().S(` `)
//line views/vemail/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vemail/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/email"><em>Email</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vemail/Detail.html:29
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vemail/Detail.html:29
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Comma-separated list of values">Recipients</th>
          <td>`)
//line views/vemail/Detail.html:33
	components.StreamDisplayStringArray(qw422016, p.Model.Recipients)
//line views/vemail/Detail.html:33
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Subject</th>
          <td>`)
//line views/vemail/Detail.html:37
	qw422016.E().S(p.Model.Subject)
//line views/vemail/Detail.html:37
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="JSON object">Data</th>
          <td>`)
//line views/vemail/Detail.html:41
	components.StreamJSON(qw422016, p.Model.Data)
//line views/vemail/Detail.html:41
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Plain</th>
          <td>`)
//line views/vemail/Detail.html:45
	qw422016.E().S(p.Model.Plain)
//line views/vemail/Detail.html:45
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">HTML</th>
          <td><pre>`)
//line views/vemail/Detail.html:49
	qw422016.E().S(p.Model.HTML)
//line views/vemail/Detail.html:49
	qw422016.N().S(`</pre></td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td class="nowrap">
            `)
//line views/vemail/Detail.html:54
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vemail/Detail.html:54
	if x := p.UsersByUserID.Get(p.Model.UserID); x != nil {
//line views/vemail/Detail.html:54
		qw422016.N().S(` (`)
//line views/vemail/Detail.html:54
		qw422016.E().S(x.TitleString())
//line views/vemail/Detail.html:54
		qw422016.N().S(`)`)
//line views/vemail/Detail.html:54
	}
//line views/vemail/Detail.html:54
	qw422016.N().S(`
            <a title="User" href="`)
//line views/vemail/Detail.html:55
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vemail/Detail.html:55
	qw422016.N().S(`">`)
//line views/vemail/Detail.html:55
	components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vemail/Detail.html:55
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Status</th>
          <td>`)
//line views/vemail/Detail.html:60
	qw422016.E().S(p.Model.Status)
//line views/vemail/Detail.html:60
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vemail/Detail.html:64
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vemail/Detail.html:64
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vemail/Detail.html:70
	qw422016.N().S(`  `)
//line views/vemail/Detail.html:71
	components.StreamJSONModal(qw422016, "email", "Email JSON", p.Model, 1)
//line views/vemail/Detail.html:71
	qw422016.N().S(`
`)
//line views/vemail/Detail.html:72
}

//line views/vemail/Detail.html:72
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/Detail.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/Detail.html:72
	p.StreamBody(qw422016, as, ps)
//line views/vemail/Detail.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/Detail.html:72
}

//line views/vemail/Detail.html:72
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vemail/Detail.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/Detail.html:72
	p.WriteBody(qb422016, as, ps)
//line views/vemail/Detail.html:72
	qs422016 := string(qb422016.B)
//line views/vemail/Detail.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/Detail.html:72
	return qs422016
//line views/vemail/Detail.html:72
}
