
# Урок №8: Понимаем context
## Ссылки для изучения
* https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
* https://golang.org/pkg/context/
* https://habr.com/ru/company/nixys/blog/461723/

## Понимаем через тесты:
* https://github.com/quii/learn-go-with-tests/blob/master/context.md
* https://github.com/quii/learn-go-with-tests/tree/master/context
  
## Задача
Дан корневой узел двоичного дерева поиска, вернуть сумму значений всех узлов со значениями между L и R (включительно).
Дерево двоичного поиска гарантированно будет иметь уникальные значения.
* Количество узлов в дереве не более 10000.
* Окончательный ответ будет гарантированно меньше 2^31.
  
#### Пример 1 
```
Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32
```
#### Пример 2

``` 
Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23
```

## Задача 
Изучаем паттерн mediator (медиатор) придумать и написать свою реализацию паттерна mediator с тестом
https://github.com/AlexanderGrom/go-patterns/tree/master/Behavioral/Mediator
  