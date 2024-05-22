package main

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"strings"
)

type Migration struct {
	Id string
	Statements []string
}

type MigrationProvider interface {
	GetMigrations() []Migration
}

type FileMigrationProvider struct {
	directory string
}

func (f *FileMigrationProvider) GetMigrations() []Migration {
	files, err := os.ReadDir(f.directory)
	if err != nil {
		log.Fatal(err)
	}
	var migrations []Migration
	for _, file := range files {
		migration := readMigrationFromFile(f.directory, file.Name())
		migrations= append(migrations, migration)
	}
	return migrations
}

func readMigrationFromFile(filePath string, fileName string) Migration {
	fullPath := fmt.Sprintf("%s/%s", filePath, fileName)
	file, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	id := strings.Split(fileName, ".")[0]
	return Migration{Id: id, Statements: lines}
}

