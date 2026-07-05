package service

import (
	"context"
	"log"
	"time"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type SalaryService struct {
	repo          *repository.RepositoryManager
	salaryRepo    *repository.SalaryRepository
	employeeRepo  *repository.EmployeeRepository
	attedanceRepo *repository.AttendanceRepository
	leaveRepo     *repository.LeaveRepository
}

func NewSalaryService(
	repo *repository.RepositoryManager,
) *SalaryService {
	return &SalaryService{
		repo:          repo,
		salaryRepo:    repo.Salary(),
		attedanceRepo: repo.Attendance(),
		leaveRepo:     repo.Leave(),
		employeeRepo:  repo.Employee(),
	}
}

func (s *SalaryService) Calculate(
	ctx context.Context,
	period time.Time,
) (*string, error) {
	emp, err1 := s.employeeRepo.FindByID(ctx, 1, model.EmployeePreloadPosition)
	log.Printf("%+v", emp)
	log.Println(err1)
	log.Printf("Calculate Payroll, Period [%s]", period)

	// return nil, err1
	// err := s.employeeRepo.ProcessInBatches(ctx, 100, func(employees []model.Employee, batch int) error {
	// 	return s.processBatch(ctx, period, batch, employees)
	// })

	err := s.employeeRepo.ProcessWithoutPayrollInBatches(ctx, 20, period, func(employees []model.Employee, batch int) error {
		return s.processBatch(ctx, period, batch, employees)
	})
	return nil, err
}

func (s *SalaryService) processBatch(
	ctx context.Context,
	period time.Time,
	batch int,
	employees []model.Employee,
) error {
	startDate := period
	endDate := period.AddDate(0, 1, 0)
	log.Printf("Batch No [%d], Employee Count %d", batch, len(employees))
	err := s.repo.Transaction(ctx, func(repo *repository.RepositoryManager) error {
		salaries := make([]model.Salary, 0, len(employees))
		for i := range employees {
			emp := &employees[i]
			log.Printf("GIRUNJAY [%d] %+v", i, emp)
			summary, err := repo.Attendance().SummaryForPayroll(ctx, startDate, endDate, emp.ID)
			if err != nil {
				log.Printf("Batch [%d], Employee ID [%d] Failed to Payroll . %+v\n", batch, emp.ID, err)
				return err
			}

			leaveApprovedCount, err := repo.Leave().CountBetweenDateAndStatus(ctx, emp.ID, startDate, endDate, model.LeaveStatusApproved)

			if err != nil {
				log.Printf("Batch [%d], Employee ID [%d] Failed to Payroll . %+v\n", batch, emp.ID, err)
				return err
			}

			empPosition, err := repo.Employee().FindByID(ctx, emp.ID, model.EmployeePreloadPosition)
			if err != nil {
				log.Printf("Batch [%d], Employee ID [%d] Failed to Payroll . %+v\n", batch, emp.ID, err)
				return err
			}

			allowance := float64(summary.Present) * 50_000.0
			deduction := float64(summary.Absent)*100_000.0 + float64(summary.Late)*20_000.0 + float64(leaveApprovedCount)*50_000
			netSalary := empPosition.Position.BaseSalary + allowance - deduction

			salaries = append(salaries, model.Salary{
				EmployeeID:  emp.ID,
				Period:      period,
				BasicSalary: empPosition.Position.BaseSalary,
				Allowance:   allowance,
				Deductions:  deduction,
				NetSalary:   netSalary,
			})
		}
		return repo.Salary().CreateInBatches(ctx, 10, salaries)
	})
	return err
}

/*
D. Proses Payroll Bulanan (Perhitungan Dinamis)
POST /api/salaries/calculate :
a/ Endpoint untuk menghitung dan menyimpan rekap gaji
bersih karyawan untuk periode tertentu.

Di dalam proses ini:
1. Ambil data gaji pokok berdasarkan jabatan karyawan.
2. Hitung jumlah kehadiran dari tabel  Attendances  (untuk menentukan  Allowance  dan Deductions ).
3. Kurangi saldo jatah cuti jika ada cuti yang disetujui ( APPROVED ) di tabel  Leaves .
4. Simpan hasil akhir ke tabel  Salaries .
*/
