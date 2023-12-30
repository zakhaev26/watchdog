package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bashScript := []byte("#!/bin/bash\necho 'Hello, world!'")

	err := ioutil.WriteFile("hotttasd.sh", bashScript, 0755)
	fmt.Println("OOGA BOOGA")

	if err != nil {
		fmt.Println("Error creating Bash script:", err)
		return
	}

	fmt.Println("Bash script created successfully.")
}
