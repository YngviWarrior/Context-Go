package main

import (
	"context"
	"fmt"
	"time"
)

// Context funciona como árvore.
func main() {
	// Contexto em branco, raíz (pai)!
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*2) // já faz o cancelamento automático.

	ctx = context.WithValue(ctx, "key", 10)
	fmt.Println(ctx.Value("key"))

	//sempre cancela no fim, por boas práticas.
	defer cancel()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	cancel()
	// }()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): //cancel ativa .Done()
		fmt.Println("Tempo excedido para bookar o quarto.")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso.")
	}
}
