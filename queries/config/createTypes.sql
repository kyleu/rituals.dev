create type system_role as enum ('guest', 'user', 'admin');

create type member_status as enum ('owner', 'member', 'observer');

create type estimate_status as enum ('new', 'active', 'complete', 'deleted');
create type story_status as enum('pending', 'active', 'complete', 'deleted');

create type standup_status as enum ('new', 'deleted');

create type retro_status as enum ('new', 'deleted');

create type invitation_type as enum ('estimate', 'retro', 'standup');
create type invitation_status as enum ('pending', 'redeemed', 'deleted');

-- <%: func CreateTypes(w io.Writer) %>
