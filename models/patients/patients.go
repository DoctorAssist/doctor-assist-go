package patients

type Sex int

const (
	Intersex = iota
	Female
	Male
)

type Patient struct {
	Name string `db:"name"`
	Sex Sex `db:"sex"`
	Age int `db:"age"`
	ID string `db:"id"`
	Phone string `db:"phone"`
	Email string  `db:"email"`
	Address string `db:"address"`
}
