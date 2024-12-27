package services

import (
    "context"
    "os/exec"
    "strings"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "raspberry-controller/minidlna_app/models"
)

// GetStatus verifica si el contenedor 'minidlna' está en ejecución utilizando la nueva API de Docker
func GetStatus() (models.MinidlnaStatus, error) {
    // Crear el cliente Docker
    apiClient, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        return models.MinidlnaStatus{Active: false}, err
    }
    defer apiClient.Close()

    // Obtener la lista de contenedores (todos, incluso detenidos)
    containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
    if err != nil {
        return models.MinidlnaStatus{Active: false}, err
    }

    // Verificar si el contenedor 'minidlna' está en ejecución
    for _, ctr := range containers {
        for _, name := range ctr.Names {
            if name == "/minidlna" {
                return models.MinidlnaStatus{Active: true}, nil
            }
        }
    }

    return models.MinidlnaStatus{Active: false}, nil
}

// StartMinidlna inicia el contenedor 'minidlna' utilizando docker-compose si no está en ejecución
func StartMinidlna() (models.MinidlnaMessage, error) {
    // Verificar si el contenedor ya está en ejecución
    status, err := GetStatus()
    if err != nil {
        return models.MinidlnaMessage{}, err
    }

    if status.Active {
        return models.MinidlnaMessage{Message: "Minidlna is already running"}, nil // El contenedor ya está en ejecución, no hacer nada
    }

    // Iniciar el contenedor utilizando docker-compose
    cmd := exec.Command("docker-compose", "up", "-d", "minidlna")
    cmd.Dir = "/home/pi/minidlna" // Reemplaza con la ruta correcta a tu archivo docker-compose.yml

    output, err := cmd.CombinedOutput()
    if err != nil {
        return models.MinidlnaMessage{}, err
    }

    if strings.Contains(string(output), "done") {
        return models.MinidlnaMessage{Message: "Minidlna started successfully"}, nil
    }

    return models.MinidlnaMessage{Message: "Minidlna started"}, nil
}

// StopMinidlna detiene el contenedor 'minidlna' utilizando docker-compose si está en ejecución
func StopMinidlna() (models.MinidlnaMessage, error) {
    // Verificar si el contenedor está en ejecución
    status, err := GetStatus()
    if err != nil {
        return models.MinidlnaMessage{}, err
    }

    if !status.Active {
        return models.MinidlnaMessage{Message: "Minidlna is not running"}, nil // El contenedor no está en ejecución, no hacer nada
    }

    // Detener el contenedor utilizando docker-compose
    cmd := exec.Command("docker-compose", "down")
    cmd.Dir = "/home/pi/minidlna" // Reemplaza con la ruta correcta a tu archivo docker-compose.yml

    output, err := cmd.CombinedOutput()
    if err != nil {
        return models.MinidlnaMessage{}, err
    }

    if strings.Contains(string(output), "done") {
        return models.MinidlnaMessage{Message: "Minidlna stopped successfully"}, nil
    }

    return models.MinidlnaMessage{Message: "Minidlna stopped"}, nil
}