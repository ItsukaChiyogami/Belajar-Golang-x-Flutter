package models

import (
	"golang/db"
	"net/http"
)

type Barang struct {
	Id         int    `json:id`
	KodeBarang string `json:kodebarang`
	NamaBarang string `json:namabarang`
}

func DataBarang() (Response, error) {
	var obj Barang
	var arrobj []Barang
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tabelbarang"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.KodeBarang, &obj.NamaBarang)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}
