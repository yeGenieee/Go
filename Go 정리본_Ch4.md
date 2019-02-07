# Ch4 함수

- 값 넘겨주고 넘겨받기
- 값으로 취급되는 함수
- 메서드
- 활용



  코드의 덩어리를 만든 다음에 그를 호출하고 귀환할 수 있는 구조를 **서브루틴**이라고 부른다. 큰 프로그램을 서브루틴으로 구분하면 코드를 재사용하여 중복된 코드를 줄일 수 있고, 서브 루틴의 내부와 외부를 분리하여 생각할 수 있어서 코드를 추상화하고 단순화할 수 있다.

  Go에서는 이러한 서브루틴을 **함수**라고 부른다.  Go 언어는 `값에 의한 호출 (Call by value)`만을 지원한다. 함수 내에서 넘겨받은 변수값을 변경하더라도 함수 밖의 변수에는 영향을 주지 않는다. 변수에 담겨 있던 값만 넘어오게된다. 따라서, 함수 밖의 변수의 값을 변경하려면 해당 값이 들어 있는 주소값을 넘겨받아서 그 주소에 있는 값을 변경하여 `참조에 의한 호출 (Call by reference)`과 비슷한 효과를 낼 수 있다.



## 4.1 값 넘겨주고 넘겨받기

  서브루틴을 호출할 때, 그 서브루틴에 값을 넘겨주거나 받을 수 있다. 이를 통해, 서브루틴을 다른 값에 대하여 재사용할 수 있고, 결과값을 받아서 이용할 수 있다.



### 4.1.1 값 넘겨주기

  3장의 **텍스트 리스트 읽고 쓰기**에서, ReadFrom 함수에서 *[]string 자료형으로 lines를 받았다. 

```go
func ReadFrom(r io.Reader, lines *[]string) error {}
```

- `[]string` 자료형이 아닌 `*[]string` 자료형으로 받은 이유는 이 `ReadFrom` 함수가 `lines` 변수의 값을 변경하고자 하기 때문
- 슬라이스는 배열에 대한 포인터, 길이 그리고 용량 이렇게 세 값으로 이루어짐
  - 만약, `*[]string`을 쓰지 않고 `[]string`을 이용하여 넘겼다면 이 포인터, 길이, 용량 이렇게 세 개의 값이 넘어가며, 세 개의 값을 담고 있던 변수와는 연관성이 없어짐
  - 즉, `lines []string` 과 같이 받았다면, 함수 내의`lines` 변수가 품고 있는 세 값을 변경해도 원래 변수의 변경과는 무관하다.

```go
func AddOne(nums []int) {
    for i := range nums {
        nums[i]++
    }
}

func ExampleAddOne() {
    n := []int{1,2,3,4}
    AddOne(n)
    fmt.Println(n)
    
    // Output :
    // [2 3 4 5]
}
```

- `AddOne` 함수에 값을 넘겨주었는데, 값을 넘겨주었는데도 호출 뒤에 `n`에 대해 변경이 일어남
- Go에서의 슬라이스가 배열에 대한 포인터, 길이, 용량 이렇게 세 값으로 이루어진 것으로 이 세 값이 넘어간 것이기 때문
  - 정수 1,2,3,4가 넘어간 것이 아님



```go
func ReadFrom(r io.Reader, lines *[]string) error { }
```

- 위의 함수가 `*[]string`과 같은 포인터를 넘겨주는 이유?
  - 함수가 넘겨준 슬라이스의 값을 변경해야 하기 때문
  - 슬라이스 값?
    - 배열 포인터
    - 길이
    - 용량
  - 포인터로 넘어온 값은 `*`을 앞에 붙여 값을 참조할 수 있음
  - 변수 앞에 `&`을 붙이면 해당 변수에 담겨 있는 포인터 값을 얻을 수 있음



### 4.1.2 둘 이상의 반환값

  Go 언어의 함수에서는 **둘 이상의 반환값**을 둘 수 있음

```go
func WriteTo(w io.Writer, lines []string) error { }
```

- 원래는 에러값 하나를 반환하는 함수



- 실제로 몇 바이트를 썼는지 알고 싶은 경우에는?

  - 기존의 C 언어에서는, 바이트 수를 반환하는 함수 하나를 더 작성해야 했음

  ```go
  func WriteTo(w io.Writer, lines []string) (int64, error) { }
  ```

  - return 값을 두 개로 둘 수 있음

  - return 값이 두 개 이상인 경우 

    - **괄호**로 둘러싸서 반환값을 구분해야 함

      

  - 함수 안에서 값을 반환할 때?
    - 쉼표로 구분하여 return

  ```go
  func WriteTo(w io.Writer, lines []string) (int64, error) {
  	...
      return n, err //  쉼표로 구분하여 반환
  }
  ```

  

  - 값을 받을 때?
    - 쉼표로 구분하여 return 값의 수에 맞게 받음

  ```
  _, err := WriteTo(w, lines)
  ```



### 4.1.3 에러 값 주고받기

- 둘 이상의 값을 리턴할 수 있게 되면서 return 값을 이용한 에러 처리가 조금 더 수월해짐

- 기존에는, 에러값을 주고 받기 위해 

  - 정상적이지 않은 경우에, 음수를 반환하여 에러를 의미하도록 규정
  - 호출하는 쪽에서 에러값을 받고 싶은 변수의 포인터나 레퍼런스를 함수로 넘겨 주어 받는 방법

  

- 첫번째, 음수를 반환하여 에러를 의미하는 경우 

  - Go 에서는, 에러가 아닌 정상적인 경우에 이를 활용
    - ex) `strings.Index*`는 문자열에서 원하는 문자열이 나타나는 위치를 돌려주는 함수
    - 인덱스는 0부터 시작하므로, 음수가 나올 수 없는 상황
    - 따라서, 문자열이 발견되지 않으면 `-1`을 return
    - 해당 경우는 에러 상황은 아니고, 정상적인 상황의 한 종류일 뿐



- Go 에서는 **에러**는 **가장 마지막 값**으로 반환

- 새로운 에러를 생성해야 하는 경우

  - `errors.New` 나 `fmt.Errorf` 이용

    ```go
    return errors.New("stringlist.ReadFrom: line is too long")
    ```

    ```go
    return fmt.Errorf("stringlist: too long line at %d", count) // fmt.Errorf로 부가 정보를 추가한 메시지를 리턴
    ```



### 4.1.4  명명된 결과 인자

  위의 함수 예시들은 넘겨받는 인자에는 이름이 있었지만, 돌려주는 값에는 이름 없이 자료형만 나열하였음

- Go에서는 반환하는 값들 역시 이름을 붙여 사용 가능

  - 넘겨받는 인자들 (파라미터)은 넘어온 값들로 초기값이 설정됨
  - 돌려주는 인자들(리턴값)은 기본값으로 초기화됨 (정수면 0, 문자열이라면 빈 문자열로 초기값 설정)

  

- 반환 시, 기존의 방식대로 결과값들을 return 뒤에 쉼표로 구분하여 나열할 수도 있고, 생략하고 `return`만 쓸 수도 있음

- 생략한 경우에는 돌려주는 인자들의 값들이 반환됨

```go
func WriteTo(w io.Writer, lines []string) (n int64, err error) {
    for _, line := range lines {
        var nw int
        nw, err = fmt.Fprintln(w, line)
        n += int64(nw)
        
        if err != nil {
            return
        }
    }
    
    return
}
```

- 위의 예제는 명명된 결과 인자 (Named return parameter)를 사용해도 눈에 띄게 좋은 것은 없음



### 4.1.5 가변 인자

- 넘겨받을 수 있는 인자의 개수가 정해져 있지 않은 함수?

  ```go
  func f(w io.Writer, nums []int) {
      ...
  }
  
  func main() {
      ...
      f(w, []int{x,y,z}) // Go는 값을 넘기므로, 그 값을 슬라이스에 담아서 넘기기
  }
  ```

  

- 슬라이스로 갖고있는 자료를 가변인자를 두고 있는 함수로 넘기는 경우?

  ```go
  lines := []string{"hello", "world", "Go language"}
  WriteTo(w, lines...)
  ```

  

## 4.2 값으로 취급되는 함수

- Go 언어에서, **함수**는 **값으로 변수에 담길 수 있고**, **다른 함수로 넘기거나 돌려받을 수 있음**



### 4.2.1 함수 리터럴

- `add` 함수의 이름이 함수를 담고 있는 변수처럼 보임

  ```go
  func add(a, b int) int {
      return a + b
  }
  ```

  

- 함수 이름을 뺀 함수 (함수 리터럴 (`Function literal`), 익명 함수)

  ```go
  func(a, b int) int {
      return a + b
  }
  ```



- 인자와 반환값이 없는 함수

  ```go
  func printHello() {
      fmt.Println("Hello")
  }
  ```



- 함수의 이름을 없애고, 함수 리터럴만을 남긴 다음 호출

  ```go
  func Example_funcLiteral() {
      func() {
          fmt.Println("Hello")
      }() // 함수 호출
      // Output:
      // Hello
  }
  ```

  - 함수 리터럴

    ```go
    func() {
    	fmt.Println("Hello")
    }
    ```

  - 함수 호출

    ```go
    ()
    ```



-  함수 리터럴을 `printHello` 변수에 담은 후 호출

  ```go
  func Example_funcLiteralVar() {
      printHello := func() {
          fmt.Println("Hello")
      }
      
      printHello() // 호출
      // Output:
      // Hello
  }
  ```





### 4.2.2 고계 함수

- `Higher-order function` : 함수를 넘기고 받는 함수 (더 고차원적인 함수라는 의미)

  

- 원래는 파일에서 한 줄 씩 읽어 슬라이스에 추가하는 `ReadFrom`, 각 사용처마다 함수를 활용하고 싶은 경우 고계 함수 이용

  ```go
  func ReadFrom(r io.Reader, f func(line string)) error {
      scanner := bufio.NewScanner(r)
      for scanner.Scan() {
          f(scanner.Text())
      }
      
      if err := scanner.Err(); err != nil {
          return err
      }
      return nil
  }
  ```



- 함수 리터럴을 넣어 호출

  ```go
  func ExampleReadFrom_Print() {
      r := strings.NewReader("bill\ntom\njane\n")
      err := ReadFrom(r, func(line string){
          fmt.Println("(", line, ")")
      })
      if err != nil {
          fmt.Println(err)
      }
      // Output:
      // (bill)
      // (tom)
      // (jane)
  }
  ```



### 4.2.3 클로저

- **닫힘**이라는 의미
- **외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드**



- 슬라이스에 추가하여 넣는 코드

  ```go
  func ExampleReadFrom_append() {
      r := strings.NewReader("bill\ntom\njane\n")
      var lines []string
      err := ReadFrom(r, func(line string) {
          lines = append(lines, line) // ReadFrom에 넘겨주는 함수 리터럴 안에서 lines 라는 문자열 슬라이스를 사용중
      })
      
      if err != nil {
          fmt.Println(err)
      }
      fmt.Println(lines)
      
      // Output:
      // [bill tom jane]
  }
  ```





### 4.2.4 생성기

- 클로저를 이용한 generator

  ```go
  func NewIntGenerator() func() int {
      var next int
      return func() int {
          next++
          return next
      }
  }
  
  func ExampleNewIntGenerator() {
      gen := NewIntGenerator()
      fmt.Println(gen(), gen(), gen(), gen(), gen())
      fmt.Println(gen(), gen(), gen(), gen(), gen())
      // Output:
      // 1 2 3 4 5
      // 6 7 8 9 10
  }
  ```

  - `NewIntGenerator`는 함수를 반환하는 고계 함수
  - 반환값의 자료형은 `func() int` : `NewIntGenerator`는 정수를 반환하는 함수를 반환하는 함수
  - `NewIntGenerator`가 반환하는 함수는 클로저
  - 반환하는 함수 리터럴이 속해있는 스코프 안에 있는 `next` 변수에 접근하고 있음
  - 이 함수는 `next` 변수와 함께 세트로 묶임



- NewIntGenerator()를 여러번 호출하여 함수 여러 개를 갖는 경우

  ```go
  func ExampleNewIntGenerator_multiple() {
      gen1 := NewIntGenerator()
      gen2 := NewIntGenerator()
      fmt.Println(gen1(), gen1(), gen1())
      fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
      fmt.Println(gen1(), gen1(), gen1(), gen1())
      // Output:
      // 1 2 3
      // 1 2 3 4 5
      // 4 5 6 7
  }
  ```

  - `gen1`과 `gen2`에 묶여있는 `next`는 서로 다르므로, 결과도 다르게 나옴
  - `gen1`과 `gen2` 두 함수의 상태는 **분리**되어 있음



### 4.2.5 명명된 자료형

- 자료형에 새로 이름을 붙일 수 있음

- `int32`의 별칭인 `rune`

- 이를 **명명된 자료형**, **Named Type** 이라고 함

  ```go
  type rune int32
  ```

- 명명되지 않은 자료형 (Unnamed Type) ?

  ```go
  type runes []rune
  type MyFunc func() int
  ```

  - `runes`와 `MyFunc` 는 이름 자체만으로 자료형을 지칭 : **명명된 자료형 (Named Type)**
  -  `[]rune` 과 `func() int` 는 이름만으로 자료형을 지칭하는 것이 아님 : **명명되지 않은 자료형 (Unnamed Type)**
  
  
  ### 4.2.5 명명된 자료형
  
  - 자료형에 새로 **이름을 붙일 수 있음**
  
  - `rune`형이 `int32`의 별칭 : **명명된 자료형 (Named Type)**
  
  ```go
  type rune int32
  ```
  
  
  
  - 명명되지 않은 자료형 (Unnamed Type)
  
  ```go
  type runes []rune
  type MyFunc func() int
  ```
  
  - runes와 MyFunc : 이름 자체만으로 자료형을 지칭 ==> 명명된 자료형
  - []rune 과 func() int : 이름만으로 자료형을 지칭 X ==> 명명되지 않은 자료형
  - 명명된 자료형과 명명되지 않은 자료형 모두에 type 예약어를 사용하여 새 이름을 붙여줄 수 있음
  
  
  
  - 자료형을 붙임으로써 장점?
  
  - 정점과 간선으로 이루어진 그래프를 다루는 코드
  - 각 정점과 간선은 정수형으로 된 ID값을 사용하여 접근하는 코드
  
  ```go
  // Vertex를 Edge로 바꾼 간선 ID Generator
  func NewVertexIDGenerator() func() int{
  var next int
  return func() int {
  next++
  return next
  }
  }
  ```
  
  - 간선의 ID를 받아 새로운 간선을 생성하는 함수 `NewEdge`
  
  ```go
  func NewEdge(eid int) {
  ...
  }
  
  func main() {
  gen := NewVertexIDGenerator()
  gen2 := NewEdgeIDGenerator()
  ...
  e := NewEdge(gen()) // 간선의 ID를 넘겨야하는데, 정점의 ID를 넘기고 있음
  }
  ```
  
  - 해당 버그를 컴파일 시간에 잡지 못하는 이유는, NewEdge를 호출할 때, 간선의 ID를 넘겨야 하는데 정점의 ID를 넘긴 것에 대한 실수를 잡지 못했기 때문 (컴파일 시간에 구분 불가)
  - 해결하려면?
  
  
  
  - 자료형에 서로 다른 이름을 붙임으로써 해결
  
  ```go
  type VertexID int
  type EdgeID int
  ```
  
  - int에 `VertextID` 와 `EdgeID` 두 이름을 붙여 서로 다른 이름이 붙은 자료형으로 만듦
  
  ```go
  func NewVertexIDGenerator() func() VertexID {
  var next VertexID
  return func() VertexID {
  next++
  return next
  }
  }
  ```
  
  - 위 함수가 리턴하는 generator는 호출될 때마다 VertexID 를 반환함
  
  - 만약, 변수 next가 VertexID 형이 아니라 int형이라면 바로 next를 반환할 수 없고, VertexID 로 형변환해주어야 함
  
  - int, VertextID가 둘 다 명명된 자료형(Named Type)이기 때문
  - 서로 다른 Named Type끼리는 호환이 되지 않음
  
  ```go
  func NewVertextIDGenerator() func() VertexID() {
  var next VertexID
  return func() VertexID {
  next++
  return VertexID(next)
  }
  }
  ```
  
  - 어떤 방식을 사용하든지 간에 이 generator가 최종적으로 반환하는 값은 `VertexID형`
  
  ```go
  func NewEdge(eid EdgeID) {
  // ...
  }
  
  func main() {
  gen := NewVertexIDGenerator()
  // ...
  e := NewEdge(gen()) // error
  }
  ```
  
  - 이제는 위와 같은 코드는 오류 발생
  
  
  
  - `[]rune` 과 `runes`와 같이 **명명되지 않은 자료형(Unnamed Type)**과 **명명된 자료형 (Named Type)** 사이에는 표현이 같으면 호환됨
  
  ```go
  type runes []rune
  
  func main() {
  var a []rune = runes{65,66}
  fmt.Println(string(a))
  }
  ```
  
  
  
  - 명명된 자료형을 이용하면 자료형을 하드 코딩하는 것에 비해 **나중에 일괄적으로 해당 자료형의 표현을 변경할 수 있다**는 점이 장점
  
  
  
  ### 4.2.6 명명된 함수형
  
  - Go 언어에서 함수는 일급 시민으로 분류됨
  
  - 따라서, 함수의 자료형도 사용자가 정의할 수 있음
  
  - 두 정수를 넘겨 받아 정수 하나를 반환하는 함수형을 `BinOp` 형으로 정의
  
  ```go
  type BinOp func(int, int) int
  ```
  
  - 명명된 함수형도 자료형 검사를 진행
  
  - BinOp를 받는 함수 OpThreeAndFour
  
  ```go
  func OpThreeAndFour(f BinOp) {
  fmt.Println(f(3,4))
  }
  ```
  
  - 이 함수를 호출할 때, 아래와 같은 방식으로 호출하면?
  
  ```go
  OpThreeAndFour(func (a, b int) int {
  return a + b
  })
  ```
  
  - `OpThreeAndFour` 는 `BinOp` 형을 인자로 받는데, 호출 시에는 그냥 정수 둘을 받고 정수 하나를 반환하는 함수를 넘겨줌 --> 컴파일 오류를 발생시키지 않음
  - 왜? `func (a, b int) int` 자료형이 명명되지 않은 자료형이기 때문
  - 양 쪽 모두 명명된 자료형이 아니면 서로 간에 호환됨
  
  
  
  - 반면, 표현이 같은 함수형이라도,  양쪽 모두 이름이 있는 경우에는 호환 X
  
  ```go
  type BinSub func(int, int) int
  ```
  
  - `BinOp` 와 동일한 표현형으로 된 명명된 함수형
  
  - 
  
  ```go
  type BinOp func(int, int) int
  type BinSub func(int, int) int
  
  func BinOpToBinSub(f BinOp) BinSub {
  var count int
  return func(a, b int) int {
  fmt.Println(f(a,b))
  count++
  return count
  }
  }
  ```
  
  - `BinOpToBinSub` 는 `BinOp` 함수를 받아 `BinSub` 함수를 반환하는 함수
  
  ```go
  func ExampleBinOpToBinSub() {
  sub := BinOpToBinSub(func(a, b int) int {
  return a + b 
  })
  
  sub(5,7)
  sub(5,7)
  count := sub(5,7)
  
  fmt.Println("count:", count)
  // Output:
  // 12
  // 12
  // 12
  // count:3
  }
  ```
  
  - BinOpToBinSub를 호출하여 sub는 BinSub의 자료형이 됨
  
  
  
  - BinOpToBiNSub로 두 번 둘러싸면?
  
  ```go
  sub := BinOpToBinSub(BinOpToBinSub(func (a, b int) int {
  return a + b
  }))
  ```
  
  - 컴파일 시 오류가 발생함
  - 함수 리터럴과 함수형 사이에는 자동으로 형변환이 일어나지만, 명명된 함수형 사이에서는 자동으로 형변환이 일어나지 않음
  
  
  
  ### 4.2.7 인자 고정
  
  - 함수의 인자를 고정하고 싶을 때
  
  ```go
  // Insert 함수는 집합에 val을 추가함
  func Insert(m map[string]int, val string)
  ```
  
  
  
  
  
  
  
  
  
  
  
  
  
  ## 4.3 메서드
  
  - 서브루틴 : 코드의 덩어리를 만든 다음에 그것을 호출하고 반환할 수 있는 구조
  - 메서드 : 리시버가 붙은 서브루틴
  
  
  
  - 자료형 T에 대하여 메서드를 호출할 때, 자료형 T에 대한 리시버가 함수 이름, 즉 메서드 이름 앞에 붙음
  
  - 리시버 부분 : 인자 목록과 같지만 함수 이름 앞에 온다는 것이 다르다
  
  ```go
  func (recv T) MethodName(p1 T1, p2 T2) R1
  ```
  
  - `(recv T)` 부분이 리시버
  
  
  
  ### 4.3.1 단순 자료형 메서드
  
  ```go
  type VertexID int // 정수형으로 명명된 타입
  ```
  
  ```go
  func ExampleVertexID_print() {
  i := VertexID(100)
  fmt.Println(i)
  // Output:
  // 100
  }
  ```
  
  - `i`는 정수형이 아닌 `VertexID` 형이지만 정수형과 마찬가지로 출력됨
  
  
  
  - 리시버가 추가
  
  ```go
  // i가 VertexID 자료형이면, i.String()과 같이 메서드를 호출할 수 있음
  func (id VertexID) String() string {
  return fmt.Sprintf("VertexID(%d)", id)
  }
  ```
  
  ```go
  func ExampleVertexID_String() {
  i := VertexID(100)
  fmt.Println(i)
  // Output:
  // VertexID(100)
  }
  ```
  
  - Go의 인터페이스라는 기능때문에 fmt.Println에 i.String()을 넘겨주지 않았는데도 이런 결과가 나옴!!
  
  
  
  
  
  ### 4.3.2 다중 문자열 집합
  
  - 문자열 다중 집합
  
  - 맵의 키는 집합의 원소인 문자열, 값은 해당 원소가 몇 번 반복되는지를 표시하는 정수
  
  ```go
  type MultiSet map[string]int // 자료형에 이름 붙이기 (이래야 메서드 정의 가능)
  ```
  
  - Insert / Erase / Count / String 메서드 (삽입 / 삭제 / 어떤 원소가 몇 개 들어있는지 확인 / 문자열로 출력)
  
  ```go
  type MultiSet map[string]int
  
  func (m MultiSet) Insert(val string) {
  m[val]++
  }
  
  func (m MultiSet) Erase(val string) {
  if m[val] <= 1 {
  delete(m, val)
  } else {
  m[val]--
  }
  }
  
  func (m MultiSet) Count(val string) int {
  return m[val]
  }
  
  func (m MultiSet) String () string {
  s := "{ "
  for val, count := range m {
  s += strings.Repeat(val+" ", count)
  }
  return s + "}"
  }
  ```
  
  
  
  ### 4.3.3 포인터 리시버
  
  - 포인터 리시버 : 자료형이 포인터형인 리시버
  
  - 포인터로 전달해야 할 경우 포인터 리시버 사용
  
  ```go
  type Graph [][]int
  ```
  
  ```go
  func WriteTo(w io.Writer, adjList [][]int) error
  func ReadFrom(r io.Reader, adjList *[][]int) error
  ```
  
  ```go
  func (adjList Graph) WriteTo(w io.Writer) error
  func (adjList *Graph) ReadFrom(r io.Reader) error
  ```
  
  
  
  
  
  ### 4.3.4 공개 및 비공개
  
  - Go 언어는 public이나 private의 예약어 없이 **식별자 이름의 첫 글자**가 **대문자**인지 / **소문자** 인지에 따라 접근 구분을 함
  - 대문자로 시작하는 메서드, 자료형, 변수, 상수, 함수 : 모듈 밖에서 보임
  - 소문자로 시작하는 메서드, 자료형, 변수, 상수, 함수 : 모듈 밖에서 보이지 않음
  
  
  
  ## 4.4 활용
  
  ### 4.4.1 타이머 활용하기
  
  -  프로그램의 수행을 잠시 멈추고 싶을 때, `time.Sleep` 함수 이용
  
  ```go
  package main
  
  import "fmt"
  
  func CountDown(seconds int) {
  for seconds > 0 {
  fmt.Println(seconds)
  time.Sleep(time.Second)
  seconds--
  }
  }
  
  func main() {
  fmt.Println("Ladies and gentlemen!")
  CountDown(5)
  }
  ```
  
  - time.Sleep을 통해 1초에 한 줄씩 출력되도록3 함
  - **Blocking Timer** : 1초라는 시간을 기다리는 동안 프로그램은 잠시 수행 중단
  - **Non-Blocking Timer** : `time`이라는 모듈안에 있는 `Timer*` 를 이용하면 됨
  
  - **Callback**
  
  - 비동기적 상황에서 어떤 조건이 만족될 때 호출해달라고 요청하는 것
  
  ```go
  time.AfterFunc(5*time.Second, func() {
  // ...
  })
  ```

