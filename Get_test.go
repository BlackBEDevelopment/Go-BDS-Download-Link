package BDSDownloadLink

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	DownloadLink, Version, err := GetWindows()
	fmt.Println(DownloadLink, Version, err)

	DownloadLink, Version, err = GetUbuntu()
	fmt.Println(DownloadLink, Version, err)
}
