package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

type File struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

func jsonFileReader(f *File) error {
	data, err := os.ReadFile("file.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, f)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	var f File
	err := jsonFileReader(&f)
	if err != nil {
		panic(err)
	}

	client, err := ftp.Dial(fmt.Sprintf("%s:%d", f.Server, f.Port), ftp.DialWithTimeout(10))
	if err != nil {
		panic(err)
	}

	defer client.Quit()

	err = client.Login(f.User, f.Password)
	if err != nil {
		panic(err)
	}

	files, err := client.List("/")
	if err != nil {
		panic(err)
	}

	fmt.Println("List of files:")
	for _, file := range files {
		fmt.Println(file.Name)
	}
}
