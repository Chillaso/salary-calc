package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	commonReductions = 0.047
	ss               = 0.0165
)

var (
	salaryFlag, irpfFlag, paysFlag *string
)

func main() {

	fmt.Println("\x1b[4m", "===== Visit https://www6.aeat.es/wlpl/PRET-R200/R210/index.zul to calc your right IRPF retention ====", "\x1b[0m")
	defineFlags()
	checkFlags()
	salary, irpf, pays := parseFlags()

	monthSalary := salary / pays
	salaryIrpf := monthSalary * (irpf / 100)
	salaryContingences := monthSalary * commonReductions
	salarySS := monthSalary * ss

	totalReduction := fmt.Sprintf("%.2f", (salaryIrpf + salaryContingences + salarySS))
	totalReductionYear := fmt.Sprintf("%.2f", ((salaryIrpf + salaryContingences + salarySS) * pays))

	netSalary := fmt.Sprintf("%.2f", (monthSalary - salaryIrpf - salarySS - salaryContingences))

	printWithColor("Total reductions per year", totalReductionYear)
	printWithColor("Gross salary per month", fmt.Sprintf("%.2f", monthSalary))
	printWithColor("Total reductions per month", totalReduction)
	printWithColor("Your net salary is", netSalary)
}

func defineFlags() {

	salaryFlag = flag.String("salary", "-1", "Your gross salary")
	irpfFlag = flag.String("irpf", "-1", "Your irpf percent")
	paysFlag = flag.String("pays", "12", "Number of pays")

	flag.Parse()
}

func checkFlags() {

	if *salaryFlag == "-1" || *irpfFlag == "-1" {
		fmt.Println("use --salary XXXX --irpf XXXX --pays 12 params to calc your salary")
		fmt.Println("if you dont define --pays will be automatically set to 12.")
		os.Exit(1)
	}
}

func parseFlags() (salary, irpf, pays float64) {
	salary, salaryerr := strconv.ParseFloat(*salaryFlag, 64)
	irpf, irpferr := strconv.ParseFloat(*irpfFlag, 64)
	pays, payserr := strconv.ParseFloat(*paysFlag, 64)
	if salaryerr != nil || irpferr != nil || payserr != nil {
		fmt.Println("Error in flags, please fix it")
	}
	return salary, irpf, pays
}

func printWithColor(text, toColor string) {
	fmt.Println(text, "\x1b[33m-", toColor, "\x1b[0m")
}
