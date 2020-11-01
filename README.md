# greyscalr

A basic Go application that applies a grayscale overlay to PNG images. Please note, due to limitations of the Go image package, this can only support PNG images. The quality isn't great but it does the job for me so I'm happy :joy:

## Prerequisites

In order to run greyscalr locally, you will need Go installed on your machine. If you don't have Go installed, please download the Go binary from [here](https://golang.org/dl/) and follow the official [installation instructions](https://golang.org/doc/install).

### Demo

In this demo, I've converted a photograph I took in 2017 of a paraglider on top of Monte Baldo:mountain: in Italy from colour to greyscale.

![](demo.gif)

### Running locally

1. `git clone` this repository.
2. Navigate to the repository.
3. Copy the image you want to greyscale into the working directory.
4. Run `go run main.go paraglider.png` (Remember to pass the image name as a parameter).
5. Open the folder and there will be a `paraglider-grayscale.png` file in the directory :slightly_smiling_face:

### Kudos

Hat tip to [Golang Programs](https://www.golangprograms.com/) for the image boundary logic.
