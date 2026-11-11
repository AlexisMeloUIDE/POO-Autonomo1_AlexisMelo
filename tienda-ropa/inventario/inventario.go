package inventario

import (
	"errors"
	"tienda-ropa/venta"
)

// FiltrarPorCategoria recorre el catálogo y devuelve las prendas que coinciden con la categoría.
func FiltrarPorCategoria(prendas []venta.Prenda, cat string) []venta.Prenda {
	var resultado []venta.Prenda
	for _, p := range prendas {
		if p.Categoria == cat {
			resultado = append(resultado, p)
		}
	}
	return resultado
}

// BuscarPorID busca una prenda específica en el slice mediante su identificador.
func BuscarPorID(prendas []venta.Prenda, id string) (venta.Prenda, bool) {
	for _, p := range prendas {
		if p.ID == id {
			return p, true
		}
	}
	return venta.Prenda{}, false
}

// ActualizarStock modifica las existencias de una prenda comprobando que no quede en negativo.
func ActualizarStock(prendas []venta.Prenda, id string, delta int) ([]venta.Prenda, error) {
	for i, p := range prendas {
		if p.ID == id {
			nuevoStock := p.Stock + delta
			if nuevoStock < 0 {
				return prendas, errors.New("stock insuficiente para realizar la operación")
			}
			prendas[i].Stock = nuevoStock
			return prendas, nil
		}
	}
	return prendas, errors.New("prenda no encontrada en el inventario")
}
