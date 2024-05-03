import { IAnswerResponse } from "./IAnswerResponse";

export interface IQuestionResponse {
    id: string,
    body: string,
    img_file: string,
    answers: IAnswerResponse[],
    test_id: string
}