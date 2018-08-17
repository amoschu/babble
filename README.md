Babble
=========

# This is a fork of
[github.com/tjarratt/babble](https://github.com/tjarratt/babble).

Babble is a small utility that generates random words for you. I found this useful because occasionally you need a random word for testing.

![tower of babel](http://image.shutterstock.com/display_pic_with_logo/518173/140700250/stock-photo-tower-of-babel-first-variant-raster-variant-140700250.jpg)

Usage
-----

```go
package your_app

import (
  "github.com/amoschu/babble"
)

func main() {
  babbler := NewBabbler()
  // specify the number of desired words to the Babble() method
  println(babbler.Babble(3)) // excessive-yak-shaving (or some other phrase)

  // optionally set your own separator
  babbler.Separator = " "
  println(babbler.Babble(3)) // "hello from nowhere" (or some other phrase)

  return
})
```
