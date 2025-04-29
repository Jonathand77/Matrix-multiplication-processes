#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/shm.h>
#include <sys/wait.h>
#include <time.h>

#define MAX_PROCESSES 8  // Número máximo de procesos que se pueden usar, ajustable

// Función para cargar una matriz desde un archivo
int** load_matrix(const char* filename, int* rows, int* cols) {
    // Abre el archivo en modo lectura
    FILE* file = fopen(filename, "r");
    if (!file) {
        perror("Archivo");
        exit(1);
    }

    // Lee las dimensiones de la matriz (filas y columnas)
    fscanf(file, "%d %d", rows, cols);
    
    // Reserva memoria dinámica para la matriz
    int** matrix = malloc(*rows * sizeof(int*));
    for (int i = 0; i < *rows; i++) {
        matrix[i] = malloc(*cols * sizeof(int));  // Reserva memoria para cada fila
        for (int j = 0; j < *cols; j++) {
            fscanf(file, "%d", &matrix[i][j]);  // Lee cada valor de la matriz
        }
    }
    fclose(file);  // Cierra el archivo
    return matrix;  // Devuelve la matriz cargada
}

// Función para guardar una matriz en un archivo
void save_matrix(const char* filename, int* C, int N, int P) {
    // Abre el archivo en modo escritura
    FILE* file = fopen(filename, "w");
    // Escribe las dimensiones de la matriz
    fprintf(file, "%d %d\n", N, P);
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < P; j++) {
            fprintf(file, "%d ", C[i * P + j]);  // Escribe el valor de cada elemento
        }
        fprintf(file, "\n");  // Salto de línea entre filas
    }
    fclose(file);  // Cierra el archivo
}

// Función para liberar la memoria de una matriz
void free_matrix(int** matrix, int rows) {
    for (int i = 0; i < rows; i++) free(matrix[i]);  // Libera cada fila
    free(matrix);  // Libera la memoria de la matriz
}

// Función para realizar una multiplicación parcial de matrices
void multiply_partial(int** A, int** B, int* C, int start_row, int end_row, int M, int P) {
    // Multiplica las matrices A y B de manera parcial, de acuerdo al rango de filas asignado
    for (int i = start_row; i < end_row; i++) {
        for (int j = 0; j < P; j++) {
            C[i * P + j] = 0;
            for (int k = 0; k < M; k++) {
                C[i * P + j] += A[i][k] * B[k][j];  // Realiza el cálculo de la multiplicación
            }
        }
    }
}

int main() {
    int N, M, M2, P;
    int num_processes = 4;  // Número de procesos hijos a crear, ajustable

    // Comienza el temporizador para medir el tiempo de ejecución
    clock_t start = clock();

    // Carga las matrices A y B desde los archivos
    int** A = load_matrix("matrix_a.txt", &N, &M);
    int** B = load_matrix("matrix_b.txt", &M2, &P);
    if (M != M2) {
        fprintf(stderr, "Error: dimensiones incompatibles.\n");
        return 1;  // Si las dimensiones no son compatibles, termina el programa
    }

    // Crear memoria compartida para la matriz C (resultado de la multiplicación)
    int shmid = shmget(IPC_PRIVATE, sizeof(int) * N * P, IPC_CREAT | 0666);  // Crea un segmento de memoria compartida
    if (shmid < 0) {
        perror("shmget");  // Si falla la creación de la memoria compartida, muestra un error
        exit(1);
    }
    int* C = (int*)shmat(shmid, NULL, 0);  // Asocia la memoria compartida al proceso

    // Dividir el trabajo entre los procesos hijos
    int rows_per_process = N / num_processes;   // Número de filas que cada proceso va a manejar
    int remaining_rows = N % num_processes;     // Filas sobrantes que no se distribuyen equitativamente

    // Crear procesos hijos
    for (int p = 0; p < num_processes; p++) {
        int start_row = p * rows_per_process;       // Fila inicial para este proceso
        int end_row = (p + 1) * rows_per_process;   // Fila final para este proceso
        if (p == num_processes - 1) {
            end_row += remaining_rows;              // El último proceso maneja las filas restantes
        }

        pid_t pid = fork();  // Crea un nuevo proceso
        if (pid < 0) {
            perror("fork");  // Si falla la creación del proceso, muestra un error
            exit(1);
        }
        if (pid == 0) {
            // Código que ejecuta el proceso hijo
            multiply_partial(A, B, C, start_row, end_row, M, P);  // Realiza la multiplicación parcial
            shmdt(C);   // Desasocia la memoria compartida
            exit(0);    // Termina el proceso hijo
        }
    }

    // Espera a que todos los procesos hijos terminen
    for (int i = 0; i < num_processes; i++) {
        wait(NULL);  // Espera que cada proceso hijo termine
    }

    // Guarda el resultado en un archivo
    save_matrix("matrix_c.txt", C, N, P);

    // Detiene el temporizador y calcula el tiempo de ejecución
    clock_t end = clock();
    double time_taken = (double)(end - start) / CLOCKS_PER_SEC;
    printf("Tiempo de ejecución paralela: %.4f segundos\n", time_taken);  // Imprime el tiempo de ejecución

    // Libera la memoria utilizada
    shmdt(C);  // Desasocia la memoria compartida
    shmctl(shmid, IPC_RMID, NULL);  // Elimina el segmento de memoria compartida
    free_matrix(A, N);              // Libera la memoria de la matriz A
    free_matrix(B, M);              // Libera la memoria de la matriz B

    return 0;  // Termina el programa
}
