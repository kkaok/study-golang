# 맵(Map)
Go 언어의 맵(Map)은 키(key)와 값(value)의 쌍으로 데이터를 저장하는 자료구조입니다. 맵은 해시 테이블(hash table)을 기반으로 구현되어 있어, 키를 기준으로 빠르게 데이터를 검색, 삽입, 삭제할 수 있는 특성을 가지고 있습니다.

## 맵의 선언과 초기화
맵은 make 함수를 사용하여 초기화하거나 리터럴 형태로 선언하고 초기화할 수 있습니다:

```
// 맵의 선언 및 초기화
var m map[keyType]valueType
m = make(map[keyType]valueType)
// 또는 짧은 선언 방식
m := make(map[keyType]valueType)
// 맵 리터럴을 사용한 초기화
m := map[keyType]valueType{
    key1: value1,
    key2: value2,
    //...
}
```

여기서 keyType는 키의 데이터 타입을, valueType는 값의 데이터 타입을 나타냅니다. 맵은 내장 함수인 make를 사용하여 초기화할 수도 있고, 짧은 선언 방식으로도 초기화할 수 있습니다. 맵 리터럴을 사용하면 선언과 초기화를 동시에 할 수 있습니다.

## 맵의 사용 예시

```
package main
import "fmt"
func main() {
    // 맵 초기화
    ages := map[string]int{
        "Alice":   31,
        "Bob":     34,
        "Charlie": 28,
    }
    // 요소 추가
    ages["David"] = 23
    // 요소 조회
    fmt.Println("Alice's age is", ages["Alice"])
    // 요소 삭제
    delete(ages, "Bob")
    // 키 체크
    if age, ok := ages["Charlie"]; ok {
        fmt.Println("Charlie's age is", age)
    } else {
        fmt.Println("Charlie's age is not present")
    }
    // 맵 순회
    for name, age := range ages {
        fmt.Printf("%s's age is %d\n", name, age)
    }
}
```

## 주요 맵 연산
- 요소 추가 및 갱신: 특정 키에 대한 값을 할당하여 요소를 추가하거나 갱신할 수 있습니다.

```
ages["Eve"] = 25 // 추가
ages["Alice"] = 32 // 갱신
```

- 요소 삭제: delete 내장 함수를 사용하여 맵에서 특정 키와 그에 해당하는 값을 삭제할 수 있습니다.

    delete(ages, "Bob")

- 키 체크: 맵에서 특정 키가 존재하는지 확인하고, 그 값에 접근할 수 있습니다.

    age, ok := ages["Alice"]
    if ok {
        fmt.Println("Alice's age is", age)
    } else {
        fmt.Println("Alice's age is not present")
    }

- 맵 순회: for range 구문을 사용하여 맵의 모든 요소를 순회할 수 있습니다.

    for name, age := range ages {
        fmt.Printf("%s's age is %d\n", name, age)
    }

## 맵의 특성
- 순서 보장 안 함: 맵은 요소의 순서를 보장하지 않습니다. 맵에 추가한 순서대로 요소가 저장되지 않고, 순회할 때마다 순서가 다를 수 있습니다.
- 키의 유일성: 맵은 각 키는 유일해야 하며, 동일한 키를 가진 요소가 존재할 수 없습니다. 키는 값의 주소 역할을 합니다.
- Nil 맵: 맵은 nil로 초기화될 수 있으며, nil 맵에 요소를 추가하거나 삭제하려 하면 런타임 에러가 발생합니다.
- 참조 타입: 맵은 참조 타입이므로 함수로 전달할 때는 맵의 복사본이 아니라 맵의 참조가 전달됩니다. 이는 메모리 사용 측면에서 유리하지만, 주의해서 사용해야 합니다.

맵은 매우 유연하고 효율적인 데이터 구조로, 데이터를 키-값 쌍으로 관리하고 검색하는 데 매우 유용합니다. Go 언어에서는 맵을 자주 사용하여 데이터를 구조화하고 효율적으로 관리하는 것이 일반적입니다.


Go 언어에서 맵(Map)에서 요소를 조회할 때 사용되는 문법 중 하나는 ok 패턴입니다. ok는 불리언(boolean) 타입의 변수로, 맵에서 특정 키에 대한 값이 존재하는지 여부를 나타냅니다. 이 패턴은 맵에서 요소를 안전하게 조회하고, 요소의 존재 여부를 확인하는 데 사용됩니다.

## ok 패턴 설명
맵에서 특정 키에 대한 값을 조회할 때, 맵이 그 키를 가지고 있는지 여부를 ok 패턴을 통해 확인할 수 있습니다. ok 패턴은 다음과 같은 문법을 가집니다:

    value, ok := map[key]

여기서:

> value: 맵에서 특정 키에 대한 값입니다.

> ok: 불리언(boolean) 타입의 변수로, 키가 맵에 존재하면 true, 존재하지 않으면 false가 할당됩니다.

## 예시
다음은 ok 패턴을 사용하여 맵에서 특정 키에 대한 값을 안전하게 조회하는 예시입니다:

```
package main
import "fmt"
func main() {
    ages := map[string]int{
        "Alice":   31,
        "Bob":     34,
        "Charlie": 28,
    }
    // Alice 키의 값 조회
    if age, ok := ages["Alice"]; ok {
        fmt.Println("Alice's age is", age)
    } else {
        fmt.Println("Alice's age is not present")
    }
    // Eve 키의 값 조회 (Eve는 존재하지 않는 경우)
    if age, ok := ages["Eve"]; ok {
        fmt.Println("Eve's age is", age)
    } else {
        fmt.Println("Eve's age is not present")
    }
}
```

위 예시에서는 ages 맵에서 Alice와 Eve 키의 값을 조회하고 있습니다.

- ages["Alice"] 조회 시, 맵에 Alice 키가 존재하므로 ok는 true가 되고, age 변수에는 31이 할당됩니다. 따라서 "Alice's age is 31"이 출력됩니다.
- ages["Eve"] 조회 시, 맵에 Eve 키가 존재하지 않으므로 ok는 false가 되고, age 변수는 0 값 또는 해당 타입의 제로 값이 할당됩니다. 따라서 "Eve's age is not present"가 출력됩니다.

이처럼 ok 패턴을 사용하면 맵에서 요소를 조회할 때 맵에 키가 존재하는지 먼저 확인하고, 안전하게 값을 가져올 수 있습니다. 이는 예상치 못한 런타임 에러를 방지하고 코드의 안정성을 높이는 데 유용합니다.