# Proyecto Go gRPC con MongoDB

Este proyecto implementa un servicio gRPC en Go conectado a MongoDB. A continuación, se detallan los pasos necesarios para configurar el entorno de desarrollo local usando Docker y ejecutar las llamadas gRPC.

---

## Requisitos previos

1. **Instalar Docker y Docker Desktop**  
   Puedes seguir los pasos [aquí](https://docs.docker.com/get-docker/) para instalar Docker.
   Problemas con docker? [Instalalo asi](#HAVING TROUBLE WITH DOCKER? INSTALL IT THIS WAY)

2. **Instalar gRPCurl**  
   Herramienta para realizar llamadas a los servicios gRPC. Puedes instalarla desde [grpc.io](https://github.com/fullstorydev/grpcurl).

3. **Clonar el repositorio**  
   Clona este repositorio en tu máquina local.

---

## Configuración y Ejecución

### Paso 1: Verificar Puertos

Asegúrate de que los puertos 50051 (para gRPC) y 27017 (para MongoDB) estén disponibles, ya que se utilizarán en este proyecto.

- **Puertos utilizados**:
  - `0.0.0.0:50051->50051/tcp`: Servidor gRPC.
  - `0.0.0.0:27017->27017/tcp`: MongoDB.

Para ver los puertos que están actualmente en uso, ejecuta:

```bash
docker ps
```

### Paso 2: Ejecutar el Proyecto con Docker

1. **Abrir Docker Desktop** y asegúrate de que Docker esté ejecutándose.
2. **Construir y levantar el contenedor**
   Desde la raíz del proyecto, ejecuta:
   ```bash
   docker compose up --build
   ```
   Nota: La primera ejecución puede tardar unos minutos.

Abre una nueva terminal y ejecuta:

```bash
docker ps
```

Esto permite confirmar que los contenedores están corriendo en los puertos correctos.

### Paso 3: Conectar a MongoDB

Para interactuar directamente con MongoDB:

1. Ejecuta:
   ```bash
   docker exec -it <nombre_contenedor_mongo> mongosh
   ```
   Ejemplo: docker exec -it go-grpc-mongo-mongodb-1 mongosh

Una vez dentro del CLI de MongoDB, selecciona la base de datos:

```bash
use argentina_office
```

Crear la base de datos y poblarla usando el siguiente script:

Copia y pega el script en MongoDB para inicializar la base de datos localmente.

```javascript
use argentina_office

db.personas.insertMany([
  {
    _id: ObjectId('672510bdfd2eb362d6daf1cc'),
    nombre: 'Juan',
    edad: 30,
    tickets: [101, 102],
    proyecto: 'proyecto alpha'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1cd'),
    nombre: 'María',
    edad: 28,
    tickets: [103],
    proyecto: 'proyecto beta'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1ce'),
    nombre: 'Pedro',
    edad: 35,
    tickets: [104, 105, 106],
    proyecto: 'proyecto gamma'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1cf'),
    nombre: 'Ana',
    edad: 25,
    tickets: [],
    proyecto: 'proyecto delta'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1d0'),
    nombre: 'Luis',
    edad: 40,
    tickets: [107, 108],
    proyecto: 'proyecto alpha'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1d1'),
    nombre: 'Elena',
    edad: 33,
    tickets: [109],
    proyecto: 'proyecto beta'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1d2'),
    nombre: 'Carlos',
    edad: 29,
    tickets: [110, 111],
    proyecto: 'proyecto gamma'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1d3'),
    nombre: 'Sofía',
    edad: 26,
    tickets: [],
    proyecto: 'proyecto delta'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1d4'),
    nombre: 'Ricardo',
    edad: 37,
    tickets: [112, 113, 114],
    proyecto: 'proyecto alpha'
  },
  {
    _id: ObjectId('672510bdfd2eb362d6daf1d5'),
    nombre: 'Julia',
    edad: 31,
    tickets: [115],
    proyecto: 'proyecto beta'
  }
])

db.tickets.insertMany([
  {
    _id: ObjectId('67251340fd2eb362d6daf1d6'),
    ticket_numero: 101,
    owner: 'Juan'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1d7'),
    ticket_numero: 102,
    owner: 'Juan'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1d8'),
    ticket_numero: 103,
    owner: 'María'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1d9'),
    ticket_numero: 104,
    owner: 'Pedro'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1da'),
    ticket_numero: 105,
    owner: 'Pedro'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1db'),
    ticket_numero: 106,
    owner: 'Pedro'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1dc'),
    ticket_numero: 107,
    owner: 'Luis'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1dd'),
    ticket_numero: 108,
    owner: 'Luis'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1de'),
    ticket_numero: 109,
    owner: 'Elena'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1df'),
    ticket_numero: 110,
    owner: 'Carlos'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1e0'),
    ticket_numero: 111,
    owner: 'Carlos'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1e1'),
    ticket_numero: 112,
    owner: 'Ricardo'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1e2'),
    ticket_numero: 113,
    owner: 'Ricardo'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1e3'),
    ticket_numero: 114,
    owner: 'Ricardo'
  },
  {
    _id: ObjectId('67251340fd2eb362d6daf1e4'),
    ticket_numero: 115,
    owner: 'Julia'
  }
])

db.proyectos.insertMany([
  {
    _id: ObjectId('672513fefd2eb362d6daf1e5'),
    nombre: 'proyecto alpha',
    colaboradores: ['Juan', 'Luis', 'Ricardo'],
    nivel_dificultad: 'medio'
  },
  {
    _id: ObjectId('672513fefd2eb362d6daf1e6'),
    nombre: 'proyecto beta',
    colaboradores: ['María', 'Elena', 'Julia'],
    nivel_dificultad: 'fácil'
  },
  {
    _id: ObjectId('672513fefd2eb362d6daf1e7'),
    nombre: 'proyecto gamma',
    colaboradores: ['Pedro', 'Carlos'],
    nivel_dificultad: 'difícil'
  },
  {
    _id: ObjectId('672513fefd2eb362d6daf1e8'),
    nombre: 'proyecto delta',
    colaboradores: ['Ana', 'Sofía'],
    nivel_dificultad: 'fácil'
  }
])
```

### Visualización de la Base de Datos

Ya puedes ver tu base de datos. Usa los siguientes comandos en el CLI de MongoDB:

```javascript
show collections
db.personas.find().pretty()
db.tickets.find().pretty()
db.proyectos.find().pretty()
```

Si prefieres una visualización más organizada, utiliza la extensión MongoDB en Visual Studio Code.

### Paso 4:

1. Abre otra terminal.
2. Realiza las gRPC calls que necesites.

---

### gRPC Calls

Muestra todos los proyectos

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.PersonasService/GetProyectos
```

Muestra todos los tickets

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.PersonasService/GetTickets
```

Muestra todas las personas

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.PersonasService/GetPersonas
```

Muestra personas que estén dentro del rango de edad pedido

```bash
grpcurl -plaintext -d '{"edadMinima": 20, "edadMaxima": 30}' localhost:50051 pb.PersonasService/GetPersonasByAgeRange
```

Muestra la persona que contenga el ticket pedido

```bash
grpcurl -plaintext -d '{"ticket_numero": 108}' localhost:50051 pb.PersonasService/GetPersonasPorNumeroDeTicket
```

Muestra la persona que tenga el nombre pedido

```bash
grpcurl -plaintext -d '{"nombre": "Juan"}' localhost:50051 pb.PersonasService/GetPersonaByNombre
```

Muestra el ticket buscando por el número de ticket pedido

```bash
grpcurl -plaintext -d '{"ticket_numero": 113}' localhost:50051 pb.PersonasService/GetTicketPorNumero
```

Muestra los tickets que pertenezcan al owner pedido

```bash
grpcurl -plaintext -d '{"dueno": "Carlos"}' localhost:50051 pb.PersonasService/GetTicketPorDueno
```

Muestra el proyecto en el que el colaborador trabaja

```bash
grpcurl -plaintext -d '{"colaborador": "Ricardo"}' localhost:50051 pb.PersonasService/GetProyectoPorColaborador
```

Muestra todos los colaboradores del proyecto pedido

```bash
grpcurl -plaintext -d '{"nombre_proyecto": "proyecto delta"}' localhost:50051 pb.PersonasService/GetColaboradoresPorProyecto
```

---

Servicios posibles de gRPC: List services

```bash
grpcurl -plaintext localhost:50051 list
```

Listar métodos de un servicio en particular: List methods from a service

```bash
grpcurl -plaintext localhost:50051 list pb.(nombreDeServicio)
```

---

## CRUD CALLS

(Create, update and delete the db, personas, tickets y proyectos)
Edita sus datos a eleccion,
Ejmplos:

#### CREATE PERSONA

```bash
grpcurl -plaintext -d '{
"nombre": "Fausto Chattas",
"edad": 21,
"tickets": [201, 202],
"proyecto": "proyecto go"
}' localhost:50051 pb.CreateService/CreatePersona
```

#### UPDATE PERSONA

```bash
grpcurl -plaintext -d '{
"id": "<ID_DE_PERSONA>",
"nombre": "Fausto Chattas Updated",
"edad": 22,
"tickets": [200, 204],
"proyecto": "proyecto actualizado"
}' localhost:50051 pb.CreateService/UpdatePersona
```

#### DELETE PERSONA

```bash
grpcurl -plaintext -d '{
"id": "<ID_DE_PERSONA>"
}' localhost:50051 pb.CreateService/DeletePersona
```

—-------------------------------

#### CREATE TICKET

```bash
grpcurl -plaintext -d '{
"ticket_numero": 301,
"owner": "Octavio"
}' localhost:50051 pb.CreateService/CreateTicket
```

#### UPDATE TICKET

```bash
grpcurl -plaintext -d '{
"id": "<ID_DEL_TICKET>", // Reemplaza con el ID del ticket
"ticket_numero": 302,
"owner": "Fausto Chattas"
}' localhost:50051 pb.CreateService/UpdateTicket
```

#### DELETE TICKET

```bash
grpcurl -plaintext -d '{
"id": "<ID_DEL_TICKET>" // Reemplaza con el ID del ticket
}' localhost:50051 pb.CreateService/DeleteTicket
```

—------------------------------

#### CREATE PROYECT

```bash
grpcurl -plaintext -d '{
"nombre": "Proyecto Nueva Era",
"colaboradores": ["Fausto", "Octavio", "Lucia"],
"nivel_dificultad": "medio"
}' localhost:50051 pb.CreateService/CreateProyecto
```

#### UPDATE PROYECT

```bash
grpcurl -plaintext -d '{
"id": "<ID_DEL_PROYECTO>",
"nombre": "Proyecto Actualizado",
"colaboradores": ["Fausto", "Octavio"],
"nivel_dificultad": "alto"
}' localhost:50051 pb.CreateService/UpdateProyecto
```

#### DELETE PROYECT

```bash
grpcurl -plaintext -d '{
"id": "<ID_DEL_PROYECTO>"
}' localhost:50051 pb.CreateService/DeleteProyecto
```

—-------------------------------

## HAVING TROUBLE WITH DOCKER? INSTALL IT THIS WAY

Install from the command line
After downloading Docker.dmg from either the download buttons at the top of the page or from the release notes, run the following commands in a terminal to install Docker Desktop in the Applications folder:

$ sudo hdiutil attach Docker.dmg

$ sudo /Volumes/Docker/Docker.app/Contents/MacOS/install

$ sudo hdiutil detach /Volumes/Docker

By default, Docker Desktop is installed at /Applications/Docker.app. As macOS typically performs security checks the first time an application is used, the install command can take several minutes to run.
