# Universidad de San Carlos de Guatemala
# Sistemas Operativos 1

# Practica 1 Contenedores de Docker
# Luis Angel Barrera Velásquez *202010223*

## Arquitectura utilizada para la solucion
Basicamente se utilizo una aplicacion completa de frontend, backend y base de datos, incluyendo una reporteria hecha con un script de bash cada una de las antes mencionadas herramientas fueron dockerizadas para utilizarla por medio de contenedores que mas adelante sera detallado. La arquitectura utilizada es la siguiente: 

![arquitectura](./imagenes/Captura%20desde%202023-02-19%2017-48-21.png)

A continuacion se explicara cada una de las capas del proyecto:

### Frontend [Repositorio](https://github.com/LuisBarrera23/SO1_202010223/tree/main/frontend)

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/2300px-React-icon.svg.png" alt="arquitectura" width="200">

Para el frontend fue utilizado React en la version node 16.19.0 creando el proyecto con el comando: 

```
creat-react-app myapp
```