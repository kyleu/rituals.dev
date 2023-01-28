// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vemember/Edit.html:2
package vemember

//line views/vestimate/vemember/Edit.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vemember/Edit.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vemember/Edit.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vemember/Edit.html:10
type Edit struct {
	layout.Basic
	Model *emember.EstimateMember
	IsNew bool
}

//line views/vestimate/vemember/Edit.html:16
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vemember/Edit.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vestimate/vemember/Edit.html:18
	if p.IsNew {
//line views/vestimate/vemember/Edit.html:18
		qw422016.N().S(`    <div class="right"><a href="/admin/db/estimate/member/random"><button>Random</button></a></div>
    <h3>`)
//line views/vestimate/vemember/Edit.html:20
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vestimate/vemember/Edit.html:20
		qw422016.N().S(` New Member</h3>
    <form action="/admin/db/estimate/member/new" class="mt" method="post">
`)
//line views/vestimate/vemember/Edit.html:22
	} else {
//line views/vestimate/vemember/Edit.html:22
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vestimate/vemember/Edit.html:23
		qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vemember/Edit.html:23
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete member [`)
//line views/vestimate/vemember/Edit.html:23
		qw422016.E().S(p.Model.String())
//line views/vestimate/vemember/Edit.html:23
		qw422016.N().S(`]?')"><button>Delete</button></a></div>
    <h3>`)
//line views/vestimate/vemember/Edit.html:24
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vestimate/vemember/Edit.html:24
		qw422016.N().S(` Edit Member [`)
//line views/vestimate/vemember/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vestimate/vemember/Edit.html:24
		qw422016.N().S(`]</h3>
    <form action="" method="post">
`)
//line views/vestimate/vemember/Edit.html:26
	}
//line views/vestimate/vemember/Edit.html:26
	qw422016.N().S(`      <table class="mt expanded">
        <tbody>
          `)
//line views/vestimate/vemember/Edit.html:29
	if p.IsNew {
//line views/vestimate/vemember/Edit.html:29
		components.StreamTableInputUUID(qw422016, "estimateID", "", "Estimate ID", &p.Model.EstimateID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vemember/Edit.html:29
	}
//line views/vestimate/vemember/Edit.html:29
	qw422016.N().S(`
          `)
//line views/vestimate/vemember/Edit.html:30
	if p.IsNew {
//line views/vestimate/vemember/Edit.html:30
		components.StreamTableInputUUID(qw422016, "userID", "", "User ID", &p.Model.UserID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vestimate/vemember/Edit.html:30
	}
//line views/vestimate/vemember/Edit.html:30
	qw422016.N().S(`
          `)
//line views/vestimate/vemember/Edit.html:31
	components.StreamTableInput(qw422016, "name", "", "Name", p.Model.Name, 5, "String text")
//line views/vestimate/vemember/Edit.html:31
	qw422016.N().S(`
          `)
//line views/vestimate/vemember/Edit.html:32
	components.StreamTableInput(qw422016, "picture", "", "Picture", p.Model.Picture, 5, "URL in string form")
//line views/vestimate/vemember/Edit.html:32
	qw422016.N().S(`
          `)
//line views/vestimate/vemember/Edit.html:33
	components.StreamTableSelect(qw422016, "role", "", "Role", string(p.Model.Role), []string{"owner", "member", "observer"}, []string{"owner", "member", "observer"}, 5, "Available options: [owner, member, observer]")
//line views/vestimate/vemember/Edit.html:33
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      rituals.autocomplete(document.getElementById("input-estimateID"), "/admin/db/estimate?estimate.l=10", "q", (o) => o["slug"] + " / " + o["title"] + " / " + o["choices"] + " (" + o["id"] + ")", (o) => o["id"]);
      rituals.autocomplete(document.getElementById("input-userID"), "/admin/db/user?user.l=10", "q", (o) => o["name"] + " (" + o["id"] + ")", (o) => o["id"]);
    });
  </script>
`)
//line views/vestimate/vemember/Edit.html:45
}

//line views/vestimate/vemember/Edit.html:45
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vemember/Edit.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vemember/Edit.html:45
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vemember/Edit.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vemember/Edit.html:45
}

//line views/vestimate/vemember/Edit.html:45
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vemember/Edit.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vemember/Edit.html:45
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vemember/Edit.html:45
	qs422016 := string(qb422016.B)
//line views/vestimate/vemember/Edit.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vemember/Edit.html:45
	return qs422016
//line views/vestimate/vemember/Edit.html:45
}
