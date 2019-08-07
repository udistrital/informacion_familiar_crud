# Informacion_familiar_crud

API CRUD para la gestion de la informaciín familiar principal del aspirante  necesarias en el proceso de inscripción del sistema de gestion academica

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/informacion_familiar_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_INFORMACION_FAMILIAR_CRUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `INFORMACION_FAMILIAR_CRUD__PGUSER`: Usuario de la base de datos
 - `INFORMACION_FAMILIAR_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `INFORMACION_FAMILIAR_CRUD__PGURLS`: Host de conexión
 - `INFORMACION_FAMILIAR_CRUD__PGDB`: Nombre de la base de datos
 - `INFORMACION_FAMILIAR_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_INFORMACION_FAMILIAR_CRUD_HTTP_PORT=8080 INFORMACION_FAMILIAR_CRUD__PGUSER=postgres INFORMACION_FAMILIAR_CRUD__PGPASS=12345678 INFORMACION_FAMILIAR_CRUD__PGURLS=localhost INFORMACION_FAMILIAR_CRUD__PGDB=informacion_familiar_1 INFORMACION_FAMILIAR_CRUD__SCHEMA=informacion_familiar bee run -downdoc=true -gendoc=true

## MODELO DE DATOS

Como modelos de datos del Api se utilizo el siguiente 