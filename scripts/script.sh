#!/bin/bash
read -p "Por favor ingrese su usuario de Github que desea consultar: " GITHUB_USER
jsonDatos=$(curl -s https://api.github.com/users/"$GITHUB_USER")

#echo $jsonDatos

id=$(echo $jsonDatos | jq -r '.id')
fecha=$(echo $jsonDatos | jq -r '.created_at')

echo "Hola $GITHUB_USER. User ID: $id. Cuenta fue creada el: $fecha."

Nombre_carpeta=$(date +"%d-%m-%Y")
#echo $Nombre_carpeta
mkdir -p "$(dirname "/tmp/$Nombre_carpeta/saludos.log")"
echo "Hola $GITHUB_USER. User ID: $id. Cuenta fue creada el: $fecha." >> "/tmp/$Nombre_carpeta/saludos.log"