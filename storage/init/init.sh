psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d "$POSTGRES_DB"  <<-EOSQL
CREATE TABLE IF NOT EXISTS "$POSTGRES_DB_TBL_USER" (
    user_id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    email varchar(40) NOT NULL CHECK (email <> ''),
    pass_hash varchar NOT NULL CHECK (pass_hash <> ''),
    refresh_hash varchar,
    expires_at bigint,
    UNIQUE(email)
);
CREATE TABLE IF NOT EXISTS "$POSTGRES_DB_TBL_CATEGORY" (
    category_id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(40) NOT NULL CHECK (name <> ''),
    code varchar(40) NOT NULL CHECK (code <> ''),
    description varchar,
    UNIQUE(code)
);
EOSQL