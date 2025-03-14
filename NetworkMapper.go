package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// Funcion para hacer ping a una IP
func ping(ip string, wg *sync.WaitGroup, results chan<- string, osInfo chan<- string) {
	defer wg.Done()
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip) // 1 paquete, timeout de 1s
	_, err := cmd.CombinedOutput()
	if err == nil {
		results <- ip
		osInfo <- detectOS(ip)
	}
}

// Funcion para detectar el sistema operativo basado en TTL
func detectOS(ip string) string {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("%s: No se pudo determinar el SO", ip)
	}

	ttlPattern := regexp.MustCompile(`ttl=(\d+)`)
	matches := ttlPattern.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return fmt.Sprintf("%s: No se pudo determinar el SO", ip)
	}

	ttl := matches[1]
	var osType string
	switch ttl {
	case "64":
		osType = "Linux/Unix"
	case "128":
		osType = "Windows"
	case "255":
		osType = "Cisco/Networking Device"
	default:
		osType = "Desconocido"
	}

	return fmt.Sprintf("%s: %s", ip, osType)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run NetworkMapper.go <rango-IP> (Ejemplo: 192.168.1.1-192.168.1.254)")
		os.Exit(1)
	}

	rangeParts := strings.Split(os.Args[1], "-")
	if len(rangeParts) != 2 {
		fmt.Println("Formato de rango incorrecto. Use: <inicio>-<fin> (Ejemplo: 192.168.1.1-192.168.1.254)")
		os.Exit(1)
	}

	startIP := netToInt(rangeParts[0])
	endIP := netToInt(rangeParts[1])
	if startIP == 0 || endIP == 0 || startIP > endIP {
		fmt.Println("Rango de IP no válido")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	results := make(chan string, 256)
	osInfo := make(chan string, 256)

	// Inicia goroutines para cada IP en el rango
	for ip := startIP; ip <= endIP; ip++ {
		ipStr := intToNet(ip)
		wg.Add(1)
		go ping(ipStr, &wg, results, osInfo)
	}

	// Goroutine para cerrar los canales cuando terminen todas las tareas
	go func() {
		wg.Wait()
		close(results)
		close(osInfo)
	}()

	// Recibir e imprimir resultados
	fmt.Println("Hosts activos y sistemas operativos detectados:")
	for range results {
		fmt.Println(<-osInfo)
	}
}

// Convierte una IP en formato string a un entero para iterar
func netToInt(ip string) int {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return 0
	}
	var num int
	for _, p := range parts {
		val, err := strconv.Atoi(p)
		if err != nil {
			return 0
		}
		num = num<<8 + val
	}
	return num
}

// Convierte un entero a IP en formato string
func intToNet(ip int) string {
	return fmt.Sprintf("%d.%d.%d.%d", (ip>>24)&255, (ip>>16)&255, (ip>>8)&255, ip&255)
}
