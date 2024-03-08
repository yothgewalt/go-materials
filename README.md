# Go: The Fundamentals and Design Patterns

### Visual Studio Code (Extensions)
- Go: https://marketplace.visualstudio.com/items?itemName=golang.Go

- Error Lens: https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens

- Go Group Imports: https://marketplace.visualstudio.com/items?itemName=aleksandra.go-group-imports

## How to use Air (Hot Reload)
1. Install `Air (Hot Reload)` with Go Installer
```cmd
$ go install github.com/cosmtrek/air@latest
```
2. Initialize an Air file.
```cmd
$ air init
```

## How to configure Air (Hot Reload)
Configuration on <b>Mac OS X</b> </br >

1. Add `;export $(grep -v '^#' .env | xargs);` into <b>.air.toml</b> at bin = "..."
```toml
bin = ";export $(grep -v '^#' .env | xargs); ./build/main"
```

Configuration on <b>Windows</b> </br >

1. Create a `.bat` file and insert the below code. (And Run it!)
```bat
@FOR /F "tokens=*" %%a IN ('FINDSTR /V /B "#" .env') DO SET "%%a"
air
```


## Go Packages (https://pkg.go.dev)
- Go Fiber: https://docs.gofiber.io

- GORM: https://gorm.io/index.html

- ZeroLog: https://github.com/rs/zerolog

## Go Guide Styles
- Uber Styles: https://github.com/uber-go/guide

- Google Styles: https://google.github.io/styleguide/go/

