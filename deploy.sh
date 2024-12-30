#!/usr/bin/env zsh

# Solicitar la contraseña una sola vez
echo -n "Ingrese la contraseña SSH: "
stty -echo
read SSHPASS
stty echo
echo

# Dirección IP y usuario del Raspberry Pi
RPI_USER="pi"
RPI_HOST="192.178.0.114"
RPI_DIR="/home/pi/raspberry_controller"

# Detener el daemon raspberry-controller
echo "Deteniendo el daemon raspberry-controller..."
sshpass -p "$SSHPASS" ssh $RPI_USER@$RPI_HOST 'sudo systemctl stop raspberry-controller' && \
echo "Daemon detenido." && \

# Hacer pull de la rama master
echo "Actualizando el repositorio..."
sshpass -p "$SSHPASS" ssh $RPI_USER@$RPI_HOST "cd $RPI_DIR && git pull origin master" && \
echo "Repositorio actualizado." && \

# Reiniciar el servicio
echo "Reiniciando el servicio raspberry-controller..."
sshpass -p "$SSHPASS" ssh $RPI_USER@$RPI_HOST 'sudo systemctl start raspberry-controller' && \
echo "Servicio reiniciado."