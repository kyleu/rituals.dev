// Code generated by qtc from "seed_team.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_team.sql:1
package seeddata

//line queries/seeddata/seed_team.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_team.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_team.sql:1
func StreamTeamSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_team.sql:1
	qw422016.N().S(`
insert into "team" (
  "id", "slug", "title", "icon", "status", "created", "updated"
) values (
  '10000000-0000-0000-0000-000000000000', 'rituals-team', 'Rituals Team', 'star', 'active', now(), null
), (
  '10000001-0000-0000-0000-000000000000', 'team-2', 'Team 2', 'action', 'active', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_team.sql:9
}

//line queries/seeddata/seed_team.sql:9
func WriteTeamSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_team.sql:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_team.sql:9
	StreamTeamSeedData(qw422016)
//line queries/seeddata/seed_team.sql:9
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_team.sql:9
}

//line queries/seeddata/seed_team.sql:9
func TeamSeedData() string {
//line queries/seeddata/seed_team.sql:9
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_team.sql:9
	WriteTeamSeedData(qb422016)
//line queries/seeddata/seed_team.sql:9
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_team.sql:9
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_team.sql:9
	return qs422016
//line queries/seeddata/seed_team.sql:9
}
