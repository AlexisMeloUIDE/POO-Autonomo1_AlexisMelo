package reportes

import "tienda-ropa/venta"

// TotalVentas suma el importe de todas las facturas procesadas en la tienda.
func TotalVentas(facturas []venta.Factura) float64 {
	total := 0.0
	for _, f := range facturas {
		total += f.Total
	}
	return total
}

// ObtenerStockBajo devuelve un slice con las prendas cuya cantidad es menor o igual al umbral.
func ObtenerStockBajo(prendas []venta.Prenda, umbral int) []venta.Prenda {
	var bajas []venta.Prenda
	for _, p := range prendas {
		if p.Stock <= umbral {
			bajas = append(bajas, p)
		}
	}
	return bajas
}
