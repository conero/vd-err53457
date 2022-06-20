package repo

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Patch struct {
	DfPaths []DiffPath
	BaseDir string
}

func (c *Patch) Gzip(targetDir string) error {
	fl, er := os.OpenFile(targetDir, os.O_CREATE|os.O_TRUNC|os.O_RDONLY, 0666)
	if er != nil {
		return errors.New(fmt.Sprintf("Gzip file creation failed，\n  %v\n", er))
	}

	zw := gzip.NewWriter(fl)
	basename := c.BaseDir
	var lines []string
	var linesDel []string
	lger := getLogger()
	for _, dp := range c.DfPaths {
		filename := strings.Replace(dp.Filename, basename, "", 1)
		switch dp.Mode {
		case DiffModeAdd, DiffModeModify:
			bys, er := ioutil.ReadFile(dp.Filename)
			if er == nil {
				zw.Name = filename
				_, er = zw.Write(bys)
				if er != nil {
					lines = append(lines, fmt.Sprintf("%v file write error, %v.", filename, er))
					lger.Warnf("%v Source file write compressed file error", dp.Filename, er)
				}
			}
		case DiffModeDel:
			linesDel = append(linesDel, filename)
		}
	}

	if err := zw.Close(); err != nil {
		return errors.New(fmt.Sprintf("gzip Close error, \n  %v\n", er))
	}

	return nil
}

func NewPatch(dps []DiffPath) *Patch {
	return &Patch{
		DfPaths: dps,
	}
}
