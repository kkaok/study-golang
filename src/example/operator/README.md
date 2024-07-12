Go 언어는 다양한 연산자를 제공하여 변수와 값 간의 연산을 수행할 수 있습니다. 연산자는 크게 산술 연산자, 비교 연산자, 논리 연산자, 비트 연산자, 할당 연산자, 기타 연산자로 나눌 수 있습니다.

# 산술 연산자 (Arithmetic Operators)
산술 연산자는 수학적 계산을 수행합니다.

|연산자|설명|예시|
|------|---|---|
| + |덧셈|a + b|
| - |뺄셈|a - b|
| \* |곱셉|a \* b|
| / |나눗셈|a / b|
| % |나머지|a % b|
| ++ |증가 (단항 연산자)|a++ (a = a + 1)|
| -- |감소 (단항 연산자)|a-- (a = a - 1)|

# 비교 연산자 (Comparison Operators)
비교 연산자는 두 값을 비교하여 논리적 값을 반환합니다.

|연산자|설명|예시|
|------|---|---|
|==|같음|a == b|
|!=|같지 않음|a != b|
|>|큼|a > b|
|<|작음|a < b|
|>=|크거나 같음|a >= b|
|<=|작거나 같음|a <= b|

# 논리 연산자 (Logical Operators)
논리 연산자는 불리언 논리 연산을 수행합니다.

|연산자|설명|예시|
|------|---|---|
|&&|논리 AND|a && b|
|\|\| | |a \|\| b|
|!|논리 NOT (단항 연산자)|!a|

# 비트 연산자 (Bitwise Operators)
비트 연산자는 정수형 값을 비트 단위로 연산합니다.

|연산자|설명|예시|
|------|---|---|
&   비트 AND  a & b
`   `   비트 OR
^   비트 XOR  a ^ b
&^  비트 AND NOT  a &^ b
<<  왼쪽 시프트  a << 1
\>\>  오른쪽 시프트 a \>\> 1

5. 할당 연산자 (Assignment Operators)
할당 연산자는 변수에 값을 할당합니다.

|연산자|설명|예시|
|------|---|---|
|=|할당|a = b|
|+=|덧셈 후 할당|a += b (a = a + b)|
|-=|뺄셈 후 할당|a -= b (a = a - b)|
|\*=|곱셈 후 할당|a \*= b (a = a * b)|
|/=|나눗셈 후 할당|a /= b (a = a / b)|
|%=|나머지 후 할당|a %= b (a = a % b)|
|&=|비트 AND 후 할당|a &= b (a = a & b)|
|^=|비트 XOR 후 할당|a ^= b (a = a ^ b)|
|<<=|왼쪽 시프트 후 할당|a <<= 1 (a = a << 1)|
|\>\>=|오른쪽 시프트 후 할당|a >>= 1 (a = a >> 1)|
|&^=|비트 AND NOT 후 할당|a &^= b (a = a &^ b)|

# 기타 연산자 (Other Operators)
기타 연산자는 주소와 포인터 연산 등에 사용됩니다.

|연산자|설명|예시|
|------|---|---|
|&|주소|&a|
|\*|포인터|*p|

# 예제 코드
다양한 연산자를 사용하는 예제 코드를 보겠습니다:

```
package main

import (
    "fmt"
)

func main() {
    // 산술 연산자
    a := 10
    b := 5
    fmt.Println("a + b =", a+b)
    fmt.Println("a - b =", a-b)
    fmt.Println("a * b =", a*b)
    fmt.Println("a / b =", a/b)
    fmt.Println("a % b =", a%b)

    // 증가 및 감소 연산자
    a++
    fmt.Println("a++ =", a)
    b--
    fmt.Println("b-- =", b)

    // 비교 연산자
    fmt.Println("a == b =", a == b)
    fmt.Println("a != b =", a != b)
    fmt.Println("a > b =", a > b)
    fmt.Println("a < b =", a < b)
    fmt.Println("a >= b =", a >= b)
    fmt.Println("a <= b =", a <= b)

    // 논리 연산자
    fmt.Println("true && false =", true && false)
    fmt.Println("true || false =", true || false)
    fmt.Println("!true =", !true)

    // 비트 연산자
    c := 6                        // 0110 in binary
    d := 11                       // 1011 in binary
    fmt.Println("c & d =", c&d)   // 0010 in binary, which is 2
    fmt.Println("c | d =", c|d)   // 1111 in binary, which is 15
    fmt.Println("c ^ d =", c^d)   // 1101 in binary, which is 13
    fmt.Println("c &^ d =", c&^d) // 0100 in binary, which is 4

    // 시프트 연산자
    e := 8                        // 1000 in binary
    fmt.Println("e << 1 =", e<<1) // 10000 in binary, which is 16
    fmt.Println("e >> 1 =", e>>1) // 0100 in binary, which is 4

    // 할당 연산자
    f := 5
    f += 3
    fmt.Println("f += 3 =", f)
    f -= 2
    fmt.Println("f -= 2 =", f)
    f *= 2
    fmt.Println("f *= 2 =", f)
    f /= 3
    fmt.Println("f /= 3 =", f)
    f %= 3
    fmt.Println("f %= 3 =", f)

    // 주소 및 포인터 연산자
    g := 42
    var p *int = &g
    fmt.Println("g의 주소:", p)
    fmt.Println("p가 가리키는 값:", *p)
}
```

결과

```
a + b = 15
a - b = 5
a * b = 50
a / b = 2
a % b = 0
a++ = 11
b-- = 4
a == b = false
a != b = true
a > b = true
a < b = false
a >= b = true
a <= b = false
true && false = false
true || false = true
!true = false
c & d = 2
c | d = 15
c ^ d = 13
c &^ d = 4
e << 1 = 16
e >> 1 = 4
f += 3 = 8
f -= 2 = 6
f *= 2 = 12
f /= 3 = 4
f %= 3 = 1
g의 주소: 0x140000a6030
p가 가리키는 값: 42
```

이 코드는 다양한 연산자를 사용하여 값을 계산하고 출력하는 예제입니다. 각 연산자의 동작 방식을 이해하는 데 도움이 될 것입니다.