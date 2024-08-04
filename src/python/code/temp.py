from typing import List

def linear_search(arr: List[any], v: any) -> bool:
    for e in arr:
        if e == v:
            return True
    return False