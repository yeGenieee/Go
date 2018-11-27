package main

import "fmt"

/**
 * 배열
 */
func Example_array() {
    fruits := [3]string{"사과", "바나나", "토마토"}
    for _, fruit := range fruits {
        fmt.Printf("%s는 맛있다.\n", fruit)
    }
    // Output:
}

/**
 * 슬라이스
 */
func Example_slice() {
	// var fruits []string : 빈 문자열 슬라이스 만들기

	fruits := make([]string, 3) // 빈 스트링 3개를 가지고 있는 슬라이스 만들기, make로 크기를 잡고 만든 슬라이스에는 해당 자료형의 기본값이 들어감

	fruits[0] = "사과"
	fruits[1] = "바나나"
	fruits[2] = "토마토"

	fruits = append(fruits, "포도") // 슬라이스 덧붙이기 ("포도" 추가)
	fruits = append(fruits, "포도", "딸기") // 여러 개 덧붙이기 (append 함수는 가변 인자를 받음)

	// Output:
}

/**
 * 슬라이싱
 */
func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5} // 정수형 슬라이스
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	// Output:
}

/**
 * 두 슬라이스 덧붙이기
 */
func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...) // 이어붙이기 (f1)
	f4 := append(f1[:2], f2...) // 토마토를 제외하고 이어붙이기

	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	
	// Output:
}

/**
 * 슬라이스 용량
 */
func Example_sliceCap() {
	nums := []int{1,2,3,4,5}

	fmt.Println(nums)
	fmt.Println("len:",len(nums))
	fmt.Println("cap:",cap(nums))
	fmt.Println()

	sliced1 := nums[:3] // 뒤에 2개를 잘라냄 [1,2,3]
	fmt.Println(sliced1)
	fmt.Println("len:",len(sliced1)) // 길이는 2만큼 줄어들게 되어 3이됨
	fmt.Println("cap:",cap(sliced1)) // 용량은 여전히 5 (뒤에 공간이 있으므로)
	fmt.Println()

	sliced2 := nums[2:] // 앞에 2개를 잘라냄 [3,4,5]
	fmt.Println(sliced2)
	fmt.Println("len:",len(sliced2)) // 길이는 3
	fmt.Println("cap:",cap(sliced2)) // 용량은 뒤에 공간이 없으므로 3이됨
	fmt.Println()

	sliced3 := sliced1[:4] // sliced3은 이미 잘라진 sliced1의 뒷 공간을 살려낸 것! , nums에서 잘라낸 것이 아님
	fmt.Println(sliced3)
	fmt.Println("len:",len(sliced3)) // 길이는 4
	fmt.Println("cap:",cap(sliced3)) // 용량은 5
	fmt.Println()

	nums[2] = 100 // nums에서 잘려진 슬라이스는 모두 동일한 메모리를 보고 있음, 따라서, nums[2]만 수정해도 sliced1, sliced2, sliced3까지 변경됨
	fmt.Println(nums, sliced1, sliced2, sliced3) 

	// Output:
}

/**
 * 슬라이스 복사 : 반복문을 통한 복사
 */
func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))

	for i := range src {
		dest[i] = src[i]
	}

	fmt.Println(dest)
	
	// Output:
}

/**
 * 슬라이스 복사 : copy를 통한 복사
 */
func Example_sliceCopy2() {
    src := []int{30, 20, 50, 10, 40}
    dest := make([]int, len(src)) // copy를 통해 슬라이스 전체 복사하기 위함
    
	if n := copy(dest, src); n != len(src) {
		fmt.Println("복사가 덜 됐습니다.")
	}
	
	fmt.Println(dest)

	// Output:
}

/**
 * 슬라이스 복사 : copy를 통한 복사, copy로 슬라이스 전체 복사하기
 */
func Example_sliceCopy3() {
	src := []int{30,20,50,10,40}
	dest := make([]int, len(src))

	copy(dest, src)

	fmt.Println(dest)

	// Output:
}

/**
 * 슬라이스 복사 : append를 통한 복사
 */
 func Example_sliceCopy4() {
	src := []int{30,20,50,10,40}
	dest := append([]int(nil), src...) // 빈 슬라이스에 src에 있는 모든 원소를 복사

	fmt.Println(dest)

	// Output:
}

/**
 *  슬라이스 삽입 : 맨 앞이나 중간에 끼워 넣기
 */
func Example_sliceInsert(){
	// a 슬라이스의 i 번째 원소로 x를 삽입하는 방법
	a := []int{1,2,3,4,5,6,7,8}
	i := 5
	x := 9

	fmt.Println(a)

	a = append(a[:i+1], a[i:]...) // 맨 뒤에 끼워넣기는 불가능 : a[i+1]가 범위가 넘어가기 때문

	a[i] = x
	
	fmt.Println(a)

	// Output:

}

/**
 *  슬라이스  : 맨 뒤에 덧붙이기
 */
 func Example_sliceInsert2(){
	// a 슬라이스의 i 번째 원소로 x를 삽입하는 방법
	a := []int{1,2,3,4,5,6,7,8}
	i := 5
	x := 9

	fmt.Println(a)

	if i < len(a) {
		a = append(a[:i+1], a[i:]...)
		a[i] = x
	} else {
		a = append(a, x)
	}
	
	fmt.Println(a)

	// Output:

}

/**
 *  슬라이스  : copy를 이용하여 덧붙이기
 */
 func Example_sliceInsert3(){
	// a 슬라이스의 i 번째 원소로 x를 삽입하는 방법
	a := []int{1,2,3,4,5,6,7,8}
	i := 5
	x := 9

	fmt.Println(a)

	a = append(a, x)
	copy(a[i+1:], a[i:])
	a[i] = x 
	
	fmt.Println(a)

	// Output:

}