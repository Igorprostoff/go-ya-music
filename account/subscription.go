package account

type Subscription struct {
	Non_auto_renewable_remainder Renewable_remainder
	Auto_renewable               []Auto_renewable
	Family_auto_renewable        []Auto_renewable
	Had_any_subscription         bool
	Operator                     []Operator
	Non_auto_renewable           Non_auto_renewable
	Can_start_trial              bool
	Mcdonalds                    bool
	End                          string
}
