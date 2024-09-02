# Prueba Técnica: Servicio de Gestión de Tareas (To-Do) en Go con Arquitectura Hexagonal

## Descripción del Proyecto

El objetivo es crear un microservicio en Go que permita crear, leer y listar tareas en una lista de tareas pendientes. El servicio debe seguir la arquitectura hexagonal y utilizar MongoDB como la base de datos para almacenar las tareas.

## Requerimientos

- Crear una API RESTful con los siguientes endpoints:
  - **POST** `/tasks`: Crear una nueva tarea.
  - **UPDATE** `/tasks/{id}`: Editar una tarea por ID.
  - **DELETE** `/tasks/{id}`: Eliminar una tarea por ID.
  - **GET** `/tasks/{id}`: Obtener una tarea por su ID.
  - **GET** `/tasks`: Listar todas las tareas.

- Las tareas deben tener la siguiente estructura:
  - **ID** (generado automáticamente)
  - **Title**: Título de la tarea.
  - **Description**: Descripción de la tarea.
  - **Completed**: Estado de la tarea (completada o no).

- Deberá estar almacenada en un contenedor de Docker, es decir, dentro del código del proyecto deberás incluir archivos de configuración de Docker.

- Manejar el avance con commits y buenas prácticas en git.

- Subirlo a un repositorio de GitHub.

## Dependencias
- Chi router
- MongoBD Driver

## Levantar el entorno - Docker (Compose)
Ejecutar el siguiente comando
```
docker-compose up
```
Para liberar el entorno ejecutar
```
docker-compose down
```

