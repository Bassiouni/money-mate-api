CREATE TYPE uset_role AS ENUM ('parent', 'child');

CREATE TABLE IF NOT EXISTS users (
    id Serial Primary Key Not Null Unique,
    firstname VarChar Not Null,
    lastname VarChar Not Null,
    email VarChar(50) Not Null Unique,
    password Text Not Null,
    role_type uset_role Not Null
);
