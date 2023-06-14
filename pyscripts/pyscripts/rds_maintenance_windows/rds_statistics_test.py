import unittest
from rds_statistics import get_maintenance_window_cpu_percentile_ranking
from maintenance_window import DB_Cluster, parse_maintenance_window
from datetime import datetime, time


class TestMaintenanceWindow(unittest.TestCase):
    def test_get_maintenance_window_cpu_percentile_ranking(self):
        date_time_list = [
            datetime(2021, 10, 10, 0, 0, 0),
            datetime(2021, 10, 10, 1, 0, 0),
            datetime(2021, 10, 10, 2, 0, 0),
            datetime(2021, 10, 10, 3, 0, 0),
            datetime(2021, 10, 10, 4, 0, 0),
            datetime(2021, 10, 10, 5, 0, 0),
            datetime(2021, 10, 10, 6, 0, 0),
            datetime(2021, 10, 10, 7, 0, 0),
            datetime(2021, 10, 10, 8, 0, 0),
            datetime(2021, 10, 10, 9, 0, 0),
            datetime(2021, 10, 10, 10, 0, 0),
            datetime(2021, 10, 10, 11, 0, 0),
            datetime(2021, 10, 10, 12, 0, 0),
            datetime(2021, 10, 10, 13, 0, 0),
            datetime(2021, 10, 10, 14, 0, 0),
        ]

        self.assertEqual(
            get_maintenance_window_cpu_percentile_ranking(
                DB_Cluster('test_db', parse_maintenance_window(
                    'test_db', 'sun:01:00-sun:02:00')),
                date_time_list),
            [2]
        )

        self.assertEqual(
            get_maintenance_window_cpu_percentile_ranking(
                DB_Cluster('test_db', parse_maintenance_window(
                    'test_db', 'sun:01:00-sun:03:00')),
                date_time_list),
            [2, 3]
        )
        self.assertEqual(
            get_maintenance_window_cpu_percentile_ranking(
                DB_Cluster('test_db', parse_maintenance_window(
                    'test_db', 'sun:01:00-sun:03:15')),
                date_time_list),
            [2, 3]
        )

        self.assertTrue(datetime(2021, 10, 10, 1, 0, 0).time() < time(hour=2))
        self.assertTrue(datetime(2021, 10, 10, 1, 30, 0).time() < time(hour=2))


if __name__ == '__main__':
    unittest.main()
