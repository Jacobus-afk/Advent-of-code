package aoc

import (
	"bufio"
	"log"
	"os"
)

func ReadFileByLine(path string) []string {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  output := []string{}

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    output = append(output, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return output
}
