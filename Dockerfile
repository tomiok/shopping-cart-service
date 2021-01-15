##la imagen de golang en docker
FROM golang

## creamos un directorio
RUN mkdir /app
## copiamos el volumen al directorio
ADD . /app
## seleccionamos app como directorio de trabajo
WORKDIR /app

## corremos el comando build del makefile
RUN make build

## corremos el binario
CMD ./shopping-service

## exponemos el puerto 3000
EXPOSE 3000