# Movie CRUD using structs and slices
This example is taken from [Learn Go Programming by Building 11 projects](https://www.youtube.com/watch?v=jFfo23yIWac&t=142s).

## Overview
- Create, Read, Update or Delete movies from the server. Nothing fancy
- Postman is employed to test the server. The Postman workspace is not included
- Update mode is dodgy, as this is just for illustration and for me to learn some Go

## Notes
- Adding the Gorilla Mux library resulted in messing up the [go.sum and go.mod](https://golangbyexample.com/go-mod-sum-module/) files
- I messed with that which resulted in a [broken import](https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#BrokenImport)
- This is a result of the way the sample projects are structured. I would prefer them to be separate concerns, each being a standalone project, that doesn't bork each other's go.sum and go.mod files
- The example runs well enough after it is built
- I also added a .gitignore file to prevent exe files from landing up in my repo