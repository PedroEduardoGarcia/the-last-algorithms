import { linearSearch } from '../code/linearSearch';

test('linear_search finds the value in the array', () => {
    const arr = [10, "hello", 12, 45.23, 111];
    expect(linearSearch(arr, "hello")).toEqual(true);
    expect(linearSearch(arr, 11)).toEqual(false);
    expect(linearSearch(arr, "not_in_array")).toEqual(false);
    expect(linearSearch(arr, 45.23)).toEqual(true);
});