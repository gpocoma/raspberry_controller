package services

import (
    "context"
    "os/exec"
    "strings"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "raspberry-controller/transmission_app/models"
)

// GetStatus verifica si el contenedor 'transmission' está en ejecución utilizando la nueva API de Docker
func GetStatus() (models.TransmissionStatus, error) {
    // Crear el cliente Docker
    apiClient, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        return models.TransmissionStatus{Active: false}, err
    }
    defer apiClient.Close()

    // Obtener la lista de contenedores (todos, incluso detenidos)
    containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
    if err != nil {
        return models.TransmissionStatus{Active: false}, err
    }

    // Verificar si el contenedor 'transmission' está en ejecución
    for _, ctr := range containers {
        for _, name := range ctr.Names {
            if strings.Contains(name, "transmission") {
                return models.TransmissionStatus{Active: true}, nil
            }
        }
    }

    return models.TransmissionStatus{Active: false}, nil
}

// StartTransmission inicia el contenedor 'transmission' utilizando docker-compose si no está en ejecución
func StartTransmission() (models.TransmissionMessage, error) {
    // Verificar si el contenedor ya está en ejecución
    status, err := GetStatus()
    if err != nil {
        return models.TransmissionMessage{}, err
    }

    if status.Active {
        return models.TransmissionMessage{Message: "Transmission is already running"}, nil // El contenedor ya está en ejecución, no hacer nada
    }

    // Iniciar el contenedor utilizando docker-compose
    cmd := exec.Command("docker", "compose", "up", "-d", "transmission")
    cmd.Dir = "/home/pi/docker/transmission"

    output, err := cmd.CombinedOutput()
    if err != nil {
        return models.TransmissionMessage{}, err
    }

    if strings.Contains(string(output), "done") {
        return models.TransmissionMessage{Message: "Transmission started successfully"}, nil
    }

    return models.TransmissionMessage{Message: "Transmission started"}, nil
}

// StopTransmission detiene el contenedor 'transmission' utilizando docker-compose si está en ejecución
func StopTransmission() (models.TransmissionMessage, error) {
    // Verificar si el contenedor está en ejecución
    status, err := GetStatus()
    if err != nil {
        return models.TransmissionMessage{}, err
    }

    if !status.Active {
        return models.TransmissionMessage{Message: "Transmission is not running"}, nil // El contenedor no está en ejecución, no hacer nada
    }

    // Detener el contenedor utilizando docker-compose
    cmd := exec.Command("docker", "compose", "down")
    cmd.Dir = "/home/pi/docker/transmission"

    output, err := cmd.CombinedOutput()
    if err != nil {
        return models.TransmissionMessage{}, err
    }

    if strings.Contains(string(output), "done") {
        return models.TransmissionMessage{Message: "Transmission stopped successfully"}, nil
    }

    return models.TransmissionMessage{Message: "Transmission stopped"}, nil
}