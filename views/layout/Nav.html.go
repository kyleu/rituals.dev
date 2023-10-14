// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vutil"
)

//line views/layout/Nav.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:11
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:11
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:13
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:13
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:13
	qw422016.E().S(util.AppName)
//line views/layout/Nav.html:13
	qw422016.N().S(` `)
//line views/layout/Nav.html:13
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:13
	qw422016.N().S(`">`)
//line views/layout/Nav.html:13
	components.StreamSVGRef(qw422016, ps.RootIcon, 32, 32, ``, ps)
//line views/layout/Nav.html:13
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a href="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:15
	qw422016.N().S(`" class="nav-root-icon" title="`)
//line views/layout/Nav.html:15
	qw422016.E().S(util.AppName)
//line views/layout/Nav.html:15
	qw422016.N().S(`">`)
//line views/layout/Nav.html:15
	components.StreamSVGRef(qw422016, ps.RootIcon, 18, 28, "breadcrumb-icon", ps)
//line views/layout/Nav.html:15
	qw422016.N().S(`</a>
    <a class="link nav-root-item" href="`)
//line views/layout/Nav.html:16
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:16
	qw422016.N().S(`">`)
//line views/layout/Nav.html:16
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:16
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:16
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:16
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:18
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:18
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:19
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:19
		qw422016.N().S(`" class="search" title="search">
    <input type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:23
	}
//line views/layout/Nav.html:23
	qw422016.N().S(`  <a class="profile" title="`)
//line views/layout/Nav.html:24
	qw422016.E().S(ps.AuthString())
//line views/layout/Nav.html:24
	qw422016.N().S(`" href="`)
//line views/layout/Nav.html:24
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:24
	qw422016.N().S(`">
`)
//line views/layout/Nav.html:25
	if i := ps.Accounts.Image(); i != "" {
//line views/layout/Nav.html:25
		qw422016.N().S(`    <img style="width: 24px; height: 24px;" src="`)
//line views/layout/Nav.html:26
		qw422016.E().S(i)
//line views/layout/Nav.html:26
		qw422016.N().S(`" />
`)
//line views/layout/Nav.html:27
	} else {
//line views/layout/Nav.html:27
		qw422016.N().S(`    `)
//line views/layout/Nav.html:28
		components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:28
		qw422016.N().S(`
`)
//line views/layout/Nav.html:29
	}
//line views/layout/Nav.html:29
	qw422016.N().S(`  </a>
`)
//line views/layout/Nav.html:31
	if !ps.HideMenu {
//line views/layout/Nav.html:31
		qw422016.N().S(`  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:34
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:34
		qw422016.N().S(`
`)
//line views/layout/Nav.html:35
	}
//line views/layout/Nav.html:35
	qw422016.N().S(`</nav>`)
//line views/layout/Nav.html:36
}

//line views/layout/Nav.html:36
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:36
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:36
}

//line views/layout/Nav.html:36
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:36
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:36
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:36
	return qs422016
//line views/layout/Nav.html:36
}

//line views/layout/Nav.html:38
func StreamNavItem(qw422016 *qt422016.Writer, link string, title string, icon string, last bool, ps *cutil.PageState) {
//line views/layout/Nav.html:39
	if link != "" {
//line views/layout/Nav.html:39
		qw422016.N().S(`<a class="link`)
//line views/layout/Nav.html:40
		if last {
//line views/layout/Nav.html:40
			qw422016.N().S(` `)
//line views/layout/Nav.html:40
			qw422016.N().S(`last`)
//line views/layout/Nav.html:40
		}
//line views/layout/Nav.html:40
		qw422016.N().S(`" href="`)
//line views/layout/Nav.html:40
		qw422016.E().S(link)
//line views/layout/Nav.html:40
		qw422016.N().S(`">`)
//line views/layout/Nav.html:41
	}
//line views/layout/Nav.html:41
	qw422016.N().S(`<span title="`)
//line views/layout/Nav.html:42
	qw422016.E().S(title)
//line views/layout/Nav.html:42
	qw422016.N().S(`">`)
//line views/layout/Nav.html:42
	components.StreamSVGRef(qw422016, icon, 18, 28, "breadcrumb-icon", ps)
//line views/layout/Nav.html:42
	qw422016.N().S(`</span><span class="nav-item-title">`)
//line views/layout/Nav.html:43
	qw422016.E().S(title)
//line views/layout/Nav.html:43
	qw422016.N().S(`</span>`)
//line views/layout/Nav.html:44
	if link != "" {
//line views/layout/Nav.html:44
		qw422016.N().S(`</a>`)
//line views/layout/Nav.html:46
	}
//line views/layout/Nav.html:47
}

//line views/layout/Nav.html:47
func WriteNavItem(qq422016 qtio422016.Writer, link string, title string, icon string, last bool, ps *cutil.PageState) {
//line views/layout/Nav.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:47
	StreamNavItem(qw422016, link, title, icon, last, ps)
//line views/layout/Nav.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:47
}

//line views/layout/Nav.html:47
func NavItem(link string, title string, icon string, last bool, ps *cutil.PageState) string {
//line views/layout/Nav.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:47
	WriteNavItem(qb422016, link, title, icon, last, ps)
//line views/layout/Nav.html:47
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:47
	return qs422016
//line views/layout/Nav.html:47
}

//line views/layout/Nav.html:49
func StreamNavItems(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:50
	for idx, bc := range ps.Breadcrumbs {
//line views/layout/Nav.html:52
		i := ps.Menu.GetByPath(ps.Breadcrumbs[:idx+1])
		if i == nil {
			i = menu.ItemFromString(bc)
		}

//line views/layout/Nav.html:57
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:57
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:59
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:60
		StreamNavItem(qw422016, i.Route, i.Title, i.Icon, idx == len(ps.Breadcrumbs)-1, ps)
//line views/layout/Nav.html:61
	}
//line views/layout/Nav.html:62
}

//line views/layout/Nav.html:62
func WriteNavItems(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:62
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:62
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:62
}

//line views/layout/Nav.html:62
func NavItems(ps *cutil.PageState) string {
//line views/layout/Nav.html:62
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:62
	WriteNavItems(qb422016, ps)
//line views/layout/Nav.html:62
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:62
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:62
	return qs422016
//line views/layout/Nav.html:62
}
