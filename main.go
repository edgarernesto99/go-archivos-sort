package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	var n int
	var str string
	var ss []string

	fmt.Print("Numero de strings: ")
	fmt.Scan(&n)
	scaner.Scan() //para ignorar salto de linea
	for i:=0; i<n; i++ {
		fmt.Print("String ", i+1,": ")
		scaner.Scan()
		str = scaner.Text()
		ss = append(ss, str)
	}
	//Ordenar de manera ascendente
	sort.Strings(ss) //forma ascendente
	//Respaldar en archivo txt
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(strings.Join(ss, "\n"))

	//Ordenar de manera descendente
	sort.Sort(sort.Reverse(sort.StringSlice(ss))) //Forma descendente
	//Respaldar en archivo txt
	file2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()
	file2.WriteString(strings.Join(ss, "\n"))
}