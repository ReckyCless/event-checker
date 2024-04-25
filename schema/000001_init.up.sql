CREATE TABLE "events" (
  "id" integer PRIMARY KEY generated always as identity,
  "title" text NOT NULL,
  "description" text NOT NULL,
  "location" text NOT NULL,
  "image_path" text,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp,
  "creator_id" integer NOT NULL,
  "type_id" integer NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "events_types" (
  "id" integer PRIMARY KEY generated always as identity,
  "name" text NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "events_tags" (
  "event_id" integer,
  "tag_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("event_id", "tag_id")
);

CREATE TABLE "tags" (
  "id" integer PRIMARY KEY generated always as identity,
  "name" text NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY generated always as identity,
  "password_hash" text NOT NULL,
  "name" text NOT NULL,
  "surname" text NOT NULL,
  "patronymic" text,
  "birth_date" date NOT NULL,
  "sex" boolean NOT NULL,
  "phone" text,
  "email" text,
  "image_path" text,
  "role_id" integer NOT NULL,
  "organisator_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "events_visitors" (
  "id" integer PRIMARY KEY generated always as identity,
  "user_id" integer,
  "event_id" integer NOT NULL,
  "name" text,
  "surname" text,
  "patronymic" text,
  "birth_date" date,
  "sex" boolean NOT NULL,
  "phone" text,
  "email" text,
  "is_visited" boolean NOT NULL DEFAULT false,
  "created_at" timestamp,
  "updated_at" timestamp,
  UNIQUE (user_id, event_id)
);

CREATE TABLE "organisators" (
  "id" integer PRIMARY KEY generated always as identity,
  "title" text NOT NULL,
  "logo_path" text,
  "site_url" text,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "roles" (
  "id" integer PRIMARY KEY generated always as identity,
  "name" text NOT NULL
);

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "events" ADD FOREIGN KEY ("type_id") REFERENCES "events_types" ("id") ON DELETE NULL ON UPDATE CASCADE;

ALTER TABLE "events_tags" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "events_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE NULL ON UPDATE CASCADE;

ALTER TABLE "users" ADD FOREIGN KEY ("organisator_id") REFERENCES "organisators" ("id") ON DELETE NULL ON UPDATE CASCADE;

ALTER TABLE "events_visitors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NULL ON UPDATE CASCADE;

ALTER TABLE "events_visitors" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE NULL ON UPDATE CASCADE;

INSERT INTO "roles" ("name") VALUES('Администратор');
INSERT INTO "roles" ("name") VALUES('Организатор');
INSERT INTO "roles" ("name") VALUES('Пользователь');

INSERT INTO "types" ("name") VALUES('Меропритие');
INSERT INTO "types" ("name") VALUES('Конкурс');
INSERT INTO "types" ("name") VALUES('Соревнование');
INSERT INTO "types" ("name") VALUES('Лекция');