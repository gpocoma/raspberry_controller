#!/usr/bin/env zsh

# Solicitar la contraseña una sola vez
echo -n "Ingrese la contraseña SSH: "
stty -echo
read SSHPASS
stty echo
echo

# Detener el daemon raspberry-controller
echo "Deteniendo el daemon raspberry-controller..."
sshpass -p "$SSHPASS" ssh pi@192.178.0.114 'sudo systemctl stop raspberry-controller'
echo "Daemon detenido."

# Copiar el nuevo controlador
echo "Copiando el nuevo controlador..."
sshpass -p "$SSHPASS" scp build/raspberry-controller pi@192.178.0.114:/home/pi/raspberry-controller
echo "Controlador copiado."

# Reiniciar el servicio
echo "Reiniciando el servicio raspberry-controller..."
sshpass -p "$SSHPASS" ssh pi@192.178.0.114 'sudo systemctl start raspberry-controller'
echo "Servicio reiniciado."