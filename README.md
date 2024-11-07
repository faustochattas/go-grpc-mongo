# Go gRPC Project with MongoDB

This project implements a gRPC service in Go connected to MongoDB. Below are the steps needed to set up the local development environment using Docker and execute gRPC calls.

---

## Prerequisites

1. **Install Docker and Docker Desktop**  
   You can follow the steps [here](https://docs.docker.com/get-docker/) to install Docker.
   Having trouble with Docker? [Install it this way](#having-trouble-with-docker-install-it-this-way)

2. **Install gRPCurl**  
   Tool to make calls to gRPC services. You can install it from [grpc.io](https://github.com/fullstorydev/grpcurl).

3. **Clone the repository**  
   Clone this repository to your local machine.

---

## Setup and Execution

### Step 1: Check Ports

Make sure that ports 50051 (for gRPC) and 27017 (for MongoDB) are available, as they will be used in this project.

- **Ports used**:
  - `0.0.0.0:50051->50051/tcp`: gRPC Server.
  - `0.0.0.0:27017->27017/tcp`: MongoDB.

To check the ports currently in use, run:

```bash
docker ps
```

### Step 2: Run the Project with Docker

1. **Open Docker Desktop** make sure Docker is running.
2. **Build and start the container**  
   From the project root, run:

```bash
   docker compose up --build
```

Note: The first run may take a few minutes.

Open a new terminal and run:

```bash
docker ps
```

This allows you to confirm that the containers are running on the correct ports.

### Step 3: Connect to MongoDB

To interact directly with MongoDB:

1. Run:

```bash
   docker exec -it <mongo_container_name> mongosh
```

Example: docker exec -it go-grpc-mongo-mongodb-1 mongosh

Once inside the MongoDB CLI, select the database:

```bash
   use argentina_office
```

Create the database and populate it using the following script:

Copy and paste the script into MongoDB to initialize the database locally.

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

### Database Visualization

You can now view your database. Use the following commands in the MongoDB CLI:

```javascript
show collections
db.personas.find().pretty()
db.tickets.find().pretty()
db.proyectos.find().pretty()
```

If you prefer a more organized visualization, use the MongoDB extension in Visual Studio Code.

### Step 4:

1. Open another terminal.
2. Make any gRPC calls you need.

---

### gRPC Calls

Show all projects

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.PersonasService/GetProyectos
```

Show all tickets

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.PersonasService/GetTickets
```

Show all people

```bash
grpcurl -plaintext -d '{}' localhost:50051 pb.PersonasService/GetPersonas
```

Show people within the specified age range

```bash
grpcurl -plaintext -d '{"edadMinima": 20, "edadMaxima": 30}' localhost:50051 pb.PersonasService/GetPersonasByAgeRange
```

Show the person who has the specified ticket

```bash
grpcurl -plaintext -d '{"ticket_numero": 108}' localhost:50051 pb.PersonasService/GetPersonasPorNumeroDeTicket
```

Show the person with the specified name

```bash
grpcurl -plaintext -d '{"nombre": "Juan"}' localhost:50051 pb.PersonasService/GetPersonaByNombre
```

Show the ticket with the specified ticket number

```bash
grpcurl -plaintext -d '{"ticket_numero": 113}' localhost:50051 pb.PersonasService/GetTicketPorNumero
```

Show tickets belonging to the specified owner

```bash
grpcurl -plaintext -d '{"dueno": "Carlos"}' localhost:50051 pb.PersonasService/GetTicketPorDueno
```

Show the project in which the specified collaborator works

```bash
grpcurl -plaintext -d '{"colaborador": "Ricardo"}' localhost:50051 pb.PersonasService/GetProyectoPorColaborador
```

Show all collaborators of the specified project

```bash
grpcurl -plaintext -d '{"nombre_proyecto": "proyecto delta"}' localhost:50051 pb.PersonasService/GetColaboradoresPorProyecto
```

Possible gRPC services: List services

```bash
grpcurl -plaintext localhost:50051 list
```

List methods of a specific service

```bash
grpcurl -plaintext localhost:50051 list pb.(ServiceName)
```

---

## CRUD CALLS

(Create, update and delete the db, personas, tickets y proyectos)
Edit the data as desired,  
Examples:

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
"id": "<ID_PERSONA>",
"nombre": "Fausto Chattas Updated",
"edad": 22,
"tickets": [200, 204],
"proyecto": "proyecto actualizado"
}' localhost:50051 pb.CreateService/UpdatePersona
```

#### DELETE PERSONA

```bash
grpcurl -plaintext -d '{
"id": "<ID_PERSONA>"
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
"id": "<ID_TICKET>", // Reemplaza con el ID del ticket
"ticket_numero": 302,
"owner": "Fausto Chattas"
}' localhost:50051 pb.CreateService/UpdateTicket
```

#### DELETE TICKET

```bash
grpcurl -plaintext -d '{
"id": "<ID_TICKET>" // Reemplaza con el ID del ticket
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
"id": "<ID_PROYECTO>",
"nombre": "Proyecto Actualizado",
"colaboradores": ["Fausto", "Octavio"],
"nivel_dificultad": "alto"
}' localhost:50051 pb.CreateService/UpdateProyecto
```

#### DELETE PROYECT

```bash
grpcurl -plaintext -d '{
"id": "<ID_PROYECTO>"
}' localhost:50051 pb.CreateService/DeleteProyecto
```

—-------------------------------

### HAVING TROUBLE WITH DOCKER? INSTALL IT THIS WAY

Install from the command line
After downloading Docker.dmg from either the download buttons at the top of the page or from the release notes, run the following commands in a terminal to install Docker Desktop in the Applications folder:

```bash
$ sudo hdiutil attach Docker.dmg

$ sudo /Volumes/Docker/Docker.app/Contents/MacOS/install

$ sudo hdiutil detach /Volumes/Docker
```

By default, Docker Desktop is installed at /Applications/Docker.app. As macOS typically performs security checks the first time an application is used, the install command can take several minutes to run.
