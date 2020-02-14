package main

import (
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

var folders = []string{"./domain",
	"./{name}/repository",
	"./{name}/handler",
	"./{name}/usecase",
}

func main() {
	var svar string

	flag.StringVar(&svar, "n", "Name", "The name of your package")
	flag.Parse()
	fields := flag.Args()
	generateFolders(svar)
	generateInterfaces(svar, fields)
	generateRepository(svar)
	generateUseCase(svar)

}

func generateFolders(name string) {
	nameLower := strings.ToLower((name))
	for _, folder := range folders {
		folder := strings.Replace(folder, "{name}", nameLower, 1)
		os.MkdirAll(folder, os.ModePerm)
	}
	fmt.Println("generated folder structure")
}

func generateInterfaces(name string, args []string) {
	f, err := os.Create("./domain/" + name + ".go")
	if err != nil {
		panic("Could not create domain file for " + name)
	}
	defer f.Close()
	domain, err := ioutil.ReadFile("./templates/domain.txt")
	if err != nil {
		fmt.Println("Could not find usecases template")
		return
	}
	domainText := string(domain)
	domainText = strings.Replace(domainText, "{name}", name, -1)
	structText := generateStructString(name, args)
	domainText = strings.Replace(domainText, "{struct}", structText, -1)
	formatedText, err := format.Source([]byte(domainText))
	f.Write(formatedText)
	fmt.Println("generated domain")
}

func generateStructString(name string, args []string) string {
	s := fmt.Sprintf("type %s struct {", name)
	for _, val := range args {
		s += strings.Replace(val, ":", "\t", 1)
		s += "\n"
	}
	s += "}"
	return s
}

func generateUseCase(name string) {
	generateFromTemplate("./"+name+"/usecase/usecase.go", "./templates/useCase.txt", name)
	fmt.Println("generated usecases")
}

func generateRepository(name string) {
	generateFromTemplate("./"+name+"/repository/repository.go", "./templates/repository.txt", name)
	fmt.Println("generated repository")
}

func generateFromTemplate(file string, template string, name string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		panic("Could not create  file  " + file)
	}
	defer f.Close()
	templateRes, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Println("Could not find template" + template)
		return
	}
	templateText := string(templateRes)
	nameLower := strings.ToLower((name))
	templateText = strings.Replace(templateText, "{name}", nameLower, -1)
	templateText = strings.Replace(templateText, "{name_u}", name, -1)
	formatedText, err := format.Source([]byte(templateText))
	f.Write(formatedText)
}
