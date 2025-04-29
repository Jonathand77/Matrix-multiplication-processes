#!/bin/bash

# Definición de colores para la salida
GREEN='\033[0;32m'  # Color verde
NC='\033[0m'        # Sin color (reset)

# Función para convertir el tiempo en formato "0m1.234s" a segundos decimales
convert_time() {
    # Usa awk para dividir el tiempo en minutos (m) y segundos (s)
    echo $1 | awk -Fm '{ split($2,s,"s"); print ($1 * 60) + s[1]; }'
}

# Mensaje de inicio de compilación, con color verde
echo -e "${GREEN}Compilando programas...${NC}"

# Compila los programas secuencial y paralelo
# Cambia los nombres si los archivos están en otra carpeta
gcc sequential.cpp -o sequential.x   # Compila el programa secuencial
gcc parallel.c -o parallel.x         # Compila el programa paralelo

# Mensaje de medición del tiempo para el programa secuencial, con color verde
echo -e "${GREEN}Midiendo tiempo para sequential.x...${NC}"
# Mide el tiempo de ejecución del programa secuencial
SEQUENTIAL_TIME=$( { time ./sequential.x > /dev/null; } 2>&1 | grep real | awk '{print $2}' )

# Mensaje de medición del tiempo para el programa paralelo, con color verde
echo -e "${GREEN}Midiendo tiempo para parallel.x...${NC}"
# Mide el tiempo de ejecución del programa paralelo
PARALLEL_TIME=$( { time ./parallel.x > /dev/null; } 2>&1 | grep real | awk '{print $2}' )

# Convierte los tiempos a segundos decimales usando la función convert_time
SEQ_SECONDS=$(convert_time $SEQUENTIAL_TIME)  # Convierte el tiempo secuencial
PAR_SECONDS=$(convert_time $PARALLEL_TIME)    # Convierte el tiempo paralelo

# Calcula el speedup (mejora de rendimiento)
SPEEDUP=$(echo "scale=2; $SEQ_SECONDS / $PAR_SECONDS" | bc)  # Calcula el speedup usando bc

# Muestra los resultados en la consola
echo ""
echo -e "${GREEN}Resultados:${NC}"
echo "Tiempo secuencial: ${SEQ_SECONDS} segundos"  # Muestra el tiempo secuencial
echo "Tiempo paralelo:   ${PAR_SECONDS} segundos"  # Muestra el tiempo paralelo
echo "Speedup:           ${SPEEDUP}x"              # Muestra el speedup
