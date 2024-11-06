package main

import (
	"context"
	"log"
	"net"

	"go-grpc-mongo/db" // Importa el paquete db
	pb "go-grpc-mongo/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var client *mongo.Client

type server struct {
	pb.UnimplementedPersonasServiceServer
	pb.UnimplementedCreateServiceServer
}

// GetPersonas - Maneja la solicitud para obtener todas las personas
func (s *server) GetPersonas(ctx context.Context, req *pb.GetPersonasRequest) (*pb.GetPersonasResponse, error) {
	log.Println("Iniciando la consulta para obtener todas las personas.")
	collection := client.Database("argentina_office").Collection("personas")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error al conectar a la base de datos para obtener personas: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var resultado []*pb.Persona
	for cursor.Next(ctx) {
		var persona struct {
			ID       string  `bson:"_id"`
			Nombre   string  `bson:"nombre"`
			Edad     int32   `bson:"edad"`
			Tickets  []int32 `bson:"tickets"`
			Proyecto string  `bson:"proyecto"`
		}
		if err := cursor.Decode(&persona); err != nil {
			log.Printf("Error al decodificar persona: %v", err)
			return nil, err
		}

		resultado = append(resultado, &pb.Persona{
			Id:       persona.ID,
			Nombre:   persona.Nombre,
			Edad:     persona.Edad,
			Tickets:  persona.Tickets,
			Proyecto: persona.Proyecto,
		})
		log.Printf("Persona encontrada: ID=%s, Nombre=%s, Edad=%d", persona.ID, persona.Nombre, persona.Edad)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error en el cursor al iterar sobre personas: %v", err)
		return nil, err
	}

	log.Println("Consulta completa. Enviando lista de personas.")
	return &pb.GetPersonasResponse{Personas: resultado}, nil
}

// GetTickets - Maneja la solicitud para obtener todos los tickets
func (s *server) GetTickets(ctx context.Context, req *pb.GetTicketsRequest) (*pb.GetTicketsResponse, error) {
	log.Println("Iniciando la consulta para obtener todos los tickets.")
	collection := client.Database("argentina_office").Collection("tickets")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error al conectar a la base de datos para obtener tickets: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var resultado []*pb.Ticket
	for cursor.Next(ctx) {
		var ticket struct {
			ID           string `bson:"_id"`
			TicketNumero int32  `bson:"ticket_numero"`
			Owner        string `bson:"owner"`
		}
		if err := cursor.Decode(&ticket); err != nil {
			log.Printf("Error al decodificar ticket: %v", err)
			return nil, err
		}

		resultado = append(resultado, &pb.Ticket{
			Id:           ticket.ID,
			TicketNumero: ticket.TicketNumero,
			Owner:        ticket.Owner,
		})
		log.Printf("Ticket encontrado: ID=%s, Número=%d, Propietario=%s", ticket.ID, ticket.TicketNumero, ticket.Owner)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error en el cursor al iterar sobre tickets: %v", err)
		return nil, err
	}

	log.Println("Consulta completa. Enviando lista de tickets.")
	return &pb.GetTicketsResponse{Tickets: resultado}, nil
}

// GetProyectos - Maneja la solicitud para obtener todos los proyectos
func (s *server) GetProyectos(ctx context.Context, req *pb.GetProyectosRequest) (*pb.GetProyectosResponse, error) {
	log.Println("Iniciando la consulta para obtener todos los proyectos.")
	collection := client.Database("argentina_office").Collection("proyectos")

	var proyectos []bson.M
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error al conectar a la base de datos para obtener proyectos: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &proyectos); err != nil {
		log.Fatalf("Error al decodificar proyectos como bson.M: %v", err)
	}
	log.Printf("Total de proyectos encontrados: %d", len(proyectos))

	var resultado []*pb.Proyecto
	for _, proyecto := range proyectos {
		resultado = append(resultado, &pb.Proyecto{
			Id:              proyecto["_id"].(primitive.ObjectID).Hex(),
			Nombre:          proyecto["nombre"].(string),
			Colaboradores:   convertToStringArray(proyecto["colaboradores"]),
			NivelDificultad: proyecto["nivel_dificultad"].(string),
		})
		log.Printf("Proyecto encontrado: ID=%s, Nombre=%s, Dificultad=%s", resultado[len(resultado)-1].Id, resultado[len(resultado)-1].Nombre, resultado[len(resultado)-1].NivelDificultad)
	}

	log.Println("Consulta completa. Enviando lista de proyectos.")
	return &pb.GetProyectosResponse{Proyectos: resultado}, nil
}

// convertToStringArray - Convierte la interfaz de MongoDB a []string
func convertToStringArray(data interface{}) []string {
	array, ok := data.([]interface{})
	if !ok {
		return []string{}
	}
	result := make([]string, len(array))
	for i, v := range array {
		result[i] = v.(string)
	}
	return result
}

// Ejemplo de otro método con logs detallados
func (s *server) GetPersonaByNombre(ctx context.Context, req *pb.GetPersonaByNombreRequest) (*pb.PersonaResponse, error) {
	log.Printf("Iniciando la consulta para obtener persona por nombre: %s.", req.Nombre)
	collection := client.Database("argentina_office").Collection("personas")

	filter := bson.M{"nombre": req.Nombre}
	var persona struct {
		ID       string  `bson:"_id"`
		Nombre   string  `bson:"nombre"`
		Edad     int32   `bson:"edad"`
		Tickets  []int32 `bson:"tickets"`
		Proyecto string  `bson:"proyecto"`
	}

	err := collection.FindOne(ctx, filter).Decode(&persona)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No se encontró persona con nombre: %s", req.Nombre)
			return nil, status.Errorf(codes.NotFound, "Persona con nombre %s no encontrada", req.Nombre)
		}
		log.Printf("Error al buscar persona: %v", err)
		return nil, err
	}

	log.Printf("Persona encontrada: ID=%s, Nombre=%s, Edad=%d", persona.ID, persona.Nombre, persona.Edad)
	return &pb.PersonaResponse{
		Persona: &pb.Persona{
			Id:       persona.ID,
			Nombre:   persona.Nombre,
			Edad:     persona.Edad,
			Tickets:  persona.Tickets,
			Proyecto: persona.Proyecto,
		},
	}, nil
}

// Obtiene personas por rango de edades
func (s *server) GetPersonasByAgeRange(ctx context.Context, req *pb.GetPersonasByAgeRangeRequest) (*pb.GetPersonasResponse, error) {
	log.Printf("Buscando personas en el rango de edad: %d - %d", req.EdadMinima, req.EdadMaxima)
	collection := client.Database("argentina_office").Collection("personas")

	filter := bson.M{
		"edad": bson.M{
			"$gte": req.EdadMinima,
			"$lte": req.EdadMaxima,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error al obtener personas por rango de edad: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var personas []*pb.Persona
	for cursor.Next(ctx) {
		var persona struct {
			ID       primitive.ObjectID `bson:"_id"`
			Nombre   string             `bson:"nombre"`
			Edad     int32              `bson:"edad"`
			Tickets  []int32            `bson:"tickets"`
			Proyecto string             `bson:"proyecto"`
		}

		if err := cursor.Decode(&persona); err != nil {
			log.Printf("Error al decodificar persona: %v", err)
			return nil, err
		}

		personas = append(personas, &pb.Persona{
			Id:       persona.ID.Hex(),
			Nombre:   persona.Nombre,
			Edad:     persona.Edad,
			Tickets:  persona.Tickets,
			Proyecto: persona.Proyecto,
		})
	}
	return &pb.GetPersonasResponse{Personas: personas}, nil
}

// Obtiene personas por número de ticket
func (s *server) GetPersonasPorNumeroDeTicket(ctx context.Context, req *pb.GetPersonasPorNumeroDeTicketRequest) (*pb.GetPersonasResponse, error) {
	log.Printf("Buscando personas con ticket número: %d", req.TicketNumero)
	collection := client.Database("argentina_office").Collection("personas")

	filter := bson.M{"tickets": req.TicketNumero}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error al obtener personas por número de ticket: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var personas []*pb.Persona
	for cursor.Next(ctx) {
		var persona struct {
			ID       primitive.ObjectID `bson:"_id"`
			Nombre   string             `bson:"nombre"`
			Edad     int32              `bson:"edad"`
			Tickets  []int32            `bson:"tickets"`
			Proyecto string             `bson:"proyecto"`
		}

		if err := cursor.Decode(&persona); err != nil {
			log.Printf("Error al decodificar persona: %v", err)
			return nil, err
		}

		personas = append(personas, &pb.Persona{
			Id:       persona.ID.Hex(),
			Nombre:   persona.Nombre,
			Edad:     persona.Edad,
			Tickets:  persona.Tickets,
			Proyecto: persona.Proyecto,
		})
	}
	return &pb.GetPersonasResponse{Personas: personas}, nil
}

// Obtiene un ticket por número de ticket
func (s *server) GetTicketPorNumero(ctx context.Context, req *pb.GetTicketPorNumeroRequest) (*pb.TicketResponse, error) {
	log.Printf("Buscando ticket con número: %d", req.TicketNumero)
	collection := client.Database("argentina_office").Collection("tickets")

	var ticket struct {
		ID           primitive.ObjectID `bson:"_id"`
		TicketNumero int32              `bson:"ticket_numero"`
		Owner        string             `bson:"owner"`
	}
	err := collection.FindOne(ctx, bson.M{"ticket_numero": req.TicketNumero}).Decode(&ticket)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "Ticket no encontrado")
		}
		log.Printf("Error al buscar el ticket: %v", err)
		return nil, err
	}

	return &pb.TicketResponse{
		Ticket: &pb.Ticket{
			Id:           ticket.ID.Hex(),
			TicketNumero: ticket.TicketNumero,
			Owner:        ticket.Owner,
		},
	}, nil
}

// Obtiene un ticket por nombre del dueño
func (s *server) GetTicketPorDueno(ctx context.Context, req *pb.GetTicketPorDuenoRequest) (*pb.TicketResponse, error) {
	log.Printf("Buscando ticket para el dueño: %s", req.Dueno)
	collection := client.Database("argentina_office").Collection("tickets")

	var ticket struct {
		ID           primitive.ObjectID `bson:"_id"`
		TicketNumero int32              `bson:"ticket_numero"`
		Owner        string             `bson:"owner"`
	}
	err := collection.FindOne(ctx, bson.M{"owner": req.Dueno}).Decode(&ticket)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "Ticket no encontrado")
		}
		log.Printf("Error al buscar el ticket: %v", err)
		return nil, err
	}

	return &pb.TicketResponse{
		Ticket: &pb.Ticket{
			Id:           ticket.ID.Hex(),
			TicketNumero: ticket.TicketNumero,
			Owner:        ticket.Owner,
		},
	}, nil
}

func (s *server) GetProyectoPorColaborador(ctx context.Context, req *pb.GetProyectoPorColaboradorRequest) (*pb.ProyectoResponse, error) {
	log.Printf("Buscando proyecto con el colaborador: %s", req.Colaborador)
	collection := client.Database("argentina_office").Collection("proyectos")

	// Crear un filtro para buscar proyectos que contengan al colaborador en la lista
	filter := bson.M{"colaboradores": req.Colaborador}

	// Definir una estructura temporal para decodificar el proyecto
	var proyecto struct {
		ID              primitive.ObjectID `bson:"_id"`
		Nombre          string             `bson:"nombre"`
		Colaboradores   []string           `bson:"colaboradores"`
		NivelDificultad string             `bson:"nivel_dificultad"`
	}

	// Buscar un único documento que coincida con el filtro
	err := collection.FindOne(ctx, filter).Decode(&proyecto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No se encontró ningún proyecto para el colaborador: %s", req.Colaborador)
			return nil, status.Errorf(codes.NotFound, "No se encontró ningún proyecto para el colaborador %s", req.Colaborador)
		}
		log.Printf("Error al buscar el proyecto: %v", err)
		return nil, err
	}

	// Retornar el proyecto encontrado
	return &pb.ProyectoResponse{
		Proyecto: &pb.Proyecto{
			Id:              proyecto.ID.Hex(),
			Nombre:          proyecto.Nombre,
			Colaboradores:   proyecto.Colaboradores,
			NivelDificultad: proyecto.NivelDificultad,
		},
	}, nil
}

func (s *server) GetColaboradoresPorProyecto(ctx context.Context, req *pb.GetColaboradoresPorProyectoRequest) (*pb.GetColaboradoresPorProyectoResponse, error) {
	log.Printf("Buscando colaboradores para el proyecto: %s", req.NombreProyecto)
	collection := client.Database("argentina_office").Collection("proyectos")

	// Filtro para encontrar el proyecto por nombre
	filter := bson.M{"nombre": req.NombreProyecto}
	var proyecto struct {
		Colaboradores []string `bson:"colaboradores"`
	}

	// Buscar el proyecto
	err := collection.FindOne(ctx, filter).Decode(&proyecto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No se encontró el proyecto: %s", req.NombreProyecto)
			return nil, status.Errorf(codes.NotFound, "No se encontró el proyecto %s", req.NombreProyecto)
		}
		log.Printf("Error al buscar el proyecto: %v", err)
		return nil, err
	}

	// Retornar la lista de colaboradores
	log.Printf("Colaboradores encontrados para el proyecto %s: %v", req.NombreProyecto, proyecto.Colaboradores)
	return &pb.GetColaboradoresPorProyectoResponse{
		Colaboradores: proyecto.Colaboradores,
	}, nil
}

func (s *server) CreatePersona(ctx context.Context, req *pb.CreatePersonaRequest) (*pb.CreatePersonaResponse, error) {
	log.Printf("Creando persona: Nombre=%s, Edad=%d", req.Nombre, req.Edad)
	collection := client.Database("argentina_office").Collection("personas")

	persona := bson.M{
		"nombre":   req.Nombre,
		"edad":     req.Edad,
		"tickets":  req.Tickets,
		"proyecto": req.Proyecto,
	}

	result, err := collection.InsertOne(ctx, persona)
	if err != nil {
		log.Printf("Error al crear persona: %v", err)
		return nil, status.Errorf(codes.Internal, "No se pudo crear la persona")
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("Persona creada con ID: %s", id)
	return &pb.CreatePersonaResponse{Id: id}, nil
}

func (s *server) UpdatePersona(ctx context.Context, req *pb.UpdatePersonaRequest) (*pb.UpdatePersonaResponse, error) {
	log.Printf("Actualizando persona con ID: %s", req.Id)
	collection := client.Database("argentina_office").Collection("personas")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID inválido")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"nombre":   req.Nombre,
		"edad":     req.Edad,
		"tickets":  req.Tickets,
		"proyecto": req.Proyecto,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		log.Printf("Error al actualizar persona: %v", err)
		return nil, status.Errorf(codes.NotFound, "Persona no encontrada")
	}

	log.Printf("Persona actualizada correctamente")
	return &pb.UpdatePersonaResponse{Success: true}, nil
}

func (s *server) DeletePersona(ctx context.Context, req *pb.DeletePersonaRequest) (*pb.DeletePersonaResponse, error) {
	log.Printf("Eliminando persona con ID: %s", req.Id)
	collection := client.Database("argentina_office").Collection("personas")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID inválido")
	}

	filter := bson.M{"_id": objID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		log.Printf("Error al eliminar persona: %v", err)
		return nil, status.Errorf(codes.NotFound, "Persona no encontrada")
	}

	log.Printf("Persona eliminada correctamente")
	return &pb.DeletePersonaResponse{Success: true}, nil
}

// Método para crear un nuevo ticket
func (s *server) CreateTicket(ctx context.Context, req *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	log.Printf("Creando un nuevo ticket: Número=%d, Owner=%s", req.TicketNumero, req.Owner)

	collection := client.Database("argentina_office").Collection("tickets")

	newTicket := bson.M{
		"ticket_numero": req.TicketNumero,
		"owner":         req.Owner,
	}

	res, err := collection.InsertOne(ctx, newTicket)
	if err != nil {
		log.Printf("Error al crear el ticket: %v", err)
		return nil, status.Error(codes.Internal, "Error al crear el ticket")
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("Ticket creado con ID: %s", id)

	return &pb.CreateTicketResponse{Id: id}, nil
}

// Método para actualizar un ticket
func (s *server) UpdateTicket(ctx context.Context, req *pb.UpdateTicketRequest) (*emptypb.Empty, error) {
	log.Printf("Actualizando ticket con ID=%s", req.Id)

	collection := client.Database("argentina_office").Collection("tickets")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("ID de ticket no válido: %v", err)
		return nil, status.Error(codes.InvalidArgument, "ID de ticket no válido")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"ticket_numero": req.TicketNumero,
			"owner":         req.Owner,
		},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error al actualizar el ticket: %v", err)
		return nil, status.Error(codes.Internal, "Error al actualizar el ticket")
	}

	if res.MatchedCount == 0 {
		log.Printf("Ticket no encontrado: %s", req.Id)
		return nil, status.Error(codes.NotFound, "Ticket no encontrado")
	}

	log.Printf("Ticket actualizado con éxito: ID=%s", req.Id)
	return &emptypb.Empty{}, nil
}

// Método para eliminar un ticket
func (s *server) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*emptypb.Empty, error) {
	log.Printf("Eliminando ticket con ID=%s", req.Id)

	collection := client.Database("argentina_office").Collection("tickets")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("ID de ticket no válido: %v", err)
		return nil, status.Error(codes.InvalidArgument, "ID de ticket no válido")
	}

	filter := bson.M{"_id": objID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error al eliminar el ticket: %v", err)
		return nil, status.Error(codes.Internal, "Error al eliminar el ticket")
	}

	if res.DeletedCount == 0 {
		log.Printf("Ticket no encontrado: %s", req.Id)
		return nil, status.Error(codes.NotFound, "Ticket no encontrado")
	}

	log.Printf("Ticket eliminado con éxito: ID=%s", req.Id)
	return &emptypb.Empty{}, nil
}

// Método para crear un proyecto
func (s *server) CreateProyecto(ctx context.Context, req *pb.CreateProyectoRequest) (*pb.CreateProyectoResponse, error) {
	log.Printf("Creando proyecto: Nombre=%s, Dificultad=%s", req.Nombre, req.NivelDificultad)

	collection := client.Database("argentina_office").Collection("proyectos")
	proyecto := bson.M{
		"nombre":           req.Nombre,
		"colaboradores":    req.Colaboradores,
		"nivel_dificultad": req.NivelDificultad,
	}

	res, err := collection.InsertOne(ctx, proyecto)
	if err != nil {
		log.Printf("Error al crear el proyecto: %v", err)
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("Proyecto creado con ID: %s", id)

	return &pb.CreateProyectoResponse{Id: id}, nil
}

// Método para actualizar un proyecto
func (s *server) UpdateProyecto(ctx context.Context, req *pb.UpdateProyectoRequest) (*emptypb.Empty, error) {
	log.Printf("Actualizando proyecto con ID: %s", req.Id)

	collection := client.Database("argentina_office").Collection("proyectos")
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("ID de proyecto inválido: %v", err)
		return nil, status.Error(codes.InvalidArgument, "ID de proyecto inválido")
	}

	update := bson.M{
		"$set": bson.M{
			"nombre":           req.Nombre,
			"colaboradores":    req.Colaboradores,
			"nivel_dificultad": req.NivelDificultad,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Error al actualizar el proyecto: %v", err)
		return nil, err
	}

	log.Printf("Proyecto actualizado con ID: %s", req.Id)
	return &emptypb.Empty{}, nil
}

// Método para eliminar un proyecto
func (s *server) DeleteProyecto(ctx context.Context, req *pb.DeleteProyectoRequest) (*emptypb.Empty, error) {
	log.Printf("Eliminando proyecto con ID: %s", req.Id)

	collection := client.Database("argentina_office").Collection("proyectos")
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("ID de proyecto inválido: %v", err)
		return nil, status.Error(codes.InvalidArgument, "ID de proyecto inválido")
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Printf("Error al eliminar el proyecto: %v", err)
		return nil, err
	}

	log.Printf("Proyecto eliminado con ID: %s", req.Id)
	return &emptypb.Empty{}, nil
}

func main() {
	var err error
	client, err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonasServiceServer(s, &server{})
	pb.RegisterCreateServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Servidor en ejecución en el puerto 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servicio: %v", err)
	}

}
