select * from pg_extension;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE  "state_enum" AS ENUM('active', 'planned', 'done', 'failed');
CREATE TYPE "role_enum" AS ENUM('manager', 'employee');
CREATE TABLE "project" ("id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),"name" character varying(50) NOT NULL, "state" "state_enum" NOT NULL DEFAULT 'planned',"progress" character varying(3),"department" character varying(50),"owner" uuid);
CREATE TABLE "participant" ("id" uuid NOT NULL PRIMARY KEY DEFAULT  uuid_generate_v4(),"role" "role_enum" NOT NULL,"department" character varying(50),"project_id" uuid REFERENCES project (id));
