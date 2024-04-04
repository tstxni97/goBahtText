# goBahtText
go lib convert numerical var to thai text based on https://github.com/tstxni97/py3bahttext


โกไลบราลีสามารถแปลงตัวเลขเป็นข้อความไทยได้ใช้หน่วยสตางค์ได้ถูกต้อง โดยแปลงโค้ดจาก  https://github.com/tstxni97/py3bahttext

# Installation

```bash
go get github.com/tstxni97/goBahtText
```

# Usage

```go
package main

import (
    "fmt"
    "github.com/tstxni97/goBahtText"
)

func main() {
    fmt.Println(goBahtText.BahtText(2.21))
    // สองบาทยี่สิบเอ็ดสตางค์
}
```
