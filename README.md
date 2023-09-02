## Library for nksdash auth microserivce

Usage:

1. Specify in .env
```
AUTH_SERVICE_HOST
AUTH_SERVICE_PORT
```

2. 
```
factory := NewHandlerFactory()

handler := factory.createCheckTokenQueryHandler()

data, error := handler.Handle("token")

if error != nill {
//
}

userId := data.UserId
```
___
For now only "Check Token" functionality implemented.

