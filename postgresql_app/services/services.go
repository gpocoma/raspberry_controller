package services

import (
    "context"
    "os/exec"
    "strings"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "raspberry-controller/postgresql_app/models"
)

// GetStatus verifica si el contenedor 'postgresql-db' está en ejecución utilizando la nueva API de Docker
func GetStatus() (models.PostgreSQLStatus, error) {
    // Crear el cliente Docker
    apiClient, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        return models.PostgreSQLStatus{Active: false}, err
    }
    defer apiClient.Close()

    // Obtener la lista de contenedores (todos, incluso detenidos)
    containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
    if err != nil {
        return models.PostgreSQLStatus{Active: false}, err
    }

    // Verificar si el contenedor 'postgresql-db' está en ejecución
    for _, ctr := range containers {
        for _, name := range ctr.Names {
            if (name == "/postgresql-db" || name == "postgresql-db") && ctr.State == "running" {
                return models.PostgreSQLStatus{Active: true}, nil
            }
        }
    }

    return models.PostgreSQLStatus{Active: false}, nil
}

// StartPostgreSQL inicia el contenedor 'postgresql-db' utilizando docker-compose si no está en ejecución
func StartPostgreSQL() (models.PostgreSQLMessage, error) {
	// Verificar si el contenedor ya está en ejecución
	status, err := GetStatus()
	if err != nil {
		return models.PostgreSQLMessage{}, err
	}

	if status.Active {
		return models.PostgreSQLMessage{Message: "PostgreSQL is already running"}, nil // El contenedor ya está en ejecución, no hacer nada
	}

	// Iniciar el contenedor utilizando docker-compose
	cmd := exec.Command("docker", "compose", "up", "-d", "postgresql-db")
	cmd.Dir = "./docker/postgresql"

	output, err := cmd.CombinedOutput()
	if err != nil {
		return models.PostgreSQLMessage{}, err
	}

	if strings.Contains(string(output), "done") {
		return models.PostgreSQLMessage{Message: "PostgreSQL has been started successfully"}, nil
	}

	return models.PostgreSQLMessage{}, nil
}

// StopPostgreSQL detiene el contenedor 'postgresql-db' utilizando docker-compose si está en ejecución
func StopPostgreSQL() (models.PostgreSQLMessage, error) {
	// Verificar si el contenedor está en ejecución
	status, err := GetStatus()
	if err != nil {
		return models.PostgreSQLMessage{}, err
	}

	if !status.Active {
		return models.PostgreSQLMessage{Message: "PostgreSQL is not running"}, nil // El contenedor no está en ejecución, no hacer nada
	}

	// Detener el contenedor utilizando docker-compose
	cmd := exec.Command("docker", "compose", "down")
	cmd.Dir = "./docker/postgresql"

	output, err := cmd.CombinedOutput()
	if err != nil {
		return models.PostgreSQLMessage{}, err
	}

	if strings.Contains(string(output), "done") {
		return models.PostgreSQLMessage{Message: "PostgreSQL has been stopped successfully"}, nil
	}

	return models.PostgreSQLMessage{}, nil
}