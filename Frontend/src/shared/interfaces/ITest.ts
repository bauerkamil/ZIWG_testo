import { IQuestion } from "./IQuestion";

export interface ITest {
  id: string;
  name: string;
  course: string;
  teacher: string;
  lastModified: Date;
  questions?: IQuestion[];
}
