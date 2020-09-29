package main

import (
	"fmt"
)

func main() {
	var mes int
	var dia int
	var signo string

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 20 {
			signo = "capricornio"
		} else {
			signo = "acuario"
		}
		break

	case 2:
		if dia > 29 {
			println("Dia invalido.")
			return
		}
		if dia <= 19 {
			signo = "acuario"
		} else {
			signo = "piscis"
		}
		break

	case 3:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 20 {
			signo = "piscis"
		} else {
			signo = "aries"
		}
		break

	case 4:
		if dia > 30 {
			println("Dia invalido.")
			return
		}
		if dia <= 20 {
			signo = "aries"
		} else {
			signo = "tauro"
		}
		break

	case 5:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 21 {
			signo = "tauro"
		} else {
			signo = "geminis"
		}
		break

	case 6:
		if dia > 30 {
			println("Dia invalido.")
			return
		}
		if dia <= 21 {
			signo = "geminis"
		} else {
			signo = "cancer"
		}
		break

	case 7:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 23 {
			signo = "cancer"
		} else {
			signo = "leo"
		}
		break

	case 8:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 23 {
			signo = "leo"
		} else {
			signo = "virgo"
		}
		break

	case 9:
		if dia > 30 {
			println("Dia invalido.")
			return
		}
		if dia <= 23 {
			signo = "virgo"
		} else {
			signo = "libra"
		}
		break

	case 10:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 23 {
			signo = "libra"
		} else {
			signo = "escorpio"
		}
		break

	case 11:
		if dia > 30 {
			println("Dia invalido.")
			return
		}
		if dia <= 22 {
			signo = "escorpio"
		} else {
			signo = "sagitario"
		}
		break

	case 12:
		if dia > 31 {
			println("Dia invalido.")
			return
		}
		if dia <= 21 {
			signo = "sagitario"
		} else {
			signo = "capricornio"
		}
		break
	}

	fmt.Println(signo)
}
