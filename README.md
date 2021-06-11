## Log

- `go mod init github.com/d-led/go-gin-heroku`
- [main.go](main.go) with gin & `go get`
- `go run .` &rarr; http://localhost:8080/ &rarr; commit
- [main_test.go](main_test.go) sanity check, `go mod tidy` to fetch the assertion library, commit
- refactor to testability, test, commit
- https://signup.heroku.com/login
- https://dashboard.heroku.com/new-app
- enable Github deployments, "wait for CI to finish", trigger deployment
- done: https://go-gin-heroku.herokuapp.com/
- added `/version` &rarr; empty in production
- installed Heroku CLI
- `heroku buildpacks:add --index 1 https://github.com/40thieves/heroku-buildpack-commit-hash.git -a go-gin-heroku` (login via the browser)
- `heroku git:remote -a go-gin-heroku`
- deploy via `git push heroku master`
- https://go-gin-heroku.herokuapp.com/version done
