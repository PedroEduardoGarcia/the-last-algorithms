from typing import List

def linear_search(arr: List[any], v: any) -> bool:
    """Performs a linear search on a list to find a value.
    Returns true if the value is found in the array arr, otherwise false

    Args:
        arr (List[any]): The list to search through.
        v (any): The value to search for in the array.

    Returns:
        bool: True if the value is found in the array, false otherwise.
    """
    for e in arr:
        if e == v:
            return True
    return False