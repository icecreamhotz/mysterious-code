// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	text := string(sd)
	splitColons := strings.Split(text, ":")

	for i, j := 0, len(splitColons)-1; i < j; i, j = i+1, j-1 {
		splitColons[i], splitColons[j] = Reverse(splitColons[j]), Reverse(splitColons[i])
	}

	for i, v := range splitColons {
		if i == 0 {
			whatIsIt = fmt.Sprintf("%s", v)
		} else {
			whatIsIt = fmt.Sprintf("%s:%s", whatIsIt, v)
		}
	}

	fmt.Println(whatIsIt)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = fmt.Sprintf("%s%s", string(v), result)
	}
	return
}
