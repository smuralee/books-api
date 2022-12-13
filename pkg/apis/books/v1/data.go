package v1

func initializeBooks() []Book {
	var Books = []Book{
		{Id: "1", Title: "Mozart in the Jungle", Author: "Blair Tindall"},
		{Id: "2", Title: "Bad Blood", Author: "John Carreyrou"},
		{Id: "3", Title: "The Feynman Lectures on Physics", Author: "Richard P. Feynman"},
	}
	return Books
}
