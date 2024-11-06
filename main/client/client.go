package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-grpc-mongo/proto"

	"google.golang.org/grpc"
)

// Función principal para ejecutar el cliente gRPC
func runClient() {
	// Conexión al servidor gRPC en localhost y puerto 50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Puse localhost porq toy corriendo localmente
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor gRPC: %v", err) // Manejo de errores en la conexión
	}
	defer conn.Close() // Asegura que la conexión se cierre al finalizar la función

	client := pb.NewPersonasServiceClient(conn)
	createClient := pb.NewCreateServiceClient(conn) // Creo nuevo cliente para el servicio de personas

	// Solicitud para obtener todas las personas
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // Establece un contexto con timeout de 1 segundo
	defer cancel()                                                        // Cancelar el contexto al final de la función

	// Llamada al método GetPersonas del servidor
	resp, err := client.GetPersonas(ctx, &pb.GetPersonasRequest{}) // Uso del alias 'pb' para la solicitud
	if err != nil {
		log.Fatalf("Error al obtener personas: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión de las personas obtenidas
	fmt.Println("Personas:")
	for _, p := range resp.Personas {
		fmt.Printf("ID: %s, Nombre: %s, Edad: %d\n", p.Id, p.Nombre, p.Edad) // Formato de salida para cada persona
	}

	// Solicitud para obtener todos los tickets
	ticketsResp, err := client.GetTickets(ctx, &pb.GetTicketsRequest{}) // Uso del alias 'pb' para la solicitud
	if err != nil {
		log.Fatalf("Error al obtener tickets: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión de los tickets obtenidos
	fmt.Println("Tickets:")
	for _, t := range ticketsResp.Tickets {
		fmt.Printf("ID: %s, Número: %d, Propietario: %s\n", t.Id, t.TicketNumero, t.Owner) // Formato de salida para cada ticket
	}

	// Solicitud para obtener todos los proyectos
	proyectosResp, err := client.GetProyectos(ctx, &pb.GetProyectosRequest{}) // Uso del alias 'pb' para la solicitud
	if err != nil {
		log.Fatalf("Error al obtener proyectos: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión de los proyectos obtenidos
	fmt.Println("Proyectos:")
	for _, p := range proyectosResp.Proyectos {
		fmt.Printf("ID: %s, Nombre: %s, Dificultad: %s\n", p.Id, p.Nombre, p.NivelDificultad) // Formato de salida para cada proyecto
	}

	// Solicitud para obtener personas dentro de un rango de edades
	ageRangeResp, err := client.GetPersonasByAgeRange(ctx, &pb.GetPersonasByAgeRangeRequest{
		EdadMinima: 25, // Edad mínima de ejemplo
		EdadMaxima: 40, // Edad máxima de ejemplo
	})
	if err != nil {
		log.Fatalf("Error al obtener personas por rango de edad: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión de personas en el rango de edad especificado
	fmt.Println("Personas en el rango de edad:")
	for _, p := range ageRangeResp.Personas {
		fmt.Printf("ID: %s, Nombre: %s, Edad: %d\n", p.Id, p.Nombre, p.Edad) // Formato de salida para cada persona
	}

	// Nueva solicitud para obtener personas por número de ticket
	ticketNumero := int32(123) // Número de ticket de ejemplo
	personasPorTicketResp, err := client.GetPersonasPorNumeroDeTicket(ctx, &pb.GetPersonasPorNumeroDeTicketRequest{
		TicketNumero: ticketNumero, // Se pasa el número de ticket como parámetro
	})
	if err != nil {
		log.Fatalf("Error al obtener personas por número de ticket: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión de personas asociadas al número de ticket
	fmt.Printf("Personas con ticket número %d:\n", ticketNumero)
	for _, persona := range personasPorTicketResp.Personas {
		fmt.Printf("ID: %s, Nombre: %s, Edad: %d\n", persona.Id, persona.Nombre, persona.Edad) // Formato de salida para cada persona
	}

	// Pedir al usuario que ingrese el nombre que desea buscar
	fmt.Print("Ingrese el nombre de la persona que desea buscar: ")
	var nombre string
	fmt.Scanln(&nombre) // Lee la entrada del usuario y la asigna a la variable 'nombre'

	// Llamada al método GetPersonaByNombre con el nombre ingresado
	personaResp, err := client.GetPersonaByNombre(ctx, &pb.GetPersonaByNombreRequest{
		Nombre: nombre, // Se usa el nombre ingresado dinámicamente
	})
	if err != nil {
		log.Fatalf("Error al obtener persona por nombre: %v", err)
	}

	// Impresión de la persona encontrada
	if personaResp != nil && personaResp.Persona != nil {
		fmt.Printf("Persona encontrada:\n")
		fmt.Printf("ID: %s, Nombre: %s, Edad: %d, Proyecto: %s\n",
			personaResp.Persona.Id,
			personaResp.Persona.Nombre,
			personaResp.Persona.Edad,
			personaResp.Persona.Proyecto)
	} else {
		fmt.Printf("No se encontró una persona con el nombre %s.\n", nombre)
	}

	// Nueva solicitud para obtener un ticket por número
	ticketResp, err := client.GetTicketPorNumero(ctx, &pb.GetTicketPorNumeroRequest{
		TicketNumero: ticketNumero, // Se pasa el número de ticket como parámetro
	})
	if err != nil {
		log.Fatalf("Error al obtener ticket por número: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión del ticket obtenido
	if ticketResp != nil && ticketResp.Ticket != nil {
		fmt.Printf("Ticket encontrado:\n")
		fmt.Printf("ID: %s, Número: %d, Propietario: %s\n",
			ticketResp.Ticket.Id,
			ticketResp.Ticket.TicketNumero,
			ticketResp.Ticket.Owner) // Formato de salida
	} else {
		fmt.Printf("No se encontró un ticket con el número %d.\n", ticketNumero)
	}

	// Nueva solicitud para obtener un proyecto por colaborador
	colaboradorNombre := "Juan Perez" // Cambia esto por cualquier nombre que desees buscar
	proyectoResp, err := client.GetProyectoPorColaborador(ctx, &pb.GetProyectoPorColaboradorRequest{
		Colaborador: colaboradorNombre, // Se pasa el nombre del colaborador como parámetro
	})
	if err != nil {
		log.Fatalf("Error al obtener proyecto por colaborador: %v", err) // Manejo de errores en la solicitud
	}

	// Impresión del proyecto obtenido
	if proyectoResp != nil && proyectoResp.Proyecto != nil {
		fmt.Printf("Proyecto encontrado:\n")
		fmt.Printf("ID: %s, Nombre: %s, Dificultad: %s, Colaboradores: %v\n",
			proyectoResp.Proyecto.Id,
			proyectoResp.Proyecto.Nombre,
			proyectoResp.Proyecto.NivelDificultad,
			proyectoResp.Proyecto.Colaboradores) // Formato de salida
	} else {
		fmt.Printf("No se encontró un proyecto para el colaborador %s.\n", colaboradorNombre)
	}

	fmt.Println("Creando nueva persona...")
	createResp, err := createClient.CreatePersona(ctx, &pb.CreatePersonaRequest{
		Nombre:   "Carlos",
		Edad:     29,
		Tickets:  []int32{100, 200},
		Proyecto: "Proyecto Nuevo",
	})
	if err != nil {
		log.Fatalf("Error al crear persona: %v", err)
	}
	fmt.Printf("Persona creada con ID: %s\n", createResp.Id)

	// Ejemplo de uso de UpdatePersona
	fmt.Println("Actualizando persona...")
	updateResp, err := createClient.UpdatePersona(ctx, &pb.UpdatePersonaRequest{
		Id:       createResp.Id,
		Nombre:   "Carlos Actualizado",
		Edad:     30,
		Tickets:  []int32{101, 201},
		Proyecto: "Proyecto Actualizado",
	})
	if err != nil {
		log.Fatalf("Error al actualizar persona: %v", err)
	}
	fmt.Printf("Actualización exitosa: %v\n", updateResp.Success)

	// Ejemplo de uso de DeletePersona
	fmt.Println("Eliminando persona...")
	deleteResp, err := createClient.DeletePersona(ctx, &pb.DeletePersonaRequest{
		Id: createResp.Id,
	})
	if err != nil {
		log.Fatalf("Error al eliminar persona: %v", err)
	}
	fmt.Printf("Eliminación exitosa: %v\n", deleteResp.Success)
}

// Función principal que inicia el cliente
func main() {
	runClient() // Llama a la función que ejecuta el cliente gRPC
}
