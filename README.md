# URL Shortener

It parses YAML using [gopkg.in/yaml.v2](https://godoc.org/gopkg.in/yaml.v2) package. It looks at the path of any incoming 
web request and determine if it should redirect the user to a new page, much like URL shortener  would.

- [x] Update the main/main.go source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string.
- [x] Build a JSONHandler that serves the same purpose, but reads from JSON data.
- [ ] Build a Handler that doesn't read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.

[Source](https://gophercises.com/exercises/urlshort)
