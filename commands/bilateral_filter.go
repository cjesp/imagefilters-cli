package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"os"
	_ "image/jpeg"
	_ "image/gif"
	"image"
	"image/png"
	// "github.com/cjesp/imagefilters"
	"runtime"
	"time"
	"github.com/cjesp/imagefilters"
)


// var input, output string
// var sigmaS, sigmaI float64

var cmdBilateralBlur = &cobra.Command{
	Use: "bilateral [sigma (spatial)] [sigma (intensity)] [input file path] [output file path]",
	Short: "Blur image with a bilateral filter",
	Long: `Blur the input image with a bilateral filter defined 
	by the sigma input parameter`,
	Run: func(cmd *cobra.Command, args []string)  {
		
		if len(args) != 4 {
			fmt.Println("usage: bilateral [sigma (spatial)] [sigma (intensity)] [input file path] [output file path].")
			return
		}
		
		sigS, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("error: couldn't parse sigma (spatial). Sigma (spatial) should be a float, e.g. 3.0.", err.Error())
			return
		}
		sigmaS = sigS

		sigI, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("error: couldn't parse sigma (intensity). Sigma (intensity) should be a float, e.g. 3.0.", err.Error())
			return
		}
		sigmaI = sigI

		inputFile, err := os.Open(args[2])
		if err != nil {
			fmt.Println("Error: couldn't open the input file.", err.Error())
			return
		}
		defer inputFile.Close()

		src, _, err := image.Decode(inputFile)
		if err != nil {
			fmt.Println("Error: couldn't parse the input image.", err.Error())
			return
		}
		
		// src, _, err := image.Decode(inputFile)
		if err != nil {
			fmt.Println("Error: couldn't parse the input image.", err.Error())
			return
		}

		outputFile, err := os.Create(args[3])
		if err != nil {
			fmt.Println("Error: couldn't open the output file.", err.Error())
			return
		}
		defer outputFile.Close()

		if Workers != 0 {
			runtime.GOMAXPROCS(Workers)
		} else {
			runtime.GOMAXPROCS(1)
		}

		var output *image.NRGBA
		if Time {
			t0 := time.Now()
			output = imagefilters.BilateralFilter(sigmaS, sigmaI, src)
			t1 := time.Since(t0)
			fmt.Println("execution time:", t1)
		} else {
			output = imagefilters.BilateralFilter(sigmaS, sigmaI, src)
		}


		if err := png.Encode(outputFile, output); err != nil {
			fmt.Println("Error: could not encode the output file.", err.Error())
		}
		
		
	},
}

func init()  {
	// cmdGaussBlur.Flags().StringVar(&input, "input", "", "path to input file")
	// cmdGaussBlur.Flags().StringVar(&input, "output", "", "path to output file (leave empty for local folder)")
	
	RootCmd.AddCommand(cmdBilateralBlur)
}