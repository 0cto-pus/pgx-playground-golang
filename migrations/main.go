package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    // Migration dosyalarının yolunu ve PostgreSQL bağlantı bilgisini veriyoruz.
    m, err := migrate.New(
        "file://./", // Migrations klasörünün yolu
        "postgres://postgres:postgres@localhost:6432/productapp?sslmode=disable") // Veritabanı bağlantısı
    if err != nil {
        log.Fatalf("Migration başlatılamadı: %v", err)
    }

    // Migration'ları çalıştırmak için `Up` metodu kullanılıyor. Bu tüm migration dosyalarını çalıştırır.
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Migration hatası: %v", err)
    }

    log.Println("Migration işlemi başarıyla tamamlandı!")
}