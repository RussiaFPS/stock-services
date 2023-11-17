package types

type Product struct {
	Id     string `json:"id"`
	Amount int    `json:"amount"`
}

type Params struct {
	Products []Product `json:"products"`
}

type Stock struct {
	Name string `json:"name"`
}

type Result struct {
	ProductsId  string
	Amount      int
	CountReserv int
}

type ResponseResult struct {
	Message []Result `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}
