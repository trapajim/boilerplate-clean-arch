## package boilerplate code generator ##

following structure will be generated

```
./domain/{name}.go
    ->	contains a struct based on the passed arguments
        and interfaces for repository and usecases
        with basic CRUD defintions
./{name}/repository/repository.go
./{name}/usecase/usecase.go
./{name}/handler
```
    
you can change the generated code within the templates
the useCase and repostory template have following placeholders:
```
{name}: specified name lowercase
{name_u}: specified name
```
domain template has following placeholders:
```
{name}: specified name
{struct}: generated struct
```

### Usage ###

use flag -n to specifiy the name of the package and
follow it up with a list of fields for the struct FieldName:Datatype

```
go run main.go -n YourName ID:int Name:string Address:string
```
