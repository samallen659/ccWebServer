package server

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func (r *Router) RouteRequest(req *Request, resp *Response) {
	route, err := r.MatchRoute(req.Header.Path)
	if err != nil {
		log.Println(err)
		//TODO: write error response
		return
	}

	var html string
	if req.Header.File == "" {
		file := fmt.Sprintf("%s/index.html", route)
		b, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Failed reading file: %s. Error: %s", file, err.Error())
			//TODO: write error response
			return
		}
		html = string(b)
	} else {
		file := fmt.Sprintf("%s/%s", route, req.Header.File)
		b, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Failed reading file: %s. Error: %s", file, err.Error())
			//TODO: write error response
			return
		}
		html = string(b)
	}

	resp.SetBody(html)
	resp.SetStatus(HTTP_OK)
	fmt.Println(resp.Header)
	return
}

func (r *Router) MatchRoute(route string) (string, error) {
	if route == "/" {
		return r.wwwPath, nil
	}
	for _, rr := range r.Routes {
		if route == strings.TrimPrefix(rr, r.wwwPath) {
			log.Printf("Route matched: %s", rr)
			return rr, nil
		}
	}

	return "", errors.New("Failed to match provided route")
}
