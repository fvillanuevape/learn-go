module helloworld

go 1.19

// Se hace referencia al paquete creado
require github.com/fvillanuevape/calculator v0.0.0

//especifica el uso de un directorio local
replace github.com/fvillanuevape/calculator => ../calculator
