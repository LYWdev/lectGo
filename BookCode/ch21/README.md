# ch21 함수 고급편

## 가변함수 인수 

함수 인수가 많은 경우 어떻게 처리해야할까?

Println 은 다양한 타임의 인수를 복수의ㅣ 개수로 받는다 어떻게 구현했을까?

```go
func sum(nums ...int)
```

위 같이 선언하면 복수의 인수를 받을 수 있다.

가변 인수와 인터페이스를 응용하면 Print 를 구현할 수 있다.

```go
func Print(args ...interface[])string {
  for _, arg := range args {
    switch f := arg.(type) {
      case bool :
        val := arg.(bool)
        Print(val)
      case float64 :
        val := arg.(bool)
        Print(val)
    }
  }
}
```