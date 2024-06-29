package person

// type alias
type Sex int

const (
	Male Sex = iota
	Female
)

// Struct
type Customer struct {
	Name      string
	age       int
	sex       Sex
	isPremium bool
}

// Constructor
func NewCustomer(name string, age int, sex Sex) Person {
	return &Customer{
		Name:      name,
		age:       age,
		sex:       sex,
		isPremium: false,
	}
}

// method
func (c Customer) Greet() string {
	return "Hello " + c.Name
}

func (c Customer) GetAge() int {
	return c.age
}

func (c Customer) GetSex() string {
	return []string{"male", "female"}[c.sex]
}

func (c *Customer) SetName(name string) {
	c.Name = name
}

func (c *Customer) SetAge(age int) {
	c.age = age
}

func (c *Customer) SetSex(sex Sex) {
	c.sex = sex
}
