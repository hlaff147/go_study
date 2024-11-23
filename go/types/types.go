package types
import "fmt"

func main() {
	// Tipos numéricos inteiros com sinal
	var intVar int = 2147483647
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807

	// Tipos numéricos inteiros sem sinal
	var uintVar uint = 4294967295
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615

	// Tipos de ponto flutuante
	var float32Var float32 = 3.14
	var float64Var float64 = 3.1415926535

	// Tipos complexos
	var complex64Var complex64 = complex(1, 2)
	var complex128Var complex128 = complex(3, 4)

	// Tipo booleano
	var boolVar bool = true

	// Tipo string
	var stringVar string = "Olá, Mundo!"

	// Tipo byte (sinônimo para uint8)
	var byteVar byte = 'A'

	// Tipo rune (sinônimo para int32)
	var runeVar rune = '🙂'

	// Arrays
	var intArray [5]int = [5]int{1, 2, 3, 4, 5}

	// Slices
	var intSlice []int = []int{6, 7, 8, 9, 10}

	// Mapas
	var mapVar map[string]int = map[string]int{"um": 1, "dois": 2}

	// Ponteiros
	var intPointer *int = &intVar

	// Estruturas (structs)
	type Pessoa struct {
		nome  string
		idade int
	}
	var pessoaVar Pessoa = Pessoa{nome: "João", idade: 30}

	// Imprimindo as variáveis
	fmt.Println("intVar:", intVar)
	fmt.Println("int8Var:", int8Var)
	fmt.Println("int16Var:", int16Var)
	fmt.Println("int32Var:", int32Var)
	fmt.Println("int64Var:", int64Var)
	fmt.Println("uintVar:", uintVar)
	fmt.Println("uint8Var:", uint8Var)
	fmt.Println("uint16Var:", uint16Var)
	fmt.Println("uint32Var:", uint32Var)
	fmt.Println("uint64Var:", uint64Var)
	fmt.Println("float32Var:", float32Var)
	fmt.Println("float64Var:", float64Var)
	fmt.Println("complex64Var:", complex64Var)
	fmt.Println("complex128Var:", complex128Var)
	fmt.Println("boolVar:", boolVar)
	fmt.Println("stringVar:", stringVar)
	fmt.Println("byteVar:", byteVar)
	fmt.Println("runeVar:", runeVar)
	fmt.Println("intArray:", intArray)
	fmt.Println("intSlice:", intSlice)
	fmt.Println("mapVar:", mapVar)
	fmt.Println("intPointer:", intPointer)
	fmt.Println("pessoaVar:", pessoaVar)
}
