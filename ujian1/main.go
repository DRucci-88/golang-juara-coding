package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*
day5-abstract-struktur-data-dinamis.pdf
Sesi 6: Soal Ujian Akhir Modul (Studi Kasus HRIS)
Studi Kasus: Sistem Manajemen Kepegawaian & Payroll Terintegrasi
*/

var (
	ErrNotNegativeValue    = errors.New("value tidak boleh negatif")
	ErrEmptyEmployeeName   = errors.New("nama karyawan tidak boleh kosong")
	ErrDuplicateEmployeeID = errors.New("karyawan dengan ID tersebut sudah terdaftar")
	ErrPercentageInvalid   = errors.New("tidak bisa dibawah 0.0 dan diatas 1.0")
)

type EmployeeTypeEnum string

const (
	FULLTIME  EmployeeTypeEnum = "FullTime"
	CONTRACT  EmployeeTypeEnum = "Contract"
	FREELANCE EmployeeTypeEnum = "Freelance"
)

type PayrollCalculator interface {
	CalculateSalary() (float64, error)
	GetEmployeeType() EmployeeTypeEnum
}

type FullTimeEmployee struct {
	BaseSalary float64
	Allowance  float64
	TaxRate    float64
}

func NewFullTimeEmployee(
	baseSalary float64,
	allowance float64,
	taxRate float64,
) (*FullTimeEmployee, error) {
	if baseSalary < 0.0 || allowance < 0.0 {
		return nil, fmt.Errorf("MonthlyRate dan PerformanceBonus [%w]", ErrNotNegativeValue)
	}
	if taxRate < 0.0 || taxRate > 1.0 {
		return nil, fmt.Errorf("TaxRate %w", ErrPercentageInvalid)
	}
	return &FullTimeEmployee{
		BaseSalary: baseSalary,
		Allowance:  allowance,
		TaxRate:    taxRate,
	}, nil
}
func (emp *FullTimeEmployee) GetEmployeeType() EmployeeTypeEnum {
	return FULLTIME
}
func (emp *FullTimeEmployee) CalculateSalary() (float64, error) {
	return (emp.BaseSalary + emp.Allowance) * (1.0 - emp.TaxRate), nil
	// return 0.0, nil
}

type ContractEmployee struct {
	MonthlyRate      float64
	PerformanceBonus float64
}

func NewContractEmployee(
	monthlyRate float64,
	performanceBonus float64,
) (*ContractEmployee, error) {
	if monthlyRate < 0.0 || performanceBonus < 0.0 {
		return nil, fmt.Errorf("MonthlyRate dan PerformanceBonus [%w]", ErrNotNegativeValue)
	}
	return &ContractEmployee{
		MonthlyRate:      monthlyRate,
		PerformanceBonus: performanceBonus,
	}, nil
}
func (emp *ContractEmployee) GetEmployeeType() EmployeeTypeEnum {
	return CONTRACT
}
func (emp *ContractEmployee) CalculateSalary() (float64, error) {
	return emp.MonthlyRate + emp.PerformanceBonus, nil
	// return 0.0, nil
}

type Freelancer struct {
	HourlyRate  float64
	HoursWorked int
}

func NewFreelancer(
	hourlyRate float64,
	hoursWorked int,
) (*Freelancer, error) {
	if hourlyRate < 0.0 || hoursWorked < 0 {
		return nil, fmt.Errorf("HourlyRate dan HourlyRate [%w]", ErrNotNegativeValue)
	}
	return &Freelancer{
		HourlyRate:  hourlyRate,
		HoursWorked: hoursWorked,
	}, nil
}
func (emp *Freelancer) GetEmployeeType() EmployeeTypeEnum {
	return FREELANCE
}
func (emp *Freelancer) CalculateSalary() (float64, error) {
	return emp.HourlyRate * float64(emp.HoursWorked), nil
	// return 0.0, nil
}

type HRIS struct {
	Employees map[string]string            // EmployeeID -> Nama Karyawan
	Payrolls  map[string]PayrollCalculator // EmployeeID -> PayrollCalculator
}

func (hris *HRIS) RegisterEmployee(
	id string, // EmployeeID
	name string,
	payroll PayrollCalculator,
) error {
	if len(name) <= 0 {
		return ErrEmptyEmployeeName
	}
	_, existEmployee := hris.Employees[id]
	_, existPayroll := hris.Payrolls[id]
	if existEmployee || existPayroll { // Jika ID tersebut sudah terdaftar di 2 map tersebut
		return fmt.Errorf("EmployeeID [%s] %w", id, ErrDuplicateEmployeeID)
	}
	hris.Employees[id] = name
	hris.Payrolls[id] = payroll
	return nil
}

func (hris *HRIS) CalculateTotalPayout() float64 {
	if len(hris.Employees) != len(hris.Payrolls) {
		fmt.Println("WADUH Jumlah map tidak sama")
		fmt.Println(hris.Employees)
		fmt.Println(hris.Payrolls)
		return 0.0
	}

	var totalSalary float64
	for _, payroll := range hris.Payrolls {
		if salary, err := payroll.CalculateSalary(); err == nil {
			totalSalary += salary
		}

	}
	return totalSalary
}

func (hris *HRIS) PrintPayrollReport() {
	fmt.Println("\n========== PAYROLL REPORT ==========")

	for employeeID, payroll := range hris.Payrolls {
		name := hris.Employees[employeeID]

		salary, err := payroll.CalculateSalary()
		if err != nil {
			fmt.Printf(
				"[ERROR] Gagal menghitung gaji karyawan %s: %v\n",
				name,
				err,
			)
			continue
		}

		fmt.Println("------------------------------------")
		fmt.Printf("Employee ID   : %s\n", employeeID)
		fmt.Printf("Nama          : %s\n", name)
		fmt.Printf("Tipe          : %s\n", payroll.GetEmployeeType())

		// Type Switch
		switch emp := payroll.(type) {

		case *FullTimeEmployee:
			fmt.Printf("Base Salary   : Rp %.2f\n", emp.BaseSalary)
			fmt.Printf("Allowance     : Rp %.2f\n", emp.Allowance)
			fmt.Printf("Tax Rate      : %.2f%%\n", emp.TaxRate*100)

		case *ContractEmployee:
			fmt.Printf("Monthly Rate  : Rp %.2f\n", emp.MonthlyRate)
			fmt.Printf("Bonus         : Rp %.2f\n", emp.PerformanceBonus)

		case *Freelancer:
			fmt.Printf("Hourly Rate   : Rp %.2f\n", emp.HourlyRate)
			fmt.Printf("Hours Worked  : %d jam\n", emp.HoursWorked)

		default:
			fmt.Println("Unknown Employee Type")
		}

		fmt.Printf("Total Salary  : Rp %.2f\n", salary)
	}

	fmt.Println("====================================")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	hris := &HRIS{
		Employees: make(map[string]string),
		Payrolls:  make(map[string]PayrollCalculator),
	}

MainLoop:
	for {
		clearConsole()

		printMainMenu()

		menu := readInt(reader, "Pilih menu: ")

		switch menu {

		case 0:
			fmt.Println("Terima kasih telah menggunakan HRIS System")
			break MainLoop

		case 1:
			menuRegisterFullTime(reader, hris)
			pause(reader)

		case 2:
			menuRegisterContract(reader, hris)
			pause(reader)

		case 3:
			menuRegisterFreelancer(reader, hris)
			pause(reader)

		case 4:
			clearConsole()
			hris.PrintPayrollReport()
			pause(reader)

		case 5:
			clearConsole()
			fmt.Printf(
				"\nTOTAL PAYOUT: Rp %.2f\n",
				hris.CalculateTotalPayout(),
			)
			pause(reader)

		case 6:
			seedDummyEmployees(hris)
			fmt.Println("Dummy employees berhasil ditambahkan")
			pause(reader)

		default:
			fmt.Println("Menu tidak valid")
			pause(reader)
		}
	}
}
func printMainMenu() {
	fmt.Println("====================================")
	fmt.Println(" HRIS & PAYROLL MANAGEMENT SYSTEM")
	fmt.Println("====================================")
	fmt.Println("1. Register FullTime Employee")
	fmt.Println("2. Register Contract Employee")
	fmt.Println("3. Register Freelancer")
	fmt.Println("4. Show Payroll Report")
	fmt.Println("5. Show Total Payout")
	fmt.Println("6. Seed Dummy Employees")
	fmt.Println("0. Exit")
	fmt.Println("====================================")
}

func pause(reader *bufio.Reader) {
	fmt.Println("\nTekan ENTER untuk melanjutkan...")
	reader.ReadString('\n')
}

func menuRegisterFullTime(
	reader *bufio.Reader,
	hris *HRIS,
) {

	id := readString(reader, "Employee ID: ")
	name := readString(reader, "Nama: ")
	baseSalary := readFloat(reader, "Base Salary: ")
	allowance := readFloat(reader, "Allowance: ")
	taxRate := readFloat(reader, "Tax Rate (0.05 = 5%): ")

	employee, err := NewFullTimeEmployee(
		baseSalary,
		allowance,
		taxRate,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = hris.RegisterEmployee(
		id,
		name,
		employee,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("FullTime Employee berhasil didaftarkan")
}

func seedDummyEmployees(hris *HRIS) {
	fulltime, _ := NewFullTimeEmployee(
		10_000_000,
		2_000_000,
		0.05,
	)

	contract, _ := NewContractEmployee(
		7_000_000,
		1_000_000,
	)

	freelancer, _ := NewFreelancer(
		150_000,
		80,
	)

	hris.RegisterEmployee(
		"EMP001",
		"Le Rucco",
		fulltime,
	)

	hris.RegisterEmployee(
		"EMP002",
		"Dewa",
		contract,
	)

	hris.RegisterEmployee(
		"EMP003",
		"Norber",
		freelancer,
	)
}

func menuRegisterContract(
	reader *bufio.Reader,
	hris *HRIS,
) {

	id := readString(reader, "Employee ID: ")
	name := readString(reader, "Nama: ")

	monthlyRate := readFloat(
		reader,
		"Monthly Rate: ",
	)

	performanceBonus := readFloat(
		reader,
		"Performance Bonus: ",
	)

	employee, err := NewContractEmployee(
		monthlyRate,
		performanceBonus,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = hris.RegisterEmployee(
		id,
		name,
		employee,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Contract Employee berhasil didaftarkan")
}

func menuRegisterFreelancer(
	reader *bufio.Reader,
	hris *HRIS,
) {

	id := readString(reader, "Employee ID: ")
	name := readString(reader, "Nama: ")

	hourlyRate := readFloat(
		reader,
		"Hourly Rate: ",
	)

	hoursWorked := readInt(
		reader,
		"Hours Worked: ",
	)

	employee, err := NewFreelancer(
		hourlyRate,
		hoursWorked,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = hris.RegisterEmployee(
		id,
		name,
		employee,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Freelancer berhasil didaftarkan")
}

// Helper untuk membaca input bertipe string secara bersih
func readString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca input bertipe integer secara bersih
func readInt(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka bulat.")
	}
}

// Helper untuk membaca input bertipe float64 secara bersih
func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readString(reader, prompt)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka desimal/nominal.")
	}
}

func clearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
