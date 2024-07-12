Go 언어의 변수와 상수에 대해 자세히 설명하겠습니다. 변수와 상수는 프로그램의 데이터를 저장하고 관리하는 데 중요한 역할을 합니다.

# 변수 (Variables)
변수는 프로그램에서 데이터를 저장하고 변경할 수 있는 메모리 위치를 나타냅니다. Go 언어에서 변수는 var 키워드를 사용하여 선언합니다.

## 변수 선언 및 초기화
- 단일 변수 선언:여기서 x는 int 타입의 변수를 선언하지만 초기화는 하지 않습니다. 기본값은 0입니다.

    var x int

- 변수 선언과 초기화:여기서 y는 int 타입의 변수로 선언되고 10으로 초기화됩니다.

    var y int = 10

- 타입 생략:여기서 z는 초기화 값에 따라 자동으로 int 타입으로 추론됩니다.

    var z = 20

- 짧은 선언:여기서 a는 := 연산자를 사용하여 선언과 동시에 초기화됩니다. 타입은 자동으로 추론됩니다.

    a := 30

## 여러 변수 선언
여러 변수를 한 번에 선언할 수 있습니다:

- 동일한 타입:

    var b, c, d int

- 다른 타입:

    var (
        e int
        f string
        g bool
    )

- 다양한 초기화:

    var h, i, j = 1, "hello", true

# 상수 (Constants)
상수는 프로그램에서 변하지 않는 값을 나타냅니다. 상수는 const 키워드를 사용하여 선언합니다. 상수는 숫자, 문자열, 부울, 문자를 가질 수 있습니다.

## 상수 선언
- 단일 상수 선언:

    const pi = 3.14

- 타입 지정 상수:

    const width int = 100

- 여러 상수 선언
여러 상수를 한 번에 선언할 수 있습니다:

> 여러 상수 선언:

    const (
        x = 1
        y = "hello"
        z = false
    )

## iota
iota는 상수 값을 자동으로 증가시키는 데 사용됩니다. 상수 그룹 내에서 첫 번째 상수부터 0으로 시작하여 자동으로 1씩 증가합니다.

    const (
        a = iota // 0
        b = iota // 1
        c = iota // 2
    )

또는

    const (
        a = iota // 0
        b           // 1
        c           // 2
    )

## 예제 코드
다음은 변수와 상수를 사용한 예제 코드입니다:

    package main
    import (
        "fmt"
    )
    func main() {
        // 변수 선언
        var x int
        x = 5
        fmt.Println("x:", x)
        // 변수 선언과 초기화
        var y int = 10
        fmt.Println("y:", y)
        // 타입 생략
        var z = 20
        fmt.Println("z:", z)
        // 짧은 선언
        a := 30
        fmt.Println("a:", a)
        // 여러 변수 선언
        var b, c, d int
        b, c, d = 1, 2, 3
        fmt.Println("b, c, d:", b, c, d)
        // 상수 선언
        const pi = 3.14
        fmt.Println("pi:", pi)
        // 여러 상수 선언
        const (
            width  = 100
            height = 200
        )
        fmt.Println("width, height:", width, height)
        // iota 사용
        const (
            e = iota // 0
            f        // 1
            g        // 2
        )
        fmt.Println("e, f, g:", e, f, g)
    }
이 코드를 실행하면 다음과 같은 결과를 출력합니다:

    go run variableConstant.go
    x: 5
    y: 10
    z: 20
    a: 30
    b, c, d: 1 2 3
    pi: 3.14
    width, height: 100 200
    e, f, g: 0 1 2

이 예제는 Go 언어에서 변수와 상수를 선언하고 사용하는 방법을 보여줍니다. 변수는 데이터를 저장하고 변경할 수 있으며, 상수는 변경되지 않는 값을 저장하는 데 사용됩니다. iota는 상수 그룹 내에서 자동으로 증가하는 값을 생성하는 데 유용합니다.
