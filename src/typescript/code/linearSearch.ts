/**
 * Performs a linear search on a list to find a value.
 * Returns true if the value is found in the array arr, otherwise false
 *
 * @param arr - The array to search through.
 * @param v - The value to search for in the array.
 * @returns True if the value is found in the array, false otherwise.
 */
export function linearSearch<T>(arr: T[], v: T): boolean {
    for (const e of arr) {
        if (e === v) {
            return true;
        }
    }
    return false;
}