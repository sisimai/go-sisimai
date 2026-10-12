//go:build ignore
package main
import "os"
import "strconv"
import "libsisimai.org/sisimai/v5"

func main() {
	path := os.Args[1]
	args := sisimai.Args()
	args.Delivered = true
	args.Vacation  = true
	sisi, _ := sisimai.Rise(path, args)
	os.Stdout.WriteString(strconv.Itoa(len(sisi)) + "\n")
}
