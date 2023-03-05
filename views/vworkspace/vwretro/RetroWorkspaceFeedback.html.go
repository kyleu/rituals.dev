// Code generated by qtc from "RetroWorkspaceFeedback.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:1
package vwretro

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:1
import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
func StreamRetroWorkspaceFeedbacks(qw422016 *qt422016.Writer, userID uuid.UUID, username string, w *workspace.FullRetro, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:14
	qw422016.N().S(`
  <div id="category-list">
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:16
	for _, grp := range w.Feedbacks.Grouped(w.Retro.Categories) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:16
		qw422016.N().S(`    <div id="category-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:17
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:17
		qw422016.N().S(`" data-category="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:17
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:17
		qw422016.N().S(`" class="category">
      <div class="right"><a class="add-feedback-link" data-category="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
		qw422016.N().S(`" href="#modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:18
		qw422016.N().S(`">
        <button>New</button>
      </a></div>
      <h4><a href="#modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:21
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:21
		qw422016.N().S(`">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:21
		qw422016.E().S(grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:21
		qw422016.N().S(`</a></h4>
      <div class="clear"></div>
      <div class="feedback-list">
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:24
		for _, f := range grp.Feedbacks {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:24
			qw422016.N().S(`      <div id="feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:25
			qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:25
			qw422016.N().S(`" class="feedback mt clear">
        <div class="right">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:26
			vwutil.StreamComments(qw422016, enum.ModelServiceFeedback, f.ID, f.TitleString(), w.Comments, w.UtilMembers, "member-icon", ps)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:26
			qw422016.N().S(`</div>
        <a href="#modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.N().S(`" class="clean modal-feedback-edit-link" data-id="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:27
			qw422016.N().S(`">
          <div class="member-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:28
			qw422016.E().S(f.UserID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:28
			qw422016.N().S(`-name">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:28
			qw422016.E().S(w.UtilMembers.Name(&f.UserID))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:28
			qw422016.N().S(`</div>
          <div class="pt feedback-content">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:29
			qw422016.N().S(f.HTML)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:29
			qw422016.N().S(`</div>
        </a>
      </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:32
		}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:32
		qw422016.N().S(`      </div>
      `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:34
		StreamRetroWorkspaceFeedbackModalAdd(qw422016, w.Retro.Categories, grp.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:34
		qw422016.N().S(`
    </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:36
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:36
	qw422016.N().S(`  </div>
  <div id="feedback-modals">
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:39
	for _, f := range w.Feedbacks {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:40
		if userID == f.UserID {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:40
			qw422016.N().S(`  `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
			StreamRetroWorkspaceFeedbackModalEdit(qw422016, f, w.Retro.Categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:41
			qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:42
		} else {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:42
			qw422016.N().S(`  `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:43
			StreamRetroWorkspaceFeedbackModalView(qw422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:43
			qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:44
		}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:45
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:45
	qw422016.N().S(`  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
func WriteRetroWorkspaceFeedbacks(qq422016 qtio422016.Writer, userID uuid.UUID, username string, w *workspace.FullRetro, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	StreamRetroWorkspaceFeedbacks(qw422016, userID, username, w, ps)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
func RetroWorkspaceFeedbacks(userID uuid.UUID, username string, w *workspace.FullRetro, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	WriteRetroWorkspaceFeedbacks(qb422016, userID, username, w, ps)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:47
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:49
func StreamRetroWorkspaceFeedbackModalAdd(qw422016 *qt422016.Writer, categories []string, category string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:49
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:50
	if !slices.Contains(categories, category) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:51
		categories = append(categories, category)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:52
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:52
	qw422016.N().S(`  <div id="modal-feedback--add-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:53
	qw422016.E().S(category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:53
	qw422016.N().S(`" class="modal modal-feedback-add" data-category="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:53
	qw422016.E().S(category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:53
	qw422016.N().S(`" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>New Feedback</h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:62
	qw422016.E().S(string(action.ActChildAdd))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:62
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:63
	components.StreamFormVerticalSelect(qw422016, "category", "", "Category", category, categories, categories, 5)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:63
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:64
	components.StreamFormVerticalTextarea(qw422016, "content", "feedback-add-content-"+category, "Content", 8, "", 5, "HTML and Markdown supported")
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:64
	qw422016.N().S(`
          <div class="right"><button type="submit">Add Feedback</button></div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
func WriteRetroWorkspaceFeedbackModalAdd(qq422016 qtio422016.Writer, categories []string, category string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	StreamRetroWorkspaceFeedbackModalAdd(qw422016, categories, category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
func RetroWorkspaceFeedbackModalAdd(categories []string, category string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	WriteRetroWorkspaceFeedbackModalAdd(qb422016, categories, category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:70
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:72
func StreamRetroWorkspaceFeedbackModalEdit(qw422016 *qt422016.Writer, f *feedback.Feedback, categories []string, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:72
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:73
	if !slices.Contains(categories, f.Category) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:74
		categories = append(categories, f.Category)

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	}
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:75
	qw422016.N().S(`  <div id="modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:76
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:76
	qw422016.N().S(`" class="modal modal-feedback-edit" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.E().S(f.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.N().S(` :: <span class="member-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.E().S(f.UserID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.E().S(username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:81
	qw422016.N().S(`</span></h2>
      </div>
      <div class="modal-body">
        <form action="#" method="post">
          <input type="hidden" name="feedbackID" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:85
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:85
	qw422016.N().S(`" />
          <table class="mt expanded">
            <tbody>
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:88
	components.StreamFormVerticalSelect(qw422016, "category", "", "Category", string(f.Category), categories, categories, 5)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:88
	qw422016.N().S(`
              `)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:89
	components.StreamFormVerticalTextarea(qw422016, "content", "input-content-"+f.ID.String(), "Content", 8, f.Content, 5, "HTML and Markdown supported")
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:89
	qw422016.N().S(`
              <tr><td colspan="2">
                <div class="right"><button class="feedback-edit-save" type="submit" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:91
	qw422016.E().S(string(action.ActChildUpdate))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:91
	qw422016.N().S(`">Save Changes</button></div>
                <button class="feedback-edit-delete" type="submit" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:92
	qw422016.E().S(string(action.ActChildRemove))
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:92
	qw422016.N().S(`" onclick="return confirm('Are you sure you want to delete this feedback?');">Delete</button>
              </td></tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
func WriteRetroWorkspaceFeedbackModalEdit(qq422016 qtio422016.Writer, f *feedback.Feedback, categories []string, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	StreamRetroWorkspaceFeedbackModalEdit(qw422016, f, categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
func RetroWorkspaceFeedbackModalEdit(f *feedback.Feedback, categories []string, username string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	WriteRetroWorkspaceFeedbackModalEdit(qb422016, f, categories, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:100
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:102
func StreamRetroWorkspaceFeedbackModalView(qw422016 *qt422016.Writer, f *feedback.Feedback, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:102
	qw422016.N().S(`
<div id="modal-feedback-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:103
	qw422016.E().S(f.ID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:103
	qw422016.N().S(`" class="modal" style="display: none;">
  <a class="backdrop" href="#"></a>
  <div class="modal-content">
    <div class="modal-header">
      <a href="#" class="modal-close">×</a>
      <h2>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.E().S(f.Category)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.N().S(` :: <span class="member-`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.E().S(f.UserID.String())
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.E().S(username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:108
	qw422016.N().S(`</span></h2>
    </div>
    <div class="modal-body">
      <div>`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qw422016.N().S(f.HTML)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:111
	qw422016.N().S(`</div>
    </div>
  </div>
</div>
`)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
func WriteRetroWorkspaceFeedbackModalView(qq422016 qtio422016.Writer, f *feedback.Feedback, username string) {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	StreamRetroWorkspaceFeedbackModalView(qw422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
}

//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
func RetroWorkspaceFeedbackModalView(f *feedback.Feedback, username string) string {
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	WriteRetroWorkspaceFeedbackModalView(qb422016, f, username)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspaceFeedback.html:115
}
