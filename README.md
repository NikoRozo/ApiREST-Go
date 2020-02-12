# apirest

API Rest en GO, este servicio permite objener los cliente dependiendo del Score cada 5 min cambia el cliente.

## Run Locally

Para Ejecutar de manera local

```
go run .
```

## Access the app 

El Endpoints a la app es `localhost:8080/customer/{endpoint}`

```text
Get Customer by Id
"/id/{id}" 

Get Customer Score
"/byScore"
Cada 5min varia el cliente

La base de cliente se encuentra en la carpeta assets/customer.csv