# Saludos en Go
Este paquete tiene dos funciones de saludo individual y grupal en el cual hay que enviar un slice de diferentes nombres. También cuenta con sus respectivas funciones de unit testing

## Instalación
Ejecuta el siguiente comando para instalar el paquete
```bash
go get -u github.com/odtheame/greetings
```

### Uso
Un pequeño ejemplo del uso del módulo
```go
package main

import (
	"fmt"
	"log"

	greetings "github.com/odtheame/ModuloGo"
)

func main() {
    // Función de saludo individual
    message, err := greetings.Hello("Nova")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)

    // Función de saludo grupal
	names := []string{"Nova", "Odtheame", "Doth"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
```