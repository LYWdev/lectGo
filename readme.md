# Tucker'Go lang Lecture Note
___

## 1.1 Go의 역사

1. Go 언어는 2009 년에 발표된 언어
2. 로버트 그리스머, 롭 파이크, 켄 톰슨 주축하에 구글에서 만든 오픈소스 프로그래밍 언어이다.


## Go언어의 특징

|개념   |존재여부|설명|
|:---:|:---:|:---:|
|클래스             | no  | 클래스는 없지만 메서드를 가지는 구조체를 지원한다.    |
|상속             | no  | 상속을 지원하지 않음.    |
|메서드           | y   |   구조체가 메서드를 가질 수 있다.     |
|인터페이스        |  y  |  상속이 없지만 인터페이스를 가질 수 있다.    |
|익명함수          | y   |  함수 리터럴이라는 이름으로 제공한다.    |
|가비지 컬렉터      | y   |고성능 가비지 컬렉터      |
|포인터           | y   | 포인터를 지원한다.|
|제네릭 프로그래밍   |   no|  제네네 프로그래로을 지원하지 않음|
|네임스페이스       |  no |   네임스페이스 제공하지 않음 모든 코드는 패키지 단위로 분리된다.|

## 변수가 타입을 갖고 있는 이유
1. 메모리 공간에 어디서부터 어디까지가 하나의 의미상 구조인지 알기 위해 필요하다. </br>
타입을 알면 어디서부터 어디까지가 하나의 변수인지 알 수 있다.  
포인터는 하나의 시작점만 알려주기 때문에 끝점을 알고자하면 결국 크기를 알아야한다.  

2. 타입을 알아야 데이터를 해석할 수 있다.


| 이진수|타입|실제 값|
|:---:|:---:|:---:|
| 1000 0000|unit8|128|
| 1000 0000|int8|-128|
| 1000 0000|float32|1.79366203434e-43|

위 표에서 처럼 같은 1000 0000이라는 이진수임에도 불구하고 타입에 따라 실제 값이 달라진다.

## 타입별 기본값

var b int  처럼 초기 값을 선언하지 않는 경우가 있다.
이런경우 타입별 기본값이 대입된다.
|타입|기본값|
|:--:|--|
|모든 정수 타입|0|
|모든 실수 타입|0.0|
|불리언 |flase|
| 문자열|""(빈 문자열 값을 넣음)|
| 나머지|nil( 정의되지 않은 메모리 주소를 나타내는 Go 키워드)|


## 타임 변환(Type Casting)주의할 점.

1. 실수 타입에서 정수 타입으로 변환하면 소수점 이하의 숫자가 사라진다.
2. 큰 범위를 갖는 타입에서 작은 범위르 갖는 타입으로 변환하면 값이 달라질 수 있다.
2.1 예를 들어 in16에서 1바이트 정수 8로 변환하면 아래처럼 비트 연산에서 오류가 있다.
최상위 한바이트가 사라지기 때문에 -128이라는 엉뚱한 값이 나온다.
타입변환할 때, 비트를 주의하자.

```go
 {
	var a int16 = 3456
	var c int8 = int8(a)
	
	fmt.Println(a)
	fmt.Println(c)
 }
```
|0|0|0|0|0|0|0|0|   0|0|0|0|0|0|0|0|
--|--|--|--|--|--|  --|--|--|--|--|--|--|--|--|--|
|0|0|0|0|1|1|0|1|   1|0|0|0|0|0|0|0|
|0|0|0|0|0|0|0|0|   1|0|0|0|0|0|0|0|

## go언어 서식문자 정리

| 서식문자 | 출력 형태                          |
|----------|------------------------------------|
| %t       | bool                               |
| %b       | 2진수 정수                         |
| %c       | 문자                               |  
| %d       | 10진수 정수                        |    
| %o       | 8진수 정수                         |    
| %x       | 16진수 정수, 소문자                |  
| %X       | 16진수 정수, 대문자                |   
| %f       | 10진수 방식의 고정 소수점 실수     |   
| %F       | 10진수 방식의 고정 소수점 실수     |   
| %e       | 지수 표현 실수, e                  |  
| %E       | 지수 표현 실수, E                  |  
| %g       | 간단한 10진수 실수                 |   
| %G       | 간단한 10진수 실수                 |   
| %s       | 문자열                             |
| %p       | 포인터                             |
| %U       | 유니코드                           |    

## 실수 오차 
6.3 실수 오차

컴퓨터에서 실수 값을 표현하는 방법의 한계점
1. 10진수가 아닌 2진수로 표현한다.
1항으로 인해 실수 연산에서 한계점이 들어난다.
 1.2 예
	0.375는 10진수로 표현하면 3*10^-1 + 7*10^-2 + 5*10^-3이다. 
	위 같은 방법을 2진수로 바꾸어 표현하면 1*2^-2 +1*2^-3이 된다.
	이같은 방법으로 모든 숫자를 표현할 수 있음 매우 좋지만,
	현실적으로 불가능하다.
	예를 들어)
	0.376를 이진수로 나타내면 위 방법으로는 표현할 수  없다.

이 책에서는 작은 오차 무시하는 방법으로 해결하는 방법을 알려준다.

컴퓨터에서 실수 값을 표현하는 방법의 한계점
1. 10진수가 아닌 2진수로 표현한다.
1항으로 인해 실수 연산에서 한계점이 들어난다

##7.2 함수를 호출하면 생기는 일
함수를 호출할 때 입력하는 값을 argument라고 한다.
함수가 외부에게 입력받는 변수를 parameter라고 부른다. 

argument 는 매개변수로 복사된다.
매개변수, 함수내에서 선언된 변수는 함수가 종료되면 변수 범위를 벗어나서 접근하지 못한다.