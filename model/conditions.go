package model

type Conditions []Condition

type Condition struct {
	Code int
	Day string
	Night string
	Icon int
	Languages[] Language
}

type Language struct {
	Lang_name string
	Lang_iso string
	Day_text string
	Night_text string
}
