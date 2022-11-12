# Setting up the project

```bash
# Enable dependency tracking by initialize `go.mod` file
go mod init DOMAIN/MODULE_NAME
```

# Setting up local dependency

Modify a `go.mod` file to replace the given module to use local files

```go.mod
replace DOMAIN/MODULES/MODULE_1 => ./LOCAL/PATH/TO/MODULE_1
replace DOMAIN/MODULES/MODULE_2 => ./LOCAL/PATH/TO/MODULE_2
...
```

# Update modules

```bash
# Add imported modules as a requirement
go mod tidy
```

# Running the application

```bash
go run PATH|DIR
```

# Build the application

```bash
go build PATH|DIR
```

# Test the modules

```bash
go test -v 
```