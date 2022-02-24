CREATE TABLE IF NOT EXISTS "items" (
   "id" serial PRIMARY KEY,
   "name" varchar(128) DEFAULT NULL,
   "price" int NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_item_name" ON items (name);