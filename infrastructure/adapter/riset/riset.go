package riset

import (
	"fmt"
	"go-hexagonal/core/model/mpay"
	port "go-hexagonal/core/port/riset"
	"go-hexagonal/utils/config"

	// "gorm.io/gorm"
	"go-hexagonal/interface/api/extl/v1/riset/request"
	"go-hexagonal/utils/paginator"
)

type RisetAdapter struct {
	mpayUrl string
}

func NewRisetAdapter() port.RisetAdapter {
	return &RisetAdapter{
		mpayUrl: "test",
	}
}

func (adapter *RisetAdapter) GetMpayCustomers(params request.MpayCustReq) (*[]mpay.MpayCustomer, int64, error) {
	paginate := paginator.Paginate(params.Page, params.Limit)
	var mpayCustomer []mpay.MpayCustomer
	var mpayCustomerCount []mpay.MpayCustomer
	var count int64

	err := config.Db.Unscoped().Scopes(paginate).Table("mpay_customer").Find(&mpayCustomer).Error
	if err != nil {

	}

	// count for pagination
	config.Db.Unscoped().Table("mpay_customer").Find(&mpayCustomerCount).Count(&count)

	fmt.Println(count)
	return &mpayCustomer, count, nil
}

func (adapter *RisetAdapter) GetMpayCustomer(id int) (*mpay.MpayCustomer, error) {
	var mpayCustomer mpay.MpayCustomer
	err := config.Db.Unscoped().Table("mpay_customer").Where("id = ?", id).First(&mpayCustomer).Error
	fmt.Println()
	if err != nil {
		return &mpayCustomer, err
	}

	return &mpayCustomer, nil
}

func (adapter *RisetAdapter) GetMpayCustomerByPhone(phone string) (*mpay.MpayCustomer, error) {
	var mpayCustomer mpay.MpayCustomer
	err := config.Db.Unscoped().Table("mpay_customer").Where("phone_number = ?", phone).First(&mpayCustomer).Error
	fmt.Println()
	if err != nil {
		return &mpayCustomer, err
	}

	return &mpayCustomer, nil
}

func (adapter *RisetAdapter) CreateMpayCustomer(payload request.MpayCustPayload) error {
	fmt.Println()
	err := config.Db.Table("mpay_customer").Create(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (adapter *RisetAdapter) UpdateMpayCustomer(id int, payload request.MpayCustPayload) error {
	fmt.Println()
	err := config.Db.Table("mpay_customer").Where("id = ?", id).Updates(payload)
	fmt.Println(err)
	if err != nil {
		return err.Error
	}
	return nil
}

func (adapter *RisetAdapter) DeleteMpayCustomer(id int) error {
	var mpayCust mpay.MpayCustomer
	err := config.Db.Table("mpay_customer").Where("id = ?", id).Delete(&mpayCust)

	if err != nil {
		return err.Error
	}
	return nil
}
