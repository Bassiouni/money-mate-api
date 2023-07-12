# Money Mate API

Don't forget to add these env variables as follows in `.env` file:

```s
POSTGRES_PASSWORD=password
POSTGRES_USER=username
POSTGRES_DB=dbname
```

Start the database usign `docker compose up postgresql -d`

Then run DB migrations using `migrations_up.sh` script.

Start the server using `docker compose up --build -d`.

The server will run on port `8080` and the postgres db on port `5432` binded to the host.

To stop everything run `docker compose down`

## API Routes

The available routes are:
- Group "/users"
    - `POST` **"/new"**: Creates new user
    - `GET` **"/:id"**: Get user by id
    - `GET` **"/"**: Get all users in the db

### User Object Scheme

The expected user object scheme is:
```
users {
    id Primary Key (Optional),
    firstname string,
    lastname string,
    email string (unique),
    password string,
    role_type enum('parent', 'child')
}
```
