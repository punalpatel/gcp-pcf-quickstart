package pivnet

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"omg-cli/config"
	"crypto/sha256"
	"io"
	"encoding/hex"
	"github.com/pivotal-cf/go-pivnet/download"
)

type TileCache struct {
	Dir string
	DisableChecksum bool
}

func NewTileCache(dir string, disableChecksum bool) (*TileCache, error) {
	if err:= os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	return &TileCache{
		Dir: dir,
		DisableChecksum: disableChecksum,
	}, nil
}

func FileName(tile config.PivnetMetadata) string {
	return fmt.Sprintf("tile-%s-%d-%d.pivotal", tile.Name, tile.ReleaseId, tile.FileId)
}

func StemcellFileName(tile config.StemcellMetadata) string {
	return fmt.Sprintf("stemcell-%s-%d-%d.pivotal", tile.Name, tile.ReleaseId, tile.FileId)
}

func (tc *TileCache) Open(tile config.PivnetMetadata) (*os.File, error) {
	if tc == nil || tc.Dir == "" {
		return nil, nil
	}

	needle := FileName(tile)
	fullPath := filepath.Join(tc.Dir, needle)

	files, err := ioutil.ReadDir(tc.Dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.Name() == needle {

			fmt.Println("tile", needle, "found in cache")

			if !tc.DisableChecksum {
				fmt.Println("checking sha256 hash")
				matched, err := checkSum(fullPath, tile.Sha256)
				if err != nil {
					return nil, fmt.Errorf("checking sha256: %v", err)
				}

				if !matched {
					fmt.Println("sha256 mismatch. Assume tile needs to be re-downloded")
					return nil, nil
				}
			}

			return os.Open(fullPath)
		}
	}

	return nil, nil
}

func checkSum(path string, hash string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	bar := download.NewBar()
	s, _ := f.Stat()

	h := sha256.New()

	var bufsize int64 = 1024 * 8
	buf := make([]byte, bufsize)

	bar.SetOutput(os.Stdout)
	bar.SetTotal(s.Size())
	bar.Kickoff()

	defer bar.Finish()

	for {
		nr, er := f.Read(buf)
		if nr > 0 {
			nw, ew := h.Write(buf[0:nr])
			if nw > 0 {
				bar.Add64(int64(nw))
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	if err != nil {
		return false, err
	}

	ha := hex.EncodeToString(h.Sum(nil))

	if ha != hash {
		return false, nil
	}

	return true, nil
}
