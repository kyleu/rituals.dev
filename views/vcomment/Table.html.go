// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vcomment/Table.html:1
package vcomment

//line views/vcomment/Table.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vcomment/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcomment/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcomment/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models comment.Comments, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vcomment/Table.html:12
	qw422016.N().S(`
`)
//line views/vcomment/Table.html:13
	prms := params.Sanitized("comment", ps.Logger)

//line views/vcomment/Table.html:13
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vcomment/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "comment", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcomment/Table.html:18
	qw422016.N().S(`
          `)
//line views/vcomment/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "comment", "svc", "Svc", enum.AllModelServices.Help(), prms, ps.URI, ps)
//line views/vcomment/Table.html:19
	qw422016.N().S(`
          `)
//line views/vcomment/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "comment", "model_id", "Model ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcomment/Table.html:20
	qw422016.N().S(`
          `)
//line views/vcomment/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "comment", "user_id", "User ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcomment/Table.html:21
	qw422016.N().S(`
          `)
//line views/vcomment/Table.html:22
	components.StreamTableHeaderSimple(qw422016, "comment", "content", "Content", "String text", prms, ps.URI, ps)
//line views/vcomment/Table.html:22
	qw422016.N().S(`
          `)
//line views/vcomment/Table.html:23
	components.StreamTableHeaderSimple(qw422016, "comment", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vcomment/Table.html:23
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vcomment/Table.html:27
	for _, model := range models {
//line views/vcomment/Table.html:27
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vcomment/Table.html:29
		qw422016.E().S(model.WebPath(paths...))
//line views/vcomment/Table.html:29
		qw422016.N().S(`">`)
//line views/vcomment/Table.html:29
		view.StreamUUID(qw422016, &model.ID)
//line views/vcomment/Table.html:29
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vcomment/Table.html:30
		qw422016.E().S(model.Svc.String())
//line views/vcomment/Table.html:30
		qw422016.N().S(`</td>
          <td>`)
//line views/vcomment/Table.html:31
		view.StreamUUID(qw422016, &model.ModelID)
//line views/vcomment/Table.html:31
		qw422016.N().S(`</td>
          <td class="nowrap">
            `)
//line views/vcomment/Table.html:33
		view.StreamUUID(qw422016, &model.UserID)
//line views/vcomment/Table.html:33
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vcomment/Table.html:33
			qw422016.N().S(` (`)
//line views/vcomment/Table.html:33
			qw422016.E().S(x.TitleString())
//line views/vcomment/Table.html:33
			qw422016.N().S(`)`)
//line views/vcomment/Table.html:33
		}
//line views/vcomment/Table.html:33
		qw422016.N().S(`
            <a title="User" href="`)
//line views/vcomment/Table.html:34
		if x := usersByUserID.Get(model.UserID); x != nil {
//line views/vcomment/Table.html:34
			qw422016.E().S(x.WebPath(paths...))
//line views/vcomment/Table.html:34
		}
//line views/vcomment/Table.html:34
		qw422016.N().S(`">`)
//line views/vcomment/Table.html:34
		components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vcomment/Table.html:34
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vcomment/Table.html:36
		view.StreamString(qw422016, model.Content)
//line views/vcomment/Table.html:36
		qw422016.N().S(`</td>
          <td>`)
//line views/vcomment/Table.html:37
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vcomment/Table.html:37
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vcomment/Table.html:39
	}
//line views/vcomment/Table.html:39
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vcomment/Table.html:43
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vcomment/Table.html:43
		qw422016.N().S(`  <hr />
  `)
//line views/vcomment/Table.html:45
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vcomment/Table.html:45
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vcomment/Table.html:47
	}
//line views/vcomment/Table.html:48
}

//line views/vcomment/Table.html:48
func WriteTable(qq422016 qtio422016.Writer, models comment.Comments, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vcomment/Table.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcomment/Table.html:48
	StreamTable(qw422016, models, usersByUserID, params, as, ps, paths...)
//line views/vcomment/Table.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/vcomment/Table.html:48
}

//line views/vcomment/Table.html:48
func Table(models comment.Comments, usersByUserID user.Users, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vcomment/Table.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcomment/Table.html:48
	WriteTable(qb422016, models, usersByUserID, params, as, ps, paths...)
//line views/vcomment/Table.html:48
	qs422016 := string(qb422016.B)
//line views/vcomment/Table.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcomment/Table.html:48
	return qs422016
//line views/vcomment/Table.html:48
}
