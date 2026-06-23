package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// concert gorm postgres
type Concert struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Date        time.Time `gorm:"not null" json:"date"`
	Venue       string    `gorm:"not null" json:"venue"`
	Status      string    `gorm:"default:active" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   time.Time `gorm:"autoDeleteTime"`
}

// config database
var db *gorm.DB

func initDB() {
	var err error
	// TODO : ganti ke postgres
	dsn := "postgres://postgres:secret45@localhost:5434/eticketdb?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// auto migrate
	db.AutoMigrate(&Concert{})

	log.Println("Database connected")

	// seeder 5 data awal
	var concerts []Concert
	db.Find(&concerts)
	if len(concerts) == 0 {
		db.Create(&[]Concert{
			{
				Title:       "Coldplay Music of the Spheres World Tour Jakarta",
				Description: "Konser perdana band asal Inggris, Coldplay, di Indonesia yang memukau ratusan ribu penonton dengan gelang Xyloband yang menyala warna-warni.",
				Date:        time.Date(2023, 11, 15, 20, 0, 0, 0, time.UTC),
				Venue:       "Stadion Utama Gelora Bung Karno (SUGBK), Jakarta",
				Status:      "completed",
			},
			{
				Title:       "Blackpink [Born Pink] World Tour Jakarta",
				Description: "Konser megah dari girlgroup K-Pop fenomenal, Blackpink, yang berhasil meremajakan Jakarta menjadi lautan cahaya merah muda selama dua hari berturut-turut.",
				Date:        time.Date(2023, 3, 11, 19, 0, 0, 0, time.UTC),
				Venue:       "Stadion Utama Gelora Bung Karno (SUGBK), Jakarta",
				Status:      "completed",
			},
			{
				Title:       "Metallica Live in Jakarta 2013",
				Description: "Konser sejarah kembalinya raja thrash metal dunia ke Indonesia setelah penantian 20 tahun, dihadiri oleh puluhan ribu pecinta musik cadas dari berbagai generasi.",
				Date:        time.Date(2013, 8, 25, 20, 0, 0, 0, time.UTC),
				Venue:       "Stadion Utama Gelora Bung Karno (SUGBK), Jakarta",
				Status:      "completed",
			},
			{
				Title:       "Bruno Mars Live in Jakarta 2026",
				Description: "Konser tur dunia dari solois legendaris Bruno Mars yang membawakan deretan lagu hitsnya dengan koreografi dan vokal yang sangat enerjik.",
				Date:        time.Date(2026, 6, 22, 20, 0, 0, 0, time.UTC),
				Venue:       "Jakarta International Stadium (JIS), Jakarta",
				Status:      "active",
			},
			{
				Title:       "Pesta Rakyat Dewa 19 - 30 Tahun Berkarya",
				Description: "Konser selebrasi 3 dekade salah satu band rock terbesar di Indonesia, Dewa 19, yang memboyong 4 vokalis dan 5 drummer dalam satu panggung.",
				Date:        time.Date(2026, 6, 22, 19, 30, 0, 0, time.UTC),
				Venue:       "Stadion Utama Gelora Bung Karno (SUGBK), Jakarta",
				Status:      "active",
			},
		})
	}
}

// middleware untuk validasi x-api-key
func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		if apiKey != "juara-coding-super-secret" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// simple rate limiter (limit 5 request per second per user)
type IPRateLimiter struct {
	mu   sync.Mutex
	hits map[string]int
}

var limiter = IPRateLimiter{
	hits: make(map[string]int),
}

func RateLimiter(maxRequest int) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		limiter.mu.Lock()

		currentHits := limiter.hits[clientIP]

		if currentHits >= maxRequest {
			limiter.mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		limiter.hits[clientIP] = currentHits + 1
		limiter.mu.Unlock()

		c.Next()
	}
}

func main() {
	initDB()

	//inisiasi gin
	r := gin.Default()

	// global middleware
	r.Use(ApiKeyAuth())
	r.Use(RateLimiter(5))

	r.GET("/api/v1/concerts", func(c *gin.Context) {
		var concerts []Concert
		db.Find(&concerts)
		c.JSON(http.StatusOK, concerts)
	})

	r.Run(":8080")

}
