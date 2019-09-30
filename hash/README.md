# Image Hashing

[![License][License-Image]][License-Url]
[![Godoc][Godoc-Image]][Godoc-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]

## Installation

```bash
go get github.com/evanoberholster/imagecolor/phash
```

## Usage

``` Go
func main() {
        file1, _ := os.Open("sample1.jpg")
        file2, _ := os.Open("sample2.jpg")
        defer file1.Close()
        defer file2.Close()

        img1, _ := jpeg.Decode(file1)
        img2, _ := jpeg.Decode(file2)
        hash1, _ := hash.AverageHash(img1)
        hash2, _ := hash.AverageHash(img2)
        distance, _ := hash1.Distance(hash2)
        fmt.Printf("Distance between images: %v\n", distance)

        hash1, _ = hash.DifferenceHash(img1)
        hash2, _ = hash.DifferenceHash(img2)
        distance, _ = hash1.Distance(hash2)
        fmt.Printf("Distance between images: %v\n", distance)
        width, height := 8, 8
        hash3, _ = hash.ExtAverageHash(img1, width, height)
        hash4, _ = hash.ExtAverageHash(img2, width, height)
        distance, _ = hash3.Distance(hash4)
        fmt.Printf("Distance between images: %v\n", distance)
        fmt.Printf("hash3 bit size: %v\n", hash3.Bits())
        fmt.Printf("hash4 bit size: %v\n", hash4.Bits())

        var b bytes.Buffer
        foo := bufio.NewWriter(&b)
        _ = hash4.Dump(foo)
        foo.Flush()
        bar := bufio.NewReader(&b)
        hash5, _ := hash.LoadExtImageHash(bar)
}
```

## Inspired By

> Inspired by [imagehash](https://github.com/JohannesBuchner/imagehash)
> Inspired by [goimagehash](https://github.com/corona10/goimagehash)

A image hashing library written in Go. ImageHash supports:

* [Average hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
* [Difference hashing](http://www.hackerfactor.com/blog/index.php?/archives/529-Kind-of-Like-That.html)
* [Perception hashing](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)

## Special thanks to

* [Haeun Kim](https://github.com/haeungun/)
* [Dong-hee Na](https://github.com/corona10/)

## AUTHORS

* [Dominik Honnef](https://github.com/dominikh) dominik@honnef.co
* [Dong-hee Na](https://github.com/corona10/) donghee.na92@gmail.com
* [Gustavo Brunoro](https://github.com/brunoro/) git@hitnail.net
* [Alex Higashino](https://github.com/TokyoWolFrog/) TokyoWolFrog@mayxyou.com


## LICENSE

BSD 2-Clause License

Copyright (c) 2019, Evan Oberholster & Contributors

Copyright (c) 2017, Dong-hee Na

[License-Url]: https://opensource.org/licenses/BSD-2-Clause
[License-Image]: https://img.shields.io/badge/license-2%20Clause%20BSD-blue.svg?maxAge=2592000
[Godoc-Url]: https://godoc.org/github.com/evanoberholster/imageColor/hash
[Godoc-Image]: https://godoc.org/github.com/evanoberholster/imageColor/hash?status.svg
[Godoc-Image]: https://godoc.org/github.com/evanoberholster/imageColor/hash?status.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/evanoberholster/imageColor/hash
[ReportCard-Image]: https://goreportcard.com/badge/github.com/evanoberholster/imageColor/hash
