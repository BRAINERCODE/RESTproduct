package Models

type Products struct {
	IdProduct int
	Nombre    string
	Stock     int
	PrecioU   float32
	IdCiudad  int
}

func (p *Products) SetIdProduct(id int) {

	p.IdProduct = id
}
func (p *Products) SetNombre(nombre string) {

	p.Nombre = nombre
}
func (p *Products) SetStock(stock int) {

	p.Stock = stock
}
func (p *Products) SetPrecioU(preciou float32) {

	p.PrecioU = preciou
}
func (p *Products) SetIdCiudad(idCiudad int) {

	p.IdCiudad = idCiudad
}

func (p *Products) GetIdProduct() int {
	return p.IdProduct
}
func (p *Products) GetNombre() string {
	return p.Nombre
}
func (p *Products) GetStock() int {
	return p.Stock
}
func (p *Products) GetPrecioU() float32 {
	return p.PrecioU
}
func (p *Products) GetIdCiudad() int {
	return p.IdCiudad
}
