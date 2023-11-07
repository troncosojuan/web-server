
package main

import (
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)


// estructura producto
type Producto struct {
	Id    int    `json:"id"`
	Nombre  string `json:"nombre"`
	Precio int    `json:"precio"`
	Stock int    `json:"stock"`
	Codigo string `json:"codigo"`
	Publicacion string `json:"publicacion"`
	FechaDeCreacion string `json:"fechaDeCreacion"`
}





func main() {
	engine := gin.Default()
	 // Define una ruta para obtener la lista de productos desde un archivo JSON
	 engine.GET("/productos", func(c *gin.Context) {
        // Abre el archivo JSON
        file, err := os.Open("productos.json")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir el archivo JSON"})
            return
        }
        defer file.Close()

        // Decodifica el contenido JSON en una estructura de datos
        var productos []Producto
        decoder := json.NewDecoder(file)
        if err := decoder.Decode(&productos); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar el archivo JSON"})
            return
        }

        // Devuelve la lista de productos como una respuesta JSON
        c.JSON(http.StatusOK, productos)
    })
	engine.Run()
}

