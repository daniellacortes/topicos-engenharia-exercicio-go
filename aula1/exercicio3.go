package main 
import "fmt"

	func IdadeLegal(idade int) bool {
		if idade <10 { 
			return false 
		}
		return true 
	}

func main() {
 fmt.Println(IdadeLegal(19))	
}