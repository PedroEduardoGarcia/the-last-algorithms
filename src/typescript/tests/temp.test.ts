import { linear_search } from '../code/temp';

test('linear_search finds the value in the array', () => {
    const arr = [10, "hello", 12, 45.23, 111];
    expect(linear_search(arr, "hello")).toEqual(true);
    expect(linear_search(arr, 45.23)).toEqual(true);
    expect(linear_search(arr, 111)).toEqual(true);
    expect(linear_search(arr, "not_in_array")).toEqual(false);
});