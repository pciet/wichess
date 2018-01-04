// Copyright 2018 Matthew Juran
// All Rights Reserved

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

const (
	PieceCount = 80000
	Scale      = 10

	// the camera is always negative x, negative y, positive z, looking at 0,0,0
	CameraX = -100
	CameraY = -100
	CameraZ = 100

	File = `#include "header.inc"
#include "camera.inc"
#include "finishes.inc"
#include "piece.inc"
`

	InputName = "album.pov"
	Width     = 1600
	Height    = 1600
	Quality   = 11
	Antialias = "on"
	Name      = "album.png"
	Filetype  = "N"
)

func main() {
	rand.Seed(129)
	f, err := os.Create(InputName)
	if err != nil {
		panic(err.Error())
	}
	out := bytes.NewBufferString(File)
	for i := 0; i < PieceCount; i++ {
		x := rand.Intn(int(math.Abs(CameraX))) * -1
		y := rand.Intn(int(math.Abs(CameraY))) * -1
		z := rand.Intn(int(math.Abs(CameraZ)))
		xRot := rand.Intn(360)
		yRot := rand.Intn(360)
		zRot := rand.Intn(360)
		s := rand.Intn(Scale)
		_, err := out.WriteString(fmt.Sprintf("object{%v translate <%v,%v,%v> rotate <%v,%v,%v> scale %v}\n", randomPiece(), x, y, z, xRot, yRot, zRot, s+20))
		if err != nil {
			panic(err.Error())
		}
	}
	_, err = f.Write(out.Bytes())
	if err != nil {
		panic(err.Error())
	}
	cmd := exec.Command("/usr/local/bin/povray",
		"Output_To_File=true",
		fmt.Sprintf("Output_File_Type=%v", Filetype),
		fmt.Sprintf("Output_File_Name=%v", Name),
		fmt.Sprintf("-w%v", Width),
		fmt.Sprintf("-h%v", Height),
		fmt.Sprintf("Quality=%v", Quality),
		fmt.Sprintf("Antialias=%v", Antialias),
		InputName)
	log.Printf("%+v", cmd.Args)
	pl, err := cmd.CombinedOutput()
	if err != nil {
		scanner := bufio.NewScanner(strings.NewReader(string(pl)))
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
		panic(err.Error())
	}
}

func randomPiece() string {
	switch rand.Intn(10) {
	case 0:
		return "Bishop texture {BlackT}"
	case 1:
		return "Bishop texture {WhiteT}"
	case 2:
		return "King texture {BlackT}"
	case 3:
		return "King texture {WhiteT}"
	case 4:
		return "Pawn texture {BlackT}"
	case 5:
		return "Pawn texture {WhiteT}"
	case 6:
		return "Queen texture {BlackT}"
	case 7:
		return "Queen texture {WhiteT}"
	case 8:
		return "Rook texture {BlackT}"
	case 9:
		return "Rook texture {WhiteT}"
	}
	panic("rng out of bounds")
	return ""
}
