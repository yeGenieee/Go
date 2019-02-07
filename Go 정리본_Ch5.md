# Ch5 구조체 및 인터페이스

**구조체**는 필드들을 묶어놓은 것!

**인터페이스** 는 메서드의 집합! 구현은 없고 메서드의 형태만 있음



## 5.1 구조체

- 구조체 : 필드들의 모음 혹은 묶음
- 필드 : 명명된 구성 요소들
- 배열은 서로 같은 자료형들을 묶어놓은 것, 구조체는 **서로 다른 자료형**의 자료들도 묶을 수 있음



### 5.1.1 구조체 사용법

- 구조체 자료형으로 된 `task` 변수

```go
var task = struct {
    title string // 제목
 	done bool // 완료 여부
    due *time.Time // 기한 (없을 수도 있으니 포인터 이용, 없으면 nil)
}{"laundry", false, nil}
```



- 자료형만 떼어내어 이름 붙이기

```go
type Task struct {
    title string
    done bool
    due *time.Time
}
```



- 새로운 Task 변수의 선언

  ```go
  var myTask Task
  ```

- Task 변수의 정의

  ```go
  var myTask = Task{"laundry", false, nil}
  ```

- 필드의 값을 넣지 않으면 기본값으로 설정 (0, false, 빈 문자열 등)

  ```go
  var myTask = Task{title:"laundry"}
  ```

- 여러 필드의 값을 정할 때 : **쉼표**로 구분

  ```go
  var myTask = Task{title:"laundry", done:true}
  ```

  - 더 보기 편하도록,

    ```go
    var myTask = Task {
        title: "laundry",
        done: true, // 쉼표가 없으면 줄이 끝나는 것처럼 보여 세미콜론이 붙음 --> 구문 오류 발생
    }
    ```



### 5.1.2 const와 iota

- 확장성을 위한 팁 : bool형을 쓸 곳에 `enum` 형을 사용

  ```go
  type Task struct {
      title string
      status status
      due *time.Time
  }
  ```

  ```go
  type status int
  ```

  - Go 에서는 enum 자료형이 따로 없고, **상수로 정의**하여 사용함

    ```go
    const UNKNOWN status = 0
    const TODO status = 1
    const DONE status = 2
    ```

  - 위의 세 상수는 서로 연관이 있으므로 아래와 같이 묶어 쓸 수 있음

    ```go
    const {
        UNKNOWN status = 0 // 기본값인 0을 dummy값으로 둠
        TODO 	status = 1
        DONE	status = 2
    }
    ```

    

  - 0,1,2 를 순서대로 붙이는 작업 ==> 귀찮을 수 있음 ==> `iota` 사용

    - 순서대로 0, 1, 2 가 붙음

    ```go
    const {
        UNKNOWN status = iota
        TODO	status = iota
        DONE	status = iota
    }
    ```

    

    - `iota` 는 한 번만 써줘도 됨
      - TODO와 DONE의 자료형도 status가 됨

    ```go
    const {
        UNKNOWN status = iota
        TODO
        DONE
    }
    ```

  - iota를 사용하는 예제

    ```go
    type ByteSize float64
    
    const {
        _			= iota // ignore first value
        KB ByteSize = 1 << (10 * iota)
        MB
        GB
        TB
        PB
    	EB
        ZB
        YB
    }
    ```

    - `iota` 가 `0` 부터 시작하는데, 첫 번째 값인 0은 버림
    - 앞에 **상수 이름**으로 **밑줄**을 쓰게 되면 아무 이름을 정의하지 않고 **버리는 것**이 됨 
    - 다음 `iota` 는 `1`이 되는데, `1 << (10 * iota)` 를 썼으므로, 1을 왼쪽으로 10 번 이동시킨, 2^10 = 1024가 됨
    - 그 다음, MB 부터, `1 << (10 * 2)` 로 계산됨

    ```go
    type ByteSize float64
    
    const {
        KB ByteSize = 1 << (10 * (1 + iota))
        MB
        GB
        TB
        PB
        EB
        ZB
        YB
    }
    ```



### 5.1.3 테이블 기반 테스트

- Assertion을 이용한 유닛 테스트 (Python 예시)

  ```python
  def test_fib(self):
  	self.assertEqual(0, fib(0)) # 0번째 수는 0와 같은지 테스트
      self.assertEqual(5, fib(5)) # 5번째 수는 5와 같은지 테스트
      self.assertEqual(8, fib(8)) # 8번째 수는 8과 같은지 테스트
  ```

  - `assertEqual` 함수는 정수형 뿐만 아니라 다른 자료형도 비교하여 테스트

  - Go 언어에서 제공하는 testing 패키지에는 이와 같은 방식으로 테스트할 수 있는 함수들이 제공되지 않음

    ```go
    func assertIntEqual(t *testing.T, a, b int) {
        if a != b {
            t.Errorf("assertion failed: %d != %d", a, b)
        }
    }
    ```

    - 위의 함수는 정수형만 비교할 수 있음 (새로운 자료형이 나타날 때 마다 새로운 함수를 만들어야한다는 단점 존재)

  

- 구조체와 배열을 이용한 테이블 기반 테스트

  ```go
  func TestFib(t *tesing.T) {
      cases := []struct {
          in, want int
      }{
          {0,0},
          {5,5},
          {6,8},
      }
      
      for _, c := range cases {
          got := seq.Fib(c.in)
          if got != c.want {
              t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
          }
      }
  }
  ```





### 5.1.4 구조체 내장

- **여러 자료형의 필드들을 가질 수 있다** : 구조체의 가장 큰 장점!!



- `Deadline` 이라는 자료형 정의

  ```go
  type Deadline time.Time
  
  // OverDue returns true if the deadline is before the current time.
  func (d Deadline) OverDue() bool {
      return time.Time(d).Before(time.Now())
  }
  
  func ExampleDeadline_OverDue() {
      d1 := Deadline(time.Now().Add(-4 * time.Hour)) // 현재 시간에서 4시간 전 마감인 것을 d1에 담음
      d2 := Deadline(time.Now().Add(4 * time.Hour)) // 현재 시간에서 4시간 후에 마감인 것을 d2에 담음
      fmt.Println(d1.OverDue()) // 마감이 지났으므로, true 
      fmt.Println(d2.OverDue()) // 마감이 지나지 않았으므로, false
      // Output:
      // true
      // false
  }
  ```

- 데드라인이 없는 경우 (`Overdue` 수정)

  - Overdue 메서드의 리시버를 포인터로 바꾸기

  ```go
  func (d *Deadline) Overdue() bool {
      return d != nil && time.Time(*d).Before(time.Now())
  }
  ```

  