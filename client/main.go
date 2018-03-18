package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
	"github.com/hybridgroup/mjpeg"
	"net/http"
	"log"
	"time"
)

var (
	deviceID = 0
	xmlFile = "data/frontalface.xml"
	stream *mjpeg.Stream
)

func main() {
	stream = mjpeg.NewStream()

	go capture()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/" + r.URL.Path[1:])
	})
	http.Handle("/stream", stream)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func capture() {
	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifing as "Human"
		for _, r := range rects {
			gocv.Rectangle(img, r, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)

		fmt.Println(time.Now())
		time.Sleep(100 * time.Millisecond)
		fmt.Println(time.Now())

	}
}
