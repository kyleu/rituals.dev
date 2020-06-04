-- user
create table if not exists "system_user" (
  "id" uuid not null primary key,
  "name" varchar(2048) not null,
  "role" system_role not null,
  "theme" varchar(32) not null,
  "nav_color" varchar(32) not null,
  "link_color" varchar(32) not null,
  "picture" text not null,
  "locale" varchar(32) not null,
  "created" timestamp not null default now()
);

create table if not exists "auth" (
  "id" uuid not null primary key,
  "user_id" uuid not null references "system_user"("id"),
  "provider" auth_provider not null,
  "provider_id" text not null,
  "user_list_id" varchar(512) not null,
  "user_list_name" varchar(2048) not null,
  "access_token" text not null,
  "expires" timestamp,
  "name" varchar(2048) not null,
  "email" varchar(2048) not null,
  "picture" text not null,
  "created" timestamp not null default now()
);

create index if not exists idx_auth_provider_provider_id on auth(provider, provider_id);

-- team
create table if not exists "team" (
  "id" uuid not null primary key,
  "slug" varchar(1024) not null unique,
  "title" varchar(2048) not null,
  "owner" uuid references "system_user"("id"),
  "created" timestamp not null default now()
);

create unique index if not exists idx_team_slug on team(slug);

create table if not exists "team_member" (
  "team_id" uuid not null references "team"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("team_id", "user_id")
);

create table if not exists "team_history" (
  "slug" varchar(1024) not null primary key,
  "model_id" uuid not null references "team"("id"),
  "model_name" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "team_permission" (
  "team_id" uuid not null references "team"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (team_id, k, v)
);

create index if not exists idx_team_permission_team_id on team_permission(team_id);

-- sprint
create table if not exists "sprint" (
  "id" uuid not null primary key,
  "slug" varchar(1024) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "owner" uuid references "system_user"("id"),
  "start_date" date,
  "end_date" date,
  "created" timestamp not null default now()
);

create unique index if not exists idx_sprint_slug on sprint(slug);

create table if not exists "sprint_member" (
  "sprint_id" uuid not null references "sprint"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("sprint_id", "user_id")
);

create table if not exists "sprint_history" (
  "slug" varchar(1024) not null primary key,
  "model_id" uuid not null references "sprint"("id"),
  "model_name" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "sprint_permission" (
  "sprint_id" uuid not null references "sprint"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (sprint_id, k, v)
);

create index if not exists idx_sprint_permission_sprint_id on sprint_permission(sprint_id);

-- estimate
create table if not exists "estimate" (
  "id" uuid not null primary key,
  "slug" varchar(1024) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "sprint_id" uuid references "sprint"("id"),
  "owner" uuid references "system_user"("id"),
  "status" estimate_status not null,
  "choices" varchar(2048)[] not null,
  "created" timestamp not null default now()
);

create unique index if not exists idx_estimate_slug on estimate(slug);

create table if not exists "estimate_member" (
  "estimate_id" uuid not null references "estimate"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("estimate_id", "user_id")
);

create table if not exists "estimate_history" (
  "slug" varchar(1024) not null primary key,
  "model_id" uuid not null references "estimate"("id"),
  "model_name" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "estimate_permission" (
  "estimate_id" uuid not null references "estimate"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (estimate_id, k, v)
);

create index if not exists idx_estimate_permission_estimate_id on estimate_permission(estimate_id);

create table if not exists "story" (
  "id" uuid not null primary key,
  "estimate_id" uuid not null references "estimate"("id"),
  "idx" int not null default 0,
  "user_id" uuid not null references "system_user"("id"),
  "title" varchar(2048) not null,
  "status" story_status not null default 'pending',
  "final_vote" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "vote" (
  "story_id" uuid not null references "story"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "choice" varchar(256) not null,
  "updated" timestamp not null default now(),
  "created" timestamp not null default now(),
  primary key ("story_id", "user_id")
);

-- standup
create table if not exists "standup" (
  "id" uuid not null primary key,
  "slug" varchar(1024) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "sprint_id" uuid references "sprint"("id"),
  "owner" uuid references "system_user"("id"),
  "status" standup_status not null,
  "created" timestamp not null default now()
);

create unique index if not exists idx_standup_slug on standup(slug);

create table if not exists "standup_member" (
  "standup_id" uuid not null references "standup"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("standup_id", "user_id")
);

create table if not exists "standup_history" (
  "slug" varchar(1024) not null primary key,
  "model_id" uuid not null references "standup"("id"),
  "model_name" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "standup_permission" (
  "standup_id" uuid not null references "standup"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (standup_id, k, v)
);

create index if not exists idx_standup_permission_standup_id on standup_permission(standup_id);

create table "report" (
  "id" uuid not null primary key,
  "standup_id" uuid not null references "standup"("id"),
  "d" date not null default now()::date,
  "user_id" uuid not null references "system_user"("id"),
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now()
);

-- retro
create table if not exists "retro" (
  "id" uuid not null primary key,
  "slug" varchar(1024) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "sprint_id" uuid references "sprint"("id"),
  "owner" uuid references "system_user"("id"),
  "status" retro_status not null,
  "categories" varchar(2048)[] not null,
  "created" timestamp not null default now()
);

create unique index if not exists idx_retro_slug on retro(slug);

create table if not exists "retro_member" (
  "retro_id" uuid not null references "retro"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("retro_id", "user_id")
);

create table if not exists "retro_history" (
  "slug" varchar(1024) not null primary key,
  "model_id" uuid not null references "retro"("id"),
  "model_name" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "retro_permission" (
  "retro_id" uuid not null references "retro"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (retro_id, k, v)
);

create index if not exists idx_retro_permission_retro_id on retro_permission(retro_id);

create table if not exists "feedback" (
  "id" uuid not null primary key,
  "retro_id" uuid not null references "retro"("id"),
  "idx" int not null default 0,
  "user_id" uuid not null references "system_user"("id"),
  "category" varchar(2048) not null,
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now()
);

-- action
create table if not exists "action" (
  "id" uuid not null primary key,
  "svc" model_service not null,
  "model_id" uuid not null,
  "user_id" uuid references "system_user"("id"),
  "act" varchar(64) not null,
  "content" json,
  "note" text,
  "created" timestamp not null default now()
);

-- comment
create table if not exists "comment" (
  "id" uuid not null primary key,
  "svc" model_service not null,
  "model_id" uuid not null,
  "target_type" varchar(128) not null,
  "target_id" uuid,
  "user_id" uuid references "system_user"("id"),
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now()
);

create index if not exists idx_comment_svc_model_id on comment(svc, model_id);

-- emails
create table if not exists "email" (
  "id" varchar(1024) not null primary key,
  "recipients" varchar(1024)[] not null,
  "subject" text not null,
  "data" json not null,
  "plain" text not null,
  "html" text not null,
  "user_id" uuid references "system_user"("id"),
  "status" varchar(128) not null,
  "created" timestamp not null default now()
);

-- migrations
create table if not exists "migration" (
  "idx" int not null primary key,
  "title" varchar(1024) not null,
  "src" text not null,
  "created" timestamp not null default now()
);

-- <%: func CreateTables(w io.Writer) %>
