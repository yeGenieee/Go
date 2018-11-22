package hangul


var (
	start = rune(44032) // '가'의 유니코드 포인트
	end = rune(55204) // '힣' 다음 글자의 유니코드 포인트
)

// HasConsonantSuffix : 받침이 있으면 true
func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false

	for _, r := range s { // s 에 있는 유니코드를 하나씩 r로 꺼냄
		if start <= r && r < end { // 한글인 문자만 검사 (가에서 힣까지의 한글)
			index := int(r-start) // 몇 번째 한글인지 검사 ("가"는 0번째, "각"은 1번째...)
			result = index % numEnds != 0 // 28로 나누어서 나머지가 0이면 받침이 없는 경우, 나머지가 0이 아니면 받침이 있음
		}
	}

	return result
}