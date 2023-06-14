import unittest
from maintenance_window import parse_maintenance_window
from datetime import time


class TestMaintenanceWindow(unittest.TestCase):
    def test_maintenance_window_to_start_end_time(self):
        self.assertEqual(parse_maintenance_window(
            'test_db', 'sun:01:00-sun:02:00'),
            (time(hour=1), time(hour=2))
        )

        self.assertEqual(parse_maintenance_window(
            'test_db', 'sun:05:00-sun:08:00'), (time(hour=5), time(hour=8)))

        self.assertEqual(
            parse_maintenance_window('test_db', 'wed:23:15-wed:23:45'),
            (time(hour=23, minute=15), time(hour=23, minute=45))
        )

        self.assertRaises(
            AssertionError, parse_maintenance_window, 'wed:23:15-thu:23:45')


if __name__ == '__main__':
    unittest.main()
