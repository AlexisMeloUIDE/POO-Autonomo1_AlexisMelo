package main

import (
	"fmt"
	"tienda-ropa/facturacion"
	"tienda-ropa/inventario"
	"tienda-ropa/reportes"
	"tienda-ropa/venta"
)

func main() {
	fmt.Println("=== SISTEMA DE GESTIÓN DE TIENDA DE ROPA ===")

	// 1. Crear prendas iniciales del inventario
	catalogo := []venta.Prenda{
		venta.CrearPrenda("P001", "Camiseta Oversize Negro", "Superior", 25.00, 15),
		venta.CrearPrenda("P002", "Pantalón Cargo Beige", "Inferior", 45.00, 4),
		venta.CrearPrenda("P003", "Chaqueta Denim Azul", "Abrigos", 60.00, 2),
	}

	// 2. Probar filtrado del módulo inventario
	fmt.Println("\n--- Filtrando Prendas por Categoría: Superior ---")
	superiores := inventario.FiltrarPorCategoria(catalogo, "Superior")
	for _, p := range superiores {
		fmt.Printf("- %s | Precio: $%.2f | Stock: %d\n", p.Nombre, p.Precio, p.Stock)
	}

	// 3. Crear Carrito y agregar prendas
	carrito := venta.NuevoCarrito()
	if len(superiores) > 0 {
		carrito = venta.AgregarItem(carrito, superiores[0], 2)
	}
	if prenda, existe := inventario.BuscarPorID(catalogo, "P002"); existe {
		carrito = venta.AgregarItem(carrito, prenda, 1)
	}

	// 4. Probar módulo facturación (10% descuento, 15% IVA)
	factura := facturacion.GenerarFactura("FAC-2026-001", carrito, 0.10, 0.15)
	fmt.Println("\n--- Resumen de Factura Generada ---")
	fmt.Printf("Factura ID: %s\n", factura.ID)
	fmt.Printf("Subtotal:   $%.2f\n", factura.Subtotal)
	fmt.Printf("IVA (15%%):  $%.2f\n", factura.IVA)
	fmt.Printf("Total:      $%.2f\n", factura.Total)

	// 5. Probar módulo reportes
	historialFacturas := []venta.Factura{factura}
	totalRecaudado := reportes.TotalVentas(historialFacturas)
	stockCritico := reportes.ObtenerStockBajo(catalogo, 5)

	fmt.Println("\n--- Reporte Comercial ---")
	fmt.Printf("Ingresos totales acumulados: $%.2f\n", totalRecaudado)
	fmt.Printf("Prendas con stock bajo (<=5 unidades): %d\n", len(stockCritico))
	for _, p := range stockCritico {
		fmt.Printf("  * Alerta: %s (Stock actual: %d)\n", p.Nombre, p.Stock)
	}
}
