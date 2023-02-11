// Code generated by qtc from "SprintWorkspace.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwsprint/SprintWorkspace.html:1
package vwsprint

//line views/vworkspace/vwsprint/SprintWorkspace.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwestimate"
	"github.com/kyleu/rituals/views/vworkspace/vwretro"
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwsprint/SprintWorkspace.html:21
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwsprint/SprintWorkspace.html:21
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwsprint/SprintWorkspace.html:21
type SprintWorkspace struct {
	layout.Basic
	FullSprint *workspace.FullSprint
	Teams      team.Teams
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:27
func streamsummary(qw422016 *qt422016.Writer, s *sprint.Sprint) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:27
	qw422016.N().S(`<span>`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:29
	if s.StartDate != nil {
//line views/vworkspace/vwsprint/SprintWorkspace.html:29
		qw422016.N().S(`starts`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:30
		qw422016.N().S(` `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:31
		components.StreamDisplayTimestampDay(qw422016, s.StartDate)
//line views/vworkspace/vwsprint/SprintWorkspace.html:32
		if s.EndDate != nil {
//line views/vworkspace/vwsprint/SprintWorkspace.html:32
			qw422016.N().S(`,`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:33
			qw422016.N().S(` `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:34
		}
//line views/vworkspace/vwsprint/SprintWorkspace.html:35
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:36
	if s.EndDate != nil {
//line views/vworkspace/vwsprint/SprintWorkspace.html:36
		qw422016.N().S(`ends`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:37
		qw422016.N().S(` `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:38
		components.StreamDisplayTimestampDay(qw422016, s.EndDate)
//line views/vworkspace/vwsprint/SprintWorkspace.html:39
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:39
	qw422016.N().S(`</span>`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:41
func writesummary(qq422016 qtio422016.Writer, s *sprint.Sprint) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	streamsummary(qw422016, s)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:41
func summary(s *sprint.Sprint) string {
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	writesummary(qb422016, s)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
	return qs422016
//line views/vworkspace/vwsprint/SprintWorkspace.html:41
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:43
func (p *SprintWorkspace) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:43
	qw422016.N().S(`
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:45
	w := p.FullSprint
	s := w.Sprint
	self, others, _ := w.UtilMembers.Split(ps.Profile.ID)

//line views/vworkspace/vwsprint/SprintWorkspace.html:48
	qw422016.N().S(`  <div class="flex-wrap">
    <div id="panel-summary">
      <div class="card">
        <div class="right">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:52
	vwutil.StreamComments(qw422016, enum.ModelServiceSprint, s.ID, s.TitleString(), w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:52
	qw422016.N().S(`</div>
        <h3><a href="#modal-sprint-config" id="modal-sprint-config-link">
          <span id="model-icon">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:54
	components.StreamSVGRefIcon(qw422016, s.IconSafe(), ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:54
	qw422016.N().S(`</span>
          <span id="model-title">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:55
	qw422016.E().S(s.TitleString())
//line views/vworkspace/vwsprint/SprintWorkspace.html:55
	qw422016.N().S(`</span>
        </a></h3>
        <div class="clear"></div>
        <div class="right" id="model-summary">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
	if s.StartDate != nil {
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
		qw422016.N().S(`starts `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
		components.StreamDisplayTimestampDay(qw422016, s.StartDate)
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
		if s.EndDate != nil {
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
			qw422016.N().S(`, `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
		}
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
	if s.EndDate != nil {
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
		qw422016.N().S(`ends `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
		components.StreamDisplayTimestampDay(qw422016, s.EndDate)
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:58
	qw422016.N().S(`</div>
        `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:59
	vwutil.StreamBanner(qw422016, w.Team, nil, util.KeySprint)
//line views/vworkspace/vwsprint/SprintWorkspace.html:59
	qw422016.N().S(`
        `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:60
	StreamSprintWorkspaceModal(qw422016, s, p.Teams, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:60
	qw422016.N().S(`
      </div>
    </div>
    <div id="panel-detail">
      `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:64
	vwestimate.StreamEstimateWorkspaceList(qw422016, w.Estimates, s.TeamID, &s.ID, true, w.Comments, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:64
	qw422016.N().S(`
      `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:65
	vwstandup.StreamStandupWorkspaceList(qw422016, w.Standups, s.TeamID, &s.ID, true, w.Comments, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:65
	qw422016.N().S(`
      `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:66
	vwretro.StreamRetroWorkspaceList(qw422016, w.Retros, s.TeamID, &s.ID, true, w.Comments, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:66
	qw422016.N().S(`
    </div>
    <div id="panel-self">
      <div class="card">
        <span id="self-id" style="display: none;">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:70
	qw422016.E().S(self.UserID.String())
//line views/vworkspace/vwsprint/SprintWorkspace.html:70
	qw422016.N().S(`</span>
        <h3><a href="#modal-self">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:71
	components.StreamSVGRefIcon(qw422016, `profile`, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:71
	qw422016.N().S(`<span id="self-name">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:71
	qw422016.E().S(self.Name)
//line views/vworkspace/vwsprint/SprintWorkspace.html:71
	qw422016.N().S(`</span></a></h3>
        <em id="self-role">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:72
	qw422016.E().S(string(self.Role))
//line views/vworkspace/vwsprint/SprintWorkspace.html:72
	qw422016.N().S(`</em>
      </div>
      `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:74
	vwutil.StreamSelfModal(qw422016, self.Name, self.Picture, self.Role, s.PublicWebPath())
//line views/vworkspace/vwsprint/SprintWorkspace.html:74
	qw422016.N().S(`
    </div>
    <div id="panel-members">
      <div class="card">
        <h3><a href="#modal-invite">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:78
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:78
	qw422016.N().S(`Members</a></h3>
        <table class="mt expanded">
          <tbody>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:81
	for _, m := range others {
//line views/vworkspace/vwsprint/SprintWorkspace.html:81
		qw422016.N().S(`            `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:82
		vwutil.StreamMemberRow(qw422016, m, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:82
		qw422016.N().S(`
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:83
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:83
	qw422016.N().S(`          </tbody>
        </table>
        <div id="member-modals">
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:87
	for _, m := range others {
//line views/vworkspace/vwsprint/SprintWorkspace.html:87
		qw422016.N().S(`          `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:88
		vwutil.StreamMemberModal(qw422016, m, s.PublicWebPath())
//line views/vworkspace/vwsprint/SprintWorkspace.html:88
		qw422016.N().S(`
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:89
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:89
	qw422016.N().S(`        </div>
      </div>
    </div>
    `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:93
	vwutil.StreamInviteModal(qw422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:93
	qw422016.N().S(`
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      initWorkspace("`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:97
	qw422016.E().S(util.KeySprint)
//line views/vworkspace/vwsprint/SprintWorkspace.html:97
	qw422016.N().S(`", "`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:97
	qw422016.E().S(s.ID.String())
//line views/vworkspace/vwsprint/SprintWorkspace.html:97
	qw422016.N().S(`");
    });
  </script>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:100
func (p *SprintWorkspace) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:100
func (p *SprintWorkspace) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
	return qs422016
//line views/vworkspace/vwsprint/SprintWorkspace.html:100
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:102
func StreamSprintWorkspaceModal(qw422016 *qt422016.Writer, s *sprint.Sprint, teams team.Teams, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:102
	qw422016.N().S(`
  <div id="modal-sprint-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Sprint</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:111
	qw422016.E().S(s.PublicWebPath())
//line views/vworkspace/vwsprint/SprintWorkspace.html:111
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:112
	qw422016.E().S(string(action.ActUpdate))
//line views/vworkspace/vwsprint/SprintWorkspace.html:112
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:113
	components.StreamFormVerticalInput(qw422016, "title", "", "Title", s.TitleString(), 5, "The name of your sprint")
//line views/vworkspace/vwsprint/SprintWorkspace.html:113
	qw422016.N().S(`
          `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:114
	components.StreamFormVerticalIconPicker(qw422016, "icon", "Icon", s.IconSafe(), ps, 5)
//line views/vworkspace/vwsprint/SprintWorkspace.html:114
	qw422016.N().S(`
          `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:115
	components.StreamFormVerticalInputTimestampDay(qw422016, "startDate", "", "Start Date", s.StartDate, 5, "The day your sprint begins")
//line views/vworkspace/vwsprint/SprintWorkspace.html:115
	qw422016.N().S(`
          `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:116
	components.StreamFormVerticalInputTimestampDay(qw422016, "endDate", "", "End Date", s.EndDate, 5, "The day your sprint ends")
//line views/vworkspace/vwsprint/SprintWorkspace.html:116
	qw422016.N().S(`
          `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:117
	components.StreamFormVerticalSelect(qw422016, util.KeyTeam, "", "Team", util.UUIDString(s.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5)
//line views/vworkspace/vwsprint/SprintWorkspace.html:117
	qw422016.N().S(`
          <em class="title">Permissions</em>
          <div>
            <label><input type="checkbox" name="perm-team" value="true"> Must be a member of this sprint's team</label>
          </div>
          <hr />
          <div class="right"><button type="submit">Save</button></div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:128
func WriteSprintWorkspaceModal(qq422016 qtio422016.Writer, s *sprint.Sprint, teams team.Teams, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	StreamSprintWorkspaceModal(qw422016, s, teams, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:128
func SprintWorkspaceModal(s *sprint.Sprint, teams team.Teams, ps *cutil.PageState) string {
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	WriteSprintWorkspaceModal(qb422016, s, teams, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
	return qs422016
//line views/vworkspace/vwsprint/SprintWorkspace.html:128
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:130
func StreamSprintWorkspaceList(qw422016 *qt422016.Writer, sprints sprint.Sprints, teamID *uuid.UUID, showComments bool, comments comment.Comments, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:130
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:132
	vwutil.StreamEditWorkspaceForm(qw422016, util.KeySprint, teamID, nil, "New Sprint")
//line views/vworkspace/vwsprint/SprintWorkspace.html:132
	qw422016.N().S(`</div>
    <h3>`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:133
	components.StreamSVGRefIcon(qw422016, util.KeySprint, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:133
	qw422016.N().S(`Sprints</h3>
    <table id="sprint-list" class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:136
	if len(sprints) == 0 {
//line views/vworkspace/vwsprint/SprintWorkspace.html:136
		qw422016.N().S(`          <tr class="empty"><td><em>no sprints</em></td></tr>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:138
	} else {
//line views/vworkspace/vwsprint/SprintWorkspace.html:139
		for _, x := range sprints {
//line views/vworkspace/vwsprint/SprintWorkspace.html:139
			qw422016.N().S(`          <tr id="sprint-list-`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:140
			qw422016.E().S(x.ID.String())
//line views/vworkspace/vwsprint/SprintWorkspace.html:140
			qw422016.N().S(`">
            <td>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:142
			if showComments {
//line views/vworkspace/vwsprint/SprintWorkspace.html:142
				qw422016.N().S(`              <div class="right">
                `)
//line views/vworkspace/vwsprint/SprintWorkspace.html:144
				vwutil.StreamComments(qw422016, enum.ModelServiceSprint, x.ID, x.TitleString(), comments, nil, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:144
				qw422016.N().S(`
              </div>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:146
			}
//line views/vworkspace/vwsprint/SprintWorkspace.html:146
			qw422016.N().S(`              <a href="`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:147
			qw422016.E().S(x.PublicWebPath())
//line views/vworkspace/vwsprint/SprintWorkspace.html:147
			qw422016.N().S(`"><div>
                <span>`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:148
			components.StreamSVGRef(qw422016, x.IconSafe(), 16, 16, "icon", ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:148
			qw422016.N().S(`</span><span>`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:148
			qw422016.E().S(x.TitleString())
//line views/vworkspace/vwsprint/SprintWorkspace.html:148
			qw422016.N().S(`</span>
              </div></a>
            </td>
          </tr>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:152
		}
//line views/vworkspace/vwsprint/SprintWorkspace.html:153
	}
//line views/vworkspace/vwsprint/SprintWorkspace.html:153
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:157
func WriteSprintWorkspaceList(qq422016 qtio422016.Writer, sprints sprint.Sprints, teamID *uuid.UUID, showComments bool, comments comment.Comments, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	StreamSprintWorkspaceList(qw422016, sprints, teamID, showComments, comments, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
}

//line views/vworkspace/vwsprint/SprintWorkspace.html:157
func SprintWorkspaceList(sprints sprint.Sprints, teamID *uuid.UUID, showComments bool, comments comment.Comments, ps *cutil.PageState) string {
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	WriteSprintWorkspaceList(qb422016, sprints, teamID, showComments, comments, ps)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
	return qs422016
//line views/vworkspace/vwsprint/SprintWorkspace.html:157
}
