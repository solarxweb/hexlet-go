package greeting

func Hello() string {
	return `Golang for Brave & ` + developer["Name"][:1] + developer["Name"][7:];
}