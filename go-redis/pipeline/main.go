// Redis Pipeline
/* Permite implementar un servidor de Solicitud/Respuesta
para que pueda procesar nuevas solicitudes, incluso si el
cliente no leyó las respuestas anteriores
De esta forma se puede enviar multiples comandos al servidor
sin esperar las respuestas.
*/
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/my/repo/config"
)

var ctx = context.Background()

func main() {
	client := config.NewClientRedis()
	pipe := client.Pipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(incr.Val())
}
