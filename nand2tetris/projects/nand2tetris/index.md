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

最后ALU的实现花了一些时间，主要还是用到了一些前面并没有出现的hdl语法问题，[这里是hdl语法连接](https://www.ic.unicamp.br/~rodolfo/mc404/HDL_Survival_Guida-Nand2tetris.pdf)，有需要可以点进去看看。

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

[此章的文档](https://www.nand2tetris.org/_files/ugd/44046b_862828b3a3464a809cda6f44d9ad2ec9.pdf)

DFF已经实现了，按照文档的要求和先后顺序一步步实现就好了。

值得一提的是PC，文章给出的只是PC abstraction，并没有给出具体的实现，所以需要自己设计一下PC的实现，可以参照下面的实现。因为PC要自增，需要有存储的结构在里面，所以这里需要一个Register来存储值。

```
CHIP PC {
    IN in[16],inc, load, reset;
    OUT out[16];
    
    PARTS:
    Inc16(in = t, out = incout);
    Mux16(a = t, b = incout, sel = inc, out = out1);
    Mux16(a = out1, b = in, sel = load, out = out2);
    Mux16(a = out2, b = false, sel = reset, out = out3);
    Register(in = out3, load = true, out = t, out = out);
}

```

**b**部分的实现都是一模一样，这里给出RAM512

```
CHIP RAM512 {
    IN in[16], load, address[9];
    OUT out[16];

    PARTS:
    DMux8Way(in = load, sel = address[6..8], a = load0, b = load1, c = load2, d = load3, e = load4, f = load5, g = load6, h = load7);

    RAM64(in = in, load = load0, address = address[0..5], out = out0);
    RAM64(in = in, load = load1, address = address[0..5], out = out1);
    RAM64(in = in, load = load2, address = address[0..5], out = out2);
    RAM64(in = in, load = load3, address = address[0..5], out = out3);
    RAM64(in = in, load = load4, address = address[0..5], out = out4);
    RAM64(in = in, load = load5, address = address[0..5], out = out5);
    RAM64(in = in, load = load6, address = address[0..5], out = out6);
    RAM64(in = in, load = load7, address = address[0..5], out = out7);

    Mux8Way16(a = out0, b = out1, c = out2, d = out3, e = out4, f = out5, g = out6, h = out7, sel = address[6..8], out = out);
}

```

## 04

### 知识点

#### 关于Hack

Hack是基于冯诺依曼架构的16位计算机，由一个CPU、一个指令内存（ROM）、一个数据内存（RAM）和两个内存映射IO设备（显示器和键盘）组成。

Hack数据总线位宽为16bit, 一次可以传输16bit的值；地址总线位宽15bit可以有2的15次方（32768）个不同的地址。

Hack RAM：
- 0-16384（16k）：data memory
- 16384-24576（8k）：screen memory map
- 24576-..：keyboard

Hack有关字母：
- **D**：Data Register, 只用来存储数据
- **A**：Address Register, 用来存储数据和地址
- **M**：Memory, 通过A中存储的地址来进行索引
- **R0-R15**：0到15号RAM地址
- **SP、LCL、ARG、THIS、THAT**：0到4号RAM地址
- **SCREEN**：屏幕基地址（16384）
- **KBD**：键盘基地址（24576）


#### Hack汇编语法

##### A指令

`@value`

给A存放一个15bit的值

##### C指令

`dest = comp; jump`

#### I/O 输入输出

**Display Unit**：*256行 * 512列*， 0表示白色， 1表示黑色，从16384地址开始。

因为一个地址是16个bit, Display Unit一行是512个像素点，所以每32个地址就是一行，所以一共有256 * 32 = 8192个地址。

**Keyboard**：按下某个键时，对应的16位ASCII码值出现在RAM\[24576\]。

### 实现

[还是先给出连接](https://www.nand2tetris.org/_files/ugd/44046b_7ef1c00a714c46768f08c459a6cab45a.pdf)

这一章需要着重注意的是A指令的语法，若value是一个不为0的十进制数，代表给A寄存器存储一个值，若value是一个identifier, 则使用M可以取出value的值。
若单独使用M, 而没有使用A指令的形式，那么M取出的就是以A为地址的值，所以语法不要用混了。

第二个实现在课程中老师给出的例子就有，pointer那部分就是。
