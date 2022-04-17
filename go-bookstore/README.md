# Golang bookstore that uses MySQL
This example is taken from [Learn Go Programming by Building 11 projects](https://www.youtube.com/watch?v=jFfo23yIWac&t=142s).

There's a similar [example that uses SQLite from Pluralsight](https://pluralsight.hashnode.dev/how-to-build-a-web-app-with-go-and-sqlite).

## Notes
- This project structure is less confusing than the other projects in the series
- Seems like there are other dialects for the Gorm package, if MySQL is not your first choice
- Golang has absolute paths. Note the paths of the controllers in particular
- Imports are fun due to the absolute paths. At least [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports?utm_source=godoc) provides a standard
- I left the exe files out of the repo
- I didn't add the database username or password as this is just a sample project, and I don't want to share sensitive information in repos. There are some recommended best practices to [clean configuration management in Golang](https://kaznacheev.me/posts/en/clean-configuration-management-golang/)