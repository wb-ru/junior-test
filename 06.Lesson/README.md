
# Урок №6: Материалы по профилированию системы
## Ссылки для изучения
	
* https://www.ardanlabs.com/blog/2017/06/language-mechanics-on-memory-profiling.html
  
* https://golang.org/pkg/runtime/pprof/

* https://golang.org/doc/diagnostics.html

* https://habr.com/ru/company/badoo/blog/301990/
  
## Очень полезная статья, к изучению обязательна

* https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/



## Задача

* Имея связанный список, определите, есть ли в нем цикл.

* Чтобы представить цикл в данном связанном списке, мы используем целое число pos, которое представляет позицию (с 0 индексами) в связанном списке, с которой соединяется tail. Если pos равно -1, то в связанном списке нет цикла.

#### Пример 1
```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

#### Пример 2

```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)
#### Пример 3

```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)
#### Шаблон 
``` Go
/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    
}
```
## Задача 

* Изучаем паттерн proxy(прокси) 
* Придумать и написать свою реализацию паттерна proxy с тестом 