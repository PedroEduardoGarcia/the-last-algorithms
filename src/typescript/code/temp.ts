export const linear_search = (arr: any[], v: any): boolean => {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === v) return true;
    }
    return false;
};