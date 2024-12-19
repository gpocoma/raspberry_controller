package services

import (
	"context"
	"os/exec"
    "strings"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// GetStatus verifica si el contenedor 'minidlna' está en ejecución utilizando la nueva API de Docker
func GetStatus() (bool, error) {
	// Crear el cliente Docker
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return false, err
	}
	defer apiClient.Close()

	// Obtener la lista de contenedores (todos, incluso detenidos)
	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return false, err
	}

	// Verificar si el contenedor 'minidlna' está en ejecución
	for _, ctr := range containers {
		for _, name := range ctr.Names {
			if name == "/minidlna" {
				return true, nil
			}
		}
	}

	return false, nil
}


// Startminidlna inicia el contenedor 'minidlna' utilizando docker-compose si no está en ejecución
func StartMinidlna() error {
    // Verificar si el contenedor ya está en ejecución
    active, err := GetStatus()
    if err != nil {
        return err
    }

    if active {
        return nil // El contenedor ya está en ejecución, no hacer nada
    }

    // Iniciar el contenedor utilizando docker-compose
    cmd := exec.Command("docker-compose", "up", "-d", "minidlna")
    cmd.Dir = "/home/pi/minidlna" // Reemplaza con la ruta correcta a tu archivo docker-compose.yml

    output, err := cmd.CombinedOutput()
    if err != nil {
        return err
    }

    if strings.Contains(string(output), "done") {
        return nil
    }

    return nil
}

// Stopminidlna detiene el contenedor 'minidlna' utilizando docker-compose si está en ejecución
func StopMinidlna() error {
    // Verificar si el contenedor está en ejecución
    active, err := GetStatus()
    if err != nil {
        return err
    }

    if !active {
        return nil // El contenedor no está en ejecución, no hacer nada
    }

    // Detener el contenedor utilizando docker-compose
    cmd := exec.Command("docker-compose", "down")
    cmd.Dir = "/home/pi/minidlna" // Reemplaza con la ruta correcta a tu archivo docker-compose.yml

    output, err := cmd.CombinedOutput()
    if err != nil {
        return err
    }

    if strings.Contains(string(output), "done") {
        return nil
    }

    return nil
}