Go 언어에는 다양한 데이터 타입이 있습니다. 데이터 타입은 변수에 저장될 수 있는 값의 종류와 그 값이 수행할 수 있는 연산을 정의합니다. Go에서 주로 사용되는 데이터 타입을 설명하겠습니다.

# 기본 데이터 타입 (Basic Types)
## 부울 (Boolean)
- 타입: bool
- 값: true 또는 false

    var isGoAwesome bool = true

## 정수 (Integer)
- 타입: int, int8, int16, int32, int64
- 부호 없는 정수 타입: uint, uint8 (byte), uint16, uint32, uint64
- 특별 타입: uintptr (포인터 크기를 저장할 수 있는 정수형)

    var age int = 30
    var height uint = 180
    var b byte = 255

## 부동 소수점 (Floating Point)
- 타입: float32, float64

    var pi float64 = 3.14

## 복소수 (Complex Numbers)
- 타입: complex64, complex128

    var c complex128 = complex(1, 2) // 1 + 2i

## 문자열 (String)
- 타입: string

    var greeting string = "Hello, World!"

## 바이트와 루너 (Byte and Rune)
- 타입: byte (uint8의 별칭), rune (int32의 별칭, 주로 유니코드 코드 포인트를 저장하는 데 사용)

    var letter byte = 'a'
    var unicodeChar rune = '한'

## 집합 데이터 타입 (Aggregate Types)
### 배열 (Array)
- 타입: [n]T (n은 배열의 길이, T는 배열 요소의 타입)

    var numbers [5]int = [5]int{1, 2, 3, 4, 5}

### 슬라이스 (Slice)
- 타입: []T (T는 슬라이스 요소의 타입)
- 특징: 배열보다 유연하고 크기가 동적으로 변경 가능

    var scores []int = []int{10, 20, 30}
    scores = append(scores, 40)

### 맵 (Map)
- 타입: map[K]V (K는 키의 타입, V는 값의 타입)

    var person map[string]int = map[string]int{"age": 30, "height": 180}

### 구조체 (Struct)
- 타입: struct

    type Person struct {
        Name string
        Age int
    }
    var p Person = Person{Name: "John", Age: 30}

### 포인터 (Pointer)
- 타입: *T (T는 포인터가 가리키는 값의 타입)
- 특징: 변수의 메모리 주소를 저장

    var x int = 10
    var p *int = &x

### 인터페이스 (Interface)
- 타입: interface{} (빈 인터페이스는 모든 타입을 저장할 수 있음)

    type Speaker interface {
        Speak() string
    }
    type Person struct {
        Name string
    }
    func (p Person) Speak() string {
        return "Hello, my name is " + p.Name
    }
    var s Speaker = Person{Name: "John"}

### 함수 타입 (Function Types)
- 타입: func

    func add(a int, b int) int {
        return a + b
    }
    var sum func(int, int) int = add

## 예제 코드
아래는 위의 다양한 데이터 타입을 사용하는 예제 코드입니다:

    package main
    import (
        "fmt"
    )
    // PersonWithName 구조체 정의
    type PersonWithName struct {
        Name string
    }
    // Speak 메서드 정의
    func (p PersonWithName) Speak() string {
        return "Hello, my name is " + p.Name
    }
    func main() {
        // 부울 타입
        var isGoAwesome bool = true
        fmt.Println("isGoAwesome:", isGoAwesome)
        // 정수 타입
        var age int = 30
        var height uint = 180
        var b byte = 255
        fmt.Println("age:", age, "height:", height, "byte:", b)
        // 부동 소수점 타입
        var pi float64 = 3.14
        fmt.Println("pi:", pi)
        // 복소수 타입
        var c complex128 = complex(1, 2)
        fmt.Println("complex number:", c)
        // 문자열 타입
        var greeting string = "Hello, World!"
        fmt.Println("greeting:", greeting)
        // 바이트와 루너 타입
        var letter byte = 'a'
        var unicodeChar rune = '한'
        fmt.Println("letter:", letter, "unicodeChar:", unicodeChar)
        // 배열 타입
        var numbers [5]int = [5]int{1, 2, 3, 4, 5}
        fmt.Println("numbers:", numbers)
        // 슬라이스 타입
        var scores []int = []int{10, 20, 30}
        scores = append(scores, 40)
        fmt.Println("scores:", scores)
        // 맵 타입
        var person map[string]int = map[string]int{"age": 30, "height": 180}
        fmt.Println("person:", person)
        // 구조체 타입
        type Person struct {
            Name string
            Age  int
        }
        var p Person = Person{Name: "John", Age: 30}
        fmt.Println("Person:", p)
        // 포인터 타입
        var x int = 10
        var pPointer *int = &x
        fmt.Println("pointer:", pPointer, "value:", *pPointer)
        // 인터페이스 타입
        type Speaker interface {
            Speak() string
        }
        var s Speaker = PersonWithName{Name: "John"}
        fmt.Println("Speaker:", s.Speak())
        // 함수 타입
        add := func(a int, b int) int {
            return a + b
        }
        var sum func(int, int) int = add
        fmt.Println("sum:", sum(1, 2))
    }

### 결과 print

    go run dataType.go
    isGoAwesome: true
    age: 30 height: 180 byte: 255
    pi: 3.14
    complex number: (1+2i)
    greeting: Hello, World!
    letter: 97 unicodeChar: 54620
    numbers: [1 2 3 4 5]
    scores: [10 20 30 40]
    person: map[age:30 height:180]
    Person: {John 30}
    pointer: 0x14000110078 value: 10
    Speaker: Hello, my name is John
    sum: 3

이 예제는 Go의 다양한 데이터 타입을 보여줍니다. 각 타입은 특정한 용도로 사용되며, 이를 통해 다양한 데이터를 저장하고 조작할 수 있습니다.



### complex 타입 예제

    package main
    import (
        "fmt"
        "math/cmplx"
    )
    func main() {
        // 복소수 선언
        var c1 complex64 = complex(5, 7)       // 5 + 7i
        var c2 complex128 = complex(3.2, -4.5) // 3.2 - 4.5i
        // 복소수의 실수 부분과 허수 부분 추출
        fmt.Println("c1:", c1)
        fmt.Println("실수 부분:", real(c1), "허수 부분:", imag(c1))
        fmt.Println("c2:", c2)
        fmt.Println("실수 부분:", real(c2), "허수 부분:", imag(c2))
        // 복소수 연산
        sum := c1 + complex64(c2)     // 복소수 덧셈
        product := c1 * complex64(c2) // 복소수 곱셈
        fmt.Println("합:", sum)
        fmt.Println("곱:", product)
        // 복소수 함수
        abs := cmplx.Abs(c2)     // 복소수의 절대값
        phase := cmplx.Phase(c2) // 복소수의 위상
        fmt.Println("c2의 절대값:", abs)
        fmt.Println("c2의 위상:", phase)
    }

### 결과

    c1: (5+7i)
    실수 부분: 5 허수 부분: 7
    c2: (3.2-4.5i)
    실수 부분: 3.2 허수 부분: -4.5
    합: (8.2+2.5i)
    곱: (47.5-0.099999666i)
    c2의 절대값: 5.52177507691141
    c2의 위상: -0.9526521008597094
