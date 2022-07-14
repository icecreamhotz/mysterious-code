// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	text := string(sd)

	for _, v := range text {
		whatIsIt = fmt.Sprintf("%s%s", string(v), whatIsIt)
	}

	fmt.Println(whatIsIt)
}
