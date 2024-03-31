// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vretro/vrhistory/Table.html:2
package vrhistory

//line views/vretro/vrhistory/Table.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
)

//line views/vretro/vrhistory/Table.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrhistory/Table.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrhistory/Table.html:12
func StreamTable(qw422016 *qt422016.Writer, models rhistory.RetroHistories, retrosByRetroID retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Table.html:12
	qw422016.N().S(`
`)
//line views/vretro/vrhistory/Table.html:13
	prms := params.Sanitized("rhistory", ps.Logger)

//line views/vretro/vrhistory/Table.html:13
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vretro/vrhistory/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "rhistory", "slug", "Slug", "String text", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:17
	qw422016.N().S(`
        `)
//line views/vretro/vrhistory/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "rhistory", "retro_id", "Retro ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:18
	qw422016.N().S(`
        `)
//line views/vretro/vrhistory/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "rhistory", "retro_name", "Retro Name", "String text", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:19
	qw422016.N().S(`
        `)
//line views/vretro/vrhistory/Table.html:20
	components.StreamTableHeaderSimple(qw422016, "rhistory", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vretro/vrhistory/Table.html:20
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vretro/vrhistory/Table.html:24
	for _, model := range models {
//line views/vretro/vrhistory/Table.html:24
		qw422016.N().S(`      <tr>
        <td><a href="/admin/db/retro/history/`)
//line views/vretro/vrhistory/Table.html:26
		qw422016.N().U(model.Slug)
//line views/vretro/vrhistory/Table.html:26
		qw422016.N().S(`">`)
//line views/vretro/vrhistory/Table.html:26
		view.StreamString(qw422016, model.Slug)
//line views/vretro/vrhistory/Table.html:26
		qw422016.N().S(`</a></td>
        <td class="nowrap">
          `)
//line views/vretro/vrhistory/Table.html:28
		view.StreamUUID(qw422016, &model.RetroID)
//line views/vretro/vrhistory/Table.html:28
		if x := retrosByRetroID.Get(model.RetroID); x != nil {
//line views/vretro/vrhistory/Table.html:28
			qw422016.N().S(` (`)
//line views/vretro/vrhistory/Table.html:28
			qw422016.E().S(x.TitleString())
//line views/vretro/vrhistory/Table.html:28
			qw422016.N().S(`)`)
//line views/vretro/vrhistory/Table.html:28
		}
//line views/vretro/vrhistory/Table.html:28
		qw422016.N().S(`
          <a title="Retro" href="`)
//line views/vretro/vrhistory/Table.html:29
		qw422016.E().S(`/admin/db/retro` + `/` + model.RetroID.String())
//line views/vretro/vrhistory/Table.html:29
		qw422016.N().S(`">`)
//line views/vretro/vrhistory/Table.html:29
		components.StreamSVGRef(qw422016, "retro", 18, 18, "", ps)
//line views/vretro/vrhistory/Table.html:29
		qw422016.N().S(`</a>
        </td>
        <td>`)
//line views/vretro/vrhistory/Table.html:31
		view.StreamString(qw422016, model.RetroName)
//line views/vretro/vrhistory/Table.html:31
		qw422016.N().S(`</td>
        <td>`)
//line views/vretro/vrhistory/Table.html:32
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vretro/vrhistory/Table.html:32
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrhistory/Table.html:34
	}
//line views/vretro/vrhistory/Table.html:35
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vretro/vrhistory/Table.html:35
		qw422016.N().S(`      <tr>
        <td colspan="4">`)
//line views/vretro/vrhistory/Table.html:37
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vretro/vrhistory/Table.html:37
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vretro/vrhistory/Table.html:39
	}
//line views/vretro/vrhistory/Table.html:39
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vretro/vrhistory/Table.html:42
}

//line views/vretro/vrhistory/Table.html:42
func WriteTable(qq422016 qtio422016.Writer, models rhistory.RetroHistories, retrosByRetroID retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Table.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrhistory/Table.html:42
	StreamTable(qw422016, models, retrosByRetroID, params, as, ps)
//line views/vretro/vrhistory/Table.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrhistory/Table.html:42
}

//line views/vretro/vrhistory/Table.html:42
func Table(models rhistory.RetroHistories, retrosByRetroID retro.Retros, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrhistory/Table.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrhistory/Table.html:42
	WriteTable(qb422016, models, retrosByRetroID, params, as, ps)
//line views/vretro/vrhistory/Table.html:42
	qs422016 := string(qb422016.B)
//line views/vretro/vrhistory/Table.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrhistory/Table.html:42
	return qs422016
//line views/vretro/vrhistory/Table.html:42
}
