package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 10000)
	fmt.Printf("Literal inteiro é :", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte é :", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println(i1)
	fmt.Println(reflect.TypeOf(i1))

	//mapeamento da tabela unicode(int32)
	var i2 rune = 'a'
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2)

	//float32 e float64

	var x float32 = 49.33
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(49.99))

	//variaveis booleanas

	bo := false
	fmt.Println(bo)
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(!bo)

	//variaveis string

	s1 := "String é aspas duplas SEMPRE"
	fmt.Println(s1, len(s1))

	//char? ----> Não existe em GO, ele gera em int32 e mapeia o valor para o unicode
	char := 'a'
	fmt.Println(reflect.TypeOf(char))

}
