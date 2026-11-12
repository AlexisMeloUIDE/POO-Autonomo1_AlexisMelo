# Sistema de Gestión para Tienda de Ropa 

Sistema modular de gestión comercial desarrollado en **Go (Golang)** como parte del **Aprendizaje Autónomo 1** de la asignatura *Programación Orientada a Objetos*. 

El proyecto automatiza las funciones funamentales de unatienda de ropa mediante el uso de 4 paquetes funcionales dentro de un único módulo de Go.

## Tabla de funciones resumida

El proyecto sigue una estructura de módulo único (`go.mod`) dividido en 4 paquetes con responsabilidades delimitadas:

| Paquete | Responsabilidad Principal | Funciones Clave |
| :--- | :--- | :--- |
| **`venta`** | Definición de prenda, factura y gestión del carrito de compras. | `CrearPrenda`, `NuevoCarrito`, `AgregarItem` |
| **`inventario`** | Filtros, búsqueda por ID y control de existencias. | `FiltrarPorCategoria`, `BuscarPorID`, `ActualizarStock` |
| **`facturacion`** | Sección financiera para cálculo de subtotales, descuentos e impestos. | `CalcularSubtotal`, `AplicarDescuento`, `CalcularIVA`, `GenerarFactura` |
| **`reportes`** | Análisis de las ventas totales e identificación de stock crítico. | `TotalVentas`, `ObtenerStockBajo` |

---

## Estructura del Repositorio

```text
tienda-ropa/
├── go.mod                  ← Módulo raíz de Go (go mod init tienda-ropa)
├── README.md               ← Documentación general del proyecto
├── main.go                 ← Paquete principal
├── venta/                  ← Paquete 1: Definición de prenda, carrito, factura
│   └── venta.go
├── inventario/             ← Paquete 2: Filtros, búsquedas y control de stock
│   └── inventario.go
├── facturacion/            ← Paquete 3: Operaciones matemáticas para las facturas
│   └── facturacion.go
└── reportes/               ← Paquete 4: Analísis de ventas y alertas de stock
    └── reportes.go

