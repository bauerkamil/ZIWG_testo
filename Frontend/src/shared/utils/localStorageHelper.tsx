export const getStoredValue = <T>(keyName: string): T | undefined => {
  try {
    const value = window.localStorage.getItem(keyName);
    if (value) {
      return JSON.parse(value);
    } else {
      return undefined;
    }
  } catch (err) {
    console.log(err);
    return undefined;
  }
};

export const setStoredValue = <T>(keyName: string, newValue: T): void => {
  try {
    window.localStorage.setItem(keyName, JSON.stringify(newValue));
  } catch (err) {
    console.log(err);
  }
};
export const removeStoredItem = (keyName: string): void => {
  try {
    window.localStorage.removeItem(keyName);
  } catch (err) {
    console.log(err);
  }
};
