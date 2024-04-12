# Ejercicio 3 - Un poco de docker-compose
## Objetivos
1. Testear un poco el funcionamiento de `docker-compose`.
2. Una pequeña introduccion a la gestion de volumnes.


## Consideraciones
 1. En la carpeta `soluciones` se creara una carpeta con el siguiente formato  `<vuestro nombre>-Ejercicio-3`.
 1. En esa carpeta se dejara el dockerfile creado y en un documento txt llamado `comandos.txt` con los comandos realiazado con sus salidas por pantalla.
 2. El ejercicio consta de un proyecto `nginx`. Tendreis que llevaros usar el codigo que se encuentra en la carpeta `./datos/Ejercicio03`. El dockerfile ya os lo proporciono yo.
## Tarea
1. Generar la imagen del proyecto.
2. Listar las imagenes y asegurarse de que lo teneis en vuestro local.
3. Generar vuestro  `docker-compose` con la capacidad de desplegar el contenedor y dar servicio (dejo a vuestra elección el puerto 😏)
4. Testear que os da servicio usando el navegador y sacar una captura de pantalla.
5. Modificar el fichero `html/index.html` para que salga vuestro nombre por ejemplo el mio quedaria asi:
``` 
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>¡Bienvenido Hector Gonazalez a mi sitio web!</title>
</head>
<body>
    <h1>¡Bienvenido Hector Gonazalez a mi sitio web desplegado con NGINX en Docker!</h1>
    <p>Este es un ejemplo básico de un sitio web desplegado utilizando NGINX en un contenedor Docker.</p>
</body>
</html>

```
una vez guardado recargad la pagina y sacad captura del resultado.

>[!TIP] Pista
La idea es que lo que renderice se encuentre en nuestro `html/index.html`. ¿Cual sera la forma para que el contenedor pueda acceder a este archivo desde nuestro equipo local?.