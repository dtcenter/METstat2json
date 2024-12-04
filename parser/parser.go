package main
import "fmt"
import "parser/pkg/structColumnDefs"

// see https://yuminlee2.medium.com/golang-import-local-packages-cbb2201c72e1
// for a brief explanation of how to define and use local imports

// cd to the root of the project and run:
// go mod init parser
// module parser
// go run .

type Ret struct {
	name string
	value string
}

func myStructParser(columnDef structColumnDefs.ColumnDef, line string) (interface{}, error) {
	fmt.Printf("myStructParser %s is parsing`%s`\n", columnDef.Name, line)
	ret := Ret{}
	ret.name = "fred"
	ret.value = "something"
	return ret, nil
}


func main() {
    fmt.Printf("\nstructColumnDefs\n")
	ret, error := structColumnDefs.ParseIt(myStructParser, "VL1L2", "this is a line of data")
	fmt.Printf("returned %s with err %s\n", ret, error)
    ret, error = structColumnDefs.ParseIt(myStructParser, "VAL1L2", "this is a line of data")
	fmt.Printf("returned %s with err %s\n", ret, error)
}