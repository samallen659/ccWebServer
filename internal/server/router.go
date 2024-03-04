package server

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

type Router struct {
	wwwPath string
	Routes  []string
}

func NewRouter(wwwPath string) (*Router, error) {
	r := &Router{wwwPath: wwwPath}
	err := r.WalkRoutes()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Router) WalkRoutes() error {
	err := filepath.WalkDir(r.wwwPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			err = fmt.Errorf("Unable to access %s err: %s", path, err)
			return err
		}
		if d.IsDir() {
			log.Printf("Adding %s to Routes", path)
			r.Routes = append(r.Routes, path)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
