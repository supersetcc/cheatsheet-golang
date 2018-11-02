# Golang workspace

## TODO
- (/) vim-go install
- (/) snippet setting => already installed in vim-go
- (v) exercises
- (x) screenshot -> to link (to use in markdown)
- (x) https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
- (x) control flow

## Vim + Go
- https://github.com/fatih/vim-go

### Add on .vimrc
```
nmap <silent> <leader>g :GoRun<CR>
```

### Add on go.snippets (Ultisnips)
```
snippet pa
package main

import "fmt"

func main() {
	${0}
}
endsnippet


snippet fmt
fmt.Println(${0})
endsnippet
```


## Reference
- [Udemy Course Outline](https://docs.google.com/document/d/1jGdUyurQhPxtr_nd7z-0GKOjhrfdtkSjFKORa44favQ/edit#heading=h.vgdf3swpywdi)
- [Udemy Course GolangTraining](https://github.com/GoesToEleven/GolangTraining)
