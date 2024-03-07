package controller

import (
	"proto-workshop/client"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControllerInterface interface {
	HitungLuasPersegi(panjang int32, lebar int32) (int32, error)
	HitungVolumeKubus(sisi int32) (int32, error)
	HitungKelilingPersegi(sisi int32) (int32, error)
	HitungLingkaran(radius float64) (float64, float64, error)
	HitungSegitiga(alas int32, tinggi int32, sisi_1 int32, sisi_2 int32, sisi_3 int32) (int32, int32, error)
	HitungUmur(tanggalLahir string) (int32, error)
}

type ControllerStruct struct {
	logger interfaces.Logger
	cl     client.ClientInterFace
}

type ControlleStruct struct {
	logger interfaces.Logger
}

func NewController(logger interfaces.Logger, cl client.ClientInterFace) ControllerInterface {
	return &ControllerStruct{
		logger: logger,
		cl:     cl,
	}
}

func (c *ControllerStruct) HitungLuasPersegi(panjang int32, lebar int32) (int32, error) {
	if panjang < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	if lebar < 1 {
		c.logger.Error("Error", "Lebar Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Lebar harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := c.cl.HitungLuasPersegi(panjang, lebar)
	c.logger.Info("LUAS", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungVolumeKubus(sisi int32) (int32, error) {
	if sisi < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", sisi)
	volumeKubus := c.cl.HitungVolumeKubus(sisi)
	c.logger.Info("LUAS", volumeKubus)
	return volumeKubus, nil
}

func (c *ControllerStruct) HitungKelilingPersegi(sisi int32) (int32, error) {
	if sisi < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", sisi)
	kelilingPersegi := c.cl.HitungKelilingPersegi(sisi)
	c.logger.Info("KELILING", kelilingPersegi)
	return kelilingPersegi, nil
}

func (c *ControllerStruct) HitungLingkaran(radius float64) (float64, float64, error) {
	if radius == 0 {
		c.logger.Error("Error", "radius tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "radius tidak boleh kosong")
	}

	luas, keliling := c.cl.HitungLingkaran(radius)
	return luas, keliling, nil
}

func (c *ControllerStruct) HitungSegitiga(alas int32, tinggi int32, sisi_1 int32, sisi_2 int32, sisi_3 int32) (int32, int32, error) {
	if alas == 0 {
		c.logger.Error("Error", "alas tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "alas tidak boleh kosong")
	}
	if tinggi == 0 {
		c.logger.Error("Error", "tinggi tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "tinggi tidak boleh kosong")
	}
	if sisi_1 == 0 {
		c.logger.Error("Error", "sisi tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "sisi tidak boleh kosong")
	}

	luas, keliling := c.cl.HitungSegitiga(alas, tinggi, sisi_1, sisi_2, sisi_3)
	return luas, keliling, nil
}

func (c *ControllerStruct) HitungUmur(tanggalLahir string) (int32, error) {
	c.logger.Info("TANGGAL LAHIR", tanggalLahir)
	umur := c.cl.HitungUmur(tanggalLahir)
	c.logger.Info("UMUR", umur)
	return umur, nil
}
