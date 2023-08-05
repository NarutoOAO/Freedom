package util

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"9900project/conf"
)

// UploadMaterialToLocalStatic1 upload material
func UploadMaterialToLocalStatic1(file multipart.File, courseNumber int, materialName string) (filePath string, err error) {
	cn := strconv.Itoa(courseNumber)
	basePath := "." + conf.CoursePath + cn + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	coursePath := basePath + materialName + ".pdf"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(coursePath, content, 0666)
	if err != nil {
		return "", err
	}
	return cn + "/" + materialName + ".pdf", err
}

// UploadMaterialToLocalStatic2 upload material
func UploadMaterialToLocalStatic2(file multipart.File, courseNumber int, materialName string) (filePath string, err error) {
	cn := strconv.Itoa(courseNumber)
	basePath := "." + conf.CoursePath + cn + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	coursePath := basePath + materialName + ".ppt"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(coursePath, content, 0666)
	if err != nil {
		return "", err
	}
	return cn + "/" + materialName + ".ppt", err
}

// UploadAvatarToLocalStatic upload avatar
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", err
}

// UploadAssignmentToLocalStatic upload assignment
func UploadAssignmentToLocalStatic(file multipart.File, courseNumber int, materialName string) (filePath string, err error) {
	cn := strconv.Itoa(courseNumber)
	basePath := "." + conf.AssignmentPath + cn + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	coursePath := basePath + materialName + ".pdf"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(coursePath, content, 0666)
	if err != nil {
		return "", err
	}
	return cn + "/" + materialName + ".pdf", err
}

// UploadAssSolutionToLocalStatic upload assignment solution
func UploadAssSolutionToLocalStatic(file multipart.File, courseNumber int, materialName string) (filePath string, err error) {
	cn := strconv.Itoa(courseNumber)
	basePath := "." + conf.AssSolutionPath + cn + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	coursePath := basePath + materialName + ".pdf"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(coursePath, content, 0666)
	if err != nil {
		return "", err
	}
	return cn + "/" + materialName + ".pdf", err
}

// DirExistOrNot
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
