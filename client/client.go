package client

import (
	"time"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type ClientInterFace interface {
	HitungLuasPersegi(panjang int32, lebar int32) int32
	HitungVolumeKubus(sisi int32) int32
	HitungKelilingPersegi(sisi int32) int32
	HitungLingkaran(radius float64) (float64, float64)
	HitungSegitiga(alas int32, tinggi int32, sisi_1 int32, sisi_2 int32, sisi_3 int32) (int32, int32)
	HitungUmur(tanggalLahir string) int32
}

type ClientStruct struct {
	logger interfaces.Logger
}

func NewClientStruct(logger interfaces.Logger) *ClientStruct {
	return &ClientStruct{
		logger: logger,
	}
}

func (c *ClientStruct) HitungLuasPersegi(panjang int32, lebar int32) int32 {
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := panjang * lebar
	c.logger.Info("LUAS", luas)
	return luas
}

func (c *ClientStruct) HitungVolumeKubus(sisi int32) int32 {
	c.logger.Info("PANJANG SISI", sisi)
	volumeKubus := sisi * sisi * sisi
	c.logger.Info("LUAS", volumeKubus)
	return volumeKubus
}

func (c *ClientStruct) HitungKelilingPersegi(sisi int32) int32 {
	c.logger.Info("PANJANG SISI", sisi)
	kelilingPersegi := sisi * sisi * sisi * sisi
	c.logger.Info("KELILING", kelilingPersegi)
	return kelilingPersegi
}

func (c *ClientStruct) HitungLingkaran(radius float64) (float64, float64) {
	c.logger.Info("RADIUS", radius)
	luas := 3.14 * radius * radius
	keliling := 2 * 3.14 * radius
	c.logger.Info("LUAS", luas)
	c.logger.Info("KELILING", keliling)
	return luas, keliling
}

func (c *ClientStruct) HitungSegitiga(alas int32, tinggi int32, sisi_1 int32, sisi_2 int32, sisi_3 int32) (int32, int32) {
	luas := (alas * tinggi) / 2
	keliling := sisi_1 + sisi_2 + sisi_3

	return luas, keliling
}

func (c *ClientStruct) HitungUmur(tanggalLahir string) int32 {
	c.logger.Info("TANGGAL LAHIR", tanggalLahir)
	formatDate := "02-01-2006"
	parsedDate, err := time.Parse(formatDate, tanggalLahir)
	if err != nil {
		c.logger.Error("Error", err)
		return 0
	}
	currentTime := time.Now()
	age := currentTime.Year() - parsedDate.Year()

	if currentTime.Month() < parsedDate.Month() || (currentTime.Month() == parsedDate.Month() && currentTime.Day() < parsedDate.Day()) {
		age--
	}
	c.logger.Info("UMUR", age)
	return int32(age)
}
