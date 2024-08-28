package imgdemo

// log "github.com/sirupsen/logrus"

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/gographics/imagick.v2/imagick"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"runtime"
)

// Server has router and db instances
type Server struct {
	router *mux.Router
}

func NewServer() Server {

	r := mux.NewRouter()
	srv := Server{
		router: r,
	}

	return srv

}

func (s *Server) Run() {
	port := ":8080"

	s.routes()

	slog.Info("used", "cpu", runtime.NumCPU())

	fmt.Printf("started on %v\n", port)

	r := s.router

	slog.Error(http.ListenAndServe(port, r).Error())
	os.Exit(0)
}

func (s *Server) handleImage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// read image

		originalWand := imagick.NewMagickWand()
		finalWand := imagick.NewMagickWand()

		err := originalWand.ReadImage("tx.gif")
		if err != nil {
			log.Fatal(err)
		}

		numberImages := originalWand.GetNumberImages()
		//if numberImages > 1 {
		//	// rebuild optimized frame in animation
		//	originalWand = originalWand.CoalesceImages()
		//}

		for i := 0; i < int(numberImages); i++ {
			originalWand.SetIteratorIndex(i)
			mw := imagick.NewMagickWand()
			mw = originalWand.GetImage()

			err = mw.ResizeImage(128, 105, imagick.FILTER_LANCZOS, 0)
			if err != nil {
				return
			}

			finalWand.AddImage(mw)
			mw.Destroy()

		}

		originalWand.Destroy()

		var data []byte
		data, err = finalWand.GetImagesBlob()
		if err != nil {
			return
		}
		finalWand.Destroy()

		reader := bytes.NewReader(data)
		_, err = io.Copy(w, reader)
		if err != nil {
			panic(err)
		}

	}

}
