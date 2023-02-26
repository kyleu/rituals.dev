// Code generated by qtc from "TeamWorkspace.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwteam/TeamWorkspace.html:1
package vwteam

//line views/vworkspace/vwteam/TeamWorkspace.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwestimate"
	"github.com/kyleu/rituals/views/vworkspace/vwretro"
	"github.com/kyleu/rituals/views/vworkspace/vwsprint"
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwteam/TeamWorkspace.html:18
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwteam/TeamWorkspace.html:18
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwteam/TeamWorkspace.html:18
type TeamWorkspace struct {
	layout.Basic
	FullTeam *workspace.FullTeam
}

//line views/vworkspace/vwteam/TeamWorkspace.html:23
func (p *TeamWorkspace) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:23
	qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:25
	w := p.FullTeam
	t := w.Team

//line views/vworkspace/vwteam/TeamWorkspace.html:27
	qw422016.N().S(`  <div class="flex-wrap">
    <div id="panel-summary">
      <div class="card">
        <div class="right">
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:32
	vwutil.StreamPermissionsLink(qw422016, enum.ModelServiceTeam, t.ID, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:32
	qw422016.N().S(`
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:33
	vwutil.StreamComments(qw422016, enum.ModelServiceTeam, t.ID, t.TitleString(), w.Comments, w.UtilMembers, "", ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:33
	qw422016.N().S(`
        </div>
        <a href="#modal-team-config" id="modal-team-config-link"><h3>
          <span id="model-icon">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:36
	components.StreamSVGRefIcon(qw422016, t.IconSafe(), ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:36
	qw422016.N().S(`</span>
          <span id="model-title">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:37
	qw422016.E().S(t.TitleString())
//line views/vworkspace/vwteam/TeamWorkspace.html:37
	qw422016.N().S(`</span>
        </h3></a>
        `)
//line views/vworkspace/vwteam/TeamWorkspace.html:39
	vwutil.StreamBanner(qw422016, nil, nil, util.KeyTeam)
//line views/vworkspace/vwteam/TeamWorkspace.html:39
	qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:40
	if w.Admin() {
//line views/vworkspace/vwteam/TeamWorkspace.html:40
		qw422016.N().S(`        `)
//line views/vworkspace/vwteam/TeamWorkspace.html:41
		StreamTeamWorkspaceModalEdit(qw422016, t, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:41
		qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:42
	} else {
//line views/vworkspace/vwteam/TeamWorkspace.html:42
		qw422016.N().S(`        `)
//line views/vworkspace/vwteam/TeamWorkspace.html:43
		StreamTeamWorkspaceModalView(qw422016, t, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:43
		qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:44
	}
//line views/vworkspace/vwteam/TeamWorkspace.html:44
	qw422016.N().S(`      </div>
    </div>
    <div id="panel-detail">
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:48
	vwsprint.StreamSprintListTable(qw422016, w.Sprints, &t.ID, true, w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:48
	qw422016.N().S(`
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:49
	vwestimate.StreamEstimateListTable(qw422016, w.Estimates, &t.ID, nil, true, w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:49
	qw422016.N().S(`
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:50
	vwstandup.StreamStandupListTable(qw422016, w.Standups, &t.ID, nil, true, w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:50
	qw422016.N().S(`
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:51
	vwretro.StreamRetroListTable(qw422016, w.Retros, &t.ID, nil, true, w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:51
	qw422016.N().S(`
    </div>
    `)
//line views/vworkspace/vwteam/TeamWorkspace.html:53
	vwutil.StreamMemberPanels(qw422016, w.UtilMembers, w.Admin(), t.PublicWebPath(), ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:53
	qw422016.N().S(`
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      initWorkspace("`)
//line views/vworkspace/vwteam/TeamWorkspace.html:57
	qw422016.E().S(util.KeyTeam)
//line views/vworkspace/vwteam/TeamWorkspace.html:57
	qw422016.N().S(`", "`)
//line views/vworkspace/vwteam/TeamWorkspace.html:57
	qw422016.E().S(t.ID.String())
//line views/vworkspace/vwteam/TeamWorkspace.html:57
	qw422016.N().S(`");
    });
  </script>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
}

//line views/vworkspace/vwteam/TeamWorkspace.html:60
func (p *TeamWorkspace) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
}

//line views/vworkspace/vwteam/TeamWorkspace.html:60
func (p *TeamWorkspace) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	return qs422016
//line views/vworkspace/vwteam/TeamWorkspace.html:60
}

//line views/vworkspace/vwteam/TeamWorkspace.html:62
func StreamTeamWorkspaceModalEdit(qw422016 *qt422016.Writer, t *team.Team, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:62
	qw422016.N().S(`
  <div id="modal-team-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwteam/TeamWorkspace.html:68
	components.StreamSVGRef(qw422016, util.KeyTeam, 24, 24, "icon", ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:68
	qw422016.N().S(` Team</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwteam/TeamWorkspace.html:71
	qw422016.E().S(t.PublicWebPath())
//line views/vworkspace/vwteam/TeamWorkspace.html:71
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwteam/TeamWorkspace.html:72
	qw422016.E().S(string(action.ActUpdate))
//line views/vworkspace/vwteam/TeamWorkspace.html:72
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:73
	components.StreamFormVerticalInput(qw422016, "title", "", "Title", t.TitleString(), 5, "The name of your team")
//line views/vworkspace/vwteam/TeamWorkspace.html:73
	qw422016.N().S(`
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:74
	components.StreamFormVerticalIconPicker(qw422016, "icon", "Icon", t.IconSafe(), ps, 5)
//line views/vworkspace/vwteam/TeamWorkspace.html:74
	qw422016.N().S(`
          <hr />
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:76
	vwutil.StreamPermissionsForm(qw422016, util.KeyTeam, perms, false, nil, false, nil, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:76
	qw422016.N().S(`
          <hr />
          <div class="right"><button type="submit">Save</button></div>
          <a href="`)
//line views/vworkspace/vwteam/TeamWorkspace.html:79
	qw422016.E().S(t.PublicWebPath())
//line views/vworkspace/vwteam/TeamWorkspace.html:79
	qw422016.N().S(`/delete" onclick="return confirm('are you sure you wish to delete this team?')"><button type="button">Delete</button></a>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
}

//line views/vworkspace/vwteam/TeamWorkspace.html:84
func WriteTeamWorkspaceModalEdit(qq422016 qtio422016.Writer, t *team.Team, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	StreamTeamWorkspaceModalEdit(qw422016, t, perms, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
}

//line views/vworkspace/vwteam/TeamWorkspace.html:84
func TeamWorkspaceModalEdit(t *team.Team, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	WriteTeamWorkspaceModalEdit(qb422016, t, perms, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:84
	return qs422016
//line views/vworkspace/vwteam/TeamWorkspace.html:84
}

//line views/vworkspace/vwteam/TeamWorkspace.html:86
func StreamTeamWorkspaceModalView(qw422016 *qt422016.Writer, t *team.Team, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:86
	qw422016.N().S(`
  <div id="modal-team-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="view-icon">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:93
	components.StreamSVGRef(qw422016, t.IconSafe(), 24, 24, "icon", ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:93
	qw422016.N().S(`</span>
          <span class="view-title">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:94
	qw422016.E().S(t.TitleString())
//line views/vworkspace/vwteam/TeamWorkspace.html:94
	qw422016.N().S(`</span>
        </h2>
      </div>
      <div class="modal-body">
        <div style="display: none;">
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:99
	components.StreamIconPicker(qw422016, "icon", t.IconSafe(), ps, 5)
//line views/vworkspace/vwteam/TeamWorkspace.html:99
	qw422016.N().S(`
        </div>
        <table>
          <tbody>
            <tr>
              <th class="shrink">Permissions</th>
              <td class="config-panel-perms">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:105
	vwutil.StreamPermissionsList(qw422016, util.KeyRetro, perms, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:105
	qw422016.N().S(`</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
}

//line views/vworkspace/vwteam/TeamWorkspace.html:112
func WriteTeamWorkspaceModalView(qq422016 qtio422016.Writer, t *team.Team, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	StreamTeamWorkspaceModalView(qw422016, t, perms, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
}

//line views/vworkspace/vwteam/TeamWorkspace.html:112
func TeamWorkspaceModalView(t *team.Team, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	WriteTeamWorkspaceModalView(qb422016, t, perms, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:112
	return qs422016
//line views/vworkspace/vwteam/TeamWorkspace.html:112
}
