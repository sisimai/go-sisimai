//go:build ignore
package main
import "os"
import "fmt"
import "libsisimai.org/sisimai/v5"

func main() {
	path := os.Args[1]
	args := sisimai.Args()
	args.Delivered = true
	args.Vacation  = true
	sisi, _ := sisimai.Rise(path, args)
	fmt.Printf("%d\n", len(sisi))
}
