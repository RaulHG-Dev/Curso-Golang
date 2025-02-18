# Comandos para testing

1. Para ejecutar tests
```cmd
go test
```

2. Para consultar el coverage de código para testear
```cmd
go test -cover
```

3. Para consultar el código sin cover para test y almacenarlo en un archivo
```cmd
go test -coverprofile=coverage.out
```

4. Para consultar el archivo de cover generado
```cmd
go tool cover -func=coverage.out
```

```cmd
go tool cover -html=coverage.out
```

