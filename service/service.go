package service

import (
	"context"
	"proto-workshop/controller"
	"proto-workshop/generated/proto"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type ServiceStruct struct {
	logger interfaces.Logger
	ctr    controller.ControllerInterface
}

func NewService(logger interfaces.Logger, ctr controller.ControllerInterface) *ServiceStruct {
	return &ServiceStruct{
		logger: logger,
		ctr:    ctr,
	}
}

func (s *ServiceStruct) HitungLuasPersegi(ctx context.Context, req *proto.RequestHitungLuasPersegi) (*proto.ResponseGlobal, error) {
	data, err := s.ctr.HitungLuasPersegi(req.GetPanjang(), req.GetLebar())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseGlobal{
		Response: data,
	}, nil
}

func (s *ServiceStruct) HitungVolumeKubus(ctx context.Context, req *proto.RequestGlobal) (*proto.ResponseGlobal, error) {
	data, err := s.ctr.HitungVolumeKubus(req.GetRequest())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseGlobal{
		Response: data,
	}, nil
}

func (s *ServiceStruct) HitungKelilingPersegi(ctx context.Context, req *proto.RequestGlobal) (*proto.ResponseGlobal, error) {
	data, err := s.ctr.HitungKelilingPersegi(req.GetRequest())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseGlobal{
		Response: data,
	}, nil
}

func (s *ServiceStruct) HitungLingkaran(ctx context.Context, req *proto.RequestLingkaran) (*proto.ResponseLingkaran, error) {
	luas, keliling, err := s.ctr.HitungLingkaran(float64(req.GetRadius()))

	if err != nil {
		return nil, err
	}

	return &proto.ResponseLingkaran{Luas: float32(luas), Keliling: float32(keliling)}, nil
}

func (s *ServiceStruct) HitungSegitiga(ctx context.Context, req *proto.RequestSegitiga) (*proto.ResponseSegitiga, error) {
	luas, keliling, err := s.ctr.HitungSegitiga(req.Alas, req.Tinggi, req.Sisi_1, req.Sisi_2, req.Sisi_3)

	if err != nil {
		return nil, err
	}

	return &proto.ResponseSegitiga{Luas: luas, Keliling: keliling}, nil
}

func (s *ServiceStruct) HitungUmur(ctx context.Context, req *proto.RequestUmur) (*proto.ResponseGlobal, error) {
	data, err := s.ctr.HitungUmur(req.GetTanggalLahir())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseGlobal{
		Response: data,
	}, nil
}
