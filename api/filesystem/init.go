package filesystem

//Files

func Rp_readFile(file string) string {

	content, err := ioutil.ReadFile(file) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	return string(content)

}

func Rp_createFile(file string) {

	_, err := os.Create(file)

	if err != nil {
		log.Fatal(err)
	}

}

func Rp_writeFile(file, content string) {

	f, err := os.Create(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func Rp_writeToFile(file, content string) {

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func Rp_removeFile(file string) {

	os.Remove(file)

}

func Rp_removeDir(Dir string) {

	os.RemoveAll(Dir)

}
