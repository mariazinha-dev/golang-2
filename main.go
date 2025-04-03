package main
 import "fmt"
func main(){	
	var nomes= []string{"Maria","billie", "camila","jao","biazin"}
	var onenomes =nomes[0:2]
 	fmt.Println(onenomes)
	var twonomes=nomes[3:5]
	fmt.Println(twonomes)
	var threenomes=nomes[2]
	fmt.Println(threenomes)
	fmt.Println(nomes,len(nomes),cap(nomes))
	}