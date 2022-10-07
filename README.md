# Desafio DEUNA

### Descripción

**Define un endpoint** con el verbo http `GET` y con la siguiente URL: `/api/v1/starships/available?passengers=600` donde le puedas enviar por query params la cantidad de pasajeros deseada para buscar naves.

**Consume el siguiente endpoint** `https://swapi.dev/api/starships` de la API de **[star wars](https://swapi.dev/)** y responde con la cantidad de naves que aceptan sobre la cantidad de pasajeros enviada en el query params.

### Construcción 🛠️
* **Language:** Golang
* **Framework:** Fiber

## Instalación y Ejecución

Pasos:

1. Clone el proyecto.
2. Ejecutar el archivo ```main.go```
3. Consumir el servicio ```available```
   ```
   curl --location --request GET '127.0.0.1:3000/api/v1/starships/available?passengers=600'
   ```
