CREATE TABLE "audit_events" (
    "id" bigserial PRIMARY KEY,
    "action" varchar NOT NULL,
    "actor" varchar NOT NULL,
    "actor_id" varchar NOT NULL,
    "catalog_service" varchar NOT NULL,
    "category_type" varchar NOT NULL,
    "timestamp" timestamptz NOT NULL,
    "server_id" varchar NOT NULL,
    "note" text NOT NULL,
    "metadata" jsonb NOT NULL DEFAULT '{}',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);