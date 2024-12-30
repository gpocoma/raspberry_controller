package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// createEnvFile crea un archivo .env solicitando los valores al usuario
func createEnvFile(envFilePath string, config Config) error {
    file, err := os.Create(envFilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := bufio.NewReader(os.Stdin)
    for key, value := range config.Variables {
        fmt.Printf("Enter %s (e.g., %s): ", key, value.Ejemplo)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "" {
            input = value.Ejemplo
        }
        _, err := file.WriteString(key + "=" + input + "\n")
        if err != nil {
            return err
        }
    }

    return nil
}