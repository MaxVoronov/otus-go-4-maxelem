# Otus Homework #4
## Поиск максимума

Написать функцию находящую максимальный элемент в слайсе с произвольными элементами (_[]interface{}_) с использованием
пользовательской функции-компаратора.

### Пример использования

```go
package main

import (
    "fmt"
    "github.com/maxvoronov/otus-go-4-maxelem"
)

func main() {
    data := []int{41, 72, 19, 36, 65}

    result, _ := maxelem.FindMax(data, func(i, j int) bool {
        return data[i] < data[j]
    })

    fmt.Printf("Max value: %v\n", result)
}

// Output:
// Max value: 72
```
