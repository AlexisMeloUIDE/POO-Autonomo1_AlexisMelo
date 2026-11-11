
package venta

// Prenda representa la estructura inmutable de un artículo de ropa.
type Prenda struct {
	ID        string
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
}
// ItemCarrito vincula una prenda con la cantidad seleccionada.
type ItemCarrito struct {
	Prenda   Prenda
	Cantidad int
}
// Carrito contiene la lista de prendas seleccionadas para la compra.
type Carrito struct {
	Items []ItemCarrito
}
// Factura almacena el detalle final del proceso de compra.
type Factura struct {
	ID       string
	Subtotal float64
	IVA      float64
	Total    float64
}
// CrearPrenda es un constructor que retorna una nueva estructura Prenda.
func CrearPrenda(id, nombre, cat string, precio float64, stock int) Prenda {
	return Prenda{
		ID:        id,
		Nombre:    nombre,
		Categoria: cat,
		Precio:    precio,
		Stock:     stock,
	}
}
// NuevoCarrito inicializa y devuelve una estructura Carrito vacía.
func NuevoCarrito() Carrito {
	return Carrito{Items: []ItemCarrito{}}
}
// AgregarItem agrega una prenda y su cantidad al carrito.
func AgregarItem(c Carrito, p Prenda, cant int) Carrito {
	c.Items = append(c.Items, ItemCarrito{Prenda: p, Cantidad: cant})
	return c
}
