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

# Verificar que todos los modulos esten correctos
```
go mod verify
```

# Reemplazar un modulo por una variacion en un path especifico
```
go mod edit -replace <module>=<path>
```
# No se debe modificar los paquetes originales del go/pkg

# Tener las dependencias en tu directorio
```
go mod vendor
```