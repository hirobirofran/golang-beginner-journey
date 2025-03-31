// mainパッケージを宣言します。
// Go では、アプリケーションとして実行されるコードはmainパッケージ内にある必要があります。
package main

// 2 つのパッケージをインポートします。example.com/greetingsおよびfmtパッケージ です。
// これにより、コードはそれらのパッケージ内の関数にアクセスできるようになります。
// example.com/greetings(先ほど作成したモジュールに含まれるパッケージ) をインポートすると、関数にアクセスできるようになります。
// また、入力テキストと出力テキストを処理する関数 (コンソールへのテキストの出力など) を含むfmtもインポートします。
import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// 挨拶メッセージを取得して印刷します。 Get a greeting message and print it.
	// greetingsパッケージの Hello関数 を呼び出して挨拶を取得します。
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
