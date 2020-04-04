# Inicializar modulo

```
 go mod init github.com/user/package
```

# Listas dependencias
```
go list -m all
```

# Varifica depencias que no se estan usando
```
go mod tidy 
```

# Ver grafico de uso de dependencias
```
go mod graph
```

# Editar un modulo
```
go mod edit -go 1.12
```

# Mostrar paquetes en json
```
go mod edit -json
```

# Descargar un paquete pero no usarlo
```
go mod download module
```