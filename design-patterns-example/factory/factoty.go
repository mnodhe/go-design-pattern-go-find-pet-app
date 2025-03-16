package main

func main() {
	//this is better than
	//factory := products.Product{}
	//product := factory.New()
	//fmt.Println("My product was created at ", product.CreatedAt.UTC())
	// better than doing this bellow
	//product := products.Product{
	//	ProductName: "widget",
	//	CreatedAt:   time.Now(),
	//	UpdatedAt:   time.Now(),
	//}
	//fmt.Println("My product was created at ", product.CreatedAt.UTC())

}
