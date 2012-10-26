package main

import (
	"./cdjm"
	"strconv"
	"fmt"
	"strings"
)

type basketMap map[int]int

func addToBasket(bM basketMap , code int){
  //is coupon code present in basket? 
  _,present := bM[code]
  if present{
    bM[code]=bM[code]+1
  //or do we need to create it?
  }else{
    bM[code]=1
  }
}

func findLonely(bM basketMap) (code int) {
  for code, number:=range bM{
    if number==1{
      return code
    }
  }
  return 0 //code cant be 0 according to codejam limits so this means, nothing found.
}


func main() {







	inArr, _ := cdjm.Input("input/A-large-practice.in")
  size, _ := strconv.Atoi(inArr[0])
	
	//a == a[0:] // its a reminder how it includes number on the left
	inArr = inArr[1:]


  outArr := make([]string, 0, len(inArr) )
  //2 lines for 1 case
  caseLen :=2
  size=size*caseLen
	for i := 0; i < size; i=i+caseLen {
		scodes := strings.Split( inArr[i+1], " ") //( space separated codes )
    
    
    bM := make(basketMap) //new basket
    
    //var codes = make([]int, 0, len(scodes) )
    for _,c:= range scodes{
      code,_ := strconv.Atoi(c)
      addToBasket(bM, code ) //code added to basket!
    }
    
		fmt.Println(bM)
		fmt.Println(findLonely(bM))
		fmt.Println("********************")
		
	  outArr = append(outArr,  strconv.Itoa(findLonely(bM))  )
		
	}

  fmt.Println(outArr)

  cdjm.Output("output/now.out", outArr)


}
