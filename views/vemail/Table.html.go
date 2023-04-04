// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vemail/Table.html:2
package vemail

//line views/vemail/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
)

//line views/vemail/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vemail/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vemail/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models email.Emails, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vemail/Table.html:11
	qw422016.N().S(`
`)
//line views/vemail/Table.html:12
	prms := params.Get("email", nil, ps.Logger).Sanitize("email")

//line views/vemail/Table.html:12
	qw422016.N().S(`  <table class="mt">
    <thead>
      <tr>
        `)
//line views/vemail/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "email", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vemail/Table.html:16
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "email", "recipients", "Recipients", "Comma-separated list of values", prms, ps.URI, ps)
//line views/vemail/Table.html:17
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "email", "subject", "Subject", "String text", prms, ps.URI, ps)
//line views/vemail/Table.html:18
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "email", "data", "Data", "JSON object", prms, ps.URI, ps)
//line views/vemail/Table.html:19
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "email", "plain", "Plain", "String text", prms, ps.URI, ps)
//line views/vemail/Table.html:20
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "email", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vemail/Table.html:21
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "email", "status", "Status", "String text", prms, ps.URI, ps)
//line views/vemail/Table.html:22
	qw422016.N().S(`
        `)
//line views/vemail/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "email", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vemail/Table.html:23
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vemail/Table.html:27
	for _, model := range models {
//line views/vemail/Table.html:27
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/email/`)
//line views/vemail/Table.html:29
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vemail/Table.html:29
		qw422016.N().S(`">`)
//line views/vemail/Table.html:29
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vemail/Table.html:29
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vemail/Table.html:30
		components.StreamDisplayStringArray(qw422016, model.Recipients)
//line views/vemail/Table.html:30
		qw422016.N().S(`</td>
        <td>`)
//line views/vemail/Table.html:31
		qw422016.E().S(model.Subject)
//line views/vemail/Table.html:31
		qw422016.N().S(`</td>
        <td>`)
//line views/vemail/Table.html:32
		components.StreamJSON(qw422016, model.Data)
//line views/vemail/Table.html:32
		qw422016.N().S(`</td>
        <td>`)
//line views/vemail/Table.html:33
		qw422016.E().S(model.Plain)
//line views/vemail/Table.html:33
		qw422016.N().S(`</td>
        <td class="nowrap">
          `)
//line views/vemail/Table.html:35
		components.StreamDisplayUUID(qw422016, &model.UserID)
//line views/vemail/Table.html:35
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vemail/Table.html:35
			qw422016.N().S(` (`)
//line views/vemail/Table.html:35
			qw422016.E().S(x.TitleString())
//line views/vemail/Table.html:35
			qw422016.N().S(`)`)
//line views/vemail/Table.html:35
		}
//line views/vemail/Table.html:35
		qw422016.N().S(`
          <a title="User" href="`)
//line views/vemail/Table.html:36
		qw422016.E().S(`/admin/db/user` + `/` + model.UserID.String())
//line views/vemail/Table.html:36
		qw422016.N().S(`">`)
//line views/vemail/Table.html:36
		components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vemail/Table.html:36
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vemail/Table.html:38
		qw422016.E().S(model.Status)
//line views/vemail/Table.html:38
		qw422016.N().S(`</td>
        <td>`)
//line views/vemail/Table.html:39
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vemail/Table.html:39
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vemail/Table.html:41
	}
//line views/vemail/Table.html:42
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vemail/Table.html:42
		qw422016.N().S(`      <tr>
        <td colspan="8">`)
//line views/vemail/Table.html:44
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vemail/Table.html:44
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vemail/Table.html:46
	}
//line views/vemail/Table.html:46
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vemail/Table.html:49
}

//line views/vemail/Table.html:49
func WriteTable(qq422016 qtio422016.Writer, models email.Emails, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vemail/Table.html:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/Table.html:49
	StreamTable(qw422016, models, usersByUserID, params, as, ps)
//line views/vemail/Table.html:49
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/Table.html:49
}

//line views/vemail/Table.html:49
func Table(models email.Emails, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vemail/Table.html:49
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/Table.html:49
	WriteTable(qb422016, models, usersByUserID, params, as, ps)
//line views/vemail/Table.html:49
	qs422016 := string(qb422016.B)
//line views/vemail/Table.html:49
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/Table.html:49
	return qs422016
//line views/vemail/Table.html:49
}
