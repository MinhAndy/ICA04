package main

import (
	"fmt"

	speech "github.com/nicolaifsf/go-speak"
)

func main() {
	speech.SetWitKey("WD5YYFIV5H2VIV24TU2SWTWEEZAE5UMM")
	a := speech.SendWitVoice("test.wav")
	fmt.Print(a)
}
