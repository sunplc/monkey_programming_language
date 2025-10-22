# 代码来自《用go语言自制解释器》

## 文本替换宏系统 语法宏系统
> 宏可以分为两大类：文本替换宏系统和语法宏系统。在我看来，它们分别相当于搜索替换和代码即数据两个类别。


## 在最后一章宏系统中，最终的代码有一个问题，即宏函数多次调用时，参数仍然是第一次调用时的参数。

### 例如：
```
// 定义宏
let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); };

// 第一次调用
unless(1>0, puts("not great"), puts("great"));
// 输出结果：great

// 第二次调用
unless(1>2, puts("not great"), puts("great"));
// 输出结果仍然是：great，有问题，实参还是第一次的参数值

```

```
let m = macro(a) { quote(unquote(a) + unquote(a)); };

m(1)
// 输出结果是2

m(2)
// 输出结果还是2，错误
```

> 为何会出现此问题？原因是第一次宏调用时，宏替换函数 ExpandMacros() 调用了 quote() 函数，quote() 函数调用了 evalUnquoteCalls() 函数，而 evalUnquoteCalls() 函数会将 AST 中所有 unquote 函数 替换为 ast.Node 节点。
后续再次宏调用时，由于AST中的 unquote函数已被替换，所以 AST 不会再发生变化。