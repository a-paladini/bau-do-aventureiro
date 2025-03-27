CREATE TABLE "armas" (
  "id" SERIAL PRIMARY KEY,
  "nome" VARCHAR NOT NULL,
  "descricao" VARCHAR NOT NULL,
  "categoria" VARCHAR NOT NULL,
  "preco" FLOAT NOT NULL,
  "espaco" FLOAT NOT NULL,
  "origem" VARCHAR NOT NULL,
  "dano" VARCHAR NOT NULL,
  "critico" VARCHAR,
  "alcance" VARCHAR,
  "tipo_dano" VARCHAR,
  "empunhadura" VARCHAR,
  "proficiencia" VARCHAR,
  "habilidades" VARCHAR
);

CREATE TABLE "armaduras" (
  "id" SERIAL PRIMARY KEY,
  "nome" VARCHAR NOT NULL,
  "descricao" VARCHAR NOT NULL,
  "categoria" VARCHAR NOT NULL,
  "preco" FLOAT NOT NULL,
  "espaco" FLOAT NOT NULL,
  "origem" VARCHAR NOT NULL,
  "bonus_na_defesa" INT,
  "penalidade_de_armadura" INT
);

CREATE TABLE "itens" (
  "id" SERIAL PRIMARY KEY,
  "nome" VARCHAR NOT NULL,
  "descricao" VARCHAR NOT NULL,
  "categoria" VARCHAR NOT NULL,
  "preco" FLOAT NOT NULL,
  "espaco" FLOAT DEFAULT 0,
  "origem" VARCHAR NOT NULL
);

CREATE INDEX ON "armas" ("categoria");

CREATE INDEX ON "armas" ("tipo_dano");

CREATE INDEX ON "armaduras" ("categoria");

CREATE INDEX ON "itens" ("categoria");

COMMENT ON COLUMN "armas"."preco" IS 'valor em moedas de prata';

COMMENT ON COLUMN "armaduras"."preco" IS 'valor em moedas de prata';

COMMENT ON COLUMN "itens"."preco" IS 'valor em moedas de prata';
