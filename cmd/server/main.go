package main

import (
    "fmt"
    "log"
    "os"
	"net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "raspberry-controller/server"
)

func main() {
    // Obtener la ruta del ejecutable
    execPath, err := os.Executable()
    if err != nil {
        log.Fatal("Error al obtener la ruta del ejecutable:", err)
    }

    // Obtener el directorio del ejecutable
    execDir := filepath.Dir(execPath)

    // Construir las rutas relativas a los archivos .env y config.json
    envFiles := []string{
        filepath.Join(execDir, "../config/minidlna/.env"),
        filepath.Join(execDir, "../config/transmission/.env"),
        // Agrega más rutas a archivos .env según sea necesario
    }

    configFiles := []string{
        filepath.Join(execDir, "../config/minidlna/config.json"),
        filepath.Join(execDir, "../config/transmission/config.json"),
        // Agrega más rutas a archivos config.json según sea necesario
    }

    // Verificar si los archivos .env necesarios existen y generarlos si no existen
    for i, envFilePath := range envFiles {
        if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
            fmt.Printf("El archivo %s no existe. Intentando generar desde config.json.\n", envFilePath)
            config, err := loadConfig(configFiles[i])
            if err != nil {
                fmt.Printf("Error al cargar el archivo %s: %v\n", configFiles[i], err)
                return
            }
            err = createEnvFile(envFilePath, config)
            if err != nil {
                fmt.Printf("Error al crear el archivo %s: %v\n", envFilePath, err)
                return
            }
            fmt.Printf("Archivo %s generado correctamente.\n", envFilePath)
        }
    }

    fmt.Println("Todos los archivos .env están presentes.")

    // Crear el enrutador Gin
    r := gin.Default()

    r.HandleMethodNotAllowed = true
    r.NoMethod(func(c *gin.Context) {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
    })

    // Configurar las rutas
    server.SetupRoutes(r)

    // Iniciar el servidor en el puerto 3000
    if err := r.Run(":3000"); err != nil {
        log.Fatal("Error al iniciar el servidor:", err)
    }
}