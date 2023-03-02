package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

var buffer bytes.Buffer

func main() {
	//diff("sometool/excel/UBC807.xlsx")
	//diff("sometool/excel/b.xlsx")
	//diff("sometool/excel/c.xlsx")
	//diff("sometool/excel/d.xlsx")
	//diff("sometool/excel/f.xlsx")
	//diff("sometool/excel/d.xlsx")
	//diff("sometool/excel/h.xlsx")
	//diff("sometool/excel/i.xlsx")
	//diff("sometool/excel/j.xlsx")
	//diff("sometool/excel/k.xlsx")
	//diff("sometool/excel/l.xlsx")
	//diff("sometool/excel/m.xlsx")
	//diffAll("sometool/excel/图谱.xlsx")
	diffAll("sometool/excel/指纹图谱.xlsx")
	os.WriteFile("sometool/excel/指纹图谱.txt", buffer.Bytes(), 777)
}

func diff(excel string) {
	buffer.WriteString("==============" + excel + "================ \n")
	x, err := xlsx.OpenFile(excel)
	if err != nil {
		return
	}
	s := x.Sheets[0]
	data := make([]string, 0)
	for _, r := range s.Rows {
		m := ""
		for _, cell := range r.Cells {
			m += cell.Value
		}
		data = append(data, m)
	}
	cnt := 0
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			record := data[i]
			other := data[j]
			if record == other {
				cnt++
				buffer.WriteString("两条完全相同的记录 \n")
				buffer.WriteString(fmt.Sprintf("A列号:%d 结果 %+v \n", i+1, record))
				buffer.WriteString(fmt.Sprintf("B列号:%d 结果 %+v \n", j+1, other))
			} else {
				//cnt := 01
				//for i, c := range record {
				//	for j, c2 := range other {
				//		if i == j && c == c2 {
				//			cnt++
				//		}
				//	}
				//}
				//fmt.Println("两条不同的记录")
				//fmt.Printf("A:%d %+v \n", i, record)
				//fmt.Printf("B:%d %+v \n", j, other)
			}
		}
	}
	if cnt == 0 {
		buffer.WriteString("这个excel没有相同的记录！！！！" + excel)
	}
}

func diffAll(excel string) {
	x, err := xlsx.OpenFile(excel)
	if err != nil {
		return
	}
	summ := ""
	for _, sheet := range x.Sheets {
		if len(sheet.Name) == 1 {
			buffer.WriteString("==============" + sheet.Name + "================ \n")

			data := make([]string, 0)
			for _, r := range sheet.Rows {
				m := ""
				for i := 2; i < len(r.Cells); i++ {
					m += r.Cells[i].Value
				}
				data = append(data, m)
			}
			cnt := 0
			for i := 0; i < len(data); i++ {
				for j := i + 1; j < len(data); j++ {
					record := data[i]
					other := data[j]
					if record == other {
						cnt++
						buffer.WriteString("两条完全相同的记录 \n")
						buffer.WriteString(fmt.Sprintf("A列号:%d 结果 %+v \n", i, record))
						buffer.WriteString(fmt.Sprintf("B列号:%d 结果 %+v \n", j, other))
					} else {
						//cnt := 01
						//for i, c := range record {
						//	for j, c2 := range other {
						//		if i == j && c == c2 {
						//			cnt++
						//		}
						//	}
						//}
						//fmt.Println("两条不同的记录")
						//fmt.Printf("A:%d %+v \n", i, record)
						//fmt.Printf("B:%d %+v \n", j, other)
					}
				}
			}
			if cnt == 0 {
				buffer.WriteString("这个excel没有相同的记录！！！！" + excel)
			}
			summ += fmt.Sprintf("%ssheet有%d条相同的记录 \n", sheet.Name, cnt)
		}
	}
	buffer.WriteString(summ)

}
