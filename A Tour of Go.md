# A Tour of Go

📝  [A Tour of Go](https://tour.golang.org/)



1️⃣  Go의 시작

2️⃣  변수와 함수

3️⃣  반복문

4️⃣  구조체

5️⃣  자료구조

6️⃣  함수 클로저

7️⃣  메소드와 인터페이스

### 01. Hello, 안녕

### 02. Go local

### 03. Go offline [Go offline](link)

### 04. Package (패키지)

### 05. Imports (임포트)

### 06. Exported names

- 패키지를 import하면 패키지가 외부로 export 한 것들 (메서드나 변수, 상수 등)에 접근할 수 있음

- Go에서는 첫 문자가 대문자로 시작하면 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됨

  ```go
  Foo // 외부에서 참조 가능
  FOO // 외부에서 참조 가능
  foo // 외부에서 참조 불가
  ```

  ```go
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  func main() {
      fmt.Println(math.Pi) // math.pi로는 패키지 참조 불가능
  }
  ```



### 07. 함수_1

- 함수는 매개변수 (인자)를 가질 수 있음

- 매개변수의 타입은 뒤에 명시

  ```go
  package main
  
  import "fmt"
  
  func add(a int, b int) int {
      return a + b
  }
  
  func main() {
      fmt.Println(add(42,13))
  }
  ```



### 08. 함수_2

- 두 개 이상의 매개변수가 같은 타입일 때, 같은 타입을 취하는 마지막 매개변수에만 타입을 명시하고 나머지는 생략 가능

  ```go
  (a int, b int)
  ```

  ```go
  (a, b int)
  ```

  ```go
  package main
  
  import "fmt"
  
  func add(a, b int) int {
      return a + b
  }
  
  func main() {
      fmt.Println(add(42,13))
  }
  ```



### 09. 여러 개의 결과 (Multiple results)

- 하나의 함수에 여러 개의 결과 반환 가능

  ```go
  package main
  
  import "fmt"
  
  func swap(x, y string) (string, string) {
      return y, x
  }
  
  func main() {
      a, b := swap{"hello","world"}
      fmt.Println(a,b) // a = world, b = hello
  }
  ```

  

### 10. 이름이 정해진 결과

- Go 에서 함수는 여러 개의 결과를 반환할 수 있는데, 리턴값에 이름을 부여하면, 변수처럼 사용할 수도 있음

- 결과에 이름을 붙이면, 반환값을 지정하지 않은 return 문장으로 결과의 현재 값을 알아서 반환함

  ```go
  package main
  
  import "fmt"
  
  // return 값에 이름을 붙였으므로, return 문장을 통해 x,y의 현재 값을 알아서 반환
  func split(sum int) (x, y int) {
      x = sum * 4 / 9
      y = sum - x
  }
  
  func main() {
      fmt.Println(split(17))
      // Output:
      // 7 10
  }
  ```



### 11. 변수

- 변수 선언을 위해 `var` 이용

- 타입은 끝에 명시

  ```go
  package main
  
  import "fmt"
  
  var x, y, z int
  var c, python, java bool
  
  func main() {
      fmt.Println(x,y,z,c,python,java)
  }
  ```



### 14. 상수 (Constants)

- `const` 키워드와 함께 변수처럼 선언

- 상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있음

  ```go
  package main
  
  import "fmt"
  
  const Pi = 3.14
  
  func main() {
      const World = "안녕"
      fmt.Println("Hello", World)
      fmt.Println("Happy", Pi, "Day")
      
      const Truth = true
      fmt.Println("Go rules?", Truth)
  }
  ```



### 24. 기본 자료형

```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8의 다른 이름(alias:별칭)
rune // int32의 다른 이름, 유니코드 코드 포인트 값을 표현
float32 float64
complex64 complex128
```

```go
package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}
```



### 25. 구조체 (Structs)

- `struct` 는 필드(데이터)들의 조합

  - type 선언으로 struct의 이름을 지정할 수 있음

  ```go
  package main
  
  import "fmt"
  
  type Vertex struct {
      X int
      Y int
  }
  
  func main() {
      fmt.Println(Vertex{1,2})
      // Output:
      // {1 2}
  }
  ```



### 26. 구조체 필드

- 구조체에 속한 필드(데이터)는 dot으로 접근

  ```go
  package main
  
  import "fmt"
  
  type Vertex struct {
      X int
      Y int
  }
  
  func main() {
      v := Vertex{1,2}
      v.X = 4
      fmt.Println(v.X)
      // Output:
      // 4
  }
  ```



### 27. 포인터 (Pointers)

- Go 에는 포인터가 있지만 포인터 연산은 불가능

- 구조체 변수는 구조체 포인터를 이용해서 접근할 수 있음

- 포인터를 이용하는 간접적인 접근은 실제 구조체에도 영향을 미침

  ```go
  package main
  
  import "fmt"
  
  type Vertex struct {
      X int
      Y int
  }
  
  func main() {
      p := Vertex{1,2}
      q := &p
      q.X = 1e9
      fmt.Println(p)
      // Output:
      // 1000000000 2
  }
  ```

### 28. 구조체 리터럴 (Struct Literals)

- 구조체 리터럴 : 필드와 값을 나열하여 구조체를 새로 할당하는 방법

- 원하는 필드를 `{Name:value}` 구문을 통해 할당 할 수 있음

  - 필드의 순서는 무관

- 접두어 `&` 을 사용하면 구조체 리터럴에 대한 포인터를 생성할 수 있음

  ```go
  package main
  
  import "fmt"
  
  type Vertex struct {
      X, Y int
  }
  
  var (
  	p = Vertex{1, 2} // Vertex 타입
      q = &Vertex{1, 2} // *Vertex 타입
      r = Vertex{X: 1} // Y:0
      s = Vertex{} // X:0 & Y:0
  )
  
  func main() {
      fmt.Println(p,q,r,s)
      // Output:
      // {1 2} &{1 2} {1 0} {0 0}
  }
  ```



### 29. new 함수

- `new(T)` 는 모든 필드가 0이 할당된 T 타입의 포인터를 반환

  - zero value : 숫자형에서는 0, 참조형에서는 nil을 의미

  ```go
  var t *T = new(T)
  ```

  or

  ```go
  t := new(T) // t는 T에서 반환된 포인터를 가짐
  ```

  

  ```go
  package main
  
  import "fmt"
  
  type Vertex struct {
      X, Y int
  }
  
  func main() {
      v := new(Vertex) // Vertex에서 반환된 포인터가 v에 들어감
  	fmt.Println(v)
  	v.X, v.Y = 11, 9
  	fmt.Println(v)
      // Output:
      // &{0 0}
      // &{11 9}
  }
  ```



### 30. 슬라이스



### 42. 함수 값 (Function values)

- 함수도 값이다

  ```go
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  func main() {
      hypot := func(x, y float64) float64 {
          return math.Sqrt(x*x + y*y)
      }
      
      fmt.Println(hypot(3,4))
      // Output:
      // 5
  }
  ```



### 43. 함수 클로져 (Function closures)

- 함수는 클로져 (full closures)

  ```go
  package main
  
  import "fmt"
  
  // adder 함수는 클로져를 반환
  func adder() func(int) int {
      sum := 0
      return func(x int) int {
          sum += x
          return sum
      }
  }
  
  // 각각의 클로저는 자신만의 sum 변수를 가짐
  func main() {
      pos, neg := adder(), adder()
      for i := 0, i < 10, i++ {
          fmt.Println(
              pos(i),
              neg(-2*i),
          )
      }
      // Output:
      /*  0  0
      	1  -2
      	3  -6
      	6  -12
      	10 -20
      	15 -30
      	21 -42
      	28 -56
      	36 -72
      	45 -90
      */
      
  }
  ```



### 44. 연습: 피보나치 클로져

```go
package main

import "fmt"

func fibonacci() func() int {
    f1, f2 := 0, 1
    return func() int {
        res := f1 + f2
        f1 ,f2 = f2, res
        return res
    }
}

func main() {
    f := fibonacci()
    for i := 0, i < 10, i++ {
        fmt.Println(f())
    }
}
```



### 45. Switch 스위치

- Go 언어에서는 case의 코드 실행을 마치면 알아서 break를 함

- case를 종료하고 싶지않다면, 끝에 `fallthrough` 를 추가

  ```go
  package main
  
  import "fmt"
  
  func main() {
      fmt.Print("Go runs on ")
      switch os := runtime.GOOS; os {
          case "darwin":
          	fmt.Println("OS X")
          case "linux":
          	fmt.Println("Linux")
          default:
          	fmt.Printf("%s", os)
      }
  }
  ```

  - 스스로 break를 하지 않는 case문

    ```go
    package main
    
    import (
        "fmt"
    )
    
    func checkWeather(degree int) {
        switch {
        	case degree > 32:
            	fmt.Println("Very Hot!")
        	case degree > 25:
            	fmt.Println("Hot!")
        	case degree > 20:
            	fmt.Println("Good!")
            case degree > 10:
            	fmt.Println("Cool!")
        	case degree > -5:
            	fmt.Println("Cold!")
        	default:
            	fmt.Println("?")
        }
    }
    
    func checkWeather_fallthrough(degree int) {
        switch {
        	case degree > 32:
            	fmt.Println("Very Hot!")
            	fallthrough
        	case degree > 25:
            	fmt.Println("Hot!")
            	fallthrough
        	case degree > 20:
            	fmt.Println("Good!")
            	fallthrough
            case degree > 10:
            	fmt.Println("Cool!")
            	fallthrough
        	case degree > -5:
            	fmt.Println("Cold!")
        	default:
            	fmt.Println("?")
        }
    }
    
    func main() {
        degree := 35
        checkWeather(degree)
        checkWeather_fallthrough(degree)
        // Output:
        // Very Hot!
        // Very Hot!
        // Hot!
        // Good!
        // Cool!
        // Cold!
    }
    ```

### 47. 조건을 생략한 스위치



### 50. Methods 메소드

- Go 에는 클래스가 없음

  - struct가 클래스를 대신함

- 대신, 메소드를 구조체에 붙일 수 있음

- 리시버 인자를 갖는 함수로 메소드를 정의하면 됨!!

  - 리시버로 구조체와 함수를 연결하여 메서드를 구현!!

  ```go
  func (리시버 인자) 함수이름 리턴타입
  ```

  - 메소드 호출로 값의 복사가 아닌, 자기 자신의 값을 바꾸려면 인자로 **포인터 리시버** 를 받아야 함

  ```go
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type Vertex struct {
      X, Y float64
  }
  
  // 리시버 인자로 v *Vertex, 즉 포인터 리시버 인자를 받음
  func (v *Vertex) Abs() float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func main() {
      v := &Vertex{3, 4}
      fmt.Println(v.Abs())
      // Output:
      // 5
  }
  ```



## 51. 메소드_2

- 메소드는 구조체 뿐만 아니라 아무 타입에나 붙일 수 있음

- 다른 패키지에 있는 타입이나 기본 타입들에 메소드를 붙이는 것은 불가능

  ```go
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type MyFloat float64 // 기본형 타입도 메소드를 만들 수 있음
  
  func (f MyFloat) Abs() float64 {
      if f < 0 {
          return float64(-f)
      }
      return float64(f)
  }
  
  func main() {
      f := MyFloat(-math.Sqrt2)
      fmt.Println(f.Abs())
      // Output:
      // 1.4142135623730951
  }
  ```



## 52. 포인터 리시버를 가지는 메소드

- 포인터를 인자로 받는 함수 : value (값) 을 전달 받을 수 없음

- 값을 인자로 받는 함수 : 포인터를 전달 받을 수 없음

  ```go
  var v Vertex
  ScaleFunc(v) // 컴파일 에러
  ScaleFunc(&v) // 정상
  ```

- 포인터 리시버 : 값과 포인터 모두 접근 가능

  ```go
  var v Vertex
  v.Scale(5) // value로 접근 가능
  p := &v
  p.Scale(10) // 포인터로도 접근 가능
  ```

  

- 메소드는 이름이 있는 타입 또는 이름이 있는 타입의 포인터와 연결할 수 있음

- 포인터 리시버를 이용하는 이유?

  - 메소드가 호출될 때 마다 값이 복사되는 것 (큰 구조체 타입인 경우 값이 복사되는 것은 비효율적이기 때문)을 방지하기 위함
  - 메소드에서 리시버 포인터가 가리키는 값을 수정하기 위함

  ```go
  package main
  
  import (
  	"fmt"
  	"math"
  )
  
  type Vertex struct {
      X, Y float64
  }
  
  // 포인터 리시버를 사용한 함수 - 메소드에서 리시버 포인터가 가르키는 값을 수정
  // 리시버를 Vertex 타입 (값 타입)으로 바꾸면 Vertex 타입의 복사본에 작업을 하므로,
  // 원래의 구조체 값은 변경되지 않음
  func (v *Vertex) Scale(f float64) {
      v.X = v.X * f
      v.Y = v.Y * f
  }
  
  // 포인터 리시버를 사용한 함수 - 메소드가 호출될 때 마다 값이 복사되는 것을 방지
  // 큰 구조체 타입인 경우 값이 복사되는 것이 비효율적이기 때문
  // Abs() 함수는 리시버를 Vertex 타입으로 바꿔도 변함없이 동작
  func (v *Vertex) Abs() float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func main() {
      v := &Vertex{3, 4}
      v.Scale(5)
      fmt.Println(v, v.Abs())
      // Output:
      // &{15 20} 25
  }
  ```



### 53. Interface (인터페이스)

- 메소드의 집합으로 인터페이스 정의
- 메소드들의 구현되어 있는 타입의 값은 모두 인터페이스 타입의 값이 될 수 있음

```go
package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser
    a = v  // a Vertex, does NOT
    // implement Abser

    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```



### 63. 고루틴 (Goroutines)

- Go 런타임에 의해 관리되는 경량 쓰레드

