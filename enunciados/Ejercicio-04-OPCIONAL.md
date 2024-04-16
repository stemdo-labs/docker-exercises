# Ejercicio 4 (OPCIONAL)

## Objetivos

1. Adquisición del concepto de "Multi-Stage".
1. Conocer distintas imágenes base livianas.
1. Cambiar el usuario por defecto en una imagen.

## Consideraciones

1. En la carpeta `soluciones` se creará una carpeta con el siguiente formato `<vuestro nombre>-Ejercicio-4`.
1. En esta carpeta se dejarán los Dockerfiles creados.

## Tarea

**NOTA**: Entre los ficheros que se encuentran en `/datos/Ejercicio04-opcional` se encuentra un
Dockerfile de ejemplo que se puede usar como base para la realización de este ejercicio.
La primera fase puede ser la misma que la del Dockerfile de ejemplo.

Partiendo de los ficheros golang que se encuentran en /datos/Ejercicio04-opcional,
crear dos Dockerfile usando multi-stage:

- El primero debe usar una imagen base "scratch" para correr el ejecutable generado.
- El segundo, debe usar una imagen de linux "alpine" para correr el ejecutable generado.
  En este segundo Dockerfile, es necesario cambiar el usuario por defecto a uno que no sea root y que solo tenga permisos de lectura y ejecución sobre el ejecutable. El resto del sistema debe ser de solo lectura.