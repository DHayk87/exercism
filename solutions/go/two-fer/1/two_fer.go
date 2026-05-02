package twofer
import "fmt"
// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name != ""{
		str := fmt.Sprintf( "One for %s, one for me." , name)
		return str
	}
	return  "One for you, one for me."
}
