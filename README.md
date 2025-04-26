# isemoji

[![go reportcard](https://goreportcard.com/badge/github.com/emoj-info/go-isemoji)](https://goreportcard.com/report/github.com/emoj-info/go-isemoji)
[![GoDoc](https://godoc.org/github.com/emoj-info/go-isemoji?status.svg)](https://godoc.org/github.com/emoj-info/go-isemoji)
[![license](https://img.shields.io/github/license/emoj-info/go-isemoji)](./LICENSE)

This is [fork](https://github.com/makeworld-the-better-one/go-isemoji).

Go library to test if a string is an emoji.

## Usage

```go
isemoji.IsEmoji("🫩")    // True
isemoji.IsEmoji("test")  // False
isemoji.IsEmoji("🙇🏼‍♂️🤗")  // False, because there are multiple emojis
isemoji.Name("🫆")       // "fingerprint"
```

Find the full documentation on [godoc](https://godoc.org/github.com/dejurin/go-isemoji).

## License

This project is under the MIT License.
