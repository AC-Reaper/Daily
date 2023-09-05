## 语言基础

### 数据类型：基本数据类型，引用数据类型（类似于指针）

#### 整型：byte(1)	short(2)	int(4)	long(8)

#### 浮点型：float	double

#### 布尔型：bollean

#### 字符型：char(2)	 



#### **标识符**：

##### 必须以字母、下划线（_）或美圆符号（$）开头的，紧跟着可以由字母、数字、下划线和美圆符号组合而成的字符序列。

####   在 Java 编程中，经常遵循以下的编码规则：

#####   （1）类名首字母应该大写，变量、方法及对象首字母小写

#####   （2）对于所有标识符，其中包含的所有单词都应紧靠在一起，而且大写中间单词的首字母。

#####   （3）Java 包（Package）属于特殊的一种，它们全都是小写，即便中间的单词也是如此

####   有关关键字值得我们注意的地方：

#####  （1）所有 Java 关键字都是小写, true、false 和 null 为小写，而不是象在 C++ 语言中那样为大写。

#####  （2）无 sizeof 运算符，因为 java 所有数据类型的长度和表示是固定的，与平台无关，不是象在 C++ 语言中那样数据类型的长度根据不同的平台而变化。这正是 Java 语言的一大特点。



#### 注意

##### 在 Java 语言中 boolean 类型只允许使用 boolean 值，在整数类型和 boolean 类型之间无转换计算。 在 C++ 语言中允许将数字值转换成逻辑值，在 Java 编程语言中是不允许的。

##### Java字符串：String

#### 数据类型转换:

##### 自动类型转换（隐形）转换规则：按精度从低到高转换。

##### 说明不同类型的数据进行运算时，按精度从低到高转换,需先转换为同一类型，然后进行运算。

##### byte→short→char→int→long→float→double  

##### 强制类型转换 

##### 如果要将高精度数据转换成低精度数据时，就要进行强制类型转换。强制类型转换的格式如下：

```java
//强制类型转换  
double x=12.34;
  int a;

	a=(int) x;

//字符串转数据 Byte, Short, Integer, Long, Float, Double, Boolean
//Integer.ParseInt(String s)；
double x = Double.parseDouble(“12.345”)
//数据转字符串
String s = "" + x;

```



#### 数据输入与输出

##### 数据输入

```java
char c = ' ';
        try{
            c = (char)System.in.read();
        }
        catch(IOException e)
        {}

```

```java
//先以字符串方式输入数据，然后转换为所需要数据
          String  str;
          BufferedReader  buf=new BufferedReader(new
                                  InputStreamReader(System.in));
          str = buf.readLine();

```

```java
          String  str;
          int a, b;
          BufferedReader  buf=new BufferedReader(new
                                  InputStreamReader(System.in));
          System.out.println(“输入第一个整数：”):
          str = buf.readLine();
          a = Integer.ParseInt(str);
          System.out.println(“输入第二个整数：”):
          str = buf.readLine();
          b = Integer.ParseInt(str);
          System.out.println(a+”+”+b+” = “+(a+b)):
//两数之和
```



```java
//优化
//Java SE 5 版本开始，在 java.util 类库提供一个简化版本的输入对象 Scanner 
Scanner reader = new Scanner(System.in);
          int a = reader.nextInt();
          double x = reader.nextDouble();
          String str = reader.next(); 

 /*         其中 nextXXX() 可以有 
          nextByte(), nextShort(), nextInt(), nextLong(), 
          nextFloat(), nextDouble(), next(), nextLine() 
          next()  是以空格，Tab 键或回车键作为结束符
          nextLine() 是以回车键作为结束符，读取一行字符串
*/
```



##### 数据输出

#####  

##### 格式化的数据输出函数 printf() 

##### 它的作用是向终端输出若干个任意类型格式化数据。它的一般形式为：**System.out.printf(格式控制，表达式列表)

##### 其中“格式控制”是由双引号括起来的字符串，也称“转换控制字符串”，它里面也包含两部分信息

##### 一是 格式说明，是由“%”和格式字符组成的，如%d等

##### 二是 普通字符，即在执行时原样输出的字符

##### 简单输出函数 print() 和 println() （加换行）

##### System.out.print(输出任意类型数据);

#### 运算符

Java中的运算符可以分为多个类别，下面是对它们的概述：

1. ##### **算术运算符**:
   
   - `+`：加法
   - `-`：减法
   - `*`：乘法
   - `/`：除法
   - `%`：模（取余）
   - `++`：递增
   - `--`：递减
   
2. ##### **关系运算符**:
   
   - `==`：等于
   - `!=`：不等于
   - `>`：大于
   - `<`：小于
   - `>=`：大于或等于
   - `<=`：小于或等于
   
3. ##### **逻辑运算符**:
   
   - `&&`：逻辑与
   - `||`：逻辑或
   - `!`：逻辑非
   
4. ##### **位运算符**:
   
   - `&`：按位与
   - `|`：按位或
   - `^`：按位异或
   - `~`：按位非
   - `<<`：左移
   - `>>`：右移
   - `>>>`：无符号右移
   
5. ##### **赋值运算符**:
   
   - `=`：赋值
   - `+=`：加等于
   - `-=`：减等于
   - `*=`：乘等于
   - `/=`：除等于
   - `%=`：模等于
   - `<<=`：左移等于
   - `>>=`：右移等于
   - `&=`：按位与等于
   - `^=`：按位异或等于
   - `|=`：按位或等于
   
6. ##### **条件运算符** (也被称为三元运算符):
   
   - `? :`：例如，`x > y ? x : y`，如果`x`大于`y`，返回`x`，否则返回`y`。
   
7. ##### **类型比较运算符**:
   
   - `instanceof`：检查对象是否是特定类型的实例。
   
8. ##### **特殊运算符**:
   
   - `()`：强制类型转换
   - `.`：访问对象的成员
   - `[]`：数组下标
   - `new`：对象和数组的实例化
   - `::`：方法引用运算符（Java 8+）

这是Java中运算符的基本概述。每种运算符都有其具体的用法和优先级。在编写复杂的表达式时，建议使用括号来确保按照预期的顺序执行操作。

#### Java源程序代码的结构

Java源程序代码的结构主要是围绕它的面向对象的特性设计的。以下是一个Java源代码文件的基本结构：

1. **包声明** (Package Declaration):
   - 在源文件的最顶部, 通常是这样的：
     ```java
     package com.example.myapp;
     ```

2. **导入模块** (Import Statements):
   - 导入其他包中的类和接口，这样我们可以在当前类中使用它们。
     ```java
     import java.util.List;
     import java.util.ArrayList;
     ```

3. **注解** (Annotations) 和 **注释** (Comments):
   - 注解例如：`@Override`, `@Deprecated`
   - 注释可以是单行 `//` 或多行 `/* */`

4. **类声明/枚举/接口声明** (Class/Enum/Interface Declaration):
   - Java源文件可以包含多个类、接口和枚举，但只能有一个public类/接口/枚举，并且它的名字必须与文件名相匹配。
     ```java
     public class MyClass {
         // class body
     }
     ```

5. **类的成员** (Class Members):
   - **变量** (Variables): 可以是实例变量（非静态）或类变量（静态）
   - **方法** (Methods): 函数或行为定义
   - **构造函数** (Constructors): 用于创建和初始化类的新对象
   - **内部类、接口、枚举** (Inner classes/interfaces/enums)

6. **主方法** (Main Method):
   - 作为应用程序的入口点，具有以下签名：
     ```java
     public static void main(String[] args) {
         // code
     }
     ```

以下是一个简单的Java源代码结构的例子：

```java
// 1. Package Declaration
package com.example.myapp;

// 2. Import Statements
import java.util.List;
import java.util.ArrayList;

// 3. Annotations/Comments
/**
 * A sample Java class.
 */
// 4. Class Declaration
public class MyClass {

    // 5. Class Members

    // Variables
    private int myVar;

    // Constructors
    public MyClass(int myVar) {
        this.myVar = myVar;
    }

    // Methods
    public int getMyVar() {
        return myVar;
    }

    // 6. Main Method
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```

这就是Java源代码的基本结构。确保理解并遵循这个结构，因为这是Java编译器和JVM期望的结构。