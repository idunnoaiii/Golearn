package iteration

func Repeat(value string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += value
	}
	return repeated
}

func main() {

}
