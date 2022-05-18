package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strings"
)

func main() {
	guidsMap := make(map[string]string)
	file, err := os.Open("/home/roman/git/LeetCode/work/correspondenceTable/tmp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		t := scanner.Text()
		ind := strings.Index(t, " ")
		if ind == -1 {
			fmt.Println("bag")
		}
		guidsMap[t[0:ind]] = t[ind+1:]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	csv, err := os.Open("/home/roman/git/LeetCode/work/correspondenceTable/csvl.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvScanner := bufio.NewScanner(csv)

	newFile, err := os.Create("newFile.txt")
	datawriter := bufio.NewWriter(newFile)
	defer newFile.Close()
	i := 0
	for csvScanner.Scan() {
		if i == 0 {
			i++
			_, _ = datawriter.WriteString(csvScanner.Text() + "\n")
			continue
		}
		t := csvScanner.Text()

		t1 := t[strings.Index(t, ";")+1:]
		t2 := t1[strings.Index(t1, ";")+1:]
		t3 := t2[strings.Index(t2, ";")+1:]
		ind := strings.Index(t3, ";")
		number := t3[:ind]
		guid := guidsMap[number]
		if guid == "" {
			guid = "null"
		}
		_, _ = datawriter.WriteString(t + guid + "\n")
	}
	datawriter.Flush()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*func readCorrespondenceTable(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Err(err).Msg("open file correspondence_table_fixture.csv failed")
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Err(err).Msg("error while closing correspondence_table_fixture.csv file")
		}
	}()

	r := csv.NewReader(f)
	r.Comma = ';'
	rows, err := r.ReadAll()
	if err != nil {
		log.Err(err).Msg("read file correspondence_table_fixture failed")
	}
	return rows
}
*/
