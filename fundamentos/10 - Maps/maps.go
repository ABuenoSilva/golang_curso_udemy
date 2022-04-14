package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuario := map[string]string{
		"nome":      "Alexandre",
		"sobrenome": "Bueno",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Alexandre",
			"segundo":  "Bueno",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Sacadura",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2["curso"])
	fmt.Println(usuario2["curso"]["nome"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2["curso"])
	fmt.Println(usuario2["curso"]["nome"])

	usuario2["signo"] = map[string]string{"nome": "Leão"}
	fmt.Println(usuario2)
	fmt.Println(usuario2["signo"])
	fmt.Println(usuario2["signo"]["nome"])
	fmt.Println(usuario2["curso"])
	fmt.Println(usuario2["curso"]["nome"])

}
