CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false
);
