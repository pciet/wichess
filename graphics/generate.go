// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	incremental_offset = 0.01
	incremental_rotate = 6

	camera_height = 2.5

	camera_file = "camera.inc"
	viewer_file = "viewer.html"
	output_dir  = "img"

	povray_width     = 256
	povray_height    = 256
	povray_quality   = 11
	povray_antialias = "off"
)

// Generate a viewer page, then generate camera.inc and render 64 times.
func main() {
	t, err := template.ParseFiles(viewer_file)
	if err != nil {
		log.Panicln(err.Error())
	}
	err = os.Mkdir(output_dir, 0777)
	if (err != nil) && (os.IsExist(err) == false) {
		log.Panicln(err.Error())
	}
	for _, file := range os.Args[1:] {
		name := file[:len(file)-len(".pov")]
		err = generateViewer(t, name)
		if err != nil {
			log.Panicln(err.Error())
		}
		err = generateImages(name)
		if err != nil {
			log.Panicln(err.Error())
		}
	}
}

type viewerTemplate struct {
	Name string
}

// Assumes only .pov files.
func generateViewer(templ *template.Template, forPiece string) error {
	f, err := os.Create(fmt.Sprintf("%v/%v_%v", output_dir, forPiece, viewer_file))
	if err != nil {
		return err
	}
	return templ.Execute(f, viewerTemplate{
		Name: forPiece,
	})
}

func generateImages(forPiece string) error {
	for i := 0; i < 64; i++ {
		err := generateCameraInc(i)
		if err != nil {
			return err
		}
		cmd := exec.Command("/usr/local/bin/povray",
			"Output_To_File=true",
			"Output_File_Type=N",
			fmt.Sprintf("Output_File_Name=%v/%v_%v.png", output_dir, forPiece, i),
			fmt.Sprintf("-w%v", povray_width),
			fmt.Sprintf("-h%v", povray_height),
			fmt.Sprintf("Quality=%v", povray_quality),
			fmt.Sprintf("Antialias=%v", povray_antialias),
			fmt.Sprintf("%v.pov", forPiece))
		log.Printf("%+v", cmd.Args)
		pl, err := cmd.CombinedOutput()
		if err != nil {
			scanner := bufio.NewScanner(strings.NewReader(string(pl)))
			for scanner.Scan() {
				log.Println(scanner.Text())
			}
			return err
		}
	}
	return nil
}

func generateCameraInc(forPoint int) error {
	if (forPoint > 63) || (forPoint < 0) {
		return errors.New(fmt.Sprintf("generate: point %v out of range", forPoint))
	}
	f, err := os.Create(camera_file)
	if err != nil {
		return err
	}
	file := float64(forPoint % 8)
	rank := float64(forPoint / 8)
	var xOffset, zOffset, xRotate, yRotate float64
	if rank < 4 {
		zOffset = (4 - rank) * incremental_offset
		xRotate = (4 - rank) * incremental_rotate
		if file < 4 { // x,z
			xOffset = (4 - file) * incremental_offset
			yRotate = (4 - file) * incremental_rotate * -1
		} else { // -x,z
			xOffset = (file - 3) * incremental_offset * -1
			yRotate = (file - 3) * incremental_rotate
		}
	} else {
		zOffset = (rank - 3) * incremental_offset * -1
		xRotate = (rank - 3) * incremental_rotate * -1
		if file < 4 { // x,-z
			xOffset = (4 - file) * incremental_offset
			yRotate = (4 - file) * incremental_rotate * -1
		} else { // -x,-z
			xOffset = (file - 3) * incremental_offset * -1
			yRotate = (file - 3) * incremental_rotate
		}
	}
	_, err = f.Write([]byte(fmt.Sprintf("camera {\nrotate <%v,0,%v>\nlocation <0,%v,0>\nlook_at <0,0,0>\ntranslate <%v,0,%v>\n}", xRotate, yRotate, camera_height, xOffset, zOffset)))
	if err != nil {
		return err
	}
	return f.Close()
}
