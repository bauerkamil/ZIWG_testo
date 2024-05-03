import { IQuestionResponse } from "./IQuestionResponse";

export interface ITestResponse {
    id: string,
    name: string,
    questions: IQuestionResponse[]
}