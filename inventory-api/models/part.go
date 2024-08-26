package models

type Part struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Quantity      int    `json:"quantity"`
	Location      string `json:"location"`
	Minimun_Stock int    `json:"minimun_stock"`
}

// las variables y tipos que empiezan con una letra mayúscula están exportadas y son accesibles desde otros paquetes
var Parts = []Part{
	{
		ID: 1, Name: "Ruleman", Description: "Descripción del ruleman", Quantity: 10, Location: "Lugar en el taller", Minimun_Stock: 10,
	},
	{
		ID: 2, Name: "Cigueñal", Description: "Descripción del Cigueñal", Quantity: 2, Location: "Lugar en el taller", Minimun_Stock: 2,
	},
	{
		ID: 3, Name: "Rotor", Description: "Descripción del Rotor", Quantity: 5, Location: "Lugar en el taller", Minimun_Stock: 5,
	},
}