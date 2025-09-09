-- Create sequence for serial column "id"
CREATE SEQUENCE IF NOT EXISTS "t_protocol_mapping_id_seq" OWNED BY "t_protocol_mapping"."id";
-- Modify "t_protocol_mapping" table
ALTER TABLE "t_protocol_mapping" ALTER COLUMN "id" SET DEFAULT nextval('"t_protocol_mapping_id_seq"');
-- Create sequence for serial column "id"
CREATE SEQUENCE IF NOT EXISTS "t_protocol_position_id_seq" OWNED BY "t_protocol_position"."id";
-- Modify "t_protocol_position" table
ALTER TABLE "t_protocol_position" ALTER COLUMN "id" SET DEFAULT nextval('"t_protocol_position_id_seq"');
-- Create sequence for serial column "id"
CREATE SEQUENCE IF NOT EXISTS "t_user_token_id_seq" OWNED BY "t_user_token"."id";
-- Modify "t_user_token" table
ALTER TABLE "t_user_token" ALTER COLUMN "id" SET DEFAULT nextval('"t_user_token_id_seq"');
