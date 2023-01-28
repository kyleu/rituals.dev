// Code generated by qtc from "Comments.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwutil/Comments.html:1
package vwutil

//line views/vworkspace/vwutil/Comments.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
)

//line views/vworkspace/vwutil/Comments.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwutil/Comments.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwutil/Comments.html:12
func StreamComments(qw422016 *qt422016.Writer, svc enum.ModelService, id uuid.UUID, title string, allComments comment.Comments, members util.Members, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Comments.html:12
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Comments.html:14
	comments := allComments.GetByModel(svc, id)

//line views/vworkspace/vwutil/Comments.html:15
	qw422016.N().S(`  `)
//line views/vworkspace/vwutil/Comments.html:16
	StreamCommentsLink(qw422016, svc, id, title, comments, ps)
//line views/vworkspace/vwutil/Comments.html:16
	qw422016.N().S(`
  `)
//line views/vworkspace/vwutil/Comments.html:17
	StreamCommentsModal(qw422016, svc, id, title, comments, members, ps)
//line views/vworkspace/vwutil/Comments.html:17
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Comments.html:18
}

//line views/vworkspace/vwutil/Comments.html:18
func WriteComments(qq422016 qtio422016.Writer, svc enum.ModelService, id uuid.UUID, title string, allComments comment.Comments, members util.Members, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Comments.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Comments.html:18
	StreamComments(qw422016, svc, id, title, allComments, members, ps)
//line views/vworkspace/vwutil/Comments.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Comments.html:18
}

//line views/vworkspace/vwutil/Comments.html:18
func Comments(svc enum.ModelService, id uuid.UUID, title string, allComments comment.Comments, members util.Members, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Comments.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Comments.html:18
	WriteComments(qb422016, svc, id, title, allComments, members, ps)
//line views/vworkspace/vwutil/Comments.html:18
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Comments.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Comments.html:18
	return qs422016
//line views/vworkspace/vwutil/Comments.html:18
}

//line views/vworkspace/vwutil/Comments.html:20
func StreamCommentsLink(qw422016 *qt422016.Writer, svc enum.ModelService, id uuid.UUID, title string, comments comment.Comments, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Comments.html:20
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Comments.html:22
	icon := "comment-dots"
	if len(comments) == 0 {
		icon = "comment-alt"
	}

//line views/vworkspace/vwutil/Comments.html:26
	qw422016.N().S(`  <a id="comment-link-`)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.E().S(string(svc))
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.E().S(id.String())
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(`" href="#modal-`)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.E().S(string(svc))
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.E().S(id.String())
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(`-comments" title="`)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().D(len(comments))
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(` `)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.E().S(util.StringPluralMaybe(`comment`, len(comments)))
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(`">`)
//line views/vworkspace/vwutil/Comments.html:27
	components.StreamSVGRef(qw422016, icon, 18, 18, "", ps)
//line views/vworkspace/vwutil/Comments.html:27
	qw422016.N().S(`</a>
`)
//line views/vworkspace/vwutil/Comments.html:28
}

//line views/vworkspace/vwutil/Comments.html:28
func WriteCommentsLink(qq422016 qtio422016.Writer, svc enum.ModelService, id uuid.UUID, title string, comments comment.Comments, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Comments.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Comments.html:28
	StreamCommentsLink(qw422016, svc, id, title, comments, ps)
//line views/vworkspace/vwutil/Comments.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Comments.html:28
}

//line views/vworkspace/vwutil/Comments.html:28
func CommentsLink(svc enum.ModelService, id uuid.UUID, title string, comments comment.Comments, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Comments.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Comments.html:28
	WriteCommentsLink(qb422016, svc, id, title, comments, ps)
//line views/vworkspace/vwutil/Comments.html:28
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Comments.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Comments.html:28
	return qs422016
//line views/vworkspace/vwutil/Comments.html:28
}

//line views/vworkspace/vwutil/Comments.html:30
func StreamCommentsModal(qw422016 *qt422016.Writer, svc enum.ModelService, id uuid.UUID, title string, comments comment.Comments, members util.Members, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Comments.html:30
	qw422016.N().S(`
  <div id="modal-`)
//line views/vworkspace/vwutil/Comments.html:31
	qw422016.E().S(string(svc))
//line views/vworkspace/vwutil/Comments.html:31
	qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Comments.html:31
	qw422016.E().S(id.String())
//line views/vworkspace/vwutil/Comments.html:31
	qw422016.N().S(`-comments" class="modal comments" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwutil/Comments.html:36
	qw422016.E().S(title)
//line views/vworkspace/vwutil/Comments.html:36
	qw422016.N().S(` Comments</h2>
      </div>
      <div class="modal-body">
        <ul id="comment-list-`)
//line views/vworkspace/vwutil/Comments.html:39
	qw422016.E().S(string(svc))
//line views/vworkspace/vwutil/Comments.html:39
	qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Comments.html:39
	qw422016.E().S(id.String())
//line views/vworkspace/vwutil/Comments.html:39
	qw422016.N().S(`" class="comment-list">
`)
//line views/vworkspace/vwutil/Comments.html:40
	for _, c := range comments {
//line views/vworkspace/vwutil/Comments.html:42
		un := c.UserID.String()
		if m := members.Get(c.UserID); m != nil {
			un = m.Name
		}

//line views/vworkspace/vwutil/Comments.html:46
		qw422016.N().S(`          <li>
            <div class="right">`)
//line views/vworkspace/vwutil/Comments.html:48
		components.StreamDisplayTimestampRelative(qw422016, &c.Created)
//line views/vworkspace/vwutil/Comments.html:48
		qw422016.N().S(`</div>
            <div>`)
//line views/vworkspace/vwutil/Comments.html:49
		qw422016.N().S(c.HTML)
//line views/vworkspace/vwutil/Comments.html:49
		qw422016.N().S(`</div>
            <div><em>`)
//line views/vworkspace/vwutil/Comments.html:50
		qw422016.E().S(un)
//line views/vworkspace/vwutil/Comments.html:50
		qw422016.N().S(`</em></div>
          </li>
`)
//line views/vworkspace/vwutil/Comments.html:52
	}
//line views/vworkspace/vwutil/Comments.html:52
	qw422016.N().S(`        </ul>
        <form action="#" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwutil/Comments.html:55
	qw422016.E().S(string(action.ActComment))
//line views/vworkspace/vwutil/Comments.html:55
	qw422016.N().S(`" />
          <input type="hidden" name="svc" value="`)
//line views/vworkspace/vwutil/Comments.html:56
	qw422016.E().S(string(svc))
//line views/vworkspace/vwutil/Comments.html:56
	qw422016.N().S(`" />
          <input type="hidden" name="modelID" value="`)
//line views/vworkspace/vwutil/Comments.html:57
	qw422016.E().S(id.String())
//line views/vworkspace/vwutil/Comments.html:57
	qw422016.N().S(`" />
          <div><textarea name="content" placeholder="Add a comment, Markdown and HTML supported"></textarea></div>
          <div class="mt right"><button type="submit">Add Comment</button></div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Comments.html:64
}

//line views/vworkspace/vwutil/Comments.html:64
func WriteCommentsModal(qq422016 qtio422016.Writer, svc enum.ModelService, id uuid.UUID, title string, comments comment.Comments, members util.Members, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Comments.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Comments.html:64
	StreamCommentsModal(qw422016, svc, id, title, comments, members, ps)
//line views/vworkspace/vwutil/Comments.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Comments.html:64
}

//line views/vworkspace/vwutil/Comments.html:64
func CommentsModal(svc enum.ModelService, id uuid.UUID, title string, comments comment.Comments, members util.Members, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Comments.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Comments.html:64
	WriteCommentsModal(qb422016, svc, id, title, comments, members, ps)
//line views/vworkspace/vwutil/Comments.html:64
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Comments.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Comments.html:64
	return qs422016
//line views/vworkspace/vwutil/Comments.html:64
}
