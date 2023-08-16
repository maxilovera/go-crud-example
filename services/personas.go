package services

import "github.com/maxilovera/go-crud-example/dto"

var personas []*dto.Persona

func ObtenerPersonas() []*dto.Persona {
	return personas
}

func ObtenerPersonaPorId(id int) *dto.Persona {
	for _, persona := range personas {
		if persona.ID == id {
			return persona
		}
	}

	return nil
}

func CrearPersona(nuevaPersona dto.Persona) *dto.Persona {

	nuevaPersona.ID = len(personas) + 1

	personas = append(personas, &nuevaPersona)

	return &nuevaPersona
}
