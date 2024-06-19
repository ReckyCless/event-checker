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

ALTER TABLE "events" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "events" ADD FOREIGN KEY ("type_id") REFERENCES "events_types" ("id") ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE "events_tags" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "events_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE "users" ADD FOREIGN KEY ("organisator_id") REFERENCES "organisators" ("id") ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE "events_visitors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE "events_visitors" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE SET NULL ON UPDATE CASCADE;

INSERT INTO "roles" ("name") VALUES('Администратор');
INSERT INTO "roles" ("name") VALUES('Организатор');
INSERT INTO "roles" ("name") VALUES('Пользователь');

INSERT INTO "events_types" ("name") VALUES('Мероприятие');
INSERT INTO "events_types" ("name") VALUES('Конкурс');
INSERT INTO "events_types" ("name") VALUES('Соревнование');
INSERT INTO "events_types" ("name") VALUES('Лекция');

INSERT INTO "users" ("name", "surname", "patronymic", "birth_date", "sex", "email", "password_hash", "role_id") VALUES('Хоро', 'Мудрая', 'Дмитриевна', '2012-12-12T00:00:00.00Z', 'false', 'horo@wise.com', '45384a33666e4d7753733033487a534b464437586e435647736e366153326b3171594e8d7fa6428e130ac023c577c37b4dc9325f271692', 1);
INSERT INTO "users" ("name", "surname", "patronymic", "birth_date", "sex", "email", "password_hash", "role_id") VALUES('Валера', 'Тоухович', 'Васильевич', '2009-02-15T00:00:00.00Z', 'true', 'roge@cnad.org', '45384a33666e4d7753733033487a534b464437586e435647736e366153326b3171594eea825c1a6416c74989d8b8dec9565df6621e91e2', 1);

INSERT INTO "events" ("title", "description", "location", "start_time", "end_time", "creator_id", "type_id") VALUES('Олимпида чтецов', 'В честь дня поэзии будет проведён конкурс на лучшее прочтение', 'Ул. Пушкина, д.12, к.4', '2024-02-15T15:30:00.00Z', '2024-02-15T17:30:00.00Z', 1, 2);
INSERT INTO "events" ("title", "description", "location", "start_time", "end_time", "creator_id", "type_id") VALUES('Выставка искуств', 'Для продвижения новых взглядов на искуство была организована выставка', 'Пр. Пикасо, д.14, к.1', '2024-03-16T10:00:00.00Z', '2024-03-16T21:30:00.00Z', 1, 4);
INSERT INTO "events" ("title", "description", "location", "start_time", "end_time", "creator_id", "type_id") VALUES('Мастер-класс по живописи', 'Опытные художники поделятся секретами своего мастерства', 'Ул. Левитана, д.8', '2024-04-20T12:00:00.00Z', '2024-04-20T14:00:00.00Z', 2, 3);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Лекция по истории искусств', 'Известный историк расскажет о развитии искусств в разных эпохах', 'Проспект Академический, д.2', '2024-02-22T16:00:00.00Z', 1, 4);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Концерт классической музыки', 'Оркестр исполнит произведения Моцарта и Бетховена', 'Ул. Лунная, д.5', '2024-05-25T19:00:00.00Z', 1, 3);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Театральная постановка', 'Премьера новой пьесы современного драматурга', 'Театральная площадь, д.1', '2024-05-28T18:30:00.00Z', 1, 1);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Фестиваль уличной еды', 'Дегустация блюд от лучших уличных поваров города', 'Центральный парк', '2024-05-01T11:00:00.00Z', 1, 2);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Кинопоказ под открытым небом', 'Просмотр классических фильмов под звёздами', 'Парк Горького, д.3', '2024-08-03T21:00:00.00Z', 2, 1);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Йога на рассвете', 'Утренние занятия йогой на свежем воздухе', 'Набережная реки, д.7', '2024-08-05T05:30:00.00Z', 2, 1);
INSERT INTO "events" ("title", "description", "location", "start_time", "creator_id", "type_id") VALUES('Благотворительный марафон', 'Сбор средств для местного приюта для животных', 'Главная площадь, д.9', '2024-08-10T08:00:00.00Z', 1, 4);
