# ttt

ttt is the server for a two-player game of Tic-Tac-Toe, written in Go.
Players sign up as guests, one of them starts a game and shares its
five-character code with the other, and they play over a WebSocket.

## How the code is organized

Each action has its own folder under `internal`. Creating an account is in
`internal/user/core/create`, for example, and playing a move is in
`internal/game/play/create`. Every action folder has the same four parts:

- `domain` has the action's types, and the interface for the storage it
  needs.
- `api` runs the action and returns one result type for each outcome.
- `infra/data` implements that storage on MongoDB.
- `infra/web` has the route and the request and response types, and for
  playing, the WebSocket session.

Code that several actions use is in `internal/base` and in each area's
`common` folder. `pkg` has the parts that know nothing about the game, such
as JWT, IDs, nicknames, the database client and the HTTP server.

The linter checks the layering with depguard, configured in
[`.golangci.yml`](.golangci.yml):

- `domain` may not import nbio, the MongoDB driver or the JWT library.
- `api` may not import nbio or the MongoDB driver.
- `infra/data` may not import nbio or the JWT library.
- `infra/web` may not import the MongoDB driver or the JWT library.
- No package may import `reflect`.

Since the domain can't reach the server, the database or the token library,
the game rules are tested on their own.

Each game runs in its own goroutine. Moves arrive on a channel, and a ticker
ends the game when a player runs out of time. The server uses
[nbio](https://github.com/lesismal/nbio), a non-blocking I/O library, for
both HTTP and WebSockets, and accepts up to 20,000 connections.

## How to run it

To run it with MongoDB in Docker:

```sh
docker compose up --build
```

The server listens on port 8080, and MongoDB 7.0 runs in a second container.
On macOS, `make compose` runs the linter and tests first, and starts Colima
if it isn't running.

To run the binary against your own MongoDB, set the variables under
[Configuration](#configuration) and run `make run`.

## How to lint and test

`make test` runs the linter and then the tests, and `make build` does both
before building the binary. `make lint` tidies the module and runs the
linter, fixing what it can. Almost every linter is enabled. CI runs the
build, the tests and the linter on each push and pull request to `main`.

## Reference

### HTTP API

Requests and responses are JSON. Tokens go in an
`Authorization: Bearer <token>` header. `GET /user`, `DELETE /user`,
`POST /games` and `POST /games/{code}/users` need the account's access
token. `GET /gameplay` needs the access token that `POST /games` or
`POST /games/{code}/users` returns, which carries the game and player IDs.

- `POST /users` takes `platform` and `appVersion`, and returns a new guest
  account with its tokens.
- `POST /logins` takes `platform`, `appVersion` and `loginId`, and returns
  the account with new tokens.
- `POST /auths` takes `refreshToken` and returns new tokens.
- `GET /user` returns the signed-in account.
- `DELETE /user` deletes it.
- `POST /games` takes `side` and returns the new game, with its
  five-character `code`, and an access token for playing it.
- `POST /games/{code}/users` joins that game and returns it with an access
  token for playing it.
- `GET /gameplay` opens the game's WebSocket.
- `GET /` returns `ok`.

### Gameplay messages

Players send moves as `{"type": "play", "data": {"position": 4}}`.
Positions are numbered 0 to 8, one row after another, so 0, 1 and 2 are the
first row. The server sends `start`, `sync`, `move`, `end` and `error`
messages. If the player whose turn it is doesn't move within 60 seconds, the
game ends and the other player wins.

### Configuration

- `MongoDB_ConnectionString`: the MongoDB connection string
- `MongoDB_DatabasePrefix`: a prefix for the database name
- `JWT_Secret`: the signing secret
- `JWT_Issuer`: the token issuer
- `JWT_Audience`: the token audience

## License

[AGPL-3.0-only](COPYING)
