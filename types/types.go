package types

type Money int64

type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Minute               = 60 * Nanosecond
	Minutes              = 60 * Microsecond
	Hour                 = 60 * Minute
)

const (
	TJS = "TJS"
	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"
)

type Card struct {
	Id       int
	Pan      string
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type Currency string

type PAN string

type Payment struct {
	Id        int
	AccountId int64
	Amount    Money
	Category  Category
	Status    PaymentStatus
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

type Category string

type PaymentStatus string

type Account struct {
	Id      int64
	phone   Phone
	Balance Money
}

type Phone string

type Client struct {
	Id       int64
	Name     string
	Lastname string
}

type Favorite struct {
	Id        int
	AccountId int64
	Name      string
	Amount    Money
	Category  Category
}
