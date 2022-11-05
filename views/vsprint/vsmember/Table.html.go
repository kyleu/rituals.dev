// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/vsmember/Table.html:2
package vsmember

//line views/vsprint/vsmember/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vsprint/vsmember/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models smember.SprintMembers, sprints sprint.Sprints, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Table.html:12
	qw422016.N().S(`
`)
//line views/vsprint/vsmember/Table.html:13
	prms := params.Get("smember", nil, ps.Logger).Sanitize("smember")

//line views/vsprint/vsmember/Table.html:13
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vsprint/vsmember/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "smember", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:17
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "smember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:18
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "smember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:19
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "smember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:20
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "smember", "role", "Role", "Available options: [owner, member, observer]", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:21
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "smember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:22
	qw422016.N().S(`
        `)
//line views/vsprint/vsmember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "smember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:23
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vsprint/vsmember/Table.html:27
	for _, model := range models {
//line views/vsprint/vsmember/Table.html:27
		qw422016.N().S(`      <tr>
        <td>
          <div class="icon"><a href="/admin/db/sprint/member/`)
//line views/vsprint/vsmember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:30
		qw422016.N().S(`/`)
//line views/vsprint/vsmember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:30
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:30
		components.StreamDisplayUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:30
		if x := sprints.Get(model.SprintID); x != nil {
//line views/vsprint/vsmember/Table.html:30
			qw422016.N().S(` (`)
//line views/vsprint/vsmember/Table.html:30
			qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Table.html:30
			qw422016.N().S(`)`)
//line views/vsprint/vsmember/Table.html:30
		}
//line views/vsprint/vsmember/Table.html:30
		qw422016.N().S(`</a></div>
          <a title="Sprint" href="`)
//line views/vsprint/vsmember/Table.html:31
		qw422016.E().S(`/sprint` + `/` + model.SprintID.String())
//line views/vsprint/vsmember/Table.html:31
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:31
		components.StreamSVGRefIcon(qw422016, "sprint", ps)
//line views/vsprint/vsmember/Table.html:31
		qw422016.N().S(`</a>
        </td>
        <td>
          <div class="icon"><a href="/admin/db/sprint/member/`)
//line views/vsprint/vsmember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:34
		qw422016.N().S(`/`)
//line views/vsprint/vsmember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:34
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:34
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:34
		if x := users.Get(model.UserID); x != nil {
//line views/vsprint/vsmember/Table.html:34
			qw422016.N().S(` (`)
//line views/vsprint/vsmember/Table.html:34
			qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Table.html:34
			qw422016.N().S(`)`)
//line views/vsprint/vsmember/Table.html:34
		}
//line views/vsprint/vsmember/Table.html:34
		qw422016.N().S(`</a></div>
          <a title="User" href="`)
//line views/vsprint/vsmember/Table.html:35
		qw422016.E().S(`/user` + `/` + model.UserID.String())
//line views/vsprint/vsmember/Table.html:35
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:35
		components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vsprint/vsmember/Table.html:35
		qw422016.N().S(`</a>
        </td>
        <td><strong>`)
//line views/vsprint/vsmember/Table.html:37
		qw422016.E().S(model.Name)
//line views/vsprint/vsmember/Table.html:37
		qw422016.N().S(`</strong></td>
        <td><a href="`)
//line views/vsprint/vsmember/Table.html:38
		qw422016.E().S(model.Picture)
//line views/vsprint/vsmember/Table.html:38
		qw422016.N().S(`" target="_blank">`)
//line views/vsprint/vsmember/Table.html:38
		qw422016.E().S(model.Picture)
//line views/vsprint/vsmember/Table.html:38
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vsprint/vsmember/Table.html:39
		qw422016.E().V(model.Role)
//line views/vsprint/vsmember/Table.html:39
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/vsmember/Table.html:40
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vsprint/vsmember/Table.html:40
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/vsmember/Table.html:41
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vsprint/vsmember/Table.html:41
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vsprint/vsmember/Table.html:43
	}
//line views/vsprint/vsmember/Table.html:44
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vsprint/vsmember/Table.html:44
		qw422016.N().S(`      <tr>
        <td colspan="7">`)
//line views/vsprint/vsmember/Table.html:46
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vsprint/vsmember/Table.html:46
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vsprint/vsmember/Table.html:48
	}
//line views/vsprint/vsmember/Table.html:48
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vsprint/vsmember/Table.html:51
}

//line views/vsprint/vsmember/Table.html:51
func WriteTable(qq422016 qtio422016.Writer, models smember.SprintMembers, sprints sprint.Sprints, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vsmember/Table.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/Table.html:51
	StreamTable(qw422016, models, sprints, users, params, as, ps)
//line views/vsprint/vsmember/Table.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/Table.html:51
}

//line views/vsprint/vsmember/Table.html:51
func Table(models smember.SprintMembers, sprints sprint.Sprints, users user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vsmember/Table.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/Table.html:51
	WriteTable(qb422016, models, sprints, users, params, as, ps)
//line views/vsprint/vsmember/Table.html:51
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/Table.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/Table.html:51
	return qs422016
//line views/vsprint/vsmember/Table.html:51
}
