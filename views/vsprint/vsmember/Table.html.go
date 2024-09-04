// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsprint/vsmember/Table.html:1
package vsmember

//line views/vsprint/vsmember/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vsprint/vsmember/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vsmember/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vsmember/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models smember.SprintMembers, sprintsBySprintID sprint.Sprints, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vsprint/vsmember/Table.html:13
	qw422016.N().S(`
`)
//line views/vsprint/vsmember/Table.html:14
	prms := params.Sanitized("smember", ps.Logger)

//line views/vsprint/vsmember/Table.html:14
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vsprint/vsmember/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "smember", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:19
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "smember", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:20
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "smember", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:21
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "smember", "picture", "Picture", "URL in string form", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:22
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "smember", "role", "Role", enum.AllMemberStatuses.Help(), prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:23
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "smember", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:24
	qw422016.N().S(`
          `)
//line views/vsprint/vsmember/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "smember", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vsprint/vsmember/Table.html:25
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vsprint/vsmember/Table.html:29
	for _, model := range models {
//line views/vsprint/vsmember/Table.html:29
		qw422016.N().S(`        <tr>
          <td class="nowrap">
            <a href="`)
//line views/vsprint/vsmember/Table.html:32
		qw422016.E().S(model.WebPath(paths...))
//line views/vsprint/vsmember/Table.html:32
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:32
		view.StreamUUID(qw422016, &model.SprintID)
//line views/vsprint/vsmember/Table.html:32
		if x := sprintsBySprintID.Get(model.SprintID); x != nil {
//line views/vsprint/vsmember/Table.html:32
			qw422016.N().S(` (`)
//line views/vsprint/vsmember/Table.html:32
			qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Table.html:32
			qw422016.N().S(`)`)
//line views/vsprint/vsmember/Table.html:32
		}
//line views/vsprint/vsmember/Table.html:32
		qw422016.N().S(`</a>
            <a title="Sprint" href="`)
//line views/vsprint/vsmember/Table.html:33
		if x := sprintsBySprintID.Get(model.SprintID); x != nil {
//line views/vsprint/vsmember/Table.html:33
			qw422016.E().S(x.WebPath(paths...))
//line views/vsprint/vsmember/Table.html:33
		}
//line views/vsprint/vsmember/Table.html:33
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:33
		components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vsprint/vsmember/Table.html:33
		qw422016.N().S(`</a>
          </td>
          <td class="nowrap">
            <a href="`)
//line views/vsprint/vsmember/Table.html:36
		qw422016.E().S(model.WebPath(paths...))
//line views/vsprint/vsmember/Table.html:36
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:36
		view.StreamUUID(qw422016, &model.UserID)
//line views/vsprint/vsmember/Table.html:36
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vsprint/vsmember/Table.html:36
			qw422016.N().S(` (`)
//line views/vsprint/vsmember/Table.html:36
			qw422016.E().S(x.TitleString())
//line views/vsprint/vsmember/Table.html:36
			qw422016.N().S(`)`)
//line views/vsprint/vsmember/Table.html:36
		}
//line views/vsprint/vsmember/Table.html:36
		qw422016.N().S(`</a>
            <a title="User" href="`)
//line views/vsprint/vsmember/Table.html:37
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vsprint/vsmember/Table.html:37
			qw422016.E().S(x.WebPath(paths...))
//line views/vsprint/vsmember/Table.html:37
		}
//line views/vsprint/vsmember/Table.html:37
		qw422016.N().S(`">`)
//line views/vsprint/vsmember/Table.html:37
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vsprint/vsmember/Table.html:37
		qw422016.N().S(`</a>
          </td>
          <td><strong>`)
//line views/vsprint/vsmember/Table.html:39
		view.StreamString(qw422016, model.Name)
//line views/vsprint/vsmember/Table.html:39
		qw422016.N().S(`</strong></td>
          <td><a href="`)
//line views/vsprint/vsmember/Table.html:40
		qw422016.E().S(model.Picture)
//line views/vsprint/vsmember/Table.html:40
		qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vsprint/vsmember/Table.html:40
		qw422016.E().S(model.Picture)
//line views/vsprint/vsmember/Table.html:40
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vsprint/vsmember/Table.html:41
		qw422016.E().S(model.Role.String())
//line views/vsprint/vsmember/Table.html:41
		qw422016.N().S(`</td>
          <td>`)
//line views/vsprint/vsmember/Table.html:42
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vsprint/vsmember/Table.html:42
		qw422016.N().S(`</td>
          <td>`)
//line views/vsprint/vsmember/Table.html:43
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vsprint/vsmember/Table.html:43
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vsprint/vsmember/Table.html:45
	}
//line views/vsprint/vsmember/Table.html:45
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vsprint/vsmember/Table.html:49
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vsprint/vsmember/Table.html:49
		qw422016.N().S(`  <hr />
  `)
//line views/vsprint/vsmember/Table.html:51
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vsprint/vsmember/Table.html:51
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vsprint/vsmember/Table.html:53
	}
//line views/vsprint/vsmember/Table.html:54
}

//line views/vsprint/vsmember/Table.html:54
func WriteTable(qq422016 qtio422016.Writer, models smember.SprintMembers, sprintsBySprintID sprint.Sprints, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vsprint/vsmember/Table.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vsmember/Table.html:54
	StreamTable(qw422016, models, sprintsBySprintID, usersByUserID, params, as, ps, paths...)
//line views/vsprint/vsmember/Table.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vsmember/Table.html:54
}

//line views/vsprint/vsmember/Table.html:54
func Table(models smember.SprintMembers, sprintsBySprintID sprint.Sprints, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vsprint/vsmember/Table.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vsmember/Table.html:54
	WriteTable(qb422016, models, sprintsBySprintID, usersByUserID, params, as, ps, paths...)
//line views/vsprint/vsmember/Table.html:54
	qs422016 := string(qb422016.B)
//line views/vsprint/vsmember/Table.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vsmember/Table.html:54
	return qs422016
//line views/vsprint/vsmember/Table.html:54
}
