CREATE TABLE "cards" (
    "id" bigserial PRIMARY KEY,
    "time" varchar NOT NULL,
    "date" varchar NOT NULL,
    "title" varchar NOT NULL,
    "client" varchar NOT NULL,
    "user" varchar NOT NULL
);
