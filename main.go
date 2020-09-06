package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
	-Primer numero de inicio
	-Segundo numero de finalizacion de contador
	-Tercer numero que numeros del array son divisibles entre este numero
*/

func solver(p *[]string, count *int) {
	start, _ := strconv.Atoi((*p)[0])
	end, _ := strconv.Atoi((*p)[1])
	div, _ := strconv.Atoi((*p)[2])

	for i := start; i <= end; i++ {
		if math.Mod(float64(i), float64(div)) == 0 {
			*count = *count + 1
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	datos := strings.Split(scanner.Text(), " ")
	count := 0
	solver(&datos, &count)
	fmt.Println("Numeros divisibles entre", datos[2], ":", count)
}
