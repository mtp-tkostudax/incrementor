package incrementor

var (
	Value int
)

func Increment(ptr *int) {
	*ptr++
}
