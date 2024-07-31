# gohomeapp
Simple Golang Home app demonstrate use of go template, db migration, use of Tilt for local deployment.


### Go Lib used :

Postgres : https://github.com/jackc/pgx

Session management : https://github.com/alexedwards/scs

CSRF : https://github.com/justinas/nosurf

Http Router : https://github.com/go-chi/chi


### DB Migration :
DB migration : https://gobuffalo.io/documentation/database/soda/

### Local deployment 

Tilt : https://tilt.dev/


### To Run application 

1. Local kubernetes installed. e.g Kind
2. Tilt Installed.
3. go to gohomeapp directory
4. run > tilt up
5. access browser home page: http://localhost:8080/
6. access browser about page: http://localhost:8080/about