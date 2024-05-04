import React from "react";
import Navbar from "@/components/navbar/Navbar";
import { IAnswear, IQuestion, ITest } from "@/shared/interfaces";
import { shuffle } from "@/shared/utils/helpers";
import SolveQuestion from "./solve-question/SolveQuestion";
import { useParams } from "react-router-dom";

const SolveTest: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  const [test, setTest] = React.useState<ITest>();
  const [questionsToSolve, setQuestionsToSolve] = React.useState<IQuestion[]>(
    []
  );
  const [currentQuestion, setCurrentQuestion] =
    React.useState<IQuestion | null>();

  React.useEffect(() => {
    const mockQuestions: IQuestion[] = [
      {
        id: "1",
        body: "Czy 2+2=4?",
        imgFile: "",
        testId: "1",
        answears: [
          { id: "1", body: "Tak", valid: true, questionId: "1" },
          { id: "2", body: "Nie", valid: false, questionId: "1" },
        ],
      },
      {
        id: "2",
        body: "Czy pwr jest jebniety?",
        imgFile: "",
        testId: "1",
        answears: [
          { id: "3", body: "Tak", valid: true, questionId: "2" },
          { id: "4", body: "Nie", valid: false, questionId: "2" },
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
    console.log(id);
  }, [id]);

  React.useEffect(() => {
    setCurrentQuestion(
      questionsToSolve && questionsToSolve.length > 0
        ? { ...questionsToSolve[0] }
        : null
    );
  }, [questionsToSolve]);

  const handleNext = (selectedAnswears: IAnswear[]) => {
    const currentAnswears = currentQuestion?.answears;
    const correct = currentAnswears?.every((answear) => {
      if (answear.valid) {
        return selectedAnswears.some((selected) => selected.id === answear.id);
      } else {
        return !selectedAnswears.some((selected) => selected.id === answear.id);
      }
    });
    if (correct) {
      setQuestionsToSolve((prev) => prev.slice(1));
    } else {
      setQuestionsToSolve((prev) => [...prev.slice(1), prev[0]]);
    }
  };

  const handleSkip = () => {
    setQuestionsToSolve((prev) => prev.slice(1));
  };

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
        {currentQuestion && (
          <SolveQuestion
            question={currentQuestion}
            onNext={handleNext}
            onSkip={handleSkip}
          />
        )}
      </main>
    </div>
  );
};

export default SolveTest;
