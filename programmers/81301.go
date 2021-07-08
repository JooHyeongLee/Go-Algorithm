/**
문제 설명
img1.png

네오와 프로도가 숫자놀이를 하고 있습니다. 네오가 프로도에게 숫자를 건넬 때 일부 자릿수를 영단어로 바꾼 카드를 건네주면 프로도는 원래 숫자를 찾는 게임입니다.

다음은 숫자의 일부 자릿수를 영단어로 바꾸는 예시입니다.

1478 → "one4seveneight"
234567 → "23four5six7"
10203 → "1zerotwozero3"
이렇게 숫자의 일부 자릿수가 영단어로 바뀌어졌거나, 혹은 바뀌지 않고 그대로인 문자열 s가 매개변수로 주어집니다. s가 의미하는 원래 숫자를 return 하도록 solution 함수를 완성해주세요.

참고로 각 숫자에 대응되는 영단어는 다음 표와 같습니다.

1 ≤ s의 길이 ≤ 50
s가 "zero" 또는 "0"으로 시작하는 경우는 주어지지 않습니다.
return 값이 1 이상 2,000,000,000 이하의 정수가 되는 올바른 입력만 s로 주어집니다.

참고로 각 숫자에 대응되는 영단어는 다음 표와 같습니다.

숫자	영단어
0	zero
1	one
2	two
3	three
4	four
5	five
6	six
7	seven
8	eight
9	nine
 */

package programmers

import (
	"fmt"
	"regexp"
	"strconv"
)

func Solution81301(s string) int {
	result := s
	m1 := regexp.MustCompile(`one`)
	m2 := regexp.MustCompile(`two`)
	m3 := regexp.MustCompile(`three`)
	m4 := regexp.MustCompile(`four`)
	m5 := regexp.MustCompile(`five`)
	m6 := regexp.MustCompile(`six`)
	m7 := regexp.MustCompile(`seven`)
	m8 := regexp.MustCompile(`eight`)
	m9 := regexp.MustCompile(`nine`)

	numCnt := regexp.MustCompile("[0-9]+")
	strCnt := regexp.MustCompile("[a-z]+")
	fmt.Println(numCnt.FindAllString(s, -1))
	fmt.Println(strCnt.FindAllString(s, -1))

	if len(m1.FindString(result)) > 0 {
		result = m1.ReplaceAllString(result, "1")
	}
	if len(m2.FindString(result)) > 0 {
		result = m2.ReplaceAllString(result, "2")
	}
	if len(m3.FindString(result)) > 0 {
		result = m3.ReplaceAllString(result, "3")
	}
	if len(m4.FindString(result)) > 0 {
		result = m4.ReplaceAllString(result, "4")
	}
	if len(m5.FindString(result)) > 0 {
		result = m5.ReplaceAllString(result, "5")
	}
	if len(m6.FindString(result)) > 0 {
		result = m6.ReplaceAllString(result, "6")
	}
	if len(m7.FindString(result)) > 0 {
		result = m7.ReplaceAllString(result, "7")
	}
	if len(m8.FindString(result)) > 0 {
		result = m8.ReplaceAllString(result, "8")
	}
	if len(m9.FindString(result)) > 0 {
		result = m9.ReplaceAllString(result, "9")
	}
	var test, _ = strconv.Atoi(result)
	return test
}
