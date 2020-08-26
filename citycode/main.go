package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	*gorm.DB
}

type City struct {
	CityCode string `json:"city_code"`
	CityName string `json:"city_name"`
}

type Province struct {
	ProvinceName string `json:"province_name"`
	ProvinceCode string `json:"province_code"`
	Citites      []City `json:"cities"`
}

type Areas struct {
	Province     string `json:"province"`
	ProvinceCode string `json:"province_code"`
	City         string `json:"city"`
	CityCode     string `json:"city_code"`
}

func (d *Dao) Get() {
	areas := make([]Areas, 0)

	// 查询省
	if err := d.Find(&areas, "level=1").Error; err != nil {
		fmt.Println(err)
		return
	}

	provinces := []Province{}
	for _, v := range areas {
		// 查询城市
		temp := make([]Areas, 0)
		if err := d.Find(&temp, "province_code=? and county='' and city !=''", v.ProvinceCode).Error; err != nil {
			fmt.Println(err)
			return
		}

		p := Province{ProvinceName: v.Province, ProvinceCode: v.ProvinceCode}
		for _, vv := range temp {
			p.Citites = append(p.Citites, City{CityName: vv.City, CityCode: vv.CityCode})
		}
		provinces = append(provinces, p)
	}

	res, _ := json.Marshal(provinces)
	fmt.Println(string(res))
}

func initDao() Dao {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"appback", "123456", "192.168.1.124:3306", "youbei_ams_test"))
	if err != nil {
		panic(err)
	}
	return Dao{db}
}

func main() {
	dao := initDao()
	dao.Get()
}
