package main

import (
	"fmt"
	"strconv"
	"time"
)

var band bool = false
var index int = -1

func Process(n int) {
	i := 0
	for {
		if(band) {
			fmt.Printf("process %d : %d \n", n, i)
		}
		i++
		time.Sleep(time.Millisecond * 500)
		if(index == n){
			index = -1
			break
		}
	}
}

func main() {
	id := 1
	opc := ""
	for opc != "4" {
		fmt.Printf("\t\tMenu\n" +
			"\t1) Crear proceso\n" +
			"\t2) Ver Procesos\n" +
			"\t3) Eliminar Proceso\n" +
			"\t4) SALIR\n" +
			"opc: ")
		fmt.Scanln(&opc)
		switch opc {
		case "1":
			go Process(id)
			id++
		case "2":
			band = true
			fmt.Scanln()
			band = false
		case "3":
			band := true
			for band {
				fmt.Printf("id : ")
				fmt.Scanln(&opc)
				i, err := strconv.Atoi(opc)
				if err != nil {
					fmt.Println("Ingresa un numero valido")
				} else {
					index = i
					band = false
					opc = "3"
				}
			}
		case "4":
			fmt.Println("Hasta pronto!")
		default:
			fmt.Println("Opcion no valida")

		}

	}
}
