package service

import (
	"fmt"
	"praktikum/dto"
	"praktikum/model"
	"praktikum/repository"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderService interface {
	Create(req *dto.CheckoutRequest) (*model.Order, error)
}

type orderServiceImpl struct {
	db            *gorm.DB
	orderRepo     repository.OrderRepository
	userRepo      repository.UserRepository
	productRepo   repository.ProductRepository
	orderItemRepo repository.OrderItemRepository
}

func NewOrderService(
	db *gorm.DB,
	orderRepo repository.OrderRepository,
	userRepo repository.UserRepository,
	productRepo repository.ProductRepository,
	orderItemRepo repository.OrderItemRepository,
) OrderService {
	return &orderServiceImpl{
		db:            db,
		orderRepo:     orderRepo,
		userRepo:      userRepo,
		productRepo:   productRepo,
		orderItemRepo: orderItemRepo,
	}
}

// Checkout Pemesanan
func (s *orderServiceImpl) Create(req *dto.CheckoutRequest) (*model.Order, error) {

	var order model.Order

	// Transactional Database otomatis
	errTx := s.db.Transaction(func(tx *gorm.DB) error {
		// Validati keberadaan user
		user, errUser := s.userRepo.FindById(tx, req.UserID)
		if errUser != nil {
			// return fmt.Errorf("User with ID [%d] not found", req.UserID)
			return errUser
		}

		// Inisialisasi total tagihan dan items array
		var totalAmount float64
		var orderItems []model.OrderItem

		// Generate nomor invoice unik
		orderNumber := fmt.Sprintf("INV/%s/%d", time.Now().Format("20060102"), time.Now().UnixNano()%10_000)

		// Simpan record order utama terlebih dahulu agar mendapatkan OrderID
		order = model.Order{
			UserID:      user.ID,
			OrderNumber: orderNumber,
			TotalAmount: 0, // Update setelah loop kalkulasi
			Status:      "COMPLETED",
		}

		if err := s.orderRepo.Create(tx, &order); err != nil {
			return err
		}

		for _, item := range req.Items {

			// Gunakan Row Level Locking (FOR UPDATE) menghindari race condition stok
			product, errLock := s.productRepo.FindByID(
				tx.Clauses(clause.Locking{Strength: "UPDATE"}),
				item.ProductID,
			)
			if errLock != nil {
				return errLock
			}

			// Validasi kecukupan stok barang
			if product.Stock < item.Quantity {
				return fmt.Errorf("Stock barang [%s] tidak mencukupi, sisa %d, permintaan %d", product.Name, product.Stock, item.Quantity)
			}

			// Update stock produck
			product.Stock = product.Stock - item.Quantity
			if err := s.productRepo.Update(tx, product); err != nil {
				return err
			}

			// Hitung subtotal dan susun object item pesanan
			subTotal := product.Price * float64(item.Quantity)
			totalAmount += subTotal

			orderItem := model.OrderItem{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			}

			if err := s.orderItemRepo.Create(tx, &orderItem); err != nil {
				return err
			}

			orderItems = append(orderItems, orderItem)

		}

		// Update total_amount akhir pada table orders
		order.TotalAmount = totalAmount
		order.Items = orderItems

		if err := s.orderRepo.Update(tx, &order); err != nil {
			return err
		}

		return nil // Transaksi sukses. GORM melakukan commit
	})

	return &order, errTx
}
