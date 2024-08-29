package adapter

import (
	"context"
	"fmt"
	"github.com/Honeymoond24/tender-analysis/internal/application"
)

type StatisticsRepository struct {
	DbPool *DBPool
}

func NewStatisticsRepository(db *DBPool) application.StatisticsRepository {
	return &StatisticsRepository{DbPool: db}
}

// MostActiveCategoryByTenders returns the category with the most active tenders
// and the number of active tenders in this category
func (s *StatisticsRepository) MostActiveCategoryByTenders() (category string, count int) {
	row := s.DbPool.conn.QueryRow(context.Background(),
		`SELECT count(a.id) AS count, l.ktru_name
		FROM announces AS a
		         JOIN lot_announces AS l ON a.id = l.announce_id
		WHERE COALESCE(a.open_start_date, a.offers_start_date) > now()
		GROUP BY l.ktru_name
		ORDER BY count DESC
		LIMIT 1`)
	err := row.Scan(&count, &category)
	if err != nil {
		return
	}
	return
}

func (s *StatisticsRepository) MostActiveCategoryByPriceSum() (category string, sum float64) {
	row := s.DbPool.conn.QueryRow(context.Background(),
		`SELECT SUM(l.planned_amount) AS tender_sum, l.ktru_name
			FROM lot_announces AS l
					 LEFT JOIN announces AS a ON a.id = l.announce_id
			WHERE (a.open_start_date IS NOT NULL AND a.open_start_date > now())
			   OR (a.open_start_date IS NULL AND a.offers_start_date > now())
				AND l.planned_amount IS NOT NULL
			GROUP BY l.ktru_name
			ORDER BY tender_sum DESC
			LIMIT 1;`)
	err := row.Scan(&sum, &category)
	fmt.Println(sum, category)
	if err != nil {
		return
	}
	return
}

func (s *StatisticsRepository) ActiveTenders() (count int64) {
	row := s.DbPool.conn.QueryRow(context.Background(),
		`SELECT count(1) FROM announces AS a
            WHERE COALESCE(a.open_start_date, a.offers_start_date) > now()`)
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func (s *StatisticsRepository) CategorySumsCounts() (result [][]interface{}) {
	rows, err := s.DbPool.conn.Query(context.Background(),
		`SELECT l.ktru_name AS category,
		       SUM(a.sum)  AS category_sum,
		       COUNT(a.id) AS tenders_count
		FROM announces AS a
		         JOIN lot_announces AS l ON a.id = l.announce_id
		GROUP BY l.ktru_name
		ORDER BY tenders_count DESC;`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var category string
		var sum float64
		var count int
		err = rows.Scan(&category, &sum, &count)
		if err != nil {
			return
		}
		result = append(result, []interface{}{category, sum, count})
	}
	fmt.Println("result", result)
	return
}

func (s *StatisticsRepository) MonthsWithMoreTendersThanAverage() (result [][]int) {
	rows, err := s.DbPool.conn.Query(context.Background(),
		`WITH YearMonths AS (SELECT count(a.id)                                                          AS tenders_count,
                           EXTRACT(YEAR FROM COALESCE(a.offers_start_date, a.open_start_date))  AS year,
                           EXTRACT(MONTH FROM COALESCE(a.offers_start_date, a.open_start_date)) AS month
                    FROM announces AS a
                    WHERE COALESCE(a.offers_start_date, a.open_start_date) IS NOT NULL
                    GROUP BY year, month
                    ORDER BY year, month),
     AveragePerMonth AS (SELECT AVG(ym.tenders_count)
                         FROM YearMonths AS ym)
		SELECT ym.tenders_count,
			   ym.year,
			   ym.month
		FROM YearMonths AS ym
		WHERE ym.tenders_count > (SELECT * FROM AveragePerMonth) * 1.1;`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tendersCount int
		var year int
		var month int
		err = rows.Scan(&tendersCount, &year, &month)
		if err != nil {
			return
		}
		result = append(result, []int{tendersCount, year, month})
	}
	fmt.Println("result", result)
	return
}
