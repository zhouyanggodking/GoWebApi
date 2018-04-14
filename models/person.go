package models

type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Title string `json:"title"`
}
