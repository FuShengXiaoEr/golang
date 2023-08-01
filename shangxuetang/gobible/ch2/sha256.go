package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
)

func main1(){
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(bitDiff(c1[:],c2[:]))
}

func bitDiff(hash1, hash2 []byte) int{
	// 把哈希值转换为二进制
	bin1 := fmt.Sprintf("%b",new(big.Int).SetBytes(hash1))
	bin2 := fmt.Sprintf("%b",new(big.Int).SetBytes(hash2))

	// 右对齐，左边用0填充
	bin1 = fmt.Sprintf("%0256s",bin1)
	bin2 = fmt.Sprint("%0256s",bin2)

	fmt.Println("-------------------")
	fmt.Println(bin1)
	fmt.Print(bin2)

	// 初始化不同位数为0
	difBits := 0

	//逐位比对
	for i := 0; i < 256; i++{
		if bin1[i] != bin2[i] {
			difBits++
		}
	}

	return difBits
}

// }

// func main2() {
// 	hashType := flag.String("type","256","hash type (256,384,512)")
// 	flag.Parse()
// 	var hashFunc func(data []byte) []byte

// 	switch *hashType {
//     case "256":
//         hashFunc = func (data []byte) []byte {  
// 			sum := &sha512.Sum384(data)
//   			return (*sum)[:]
// 		}
//     case "384":
//         hashFunc = sha512.Sum384()
//     case "512":
//         hashFunc = sha512.Sum512()
//     default:
//         fmt.Fprintf(os.Stderr, "Invalid hash type: %s\n", *hashType)
//         os.Exit(1)
//     }
// }


func main() {
	hashType := flag.String("type", "256", "hash type (256, 384, 512)")
    flag.Parse()

    var hashFunc func(data []byte) []byte

    switch *hashType {
    case "256":
        hashFunc = func(data []byte) []byte {
            var sum [32]byte
            sum = sha256.Sum256(data)
            return (&sum)[:]
        }

    case "384":
        hashFunc = func(data []byte) []byte {
            var sum [48]byte
            sum = sha512.Sum384(data)
            return (&sum)[:]
        }

    case "512":
        hashFunc = func(data []byte) []byte {
            var sum [64]byte
            sum = sha512.Sum512(data)
            return (&sum)[:]
        }

    default:
        fmt.Fprintf(os.Stderr, "Invalid hash type: %s\n", *hashType)
        os.Exit(1)
    }

	
    var buffer bytes.Buffer

    hashWriter := func(writer io.Writer) io.Reader {
        io.Copy(writer, &buffer)
        return &buffer
    }

    io.Copy(os.Stdout, hashWriter(os.Stdin))

    hash := hashFunc(buffer.Bytes())


	 // 格式化hash为字符串输出
	 fmt.Printf("%x", hash)
}