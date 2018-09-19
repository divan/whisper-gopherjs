Proof-of-concept Whisper encryption code compiled directly from Go go-ethereum into JS using [GopherJS](https://github.com/gopherjs/gopherjs).

# Usage
Assuming you have [gopherjs](https://github.com/gopherjs/gopherjs) installed, run:

```
gopherjs build
```

The output will be `whisper-gopherjs.js` and `whisper-gopherjs.js.map` files. Include them into your project and use.

# Example

```
$ node
> const whisper = require('./whisper-gopherjs.js')
> undefined
> whisper
> { EncryptSymmetric: [Function], Key: [Function] }
> whisper.Key("password")
> '���\r\u0019\u0015�l �<\\&�en�qA�g��\t�_W\u001c��\u0018$'
```
