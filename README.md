# Zen Interpreter

A simple, for-fun language, with a simple syntax (hence the name)!

```javascript
fn main() {
    log("Hello World!")
}
```

## Progress

1. Tokenizer ✓
2. Lexer ✓
3. REPL ☡ (incredibly simple implementation)
4. Parser ❌
5. Semantic Analyzer ❌
6. IR ❌
7. Assembler :: will use [NASM Assembler](https://www.nasm.us/)
8. Linker :: will use [LD](https://ftp.gnu.org/old-gnu/Manuals/ld-2.9.1/html_mono/ld.html)

## To run

A REPL that can only handle one line of code:

```bash
go run main.go
```

To run against a file.

```
go run main.go -f hello_world.zen
```

Once again, so far can only tokenize.
