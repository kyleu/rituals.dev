-- User
create table if not exists "system_user" (
  "id" uuid not null primary key,
  "name" varchar(2048) not null,
  "role" system_role not null,
  "theme" varchar(32) not null,
  "nav_color" varchar(32) not null,
  "link_color" varchar(32) not null,
  "locale" varchar(32) not null,
  "created" timestamp not null default now()
);

-- Estimate
create table if not exists "estimate" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "password" varchar(2048) not null default '',
  "title" varchar(2048) not null,
  "owner" uuid references "system_user"("id"),
  "status" estimate_status not null,
  "choices" varchar(2048)[] not null,
  "options" json not null,
  "created" timestamp not null default now()
);

create table if not exists "estimate_member" (
  "estimate_id" uuid not null references "estimate"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" varchar(64) not null,
  "created" timestamp not null default now(),
  primary key ("estimate_id", "user_id")
);

create table "poll" (
  "id" uuid not null primary key,
  "estimate_id" uuid not null references "estimate"("id"),
  "idx" int not null default 0,
  "author_id" uuid not null references "system_user"("id"),
  "title" varchar(2048),
  "status" poll_status not null,
  "final_vote" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table "vote" (
  "poll_id" uuid not null references "poll"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "choice" varchar(256) not null,
  "updated" timestamp not null default now(),
  "created" timestamp not null default now(),
  primary key ("poll_id", "user_id")
);

-- Standup
create table if not exists "standup" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "owner" uuid references "system_user"("id"),
  "status" standup_status not null,
  "created" timestamp not null default now()
);

create table if not exists "standup_member" (
  "standup_id" uuid not null references "standup"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" varchar(64) not null,
  "created" timestamp not null default now(),
  primary key ("standup_id", "user_id")
);

-- Retro
create table if not exists "retro" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "owner" uuid references "system_user"("id"),
  "status" retro_status not null,
  "created" timestamp not null default now()
);

create table if not exists "retro_member" (
  "retro_id" uuid not null references "retro"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" varchar(64) not null,
  "created" timestamp not null default now(),
  primary key ("retro_id", "user_id")
);

-- Invite
create table if not exists "invitation" (
  "key" varchar(128) not null primary key,
  "k" invitation_type not null,
  "v" uuid not null,
  "src" uuid references "system_user"("id"),
  "tgt" uuid references "system_user"("id"),
  "note" text,
  "status" invitation_status not null,
  "redeemed" timestamp,
  "created" timestamp not null default now()
);

-- <%: func CreateTables(w io.Writer) %>
