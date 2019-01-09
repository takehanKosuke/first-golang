package goland

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)


type Person struct {
	//隠したい時は`json:"-"`とする事で隠せる
	//Name string 	`json:"name, omitenpty"`とする事で値が0のときは隠す事が出来る
	//Age  int		`json:"age, string"`とするとstringとしてあつかえる

	Name string 	`json:"name"`
	Age  int		`json:"age"`
}

//入ってきたjsonの値を弄るためにはMarshalJSONを定義するとMarshalされるときにかわりによばれる
func (p Person) MarshalJSON() ([]byte, error)  {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}


var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
	}
func main() {
	b := []byte(`{"name": "koura", "age": 24}`)
i9
	var p Person
	//↓.json.Unmarshalはjsonから構造体に変化させるもの
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age)
	//.json.Marshalは構造体からjsonに変化させるもの
	//構造体のほうでjsonにしたときのキーの指定をしないと単語の先頭が大文字になってしまうので指定するひつようがある
	v, _ := json.Marshal(p)
	fmt.Println(string(v))



	//hmac
	const apiKey = "User1key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)
	Server(apiKey, sign, data)
}
