package facturacion

import "tienda-ropa/venta"

// CalcularSubtotal recorre el carrito acumulando (precio * cantidad) de cada ítem.
func CalcularSubtotal(c venta.Carrito) float64 {
	subtotal := 0.0
	for _, item := range c.Items {
		subtotal += item.Prenda.Precio * float64(item.Cantidad)
	}
	return subtotal
}

// AplicarDescuento calcula el monto resultante tras aplicar un porcentaje de descuento.
func AplicarDescuento(monto float64, pct float64) float64 {
	return monto - (monto * pct)
}

// CalcularIVA determina el impuesto sobre el valor agregado en función de la tasa dada.
func CalcularIVA(monto float64, tasa float64) float64 {
	return monto * tasa
}

// GenerarFactura integra la secuencia de cálculos y ensambla la estructura Factura.
func GenerarFactura(id string, c venta.Carrito, descPct float64, tasaIVA float64) venta.Factura {
	subtotal := CalcularSubtotal(c)
	subtotalConDesc := AplicarDescuento(subtotal, descPct)
	iva := CalcularIVA(subtotalConDesc, tasaIVA)
	total := subtotalConDesc + iva

	return venta.Factura{
		ID:       id,
		Subtotal: subtotal,
		IVA:      iva,
		Total:    total,
	}
}
