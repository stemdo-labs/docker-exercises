# Ejercicio 2 - Creación de imagenes y despliegue de contenedor
## Objetivos
- Familiarse con comandos basicos de docker.
- Preparación básica de un dokerfile.
- Creación de una imagen de docker valida.
- Desplegar un contenedor.

## Consideraciones
 1. En la carpeta `soluciones` se creara una carpeta con el siguiente formato  `<vuestro nombre>-Ejercicio-2`.
 1. En esa carpeta se dejara el dockerfile creado y en un documento txt llamado `comandos.txt` con los comandos realiazado con sus salidas por pantalla. Ejemplo:
 ``` 
 > docker ps
 CONTAINER ID   IMAGE                                                  COMMAND                  CREATED        STATUS        PORTS                  NAMES
 1d5f6384f902   harbor.thelmalouise.ddns.net/docker_test/nginx:0.0.1   "/docker-entrypoint.…"   23 hours ago   Up 23 hours   0.0.0.0:8080->80/tcp   mi_nginx
 
 ```
 3. El ejercicio consta de dos proyectos por separado escritos en nodejs, uno actua de backend y el otro de fronend.
    
    1. ./datos/ui-web  -> Fronend
    1. ./datos/api-web -> Backend
    
    Dentro ira el dokerfile correspondiente, ajustado a las necesidades de cada proyecto. Os lo teneis que traer en local y generar el dockerfile.
    >[!TIP] Instrucciones para poner en marcha el proyecto de ui-web
    ```
    > yarn
    > yarn dev --host
     ```
    >[!TIP] Instrucciones para poner en marcha el proyecto de api-web
    ```
    > npm install
    > node index.js
     ```

## Tarea
 1. Crear los fichero dockerfile en cada proyecto.
 2. Ejecutar los comandos necesarios para generar las imagens.
 Realizar el retag de la imagenes para que sea con la siguiente nomenclatura `<vuestro nombre>/<nombre proyecto>:0.0.1`
 1. Sacar una captura de pantalla del listado de imagenes, con el objetivo de comprobar que la imagen se ha creado y nombrado de forma correcta.
 2. Desplegar 2 contenedores, uno el front y el otro el back.(!!ojo con los puertos )
 3. Listar mediante comandos los contenedores existentes.
 4. Usar el navegador para verificar el funcionamiento.
 5. Eliminar los contenedores.
 6. Eliminar las imagenes.

