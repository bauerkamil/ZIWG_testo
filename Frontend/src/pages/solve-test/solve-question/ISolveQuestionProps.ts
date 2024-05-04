import { IAnswear, IQuestion } from "@/shared/interfaces";

export interface ISolveQuestionProps {
  question: IQuestion;
  onNext: (selectedAnswears: IAnswear[]) => void;
  onSkip: () => void;
}
