import { getStoredValue } from "@/shared/utils/localStorageHelper";
import axios from "axios";
import { DEFAULT_EP } from "@/shared/utils/constants";
import { ITestsResponse } from "@/shared/interfaces/ITestsResponse";
import { IUser } from "@/shared/interfaces";
import { ITestResponse } from "@/shared/interfaces/ITestResponse";


axios.interceptors.request.use((setup) => {
    const userString = getStoredValue<IUser>("user")

    if (userString) {
        if (userString.token) {
            setup.headers.Authorization = `Bearer ${userString.token}`;
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