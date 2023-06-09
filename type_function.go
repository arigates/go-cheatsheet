package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunction func(string) string

type Server struct {
	filenameTransformFunc TransformFunction
}

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenameTransformFunc(filename)

	fmt.Println("new filename: ", newFilename)

	return nil
}

func hashFileName(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])

	return newFilename
}

func prefixFilename(prefix string) TransformFunction {
	return func(filename string) string {
		return prefix + filename
	}
}

func main() {
	s := &Server{
		filenameTransformFunc: prefixFilename("test_"),
	}

	s.handleRequest("cool_picture.jpg")
}
