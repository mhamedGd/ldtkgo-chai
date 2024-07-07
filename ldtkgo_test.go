package ldtkgo

import (
	"fmt"
	"os"
	"testing"
)

func TestOpacityExistence(t *testing.T) {
	file_body, err := os.ReadFile("TestAssets/test.ldtk")
	if err != nil {
		fmt.Printf("[Issue Finidng File]: %v", err)
	}

	reader, err := Read(file_body)
	if err != nil {
		fmt.Printf("[Opening LDTK file Error]: %v", err.Error())
	}

	wanted_layer := reader.Levels[0].LayerByIdentifier("Water")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("==== Testing Output ===================================")
	fmt.Println(wanted_layer.Opacity)
	fmt.Println("=======================================================")
	fmt.Println("")

}
