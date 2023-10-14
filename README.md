# <img src="https://raw.githubusercontent.com/JuniorBecari10/CSimple-Old/main/logo.png">

The compiled version of [Simple](https://github.com/JuniorBecari10/Simple).

Simple is a simple, interpreted programming language. <br />
It's very easy to use.

# About

CSimple is a compiler for the Simple programming language written in Go.

It has its own bytecode instructions, which you can write and assemble them directly using the built-in assembler.

todo! make performance tests

# Usage
```
csimple compile / build <file> | run <file> | assemble <file> | -v / --version | -h / --help
```

Command|Description
---|---
compile / build \<file>|compile source code in 'file' to bytecode
run \<file>|run bytecode in 'file'
assemble \<file>|assemble code in 'file' to bytecode
-v / --version|show the version number
-h / --help|show the help message

# Syntax

Simple has a simple, assembly-like syntax.

## Hello World and Printing

You can print to the screen using the `print` and `println` statements.

```
println "Hello, World!"
```

`println` inserts a newline character (`\n`) at the end, `print` does not.

## Variables

Declaring variables is simple also! Use this syntax:
```python
name = "value"
```

There are 3 different variable types in Simple: `str`, `num` and `bool`.

### User Input

You can use the `input` keyword.
```
name = input
```

You can also specify the data type you expect the user to type. If it doesn't match, the language will keep prompting until they type the correct one.

```
name = input str
age = input num

println "Your name is " + name + " and your age is " + age
```

## Labels

Labels are the core of Simple. With them, you don't need loops.


