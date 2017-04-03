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
	"github.com/cjesp/imagefilters"
	"runtime"
	"time"
)

// var input, output string
// var sigma float64

var cmdGaussBlur = &cobra.Command{
	Use: "gauss [sigma] [input file path] [output file path]",
	Short: "Blur image with a gaussian filter",
	Long: `Blur the input image with a gaussian mask defined 
	by the sigma input parameter`,
	Run: func(cmd *cobra.Command, args []string)  {
		// fmt.Println(strings.Join(args, " "))
		
		if len(args) != 3 {
			fmt.Println("usage: gauss [sigma] [input file path] [output file path].")
			return
		}
		
		sigArg, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("error: couldn't parse sigma. Sigma is a float, e.g. 3.0.")
			return
		}
		sigmaI = sigArg

		inputFile, err := os.Open(args[1])
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

		outputFile, err := os.Create(args[2])
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
			output = imagefilters.GaussianBlur1D(sigmaI, src)
			t1 := time.Since(t0)
			fmt.Println("execution time:", t1)
		} else {
			output = imagefilters.GaussianBlur1D(sigmaI, src)
		}


		if err := png.Encode(outputFile, output); err != nil {
			fmt.Println("Error: could not encode the output file.", err.Error())
		}
		
		
	},
}

func init()  {
	// cmdGaussBlur.Flags().StringVar(&input, "input", "", "path to input file")
	// cmdGaussBlur.Flags().StringVar(&input, "output", "", "path to output file (leave empty for local folder)")
	
	RootCmd.AddCommand(cmdGaussBlur)
}