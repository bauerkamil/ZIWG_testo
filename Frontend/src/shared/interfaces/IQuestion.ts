import { IAnswear } from "./IAnswear";

export interface IQuesiton {
  id: string;
  body: string;
  imgFile: string;
  testId: string;
  answears: IAnswear[];
}
