# go-agency

The meaning of this project is learn GO in phases. On each phase, the complexity and libraries that are being used on 
this project are changes. I started with native queries and `go-chi` library to create my APIs (until release 0.1.1).

On version 0.2.0, `go-chi` was replaced by fiber, and native queries are also replaced by `GORM`. Moreover, the database
complexity was increased, in order to bring a better comprehension about how GORM manage and query data.

Version 0.2.2 open api documentation was added. Using `SWAGGO` project (https://github.com/swaggo/swag), more 
specifically the experimental branch (v2), the open api documentation is being generated (there are more details about 
this below). 

## GORM

Currently, this project is using GORM, and the logger is configured to write all queries for the purpose of understand
how queries are built.

## Open API V3

Another purpose of this project is perform a small POC using libraries that can be used to generate the API 
documentation from code. So, here we are using `swaggo` project, not the main one but beta branch: 
https://github.com/swaggo/swag/releases/tag/v2.0.0-rc3

### Using swaggo v2:

First of all, we need to install the generator:

```
$ go install github.com/swaggo/swag/v2/cmd/swag@latest
## OR using the specific version:
$ go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc3
```

After this installation we will have access to `swag` command:
```
$ swag
Swag version:  v2.0.0
NAME:
COMMANDS:
   init, i  Create docs.go
   fmt, f   format swag comments
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

To be able to generate open api file, check the documentation: 
https://github.com/swaggo/swag/tree/v2#declarative-comments-format

After define all necessary declarations on our project, we can make the documentation files using the following command:
```
$ swag init --v3.1
```

Once we are using beta version, and expecting OA3 file template, the tag `--v3.1` is required, otherwise the swagger v2
will be generated.

If there are no missing tags, a folder named docs will be created.

### Rendering swagger on browse - troubleshooting 

The library `fiber-swagger` uses by default the last version from swaggo v1.8 (not the beta one - v2). In the other 
hand, when we perform the `swag init --v3.1` (using v3.1 parameter), it will generate `docs.go` file declaring
all instances using `swaggo` version 2, and once `fiber-swagger` is using a different version than `swaggo`, the render 
instances will not work properly, so we must change the version from `docs.go` file. So, how we can do that?

After generate all `docs` files using `swag init --v3.1` command, we must go to docs and change those lines:

1. On import, change the import library:
    ```go
    // FROM this:
    import "github.com/swaggo/swag/v2"
    
    // TO this:
    import "github.com/swaggo/swag"
    ```

2. Now, change the openapi version - because the previous version only support swagger 2.0 or open api 3.0.x files):
    ```go
    //FROM this
    const docTemplate = `{
        ...
        "openapi": "3.1.0",
        ...
    }`
    
    // TO this
    const docTemplate = `{
        ...
        "openapi": "3.0.2",
        ...
    }`
    ```

3. Now we must change the generated swagger files (json and yml). In those files, replace the `openapi` attribute value 
from `3.1.0` to `3.0.2`.

After these changes, the swagger file will be rendered successfully.

## Links:

* Serve APi documentation: https://github.com/swaggo/fiber-swagger#canonical-example
* Declarative comments format: https://github.com/swaggo/swag/tree/v2#declarative-comments-format
