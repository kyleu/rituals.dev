// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vepermission/Table.html:2
package vepermission

//line views/vestimate/vepermission/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vestimate/vepermission/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vepermission/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vepermission/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models epermission.EstimatePermissions, estimatesByEstimateID estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Table.html:12
	qw422016.N().S(`
`)
//line views/vestimate/vepermission/Table.html:13
	prms := params.Sanitized("epermission", ps.Logger)

//line views/vestimate/vepermission/Table.html:13
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vestimate/vepermission/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "epermission", "estimate_id", "Estimate ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:17
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "epermission", "key", "Key", "String text", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:18
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "epermission", "value", "Value", "String text", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:19
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "epermission", "access", "Access", "String text", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:20
	qw422016.N().S(`
        `)
//line views/vestimate/vepermission/Table.html:21
	components.StreamTableHeaderSimple(qw422016, "epermission", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vestimate/vepermission/Table.html:21
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vestimate/vepermission/Table.html:25
	for _, model := range models {
//line views/vestimate/vepermission/Table.html:25
		qw422016.N().S(`      <tr>
        <td class="nowrap">
          <a href="/admin/db/estimate/permission/`)
//line views/vestimate/vepermission/Table.html:28
		view.StreamUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().U(model.Key)
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().U(model.Value)
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:28
		view.StreamUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:28
		if x := estimatesByEstimateID.Get(model.EstimateID); x != nil {
//line views/vestimate/vepermission/Table.html:28
			qw422016.N().S(` (`)
//line views/vestimate/vepermission/Table.html:28
			qw422016.E().S(x.TitleString())
//line views/vestimate/vepermission/Table.html:28
			qw422016.N().S(`)`)
//line views/vestimate/vepermission/Table.html:28
		}
//line views/vestimate/vepermission/Table.html:28
		qw422016.N().S(`</a>
          <a title="Estimate" href="`)
//line views/vestimate/vepermission/Table.html:29
		qw422016.E().S(`/admin/db/estimate` + `/` + model.EstimateID.String())
//line views/vestimate/vepermission/Table.html:29
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:29
		components.StreamSVGRef(qw422016, "estimate", 18, 18, "", ps)
//line views/vestimate/vepermission/Table.html:29
		qw422016.N().S(`</a>
        </td>
        <td><a href="/admin/db/estimate/permission/`)
//line views/vestimate/vepermission/Table.html:31
		view.StreamUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().U(model.Key)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().U(model.Value)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:31
		view.StreamString(qw422016, model.Key)
//line views/vestimate/vepermission/Table.html:31
		qw422016.N().S(`</a></td>
        <td><a href="/admin/db/estimate/permission/`)
//line views/vestimate/vepermission/Table.html:32
		view.StreamUUID(qw422016, &model.EstimateID)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().U(model.Key)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().S(`/`)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().U(model.Value)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Table.html:32
		view.StreamString(qw422016, model.Value)
//line views/vestimate/vepermission/Table.html:32
		qw422016.N().S(`</a></td>
        <td>`)
//line views/vestimate/vepermission/Table.html:33
		view.StreamString(qw422016, model.Access)
//line views/vestimate/vepermission/Table.html:33
		qw422016.N().S(`</td>
        <td>`)
//line views/vestimate/vepermission/Table.html:34
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vestimate/vepermission/Table.html:34
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vepermission/Table.html:36
	}
//line views/vestimate/vepermission/Table.html:37
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vestimate/vepermission/Table.html:37
		qw422016.N().S(`      <tr>
        <td colspan="5">`)
//line views/vestimate/vepermission/Table.html:39
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vestimate/vepermission/Table.html:39
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vestimate/vepermission/Table.html:41
	}
//line views/vestimate/vepermission/Table.html:41
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vestimate/vepermission/Table.html:44
}

//line views/vestimate/vepermission/Table.html:44
func WriteTable(qq422016 qtio422016.Writer, models epermission.EstimatePermissions, estimatesByEstimateID estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Table.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vepermission/Table.html:44
	StreamTable(qw422016, models, estimatesByEstimateID, params, as, ps)
//line views/vestimate/vepermission/Table.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vepermission/Table.html:44
}

//line views/vestimate/vepermission/Table.html:44
func Table(models epermission.EstimatePermissions, estimatesByEstimateID estimate.Estimates, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vepermission/Table.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vepermission/Table.html:44
	WriteTable(qb422016, models, estimatesByEstimateID, params, as, ps)
//line views/vestimate/vepermission/Table.html:44
	qs422016 := string(qb422016.B)
//line views/vestimate/vepermission/Table.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vepermission/Table.html:44
	return qs422016
//line views/vestimate/vepermission/Table.html:44
}
