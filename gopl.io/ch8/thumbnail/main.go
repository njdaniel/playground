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

//need to return err
// this causes a goroutine leak if the err is not the last filename
// two fixes
// 1) make a buffer that is big enough to allow all to finish
// 2) create goroutine to drain the channel while the goroutine returns the err w/o delay
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

//use buffer channel to return thumbfile with errors
//returns generated file names in arbitrary order
//or an error if any step fails
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err error
	}
	
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	
	return thumbfiles, nil
}

func makeThumbnail6()  {
	
}

