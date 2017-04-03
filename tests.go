package main

import (
	// "github.com/cjesp/imagefilters"
	// "os"
	// "fmt"
	// _ "image/jpeg"
	// _ "image/gif"
	// "image"
	// "image/png"
	// "strconv"
	// "time"
	// "runtime"
	"imagetests/commands"
	"fmt"
	"os"
)

func main()  {
	// if len(os.Args) != 4 {
	// 	fmt.Println("Usage: gauss.exe <filename> <stdDev> <maskSize>")
	// 	return
	// }

	// stdDev, err := strconv.ParseFloat(os.Args[2], 64) 

	// if err != nil {
	// 	fmt.Println("Error: the standard deviation needs to be a float")
	// }

	// mskSize, err := strconv.Atoi(os.Args[3])
	
	// if err != nil || mskSize % 2 == 0 {
	// 	fmt.Println("Error: mask size needs to be an integer and an ueven number")
	// }

	// infile, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("Error: %v doesn't seem to exist.", os.Args[1])
	// 	return
	// }
	// defer infile.Close()

	// src, _, err := image.Decode(infile)
	// if err != nil {
	// 	fmt.Println("Error: couldn't decode the file")
	// 	return
	// }

	// t0 := time.Now()
	// runtime.GOMAXPROCS(1)
	// out := imagefilters.GaussianBlur1D(stdDev, src)
	// t1 := time.Since(t0)
	// fmt.Printf("time (1 core): %v\n", t1)

	// t0 = time.Now()
	// runtime.GOMAXPROCS(8)
	// out = imagefilters.GaussianBlur1D(stdDev, src)
	// t1 = time.Since(t0)
	// fmt.Printf("time (8 core): %v\n", t1)
	
	// outFile, err  := os.Create("blurred.png")
	// if err != nil {
	// 	fmt.Println("Error: couldn't write the image")
	// 	return
	// }
	// defer outFile.Close()
	
	
	
	// if err := png.Encode(outFile, out); err != nil {
	// 	fmt.Println("Error: couldn't encode the image")
	// }

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

