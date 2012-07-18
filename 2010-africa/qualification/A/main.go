package main

import (
	"./cdjm"
	"./rfile"
	"fmt"
	"strconv"
	"strings"
)

func spendMoney(credit int, items []int) (outInd1 int, outInd2 int, notFound bool) {

	//fmt.Println(items)

	//items'de slice'ları kullandığımız için buna gerek kalmadı range kullanabiliriz
	// for i:=0; i<itemNo; i++{}

	for i, v := range items {

		itmsRgt := items[(i + 1):]
		for i2, v2 := range itmsRgt {
			if v+v2 == credit {

				outInd1 = i + 1          //output formatındaki index 1 ile başlıyor ve slice düzenlemeleri yapmamız gerek. 
				outInd2 = i + 1 + i2 + 1 //bu slice'ın da 1 ile başladığını kabul etmeliyiz, kesimine göre  
				return

				// index2 her zaman büyük olacağı için

				/*
					switch {
					case oindex1 > :
						no1 = v2
						no2 = v
						return
					default:
						no1 = v
						no2 = v2
						return
					}
				*/
			}
		}

	}

	//no1 = credit
	//no2 = itemNo
	notFound = true
	return

}

func main() {

	fmt.Println("started!")
	rfile.PrintE("testing imports...")
	//lnArr := make([]int, 100)
	lnArr, err := rfile.Readlns("A-large-practice.in")

	if err != nil {
		fmt.Println(err)
	}

	count, err := strconv.Atoi(lnArr[0])
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(n+1) 

	/*
		var countStr int
		countStr = lnArr[0]
	*/

	//fmt.Println(lnArr)

	/*
		count , err := strconv.ParseInt(countStr,10,0)

		if err!=nil{
		fmt.Println(err)
		}
		fmt.Println(count, "***")
	*/

	//burada 1'den baslatabilirdik ama array slice'ının özelliğini kullanmak istedik diyelim:
	lnArr2 := lnArr[1:]

	var credit int
	var itemNo int
	itemsS := make([]string, 2000)
	items := make([]int, 2000)

	resultLines := make([]string, count)

	//count = 1 //TEST

	for i := 0; i < count; {
		/*fmt.Println(lnArr2[i*3])
		fmt.Println(lnArr2[i*3+1])
		fmt.Println(lnArr2[i*3+2])
		*/
		credit, _ = strconv.Atoi(lnArr2[i*3]) //kredi
		itemNo, _ = strconv.Atoi(lnArr2[i*3+1])

		itemsS = strings.Split(lnArr2[i*3+2], " ")
		for i, v := range itemsS {
			items[i], _ = strconv.Atoi(v)
		}

		item1, item2, notFound := spendMoney(credit, items[:itemNo])
		fmt.Println(item1, item2, notFound)
		resultLines[i] = strconv.Itoa(item1) + " " + strconv.Itoa(item2)
		i = i + 1

	}

	cdjm.Output("./output.txt", resultLines)

}
