package test

import (
	"log"
	"testing"
)



func TestPrint(t *testing.T) {
	log.Print("testPrint")
}

func TestPrintf(t *testing.T) {
	log.Printf("testPrintf")
}

func TestFatal(t *testing.T) {
	log.Fatal("Fatal")
}