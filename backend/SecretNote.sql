CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "type_of_login" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "post_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

ALTER TABLE "posts" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("email");

CREATE INDEX ON "posts" ("title");
