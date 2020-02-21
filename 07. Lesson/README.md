
# Урок №7: Изучаем пакет стандартной библиотеки net/http, библиотеку fasthttp
## Ссылки для изучения
	
* net/http: https://golang.org/pkg/net/http/
  
* fasthttp: https://github.com/valyala/fasthttp

            https://habr.com/ru/post/443378/
  
## Очень полезная статья, к изучению обязательна

* https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/



## Задача
Даны два двоичных дерева и представьте, что когда вы помещаете одно из них в другое, некоторые узлы двух деревьев перекрываются, а другие нет. Вам необходимо объединить их в новое двоичное дерево. Правило слияния состоит в том, что если два узла перекрываются, то сумма значений узла увеличивается как новое значение объединенного узла. В противном случае нулевой узел NOT будет использоваться в качестве узла нового дерева.
  
#### Пример 1
```
 Input: 
	Tree 1                     Tree 2                  
          1                         2                             
         / \                       / \                            
        3   2                     1   3                        
        /                           \   \                      
        5                             4   7                  
 Output: 
 Merged tree:
	     3
	    / \
	   4   5
	  / \   \ 
	  5   4   7
```
#### Шаблон 
``` Go
 /**
  Definition for a binary tree node.
 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    
}
```

#### Обратите внимание, что само бинарное дерево у нас является структурой
```Go
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
```

#### Cтатья на вики про двоичное дерево поиска:
 * https://ru.wikipedia.org/wiki/%D0%94%D0%B2%D0%BE%D0%B8%D1%87%D0%BD%D0%BE%D0%B5_%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE_%D0%BF%D0%BE%D0%B8%D1%81%D0%BA%D0%B0

## Задача 
Изучаем паттерн ChainOfResponsibility(цепочка ответственности) придумать и написать свою реализацию паттерна ChainOfResponsibility с тестом 
https://github.com/AlexanderGrom/go-patterns/tree/master/Behavioral/ChainOfResponsibility