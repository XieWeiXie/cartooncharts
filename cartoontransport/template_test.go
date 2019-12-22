package cartoontransport

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	tmp := Template{Path: "../cartoon/template"}
	content, _ := tmp.Read("plot.html")
	fmt.Println(string(content))

}
