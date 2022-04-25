# filez
A tiny package showing you File info

# Install
```
go get -v github.com/Cne3Rd/filez
```

# Usage
```
import (
    "github.com/Cne3Rd/filez"
)

file := "C:\\users\\views\\bar.txt"
f, err := filez.FileMode(file)
if err != nil {
	fmt.Println(err)
}

fmt.Println(f)

```
