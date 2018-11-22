

# Go 도구 사용하기

- godoc
- Oracle
- Vet
- Fix
- Test

## godoc

- Go 프로그램의 문서를 볼 수 있는 도구

- 주석을 문서화 할 수 있음

- 다른 패키지의 문서도 볼 수 있음

  - command line 으로 보는 방법

    ```go
    godoc fmt
    godoc ftm Printf
    ```

    <결과>

    ```shell
    PACKAGE DOCUMENTATION
    
    package fmt
        import "fmt"
    
        Package fmt implements formatted I/O with functions analogous to C's
        printf and scanf. The format 'verbs' are derived from C's but are
        simpler.
    
    
        Printing
    
        The verbs:
    
        General:
    
    	%v	the value in a default format
    		when printing structs, the plus flag (%+v) adds field names
    	%#v	a Go-syntax representation of the value
    	%T	a Go-syntax representation of the type of the value
    	%%	a literal percent sign; consumes no value
    
        Boolean:
    
    	%t	the word true or false
    
        Integer:
    
    	%b	base 2
    	%c	the character represented by the corresponding Unicode code point
    	%d	base 10
    	%o	base 8
    	%q	a single-quoted character literal safely escaped with Go syntax.
    	%x	base 16, with lower-case letters for a-f
    	%X	base 16, with upper-case letters for A-F
    	%U	Unicode format: U+1234; same as "U+%04X"
    ...(후략)...
    ```

  - 웹 서버로 만드는 방법

    ```
    godoc -http=:6060
    ```

    - -http옵션을 통해 웹 서버 돌릴 수 있음

    - 브라우저에서 localhost:6060 으로 접속

    - < 결과 >

      - Go 홈페이지와 비슷한 화면![image-20181121143656902](/Users/user/Desktop/스크린샷 2018-11-21 오후 2.36.34.png)

    - http://localhost:6060/pkg/github.com/yeGenieee/gogo/seq/

      - 패키지 seq에 대한 문서 결과 (코드 작성 시 주석을 달아놨기 때문에 문서 출력 가능)
      - <결과>

      ![image-20181121143656902](/Users/user/Desktop/스크린샷 2018-11-21 오후 2.39.57.png)





    - 오픈 소스 Go 코드에 대해 문서들을 모두 모아놓은 사이트

      - **https://godoc.org**
      - 알려진 Go 패키지들의 문서를 검색하고 볼 수 있음


## Oracle

- 소스 코드에 대해 여러 가지를 물어볼 수 있는 강력한 도구

  - $GOPATH로 이동하여 다음 command 수행해보기

    ```shell
    oracle -pos=src/github.com/yeGenieee/gogo/seq/seq.go:#134 describe git hub.com/yeGenieee/seq
    ```

    - 134 : 소스코드에서 func Fib가 나오는 부분의 F 바이트 위치


## Vet

- 소스 코드 검사 도구

  ```go
  go tool vet github.com/yeGenieee/gogo/seq
  go tool vet *.go
  ```


## Fix

- 이미 변경된 옛 API 호출 등을 자동으로 고쳐주는 도구

   ```go
  go tool fix github.com/yeGenieee/gogo/seq
   ```


## Test

- 테스트를 수행하는 도구

  ```go
  go test github.com/yeGenieee/gogo/seq
  ```

  ```go
  go test {패키지 경로}
  ```



### Go

- go get : 오픈 소스로 공개 되어 있는 패키지들을 가져올 수 있음
- 편집기에 gofmt을 설정하여 자동으로 형식을 맞추는 것을 권장
- goimport도 설정하여 import도 자동으로 할 수 있게 설정 가능