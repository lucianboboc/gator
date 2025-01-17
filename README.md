# How to run the app

Requirements:
- Golang
- Postgres
- `~/.gatorconfig.json` config file

`gatorconfig.json` example:
```json
{
  "db_url": "pg_connection_string",
  "current_user_name": "username"
}
```

Use `go install` to install the gator CLI.
Available commands:
```
- register
- login
- reset
- addfeed
- agg
- browse
- feeds
- follow
- unfollow
- following