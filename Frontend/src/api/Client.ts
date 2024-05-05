import { getStoredValue } from "@/shared/utils/localStorageHelper";
import axios from "axios";
import { DEFAULT_EP } from "@/shared/utils/constants";
import { ITestsResponse } from "@/shared/interfaces/ITestsResponse";
import { IUser } from "@/shared/interfaces";
import { ITestResponse } from "@/shared/interfaces/ITestResponse";
import { LocalStorageElements } from "@/shared/enums";

axios.interceptors.request.use((setup) => {
  const user = getStoredValue<IUser>(LocalStorageElements.User);

  if (user) {
    if (user.token) {
      setup.headers.Authorization = `Bearer ${user.token}`;
    }
  }

  return setup;
});

const Client = {
  getTests: async () =>
    axios
      .get<ITestsResponse[]>(`${DEFAULT_EP}/test`)
      .then((response) => response.data),
  getTest: async (test_id: string) =>
    axios
      .get<ITestResponse>(`${DEFAULT_EP}/test/${test_id}`)
      .then((response) => response.data),
};

export default Client;
