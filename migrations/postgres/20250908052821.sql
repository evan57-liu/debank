-- Create "t_protocol_mapping" table
CREATE TABLE "t_protocol_mapping" (
  "id" bigint NOT NULL,
  "internal_protocol_name" character varying(128) NOT NULL,
  "internal_protocol_id" character varying(64) NOT NULL,
  "protocol_pools" json NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_t_protocol_mapping_internal_protocol_id" UNIQUE ("internal_protocol_id")
);
-- Create index "idx_t_protocol_mapping_internal_protocol_name" to table: "t_protocol_mapping"
CREATE INDEX "idx_t_protocol_mapping_internal_protocol_name" ON "t_protocol_mapping" ("internal_protocol_name");
-- Create "t_protocol_position" table
CREATE TABLE "t_protocol_position" (
  "id" bigint NOT NULL,
  "internal_protocol_name" character varying(128) NOT NULL,
  "internal_protocol_id" character varying(128) NOT NULL,
  "address" character varying(128) NOT NULL,
  "chain_id" character varying(64) NOT NULL,
  "protocol_id" character varying(64) NOT NULL,
  "pool_id" character varying(64) NOT NULL,
  "custom_id" character varying(64) NOT NULL,
  "asset_tokens" character varying(2048) NOT NULL DEFAULT '',
  "sync_time" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_t_protocol_position_custom_id" UNIQUE ("custom_id")
);
-- Create index "idx_t_protocol_position_address" to table: "t_protocol_position"
CREATE INDEX "idx_t_protocol_position_address" ON "t_protocol_position" ("address");
-- Create index "idx_t_protocol_position_chain_id" to table: "t_protocol_position"
CREATE INDEX "idx_t_protocol_position_chain_id" ON "t_protocol_position" ("chain_id");
-- Create index "idx_t_protocol_position_internal_protocol_id" to table: "t_protocol_position"
CREATE INDEX "idx_t_protocol_position_internal_protocol_id" ON "t_protocol_position" ("internal_protocol_id");
-- Create index "idx_t_protocol_position_internal_protocol_name" to table: "t_protocol_position"
CREATE INDEX "idx_t_protocol_position_internal_protocol_name" ON "t_protocol_position" ("internal_protocol_name");
-- Create index "idx_t_protocol_position_pool_id" to table: "t_protocol_position"
CREATE INDEX "idx_t_protocol_position_pool_id" ON "t_protocol_position" ("pool_id");
-- Create index "idx_t_protocol_position_protocol_id" to table: "t_protocol_position"
CREATE INDEX "idx_t_protocol_position_protocol_id" ON "t_protocol_position" ("protocol_id");
-- Create "t_user_token" table
CREATE TABLE "t_user_token" (
  "id" bigint NOT NULL,
  "address" character varying(128) NOT NULL,
  "contract_id" character varying(128) NOT NULL,
  "chain_id" character varying(64) NOT NULL,
  "symbol" character varying(64) NOT NULL,
  "decimals" bigint NOT NULL,
  "logo_url" text NULL,
  "price" numeric(40,30) NOT NULL DEFAULT 0,
  "price24_h_change" numeric(40,30) NULL DEFAULT NULL::numeric,
  "time_at" numeric NOT NULL,
  "amount" numeric(40,30) NOT NULL DEFAULT 0,
  "sync_time" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_t_user_token_address" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_address" ON "t_user_token" ("address");
-- Create index "idx_t_user_token_amount" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_amount" ON "t_user_token" ("amount");
-- Create index "idx_t_user_token_chain_id" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_chain_id" ON "t_user_token" ("chain_id");
-- Create index "idx_t_user_token_contract_id" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_contract_id" ON "t_user_token" ("contract_id");
-- Create index "idx_t_user_token_price" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_price" ON "t_user_token" ("price");
-- Create index "idx_t_user_token_price24_h_change" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_price24_h_change" ON "t_user_token" ("price24_h_change");
-- Create index "idx_t_user_token_symbol" to table: "t_user_token"
CREATE INDEX "idx_t_user_token_symbol" ON "t_user_token" ("symbol");
