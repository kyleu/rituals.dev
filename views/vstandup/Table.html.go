// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vstandup/Table.html:1
package vstandup

//line views/vstandup/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vstandup/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/Table.html:13
func StreamTable(qw422016 *qt422016.Writer, models standup.Standups, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vstandup/Table.html:13
	qw422016.N().S(`
`)
//line views/vstandup/Table.html:14
	prms := params.Sanitized("standup", ps.Logger)

//line views/vstandup/Table.html:14
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vstandup/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "standup", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vstandup/Table.html:19
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "standup", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vstandup/Table.html:20
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "standup", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vstandup/Table.html:21
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "standup", "icon", "Icon", "String text", prms, ps.URI, ps)
//line views/vstandup/Table.html:22
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "standup", "status", "Status", enum.AllSessionStatuses.Help(), prms, ps.URI, ps)
//line views/vstandup/Table.html:23
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "standup", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vstandup/Table.html:24
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "standup", "sprint_id", "Sprint ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vstandup/Table.html:25
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:26
	components.StreamTableHeaderSimple(qw422016, "standup", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vstandup/Table.html:26
	qw422016.N().S(`
          `)
//line views/vstandup/Table.html:27
	components.StreamTableHeaderSimple(qw422016, "standup", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vstandup/Table.html:27
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vstandup/Table.html:31
	for _, model := range models {
//line views/vstandup/Table.html:31
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vstandup/Table.html:33
		qw422016.E().S(model.WebPath(paths...))
//line views/vstandup/Table.html:33
		qw422016.N().S(`">`)
//line views/vstandup/Table.html:33
		view.StreamUUID(qw422016, &model.ID)
//line views/vstandup/Table.html:33
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vstandup/Table.html:34
		view.StreamString(qw422016, model.Slug)
//line views/vstandup/Table.html:34
		qw422016.N().S(`</td>
          <td><strong>`)
//line views/vstandup/Table.html:35
		view.StreamString(qw422016, model.Title)
//line views/vstandup/Table.html:35
		qw422016.N().S(`</strong></td>
          <td>`)
//line views/vstandup/Table.html:36
		view.StreamString(qw422016, model.Icon)
//line views/vstandup/Table.html:36
		qw422016.N().S(`</td>
          <td>`)
//line views/vstandup/Table.html:37
		qw422016.E().S(model.Status.String())
//line views/vstandup/Table.html:37
		qw422016.N().S(`</td>
          <td class="nowrap">
            `)
//line views/vstandup/Table.html:39
		if x := teamsByTeamID.Get(*model.TeamID); x != nil {
//line views/vstandup/Table.html:39
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:40
			qw422016.E().S(x.TitleString())
//line views/vstandup/Table.html:40
			qw422016.N().S(` `)
//line views/vstandup/Table.html:40
			if model.TeamID != nil {
//line views/vstandup/Table.html:40
				qw422016.N().S(`<a title="Team" href="`)
//line views/vstandup/Table.html:40
				qw422016.E().S(x.WebPath(paths...))
//line views/vstandup/Table.html:40
				qw422016.N().S(`">`)
//line views/vstandup/Table.html:40
				components.StreamSVGLink(qw422016, `team`, ps)
//line views/vstandup/Table.html:40
				qw422016.N().S(`</a>`)
//line views/vstandup/Table.html:40
			}
//line views/vstandup/Table.html:40
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:41
		} else {
//line views/vstandup/Table.html:41
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:42
			view.StreamUUID(qw422016, model.TeamID)
//line views/vstandup/Table.html:42
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:43
		}
//line views/vstandup/Table.html:43
		qw422016.N().S(`
          </td>
          <td class="nowrap">
            `)
//line views/vstandup/Table.html:46
		if x := sprintsBySprintID.Get(*model.SprintID); x != nil {
//line views/vstandup/Table.html:46
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:47
			qw422016.E().S(x.TitleString())
//line views/vstandup/Table.html:47
			qw422016.N().S(` `)
//line views/vstandup/Table.html:47
			if model.SprintID != nil {
//line views/vstandup/Table.html:47
				qw422016.N().S(`<a title="Sprint" href="`)
//line views/vstandup/Table.html:47
				qw422016.E().S(x.WebPath(paths...))
//line views/vstandup/Table.html:47
				qw422016.N().S(`">`)
//line views/vstandup/Table.html:47
				components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vstandup/Table.html:47
				qw422016.N().S(`</a>`)
//line views/vstandup/Table.html:47
			}
//line views/vstandup/Table.html:47
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:48
		} else {
//line views/vstandup/Table.html:48
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:49
			view.StreamUUID(qw422016, model.SprintID)
//line views/vstandup/Table.html:49
			qw422016.N().S(`
            `)
//line views/vstandup/Table.html:50
		}
//line views/vstandup/Table.html:50
		qw422016.N().S(`
          </td>
          <td>`)
//line views/vstandup/Table.html:52
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vstandup/Table.html:52
		qw422016.N().S(`</td>
          <td>`)
//line views/vstandup/Table.html:53
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vstandup/Table.html:53
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vstandup/Table.html:55
	}
//line views/vstandup/Table.html:55
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vstandup/Table.html:59
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vstandup/Table.html:59
		qw422016.N().S(`  <hr />
  `)
//line views/vstandup/Table.html:61
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vstandup/Table.html:61
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vstandup/Table.html:63
	}
//line views/vstandup/Table.html:64
}

//line views/vstandup/Table.html:64
func WriteTable(qq422016 qtio422016.Writer, models standup.Standups, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vstandup/Table.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Table.html:64
	StreamTable(qw422016, models, teamsByTeamID, sprintsBySprintID, params, as, ps, paths...)
//line views/vstandup/Table.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Table.html:64
}

//line views/vstandup/Table.html:64
func Table(models standup.Standups, teamsByTeamID team.Teams, sprintsBySprintID sprint.Sprints, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vstandup/Table.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Table.html:64
	WriteTable(qb422016, models, teamsByTeamID, sprintsBySprintID, params, as, ps, paths...)
//line views/vstandup/Table.html:64
	qs422016 := string(qb422016.B)
//line views/vstandup/Table.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Table.html:64
	return qs422016
//line views/vstandup/Table.html:64
}
