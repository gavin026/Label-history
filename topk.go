package main

import (
	"fmt"
	//"flag"
	//"os"
	//"io"
	//"bufio"
	//"strconv"
	//"strings"
)

//var flagFile  = flag.String("f","","data files")

type MinHeap struct {
	maxsize int
	arr		[]int
}

func (this *MinHeap) createMinHeap(a []int) {
	for i:= 0; i< this.maxsize; i++ {
		this.arr = append(this.arr,a[i])
	}
	k := len(this.arr)/2 - 1
	for i := k ;i >= 0; i-- {
		this.filterDown(i)
	}
}

func (this *MinHeap) insert(val int) {
	if val > this.getTop() {
		this.arr[0] = val;
		this.filterDown(0)
	}
}

func (this *MinHeap) filterDown(current int) {
	end := len(this.arr) - 1
	child := current*2 + 1
	val := this.arr[current]
	for child <= end {
		if child < end && this.arr[child+1] < this.arr[child] {
			child++
		}
		if val < this.arr[child] {
			break
		} else {
			this.arr[current] = this.arr[child]
			current = child
			child = child*2 +1
		}
	}
	this.arr[current] = val
} 

func (this *MinHeap) getTop() int {
	if len(this.arr) != 0 {
		return this.arr[0]
	}
	panic("nil!")
}

func (this *MinHeap) getHeap() []int {
	var heap []int
	size := len(this.arr)
	for i := 0; i < size; i++ {
		heap = append(heap,this.arr[i])
	}
	return heap
}

func main() {
	// array := make([]int, 11, 100)
	// failf := func(msg string, args ...interface{}) {
		// fmt.Fprintf(os.Stderr, msg+"\n", args...)
		// os.Exit(1)
	// }
	// flag.Parse()
	// if *flagFile == ""  {
		// failf("no param!")
	// }
	// inputFile, inputError := os.Open(*flagFile)
    // if inputError != nil {
        // failf("An error occurred on opening the inputfile\n" +
            // "Does the file exist?\n" +
            // "Have you got acces to it?\n")
        // return // exit the function on error
    // }
    // defer inputFile.Close()

    // inputReader := bufio.NewReader(inputFile)
	
    // for {
        // inputString, readerError := inputReader.ReadString('\n')
		// if readerError == io.EOF {
            // break
        // }
		// newdata := strings.TrimSuffix(inputString,"\n")
		// data, _ := strconv.Atoi(newdata)
		// array = append(array,data)
    // }
	var array = []int{7,6,1,2,9,8,5,12,23,44,21,33,0,4,3}
	MHP := &MinHeap {
		maxsize: 4,
	}
	MHP.createMinHeap(array)
	
	for i:= MHP.maxsize;i<15;i++ {
		MHP.insert(array[i])
	}
	
	fmt.Println("最大的四个元素是：")
	v := MHP.getHeap()
	for _, val := range v {
		fmt.Println(val)
	}
}

