package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve saludo a persona en específico
func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("Nombre vacío")
	}
	// Se retorna el saludo
	message := fmt.Sprint(randomFormat())
	return message, nil
}

// Devolver varios saludos
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"Qué bueno verte, %v!",
		"Saludo, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}
