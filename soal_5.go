package main

// mengimpor package bufio, fmt, dan os
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// membuat scanner dan memanggil package bufio dan os 
	scanner := bufio.NewScanner(os.Stdin)
	// mencetak string sebelum meminta input
	fmt.Println("Masukkan nama Anda: ")
	// meminta input pengguna
	scanner.Scan()
	// mencetak input pengguna
	fmt.Println("Nama saya adalah:",scanner.Text())
}