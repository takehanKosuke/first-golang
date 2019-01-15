package main

import (
  // rubyのputs的なことをしてくれるライブラリだと思う
  "fmt"
  // rubyのsleep処理を行うライブラリ
  "time"
  // サーバーを立ち上げるためのライブラリ
  "net/http"
)



// 構造体
// 多分modeleみたいな感じだと思う
type user struct {
  name string
  score int
}

func init()  {
  fmt.Println("init()はmain関数より先に呼ばれるファイルだよ〜")
}


// 構造体メソッド
func (u user) show() {
  fmt.Println("name:%s, score:%d", u.name, u.score)
}

// インターフェイス
// 構造体をメソッドを用いてグルーピングできる
type  greeter interface {
  greet()
}

type japanese struct {}
type USA struct {}

func (j japanese) greet()  {
  fmt.Println("こんにちは")
}
func (u USA) greet()  {
  fmt.Println("hello")
}

// 空のインターフェース
// なんでもインターフェース化できるので入ってきた方に構造体に応じて処理を変更させる
func show(s interface{}) {
  // if文で処理を変更させる
  _, ok := s.(japanese)
  if ok {
    fmt.Println("私は日本人です")
  } else {
    fmt.Println("私は日本人ではありません")
  }


  // case文で処理を分岐させる
  switch s.(type) {
  case japanese:
    fmt.Println("私は日本人です")
  default:
    fmt.Println("私は日本人ではありません")
  }
}





// 構造体のデータを書き換えてい場合は＊をつける
func (u *user) hit() {
  u.score++
}

func main() {
  // 変更するものはver
  // 変更させないものはconst
  // を用いて定義する
  const name = "田口"
  fmt.Println(name)

// 【iota】を使うと勝手に　１，２，３，４・・・・というように数字を振ってくれる
  const (
    sun = iota
    mon
    tue
  )
  fmt.Println(sun, mon, tue)


  // ポインタとは、変数とかのメモリの住所を記録するもの
  // ポインタの演算はできない
  a := 5
  var pa *int //ポインタの確保
  pa = &a
  fmt.Println(pa) //メモリの住所を表示
  fmt.Println(*pa) //中に入ってる変数を表示


  fmt.Println(swap(3, 8))

  // main関数の中に関数を作ることもできる
  f := func(c, d int) (int, int) {
    return d, c
  }
  fmt.Println(f(10, 200))

  // js即時関数みたいにできる。その場合は関数の後ろに（）をつけてあげると実行できる。
  func (msg string) {
    fmt.Println(msg)
  }("山田")

  //time関数について
  t := time.Now()
  fmt.Println(t)
  //「RFC3339」でいい感じのフォーマットになる
  fmt.Println(t.fomat(time.RFC3339))

  // スライス
  // 基本的にgolangでは配列ではなくスライスを使う
  c := []int{1, 3, 7} //こんな感じに定義できる
  fmt.Println(c)

  // golangにおけるmapとはrubyにおけるハッシュみたいなイメージ
  m := map[int]string{1: "こーすけ", 2: "ありま", 3: "こうら", 4: "けい"}
  fmt.Println(m)
  fmt.Println(len(m))//要素の個数を返す

  delete(m, 1)//要素の削除
  fmt.Println(m)

  v, ok := m[2]//要素が配列に存在するかを調べる（vに値がokに真偽値が帰ってくる）
  fmt.Println(v)
  fmt.Println(ok)

  // rangeはrailsのeachみたいなやつ
  for i, v := range c {
    fmt.Println(i, v)
  }
  // スライスの値だけ返したい場合「_」を与えてあげると値だけを返してくれる
  for _, v := range c {
    fmt.Println(v)
  }

  // railsでいうハッシュにもeachを使うこともできる
  for i, v := range m {
    fmt.Println(i, v)
  }

  // 構造体を使ってみる
  u := user{name: "こーすけ", score: 21}
  fmt.Println(u)
  // 生成した構造体のメソッドを使ってみる
  u.show()
  // 構造体の値を変化させる
  u.hit()
  fmt.Println(u.score)

  // インターフェースを使ってrangeを使って同じ処理を実行させる
  humans := []greeter{japanese{}, USA{}}
  for _, human := range humans {
    human.greet()
    show(human)
  }

  // 並行処理を行う goroutine
  // メソッド前に[go]をつけることで並行で処理を行うことができる
  // ただし返り値は受け取れない！！！
  // 返り値を受け取りたい場合はchanを使う必要がある
  result := make(chan string)

    go task1(result)
    go task2()
    fmt.Println(<-result)
    time.Sleep(time.Second * 3)


  // サーバーを立ち上げる
  // ルートでアクセスされたらhandlerを立ち上げてねっていうメソッド
  http.HandleFunc("/", handler)
  //立ち上げるポートの指定
  http.ListenAndServe(":8080", nil)
}

// handlerの記述この辺は定型文らしい。。。

func handler(w http.ResponseWriter, r*http.Request)  {
  // ここの出力するときの書き方は他のところとちょっと違う
  fmt.Fprintf(w, "hi %s!", r.URL.Path[1:])
}

func task1(result chan string)  {
  time.Sleep(time.Second * 2)
  // fmt.Println("task1 finished")
  // chanに値を入れて返す場合は以下のようにしなければならない
  result <- ("task1 finished")
}
func task2()  {
  fmt.Println("task2 finished")
}

// 複数の返り値を返す
func swap (a, b int) (int, int){
  return b,a
}

// 引数の数をよしなに変えて受け付ける「可変長引数」って言うらしい
// 引数を(params ...データ型)とすると可変の引数で処理できる
// 逆に呼び出すときにover(配列名...)とすることで呼び出すことができる
// とりあえず可変の引数を持つfuncを作るときは「...」使っとけってことですね
func over(params ...int)  {
  for _, param := range params {
    fmt.Println(param)
  }
}


// deferは呼び出された関数の一番最後に実行される
// deferが複数個あるときは下から実行される（スタッキングdeferっていうらしい）
// ファイルをオープンするときには必ずクローズの処理を挟まなければいけないのでそのときにdeferを入れることでクローズをわすれる心配をなくしている
func late()  {
  defer fmt.Println("呼び出された関数の一番最後に実行されるよ〜")
}

// ↓ファイルをよみこむときはこんな感じ「os」っていうライブラリ使ってるのかな？
func open_file()  {
  file, _ := os.OPEN("./lesson.go")
  defer file.Close()
  data := make([]byte, 100)
  file.READ(data)
  fmt.Println(string(data))
}




// ↓Goはenumはないけどenumっぽいものはできるよ

type Status int

type User struct {
  Name    string
  Age     int
  Status  Status
}

const (
  Nomal Status = iota
  Admin
  Ban
)

func main(){
  User1 := User{"kosuke", 21, Nomal}
  fmt.Println(User1)
}

// 〜出力〜
// {kosuke 21 0}
