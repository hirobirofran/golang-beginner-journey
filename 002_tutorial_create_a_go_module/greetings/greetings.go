package greetings

import "fmt"

// Hello は指定された人への挨拶を返します。 Hello returns a greeting for the named person.
func Hello(name string) string {
	// メッセージに名前を埋め込んだ挨拶を返します。 Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
