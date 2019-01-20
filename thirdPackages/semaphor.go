package main
// セマフォは並行処理のかずを制限するものです

import (
  "context"
  "fmt"
  "time"

  "golang.org/x/sync/semaphore"
)

var s*semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
  // [s.Acquire]でロックする感じ
  if err := s.Acquire(ctx, 1); err != nil {
    fmt.Println(err)
    return
  }
  // [s.Release]でロック解除
  defer s.Release(1)
  fmt.Println("Waite...")
  time.Sleep(1*time.Second)
  fmt.Println("Done")
}

func main()  {
  ctx := context.TODO()
  go longProcess(ctx)
  go longProcess(ctx)
  go longProcess(ctx)
  time.Sleep(5*time.Second)
}
