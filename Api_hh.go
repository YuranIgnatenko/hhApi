package Api_hh

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HH struct {
	data map[int][]string
}

func (hh *HH) GetData(city, spec string) []string {
	// only english search-word
	// only ru (temp.)
	passed := 3
	cntr := 0
	dataRes := make([]string, 0)
	site := fmt.Sprintf("http://%v.hh.ru/search/vacancy?clusters=true&area=77&ored_clusters=true&enable_snippets=true&salary=&text=%v", city, spec)
	res, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	r := doc.Find("div span").Text()
	dt := strings.Split(r, "\n")
	dt = strings.Split(dt[0], "Откликнуться")
	for _, val := range dt {
		cntr++
		if cntr <= passed {
			continue
		}
		val = hh.parseLine(val)
		dataRes = append(dataRes, val)
	}
	return dataRes
}

func (hh *HH) getValue(NowLine string) string {
	// assert parametres : reverse line ("456 - 321 cba") --> [abc][123][456],
	if NowLine == " " || NowLine == " " || NowLine == "" {
		return "pass"
	}
	line := NowLine
	if strings.Index(line, ".бур") != -1 {
		line = line[7:]
	}
	depos1 := ""
	depos2 := ""
	vacancy := ""
	isSpace := false
	isStop := false
	for _, v := range line {
		val := string(v)
		if isStop == false {
			switch val {
			case " ", "–", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", " ":
				if _, e := strconv.Atoi(val); e == nil {
					if isSpace == false {
						depos1 += val
					} else {
						depos2 += val
					}
				} else if val == "–" {
					isSpace = true
				} else if val == " " {
					depos1 += " "
					depos2 += " "
				} else if val == " " {
					depos1 += " "
					depos2 += " "
				}
			default:
				isStop = true
			}
		} else {
			switch val {
			case " ":
				vacancy += " "
			default:
				vacancy += val
			}
		}
	}

	depos1 = strings.Trim(depos1, " ")
	depos2 = strings.Trim(depos2, " ")
	depos1 = strings.Trim(depos1, " ")
	depos2 = strings.Trim(depos2, " ")

	depos1 = hh.reverseLine(depos1)

	depos2 = hh.reverseLine(depos2)
	vacancy = hh.reverseLine(NowLine)
	if depos2 != "" {
		indVacancy := strings.Index(vacancy, depos2)
		vacancy = vacancy[0:indVacancy]

	} else {
		indVacancy := strings.Index(vacancy, depos1)
		vacancy = vacancy[0:indVacancy]
	}
	if depos2 == "" {
		vacancy = hh.trimEnd(vacancy)
	}
	vacancy = hh.splitterString(vacancy)
	return fmt.Sprintf("[%v][%v][%v]", vacancy, depos1, depos2)
}

func (hh *HH) splitterString(line string) string {
	data := []rune(line)
	strData := make([]string, 0)
	lenData := len(data) / 2
	for i, val := range data {
		if i < lenData {
			strData = append(strData, string(val))
		}
	}
	return strings.Join(strData, "")
}

func (hh *HH) trimEnd(line string) string {
	data := []rune(line)
	strData := make([]string, 0)
	lenData := len(data) - 2
	for i, val := range data {
		if i < lenData {
			strData = append(strData, string(val))
		}
	}
	return strings.Join(strData, "")
}

func (hh *HH) reverseLine(line string) string {
	reverseLine := ""
	for _, v := range line {
		val := string(v)
		temp := ""
		switch val {
		case " ":
			symbol := " "
			temp = symbol + reverseLine
			reverseLine = temp
		default:
			symbol := val
			temp = symbol + reverseLine
			reverseLine = temp
		}
	}
	return reverseLine
}

func (hh *HH) parseLine(line string) string {
	rl := line
	if strings.Index(line, "руб.") != -1 {
		rl := hh.reverseLine(line)
		return hh.getValue(rl)
	} else {
		if ind := strings.Index(hh.splitterString(line), "Фильтры"); ind == -1 {
			return "[" + hh.splitterString(line) + "][][]"
		}
		return ""
	}
	return "error: line [" + rl + "][][]"
}

func (hh *HH) View(dt []string) {
	for i, val := range dt {
		if val != "" {
			fmt.Printf("%v:%v\n", i, val)
		}
	}
	fmt.Println()
}
