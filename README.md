# Proyecto de Multiplicación de Matrices Paralela

## 👥 Integrantes

| 👨‍💻 Nombre | 📧 Correo | 🐙 Usuario GitHub |
|---|---|---|
| **Jonathan David Fernandez Vargas** | jonathand.fernandez@udea.edu.co | [jonathand77](https://github.com/jonathand77) |
| **Valeria Alvarez Fernandez** | valeria.alvarezf@udea.edu.co | [vaf88](https://github.com/vaf88) |

---

# 🏆 Programa en Go y C++

Este proyecto implementa una multiplicación de matrices utilizando procesos paralelos en C. Se dividen las filas de la matriz resultante entre varios procesos hijos, los cuales realizan la multiplicación de manera independiente y luego combinan los resultados.

## Archivos del Proyecto

El proyecto contiene los siguientes archivos:

1. **`matrix_multiplication.c`**: Implementa la multiplicación de matrices utilizando múltiples procesos en paralelo.
2. **`matrix_a.txt`**: Archivo que contiene la primera matriz (A) a multiplicar.
3. **`matrix_b.txt`**: Archivo que contiene la segunda matriz (B) a multiplicar.
4. **`matrix_c.txt`**: Archivo de salida donde se guardará el resultado de la multiplicación de matrices.
5. **`README.md`**: Este archivo, que contiene documentación sobre el proyecto.

## Descripción del Proyecto

Este proyecto realiza la multiplicación de dos matrices cuadradas A y B, dividiendo el trabajo en varios procesos hijos para realizarlo de manera paralela, lo que mejora el rendimiento. Se utiliza memoria compartida para almacenar el resultado de la multiplicación y luego se guarda en un archivo de texto.

## Estructura del Proyecto

El proyecto está organizado de la siguiente manera:

```
project/
│
├── matrix_multiplication.c  # Código fuente en C para la multiplicación paralela de matrices.
├── matrix_a.txt             # Archivo de entrada con la matriz A.
├── matrix_b.txt             # Archivo de entrada con la matriz B.
├── matrix_c.txt             # Archivo de salida con el resultado de la multiplicación.
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
gcc matrix_multiplication.c -o matrix_multiplication -lm
```

Este comando compilará el archivo `matrix_multiplication.c` y generará un ejecutable llamado `matrix_multiplication`.

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

Una vez que hayas compilado el código y creado los archivos de entrada, puedes ejecutar el programa con el siguiente comando:

```bash
./matrix_multiplication
```

Este comando ejecutará el programa, realizará la multiplicación de matrices en paralelo y guardará el resultado en `matrix_c.txt`.

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
$ gcc matrix_multiplication.c -o matrix_multiplication -lm
$ ./matrix_multiplication
Tiempo de ejecución paralela: 0.0123 segundos
```

El archivo `matrix_c.txt` contendrá la matriz resultante.

## Explicación del Código

### `matrix_multiplication.c`

El archivo principal del proyecto es `matrix_multiplication.c`. Este archivo contiene la lógica para realizar la multiplicación de matrices de manera paralela utilizando procesos hijos. A continuación, se describen las principales funciones en el archivo:

1. **`load_matrix`**: Carga una matriz desde un archivo de texto. Lee las dimensiones de la matriz y luego llena la matriz con los valores leídos.
   
2. **`save_matrix`**: Guarda una matriz en un archivo de texto. Escribe las dimensiones de la matriz y luego los elementos de la matriz en el archivo.

3. **`free_matrix`**: Libera la memoria dinámica utilizada por una matriz.

4. **`multiply_partial`**: Realiza una multiplicación parcial de las matrices A y B. Cada proceso hijo maneja un conjunto de filas para la multiplicación.

5. **`main`**: Función principal que maneja la carga de las matrices, la creación de memoria compartida, la creación de procesos hijos, y la ejecución de la multiplicación paralela. Además, mide el tiempo de ejecución y guarda el resultado en un archivo de texto.

### Memoria Compartida

Para optimizar la ejecución, el programa utiliza memoria compartida para almacenar el resultado de la multiplicación de matrices. Esto permite que todos los procesos hijos accedan a la misma área de memoria y almacenen los resultados de forma concurrente.

### Multiplicación Paralela

El trabajo se divide entre varios procesos hijos, donde cada proceso maneja una porción de las filas de la matriz resultante. Esta división se realiza de manera equitativa, asegurando que cada proceso tenga una carga de trabajo similar.

## Conclusiones

Este proyecto demuestra cómo realizar la multiplicación de matrices de manera eficiente utilizando programación paralela en C. El uso de procesos y memoria compartida mejora significativamente el tiempo de ejecución al distribuir la carga de trabajo entre múltiples núcleos de la CPU.

---