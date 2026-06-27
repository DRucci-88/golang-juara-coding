package service

import (
	"fmt"
	"log"
	"praktikum/dto"
	"praktikum/model"
	"praktikum/repository"
	"runtime/debug"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderService interface {
	Create(req *dto.CheckoutRequest) (*model.Order, error)
	FindById(id int) (*model.Order, error)
	Cancel(id int) (*model.Order, error)
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
	log.Println(req)

	// Transactional Database otomatis
	errTx := s.db.Transaction(func(tx *gorm.DB) error {
		// Validati keberadaan user
		user, errUser := s.userRepo.FindById(tx, req.UserID)
		if errUser != nil {
			log.Println(errUser)
			debug.PrintStack()
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
			log.Println(err)
			debug.PrintStack()
			return err
		}

		for _, item := range req.Items {

			// Gunakan Row Level Locking (FOR UPDATE) menghindari race condition stok
			product, errLock := s.productRepo.FindByID(
				tx.Clauses(clause.Locking{Strength: "UPDATE"}),
				item.ProductID,
			)
			if errLock != nil {
				log.Println(errLock)
				debug.PrintStack()
				return errLock
			}

			// Validasi kecukupan stok barang
			if product.Stock < item.Quantity {
				return fmt.Errorf("Stock barang [%s] tidak mencukupi, sisa %d, permintaan %d", product.Name, product.Stock, item.Quantity)
			}

			// Update stock produck
			product.Stock = product.Stock - item.Quantity
			if err := s.productRepo.Update(tx, product); err != nil {
				log.Println(err)
				debug.PrintStack()
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
				log.Println(err)
				debug.PrintStack()
				return err
			}

			orderItems = append(orderItems, orderItem)

		}

		// Update total_amount akhir pada table orders
		order.TotalAmount = totalAmount
		order.Items = orderItems

		if err := s.orderRepo.Update(tx, &order); err != nil {
			log.Println(err)
			debug.PrintStack()
			return err
		}

		return nil // Transaksi sukses. GORM melakukan commit
	})

	return &order, errTx
}

func (s *orderServiceImpl) FindById(id int) (*model.Order, error) {
	return s.orderRepo.FindByID(s.db, uint(id), "User", "Items.Product")
}

func (s *orderServiceImpl) Cancel(id int) (*model.Order, error) {
	order, err := s.orderRepo.FindByID(s.db, uint(id), "Items.Product")

	if err != nil {
		return nil, err
	}

	if order.Status == "CANCELLED" {
		return nil, fmt.Errorf("This Order with ID [%d] already CANCELLED", order.ID)
	}
	errTx := s.db.Transaction(func(tx *gorm.DB) error {
		for _, orderItem := range order.Items {
			product, err := s.productRepo.FindByID(tx, orderItem.ProductID)
			if err != nil {
				return err
			}

			product.Stock += orderItem.Quantity
			if err := s.productRepo.Update(tx, product); err != nil {
				return err
			}
		}

		order.Status = "CANCELLED"
		if err := s.orderRepo.Update(tx, order); err != nil {
			return err
		}
		return nil
	})

	return order, errTx

}
