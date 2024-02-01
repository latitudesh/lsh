package output

import "fmt"

func SuccessfulDeletion(resource string) {
	fmt.Printf("\n\n%v deleted successfully!\n\n", resource)
}
