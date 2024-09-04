// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vstandup/vupermission/Detail.html:1
package vupermission

//line views/vstandup/vupermission/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vstandup/vupermission/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/vupermission/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/vupermission/Detail.html:11
type Detail struct {
	layout.Basic
	Model              *upermission.StandupPermission
	StandupByStandupID *standup.Standup
	Paths              []string
}

//line views/vstandup/vupermission/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standupPermission"><button type="button">`)
//line views/vstandup/vupermission/Detail.html:21
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vstandup/vupermission/Detail.html:21
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vstandup/vupermission/Detail.html:22
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vstandup/vupermission/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/vupermission/Detail.html:22
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vstandup/vupermission/Detail.html:22
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/vupermission/Detail.html:24
	components.StreamSVGIcon(qw422016, `permission`, ps)
//line views/vstandup/vupermission/Detail.html:24
	qw422016.N().S(` `)
//line views/vstandup/vupermission/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/vupermission/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vstandup/vupermission/Detail.html:25
	qw422016.E().S(upermission.Route(p.Paths...))
//line views/vstandup/vupermission/Detail.html:25
	qw422016.N().S(`"><em>Permission</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Standup ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/vupermission/Detail.html:32
	view.StreamUUID(qw422016, &p.Model.StandupID)
//line views/vstandup/vupermission/Detail.html:32
	if p.StandupByStandupID != nil {
//line views/vstandup/vupermission/Detail.html:32
		qw422016.N().S(` (`)
//line views/vstandup/vupermission/Detail.html:32
		qw422016.E().S(p.StandupByStandupID.TitleString())
//line views/vstandup/vupermission/Detail.html:32
		qw422016.N().S(`)`)
//line views/vstandup/vupermission/Detail.html:32
	}
//line views/vstandup/vupermission/Detail.html:32
	qw422016.N().S(`
              <a title="Standup" href="`)
//line views/vstandup/vupermission/Detail.html:33
	if x := p.StandupByStandupID; x != nil {
//line views/vstandup/vupermission/Detail.html:33
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vstandup/vupermission/Detail.html:33
	}
//line views/vstandup/vupermission/Detail.html:33
	qw422016.N().S(`">`)
//line views/vstandup/vupermission/Detail.html:33
	components.StreamSVGLink(qw422016, `standup`, ps)
//line views/vstandup/vupermission/Detail.html:33
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Key</th>
            <td>`)
//line views/vstandup/vupermission/Detail.html:38
	view.StreamString(qw422016, p.Model.Key)
//line views/vstandup/vupermission/Detail.html:38
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Value</th>
            <td>`)
//line views/vstandup/vupermission/Detail.html:42
	view.StreamString(qw422016, p.Model.Value)
//line views/vstandup/vupermission/Detail.html:42
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Access</th>
            <td>`)
//line views/vstandup/vupermission/Detail.html:46
	view.StreamString(qw422016, p.Model.Access)
//line views/vstandup/vupermission/Detail.html:46
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vstandup/vupermission/Detail.html:50
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/vupermission/Detail.html:50
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vstandup/vupermission/Detail.html:57
	qw422016.N().S(`  `)
//line views/vstandup/vupermission/Detail.html:58
	components.StreamJSONModal(qw422016, "standupPermission", "Permission JSON", p.Model, 1)
//line views/vstandup/vupermission/Detail.html:58
	qw422016.N().S(`
`)
//line views/vstandup/vupermission/Detail.html:59
}

//line views/vstandup/vupermission/Detail.html:59
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/vupermission/Detail.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/vupermission/Detail.html:59
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/vupermission/Detail.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/vupermission/Detail.html:59
}

//line views/vstandup/vupermission/Detail.html:59
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/vupermission/Detail.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/vupermission/Detail.html:59
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/vupermission/Detail.html:59
	qs422016 := string(qb422016.B)
//line views/vstandup/vupermission/Detail.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/vupermission/Detail.html:59
	return qs422016
//line views/vstandup/vupermission/Detail.html:59
}
