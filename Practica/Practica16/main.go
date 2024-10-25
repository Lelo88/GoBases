package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type Contact struct {
    Name  string
    Phone string
    Email string
}

func saveContactsToFile(contacts []Contact) error {
    file, err := os.Create("contacts.json")
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(contacts)
    if err != nil {
        return err
    }

    return nil
}

func loadContactsFromFile() ([]Contact, error) {
    file, err := os.Open("contacts.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var contacts []Contact
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&contacts)
    if err != nil {
        return nil, err
    }

    return contacts, nil
}

func main() {
    contacts, err := loadContactsFromFile()
    if err != nil {
        fmt.Println("Error al cargar los contactos")
        contacts = []Contact{}
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("=== GESTOR DE CONTACTOS ===\n")
        fmt.Print("1. Agregar contacto\n")
        fmt.Print("2. Mostrar contactos\n")
        fmt.Print("3. Salir\n")

        var option int
        _, err := fmt.Scanln(&option)
        if err != nil {
            fmt.Println("Error al leer la opción")
            return
        }

        switch option {
        case 1:
            var c Contact
            fmt.Print("Nombre: ")
            c.Name, _ = reader.ReadString('\n')
            c.Name = strings.TrimSpace(c.Name)

            fmt.Print("Teléfono: ")
            c.Phone, _ = reader.ReadString('\n')
            c.Phone = strings.TrimSpace(c.Phone)

            fmt.Print("Email: ")
            c.Email, _ = reader.ReadString('\n')
            c.Email = strings.TrimSpace(c.Email)

            contacts = append(contacts, c)
            err = saveContactsToFile(contacts)
            if err != nil {
                fmt.Println("Error al guardar el contacto")
                return
            }
        case 2:
            for i, c := range contacts {
                fmt.Printf("%d. %s - %s - %s\n", i+1, c.Name, c.Phone, c.Email)
            }
        case 3:
            return
        default:
            fmt.Println("Opción no válida")
        }
    }
}