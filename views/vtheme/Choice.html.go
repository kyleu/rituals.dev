// Code generated by qtc from "Choice.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/Choice.html:2
package vtheme

//line views/vtheme/Choice.html:2
import (
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/theme"
	"github.com/kyleu/rituals/views/components"
)

//line views/vtheme/Choice.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Choice.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Choice.html:8
func StreamChoicePanel(qw422016 *qt422016.Writer, themes theme.Themes, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Choice.html:9
	components.StreamIndent(qw422016, true, indent)
//line views/vtheme/Choice.html:9
	qw422016.N().S(`<tr>`)
//line views/vtheme/Choice.html:11
	components.StreamIndent(qw422016, true, indent+1)
//line views/vtheme/Choice.html:11
	qw422016.N().S(`<th class="shrink"><label>Theme</label></th>`)
//line views/vtheme/Choice.html:13
	components.StreamIndent(qw422016, true, indent+1)
//line views/vtheme/Choice.html:13
	qw422016.N().S(`<td>`)
//line views/vtheme/Choice.html:15
	components.StreamIndent(qw422016, true, indent+2)
//line views/vtheme/Choice.html:15
	qw422016.N().S(`<div class="right"><a href="/theme">Edit Themes</a></div><div class="clear"></div>`)
//line views/vtheme/Choice.html:19
	sel := ps.Profile.Theme
	if sel == "" {
		sel = "default"
	}

//line views/vtheme/Choice.html:24
	StreamChoice(qw422016, themes, sel, icon, indent+2, ps)
//line views/vtheme/Choice.html:24
	qw422016.N().S(`</td>`)
//line views/vtheme/Choice.html:26
	components.StreamIndent(qw422016, true, indent)
//line views/vtheme/Choice.html:26
	qw422016.N().S(`</tr>`)
//line views/vtheme/Choice.html:28
}

//line views/vtheme/Choice.html:28
func WriteChoicePanel(qq422016 qtio422016.Writer, themes theme.Themes, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Choice.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Choice.html:28
	StreamChoicePanel(qw422016, themes, icon, indent, ps)
//line views/vtheme/Choice.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Choice.html:28
}

//line views/vtheme/Choice.html:28
func ChoicePanel(themes theme.Themes, icon string, indent int, ps *cutil.PageState) string {
//line views/vtheme/Choice.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Choice.html:28
	WriteChoicePanel(qb422016, themes, icon, indent, ps)
//line views/vtheme/Choice.html:28
	qs422016 := string(qb422016.B)
//line views/vtheme/Choice.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Choice.html:28
	return qs422016
//line views/vtheme/Choice.html:28
}

//line views/vtheme/Choice.html:30
func StreamChoice(qw422016 *qt422016.Writer, themes theme.Themes, selected string, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Choice.html:31
	components.StreamIndent(qw422016, true, indent)
//line views/vtheme/Choice.html:31
	qw422016.N().S(`<div class="choice">`)
//line views/vtheme/Choice.html:33
	for _, t := range themes {
//line views/vtheme/Choice.html:34
		components.StreamIndent(qw422016, true, indent+1)
//line views/vtheme/Choice.html:34
		qw422016.N().S(`<label title="`)
//line views/vtheme/Choice.html:35
		qw422016.E().S(t.Key)
//line views/vtheme/Choice.html:35
		qw422016.N().S(`">`)
//line views/vtheme/Choice.html:36
		if t.Key == selected {
//line views/vtheme/Choice.html:36
			qw422016.N().S(`<input type="radio" name="theme" value="`)
//line views/vtheme/Choice.html:37
			qw422016.E().S(t.Key)
//line views/vtheme/Choice.html:37
			qw422016.N().S(`" checked="checked" />`)
//line views/vtheme/Choice.html:38
		} else {
//line views/vtheme/Choice.html:38
			qw422016.N().S(`<input type="radio" name="theme" value="`)
//line views/vtheme/Choice.html:39
			qw422016.E().S(t.Key)
//line views/vtheme/Choice.html:39
			qw422016.N().S(`" />`)
//line views/vtheme/Choice.html:40
		}
//line views/vtheme/Choice.html:41
		StreamMockupTheme(qw422016, t, true, icon, indent+2, ps)
//line views/vtheme/Choice.html:41
		qw422016.N().S(`</label>`)
//line views/vtheme/Choice.html:43
	}
//line views/vtheme/Choice.html:44
	components.StreamIndent(qw422016, true, indent)
//line views/vtheme/Choice.html:44
	qw422016.N().S(`</div>`)
//line views/vtheme/Choice.html:46
}

//line views/vtheme/Choice.html:46
func WriteChoice(qq422016 qtio422016.Writer, themes theme.Themes, selected string, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Choice.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Choice.html:46
	StreamChoice(qw422016, themes, selected, icon, indent, ps)
//line views/vtheme/Choice.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Choice.html:46
}

//line views/vtheme/Choice.html:46
func Choice(themes theme.Themes, selected string, icon string, indent int, ps *cutil.PageState) string {
//line views/vtheme/Choice.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Choice.html:46
	WriteChoice(qb422016, themes, selected, icon, indent, ps)
//line views/vtheme/Choice.html:46
	qs422016 := string(qb422016.B)
//line views/vtheme/Choice.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Choice.html:46
	return qs422016
//line views/vtheme/Choice.html:46
}
