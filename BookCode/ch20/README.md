# ch20 인터페이스

## 인터페이페란 구현을 포함하지 않는 메서드 집합을 의미.

인터페이스만으로 메서드를 호출할 수 있음. 따라서 프로그램 요구사항 변경 시 유연하게 대처할 수 있음.

선언하는 방법
```go
type DuckInterface interface{
  Fly()
  Walk(distance int)int
}
```
인터페이스도 구조체 처럼 type 를 사용해야한다.
