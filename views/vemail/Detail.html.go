// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vemail/Detail.html:1
package vemail

//line views/vemail/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
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
	Model        *email.Email
	UserByUserID *user.User
	Paths        []string
}

//line views/vemail/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-email"><button type="button">`)
//line views/vemail/Detail.html:21
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vemail/Detail.html:21
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vemail/Detail.html:22
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vemail/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vemail/Detail.html:22
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vemail/Detail.html:22
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vemail/Detail.html:24
	components.StreamSVGIcon(qw422016, `email`, ps)
//line views/vemail/Detail.html:24
	qw422016.N().S(` `)
//line views/vemail/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vemail/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vemail/Detail.html:25
	qw422016.E().S(email.Route(p.Paths...))
//line views/vemail/Detail.html:25
	qw422016.N().S(`"><em>Email</em></a></div>
    `)
//line views/vemail/Detail.html:26
	StreamDetailTable(qw422016, p, ps)
//line views/vemail/Detail.html:26
	qw422016.N().S(`
  </div>
`)
//line views/vemail/Detail.html:29
	qw422016.N().S(`  `)
//line views/vemail/Detail.html:30
	components.StreamJSONModal(qw422016, "email", "Email JSON", p.Model, 1)
//line views/vemail/Detail.html:30
	qw422016.N().S(`
`)
//line views/vemail/Detail.html:31
}

//line views/vemail/Detail.html:31
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/Detail.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/Detail.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vemail/Detail.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/Detail.html:31
}

//line views/vemail/Detail.html:31
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vemail/Detail.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/Detail.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vemail/Detail.html:31
	qs422016 := string(qb422016.B)
//line views/vemail/Detail.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/Detail.html:31
	return qs422016
//line views/vemail/Detail.html:31
}

//line views/vemail/Detail.html:33
func StreamDetailTable(qw422016 *qt422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vemail/Detail.html:33
	qw422016.N().S(`
  <div class="mt overflow full-width">
    <table>
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vemail/Detail.html:39
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vemail/Detail.html:39
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Comma-separated list of values">Recipients</th>
          <td>`)
//line views/vemail/Detail.html:43
	view.StreamStringArray(qw422016, p.Model.Recipients)
//line views/vemail/Detail.html:43
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Subject</th>
          <td>`)
//line views/vemail/Detail.html:47
	view.StreamString(qw422016, p.Model.Subject)
//line views/vemail/Detail.html:47
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="JSON object">Data</th>
          <td>`)
//line views/vemail/Detail.html:51
	components.StreamJSON(qw422016, p.Model.Data)
//line views/vemail/Detail.html:51
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Plain</th>
          <td>`)
//line views/vemail/Detail.html:55
	view.StreamString(qw422016, p.Model.Plain)
//line views/vemail/Detail.html:55
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="HTML code, in string form">HTML</th>
          <td>`)
//line views/vemail/Detail.html:59
	view.StreamFormat(qw422016, p.Model.HTML, "html")
//line views/vemail/Detail.html:59
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td class="nowrap">
            `)
//line views/vemail/Detail.html:64
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vemail/Detail.html:64
	if p.UserByUserID != nil {
//line views/vemail/Detail.html:64
		qw422016.N().S(` (`)
//line views/vemail/Detail.html:64
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vemail/Detail.html:64
		qw422016.N().S(`)`)
//line views/vemail/Detail.html:64
	}
//line views/vemail/Detail.html:64
	qw422016.N().S(`
            <a title="User" href="`)
//line views/vemail/Detail.html:65
	if x := p.UserByUserID; x != nil {
//line views/vemail/Detail.html:65
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vemail/Detail.html:65
	}
//line views/vemail/Detail.html:65
	qw422016.N().S(`">`)
//line views/vemail/Detail.html:65
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vemail/Detail.html:65
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Status</th>
          <td>`)
//line views/vemail/Detail.html:70
	view.StreamString(qw422016, p.Model.Status)
//line views/vemail/Detail.html:70
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vemail/Detail.html:74
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vemail/Detail.html:74
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vemail/Detail.html:79
}

//line views/vemail/Detail.html:79
func WriteDetailTable(qq422016 qtio422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vemail/Detail.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/Detail.html:79
	StreamDetailTable(qw422016, p, ps)
//line views/vemail/Detail.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/Detail.html:79
}

//line views/vemail/Detail.html:79
func DetailTable(p *Detail, ps *cutil.PageState) string {
//line views/vemail/Detail.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/Detail.html:79
	WriteDetailTable(qb422016, p, ps)
//line views/vemail/Detail.html:79
	qs422016 := string(qb422016.B)
//line views/vemail/Detail.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/Detail.html:79
	return qs422016
//line views/vemail/Detail.html:79
}
