// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vteam/vtpermission/Table.html:1
package vtpermission

//line views/vteam/vtpermission/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vteam/vtpermission/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/vtpermission/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/vtpermission/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models tpermission.TeamPermissions, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vteam/vtpermission/Table.html:11
	qw422016.N().S(`
`)
//line views/vteam/vtpermission/Table.html:12
	prms := params.Sanitized("tpermission", ps.Logger)

//line views/vteam/vtpermission/Table.html:12
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vteam/vtpermission/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "tpermission", "team_id", "Team ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vteam/vtpermission/Table.html:17
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "tpermission", "key", "Key", "String text", prms, ps.URI, ps)
//line views/vteam/vtpermission/Table.html:18
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "tpermission", "value", "Value", "String text", prms, ps.URI, ps)
//line views/vteam/vtpermission/Table.html:19
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "tpermission", "access", "Access", "String text", prms, ps.URI, ps)
//line views/vteam/vtpermission/Table.html:20
	qw422016.N().S(`
          `)
//line views/vteam/vtpermission/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "tpermission", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vteam/vtpermission/Table.html:21
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vteam/vtpermission/Table.html:25
	for _, model := range models {
//line views/vteam/vtpermission/Table.html:25
		qw422016.N().S(`        <tr>
          <td class="nowrap">
            `)
//line views/vteam/vtpermission/Table.html:28
		if x := teamsByTeamID.Get(model.TeamID); x != nil {
//line views/vteam/vtpermission/Table.html:28
			qw422016.N().S(`
            `)
//line views/vteam/vtpermission/Table.html:29
			qw422016.E().S(x.TitleString())
//line views/vteam/vtpermission/Table.html:29
			qw422016.N().S(` <a title="Team" href="`)
//line views/vteam/vtpermission/Table.html:29
			qw422016.E().S(x.WebPath(paths...))
//line views/vteam/vtpermission/Table.html:29
			qw422016.N().S(`">`)
//line views/vteam/vtpermission/Table.html:29
			components.StreamSVGLink(qw422016, `team`, ps)
//line views/vteam/vtpermission/Table.html:29
			qw422016.N().S(`</a>
            `)
//line views/vteam/vtpermission/Table.html:30
		} else {
//line views/vteam/vtpermission/Table.html:30
			qw422016.N().S(`
            `)
//line views/vteam/vtpermission/Table.html:31
			view.StreamUUID(qw422016, &model.TeamID)
//line views/vteam/vtpermission/Table.html:31
			qw422016.N().S(`
            `)
//line views/vteam/vtpermission/Table.html:32
		}
//line views/vteam/vtpermission/Table.html:32
		qw422016.N().S(`
          </td>
          <td><a href="`)
//line views/vteam/vtpermission/Table.html:34
		qw422016.E().S(model.WebPath(paths...))
//line views/vteam/vtpermission/Table.html:34
		qw422016.N().S(`">`)
//line views/vteam/vtpermission/Table.html:34
		view.StreamString(qw422016, model.Key)
//line views/vteam/vtpermission/Table.html:34
		qw422016.N().S(`</a></td>
          <td><a href="`)
//line views/vteam/vtpermission/Table.html:35
		qw422016.E().S(model.WebPath(paths...))
//line views/vteam/vtpermission/Table.html:35
		qw422016.N().S(`">`)
//line views/vteam/vtpermission/Table.html:35
		view.StreamString(qw422016, model.Value)
//line views/vteam/vtpermission/Table.html:35
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vteam/vtpermission/Table.html:36
		view.StreamString(qw422016, model.Access)
//line views/vteam/vtpermission/Table.html:36
		qw422016.N().S(`</td>
          <td>`)
//line views/vteam/vtpermission/Table.html:37
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vteam/vtpermission/Table.html:37
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vteam/vtpermission/Table.html:39
	}
//line views/vteam/vtpermission/Table.html:39
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vteam/vtpermission/Table.html:43
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vteam/vtpermission/Table.html:43
		qw422016.N().S(`  <hr />
  `)
//line views/vteam/vtpermission/Table.html:45
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vteam/vtpermission/Table.html:45
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vteam/vtpermission/Table.html:47
	}
//line views/vteam/vtpermission/Table.html:48
}

//line views/vteam/vtpermission/Table.html:48
func WriteTable(qq422016 qtio422016.Writer, models tpermission.TeamPermissions, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vteam/vtpermission/Table.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/vtpermission/Table.html:48
	StreamTable(qw422016, models, teamsByTeamID, params, as, ps, paths...)
//line views/vteam/vtpermission/Table.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/vtpermission/Table.html:48
}

//line views/vteam/vtpermission/Table.html:48
func Table(models tpermission.TeamPermissions, teamsByTeamID team.Teams, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vteam/vtpermission/Table.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/vtpermission/Table.html:48
	WriteTable(qb422016, models, teamsByTeamID, params, as, ps, paths...)
//line views/vteam/vtpermission/Table.html:48
	qs422016 := string(qb422016.B)
//line views/vteam/vtpermission/Table.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/vtpermission/Table.html:48
	return qs422016
//line views/vteam/vtpermission/Table.html:48
}
