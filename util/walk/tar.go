package walk

import (
	"archive/tar"
	"compress/bzip2"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"util/logs"
)

func xtar(input, output string) bool {

	if !checkFileExt(input, ".tar") {
		return false
	}

	f, err := os.Open(input)
	if err != nil {
		logs.Warn(err)
		return false
	}
	defer f.Close()

	fr := tar.NewReader(f)
	for {

		fh, err := fr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logs.Warn(err)
			continue
		}

		fp := filepath.Join(output, fh.Name)
		if fh.Typeflag == tar.TypeDir {
			os.MkdirAll(fp, 0755)
			continue
		}

		fw, err := os.Create(fp)
		if err != nil {
			logs.Warn(err)
			continue
		}

		io.Copy(fw, fr)
		fw.Close()
	}
	return true
}

func xgz(input, output string) bool {

	if !checkFileHead(input, M_GZ) {
		return false
	}

	f, err := os.Open(input)
	if err != nil {
		logs.Warn(err)
		return false
	}
	defer f.Close()

	fr, err := gzip.NewReader(f)
	if err != nil {
		logs.Warn(err)
		return false
	}
	defer fr.Close()

	fp := filepath.Join(output, strings.TrimSuffix(filepath.Base(input), filepath.Ext(input)))
	fw, err := os.Create(fp)
	if err != nil {
		logs.Warn(err)
		return false
	}

	_, err = io.Copy(fw, fr)
	fw.Close()

	return err == nil
}

func xbz2(input, output string) bool {

	if !checkFileHead(input, M_BZ2) {
		return false
	}

	f, err := os.Open(input)
	if err != nil {
		logs.Warn(err)
		return false
	}
	defer f.Close()

	fr := bzip2.NewReader(f)

	fp := filepath.Join(output, strings.TrimSuffix(filepath.Base(input), filepath.Ext(input)))
	fw, err := os.Create(fp)
	if err != nil {
		logs.Warn(err)
		return false
	}

	_, err = io.Copy(fw, fr)
	fw.Close()

	return err == nil
}
