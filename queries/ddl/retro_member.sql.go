// Code generated by qtc from "retro_member.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/retro_member.sql:2
package ddl

//line queries/ddl/retro_member.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/retro_member.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/retro_member.sql:2
func StreamRetroMemberDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/retro_member.sql:2
	qw422016.N().S(`
drop table if exists "retro_member";
-- `)
//line queries/ddl/retro_member.sql:4
}

//line queries/ddl/retro_member.sql:4
func WriteRetroMemberDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/retro_member.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/retro_member.sql:4
	StreamRetroMemberDrop(qw422016)
//line queries/ddl/retro_member.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/retro_member.sql:4
}

//line queries/ddl/retro_member.sql:4
func RetroMemberDrop() string {
//line queries/ddl/retro_member.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/retro_member.sql:4
	WriteRetroMemberDrop(qb422016)
//line queries/ddl/retro_member.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/retro_member.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/retro_member.sql:4
	return qs422016
//line queries/ddl/retro_member.sql:4
}

// --

//line queries/ddl/retro_member.sql:6
func StreamRetroMemberCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/retro_member.sql:6
	qw422016.N().S(`
create table if not exists "retro_member" (
  "retro_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("retro_id") references "retro" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("retro_id", "user_id")
);

create index if not exists "retro_member__retro_id_idx" on "retro_member" ("retro_id");

create index if not exists "retro_member__user_id_idx" on "retro_member" ("user_id");
-- `)
//line queries/ddl/retro_member.sql:23
}

//line queries/ddl/retro_member.sql:23
func WriteRetroMemberCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/retro_member.sql:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/retro_member.sql:23
	StreamRetroMemberCreate(qw422016)
//line queries/ddl/retro_member.sql:23
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/retro_member.sql:23
}

//line queries/ddl/retro_member.sql:23
func RetroMemberCreate() string {
//line queries/ddl/retro_member.sql:23
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/retro_member.sql:23
	WriteRetroMemberCreate(qb422016)
//line queries/ddl/retro_member.sql:23
	qs422016 := string(qb422016.B)
//line queries/ddl/retro_member.sql:23
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/retro_member.sql:23
	return qs422016
//line queries/ddl/retro_member.sql:23
}
