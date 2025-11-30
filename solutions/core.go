package solutions

type Solution interface {
	Day() string
	Execute1(input string) string
	Execute2(input string) string
}
