package main

import "fmt"
import "github.com/gojp/kana"

func writeToJapanese(x string) {
       s := kana.RomajiToKatakana(x)
        fmt.Println(x," in Japanese Katakana :", s)
        r := kana.RomajiToHiragana(x)
        fmt.Println(x, " in Japanese Hiragana :", r)
}

func main() {
	writeToJapanese("benoy")
        writeToJapanese("asha")
}
