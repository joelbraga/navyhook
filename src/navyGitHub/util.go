package navyGitHub
import (
	"net/http"
	"github.com/andrepinto/navyhook/src/curl"
	"fmt"
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func GetReleaseDownloadLink(url string) (string, error){
	resp, err := http.Get(url)
	var finalURL string

	if err == nil {
		finalURL = resp.Request.URL.String()
	}

	return finalURL, err
}


func DownloadDoc(url string, fileName, token string){

	hd := make(http.Header)
	hd.Add("Authorization", "token "+token)
	curl.File(
		url,
		fileName,
		func(st curl.IoCopyStat) error {
			fmt.Println(st.Stat, st.Perstr, st.Sizestr, st.Lengthstr, st.Speedstr, st.Durstr)
			//fmt.Println(st.Header["Date"])
			return nil
		},
		"maxspeed=", 3*1024*1024,
		"followredirects=", true,
		"cbinterval=", 0.5, // call the callback 0.5 second
		"header=", hd,
	)

	fmt.Println("end")
}


func Unzip(src, dest string) (error, string) {
	r, err := zip.OpenReader(src)
	var name string
	first := true
	if err != nil {
		return err, name
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err, name
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			if first{
				name = f.Name
				first = false
			}
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath,string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err, name
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err, name
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err, name
			}
		}
	}
	return nil, name
}