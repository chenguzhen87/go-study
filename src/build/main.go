package main
import "fmt"
import "./de"

func main()  {
	for {
		fmt.Printf("%d name\n",de.Name)
		fmt.Printf("%d age\n",de.Age)
	}
}