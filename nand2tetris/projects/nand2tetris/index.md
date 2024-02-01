+++
title = 'Nand2Tetris'
date = 2024-01-29T13:55:12+08:00
description = "关于Coursera课程Nand2Tetris的projects的一些解决办法和想法"
draft = true
categories = [
    "Course-CS"
]
tags = [
    "Nand2Tetris"
]
+++


# projects

## 00

略过

## 01

在上第一周课的时候，有一门课中间出现了一个习题：**Nand(a, a)**还可以用什么方式表示，那道题的答案是：**Not(a)**, 所以proj1最先实现的应该是Not逻辑门，有了Not逻辑门和Nand逻辑门，后续的就按照逻辑走就很好实现出来了。

## 02

### 知识点

#### 关于原码、反码和补码

* 原码：最高位为符号位，0表示正数，1表示负数，其余位表示数值。例如，+7的原码是0111，-7的原码是1111（假设我们只用4位表示）。

* 反码：正数的反码与其原码相同，负数的反码是其原码符号位不变，其余位取反。例如，+7的反码是0111，-7的反码是1000。

* 补码：正数的补码与其原码相同，负数的补码是其反码加1。例如，+7的补码是0111，-7的补码是1001。

补码的引入主要是为了解决原码和反码在进行加减运算时需要考虑符号的问题，而补码可以直接进行加减运算，大大简化了计算机内部的运算过程。

**反码的英文名称是 "One's Complement"，补码的英文名称是 "Two's Complement"。 反码有正负0的区别，补码只有一个0。**

#### 关于proj2具体要求和相关知识点

[官方原文链接](https://www.nand2tetris.org/_files/ugd/44046b_f0eaab042ba042dcb58f3e08b46bb4d7.pdf)

其中包括了各种chip的结构和逻辑，还有相关真值表等等。

### 实现

根据上述连接，按照它给的顺序来实现即可，首先第一个是HalfAdder，仔细观察一下给出的真值表就可以发现规律了，**carry位可以用And(a, b)来实现，sum位可以用Xor(a, b)来实现。**剩下的就照葫芦画瓢就好了。

Inc16的实现需要一个之前没有遇到过的语法，这个貌似在第一周讲hdl的时候提到了一嘴。

```
CHIP Inc16 {
    IN in[16];
    OUT out[16];

    PARTS:
    Add16(a = in, b[0] = true, b[1..15] = false, out = out);
}
```

最后ALU的实现花了一些时间，主要还是用到了一些前面并没有出现的hdl语法问题，[这里是hdl语法连接](https://www.nand2tetris.org/_files/ugd/44046b_f0eaab042ba042dcb58f3e08b46bb4d7.pdf)，有需要可以点进去看看。

最主要的两个语法：

```
// 这样写错误！
CHIP Foo {
    IN in[16];
    OUT out;
    PARTS:
    Something16 (in=in, out=notIn);
    Or8Way (in=notIn[4..11], out=out);
}

// 这样写正确
    Something16 (in=in, out[4..11]=notIn);
    Or8Way (in=notIn, out=out);
```

```
// 当直接指定为最终输出后，后面若再想使用out就会有问题，这里可以指定多个最终输出，以供后文使用。
CHIP Foo {
    IN a, b, c;
    OUT out1, out2;
    PARTS:
    Something (a=a, b=b, out=x, out=out1);
    Whatever (a=x, b=c, out=out2);
}

```

## 03

