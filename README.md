

## 语法分析

monkey语言中，语句只有let和return两种，除了语句之外的所有内容都是表达式。

普拉特解析法。（自上而下的运算符优先级解析法）




# monkey中的表达式语句（把表达式也看成是一种语句）

- 标识符: `foobar;` `x;`

   标识符也产生值，得到的就是标识符绑定的值 

- 整数字面量: `5;` `10;`
- 带运算符的表达式
  - 前缀运算符表达式：`结构：<前缀运算符><表达式>`
    ```
    -5;
    -15;
    !isTrue()
    ```
  - 中缀运算符表达式：`结构：<表达式> <中缀运算符> <表达式>`
    ```
    5 + 5;
    5 - 5;
    5 * 5;
    5 / 5;
    5 > 5;
    5 < 5;
    5 == 5;
    5 != 5;
    ```
    

