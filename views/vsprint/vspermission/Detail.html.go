// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsprint/vspermission/Detail.html:1
package vspermission

//line views/vsprint/vspermission/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vspermission/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vspermission/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vspermission/Detail.html:11
type Detail struct {
	layout.Basic
	Model            *spermission.SprintPermission
	SprintBySprintID *sprint.Sprint
	Paths            []string
}

//line views/vsprint/vspermission/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vspermission/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprintPermission"><button type="button">`)
//line views/vsprint/vspermission/Detail.html:21
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vsprint/vspermission/Detail.html:21
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vsprint/vspermission/Detail.html:22
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vsprint/vspermission/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/vspermission/Detail.html:22
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vsprint/vspermission/Detail.html:22
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/vspermission/Detail.html:24
	components.StreamSVGIcon(qw422016, `permission`, ps)
//line views/vsprint/vspermission/Detail.html:24
	qw422016.N().S(` `)
//line views/vsprint/vspermission/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/vspermission/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vsprint/vspermission/Detail.html:25
	qw422016.E().S(spermission.Route(p.Paths...))
//line views/vsprint/vspermission/Detail.html:25
	qw422016.N().S(`"><em>Permission</em></a></div>
    `)
//line views/vsprint/vspermission/Detail.html:26
	StreamDetailTable(qw422016, p, ps)
//line views/vsprint/vspermission/Detail.html:26
	qw422016.N().S(`
  </div>
`)
//line views/vsprint/vspermission/Detail.html:29
	qw422016.N().S(`  `)
//line views/vsprint/vspermission/Detail.html:30
	components.StreamJSONModal(qw422016, "sprintPermission", "Permission JSON", p.Model, 1)
//line views/vsprint/vspermission/Detail.html:30
	qw422016.N().S(`
`)
//line views/vsprint/vspermission/Detail.html:31
}

//line views/vsprint/vspermission/Detail.html:31
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vspermission/Detail.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vspermission/Detail.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vspermission/Detail.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vspermission/Detail.html:31
}

//line views/vsprint/vspermission/Detail.html:31
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vspermission/Detail.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vspermission/Detail.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vspermission/Detail.html:31
	qs422016 := string(qb422016.B)
//line views/vsprint/vspermission/Detail.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vspermission/Detail.html:31
	return qs422016
//line views/vsprint/vspermission/Detail.html:31
}

//line views/vsprint/vspermission/Detail.html:33
func StreamDetailTable(qw422016 *qt422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vsprint/vspermission/Detail.html:33
	qw422016.N().S(`
  <div class="mt overflow full-width">
    <table>
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Sprint ID</th>
          <td class="nowrap">
            `)
//line views/vsprint/vspermission/Detail.html:40
	view.StreamUUID(qw422016, &p.Model.SprintID)
//line views/vsprint/vspermission/Detail.html:40
	if p.SprintBySprintID != nil {
//line views/vsprint/vspermission/Detail.html:40
		qw422016.N().S(` (`)
//line views/vsprint/vspermission/Detail.html:40
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vsprint/vspermission/Detail.html:40
		qw422016.N().S(`)`)
//line views/vsprint/vspermission/Detail.html:40
	}
//line views/vsprint/vspermission/Detail.html:40
	qw422016.N().S(`
            <a title="Sprint" href="`)
//line views/vsprint/vspermission/Detail.html:41
	if x := p.SprintBySprintID; x != nil {
//line views/vsprint/vspermission/Detail.html:41
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vsprint/vspermission/Detail.html:41
	}
//line views/vsprint/vspermission/Detail.html:41
	qw422016.N().S(`">`)
//line views/vsprint/vspermission/Detail.html:41
	components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vsprint/vspermission/Detail.html:41
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Key</th>
          <td>`)
//line views/vsprint/vspermission/Detail.html:46
	view.StreamString(qw422016, p.Model.Key)
//line views/vsprint/vspermission/Detail.html:46
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Value</th>
          <td>`)
//line views/vsprint/vspermission/Detail.html:50
	view.StreamString(qw422016, p.Model.Value)
//line views/vsprint/vspermission/Detail.html:50
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Access</th>
          <td>`)
//line views/vsprint/vspermission/Detail.html:54
	view.StreamString(qw422016, p.Model.Access)
//line views/vsprint/vspermission/Detail.html:54
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vsprint/vspermission/Detail.html:58
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/vspermission/Detail.html:58
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsprint/vspermission/Detail.html:63
}

//line views/vsprint/vspermission/Detail.html:63
func WriteDetailTable(qq422016 qtio422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vsprint/vspermission/Detail.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vspermission/Detail.html:63
	StreamDetailTable(qw422016, p, ps)
//line views/vsprint/vspermission/Detail.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vspermission/Detail.html:63
}

//line views/vsprint/vspermission/Detail.html:63
func DetailTable(p *Detail, ps *cutil.PageState) string {
//line views/vsprint/vspermission/Detail.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vspermission/Detail.html:63
	WriteDetailTable(qb422016, p, ps)
//line views/vsprint/vspermission/Detail.html:63
	qs422016 := string(qb422016.B)
//line views/vsprint/vspermission/Detail.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vspermission/Detail.html:63
	return qs422016
//line views/vsprint/vspermission/Detail.html:63
}
