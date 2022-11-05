// Code generated by qtc from "Members.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/Members.html:1
package vworkspace

//line views/vworkspace/Members.html:1
import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/enum"
)

//line views/vworkspace/Members.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/Members.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/Members.html:9
func StreamMemberRow(qw422016 *qt422016.Writer, userID uuid.UUID, name string, picture string, role enum.MemberStatus, updated *time.Time) {
//line views/vworkspace/Members.html:9
	qw422016.N().S(`
  <tr>
    <td><a href="#modal-member-`)
//line views/vworkspace/Members.html:11
	qw422016.E().S(userID.String())
//line views/vworkspace/Members.html:11
	qw422016.N().S(`">`)
//line views/vworkspace/Members.html:11
	qw422016.E().S(name)
//line views/vworkspace/Members.html:11
	qw422016.N().S(`</a></td>
    <td class="shrink">`)
//line views/vworkspace/Members.html:12
	qw422016.E().S(string(role))
//line views/vworkspace/Members.html:12
	qw422016.N().S(`</td>
  </tr>
`)
//line views/vworkspace/Members.html:14
}

//line views/vworkspace/Members.html:14
func WriteMemberRow(qq422016 qtio422016.Writer, userID uuid.UUID, name string, picture string, role enum.MemberStatus, updated *time.Time) {
//line views/vworkspace/Members.html:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/Members.html:14
	StreamMemberRow(qw422016, userID, name, picture, role, updated)
//line views/vworkspace/Members.html:14
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/Members.html:14
}

//line views/vworkspace/Members.html:14
func MemberRow(userID uuid.UUID, name string, picture string, role enum.MemberStatus, updated *time.Time) string {
//line views/vworkspace/Members.html:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/Members.html:14
	WriteMemberRow(qb422016, userID, name, picture, role, updated)
//line views/vworkspace/Members.html:14
	qs422016 := string(qb422016.B)
//line views/vworkspace/Members.html:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/Members.html:14
	return qs422016
//line views/vworkspace/Members.html:14
}

//line views/vworkspace/Members.html:16
func StreamMemberModal(qw422016 *qt422016.Writer, userID uuid.UUID, name string, picture string, role enum.MemberStatus, updated *time.Time) {
//line views/vworkspace/Members.html:16
	qw422016.N().S(`
  <div id="modal-member-`)
//line views/vworkspace/Members.html:17
	qw422016.E().S(userID.String())
//line views/vworkspace/Members.html:17
	qw422016.N().S(`" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/Members.html:22
	qw422016.E().S(name)
//line views/vworkspace/Members.html:22
	qw422016.N().S(`</h2>
      </div>
      <div class="modal-body">
        TODO
      </div>
    </div>
  </div>
`)
//line views/vworkspace/Members.html:29
}

//line views/vworkspace/Members.html:29
func WriteMemberModal(qq422016 qtio422016.Writer, userID uuid.UUID, name string, picture string, role enum.MemberStatus, updated *time.Time) {
//line views/vworkspace/Members.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/Members.html:29
	StreamMemberModal(qw422016, userID, name, picture, role, updated)
//line views/vworkspace/Members.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/Members.html:29
}

//line views/vworkspace/Members.html:29
func MemberModal(userID uuid.UUID, name string, picture string, role enum.MemberStatus, updated *time.Time) string {
//line views/vworkspace/Members.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/Members.html:29
	WriteMemberModal(qb422016, userID, name, picture, role, updated)
//line views/vworkspace/Members.html:29
	qs422016 := string(qb422016.B)
//line views/vworkspace/Members.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/Members.html:29
	return qs422016
//line views/vworkspace/Members.html:29
}
