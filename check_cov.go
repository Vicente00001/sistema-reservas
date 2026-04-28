package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("cov.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var totalStmts int
	var coveredStmts int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mode:") {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			continue
		}
		
		stmtsAndCount := strings.Split(parts[1], " ")
		if len(stmtsAndCount) > 0 {
		    // The line format is: name.go:line.column,line.column numberOfStatements count
		    // Wait, it is: name.go:line.column,line.column numberOfStatements count
		}
		
		// Let's just use simple parsing
		fields := strings.Fields(line)
		if len(fields) != 3 {
		    continue
		}
		stmts, _ := strconv.Atoi(fields[1])
		count, _ := strconv.Atoi(fields[2])
		
		totalStmts += stmts
		if count > 0 {
			coveredStmts += stmts
		}
	}

	if totalStmts == 0 {
		fmt.Println("0.0%")
		return
	}

	fmt.Printf("Total statements: %d, Covered: %d\n", totalStmts, coveredStmts)
	fmt.Printf("Total coverage: %.1f%%\n", float64(coveredStmts)/float64(totalStmts)*100)
}
