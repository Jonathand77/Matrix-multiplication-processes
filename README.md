# Proyecto de Multiplicación de Matrices Paralela

## 👥 Integrantes

| 👨‍💻 Nombre | 📧 Correo | 🐙 Usuario GitHub |
|---|---|---|
| **Jonathan David Fernandez Vargas** | jonathand.fernandez@udea.edu.co | [jonathand77](https://github.com/jonathand77) |
| **Valeria Alvarez Fernandez** | valeria.alvarezf@udea.edu.co | [vaf88](https://github.com/vaf88) |

---

# 🏆 Programa en Go y C++

Este proyecto tiene como objetivo comparar el rendimiento entre versiones secuenciales y paralelas de multiplicación de matrices. Incluye implementaciones en C, C++ y Go, así como un script para automatizar pruebas de rendimiento.

## Archivos del Proyecto

El proyecto contiene los siguientes archivos:

1. **`sequential.cpp`**: Implementación secuencial de la multiplicación de matrices en C++.
2. **`parallel.c`**: Implementación paralela en C, usando procesos (fork) y memoria compartida (shm).
3. **`sequential.go`**: Implementación secuencial de multiplicación en lenguaje Go.
4. **`parallel.go`**: Versión paralela escrita en Go.
5. **`sequential.x`**: Archivo binario compilado a partir de sequential.cpp.
6. **`parallel.x`**: Archivo binario compilado a partir de parallel.c.
7. **`benchmark.sh`**: Script Bash que compila los programas, ejecuta las versiones secuencial y paralela, mide el tiempo de ejecución y calcula el speedup.
8. **`matrix_a.txt`**: Archivo de texto que contiene la primera matriz a multiplicar (formato: filas, columnas, luego los datos).
9. **`matrix_b.txt`**: Archivo de texto que contiene la segunda matriz a multiplicar.
10. **`matrix_c.txt`**: Archivo generado como salida, contiene la matriz resultado de la multiplicación.
11. **`README.md`**: Este archivo, que contiene documentación sobre el proyecto.

## Descripción del Proyecto

Este proyecto realiza la multiplicación de dos matrices cuadradas A y B, dividiendo el trabajo en varios procesos hijos para realizarlo de manera paralela, lo que mejora el rendimiento. Se utiliza memoria compartida para almacenar el resultado de la multiplicación y luego se guarda en un archivo de texto.

## Estructura del Proyecto

El proyecto está organizado de la siguiente manera:

```
project/
│
├── sequential.cpp           #Implementación secuencial de la multiplicación de matrices en C++.
├── parallel.c               #Implementación paralela en C, usando procesos (fork) y memoria compartida (shm)..
├── sequential.go            #Implementación secuencial de multiplicación en lenguaje Go.
├── parallel.go              #Versión paralela escrita en Go.
├── sequential.x             #Archivo binario compilado a partir de sequential.cpp.
├── parallel.x               #Archivo binario compilado a partir de parallel.c.
├── benchmark.sh             #Script Bash que compila los programas, ejecuta las versiones secuencial y paralela, mide el tiempo de ejecución y calcula el speedup.
├── matrix_a.txt             #Archivo de texto que contiene la primera matriz a multiplicar (formato: filas, columnas, luego los datos).
├── matrix_b.txt             #Archivo de texto que contiene la segunda matriz a multiplicar.
├── matrix_c.txt             #Archivo generado como salida, contiene la matriz resultado de la multiplicación.
└── README.md                # Este archivo.
```

## Requisitos

Para ejecutar este proyecto, necesitas tener un compilador de C y las herramientas estándar de Linux. El código utiliza la biblioteca estándar de C y las llamadas al sistema para la manipulación de procesos e IPC (memoria compartida).

### Requisitos previos:

- Un sistema Linux o similar.
- Un compilador de C (por ejemplo, `gcc`).
- Herramientas estándar de desarrollo en C (como `make`, `gcc`).

## Compilación del Proyecto

Para compilar el código, abre una terminal y navega hasta el directorio del proyecto. Luego, usa el siguiente comando para compilar el archivo C:

```bash
gcc sequential.cpp -o sequential.x
gcc parallel.c -o parallel.x
```

```bash
chmod +x benchmark.sh
./benchmark.sh
```

Este comando compilará el archivo `sequential.cpp` y `parallel.c` usando gcc, y guarda los binarios como `sequential.x` y `parallel.x`.

## Ejecución del Proyecto

### Paso 1: Crear las matrices de entrada

Antes de ejecutar el programa, necesitas crear los archivos `matrix_a.txt` y `matrix_b.txt`. Estos archivos deben contener las matrices A y B, respetando el siguiente formato:

- La primera línea debe contener dos enteros: el número de filas y columnas de la matriz (en el caso de matrices cuadradas, estos valores serán iguales).
- Las siguientes líneas deben contener los elementos de la matriz, separados por espacios.

Ejemplo de `matrix_a.txt`:
```
3 3
1 2 3
4 5 6
7 8 9
```

Ejemplo de `matrix_b.txt`:
```
3 3
9 8 7
6 5 4
3 2 1
```

### Paso 2: Ejecutar el programa

Una vez que hayas compilado el código y creado los archivos de entrada, puedes ejecutar el programa con los siguientes comandos:

```bash
./sequential.x
# o
./parallel.x
```

Estos comandos ejecutará el programa, realizará la multiplicación de matrices en paralelo y guardará el resultado en `matrix_c.txt`.

### Paso 3: Ver el resultado

Después de la ejecución, el resultado de la multiplicación de matrices se guardará en el archivo `matrix_c.txt` en el siguiente formato:

```
3 3
<elemento_1_1> <elemento_1_2> <elemento_1_3>
<elemento_2_1> <elemento_2_2> <elemento_2_3>
<elemento_3_1> <elemento_3_2> <elemento_3_3>
```

### Ejemplo de ejecución:

```
$ gcc sequential.cpp -o sequential.x
$ ./sequential.x
Tiempo de ejecución paralela: 0.0123 segundos
```

El archivo `matrix_c.txt` contendrá la matriz resultante.

## Explicación del Código

### `matrix_multiplication.c`

El archivo principal del proyecto es `matrix_c.txt`. Este archivo contiene la lógica para realizar la multiplicación de matrices de manera paralela utilizando procesos hijos. A continuación, se describen las principales funciones en el archivo:

1. **`load_matrix`**: Carga una matriz desde un archivo de texto. Lee las dimensiones de la matriz y luego llena la matriz con los valores leídos.
   
2. **`save_matrix`**: Guarda una matriz en un archivo de texto. Escribe las dimensiones de la matriz y luego los elementos de la matriz en el archivo.

3. **`free_matrix`**: Libera la memoria dinámica utilizada por una matriz.

4. **`multiply_partial`**: Realiza una multiplicación parcial de las matrices A y B. Cada proceso hijo maneja un conjunto de filas para la multiplicación.

5. **`main`**: Función principal que maneja la carga de las matrices, la creación de memoria compartida, la creación de procesos hijos, y la ejecución de la multiplicación paralela. Además, mide el tiempo de ejecución y guarda el resultado en un archivo de texto.

### Memoria Compartida

Para optimizar la ejecución, el programa utiliza memoria compartida para almacenar el resultado de la multiplicación de matrices. Esto permite que todos los procesos hijos accedan a la misma área de memoria y almacenen los resultados de forma concurrente.

### Multiplicación Paralela

El trabajo se divide entre varios procesos hijos, donde cada proceso maneja una porción de las filas de la matriz resultante. Esta división se realiza de manera equitativa, asegurando que cada proceso tenga una carga de trabajo similar.

## ✅Conclusiones

1. Paralelismo mejora el rendimiento significativamente
La versión paralela, implementada en C usando fork() y memoria compartida (shmget), demostró una reducción notable en el tiempo de ejecución al distribuir el trabajo entre varios procesos. Esto confirma que dividir tareas en múltiples núcleos es una estrategia eficiente para operaciones intensivas como la multiplicación de matrices.

2. El uso de memoria compartida fue clave
La correcta sincronización mediante memoria compartida permitió que todos los procesos hijos escribieran en la misma matriz resultado sin necesidad de mecanismos complejos de comunicación, logrando una implementación paralela funcional y coherente.

3. La versión secuencial es útil como línea base
La implementación secuencial permite comparar el rendimiento base del algoritmo sin paralelismo. Es esencial para calcular el speedup y validar la corrección de la versión paralela.

4. El script benchmark.sh facilita las pruebas y análisis
Automatizar la compilación, ejecución, medición de tiempos y cálculo de speedup permitió un análisis más claro y reproducible del comportamiento de ambas versiones.

5. Escalabilidad limitada por el número de filas y procesos
Si el número de filas de la matriz no es divisible equitativamente entre los procesos, uno de ellos terminará trabajando más que los otros. A mayor tamaño de matrices y procesos, se podrían investigar técnicas de balance de carga más avanzadas.

6. Potencial para extenderse a otras plataformas o lenguajes
El diseño modular del proyecto (archivos separados, entrada/salida clara, script automatizado) permite fácilmente portar o extender la solución a otros lenguajes como Go o Python, o incluso paralelismo con hilos en lugar de procesos.

---
