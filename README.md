# Network Mapper

Este es un script en **Go** para mapear una red, detectar hosts activos y estimar el sistema operativo de cada host basado en el TTL de la respuesta del ping.

## Características
- **Escaneo de red**: Detecta qué hosts están activos en un rango de IPs especificado.
- **Detección de sistema operativo**: Estima el SO en función del TTL en las respuestas del ping.
- **Paralelización**: Usa goroutines para mejorar el rendimiento.
- **Uso de argumentos**: Permite especificar el rango de IPs a escanear desde la línea de comandos.

## Requisitos
- **Go** instalado en tu sistema (`go version` para verificar).
- Permisos de **administrador/root** si es necesario para ejecutar `ping`.

## Instalación
1. Clona este repositorio o copia el código en un archivo:
   ```sh
   git clone https://github.com/fervigz090/NetworkMapper.git
   cd NetworkMapper
   ```
2. Compila el script:
   ```sh
   go build -o network_mapper NetworkMapper.go
   ```

## Uso
Ejecuta el script con el rango de IPs que quieres escanear:

```sh
sudo ./network_mapper 192.168.1.1-192.168.1.254
```

### Ejemplo de salida
```
Hosts activos y sistemas operativos detectados:
192.168.1.100: Windows
192.168.1.101: Linux/Unix
192.168.1.105: Desconocido
```

## Mejoras futuras
- Escaneo de puertos abiertos.
- Integración con bases de datos para guardar resultados.
- Generación de reportes en formato JSON o CSV.

## Licencia
Este proyecto se distribuye bajo la licencia **MIT**.

---
¡Disfruta escaneando redes con Go!

