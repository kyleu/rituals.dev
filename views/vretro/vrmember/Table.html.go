// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrmember/Table.html:2
package vrmember

//line views/vretro/vrmember/Table.html:2
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

//line views/vretro/vrmember/Table.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrmember/Table.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrmember/Table.html:14
func StreamTable(qw422016 *qt422016.Writer, models rmember.RetroMembers, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Table.html:14
	qw422016.N().S(`
`)
//line views/vretro/vrmember/Table.html:15
	prms := params.Sanitized("rmember", ps.Logger)

//line views/vretro/vrmember/Table.html:15
	qw422016.N().S(`  <table>
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
		qw422016.N().S(`      <tr>
        <td class="nowrap">
          <a href="/admin/db/retro/member/`)
//line views/vretro/vrmember/Table.html:32
		view.StreamUUID(qw422016, &model.RetroID)
//line views/vretro/vrmember/Table.html:32
		qw422016.N().S(`/`)
//line views/vretro/vrmember/Table.html:32
		view.StreamUUID(qw422016, &model.UserID)
//line views/vretro/vrmember/Table.html:32
		qw422016.N().S(`">`)
//line views/vretro/vrmember/Table.html:32
		view.StreamUUID(qw422016, &model.RetroID)
//line views/vretro/vrmember/Table.html:32
		if x := retrosByRetroID.Get(model.RetroID); x != nil {
//line views/vretro/vrmember/Table.html:32
			qw422016.N().S(` (`)
//line views/vretro/vrmember/Table.html:32
			qw422016.E().S(x.TitleString())
//line views/vretro/vrmember/Table.html:32
			qw422016.N().S(`)`)
//line views/vretro/vrmember/Table.html:32
		}
//line views/vretro/vrmember/Table.html:32
		qw422016.N().S(`</a>
          <a title="Retro" href="`)
//line views/vretro/vrmember/Table.html:33
		qw422016.E().S(`/admin/db/retro` + `/` + model.RetroID.String())
//line views/vretro/vrmember/Table.html:33
		qw422016.N().S(`">`)
//line views/vretro/vrmember/Table.html:33
		components.StreamSVGRef(qw422016, "retro", 18, 18, "", ps)
//line views/vretro/vrmember/Table.html:33
		qw422016.N().S(`</a>
        </td>
        <td class="nowrap">
          <a href="/admin/db/retro/member/`)
//line views/vretro/vrmember/Table.html:36
		view.StreamUUID(qw422016, &model.RetroID)
//line views/vretro/vrmember/Table.html:36
		qw422016.N().S(`/`)
//line views/vretro/vrmember/Table.html:36
		view.StreamUUID(qw422016, &model.UserID)
//line views/vretro/vrmember/Table.html:36
		qw422016.N().S(`">`)
//line views/vretro/vrmember/Table.html:36
		view.StreamUUID(qw422016, &model.UserID)
//line views/vretro/vrmember/Table.html:36
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vretro/vrmember/Table.html:36
			qw422016.N().S(` (`)
//line views/vretro/vrmember/Table.html:36
			qw422016.E().S(x.TitleString())
//line views/vretro/vrmember/Table.html:36
			qw422016.N().S(`)`)
//line views/vretro/vrmember/Table.html:36
		}
//line views/vretro/vrmember/Table.html:36
		qw422016.N().S(`</a>
          <a title="User" href="`)
//line views/vretro/vrmember/Table.html:37
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vretro/vrmember/Table.html:37
		qw422016.N().S(`">`)
//line views/vretro/vrmember/Table.html:37
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vretro/vrmember/Table.html:37
		qw422016.N().S(`</a>
        </td>
        <td><strong>`)
//line views/vretro/vrmember/Table.html:39
		view.StreamString(qw422016, model.Name)
//line views/vretro/vrmember/Table.html:39
		qw422016.N().S(`</strong></td>
        <td><a href="`)
//line views/vretro/vrmember/Table.html:40
		qw422016.E().S(model.Picture)
//line views/vretro/vrmember/Table.html:40
		qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vretro/vrmember/Table.html:40
		qw422016.E().S(model.Picture)
//line views/vretro/vrmember/Table.html:40
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vretro/vrmember/Table.html:41
		qw422016.E().S(model.Role.String())
//line views/vretro/vrmember/Table.html:41
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vrmember/Table.html:42
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vretro/vrmember/Table.html:42
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vrmember/Table.html:43
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vretro/vrmember/Table.html:43
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrmember/Table.html:45
	}
//line views/vretro/vrmember/Table.html:46
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/vrmember/Table.html:46
		qw422016.N().S(`      <tr>
        <td colspan="7">`)
//line views/vretro/vrmember/Table.html:48
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/vrmember/Table.html:48
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrmember/Table.html:50
	}
//line views/vretro/vrmember/Table.html:50
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vretro/vrmember/Table.html:53
}

//line views/vretro/vrmember/Table.html:53
func WriteTable(qq422016 qtio422016.Writer, models rmember.RetroMembers, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrmember/Table.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrmember/Table.html:53
	StreamTable(qw422016, models, retrosByRetroID, usersByUserID, params, as, ps)
//line views/vretro/vrmember/Table.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrmember/Table.html:53
}

//line views/vretro/vrmember/Table.html:53
func Table(models rmember.RetroMembers, retrosByRetroID retro.Retros, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrmember/Table.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrmember/Table.html:53
	WriteTable(qb422016, models, retrosByRetroID, usersByUserID, params, as, ps)
//line views/vretro/vrmember/Table.html:53
	qs422016 := string(qb422016.B)
//line views/vretro/vrmember/Table.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrmember/Table.html:53
	return qs422016
//line views/vretro/vrmember/Table.html:53
}
