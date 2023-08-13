CREATE TABLE messenger1.users (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger2.users (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger3.users (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger1.channels (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "name" varchar NOT NULL,
    "creator_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger1.participants (
    "id" bigserial PRIMARY KEY,
    "channel_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    "can_write" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger1.messages (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "channel_id" bigint NOT NULL,
    "sender_id" bigint NOT NULL,
    "message" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger2.messages (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "channel_id" bigint NOT NULL,
    "sender_id" bigint NOT NULL,
    "message" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger3.messages (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "channel_id" bigint NOT NULL,
    "sender_id" bigint NOT NULL,
    "message" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE messenger1.attachments (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "message_id" bigint NOT NULL,
    "file_url" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE messenger1.channels ADD FOREIGN KEY ("creator_id") REFERENCES messenger1.users ("id");

ALTER TABLE messenger1.participants ADD FOREIGN KEY ("user_id") REFERENCES messenger1.users ("id");

ALTER TABLE messenger1.participants ADD FOREIGN KEY ("channel_id") REFERENCES messenger1.channels ("id");

CREATE INDEX ON messenger1.participants ("user_id");

CREATE INDEX ON messenger1.participants ("channel_id");

CREATE INDEX ON messenger1.messages ("channel_id", "sender_id");

CREATE INDEX ON messenger2.messages ("channel_id", "sender_id");

CREATE INDEX ON messenger3.messages ("channel_id", "sender_id");

CREATE INDEX ON messenger1.attachments ("message_id");