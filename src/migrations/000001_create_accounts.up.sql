CREATE TABLE IF NOT EXISTS "accounts" (
   "id" serial PRIMARY KEY,
   "name" varchar(128) DEFAULT NULL,
   "email" varchar(128) DEFAULT NULL,
   "password" varchar(128) DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS "idx_account_name" ON accounts (name);
CREATE INDEX IF NOT EXISTS "idx_account_email" ON accounts (email);