
-- ----------------------------
--  Sequence structure for ideas_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."ideas_id_seq";
CREATE SEQUENCE "public"."ideas_id_seq" INCREMENT 1 START 5 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;

-- ----------------------------
--  Table structure for ideas
-- ----------------------------
DROP TABLE IF EXISTS "public"."ideas";
CREATE TABLE "public"."ideas" (
  "id" int4 NOT NULL DEFAULT nextval('ideas_id_seq'::regclass),
  "title" varchar(255) NOT NULL COLLATE "default",
  "description" text COLLATE "default"
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Records of ideas
-- ----------------------------
BEGIN;
INSERT INTO "public"."ideas" VALUES ('1', 'Bird Internet', 'Lots of tweets.');
INSERT INTO "public"."ideas" VALUES ('2', 'Go Web App', 'Should be fun.');
INSERT INTO "public"."ideas" VALUES ('3', 'Robotic Suit', 'Save the world.');
INSERT INTO "public"."ideas" VALUES ('4', 'Things For Dinner', 'Pizza and beer.');
COMMIT;


-- ----------------------------
--  Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."ideas_id_seq" RESTART 4 OWNED BY "ideas"."id";

-- ----------------------------
--  Primary key structure for table ideas
-- ----------------------------
ALTER TABLE "public"."ideas" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;