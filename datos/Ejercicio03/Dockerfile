# Usa la imagen de NGINX como base
FROM nginx:latest

# Copia tu configuración personalizada de NGINX al contenedor
COPY nginx.conf /etc/nginx/nginx.conf

# Expone el puerto 80 para que pueda ser accesible desde fuera del contenedor
EXPOSE 80

# Comando para iniciar NGINX cuando se ejecute el contenedor
CMD ["nginx", "-g", "daemon off;"]
