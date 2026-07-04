package service

import (
	"context"
	"log"
	"time"
	"ujian2_rematch/helper"
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
	}
}

func (s *SalaryService) Calculate(
	ctx context.Context,
	year int,
	month int,
) (*string, error) {
	period := helper.PayrollPeriod(year, time.Month(month))
	// err := s.employeeRepo.ProcessInBatches(ctx, 100, func(employees []model.Employee, batch int) error {
	// 	return s.processBatch(ctx, period, batch, employees)
	// })
	err := s.employeeRepo.ProcessWithoutPayrollInBatches(ctx, 100, period, func(employees []model.Employee, batch int) error {
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

			allowance := float64(summary.Present) * 50_000.0
			deduction := float64(summary.Absent)*100_000.0 + float64(summary.Late)*20_000.0 + float64(leaveApprovedCount)*50_000
			netSalary := emp.Position.BaseSalary + allowance - deduction

			salaries = append(salaries, model.Salary{
				EmployeeID:  emp.ID,
				Period:      period,
				BasicSalary: emp.Position.BaseSalary,
				Allowance:   allowance,
				Deductions:  deduction,
				NetSalary:   netSalary,
			})
		}
		return repo.Salary().CreateInBatches(ctx, 50, salaries)
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
