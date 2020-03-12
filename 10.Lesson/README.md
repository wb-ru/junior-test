
# Урок №10: Изучаем, что за зверь такой escape analysis и приватные интерфейсы
## Ссылки для изучения
* escape analysis:
https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html
https://habr.com/ru/company/intel/blog/422447/

* private interfaces:
blog.j7mbo.com/bypassing-golangs-lack-of-constructors/

* Interface segregation principle:
https://en.wikipedia.org/wiki/Interface_segregation_principle

* solid go design
https://dave.cheney.net/2016/08/20/solid-go-design
  
## Задача
Дано два массива, напишите функцию для вычисления их пересечения.
* Каждый элемент в результате должен появляться столько раз, сколько раз он присутствует в обоих массивах.
* Результат может быть в любом порядке.
#### Пример 1 
```
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
```
#### Пример 2 

```
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
```

## Замечания:


Что если данный массив уже отсортирован? Как бы вы оптимизировали свой алгоритм?
Что если размер nums1 маленький по сравнению с размером nums2? Какой алгоритм лучше?
Что если элементы nums2 хранятся на диске, а память ограничена, так что вы не можете загрузить все элементы в память одновременно?

## Шаблон 

``` Go
func intersect(nums1 []int, nums2 []int) []int {
    
}
```
## Задача 
Изучаем паттерн memento(хранитель) придумать и написать свою реализацию паттерна memento с тестом
https://github.com/AlexanderGrom/go-patterns/tree/master/Behavioral/Memento