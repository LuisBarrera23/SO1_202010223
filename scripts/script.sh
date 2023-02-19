#!/bin/bash

ruta="/registros/archivo.txt"

echo "---------------------------------------------------------"
echo "Reporte 1:"
wc -l $ruta | awk '{print "La cantidad de logs registrados es: "$1}'
echo "---------------------------------------------------------"

errores=$(grep -o "error" $ruta | wc -l)
echo "---------------------------------------------------------"
echo "Reporte 2:"
echo "La cantidad total de operaciones que resultaron en error fueron: $errores"
echo "---------------------------------------------------------"

Sumas=$(grep -o "suma" $ruta | wc -l)
Restas=$(grep -o "resta" $ruta | wc -l)
Multiplicacion=$(grep -o "multiplicacion" $ruta | wc -l)
Division=$(grep -o "division" $ruta | wc -l)
echo "---------------------------------------------------------"
echo "Reporte 3:"
echo "La cantidad total de cada para cada tipo de operacion:"
echo "Sumas: $Sumas"
echo "Restas: $Restas"
echo "Multiplicaciones: $Multiplicacion"
echo "Divisiones: $Division"
echo "---------------------------------------------------------"

echo "---------------------------------------------------------"
echo "Reporte 4:"
fecha=$(date +"%d/%m/%Y")
echo "Los Logs realizados el dia de hoy $fecha son:"

while read linea; do
if echo "$linea" | grep -q $fecha; then
echo "$linea"
fi
done < $ruta
echo "---------------------------------------------------------"

