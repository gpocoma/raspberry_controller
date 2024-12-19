#!/usr/bin/env zsh

# Crear la carpeta build si no existe
mkdir -p build

# Compilar el binario para la Raspberry Pi y guardarlo en la carpeta build
GOARCH=arm GOOS=linux go build -o build/raspberry-controller cmd/server/main.go

# Verificar si la compilación fue exitosa
if [ $? -eq 0 ]; then
  echo "Compilación exitosa. El archivo binario se encuentra en la carpeta 'build'."
else
  echo "Hubo un error durante la compilación."
fi
