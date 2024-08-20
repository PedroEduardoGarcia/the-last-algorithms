import unittest
from typing import Any
from code import linear_search

class TestLinearSerch(unittest.TestCase):
    def test_linear_search(self):
        arr: Any = [10, "hello", 12, 45.23, 111]
        self.assertTrue(linear_search(arr, "hello"))
        self.assertFalse(linear_search(arr, 11))
        self.assertFalse(linear_search(arr, "not_in_array"))
        self.assertTrue(linear_search(arr, 45.23))

if __name__ == '__main__':
    unittest.main()
