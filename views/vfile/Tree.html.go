// Code generated by qtc from "Tree.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/Tree.html:2
package vfile

//line views/vfile/Tree.html:2
import (
	"path"
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filesystem"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vutil"
)

//line views/vfile/Tree.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/Tree.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/Tree.html:14
func StreamTree(qw422016 *qt422016.Writer, t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Tree.html:15
	vutil.StreamIndent(qw422016, true, 2)
//line views/vfile/Tree.html:15
	qw422016.N().S(`<ul class="accordion">`)
//line views/vfile/Tree.html:17
	streamtreeNodes(qw422016, t.Nodes, "", urlPrefix, actions, as, ps, 2)
//line views/vfile/Tree.html:18
	vutil.StreamIndent(qw422016, true, 2)
//line views/vfile/Tree.html:18
	qw422016.N().S(`</ul>`)
//line views/vfile/Tree.html:20
}

//line views/vfile/Tree.html:20
func WriteTree(qq422016 qtio422016.Writer, t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Tree.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:20
	StreamTree(qw422016, t, urlPrefix, actions, as, ps)
//line views/vfile/Tree.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:20
}

//line views/vfile/Tree.html:20
func Tree(t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/Tree.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:20
	WriteTree(qb422016, t, urlPrefix, actions, as, ps)
//line views/vfile/Tree.html:20
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:20
	return qs422016
//line views/vfile/Tree.html:20
}

//line views/vfile/Tree.html:22
func streamtreeNode(qw422016 *qt422016.Writer, n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:24
	pathID := n.Name
	if pth != "" {
		pathID = pth + "/" + pathID
	}
	pathID = strings.ReplaceAll(pathID, "/", "__")

//line views/vfile/Tree.html:30
	vutil.StreamIndent(qw422016, true, indent)
//line views/vfile/Tree.html:30
	qw422016.N().S(`<li>`)
//line views/vfile/Tree.html:32
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:32
	qw422016.N().S(`<input id="accordion-`)
//line views/vfile/Tree.html:33
	qw422016.E().S(pathID)
//line views/vfile/Tree.html:33
	qw422016.N().S(`" type="checkbox" hidden />`)
//line views/vfile/Tree.html:34
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:34
	qw422016.N().S(`<label for="accordion-`)
//line views/vfile/Tree.html:35
	qw422016.E().S(pathID)
//line views/vfile/Tree.html:35
	qw422016.N().S(`">`)
//line views/vfile/Tree.html:37
	var acts map[string]string
	if actions != nil {
		acts = actions(pth, n)
	}

//line views/vfile/Tree.html:41
	qw422016.N().S(`<div class="right">`)
//line views/vfile/Tree.html:43
	if n.Size > 0 {
//line views/vfile/Tree.html:43
		qw422016.N().S(`<em>`)
//line views/vfile/Tree.html:44
		qw422016.E().S(util.ByteSizeSI(int64(n.Size)))
//line views/vfile/Tree.html:44
		qw422016.N().S(`</em>`)
//line views/vfile/Tree.html:44
		qw422016.N().S(` `)
//line views/vfile/Tree.html:45
	}
//line views/vfile/Tree.html:46
	for k, v := range acts {
//line views/vfile/Tree.html:46
		qw422016.N().S(`<a href="`)
//line views/vfile/Tree.html:47
		qw422016.E().S(v)
//line views/vfile/Tree.html:47
		qw422016.N().S(`">`)
//line views/vfile/Tree.html:47
		qw422016.E().S(k)
//line views/vfile/Tree.html:47
		qw422016.N().S(`</a>`)
//line views/vfile/Tree.html:48
	}
//line views/vfile/Tree.html:48
	qw422016.N().S(`</div>`)
//line views/vfile/Tree.html:50
	components.StreamExpandCollapse(qw422016, indent+2, ps)
//line views/vfile/Tree.html:51
	if n.Dir {
//line views/vfile/Tree.html:52
		components.StreamSVGRef(qw422016, `folder`, 15, 15, ``, ps)
//line views/vfile/Tree.html:53
	} else {
//line views/vfile/Tree.html:54
		components.StreamSVGRef(qw422016, `file`, 15, 15, ``, ps)
//line views/vfile/Tree.html:55
	}
//line views/vfile/Tree.html:56
	qw422016.N().S(` `)
//line views/vfile/Tree.html:56
	qw422016.E().S(n.Name)
//line views/vfile/Tree.html:57
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:57
	qw422016.N().S(`</label>`)
//line views/vfile/Tree.html:59
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:59
	qw422016.N().S(`<div class="bd"><div><div>`)
//line views/vfile/Tree.html:61
	if len(n.Children) == 0 {
//line views/vfile/Tree.html:62
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:62
		qw422016.N().S(`<div>`)
//line views/vfile/Tree.html:63
		qw422016.E().S(n.Name)
//line views/vfile/Tree.html:63
		qw422016.N().S(`</div>`)
//line views/vfile/Tree.html:64
	} else {
//line views/vfile/Tree.html:65
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:65
		qw422016.N().S(`<ul class="accordion" style="margin-left:`)
//line views/vfile/Tree.html:66
		qw422016.N().D((indent / 3) * 6)
//line views/vfile/Tree.html:66
		qw422016.N().S(`px; margin-bottom: 0;">`)
//line views/vfile/Tree.html:67
		streamtreeNodes(qw422016, n.Children, path.Join(pth, n.Name), urlPrefix, actions, as, ps, indent+3)
//line views/vfile/Tree.html:68
		vutil.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:68
		qw422016.N().S(`</ul>`)
//line views/vfile/Tree.html:70
	}
//line views/vfile/Tree.html:71
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:71
	qw422016.N().S(`</div></div></div>`)
//line views/vfile/Tree.html:73
	vutil.StreamIndent(qw422016, true, indent)
//line views/vfile/Tree.html:73
	qw422016.N().S(`</li>`)
//line views/vfile/Tree.html:75
}

//line views/vfile/Tree.html:75
func writetreeNode(qq422016 qtio422016.Writer, n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:75
	streamtreeNode(qw422016, n, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:75
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:75
}

//line views/vfile/Tree.html:75
func treeNode(n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) string {
//line views/vfile/Tree.html:75
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:75
	writetreeNode(qb422016, n, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:75
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:75
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:75
	return qs422016
//line views/vfile/Tree.html:75
}

//line views/vfile/Tree.html:77
func streamtreeNodes(qw422016 *qt422016.Writer, nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:78
	for _, n := range nodes {
//line views/vfile/Tree.html:79
		streamtreeNode(qw422016, n, pth, urlPrefix, actions, as, ps, indent+1)
//line views/vfile/Tree.html:80
	}
//line views/vfile/Tree.html:81
}

//line views/vfile/Tree.html:81
func writetreeNodes(qq422016 qtio422016.Writer, nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:81
	streamtreeNodes(qw422016, nodes, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:81
}

//line views/vfile/Tree.html:81
func treeNodes(nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) string {
//line views/vfile/Tree.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:81
	writetreeNodes(qb422016, nodes, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:81
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:81
	return qs422016
//line views/vfile/Tree.html:81
}
