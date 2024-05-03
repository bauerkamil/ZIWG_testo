import React from "react";
import Navbar from "@/components/navbar/Navbar";
import { IQuestion, ITest } from "@/shared/interfaces";
import { shuffle } from "@/shared/utils/helpers";
import SolveQuestion from "./solve-question/SolveQuestion";

const SolveTest: React.FC = () => {
  const [test, setTest] = React.useState<ITest>();
  const [questionsToSolve, setQuestionsToSolve] = React.useState<IQuestion[]>(
    []
  );

  React.useEffect(() => {
    const mockQuestions: IQuestion[] = [
      {
        id: "1",
        body: "Czy 2+2=4?",
        imgFile: "",
        testId: "1",
        answears: [
          { id: "1", body: "Tak", isCorrect: true, questionId: "1" },
          { id: "2", body: "Nie", isCorrect: false, questionId: "1" },
        ],
      },
      {
        id: "2",
        body: "Czy 2+2=5?",
        imgFile: "",
        testId: "1",
        answears: [
          { id: "3", body: "Tak", isCorrect: false, questionId: "2" },
          { id: "4", body: "Nie", isCorrect: true, questionId: "2" },
        ],
      },
    ];
    const mockTest: ITest = {
      id: "123",
      name: "Testownik 1",
      course: "Analiza",
      teacher: "mgr. Jan Dupa",
      lastModified: new Date(),
      questions: mockQuestions,
    };
    setTest(mockTest);
    setQuestionsToSolve(shuffle(mockQuestions));
  }, []);

  return (
    <div className="flex min-h-screen w-full flex-col">
      <Navbar />
      <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <div className="flex items-center justify-center">
          <div className="flex flex-row gap-8  w-[550px] justify-between">
            <div className="font-semibold leading-none tracking-tight text-2xl ">
              {test?.name}
            </div>
            <div className="text-primary font-bold leading-none text-2xl">
              {test?.course}
            </div>
          </div>
        </div>
        {questionsToSolve && questionsToSolve.length > 0 ? (
          <SolveQuestion question={questionsToSolve[0]} />
        ) : (
          <></>
        )}
      </main>
    </div>
  );
};

export default SolveTest;
