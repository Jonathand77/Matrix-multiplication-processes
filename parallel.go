package main

import (
	"bufio"  	// Importa el paquete bufio para leer archivos de manera eficiente
	"fmt"    	// Importa el paquete fmt para imprimir en consola y manejar cadenas
	"os"     	// Importa el paquete os para manejar archivos
	"strconv" 	// Importa el paquete strconv para convertir entre cadenas y otros tipos de datos
	"strings"  	// Importa el paquete strings para trabajar con cadenas de texto
	"sync"    	// Importa el paquete sync para sincronización de goroutines
	"time"    	// Importa el paquete time para medir el tiempo de ejecución
)

// Función para cargar una matriz desde un archivo
func loadMatrix(filename string) ([][]int, int, int) {
	// Abre el archivo para leer
	file, err := os.Open(filename)
	if err != nil {
		panic(err) 		// Si hay un error al abrir el archivo, termina el programa
	}
	defer file.Close() // Asegura que el archivo se cierre después de usarlo

	// Usa un scanner para leer línea por línea
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Lee la primera línea con las dimensiones de la matriz
	// Se separan las dimensiones (filas y columnas)
	dims := strings.Fields(scanner.Text()) // Divide la línea en partes
	rows, _ := strconv.Atoi(dims[0])  		// Convierte el primer valor (filas) a entero
	cols, _ := strconv.Atoi(dims[1])  		// Convierte el segundo valor (columnas) a entero

	// Crea una matriz vacía con el tamaño adecuado
	matrix := make([][]int, rows)
	for i := 0; i < rows && scanner.Scan(); i++ { 	// Lee cada fila de la matriz
		line := strings.Fields(scanner.Text()) 		// Divide la línea en valores individuales
		matrix[i] = make([]int, cols) 				// Inicializa un slice vacío para cada fila
		for j := 0; j < cols; j++ { 				// Llena las columnas de la matriz
			matrix[i][j], _ = strconv.Atoi(line[j]) // Convierte cada valor a entero
		}
	}
	return matrix, rows, cols // Devuelve la matriz y sus dimensiones
}

// Función para guardar una matriz en un archivo
func saveMatrix(filename string, matrix [][]int) {
	// Crea el archivo para guardar la matriz
	file, err := os.Create(filename)
	if err != nil {
		panic(err) 		// Si hay un error al crear el archivo, termina el programa
	}
	defer file.Close() // Asegura que el archivo se cierre después de escribir

	// Escribe las dimensiones de la matriz en el archivo
	fmt.Fprintf(file, "%d %d\n", len(matrix), len(matrix[0]))
	// Itera sobre cada fila de la matriz y escribe sus valores
	for _, row := range matrix {
		for _, val := range row {
			fmt.Fprintf(file, "%d ", val)
		}
		fmt.Fprintln(file) // Nueva línea al final de cada fila
	}
}

// Función que ejecuta la multiplicación de una parte de la matriz en paralelo
func multiplyPart(A [][]int, B [][]int, C [][]int, startRow, endRow int, wg *sync.WaitGroup) {
	defer wg.Done() // Asegura que la goroutine se marque como terminada cuando finalice
	n := len(B)     // Número de filas de B (también columnas de A)
	p := len(B[0])  // Número de columnas de B (también columnas de C)

	// Realiza la multiplicación para el rango de filas especificado
	for i := startRow; i < endRow; i++ { // Itera sobre las filas de A
		for j := 0; j < p; j++ { // Itera sobre las columnas de B
			sum := 0
			// Realiza la suma de productos para la multiplicación de matrices
			for k := 0; k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum // Guarda el resultado en la matriz C
		}
	}
}

func main() {
	// Carga las matrices A y B desde los archivos
	A, N, M := loadMatrix("matrix_a.txt")
	B, M2, P := loadMatrix("matrix_b.txt")

	// Verifica si las matrices son compatibles para la multiplicación
	if M != M2 {
		fmt.Println("Error: matrices no compatibles")
		return
	}

	// Crea la matriz C para almacenar el resultado de la multiplicación
	C := make([][]int, N)
	for i := range C {
		C[i] = make([]int, P) // Inicializa cada fila de C con ceros
	}

	// Configuración de las goroutines
	numWorkers := 4             	// Número de trabajadores (goroutines)
	rowsPerWorker := N / numWorkers // Cuántas filas procesará cada trabajador
	remainder := N % numWorkers   	// El resto de filas que no se distribuyen equitativamente

	// Comienza a medir el tiempo de ejecución
	start := time.Now()

	var wg sync.WaitGroup // Crea un objeto WaitGroup para esperar a que terminen las goroutines
	// Lanza una goroutine para cada "trabajador"
	for w := 0; w < numWorkers; w++ {
		// Calcula el rango de filas que le toca a este trabajador
		startRow := w * rowsPerWorker
		endRow := startRow + rowsPerWorker
		if w == numWorkers-1 {
			// Si es el último trabajador, asigna las filas restantes
			endRow += remainder
		}
		wg.Add(1) // Incrementa el contador del WaitGroup
		go multiplyPart(A, B, C, startRow, endRow, &wg) // Lanza la goroutine para procesar una parte
	}
	wg.Wait() // Espera a que todas las goroutines terminen

	// Calcula el tiempo de ejecución
	elapsed := time.Since(start)

	// Guarda la matriz resultante en un archivo
	saveMatrix("matrix_c.txt", C)

	// Imprime el tiempo de ejecución
	fmt.Printf("Tiempo de ejecución paralela (Go): %.6f segundos\n", elapsed.Seconds())
}
