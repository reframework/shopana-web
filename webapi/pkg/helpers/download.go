package helpers

import (
	"io"
	"net/http"
	"os"
)

func Download(url string) (int64, error) {
	out, err := os.Create("output.txt")
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	return n, err
}
