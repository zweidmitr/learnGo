package main

import (
	"fmt"
)

type Builder interface {
	Build()
}
type Person struct {
	Name string
	Age  int
}
type WorkExperience struct {
	Name string
	Age  int
}

func (p Person) printName() {
	fmt.Println(p.Name)
}
func (wb WoodBuilder) printName() {
	fmt.Println(wb.Name)
}

type WoodBuilder struct {
	Person
	//Name string
	//WorkExperience
}

type BrickBuilder struct {
	Person
}

func (wb WoodBuilder) Build() {
	fmt.Println("Строю дом из дерева")
}
func (bb BrickBuilder) Build() {
	fmt.Println("Строю из кирпича")
}

type Building struct {
	Builder
	Name string
}

func main() {
	//explanation()
	usecase()
}

/*func explanation() {
	//builder := WoodBuilder{Person{Name: "Vasya", Age: 30}}
	//builder := WoodBuilder{Person{Name: "Vasya", Age: 30}, "stroitel"}
	builder := WoodBuilder{
		Person{Name: "Vasya", Age: 30},
		"stroitel",
		WorkExperience{"taxist", 3}}
	fmt.Printf("Type: %T Value: %#v \n", builder, builder)

	// shorthands / colliding
	fmt.Println(builder.Person.Age)
	fmt.Println(builder.WorkExperience.Age)

	// shadowing
	fmt.Println(builder.Name)

	builder.printName()
	builder.Person.printName()
}*/

func usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			"Vasya", 40,
		}},
		Name: "Деревянная изба",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{Person{
			"Petya", 30,
		}},
		Name: "Кирпичныйй дом",
	}
	brickBuilding.Build()
}
