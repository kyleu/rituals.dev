// Code generated by qtc from "seed_estimate.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_estimate.sql:1
package seeddata

//line queries/seeddata/seed_estimate.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_estimate.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_estimate.sql:1
func StreamEstimateSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_estimate.sql:1
	qw422016.N().S(`
insert into "estimate" (
  "id", "slug", "title", "icon", "status", "team_id", "sprint_id", "owner", "choices", "created", "updated"
) values (
  '30000000-0000-0000-0000-000000000000', 'estimate-1', 'Estimate 1', 'star', 'active', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '["0","1","2","3","5","8","13","100"]', now(), null
), (
  '30000001-0000-0000-0000-000000000000', 'estimate-2', 'Estimate 2', 'dot-circle', 'active', '10000001-0000-0000-0000-000000000000', '20000001-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '["0","1","2","3","5","8","13","100"]', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_estimate.sql:9
}

//line queries/seeddata/seed_estimate.sql:9
func WriteEstimateSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_estimate.sql:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_estimate.sql:9
	StreamEstimateSeedData(qw422016)
//line queries/seeddata/seed_estimate.sql:9
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_estimate.sql:9
}

//line queries/seeddata/seed_estimate.sql:9
func EstimateSeedData() string {
//line queries/seeddata/seed_estimate.sql:9
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_estimate.sql:9
	WriteEstimateSeedData(qb422016)
//line queries/seeddata/seed_estimate.sql:9
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_estimate.sql:9
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_estimate.sql:9
	return qs422016
//line queries/seeddata/seed_estimate.sql:9
}
