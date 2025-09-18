// Aplicación Wiki Web en Go
// Esta aplicación crea un servidor web que permite crear, editar y ver páginas wiki
package main

import (
	"html/template" // Para procesar plantillas HTML
	"log"           // Para logging de errores
	"net/http"      // Para el servidor HTTP
	"os"            // Para operaciones del sistema de archivos
	"regexp"        // Para validar URLs con expresiones regulares
)

// Page representa una página wiki con título y contenido
// Los tags json permiten serializar/deserializar a JSON si fuera necesario
type Page struct {
	Title string // Título de la página
	Body  []byte // Contenido de la página en bytes (permite cualquier tipo de texto)
}

// save es un método de Page que guarda la página en un archivo .txt
// Retorna error si no puede escribir el archivo
func (p *Page) save() error {
	filename := p.Title + ".txt" // Crea el nombre del archivo basado en el título
	// WriteFile crea/sobrescribe el archivo con permisos 0600 (solo el propietario puede leer/escribir)
	return os.WriteFile(filename, p.Body, 0600)
}

// loadPage carga una página desde un archivo .txt
// Retorna un puntero a Page y un error si no puede leer el archivo
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"         // Construye el nombre del archivo
	body, err := os.ReadFile(filename) // Lee todo el contenido del archivo
	if err != nil {
		return nil, err // Si hay error (ej: archivo no existe), retorna nil y el error
	}
	// Crea y retorna una nueva Page con el título y contenido leído
	return &Page{Title: title, Body: body}, nil
}

// viewHandler maneja las peticiones GET para ver una página
// Si la página no existe, redirige al editor para crearla
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title) // Intenta cargar la página
	if err != nil {
		// Si la página no existe, redirige a la página de edición
		// StatusFound = 302 (redirección temporal)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// Si la página existe, la muestra usando la plantilla "view"
	renderTemplate(w, "view", p)
}

// editHandler maneja las peticiones GET para editar una página
// Muestra el formulario de edición con el contenido actual (si existe)
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title) // Intenta cargar la página existente
	if err != nil {
		// Si la página no existe, crea una nueva página vacía
		p = &Page{Title: title}
	}
	// Muestra el formulario de edición usando la plantilla "edit"
	renderTemplate(w, "edit", p)
}

// saveHandler maneja las peticiones POST para guardar una página
// Recibe el contenido del formulario y guarda la página
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body") // Obtiene el valor del campo "body" del formulario HTML
	// Crea una nueva página con el título y el contenido recibido
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save() // Guarda la página en un archivo
	if err != nil {
		// Si hay error al guardar, responde con error 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Si se guarda exitosamente, redirige a la página de visualización
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// templates carga y parsea las plantillas HTML al iniciar la aplicación
// Must() hace que el programa termine con panic si hay error al cargar las plantillas
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// renderTemplate ejecuta una plantilla HTML específica con los datos de una página
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// ExecuteTemplate ejecuta la plantilla especificada (tmpl+".html")
	// y pasa los datos de la página (p) para rellenar las variables de la plantilla
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		// Si hay error al ejecutar la plantilla, responde con error 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// validPath es una expresión regular que valida las URLs permitidas
// Solo acepta rutas como: /edit/PageName, /save/PageName, /view/PageName
// donde PageName solo puede contener letras y números
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// makeHandler es una función de orden superior que crea handlers HTTP con validación
// Toma una función handler y retorna una nueva función que valida la URL antes de ejecutar el handler
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Valida que la URL coincida con el patrón esperado
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			// Si la URL no es válida, responde con 404 Not Found
			http.NotFound(w, r)
			return
		}
		// Si la URL es válida, ejecuta el handler pasando el título de la página
		// m[2] contiene el segundo grupo capturado en la regex (el nombre de la página)
		fn(w, r, m[2])
	}
}

// main es la función principal que configura y ejecuta el servidor web
func main() {
	// Registra los handlers para cada ruta, envueltos en makeHandler para validación
	// /view/ -> muestra una página existente
	http.HandleFunc("/view/", makeHandler(viewHandler))
	// /edit/ -> muestra el formulario para editar una página
	http.HandleFunc("/edit/", makeHandler(editHandler))
	// /save/ -> procesa el formulario y guarda la página
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// Inicia el servidor HTTP en el puerto 8080
	// log.Fatal registra el error y termina el programa si el servidor falla
	log.Fatal(http.ListenAndServe(":8080", nil))
}
