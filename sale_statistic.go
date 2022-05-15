package rdg

import (
	"database/sql"

	"ondrejmaksi.com/db2project/db"
)

/*
Detailnosť popisu produktu je počet atribútov, ktorý má produkt v systéme vyplnený. Pre každú detailnosť popisu produktu vypočítajte počet predaných kusov produktov s danou detailnosťou produktu.
*/
type SaleStatistic struct {
	AttributeCount int
	TotalSold      int
}

func rowsToSaleStatistics(rows *sql.Rows) []SaleStatistic {
	res := []SaleStatistic{}
	for rows.Next() {
		ss := SaleStatistic{}
		err := rows.Scan(&ss.AttributeCount, &ss.TotalSold)
		if err != nil {
			panic(err)
		}
		res = append(res, ss)
	}

	return res
}

func GetSaleStatistics() []SaleStatistic {
	sql := `with attribute_count as (
		select 
			p.id, count(av.id) as a_count
		from 
			products p
		left join 
			product_attributes pa on p.id = pa.product_id
		left join 
			attribute_values av on av.id = pa.attribute_value_id
		group by 
			p.id
		order by 
			p.id
	), sold_count as (
		select 
			p.id, coalesce(sum(oi.quantity), 0) as s_count
		from 
			products p
		join 
			order_items oi on oi.product_id = p.id
		join 
			orders o on oi.order_id = o.id
		where 
			o.status = 'pending'
		group by 
			p.id
		order by 
			p.id
	)
	select  ac.a_count, sum(sc.s_count) as s_count_total
	from attribute_count ac
	left join sold_count sc on ac.id = sc.id 
	group by ac.a_count
	order by ac.a_count
	`
	rows, err := db.GetDatabase().Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToSaleStatistics(rows)
}
