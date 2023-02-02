package layout

import (
	"fmt"
	"os"
)

type Project struct {
	Path string
	Name string
}

// func New(path string, name string) *Project {
// 	return &Project{
// 		Path: path,
// 		Name: name,
// 	}
// }


func (p *Project) CreateStructure() error {
	
	err := os.Mkdir(p.Name, 0777)
	if err != nil {
		return fmt.Errorf("error in creating directory internal")
	}

	// TODO: create check path 
	err = os.Chdir(p.Path + "/" + p.Name)
	if err != nil {
		return fmt.Errorf("error in change work directory for genos")
	}
	return nil
}

