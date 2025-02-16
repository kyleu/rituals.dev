// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vretro/vrmember/Table.html:1
package vrmember

//line views/vretro/vrmember/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vretro/vrmember/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrmember/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrmember/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models rmember.RetroMembers, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vretro/vrmember/Table.html:13
	qw422016.N().S(`
`)
//line views/vretro/vrmember/Table.html:14
	prms := params.Sanitized("rmember", ps.Logger)

//line views/vretro/vrmember/Table.html:14
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vretro/vrmember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "rmember", "retro_id", "Retro ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:19
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "rmember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:20
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "rmember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:21
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "rmember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:22
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "rmember", "role", "Role", enum.AllMemberStatuses.Help(), prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:23
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "rmember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:24
	qw422016.N().S(`
          `)
//line views/vretro/vrmember/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "rmember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vretro/vrmember/Table.html:25
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vretro/vrmember/Table.html:29
	for _, model := range models {
//line views/vretro/vrmember/Table.html:29
		qw422016.N().S(`        <tr>
          <td class="nowrap">
            `)
//line views/vretro/vrmember/Table.html:32
		if x := retrosByRetroID.Get(model.RetroID); x != nil {
//line views/vretro/vrmember/Table.html:32
			qw422016.N().S(`
            `)
//line views/vretro/vrmember/Table.html:33
			qw422016.E().S(x.TitleString())
//line views/vretro/vrmember/Table.html:33
			qw422016.N().S(` <a title="Retro" href="`)
//line views/vretro/vrmember/Table.html:33
			qw422016.E().S(x.WebPath(paths...))
//line views/vretro/vrmember/Table.html:33
			qw422016.N().S(`">`)
//line views/vretro/vrmember/Table.html:33
			components.StreamSVGLink(qw422016, `retro`, ps)
//line views/vretro/vrmember/Table.html:33
			qw422016.N().S(`</a>
            `)
//line views/vretro/vrmember/Table.html:34
		} else {
//line views/vretro/vrmember/Table.html:34
			qw422016.N().S(`
            `)
//line views/vretro/vrmember/Table.html:35
			view.StreamUUID(qw422016, &model.RetroID)
//line views/vretro/vrmember/Table.html:35
			qw422016.N().S(`
            `)
//line views/vretro/vrmember/Table.html:36
		}
//line views/vretro/vrmember/Table.html:36
		qw422016.N().S(`
          </td>
          <td class="nowrap">
            `)
//line views/vretro/vrmember/Table.html:39
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vretro/vrmember/Table.html:39
			qw422016.N().S(`
            `)
//line views/vretro/vrmember/Table.html:40
			qw422016.E().S(x.TitleString())
//line views/vretro/vrmember/Table.html:40
			qw422016.N().S(` <a title="User" href="`)
//line views/vretro/vrmember/Table.html:40
			qw422016.E().S(x.WebPath(paths...))
//line views/vretro/vrmember/Table.html:40
			qw422016.N().S(`">`)
//line views/vretro/vrmember/Table.html:40
			components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vretro/vrmember/Table.html:40
			qw422016.N().S(`</a>
            `)
//line views/vretro/vrmember/Table.html:41
		} else {
//line views/vretro/vrmember/Table.html:41
			qw422016.N().S(`
            `)
//line views/vretro/vrmember/Table.html:42
			view.StreamUUID(qw422016, &model.UserID)
//line views/vretro/vrmember/Table.html:42
			qw422016.N().S(`
            `)
//line views/vretro/vrmember/Table.html:43
		}
//line views/vretro/vrmember/Table.html:43
		qw422016.N().S(`
          </td>
          <td><strong>`)
//line views/vretro/vrmember/Table.html:45
		view.StreamString(qw422016, model.Name)
//line views/vretro/vrmember/Table.html:45
		qw422016.N().S(`</strong></td>
          <td><a href="`)
//line views/vretro/vrmember/Table.html:46
		qw422016.E().S(model.Picture)
//line views/vretro/vrmember/Table.html:46
		qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vretro/vrmember/Table.html:46
		qw422016.E().S(model.Picture)
//line views/vretro/vrmember/Table.html:46
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vretro/vrmember/Table.html:47
		qw422016.E().S(model.Role.String())
//line views/vretro/vrmember/Table.html:47
		qw422016.N().S(`</td>
          <td>`)
//line views/vretro/vrmember/Table.html:48
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vretro/vrmember/Table.html:48
		qw422016.N().S(`</td>
          <td>`)
//line views/vretro/vrmember/Table.html:49
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vretro/vrmember/Table.html:49
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vretro/vrmember/Table.html:51
	}
//line views/vretro/vrmember/Table.html:51
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vretro/vrmember/Table.html:55
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/vrmember/Table.html:55
		qw422016.N().S(`  <hr />
  `)
//line views/vretro/vrmember/Table.html:57
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/vrmember/Table.html:57
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vretro/vrmember/Table.html:59
	}
//line views/vretro/vrmember/Table.html:60
}

//line views/vretro/vrmember/Table.html:60
func WriteTable(qq422016 qtio422016.Writer, models rmember.RetroMembers, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vretro/vrmember/Table.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrmember/Table.html:60
	StreamTable(qw422016, models, retrosByRetroID, usersByUserID, params, as, ps, paths...)
//line views/vretro/vrmember/Table.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrmember/Table.html:60
}

//line views/vretro/vrmember/Table.html:60
func Table(models rmember.RetroMembers, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vretro/vrmember/Table.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrmember/Table.html:60
	WriteTable(qb422016, models, retrosByRetroID, usersByUserID, params, as, ps, paths...)
//line views/vretro/vrmember/Table.html:60
	qs422016 := string(qb422016.B)
//line views/vretro/vrmember/Table.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrmember/Table.html:60
	return qs422016
//line views/vretro/vrmember/Table.html:60
}
