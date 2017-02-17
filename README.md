# Go Init
An initial structure for building Go web apps.

## Installation
- Clone this repository in a location from your `GOPATH`
- Replace all occurrences of `go-init` with the name of the project

> If you wish to change the port that the app will listen to, go to `main.go`

## Usage
### Routing
- Install the [routing package](https://github.com/husobee/vestigo)
```
go get github.com/husobee/vestigo
```
- Import the `features` package
```go
import f "github.com/[vendor]/[project]/features"
```
- Routes can be defined in `routes/routes.go` in the `Register` method as follows:
```go
r.Get('/path', Serve(f.DoSomething))
```

### Parsing & Validation

- Install the [validation package](https://github.com/go-playground/validator)
```
go get gopkg.in/go-playground/validator.v9
```
- Insall the [parsing package](http://www.gorillatoolkit.org/pkg/schema)
```
go get github.com/gorilla/schema
```
- Use the following snippet to parse and validate given input as a struct (see [here](#errors) for errors)
```go
type ListingInput struct {
	Limit  uint `schema:"limit" validate:"min=0,max=100"`
	Offset uint `schema:"offset" validate:"min=0"`
}

func ParseListingInput(r *http.Request) *ListingInput {
	input := new(ListingInput)
	schema.NewDecoder().Decode(input, r.URL.Query())

	return input
}

func Validate(input *ListingInput) {
	if err := validator.New().Struct(input); err != nil {
		panic(throw.InvalidInputError(err.Error(), 1000, 400))
	}
}
```

### Responses

- Import the `response` package
```go
import "github.com/[vendor]/[project]/response
```
#### Success
```go
response.JSON(w, data)
```
#### Failure
```go
response.JSONError(w, status, errorCode, errorMessage)
```

### Error Handling
- Define errors in the `throw` package at `throw/errors.go` as a struct
```go
type SomethingWrong struct {
	*HTTPError
}
```
- For each error, better have a factory method to help make a new instance of the error, i.e.
```go
func SomethingWrongError(m string, c, s uint16) SomethingWrong {
	return InvalidInput{&HTTPError{Message: m, Code: c, Status: s}r}
}
```
