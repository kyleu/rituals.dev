// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vuser/Detail.html:2
package vuser

//line views/vuser/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vaction"
	"github.com/kyleu/rituals/views/vcomment"
	"github.com/kyleu/rituals/views/vemail"
	"github.com/kyleu/rituals/views/vestimate/vemember"
	"github.com/kyleu/rituals/views/vestimate/vstory"
	"github.com/kyleu/rituals/views/vestimate/vstory/vvote"
	"github.com/kyleu/rituals/views/vretro/vfeedback"
	"github.com/kyleu/rituals/views/vretro/vrmember"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
	"github.com/kyleu/rituals/views/vstandup/vreport"
	"github.com/kyleu/rituals/views/vstandup/vumember"
	"github.com/kyleu/rituals/views/vteam/vtmember"
)

//line views/vuser/Detail.html:36
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vuser/Detail.html:36
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vuser/Detail.html:36
type Detail struct {
	layout.Basic
	Model                      *user.User
	Params                     filter.ParamSet
	RelActionsByUserID         action.Actions
	RelCommentsByUserID        comment.Comments
	RelEmailsByUserID          email.Emails
	RelEstimateMembersByUserID emember.EstimateMembers
	RelFeedbacksByUserID       feedback.Feedbacks
	RelReportsByUserID         report.Reports
	RelRetroMembersByUserID    rmember.RetroMembers
	RelSprintMembersByUserID   smember.SprintMembers
	RelStandupMembersByUserID  umember.StandupMembers
	RelStoriesByUserID         story.Stories
	RelTeamMembersByUserID     tmember.TeamMembers
	RelVotesByUserID           vote.Votes
}

//line views/vuser/Detail.html:54
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/Detail.html:54
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-user"><button type="button">JSON</button></a>
      <a href="`)
//line views/vuser/Detail.html:58
	qw422016.E().S(p.Model.WebPath())
//line views/vuser/Detail.html:58
	qw422016.N().S(`/edit"><button>`)
//line views/vuser/Detail.html:58
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vuser/Detail.html:58
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vuser/Detail.html:60
	components.StreamSVGRefIcon(qw422016, `profile`, ps)
//line views/vuser/Detail.html:60
	qw422016.N().S(` `)
//line views/vuser/Detail.html:60
	qw422016.E().S(p.Model.TitleString())
//line views/vuser/Detail.html:60
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/user"><em>User</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vuser/Detail.html:66
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vuser/Detail.html:66
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Name</th>
          <td><strong>`)
//line views/vuser/Detail.html:70
	qw422016.E().S(p.Model.Name)
//line views/vuser/Detail.html:70
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="URL in string form">Picture</th>
          <td><a href="`)
//line views/vuser/Detail.html:74
	qw422016.E().S(p.Model.Picture)
//line views/vuser/Detail.html:74
	qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vuser/Detail.html:74
	qw422016.E().S(p.Model.Picture)
//line views/vuser/Detail.html:74
	qw422016.N().S(`</a></td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vuser/Detail.html:78
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vuser/Detail.html:78
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vuser/Detail.html:82
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vuser/Detail.html:82
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vuser/Detail.html:88
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-ActionsByUserID" type="checkbox" hidden />
        <label for="accordion-ActionsByUserID">
          `)
//line views/vuser/Detail.html:95
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:95
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:96
	components.StreamSVGRefIcon(qw422016, `action`, ps)
//line views/vuser/Detail.html:96
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:97
	qw422016.E().S(util.StringPlural(len(p.RelActionsByUserID), "Action"))
//line views/vuser/Detail.html:97
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:100
	if len(p.RelActionsByUserID) == 0 {
//line views/vuser/Detail.html:100
		qw422016.N().S(`          <em>no related Actions</em>
`)
//line views/vuser/Detail.html:102
	} else {
//line views/vuser/Detail.html:102
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:104
		vaction.StreamTable(qw422016, p.RelActionsByUserID, nil, p.Params, as, ps)
//line views/vuser/Detail.html:104
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:106
	}
//line views/vuser/Detail.html:106
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-CommentsByUserID" type="checkbox" hidden />
        <label for="accordion-CommentsByUserID">
          `)
//line views/vuser/Detail.html:112
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:112
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:113
	components.StreamSVGRefIcon(qw422016, `comments`, ps)
//line views/vuser/Detail.html:113
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:114
	qw422016.E().S(util.StringPlural(len(p.RelCommentsByUserID), "Comment"))
//line views/vuser/Detail.html:114
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:117
	if len(p.RelCommentsByUserID) == 0 {
//line views/vuser/Detail.html:117
		qw422016.N().S(`          <em>no related Comments</em>
`)
//line views/vuser/Detail.html:119
	} else {
//line views/vuser/Detail.html:119
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:121
		vcomment.StreamTable(qw422016, p.RelCommentsByUserID, nil, p.Params, as, ps)
//line views/vuser/Detail.html:121
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:123
	}
//line views/vuser/Detail.html:123
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EmailsByUserID" type="checkbox" hidden />
        <label for="accordion-EmailsByUserID">
          `)
//line views/vuser/Detail.html:129
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:129
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:130
	components.StreamSVGRefIcon(qw422016, `email`, ps)
//line views/vuser/Detail.html:130
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:131
	qw422016.E().S(util.StringPlural(len(p.RelEmailsByUserID), "Email"))
//line views/vuser/Detail.html:131
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:134
	if len(p.RelEmailsByUserID) == 0 {
//line views/vuser/Detail.html:134
		qw422016.N().S(`          <em>no related Emails</em>
`)
//line views/vuser/Detail.html:136
	} else {
//line views/vuser/Detail.html:136
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:138
		vemail.StreamTable(qw422016, p.RelEmailsByUserID, nil, p.Params, as, ps)
//line views/vuser/Detail.html:138
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:140
	}
//line views/vuser/Detail.html:140
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EstimateMembersByUserID" type="checkbox" hidden />
        <label for="accordion-EstimateMembersByUserID">
          `)
//line views/vuser/Detail.html:146
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:146
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:147
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vuser/Detail.html:147
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:148
	qw422016.E().S(util.StringPlural(len(p.RelEstimateMembersByUserID), "Member"))
//line views/vuser/Detail.html:148
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:151
	if len(p.RelEstimateMembersByUserID) == 0 {
//line views/vuser/Detail.html:151
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:153
	} else {
//line views/vuser/Detail.html:153
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:155
		vemember.StreamTable(qw422016, p.RelEstimateMembersByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:155
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:157
	}
//line views/vuser/Detail.html:157
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-FeedbacksByUserID" type="checkbox" hidden />
        <label for="accordion-FeedbacksByUserID">
          `)
//line views/vuser/Detail.html:163
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:163
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:164
	components.StreamSVGRefIcon(qw422016, `comment`, ps)
//line views/vuser/Detail.html:164
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:165
	qw422016.E().S(util.StringPlural(len(p.RelFeedbacksByUserID), "Feedback"))
//line views/vuser/Detail.html:165
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:168
	if len(p.RelFeedbacksByUserID) == 0 {
//line views/vuser/Detail.html:168
		qw422016.N().S(`          <em>no related Feedbacks</em>
`)
//line views/vuser/Detail.html:170
	} else {
//line views/vuser/Detail.html:170
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:172
		vfeedback.StreamTable(qw422016, p.RelFeedbacksByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:172
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:174
	}
//line views/vuser/Detail.html:174
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-ReportsByUserID" type="checkbox" hidden />
        <label for="accordion-ReportsByUserID">
          `)
//line views/vuser/Detail.html:180
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:180
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:181
	components.StreamSVGRefIcon(qw422016, `file-alt`, ps)
//line views/vuser/Detail.html:181
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:182
	qw422016.E().S(util.StringPlural(len(p.RelReportsByUserID), "Report"))
//line views/vuser/Detail.html:182
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:185
	if len(p.RelReportsByUserID) == 0 {
//line views/vuser/Detail.html:185
		qw422016.N().S(`          <em>no related Reports</em>
`)
//line views/vuser/Detail.html:187
	} else {
//line views/vuser/Detail.html:187
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:189
		vreport.StreamTable(qw422016, p.RelReportsByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:189
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:191
	}
//line views/vuser/Detail.html:191
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetroMembersByUserID" type="checkbox" hidden />
        <label for="accordion-RetroMembersByUserID">
          `)
//line views/vuser/Detail.html:197
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:197
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:198
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vuser/Detail.html:198
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:199
	qw422016.E().S(util.StringPlural(len(p.RelRetroMembersByUserID), "Member"))
//line views/vuser/Detail.html:199
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:202
	if len(p.RelRetroMembersByUserID) == 0 {
//line views/vuser/Detail.html:202
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:204
	} else {
//line views/vuser/Detail.html:204
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:206
		vrmember.StreamTable(qw422016, p.RelRetroMembersByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:206
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:208
	}
//line views/vuser/Detail.html:208
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-SprintMembersByUserID" type="checkbox" hidden />
        <label for="accordion-SprintMembersByUserID">
          `)
//line views/vuser/Detail.html:214
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:214
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:215
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vuser/Detail.html:215
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:216
	qw422016.E().S(util.StringPlural(len(p.RelSprintMembersByUserID), "Member"))
//line views/vuser/Detail.html:216
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:219
	if len(p.RelSprintMembersByUserID) == 0 {
//line views/vuser/Detail.html:219
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:221
	} else {
//line views/vuser/Detail.html:221
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:223
		vsmember.StreamTable(qw422016, p.RelSprintMembersByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:223
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:225
	}
//line views/vuser/Detail.html:225
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupMembersByUserID" type="checkbox" hidden />
        <label for="accordion-StandupMembersByUserID">
          `)
//line views/vuser/Detail.html:231
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:231
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:232
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vuser/Detail.html:232
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:233
	qw422016.E().S(util.StringPlural(len(p.RelStandupMembersByUserID), "Member"))
//line views/vuser/Detail.html:233
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:236
	if len(p.RelStandupMembersByUserID) == 0 {
//line views/vuser/Detail.html:236
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:238
	} else {
//line views/vuser/Detail.html:238
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:240
		vumember.StreamTable(qw422016, p.RelStandupMembersByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:240
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:242
	}
//line views/vuser/Detail.html:242
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StoriesByUserID" type="checkbox" hidden />
        <label for="accordion-StoriesByUserID">
          `)
//line views/vuser/Detail.html:248
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:248
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:249
	components.StreamSVGRefIcon(qw422016, `story`, ps)
//line views/vuser/Detail.html:249
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:250
	qw422016.E().S(util.StringPlural(len(p.RelStoriesByUserID), "Story"))
//line views/vuser/Detail.html:250
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:253
	if len(p.RelStoriesByUserID) == 0 {
//line views/vuser/Detail.html:253
		qw422016.N().S(`          <em>no related Stories</em>
`)
//line views/vuser/Detail.html:255
	} else {
//line views/vuser/Detail.html:255
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:257
		vstory.StreamTable(qw422016, p.RelStoriesByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:257
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:259
	}
//line views/vuser/Detail.html:259
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-TeamMembersByUserID" type="checkbox" hidden />
        <label for="accordion-TeamMembersByUserID">
          `)
//line views/vuser/Detail.html:265
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:265
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:266
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vuser/Detail.html:266
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:267
	qw422016.E().S(util.StringPlural(len(p.RelTeamMembersByUserID), "Member"))
//line views/vuser/Detail.html:267
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:270
	if len(p.RelTeamMembersByUserID) == 0 {
//line views/vuser/Detail.html:270
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vuser/Detail.html:272
	} else {
//line views/vuser/Detail.html:272
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:274
		vtmember.StreamTable(qw422016, p.RelTeamMembersByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:274
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:276
	}
//line views/vuser/Detail.html:276
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-VotesByUserID" type="checkbox" hidden />
        <label for="accordion-VotesByUserID">
          `)
//line views/vuser/Detail.html:282
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vuser/Detail.html:282
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:283
	components.StreamSVGRefIcon(qw422016, `vote-yea`, ps)
//line views/vuser/Detail.html:283
	qw422016.N().S(`
          `)
//line views/vuser/Detail.html:284
	qw422016.E().S(util.StringPlural(len(p.RelVotesByUserID), "Vote"))
//line views/vuser/Detail.html:284
	qw422016.N().S(` by [User ID]
        </label>
        <div class="bd-animated"><div><div>
`)
//line views/vuser/Detail.html:287
	if len(p.RelVotesByUserID) == 0 {
//line views/vuser/Detail.html:287
		qw422016.N().S(`          <em>no related Votes</em>
`)
//line views/vuser/Detail.html:289
	} else {
//line views/vuser/Detail.html:289
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vuser/Detail.html:291
		vvote.StreamTable(qw422016, p.RelVotesByUserID, nil, nil, p.Params, as, ps)
//line views/vuser/Detail.html:291
		qw422016.N().S(`
          </div>
`)
//line views/vuser/Detail.html:293
	}
//line views/vuser/Detail.html:293
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vuser/Detail.html:298
	components.StreamJSONModal(qw422016, "user", "User JSON", p.Model, 1)
//line views/vuser/Detail.html:298
	qw422016.N().S(`
`)
//line views/vuser/Detail.html:299
}

//line views/vuser/Detail.html:299
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vuser/Detail.html:299
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vuser/Detail.html:299
	p.StreamBody(qw422016, as, ps)
//line views/vuser/Detail.html:299
	qt422016.ReleaseWriter(qw422016)
//line views/vuser/Detail.html:299
}

//line views/vuser/Detail.html:299
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vuser/Detail.html:299
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vuser/Detail.html:299
	p.WriteBody(qb422016, as, ps)
//line views/vuser/Detail.html:299
	qs422016 := string(qb422016.B)
//line views/vuser/Detail.html:299
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vuser/Detail.html:299
	return qs422016
//line views/vuser/Detail.html:299
}
