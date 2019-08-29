package thumbnail

import (
	"gopl.io/ch8/thumbnail"
	"log"
)

// ImageFile reads an image form infile and
// writes a thumbnail-size image version in the same dir
func ImageFile(infile string) (string, error) {
	return "", nil
}

func makeThumbnails(filenames []string)  {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//incorrect
//doesnt wait for go func to finish
func makeThumbnails2(filenames []string)  {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string)  {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)

		//wait to finish
		for range filenames {
			<-ch
		}
	}
}
