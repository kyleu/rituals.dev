// Code generated by qtc from "seed_team_member.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_team_member.sql:1
package seeddata

//line queries/seeddata/seed_team_member.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_team_member.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_team_member.sql:1
func StreamTeamMemberSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_team_member.sql:1
	qw422016.N().S(`
insert into "team_member" (
  "team_id", "user_id", "name", "picture", "role", "created", "updated"
) values (
  '10000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Test User', 'https://google.com', 'owner', now(), null
), (
  '10000000-0000-0000-0000-000000000000', '90000001-0000-0000-0000-000000000000', 'Test User 2', 'https://google.com', 'member', now(), null
), (
  '10000001-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Test User', 'https://google.com', 'owner', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_team_member.sql:11
}

//line queries/seeddata/seed_team_member.sql:11
func WriteTeamMemberSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_team_member.sql:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_team_member.sql:11
	StreamTeamMemberSeedData(qw422016)
//line queries/seeddata/seed_team_member.sql:11
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_team_member.sql:11
}

//line queries/seeddata/seed_team_member.sql:11
func TeamMemberSeedData() string {
//line queries/seeddata/seed_team_member.sql:11
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_team_member.sql:11
	WriteTeamMemberSeedData(qb422016)
//line queries/seeddata/seed_team_member.sql:11
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_team_member.sql:11
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_team_member.sql:11
	return qs422016
//line queries/seeddata/seed_team_member.sql:11
}
