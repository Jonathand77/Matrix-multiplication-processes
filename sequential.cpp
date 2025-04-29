#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// Función para cargar una matriz desde un archivo
int** load_matrix(const char* filename, int* rows, int* cols) {
    // Abre el archivo en modo lectura
    FILE* file = fopen(filename, "r");
    if (!file) {
        perror("Error abriendo archivo");
        exit(1);
    }

    // Lee las dimensiones de la matriz (filas y columnas)
    fscanf(file, "%d %d", rows, cols);

    // Reserva memoria para la matriz
    int** matrix = (int**)malloc(*rows * sizeof(int*));
    for (int i = 0; i < *rows; i++) {
        matrix[i] = (int*)malloc(*cols * sizeof(int));  // Reserva memoria para cada fila
        for (int j = 0; j < *cols; j++) {
            fscanf(file, "%d", &matrix[i][j]);          // Lee el valor de cada elemento de la matriz
        }
    }

    // Cierra el archivo después de leerlo
    fclose(file);
    return matrix;  // Devuelve la matriz cargada
}

// Función para multiplicar dos matrices
int** multiply_matrices(int** A, int** B, int N, int M, int P) {
    // Reserva memoria para la matriz resultante
    int** C = (int**)malloc(N * sizeof(int*));
    for (int i = 0; i < N; i++) {
        C[i] = (int*)calloc(P, sizeof(int));  // Inicializa la matriz resultante en 0
        for (int j = 0; j < P; j++) {
            for (int k = 0; k < M; k++) {
                C[i][j] += A[i][k] * B[k][j];  // Calcula el valor de la matriz C
            }
        }
    }
    return C;  // Devuelve la matriz resultante
}

// Función para guardar una matriz en un archivo
void save_matrix(const char* filename, int** matrix, int rows, int cols) {
    // Abre el archivo en modo escritura
    FILE* file = fopen(filename, "w");
    // Escribe las dimensiones de la matriz en el archivo
    fprintf(file, "%d %d\n", rows, cols);
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            fprintf(file, "%d ", matrix[i][j]);  // Escribe cada valor de la matriz
        }
        fprintf(file, "\n");    // Salto de línea entre filas
    }
    fclose(file);               // Cierra el archivo después de escribir
}

// Función para liberar la memoria de una matriz
void free_matrix(int** matrix, int rows) {
    for (int i = 0; i < rows; i++) free(matrix[i]);  // Libera cada fila
    free(matrix);  // Libera la memoria de la matriz
}

int main() {
    int N, M, M2, P;  // Variables para las dimensiones de las matrices

    // Comienza el temporizador para medir el tiempo de ejecución
    clock_t start = clock();

    // Carga las matrices A y B desde los archivos
    int** A = load_matrix("matrix_a.txt", &N, &M);
    int** B = load_matrix("matrix_b.txt", &M2, &P);

    // Verifica que las dimensiones de las matrices sean compatibles para la multiplicación
    if (M != M2) {
        fprintf(stderr, "Error: dimensiones incompatibles.\n");
        return 1;
    }

    // Multiplica las matrices A y B y guarda el resultado en C
    int** C = multiply_matrices(A, B, N, M, P);

    // Guarda la matriz resultante C en un archivo
    save_matrix("matrix_c.txt", C, N, P);

    // Detiene el temporizador
    clock_t end = clock();
    // Calcula el tiempo de ejecución en segundos
    double time_taken = (double)(end - start) / CLOCKS_PER_SEC;
    // Imprime el tiempo de ejecución secuencial
    printf("Tiempo de ejecución secuencial: %.4f segundos\n", time_taken);

    // Libera la memoria de las matrices A, B y C
    free_matrix(A, N);
    free_matrix(B, M);
    free_matrix(C, N);

    return 0;
}
