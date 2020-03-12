
# Урок №5: Понимаем поинтеры и типы данных
## Ссылки для изучения
	
* https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85
* https://rtfm.co.ua/golang-ukazateli-podrobnyj-razbor/
* https://habr.com/ru/post/339192/
* https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html

## Поинтеры через тесты

* https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md
* https://github.com/quii/learn-go-with-tests/tree/master/pointers
  
## Изучаем исходники

* https://github.com/golang/go/blob/master/src/runtime/proc.go
 
## Включая подссылки в статье

* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
  
## Типы в сорсах
* https://github.com/golang/go/blob/0ae9389609f23dc905c58fc2ad7bcc16b770f337/src/runtime/type.


## Задача
Дана строка, содержащую только символы '(', ')', '{', '}', '[' и ']', определите, является ли входная строка допустимой.

Входная строка действительна, если:
* Открытые скобки должны быть закрыты скобками того же типа.
* Открытые скобки должны быть закрыты в правильном порядке.
* Обратите внимание, что пустая строка также считается допустимой.

#### Пример 1
```
Input: "()"
Output: true
```
#### Пример 2 

```
Input: "()[]{}"
Output: true
```
#### Пример 3
```
Input: "(]"
Output: false
```
#### Пример 4
```
Input: "([)]"
Output: false
```
#### Пример 5 
```
Input: "{[]}"
Output: true
```
### Шаблон 
```Go
func isValid(s string) bool {
    
}
```

## Задача 

Изучаем паттерн singleton(одиночка) 

Придумать и написать свою реализацию паттерна siingleton
https://github.com/AlexanderGrom/go-patterns/tree/master/Creational/Singleton





