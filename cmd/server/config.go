package main

import (
    "encoding/json"
    "os"
)

// Config estructura para almacenar las claves de las variables de entorno
type Config struct {
    Variables map[string]struct {
        Ejemplo string `json:"ejemplo"`
    } `json:"variables"`
}

// loadConfig carga la configuración desde un archivo JSON
func loadConfig(filePath string) (Config, error) {
    var config Config
    file, err := os.Open(filePath)
    if err != nil {
        return config, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    return config, err
}