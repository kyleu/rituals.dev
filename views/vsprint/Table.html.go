// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/Table.html:2
package vsprint

//line views/vsprint/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
)

//line views/vsprint/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models sprint.Sprints, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Table.html:11
	qw422016.N().S(`
`)
//line views/vsprint/Table.html:12
	prms := params.Get("sprint", nil, ps.Logger).Sanitize("sprint")

//line views/vsprint/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vsprint/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "sprint", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vsprint/Table.html:16
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "sprint", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vsprint/Table.html:17
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "sprint", "title", "Title", "String text", prms, ps.URI, ps)
//line views/vsprint/Table.html:18
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "sprint", "icon", "Icon", "String text", prms, ps.URI, ps)
//line views/vsprint/Table.html:19
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "sprint", "status", "Status", "Available options: [new, active, complete]", prms, ps.URI, ps)
//line views/vsprint/Table.html:20
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "sprint", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000) (optional)", prms, ps.URI, ps)
//line views/vsprint/Table.html:21
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "sprint", "start_date", "Start Date", "Calendar date (optional)", prms, ps.URI, ps)
//line views/vsprint/Table.html:22
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "sprint", "end_date", "End Date", "Calendar date (optional)", prms, ps.URI, ps)
//line views/vsprint/Table.html:23
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:24
	components.StreamTableHeaderSimple(qw422016, "sprint", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vsprint/Table.html:24
	qw422016.N().S(`
        `)
//line views/vsprint/Table.html:25
	components.StreamTableHeaderSimple(qw422016, "sprint", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vsprint/Table.html:25
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vsprint/Table.html:29
	for _, model := range models {
//line views/vsprint/Table.html:29
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/sprint/`)
//line views/vsprint/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vsprint/Table.html:31
		qw422016.N().S(`">`)
//line views/vsprint/Table.html:31
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vsprint/Table.html:31
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vsprint/Table.html:32
		qw422016.E().S(model.Slug)
//line views/vsprint/Table.html:32
		qw422016.N().S(`</td>
        <td><strong>`)
//line views/vsprint/Table.html:33
		qw422016.E().S(model.Title)
//line views/vsprint/Table.html:33
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vsprint/Table.html:34
		qw422016.E().S(model.Icon)
//line views/vsprint/Table.html:34
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/Table.html:35
		qw422016.E().V(model.Status)
//line views/vsprint/Table.html:35
		qw422016.N().S(`</td>
        <td class="nowrap">
          `)
//line views/vsprint/Table.html:37
		components.StreamDisplayUUID(qw422016, model.TeamID)
//line views/vsprint/Table.html:37
		if model.TeamID != nil {
//line views/vsprint/Table.html:37
			if x := teamsByTeamID.Get(*model.TeamID); x != nil {
//line views/vsprint/Table.html:37
				qw422016.N().S(` (`)
//line views/vsprint/Table.html:37
				qw422016.E().S(x.TitleString())
//line views/vsprint/Table.html:37
				qw422016.N().S(`)`)
//line views/vsprint/Table.html:37
			}
//line views/vsprint/Table.html:37
		}
//line views/vsprint/Table.html:37
		qw422016.N().S(`
          `)
//line views/vsprint/Table.html:38
		if model.TeamID != nil {
//line views/vsprint/Table.html:38
			qw422016.N().S(`<a title="Team" href="`)
//line views/vsprint/Table.html:38
			qw422016.E().S(`/admin/db/team` + `/` + model.TeamID.String())
//line views/vsprint/Table.html:38
			qw422016.N().S(`">`)
//line views/vsprint/Table.html:38
			components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vsprint/Table.html:38
			qw422016.N().S(`</a>`)
//line views/vsprint/Table.html:38
		}
//line views/vsprint/Table.html:38
		qw422016.N().S(`
        </td>
        <td>`)
//line views/vsprint/Table.html:40
		components.StreamDisplayTimestampDay(qw422016, model.StartDate)
//line views/vsprint/Table.html:40
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/Table.html:41
		components.StreamDisplayTimestampDay(qw422016, model.EndDate)
//line views/vsprint/Table.html:41
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/Table.html:42
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vsprint/Table.html:42
		qw422016.N().S(`</td>
        <td>`)
//line views/vsprint/Table.html:43
		components.StreamDisplayTimestamp(qw422016, model.Updated)
//line views/vsprint/Table.html:43
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vsprint/Table.html:45
	}
//line views/vsprint/Table.html:46
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vsprint/Table.html:46
		qw422016.N().S(`      <tr>
        <td colspan="10">`)
//line views/vsprint/Table.html:48
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vsprint/Table.html:48
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vsprint/Table.html:50
	}
//line views/vsprint/Table.html:50
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vsprint/Table.html:53
}

//line views/vsprint/Table.html:53
func WriteTable(qq422016 qtio422016.Writer, models sprint.Sprints, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Table.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/Table.html:53
	StreamTable(qw422016, models, teamsByTeamID, params, as, ps)
//line views/vsprint/Table.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/Table.html:53
}

//line views/vsprint/Table.html:53
func Table(models sprint.Sprints, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vsprint/Table.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/Table.html:53
	WriteTable(qb422016, models, teamsByTeamID, params, as, ps)
//line views/vsprint/Table.html:53
	qs422016 := string(qb422016.B)
//line views/vsprint/Table.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/Table.html:53
	return qs422016
//line views/vsprint/Table.html:53
}
