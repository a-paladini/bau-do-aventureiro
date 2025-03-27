CREATE TABLE "weapons" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL,
  "category" VARCHAR NOT NULL,
  "price" FLOAT NOT NULL,
  "slot" FLOAT NOT NULL,
  "origin" VARCHAR NOT NULL,
  "damage" VARCHAR NOT NULL,
  "critical" VARCHAR NOT NULL,
  "range" VARCHAR NOT NULL,
  "type_damage" VARCHAR NOT NULL,
  "property" VARCHAR NOT NULL,
  "proficiency" VARCHAR NOT NULL,
  "special" VARCHAR NOT NULL
);

CREATE TABLE "armours" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL,
  "category" VARCHAR NOT NULL,
  "price" FLOAT NOT NULL,
  "slot" FLOAT NOT NULL,
  "origin" VARCHAR NOT NULL,
  "ca_bonus" INT NOT NULL,
  "penality" INT NOT NULL
);

CREATE TABLE "items" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL,
  "category" VARCHAR NOT NULL,
  "price" FLOAT NOT NULL,
  "slot" FLOAT NOT NULL,
  "origin" VARCHAR NOT NULL
);

CREATE INDEX ON "weapons" ("category");

CREATE INDEX ON "weapons" ("type_damage");

CREATE INDEX ON "armours" ("category");

CREATE INDEX ON "items" ("category");