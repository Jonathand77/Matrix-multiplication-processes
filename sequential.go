package main

import (
	"bufio"  	// Importa el paquete bufio para la lectura eficiente de archivos
	"fmt"    	// Importa el paquete fmt para imprimir en consola y manejar texto
	"os"     	// Importa el paquete os para manejar archivos
	"strconv" 	// Importa el paquete strconv para convertir entre cadenas y otros tipos
	"strings"  	// Importa el paquete strings para manipular cadenas de texto
	"time"    	// Importa el paquete time para medir el tiempo de ejecución
)

// Función para cargar una matriz desde un archivo
func loadMatrix(filename string) ([][]int, int, int) {
	// Abre el archivo para leer
	file, err := os.Open(filename)
	if err != nil {
		panic(err) 		// Si ocurre un error al abrir el archivo, se detiene el programa
	}
	defer file.Close() 	// Asegura que el archivo se cierre después de que termine de usarse

	// Usa un scanner para leer línea por línea
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Lee la primera línea con las dimensiones de la matriz
	// Se separan las dimensiones de la matriz (filas y columnas)
	dims := strings.Fields(scanner.Text()) 	// Dividir la línea por espacios
	rows, _ := strconv.Atoi(dims[0])  		// Convierte el primer valor (filas) a entero
	cols, _ := strconv.Atoi(dims[1])  		// Convierte el segundo valor (columnas) a entero

	// Se crea una matriz vacía con el tamaño adecuado
	matrix := make([][]int, rows)
	for i := 0; i < rows && scanner.Scan(); i++ { 	// Lee cada fila de la matriz
		line := strings.Fields(scanner.Text()) 		// Divide la línea en valores individuales
		matrix[i] = make([]int, cols) 				// Asigna un nuevo slice para cada fila
		for j := 0; j < cols; j++ { 				// Llena cada columna con el valor correspondiente
			matrix[i][j], _ = strconv.Atoi(line[j]) // Convierte cada valor a entero
		}
	}
	return matrix, rows, cols // Devuelve la matriz junto con las dimensiones
}

// Función para guardar una matriz en un archivo
func saveMatrix(filename string, matrix [][]int) {
	// Crea un archivo para guardar la matriz
	file, err := os.Create(filename)
	if err != nil {
		panic(err) 			// Si hay un error al crear el archivo, se detiene el programa
	}
	defer file.Close() 		// Asegura que el archivo se cierre después de escribir

	// Escribe las dimensiones de la matriz en el archivo
	fmt.Fprintf(file, "%d %d\n", len(matrix), len(matrix[0]))
	// Itera sobre cada fila de la matriz
	for _, row := range matrix {
		// Itera sobre cada valor en la fila y lo escribe en el archivo
		for _, val := range row {
			fmt.Fprintf(file, "%d ", val)
		}
		// Escribe una nueva línea después de cada fila
		fmt.Fprintln(file)
	}
}

// Función principal
func main() {
	// Carga las matrices A y B desde los archivos
	A, N, M := loadMatrix("matrix_a.txt")
	B, M2, P := loadMatrix("matrix_b.txt")

	// Verifica si las matrices tienen dimensiones compatibles para multiplicación
	if M != M2 {
		fmt.Println("Error: las matrices no son compatibles para multiplicación")
		return
	}

	// Crea una matriz C para almacenar el resultado de la multiplicación
	C := make([][]int, N)
	for i := range C {
		C[i] = make([]int, P) 	// Inicializa cada fila de C con ceros
	}

	// Comienza la medición del tiempo de ejecución
	start := time.Now()

	// Realiza la multiplicación de matrices
	for i := 0; i < N; i++ { 		// Itera sobre las filas de A
		for j := 0; j < P; j++ { 	// Itera sobre las columnas de B
			sum := 0
			// Realiza la suma de productos (A[i][k] * B[k][j])
			for k := 0; k < M; k++ {
				sum += A[i][k] * B[k][j]
			}
			// Guarda el resultado en la matriz C
			C[i][j] = sum
		}
	}

	// Calcula el tiempo de ejecución
	elapsed := time.Since(start)

	// Guarda el resultado en el archivo matrix_c.txt
	saveMatrix("matrix_c.txt", C)
	// Imprime el tiempo de ejecución en consola
	fmt.Printf("Tiempo de ejecución secuencial (Go): %.6f segundos\n", elapsed.Seconds())
}
